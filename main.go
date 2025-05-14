package main

import (
	"qris-api/config"
	"qris-api/model"
	"qris-api/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	config.ConnectDB()
	config.DB.AutoMigrate(&model.QRISRequestModel{})

	router := gin.Default()
	routes.RegisterRoutes(router)

	router.Run(":8080")
}
