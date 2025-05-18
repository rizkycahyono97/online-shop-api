package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/rizkycahyono97/online-shop-api/helpers"
	"github.com/rizkycahyono97/online-shop-api/model/web"
	"github.com/rizkycahyono97/online-shop-api/services"
	"net/http"
	"strconv"
)

type OrderController struct {
	orderService services.OrderService
}

func NewOrderController(orderService services.OrderService) *OrderController {
	return &OrderController{orderService: orderService}
}

// CreateOrder -> untuk membuat order
func (oc *OrderController) CreateOrder(c *gin.Context) {
	userID := uint(c.MustGet("user_id").(float64))

	var req web.CreateOrderRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, web.ApiResponse{
			Code:    "BAD_REQUEST",
			Message: "Invalid Request",
			Data:    nil,
		})
		return
	}

	order, err := oc.orderService.CreateOrder(userID, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, web.ApiResponse{
			Code:    "INTERNAL_ERROR",
			Message: err.Error(),
			Data:    nil,
		})
		return
	}

	c.JSON(http.StatusOK, web.ApiResponse{
		Code:    "SUCCESS",
		Message: "Order Created Successfully",
		Data:    web.OrderResponseFromModels(order),
	})
}

// GetUserOrder -> untuk mendapatkan semua order (user)
func (oc *OrderController) GetUserOrders(c *gin.Context) {
	userID := uint(c.MustGet("user_id").(float64))

	orders, err := oc.orderService.GetOrderByUserID(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, web.ApiResponse{
			Code:    "INTERNAL_ERROR",
			Message: "Failed to fetch orders",
			Data:    nil,
		})
		return
	}

	var response []web.OrderResponse
	for _, order := range orders {
		response = append(response, *web.OrderResponseFromModels(order))
	}

	c.JSON(http.StatusOK, web.ApiResponse{
		Code:    "SUCCESS",
		Message: "Successfully fetched orders",
		Data:    response,
	})
}

// GetOrderById -> untuk mengambil order berdasarkan id
func (oc *OrderController) GetOrderByID(c *gin.Context) {
	orderIDParam := c.Param("id")
	orderIDUint, err := strconv.ParseUint(orderIDParam, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, web.ApiResponse{
			Code:    "BAD_REQUEST",
			Message: "Invalid order ID",
		})
		return
	}

	// Ambil userID dari JWT context
	userIDFloat := c.MustGet("user_id").(float64)
	userID := uint(userIDFloat)

	order, err := oc.orderService.GetOrderByID(uint(orderIDUint), userID)
	if err != nil {
		c.JSON(http.StatusForbidden, web.ApiResponse{
			Code:    "FORBIDDEN",
			Message: "You are not allowed to access this order",
		})
		return
	}

	c.JSON(http.StatusOK, web.ApiResponse{
		Code:    "SUCCESS",
		Message: "Order fetched successfully",
		Data:    order,
	})
}

// CancelledOrder -> membatalkan order
func (oc *OrderController) CancelOrder(c *gin.Context) {
	orderID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, web.ApiResponse{
			Code:    "BAD_REQUEST",
			Message: "Invalid Order ID",
		})
		return
	}

	userIDFloat := c.MustGet("user_id").(float64)
	userID := uint(userIDFloat)

	err = oc.orderService.CancelOrder(userID, uint(orderID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, web.ApiResponse{
			Code:    "INTERNAL_ERROR",
			Message: err.Error(),
			Data:    nil,
		})
		return
	}

	c.JSON(http.StatusOK, web.ApiResponse{
		Code:    "SUCCESS",
		Message: "Successfully cancelled order",
	})
}

// ConfirmOrder -> konfirm order jika sudah sampai
func (oc *OrderController) ConfirmOrder(c *gin.Context) {
	orderID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, web.ApiResponse{
			Code:    "BAD_REQUEST",
			Message: "Invalid Order ID",
			Data:    nil,
		})
		return
	}

	userIDFloat := c.MustGet("user_id").(float64)
	userID := uint(userIDFloat)

	err = oc.orderService.ConfirmOrder(userID, uint(orderID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, web.ApiResponse{
			Code:    "INTERNAL_ERROR",
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, web.ApiResponse{
		Code:    "SUCCESS",
		Message: "Order Confirmed Successfully",
	})
}

// GetOrderByUserID -> mendapatkan order berdsarkan userID khusus untuk admin
func (oc *OrderController) GetOrderByUserID(c *gin.Context) {
	userIDParam := c.Param("user_id")
	userIDUint, err := strconv.Atoi(userIDParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, web.ApiResponse{
			Code:    "BAD_REQUEST",
			Message: "Invalid user ID",
		})
		return
	}

	targetUserID := uint(userIDUint)

	if !helpers.IsOwnerOrAdmin(c, targetUserID) {
		c.JSON(http.StatusForbidden, web.ApiResponse{
			Code:    "FORBIDDEN",
			Message: "You are not allowed to access this order",
		})
		return
	}

	orders, err := oc.orderService.GetOrderByUserID(targetUserID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, web.ApiResponse{
			Code:    "INTERNAL_ERROR",
			Message: err.Error(),
			Data:    nil,
		})
		return
	}

	c.JSON(http.StatusOK, web.ApiResponse{
		Code:    "SUCCESS",
		Message: "Successfully fetched orders",
		Data:    web.OrderResponseListFromModels(orders),
	})

}

// UpdateOrderStatus -> update status order for admin
func (oc *OrderController) UpdateOrderStatus(c *gin.Context) {
	var req web.UpdateOrderStatusRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, web.ApiResponse{
			Code:    "BAD_REQUEST",
			Message: err.Error(),
		})
		return
	}

	orderIDParam := c.Param("id")
	orderIDUint, err := strconv.ParseUint(orderIDParam, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, web.ApiResponse{
			Code:    "BAD_REQUEST",
			Message: "Invalid order ID",
		})
		return
	}

	updatedOrder, err := oc.orderService.UpdateOrderStatus(uint(orderIDUint), req.Status)
	if err != nil {
		c.JSON(http.StatusInternalServerError, web.ApiResponse{
			Code:    "INTERNAL_ERROR",
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, web.ApiResponse{
		Code:    "SUCCESS",
		Message: "Successfully updated order",
		Data:    web.OrderResponseFromModels(updatedOrder),
	})
}
