package handlers

import (
	"context"
	"github.com/Pos-tech-FIAP-GO-HORSE/order-management/internal/core/ports/payment"
	"github.com/Pos-tech-FIAP-GO-HORSE/order-management/internal/core/usecases/payments"
	"github.com/Pos-tech-FIAP-GO-HORSE/order-management/internal/infra/gateway/payments_processor"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

type PaymentHandler struct {
	createPaymentUseCase    payment.ICreatePaymentUseCase
	getStatusPaymentUseCase payment.IGetStatusPaymentUseCase
}

func NewPaymentHandler(paymentProcessor payments_processor.IPaymentProcessor) *PaymentHandler {
	return &PaymentHandler{
		createPaymentUseCase:    payments.NewCreatePaymentUseCase(paymentProcessor),
		getStatusPaymentUseCase: payments.NewGetStatusPaymentUseCase(paymentProcessor),
	}
}

type ResponseCreatePayment struct {
	Message payments_processor.ResponseCreatePayment `json:"result"`
	Error   string                                   `json:"error,omitempty"`
}

type ResponseStatusPayment struct {
	Message payments_processor.ResponseStatusPayment `json:"result"`
	Error   string                                   `json:"error,omitempty"`
}

// CreatePayment godoc
// @Summary      Create a new payment
// @Description  Add a new payment to order
// @Tags         Payments
// @Accept       json
// @Produce      json
// @Param        create_payment   body      payment.Input  true  "Payment Data"
// @Success      200     {object}  ResponseCreatePayment
// @Failure      400     {object}  ResponseCreatePayment
// @Failure      500     {object}  ResponseCreatePayment
// @Router       /api/v1/payments [post]
func (h *PaymentHandler) CreatePayment(c *gin.Context) {
	var input payment.Input
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx, cancel := context.WithTimeout(c, time.Second*5)
	defer cancel()

	paymentInfos, err := h.createPaymentUseCase.Execute(ctx, input)
	if err != nil {
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":       "payment created successfully",
		"paymentQRCode": paymentInfos.QRCode,
		"paymentId":     paymentInfos.ID,
	})
}

// GetStatusPayment StatusPayment godoc
// @Summary      Get a payment status
// @Description  Get a payment order status
// @Tags         Payments
// @Accept       json
// @Produce      json
// @Param        id   path     string  true  "Payment ID"
// @Success      200     {object}  ResponseStatusPayment
// @Failure      400     {object}  ResponseStatusPayment
// @Failure      500     {object}  ResponseStatusPayment
// @Router       /api/v1/payments/{id} [get]
func (h *PaymentHandler) GetStatusPayment(c *gin.Context) {
	param := c.Param("id")
	if param == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Payment ID is required",
		})
		return
	}

	paymentID, err := strconv.Atoi(param)

	ctx, cancel := context.WithTimeout(c, time.Second*10)
	defer cancel()

	paymentStatusResponse, err := h.getStatusPaymentUseCase.Execute(ctx, paymentID)
	if err != nil {
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"result": paymentStatusResponse,
	})
}
