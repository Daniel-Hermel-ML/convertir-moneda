package template

import (
	"github.com/gin-gonic/gin"
	"github.com/melisource/fury_convertir-divisas/cmd/src/controller"
)

func SetupRoutes(router *gin.Engine) {
	v1 := router.Group("/V1")
	{
		v1.GET("/convert/real-to-dollar", controller.GetDollar)
		v1.GET("/convert/real-to-euro", controller.GetEuro)
		v1.GET("/convert/real-to-peso", controller.GetPeso)
		v1.GET("/ping", controller.PongFunction)
	}
}
