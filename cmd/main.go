package main

import (
	"cars_api/controller"
	"cars_api/db"
	"cars_api/repository"
	"cars_api/usecase"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	dbConnection, err := db.ConnectDB()
	if err != nil{
		panic(err)
	}
	CarRepository := repository.NewCarRepository(dbConnection)

	CarUsecase := usecase.NewCarUseCase(CarRepository)

	carController := controller.NewCarController(CarUsecase)

	server.GET("/ping", func (ctx *gin.Context) {
		ctx.JSON(200, gin.H {
			"message": "TESTE",
		})
	})

	server.GET("/cars", carController.GetCar)

	server.Run(":3333")
}