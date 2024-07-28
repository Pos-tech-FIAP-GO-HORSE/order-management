package routes

import (
	"github.com/Pos-tech-FIAP-GO-HORSE/order-management/internal/handlers"
	"github.com/gin-gonic/gin"
)

func AddProductsRoutes(app *gin.Engine, handler *handlers.ProductHandler) {
	api := app.Group("/api")
	v1 := api.Group("/v1")
	{
		v1.POST("/products", handler.CreateProduct)
		v1.GET("/products", handler.FindAllProducts)
		v1.GET("/products/:id", handler.FindProductByID)
		v1.GET("/products/:category", handler.FindProductByCategory)
		v1.PATCH("/products/:id", handler.UpdateProduct)
		v1.PATCH("/products/toggle/:id/", handler.UpdateProductAvalability)
		v1.DELETE("/products/:id", handler.DeleteProduct)
	}
}
