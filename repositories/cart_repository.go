package repositories

import "github.com/rizkycahyono97/online-shop-api/model/domain"

type CartRepository interface {
	AddCart(userID uint) (*domain.Cart, error)
	GetCartByUserID(userID uint) (*domain.Cart, error)
	AddItemToCart(cartID uint, productID uint, quantity uint) (*domain.Cart, error)
	RemoveItemFromCart(cartID uint, productID uint) error
	DeleteCart(cartID uint) error
}
