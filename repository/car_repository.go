package repository

import (
	"cars_api/model"
	"database/sql"
	"fmt"
)

type CarRepository struct {
	connection *sql.DB
}

func NewCarRepository(connection *sql.DB) CarRepository {
	return CarRepository{
		connection: connection,
	}
}

func (pr *CarRepository) GetCar() ([]model.Cars, error) {
	query := "SELECT id, car_name, price FROM cars"
	rows, err := pr.connection.Query(query)
	if err != nil {
		fmt.Println(err)
		return []model.Cars{}, err
	}

	var carList []model.Cars
	var carObj model.Cars

	for rows.Next() {
		err = rows.Scan(
			&carObj.ID,
			&carObj.Name,
			&carObj.Price)

		if err != nil {
			fmt.Println(err)
			return []model.Cars{}, err
		}

		carList = append(carList, carObj)
	}

	rows.Close()

	return carList, nil
}
