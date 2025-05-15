package repositories

import (
	"github.com/rizkycahyono97/online-shop-api/model/domain"
	"gorm.io/gorm"
)

type CartRepositoryImpl struct {
	db *gorm.DB
}

func NewCartRepository(db *gorm.DB) CartRepository {
	return &CartRepositoryImpl{db: db}
}

func (repo CartRepositoryImpl) GetCartByUserID(userID uint) (*domain.Cart, error) {
	var cart domain.Cart
	if err := repo.db.Preload("CartItems").Where("user_id = ?", userID).First(&cart).Error; err != nil {
		return nil, err
	}
	return &cart, nil
}

// GetItemsCartByUserID -> mengambil semua item yang di cart
func (repo CartRepositoryImpl) GetItemsCartByUserID(userID uint) ([]*domain.CartItems, error) {
	// ambil cart berdasarkan userID
	var cart *domain.Cart
	if err := repo.db.Where("user_id = ?", userID).First(&cart).Error; err != nil {
		return nil, err
	}

	// ambil cartitems berdasarkan cartID, sekaligus ambil products juga
	var cartItems []*domain.CartItems
	if err := repo.db.Preload("Product").Where("cart_id = ?", cart.ID).Find(&cartItems).Error; err != nil {
		return nil, err
	}

	return cartItems, nil
}

// AddItemToCart -> menambahkan item ke cart
func (repo CartRepositoryImpl) AddItemToCart(cartID, productID uint, quantity uint) (*domain.CartItems, error) {
	var cartItem domain.CartItems

	// Cek apakah produk sudah ada di cart
	err := repo.db.Where("cart_id = ? AND product_id = ?", cartID, productID).First(&cartItem).Error
	if err == nil {
		// Jika item sudah ada, update quantity
		cartItem.Quantity += quantity
		if err := repo.db.Save(&cartItem).Error; err != nil {
			return nil, err
		}

		// preload product
		if err := repo.db.Preload("Product").First(&cartItem, cartItem.ID).Error; err != nil {
			return nil, err
		}

		return &cartItem, nil
	}

	// Jika item belum ada, tambahkan item baru
	cartItem = domain.CartItems{
		CartID:    cartID,
		ProductID: productID,
		Quantity:  quantity,
	}

	// Simpan item ke cart
	if err := repo.db.Create(&cartItem).Error; err != nil {
		return nil, err
	}

	// Preload Product setelah create, pakai ID
	if err := repo.db.Preload("Product").First(&cartItem, cartItem.ID).Error; err != nil {
		return nil, err
	}

	return &cartItem, nil
}

func (repo CartRepositoryImpl) RemoveItemFromCart(cartID uint, productID uint) error {
	// menghapus item dari cart berdasarkan cart_id dan product_id
	err := repo.db.Where("cart_id = ? AND product_id = ?", cartID, productID).Delete(&domain.CartItems{}).Error
	if err != nil {
		return err
	}
	return nil
}
