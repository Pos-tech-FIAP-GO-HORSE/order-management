package handlers

import (
	"context"
	"github.com/Pos-tech-FIAP-GO-HORSE/order-management/internal/core/ports/order/update_order"
	"net/http"
	"time"

	"github.com/Pos-tech-FIAP-GO-HORSE/order-management/internal/core/ports/order/create_order"
	"github.com/Pos-tech-FIAP-GO-HORSE/order-management/internal/core/ports/order/find_all_orders"

	"github.com/Pos-tech-FIAP-GO-HORSE/order-management/internal/core/usecases/orders"
	"github.com/Pos-tech-FIAP-GO-HORSE/order-management/internal/infra/repositories"
	"github.com/gin-gonic/gin"
)

type OrderHandler struct {
	createOrderUseCase   create_order.ICreateOrderUseCase
	updateOrderUseCase   update_order.IUpdateOrderUseCase
	findAllOrdersUseCase find_all_orders.IFindAllOrdersUseCase
}

func NewOrderHandler(orderRepository repositories.IOrderRepository, productRepository repositories.IProductRepository, userRepository repositories.IUserRepository) *OrderHandler {
	return &OrderHandler{
		createOrderUseCase:   orders.NewCreateProductUseCase(orderRepository, productRepository, userRepository),
		updateOrderUseCase:   orders.NewUpdateOrderUseCase(orderRepository),
		findAllOrdersUseCase: orders.NewFindAllOrdersUseCase(orderRepository),
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

func (h *OrderHandler) UpdateOrder(c *gin.Context) {
	var input update_order.Input
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := c.ShouldBindUri(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx, cancel := context.WithTimeout(c, time.Second*5)
	defer cancel()

	if err := h.updateOrderUseCase.Execute(ctx, input); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "order has been updated successfully",
	})
}

func (h *OrderHandler) FindAllOrders(c *gin.Context) {
	var input find_all_orders.Input
	if err := c.Bind(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx, cancel := context.WithTimeout(c, time.Second*5)
	defer cancel()

	orders, err := h.findAllOrdersUseCase.Execute(ctx, input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, orders)
}
