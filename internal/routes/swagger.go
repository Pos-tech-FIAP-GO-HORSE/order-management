package routes

import (
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

func AddSwaggerRoute(app *gin.Engine) {
	api := app.Group("/swagger")
	{
		api.GET("/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}
}
