package services

import (
	"errors"
	"github.com/rizkycahyono97/online-shop-api/model/domain"
	"github.com/rizkycahyono97/online-shop-api/model/web"
	"github.com/rizkycahyono97/online-shop-api/repositories"
	"time"
)

type OrderServiceImpl struct {
	orderRepo   repositories.OrderRepository
	cartRepo    repositories.CartRepository
	productRepo repositories.ProductRepository
}

func NewOrderService(
	orderRepo repositories.OrderRepository,
	cartRepo repositories.CartRepository,
	productRepo repositories.ProductRepository,
) OrderService {
	return &OrderServiceImpl{
		orderRepo:   orderRepo,
		cartRepo:    cartRepo,
		productRepo: productRepo,
	}
}

func (s OrderServiceImpl) CreateOrder(userID uint, req *web.CreateOrderRequest) (*domain.Order, error) {
	// ambil cart dan item dari user
	cartItems, err := s.cartRepo.GetItemsCartByUserID(userID)
	if err != nil || len(cartItems) == 0 {
		return nil, errors.New("cart is empty or failed to fetch cart items")
	}

	// hitung total amount dan persiapkan order_items
	var totalAmount float64
	var orderItems []domain.OrderItem
	for _, item := range cartItems {
		product := item.Product

		// validasi kembali stock
		if product.Stock < item.Quantity {
			return nil, errors.New("product stock insufficient")
		}

		// update stock product
		product.Stock -= item.Quantity
		if _, err := s.productRepo.UpdateProduct(&product); err != nil {
			return nil, errors.New("failed to update product")
		}

		orderItems = append(orderItems, domain.OrderItem{
			ProductID: product.ID,
			Quantity:  item.Quantity,
			Price:     float32(product.Price),
		})
		totalAmount += product.Price * float64(item.Quantity)
	}

	// buat order
	order := &domain.Order{
		UserID:      userID,
		TotalAmount: totalAmount,
		Status:      "pending",
		Address:     req.Address,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		OrderItems:  orderItems,
	}

	createOrder, err := s.orderRepo.CreateOrder(order)
	if err != nil {
		return nil, errors.New("failed to create order")
	}

	orderWithItems, err := s.orderRepo.GetOrderByID(createOrder.ID)
	if err != nil {
		return nil, errors.New("Failed to fetch created order")
	}

	// Hapus cart setelah checkout
	if err := s.cartRepo.ClearCart(userID); err != nil {
		return nil, err
	}

	return orderWithItems, nil
}

func (s OrderServiceImpl) GetOrderByID(orderID uint) (*domain.Order, error) {
	return s.orderRepo.GetOrderByID(orderID)
}

func (s OrderServiceImpl) GetOrderByUserID(userID uint) ([]*domain.Order, error) {
	return s.orderRepo.GetOrderByUserID(userID)
}

func (s OrderServiceImpl) GetAllOrders() ([]*domain.Order, error) {
	return s.orderRepo.GetAllOrders()
}

func (s OrderServiceImpl) CancelOrder(UserID uint) (*domain.Order, error) {
	order, err := s.orderRepo.GetOrderByID(UserID)
	if err != nil {
		return nil, errors.New("order not found")
	}

	if order.UserID != UserID {
		return nil, errors.New("unauthorized to cancel this order")
	}

	if order.Status != "pending" {
		return nil, errors.New("only pending orders can be cancelled")
	}

	order.Status = "cancelled"
	order.UpdatedAt = time.Now()

	return s.orderRepo.UpdateOrder(order)
}
