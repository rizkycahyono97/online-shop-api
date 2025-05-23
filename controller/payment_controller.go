package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/rizkycahyono97/online-shop-api/helpers"
	"github.com/rizkycahyono97/online-shop-api/model/web"
	"github.com/rizkycahyono97/online-shop-api/services"
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
	userID := uint(c.MustGet("user_id").(float64))

	var req web.PaymentCreateRequest
	if err := c.BindJSON(&req); err != nil {
		helpers.JSONBadRequestResponse(c, "", err)
		return
	}

	payment := web.PaymentRequestToModel(req)

	created, err := pc.paymentService.CreatePayment(userID, payment)
	if err != nil {
		helpers.JSONInternalErrorResponse(c, "", err)
		return
	}

	helpers.JSONSuccessResponse(c, "Payment Created", created)
}

// UpdatePaymentStatus (Admin)
func (pc *PaymentController) UpdatePaymentStatus(c *gin.Context) {
	paymentIDStr := c.Param("order_id")
	paymentID, err := strconv.Atoi(paymentIDStr)
	if err != nil {
		helpers.JSONBadRequestResponse(c, "", err)
		return
	}

	var req web.PaymentUpdateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		helpers.JSONBadRequestResponse(c, "", err)
		return
	}

	updated, err := pc.paymentService.UpdatePaymentStatus(uint(paymentID), req.Status)
	if err != nil {
		helpers.JSONInternalErrorResponse(c, "", err)
		return
	}

	helpers.JSONSuccessResponse(c, "Payment Updated", updated)
}

// GET /api/v1/payments/:id
func (pc *PaymentController) GetPaymentByID(c *gin.Context) {
	paymentIDStr := c.Param("order_id")
	paymentID, err := strconv.Atoi(paymentIDStr)
	if err != nil {
		helpers.JSONBadRequestResponse(c, "", err)
		return
	}

	payment, err := pc.paymentService.GetPaymentByOrderID(uint(paymentID))
	if err != nil {
		helpers.JSONInternalErrorResponse(c, "", err)
		return
	}

	helpers.JSONSuccessResponse(c, "Payment Found", payment)
}

// GET /api/v1/users/payments
func (pc *PaymentController) GetPaymentForUser(c *gin.Context) {
	userIDFloat, exists := c.Get("user_id")
	if !exists {
		helpers.JSONBadRequestResponse(c, "user id not found", nil)
		return
	}

	userID := uint(userIDFloat.(float64))

	payments, err := pc.paymentService.GetPaymentsForUser(userID)
	if err != nil {
		helpers.JSONInternalErrorResponse(c, "", err)
		return
	}

	helpers.JSONSuccessResponse(c, "Payments Found", payments)
}

// GET /api/v1/payments (admin only)
func (pc *PaymentController) GetAllPayment(c *gin.Context) {
	payments, err := pc.paymentService.GetAllPayment()
	if err != nil {
		helpers.JSONInternalErrorResponse(c, "", err)
		return
	}

	helpers.JSONSuccessResponse(c, "Payments Found", web.PaymentResponseFromModels(payments))
}

// GET /api/v1/payments/:user_id
func (pc *PaymentController) GetPaymentsByUserID(c *gin.Context) {
	userIDStr := c.Param("user_id")
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		helpers.JSONBadRequestResponse(c, "invalid user ID", nil)
		return
	}

	payments, err := pc.paymentService.GetPaymentsForUser(uint(userID))
	if err != nil {
		helpers.JSONInternalErrorResponse(c, "", err)
		return
	}

	helpers.JSONSuccessResponse(c, "Payments Found", web.PaymentResponseFromModels(payments))
}
