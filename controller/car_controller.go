package controller

import (
	"cars_api/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type carController struct {
	CarUsecase usecase.CarUsecase
}

func NewCarController(usecase usecase.CarUsecase) carController {
	return carController{
		CarUsecase: usecase,
	}
}

func (p *carController) GetCar(ctx *gin.Context) {
	cars, err := p.CarUsecase.GetCar()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
	}

	ctx.JSON(http.StatusOK, cars)
}