package main

import (
	"cars_api/controller"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	carController := controller.NewCarController()

	server.GET("/ping", func (ctx *gin.Context) {
		ctx.JSON(200, gin.H {
			"message": "TESTE",
		})
	})

	server.GET("/cars", carController.GetCar)

	server.Run(":3333")
}