package routes

import (
	"github.com/Pos-tech-FIAP-GO-HORSE/order-management/internal/handlers"
	"github.com/gin-gonic/gin"
)

func AddUseRoutes(app *gin.Engine, handler *handlers.UserHandler) {
	api := app.Group("/api")
	v1 := api.Group("/v1")

	v1.POST("/user", handler.CreateUser)
}
