package controller

import (
	"github.com/gin-gonic/gin"
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
			Message: "Failed to create order",
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

// GetUserOrder -> untuk mendapatkan semua order
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
		response = append(response, web.OrderResponseFromModels(order))
	}

	c.JSON(http.StatusOK, web.ApiResponse{
		Code:    "SUCCESS",
		Message: "Successfully fetched orders",
		Data:    response,
	})
}

// GetOrderById -> untuk mengambil order berdasarkan id
func (oc *OrderController) GetOrderByID(c *gin.Context) {
	orderIDStr := c.Param("id")
	orderID, err := strconv.Atoi(orderIDStr)

	order, err := oc.orderService.GetOrderByID(uint(orderID))
	if err != nil {
		c.JSON(http.StatusNotFound, web.ApiResponse{
			Code:    "NOT_FOUND",
			Message: "Order Not Found",
			Data:    nil,
		})
		return
	}

	c.JSON(http.StatusOK, web.ApiResponse{
		Code:    "SUCCESS",
		Message: "Successfully fetched order",
		Data:    web.OrderResponseFromModels(order),
	})
}
