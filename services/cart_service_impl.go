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

	return item, nil
}

// Delete item from cart
func (s CartServiceImpl) RemoveItemFromCart(cartID uint, productID uint) error {
	// cari berdasarkan user
	cart, err := s.cartRepo.GetCartByUserID(cartID)
	if err != nil {
		return err
	}

	return s.cartRepo.RemoveItemFromCart(cart.ID, productID)
}
