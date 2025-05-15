package repositories

import "github.com/rizkycahyono97/online-shop-api/model/domain"

type CartRepository interface {
	GetCartByUserID(userID uint) (*domain.Cart, error)
	GetItemsCartByUserID(userID uint) ([]*domain.CartItems, error)
	AddItemToCart(cartID uint, productID uint, quantity uint) (*domain.CartItems, error)
	UpdateCartItem(item *domain.CartItems) error
	RemoveItemFromCart(cartID uint, productID uint) error
}
