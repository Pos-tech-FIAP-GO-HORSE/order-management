package routes

import (
	_ "github.com/Pos-tech-FIAP-GO-HORSE/order-management/cmd/docs"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func AddSwaggerRoute(app *gin.Engine) {
	api := app.Group("/swagger")
	{
		api.GET("/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}
}
