package handlers

import (
	"context"
	"github.com/Pos-tech-FIAP-GO-HORSE/order-management/internal/core/ports/payment/process_payment"
	"github.com/Pos-tech-FIAP-GO-HORSE/order-management/internal/core/usecases/payment"
	"github.com/Pos-tech-FIAP-GO-HORSE/order-management/internal/infra/repositories"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type PaymentHandler struct {
	ProcessPaymentUseCase process_payment.IProcessPaymentUseCase
}

func NewPaymentHandler(paymentRepository repositories.IPaymentRepository, mercadoPagoClient payment.MercadoPagoClient) *PaymentHandler {
	return &PaymentHandler{
		ProcessPaymentUseCase: payment.NewProcessPaymentUseCase(paymentRepository, mercadoPagoClient),
	}
}

func (h *PaymentHandler) ProcessPayment(c *gin.Context) {
	var input process_payment.Input
	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx, cancel := context.WithTimeout(c, time.Second*5)
	defer cancel()

	if err := h.ProcessPaymentUseCase.Execute(ctx, input); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "payment created successfully",
	})
}
