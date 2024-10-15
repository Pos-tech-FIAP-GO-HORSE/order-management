package handlers

import (
	"context"
	"github.com/Pos-tech-FIAP-GO-HORSE/order-management/internal/core/ports/create_payment"
	"github.com/Pos-tech-FIAP-GO-HORSE/order-management/internal/core/usecases/payments"
	"github.com/Pos-tech-FIAP-GO-HORSE/order-management/internal/infra/gateway/payment"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type PaymentHandler struct {
	createPaymentUseCase create_payment.ICreatePaymentUseCase
}

func NewPaymentHandler(paymentProcessor payment.IPaymentProcessor) *PaymentHandler {
	return &PaymentHandler{
		createPaymentUseCase: payments.NewCreatePaymentUseCase(paymentProcessor),
	}
}

// CreatePayment godoc
// @Summary      Create a new payment
// @Description  Add a new payment to order
// @Tags         Payments
// @Accept       json
// @Produce      json
// @Param        create_payment   body      create_payment.Input  true  "Payment Data"
// @Success      201     {object}  ResponseMessage
// @Failure      400     {object}  ResponseMessage
// @Failure      500     {object}  ResponseMessage
// @Router       /api/v1/payments [post]
func (h *PaymentHandler) CreatePayment(c *gin.Context) {
	var input create_payment.Input
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
