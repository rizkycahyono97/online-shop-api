package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/rizkycahyono97/online-shop-api/helpers"
	"github.com/rizkycahyono97/online-shop-api/model/web"
	"github.com/rizkycahyono97/online-shop-api/services"
	"strconv"
)

type CartController struct {
	cartService services.CartService
}

func NewCartController(cartService services.CartService) *CartController {
	return &CartController{cartService: cartService}
}

// Get all items in authenticated user's cart
func (cc *CartController) GetCartItems(c *gin.Context) {
	// ambil user_id dari JWT
	userIDFloat := c.MustGet("user_id").(float64)
	userID := uint(userIDFloat)

	items, err := cc.cartService.GetItemsCartByUserID(userID)
	if err != nil {
		helpers.JSONInternalErrorResponse(c, "Failed to fetch cart items", nil)
		return
	}

	helpers.JSONSuccessResponse(c, "Cart Items Fetch Success", web.CartItemsResponseFromModels(items))
}

// mengemnalikan semua carts dan juga semua users
func (cc *CartController) GetAllCarts(c *gin.Context) {
	carts, err := cc.cartService.GetAllCartsWithItems()
	helpers.JSONInternalErrorResponse(c, "", err)
	helpers.JSONSuccessResponse(c, "Cart Items Fetch Successfully", web.CartAdminResponseFromModels(carts))
}

// Add item to cart
func (cc *CartController) AddItem(c *gin.Context) {
	userIDFloat := c.MustGet("user_id").(float64)
	userID := uint(userIDFloat)

	var req web.AddItemToCartRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		helpers.JSONBadRequestResponse(c, "Invalid Request", nil)
		return
	}

	item, err := cc.cartService.AddItemToCart(userID, req.ProductID, req.Quantity)
	if err != nil {
		helpers.JSONInternalErrorResponse(c, "Failed to add item to cart, Stock is 0", nil)
		return
	}

	helpers.JSONSuccessResponse(c, "Cart Items Add Success", web.CartItemResponseFromModels(item))
}

// remove item from cart
func (cc *CartController) RemoveItemFromCart(c *gin.Context) {
	userIDFloat := c.MustGet("user_id").(float64)
	userID := uint(userIDFloat)

	productIDParam := c.Param("product_id")
	productIDUint, err := strconv.Atoi(productIDParam)
	if err != nil {
		helpers.JSONBadRequestResponse(c, "Invalid Product ID", nil)
		return
	}

	productID := uint(productIDUint)
	var req web.RemoveItemFromCartRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		helpers.JSONBadRequestResponse(c, "Invalid Request or missing quantity", nil)
		return
	}

	err = cc.cartService.RemoveItemFromCart(userID, productID, req.Quantity)
	if err != nil {
		helpers.JSONInternalErrorResponse(c, "Failed to add item to cart, Stock is 0", err)
		return
	}

	helpers.JSONSuccessResponse(c, "Item Remove Successfully", nil)
}
