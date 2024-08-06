package handlers

import (
	"context"
	"net/http"
	"time"

	"github.com/Pos-tech-FIAP-GO-HORSE/order-management/internal/core/ports/order/create_order"

	"github.com/Pos-tech-FIAP-GO-HORSE/order-management/internal/core/usecases/orders"
	"github.com/Pos-tech-FIAP-GO-HORSE/order-management/internal/infra/repositories"
	"github.com/gin-gonic/gin"
)

type OrderHandler struct {
	createOrderUseCase create_order.ICreateOrderUseCase
}

func NewOrderHandler(orderRepository repositories.IOrderRepository, productRepository repositories.IProductRepository, userRepository repositories.IUserRepository) *OrderHandler {
	return &OrderHandler{
		createOrderUseCase: orders.NewCreateProductUseCase(orderRepository, productRepository, userRepository),
	}
}

func (h *OrderHandler) CreateOrder(c *gin.Context) {
	var input create_order.Input
	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx, cancel := context.WithTimeout(c, time.Second*5)
	defer cancel()

	if err := h.createOrderUseCase.Execute(ctx, input); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "order created successfully",
	})
}
