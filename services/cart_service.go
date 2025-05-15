package services

import "github.com/rizkycahyono97/online-shop-api/model/domain"

type CartService interface {
	GetItemsCartByUserID(userID uint) ([]*domain.CartItems, error)
	AddItemToCart(userID uint, productID uint, quantity uint) (*domain.CartItems, error)
	RemoveItemFromCart(userID uint, productID uint, quantity uint) error
}
