package db

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Car struct {
	ID       string  `json:"id" db:"car_id"`
	CarName  string  `json:"carname" db:"car_name"`
	CarType  string  `json:"cartype" db:"car_type"`
	Capacity string  `json:"capacity" db:"capacity"`
	Price    float64 `json:"price" db:"price"`
	GasType  string  `json:"gastype" db:"gas_type"`
}

func DBConnect() *sqlx.DB {
	db, err := sqlx.Connect("postgres", "user=postgres password=cactus1470 host=localhost port=5432 dbname=carrental sslmode=disable")
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return db
}
func GetCarsfromDB() ([]*Car, error) {
	db := DBConnect()
	cars := make([]*Car, 0)
	err := db.Select(&cars, `SELECT car_id, car_name, car_type, capacity,price,gas_type FROM carrental.cars`)
	if err != nil {
		return nil, err
	}
	return cars, nil
}
func PostcarinDB(car Car) error {
	database := DBConnect()
	defer database.Close()
	fmt.Println(car)
	query := `INSERT INTO carrental.cars (car_id, car_name, car_type, capacity,price,gas_type) VALUES (:car_id, :car_name, :car_type, :capacity,:price,:gas_type)`
	_, err := database.NamedExec(query, car)
	if err != nil {
		return err
	}
	return nil
}
func GetcarfromDB(id string) ([]*Car, error) {
	database := DBConnect()
	defer database.Close()
	car := make([]*Car, 0)
	query := `SELECT car_id, car_name, car_type, capacity,price,gas_type FROM carrental.cars WHERE car_id= $1`
	err := database.Select(&car, query, id)
	if err != nil {
		return nil, err
	}
	return car, nil
}
func DeletecarfromDB(id string) error {
	database := DBConnect()
	defer database.Close()
	query := `DELETE FROM carrental.cars WHERE car_id=$1`
	_, err := database.Exec(query, id)
	if err != nil {
		return err
	}
	fmt.Println("Car Deleted")
	return nil
}
func UpdatecarinDB(id string, cardata Car) error {
	database := DBConnect()
	defer database.Close()
	query := `UPDATE carrental.cars SET car_id=:car_id,car_name=:car_name,car_type=:car_type,capacity=:capacity,price=:price,gas_type=:gas_type WHERE car_id=$1`
	_, err := database.NamedExec(query, cardata)
	if err != nil {
		return err
	}
	return nil
}
