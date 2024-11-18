package usecase

import (
	"cars_api/model"
	"cars_api/repository"
)

type CarUsecase struct {
	repository repository.CarRepository
}

func NewCarUseCase(repo repository.CarRepository) CarUsecase {
	return CarUsecase{
		repository: repo,
	}
}

func (pu *CarUsecase) GetCar() ([]model.Cars, error){
	return pu.repository.GetCar()
}