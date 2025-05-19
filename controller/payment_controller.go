package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/rizkycahyono97/online-shop-api/model/domain"
	"github.com/rizkycahyono97/online-shop-api/model/web"
	"github.com/rizkycahyono97/online-shop-api/services"
	"net/http"
	"strconv"
)

type PaymentController struct {
	paymentService services.PaymentService
}

func NewPaymentController(paymentService services.PaymentService) *PaymentController {
	return &PaymentController{
		paymentService: paymentService,
	}
}

// CreatePayment (Customer)
func (pc *PaymentController) CreatePayment(c *gin.Context) {
	var payment domain.Payment
	if err := c.ShouldBindJSON(&payment); err != nil {
		c.JSON(http.StatusBadRequest, web.ApiResponse{
			Code:    "BAD_REQUEST",
			Message: err.Error(),
			Data:    nil,
		})
		return
	}

	created, err := pc.paymentService.CreatePayment(&payment)
	if err != nil {
		c.JSON(http.StatusInternalServerError, web.ApiResponse{
			Code:    "INTERNAL_ERROR",
			Message: err.Error(),
			Data:    nil,
		})
		return
	}

	c.JSON(http.StatusCreated, web.ApiResponse{
		Code:    "SUCCESS",
		Message: "Payment created",
		Data:    web.PaymentResponseFromModel(created),
	})
}

// UpdatePaymentStatus (Admin)
func (pc *PaymentController) UpdatePaymentStatus(c *gin.Context) {
	paymentIDStr := c.Param("order_id")
	paymentID, err := strconv.Atoi(paymentIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, web.ApiResponse{
			Code:    "BAD_REQUEST",
			Message: err.Error(),
			Data:    nil,
		})
		return
	}

	var req web.PaymentUpdateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, web.ApiResponse{
			Code:    "BAD_REQUEST",
			Message: err.Error(),
			Data:    nil,
		})
		return
	}

	updated, err := pc.paymentService.UpdatePaymentStatus(uint(paymentID), req.Status)
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
		Message: "Payment updated",
		Data:    web.PaymentResponseFromModel(updated),
	})
}

func (pc *PaymentController) GetPaymentByID(c *gin.Context) {
	paymentIDStr := c.Param("order_id")
	paymentID, err := strconv.Atoi(paymentIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, web.ApiResponse{
			Code:    "BAD_REQUEST",
			Message: err.Error(),
			Data:    nil,
		})
		return
	}

	payment, err := pc.paymentService.GetPaymentByOrderID(uint(paymentID))
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
		Message: "Payment found",
		Data:    web.PaymentResponseFromModel(payment),
	})
}

func (pc *PaymentController) GetPaymentByUserID(c *gin.Context) {
	userIDFloat, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusBadRequest, web.ApiResponse{
			Code:    "BAD_REQUEST",
			Message: "User ID not found",
			Data:    nil,
		})
		return
	}

	userID := uint(userIDFloat.(float64))

	payments, err := pc.paymentService.GetPaymentByUserID(userID)
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
		Message: "Payments found",
		Data:    web.PaymentResponseFromModels(payments),
	})
}

func (pc *PaymentController) GetAllPayment(c *gin.Context) {
	payments, err := pc.paymentService.GetAllPayment()
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
		Message: "Payments found",
		Data:    web.PaymentResponseFromModels(payments),
	})
}
