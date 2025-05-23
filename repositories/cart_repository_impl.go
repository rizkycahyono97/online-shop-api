package repositories

import (
	"errors"
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

// GetItemsCartByUserID -> mengambil semua items yang di cart berdasarkan user_id
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

// GetAllCartWithItems -> mengembalikan semua carts beserta user dan items
func (repo CartRepositoryImpl) GetAllCartsWithItems() ([]*domain.Cart, error) {
	var carts []*domain.Cart
	err := repo.db.Preload("User").Preload("CartItems.Product").Find(&carts).Error
	if err != nil {
		return nil, errors.New("error di repo")
	}
	return carts, nil
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

// updateCartItem -> update item in cart
func (repo CartRepositoryImpl) UpdateCartItem(item *domain.CartItems) error {
	if err := repo.db.Save(item).Error; err != nil {
		return err
	}
	return nil
}

// removeItemFromCart -> menghapus item di cart
func (repo CartRepositoryImpl) RemoveItemFromCart(cartID uint, productID uint) error {
	// menghapus item dari cart berdasarkan cart_id dan product_id
	err := repo.db.Where("cart_id = ? AND product_id = ?", cartID, productID).Delete(&domain.CartItems{}).Error
	if err != nil {
		return err
	}
	return nil
}

// ClearCart -> untuk menghapus historycart
func (repo CartRepositoryImpl) ClearCart(cartID uint) error {
	return repo.db.Where("cart_id = ?", cartID).Delete(&domain.CartItems{}).Error
}
