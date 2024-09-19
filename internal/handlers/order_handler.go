package handlers

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/Pos-tech-FIAP-GO-HORSE/order-management/internal/core/ports/order/create_order"
	"github.com/Pos-tech-FIAP-GO-HORSE/order-management/internal/core/ports/order/find_all_orders"
	"github.com/Pos-tech-FIAP-GO-HORSE/order-management/internal/core/ports/order/update_order"
	"github.com/Pos-tech-FIAP-GO-HORSE/order-management/internal/core/ports/order/update_order_status"
	"github.com/Pos-tech-FIAP-GO-HORSE/order-management/internal/core/usecases/orders"
	"github.com/Pos-tech-FIAP-GO-HORSE/order-management/internal/infra/repositories"
)

type OrderHandler struct {
	createOrderUseCase       create_order.ICreateOrderUseCase
	findAllOrdersUseCase     find_all_orders.IFindAllOrdersUseCase
	updateOrderUseCase       update_order.IUpdateOrderUseCase
	updateOrderStatusUseCase update_order_status.IUpdateOrderStatusUseCase
}

func NewOrderHandler(orderRepository repositories.IOrderRepository, productRepository repositories.IProductRepository, userRepository repositories.IUserRepository) *OrderHandler {
	return &OrderHandler{
		createOrderUseCase:       orders.NewCreateProductUseCase(orderRepository, productRepository, userRepository),
		findAllOrdersUseCase:     orders.NewFindAllOrdersUseCase(orderRepository),
		updateOrderUseCase:       orders.NewUpdateOrderUseCase(orderRepository),
		updateOrderStatusUseCase: orders.NewUpdateOrderStatusUseCase(orderRepository),
	}
}

// CreateOrder godoc
// @Summary      Create a new order
// @Description  Add a new order to the system
// @Tags         Orders
// @Accept       json
// @Produce      json
// @Param        order   body      create_order.Input  true  "Order Data"
// @Success      201     {object}  ResponseMessage
// @Failure      400     {object}  ResponseMessage
// @Failure      500     {object}  ResponseMessage
// @Router       /api/v1/orders [post]
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

// FindAllOrders godoc
// @Summary      Get all orders
// @Description  Retrieve a list of all orders in the system
// @Tags         Orders
// @Accept       json
// @Produce      json
// @Param        query  query     find_all_orders.Input  false  "Query Parameters"
// @Success      200    {array}   find_all_orders.Order
// @Failure      400    {object}  ResponseMessage
// @Failure      500    {object}  ResponseMessage
// @Router       /api/v1/orders [get]
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

// UpdateOrder godoc
// @Summary      Update an existing order
// @Description  Update the details of an existing order. All the fields are optional.
// @Tags         Orders
// @Accept       json
// @Produce      json
// @Param        id      path      string              true  "Order ID"
// @Param        order   body      update_order.Input  true  "Updated Order Data"
// @Success      200     {object}  ResponseMessage
// @Failure      400     {object}  ResponseMessage
// @Failure      500     {object}  ResponseMessage
// @Router       /api/v1/orders/{id} [patch]
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

// UpdateOrderStatus godoc
// @Summary      Update status of an existing order
// @Description  Update the status of an existing order.
// @Tags         Orders
// @Accept       json
// @Produce      json
// @Param        id       path      string  true  "Order ID"
// @Param        status   body      update_order_status.Input  true  "Updated Order Status Data"
// @Success      200     {object}  ResponseMessage
// @Failure      400     {object}  ResponseMessage
// @Failure      500     {object}  ResponseMessage
// @Router       /api/v1/orders/{id}/status [patch]
func (h *OrderHandler) UpdateOrderStatus(c *gin.Context) {
	var input update_order_status.Input
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

	if err := h.updateOrderStatusUseCase.Execute(ctx, input); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "order status has been updated successfully",
	})
}
