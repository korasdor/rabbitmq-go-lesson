package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/korasdor/rabbitmq-go/producer/controllers"
)

func SetupRouters(app *gin.Engine) {
	v1 := app.Group("/v1")
	{
		v1.GET("ping", controllers.Ping)
		v1.POST("example", controllers.Example)
	}
}
