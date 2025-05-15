package repositories

import (
	"github.com/rizkycahyono97/online-shop-api/model/domain"
	"gorm.io/gorm"
)

type CartRepositoryImpl struct {
	db *gorm.DB
}

func NewCartRepositoryImpl(db *gorm.DB) CartRepository {
	return &CartRepositoryImpl{db: db}
}

// GetCartByUserID
func (repo CartRepositoryImpl) GetCartByUserID(userID uint) (*domain.Cart, error) {
	//TODO implement me
	panic("implement me")
}

func (repo CartRepositoryImpl) AddItemToCart(cartID uint, productID uint, quantity uint) (*domain.Cart, error) {
	//TODO implement me
	panic("implement me")
}

func (repo CartRepositoryImpl) RemoveItemFromCart(cartID uint, productID uint) error {
	//TODO implement me
	panic("implement me")
}

func (repo CartRepositoryImpl) DeleteCart(cartID uint) error {
	//TODO implement me
	panic("implement me")
}
