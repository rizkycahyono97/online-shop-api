package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/rizkycahyono97/online-shop-api/model/web"
	"github.com/rizkycahyono97/online-shop-api/services"
	"net/http"
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
		c.JSON(http.StatusInternalServerError, web.ApiResponse{
			Code:    "INTERNAL_SERVER_ERROR",
			Message: "Failed to fetch cart items",
			Data:    nil,
		})
		return
	}

	c.JSON(http.StatusOK, web.ApiResponse{
		Code:    "SUCCESS",
		Message: "Cart Items Fetch Success",
		Data:    web.CartItemsResponseFromModels(items),
	})
}

// Add item to cart
func (cc *CartController) AddItem(c *gin.Context) {
	userIDFloat := c.MustGet("user_id").(float64)
	userID := uint(userIDFloat)

	var req web.AddItemToCartRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, web.ApiResponse{
			Code:    "BAD_REQUEST",
			Message: "Invalid Request",
			Data:    nil,
		})
		return
	}

	item, err := cc.cartService.AddItemToCart(userID, req.ProductID, req.Quantity)
	if err != nil {
		c.JSON(http.StatusInternalServerError, web.ApiResponse{
			Code:    "INTERNAL_SERVER_ERROR",
			Message: "Failed to add item to cart, Stock is 0",
			Data:    nil,
		})
		return
	}

	c.JSON(http.StatusCreated, web.ApiResponse{
		Code:    "CREATED",
		Message: "Cart Items Add Success",
		Data:    web.CartItemResponseFromModels(item),
	})
}

// remove item from cart
func (cc *CartController) RemoveItemFromCart(c *gin.Context) {
	userIDFloat := c.MustGet("user_id").(float64)
	userID := uint(userIDFloat)

	productIDParam := c.Param("product_id")
	productIDUint, err := strconv.Atoi(productIDParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, web.ApiResponse{
			Code:    "BAD_REQUEST",
			Message: "Invalid Product Id",
			Data:    nil,
		})
		return
	}

	productID := uint(productIDUint)
	var req web.RemoveItemFromCartRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, web.ApiResponse{
			Code:    "BAD_REQUEST",
			Message: "Invalid Request or missing quantity",
			Data:    nil,
		})
		return
	}

	err = cc.cartService.RemoveItemFromCart(userID, productID, req.Quantity)
	if err != nil {
		c.JSON(http.StatusInternalServerError, web.ApiResponse{
			Code:    "INTERNAL_SERVER_ERROR",
			Message: "Failed to add item to cart, Stock is 0",
			Data:    nil,
		})
		return
	}

	c.JSON(http.StatusOK, web.ApiResponse{
		Code:    "SUCCESS",
		Message: "Item Remove Successfully",
		Data:    nil,
	})
}
