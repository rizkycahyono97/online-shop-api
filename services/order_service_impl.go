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

	orderWithItems, err := s.orderRepo.GetOrderByID(createOrder.ID, userID)
	if err != nil {
		return nil, errors.New("Failed to fetch created order")
	}

	// Hapus cart setelah checkout
	if err := s.cartRepo.ClearCart(userID); err != nil {
		return nil, err
	}

	return orderWithItems, nil
}

func (s OrderServiceImpl) GetOrderByID(orderID, userID uint) (*web.OrderResponse, error) {
	order, err := s.orderRepo.GetOrderByID(orderID, userID)
	if err != nil {
		return nil, errors.New("order not found or you are not authorized")
	}

	return web.OrderResponseFromModels(order), nil
}

func (s OrderServiceImpl) GetOrderByUserID(userID uint) ([]*domain.Order, error) {
	return s.orderRepo.GetOrderByUserID(userID)
}

func (s OrderServiceImpl) GetAllOrders() ([]*domain.Order, error) {
	return s.orderRepo.GetAllOrders()
}

func (s *OrderServiceImpl) CancelOrder(userID, orderID uint) error {
	order, err := s.orderRepo.GetOrderByID(orderID, userID)
	if err != nil {
		return errors.New("order not found")
	}
	if order.UserID != userID {
		return errors.New("unauthorized")
	}
	if order.Status != "pending" {
		return errors.New("order cannot be cancelled")
	}
	return s.orderRepo.UpdateOrderStatus(orderID, "cancelled")
}

func (s *OrderServiceImpl) ConfirmOrder(userID uint, orderID uint) error {
	order, err := s.orderRepo.GetOrderByID(orderID, userID)
	if err != nil {
		return errors.New("Order not found")
	}
	if order.UserID != userID {
		return errors.New("unauthorized")
	}
	if order.Status != "shipped" {
		return errors.New("order is still in shipping")
	}

	return s.orderRepo.UpdateOrderStatus(orderID, "paid")
}

func (s *OrderServiceImpl) UpdateOrderStatus(orderID uint, status string) (*domain.Order, error) {
	order, err := s.orderRepo.GetOrderByIDOnly(orderID)
	if err != nil {
		return nil, err
	}

	if order.Status == "cancelled" {
		return nil, errors.New("cannot update, order is cancelled")
	}

	order.Status = status

	if _, err := s.orderRepo.UpdateOrder(order); err != nil {
		return nil, errors.New("failed to update order status")
	}

	return order, nil
}
