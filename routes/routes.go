package routes

import (
	"qris-api/controller"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine){
	router.POST("/qris", controller.CreateQRIS)
}

