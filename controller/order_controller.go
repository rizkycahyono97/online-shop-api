package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/rizkycahyono97/online-shop-api/helpers"
	"github.com/rizkycahyono97/online-shop-api/model/web"
	"github.com/rizkycahyono97/online-shop-api/services"
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
		helpers.JSONBadRequestResponse(c, "Invalid Request", nil)
		return
	}

	order, err := oc.orderService.CreateOrder(userID, &req)
	if err != nil {
		helpers.JSONInternalErrorResponse(c, "", err)
		return
	}

	helpers.JSONSuccessResponse(c, "Order Created Successfully", web.OrderResponseFromModels(order))
}

// GetUserOrder -> untuk mendapatkan semua order (user)
func (oc *OrderController) GetUserOrders(c *gin.Context) {
	userID := uint(c.MustGet("user_id").(float64))

	orders, err := oc.orderService.GetOrderByUserID(userID)
	if err != nil {
		helpers.JSONInternalErrorResponse(c, "Failed to fetch orders", nil)
		return
	}

	var response []web.OrderResponse
	for _, order := range orders {
		response = append(response, *web.OrderResponseFromModels(order))
	}

	helpers.JSONSuccessResponse(c, "Successfully fetched orders", response)
}

// GetallUserOrder -> untuk mendapatkan semua order (admin)
func (oc *OrderController) GetAllUserOrders(c *gin.Context) {
	orders, err := oc.orderService.GetAllUserOrders()
	if err != nil {
		helpers.JSONInternalErrorResponse(c, "", err)
		return
	}

	helpers.JSONSuccessResponse(c, "Successfully fetched orders", orders)
}

// GetOrderById -> untuk mengambil order berdasarkan id
func (oc *OrderController) GetOrderByID(c *gin.Context) {
	orderIDParam := c.Param("id")
	orderIDUint, err := strconv.ParseUint(orderIDParam, 10, 32)
	if err != nil {
		helpers.JSONBadRequestResponse(c, "Invalid order ID", nil)
		return
	}

	// Ambil userID dari JWT context
	userIDFloat := c.MustGet("user_id").(float64)
	userID := uint(userIDFloat)

	order, err := oc.orderService.GetOrderByID(uint(orderIDUint), userID)
	if err != nil {
		helpers.JSONInternalErrorResponse(c, "You are not allowed to access this order", nil)
		return
	}

	helpers.JSONSuccessResponse(c, "Order fetched successfully", order)
}

// CancelledOrder -> membatalkan order
func (oc *OrderController) CancelOrder(c *gin.Context) {
	orderID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		helpers.JSONBadRequestResponse(c, "Invalid order ID", nil)
		return
	}

	userIDFloat := c.MustGet("user_id").(float64)
	userID := uint(userIDFloat)

	err = oc.orderService.CancelOrder(userID, uint(orderID))
	if err != nil {
		helpers.JSONBadRequestResponse(c, "", err)
		return
	}

	helpers.JSONSuccessResponse(c, "Successfully cancelled order", nil)
}

// ConfirmOrder -> konfirm order jika sudah sampai
func (oc *OrderController) ConfirmOrder(c *gin.Context) {
	orderID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		helpers.JSONBadRequestResponse(c, "Invalid order ID", nil)
		return
	}

	userIDFloat := c.MustGet("user_id").(float64)
	userID := uint(userIDFloat)

	err = oc.orderService.ConfirmOrder(userID, uint(orderID))
	if err != nil {
		helpers.JSONInternalErrorResponse(c, "", err)
		return
	}

	helpers.JSONSuccessResponse(c, "Successfully confirmed order", nil)
}

// GetOrderByUserID -> mendapatkan order berdsarkan userID khusus untuk admin
func (oc *OrderController) GetOrderByUserID(c *gin.Context) {
	userIDParam := c.Param("user_id")
	userIDUint, err := strconv.Atoi(userIDParam)
	if err != nil {
		helpers.JSONBadRequestResponse(c, "Invalid order ID", nil)
		return
	}

	targetUserID := uint(userIDUint)

	if !helpers.IsOwnerOrAdmin(c, targetUserID) {
		helpers.JSONForbiddenResponse(c, "You are not allowed to access this order", nil)
		return
	}

	orders, err := oc.orderService.GetOrderByUserID(targetUserID)
	if err != nil {
		helpers.JSONInternalErrorResponse(c, "", err)
		return
	}

	helpers.JSONSuccessResponse(c, "Successfully fetched orders", web.OrderResponseListFromModels(orders))
}

// UpdateOrderStatus -> update status order for admin
func (oc *OrderController) UpdateOrderStatus(c *gin.Context) {
	var req web.UpdateOrderStatusRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		helpers.JSONBadRequestResponse(c, "", err)
		return
	}

	orderIDParam := c.Param("id")
	orderIDUint, err := strconv.ParseUint(orderIDParam, 10, 32)
	if err != nil {
		helpers.JSONBadRequestResponse(c, "Invalid order ID", nil)
		return
	}

	updatedOrder, err := oc.orderService.UpdateOrderStatus(uint(orderIDUint), req.Status)
	if err != nil {
		helpers.JSONInternalErrorResponse(c, "", err)
		return
	}

	helpers.JSONSuccessResponse(c, "Successfully updated order", web.OrderResponseFromModels(updatedOrder))
}
