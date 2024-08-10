package routes

import (
	"github.com/Pos-tech-FIAP-GO-HORSE/order-management/internal/handlers"
	"github.com/gin-gonic/gin"
)

func AddPaymentRoutes(app *gin.Engine, handler *handlers.PaymentHandler) {
	api := app.Group("/api")
	v1 := api.Group("/v1")
	{
		v1.POST("/payment", handler.ProcessPayment)

	}
}
