package services

import (
	"errors"
	"github.com/go-playground/validator/v10"
	"github.com/rizkycahyono97/online-shop-api/model/domain"
	"github.com/rizkycahyono97/online-shop-api/repositories"
)

type CartServiceImpl struct {
	cartRepo    repositories.CartRepository
	productRepo repositories.ProductRepository
	validate    *validator.Validate
}

func NewCartService(cartRepo repositories.CartRepository, productRepo repositories.ProductRepository) CartService {
	return &CartServiceImpl{
		cartRepo:    cartRepo,
		productRepo: productRepo,
	}
}

// Get items in cart by user ID
func (s CartServiceImpl) GetItemsCartByUserID(userID uint) ([]*domain.CartItems, error) {
	items, err := s.cartRepo.GetItemsCartByUserID(userID)
	if err != nil {
		return nil, err
	}
	return items, nil
}

// Get All Cart with Item -> mengemnalikan semua carts dan juga semua users
func (s CartServiceImpl) GetAllCartsWithItems() ([]*domain.Cart, error) {
	return s.cartRepo.GetAllCartsWithItems()
}

// Add item to cart with validation
func (s CartServiceImpl) AddItemToCart(userID uint, productID uint, quantity uint) (*domain.CartItems, error) {
	if quantity == 0 {
		return nil, errors.New("quantity must be greater than zero")
	}

	// pastikan product ada dalam stock
	product, err := s.productRepo.FindProductByID(productID)
	if err != nil {
		return nil, errors.New("product not found")
	}

	// validasi stock
	if product.Stock < quantity {
		return nil, errors.New("insufficient product stock")
	}

	// cari cart berdasarkan user
	cart, err := s.cartRepo.GetCartByUserID(userID)
	if err != nil {
		return nil, err
	}

	// tambahkan item
	item, err := s.cartRepo.AddItemToCart(cart.ID, productID, quantity)
	if err != nil {
		return nil, err
	}

	// update stock di product
	product.Stock -= quantity
	if _, err := s.productRepo.UpdateProduct(product); err != nil {
		return nil, errors.New("failed to update product stock")
	}

	return item, nil
}

// Delete item from cart
func (s CartServiceImpl) RemoveItemFromCart(userID uint, productID uint, quantity uint) error {
	// cari berdasarkan user
	cart, err := s.cartRepo.GetCartByUserID(userID)
	if err != nil {
		return errors.New("Cart Not Found")
	}

	// ambil semua item di cart
	cartItems, err := s.cartRepo.GetItemsCartByUserID(userID)
	if err != nil {
		return errors.New("Failed to get cart items")
	}

	var cartItem *domain.CartItems
	for _, item := range cartItems {
		if item.ProductID == productID {
			cartItem = item
			break
		}
	}

	if cartItem == nil {
		return errors.New("cart item not found")
	}

	// ambil product untuk update stock
	product, err := s.productRepo.FindProductByID(productID)
	if err != nil {
		return errors.New("Product not found")
	}

	// jika quantity >= current, maka hapus item
	if quantity >= cartItem.Quantity {
		// kembalikan seluruh stock ke product
		product.Stock += cartItem.Quantity

		if _, err := s.productRepo.UpdateProduct(product); err != nil {
			return errors.New("failed to update product stock")
		}

		// hapus item dari cart
		return s.cartRepo.RemoveItemFromCart(cart.ID, productID)
	}

	// jika quantity < current, kurangi cart item dan kembalikan stock sesuai quantity
	cartItem.Quantity -= quantity
	product.Stock += quantity

	if err := s.cartRepo.UpdateCartItem(cartItem); err != nil {
		return errors.New("failed to update cart item quantity")
	}

	if _, err := s.productRepo.UpdateProduct(product); err != nil {
		return errors.New("failed to update product stock")
	}

	return nil
}
