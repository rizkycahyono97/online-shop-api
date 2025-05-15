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

// Get items in authenticated user's cart
func (cc *CartController) GetCartItems(c *gin.Context) {
	userID := c.MustGet("userID").(uint)

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
		Data:    items,
	})
}

// Add item to cart
func (cc *CartController) AddItem(c *gin.Context) {
	userID := c.MustGet("userID").(uint)

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
			Message: "Failed to add item to cart",
			Data:    nil,
		})
		return
	}

	c.JSON(http.StatusCreated, web.ApiResponse{
		Code:    "CREATED",
		Message: "Cart Items Add Success",
		Data:    item,
	})
}

// remove item from cart
func (cc *CartController) RemoveItemFromCart(c *gin.Context) {
	userID := c.MustGet("userID").(uint)

	productIDParam := c.Param("product_id")
	productID, err := strconv.Atoi(productIDParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, web.ApiResponse{
			Code:    "BAD_REQUEST",
			Message: "Invalid Product Id",
			Data:    nil,
		})
		return
	}

	err = cc.cartService.RemoveItemFromCart(userID, uint(productID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, web.ApiResponse{
			Code:    "INTERNAL_SERVER_ERROR",
			Message: "Failed to remove item from cart",
			Data:    nil,
		})
		return
	}

	c.JSON(http.StatusOK, web.ApiResponse{
		Code:    "SUCCESS",
		Message: "Cart Items Remove Success",
		Data:    nil,
	})
}
