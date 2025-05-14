package main

import (
	"qris-api/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	routes.RegisterRoutes(router)

	router.Run(":8080")
}
