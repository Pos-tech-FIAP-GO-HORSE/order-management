package routes

import (
	"github.com/Pos-tech-FIAP-GO-HORSE/order-management/internal/handlers"
	"github.com/gin-gonic/gin"
)

func AddOrdersRoutes(app *gin.Engine, handler *handlers.OrderHandler) {
	api := app.Group("/api")
	v1 := api.Group("/v1")
	{
		v1.POST("/orders", handler.CreateOrder)
	}
}
