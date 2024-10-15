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

type ResponseMessagePayment struct {
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
// @Success      201     {object}  ResponseMessage
// @Failure      400     {object}  ResponseMessage
// @Failure      500     {object}  ResponseMessage
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
// @Summary      Get a new payments_processor status
// @Description  Get a order payments_processor status
// @Tags         Payments
// @Accept       json
// @Produce      json
// @Param        id   query     string  true  "Payment ID"
// @Success      201     {object}  ResponseMessagePayment
// @Failure      400     {object}  ResponseMessagePayment
// @Failure      500     {object}  ResponseMessagePayment
// @Router       /api/v1/payments [get]
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
