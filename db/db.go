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