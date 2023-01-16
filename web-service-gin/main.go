package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	uuid "github.com/satori/go.uuid"
)

type Car struct {
	ID       string  `json:"id" db:"car_id"`
	CarName  string  `json:"carname" db:"car_name"`
	CarType  string  `json:"cartype" db:"car_type"`
	Capacity string  `json:"capacity" db:"capacity"`
	Price    float64 `json:"price" db:"price"`
	GasType  string  `json:"gastype" db:"gas_type"`
}

func getallCars(c *gin.Context) {
	db := DBConnect()
	defer db.Close()
	cars, err := getCarsfromDB()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.IndentedJSON(http.StatusOK, cars)
}
func getCarsfromDB() ([]*Car, error) {
	db := DBConnect()
	cars := make([]*Car, 0)
	err := db.Select(&cars, `SELECT car_id, car_name, car_type, capacity,price,gas_type FROM carrental.cars`)
	if err != nil {
		return nil, err
	}
	return cars, nil

}

func postCar(c *gin.Context) {
	var newCar Car
	db := DBConnect()
	defer db.Close()
	id := uuid.NewV4()
	newCar.ID = id.String()
	if err := c.BindJSON(&newCar); err != nil {
		return
	}
	err := postcarinDB(newCar)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.IndentedJSON(http.StatusCreated, newCar)
}
func postcarinDB(car Car) error {
	db := DBConnect()
	fmt.Println(car)
	query := `INSERT INTO carrental.cars (car_id, car_name, car_type, capacity,price,gas_type) VALUES (:car_id, :car_name, :car_type, :capacity,:price,:gas_type)`
	_, err := db.NamedExec(query, car)
	if err != nil {
		return err
	}
	return nil
}
func getCar(c *gin.Context) {
	id := c.Param("id")
	car, err := getcarfromDB(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.IndentedJSON(http.StatusOK, car)
}
func getcarfromDB(id string) ([]*Car, error) {
	db := DBConnect()
	defer db.Close()
	car := make([]*Car, 0)
	query := `SELECT car_id, car_name, car_type, capacity,price,gas_type FROM carrental.cars WHERE car_id= $1`
	err := db.Select(&car, query, id)
	if err != nil {
		return nil, err
	}
	return car, nil
}
func deleteCar(c *gin.Context) {
	i := c.Param("id")
	err := deletecarfromDB(i)
	if err != nil {
		c.JSON(http.StatusOK, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Car deleted!!"})
}
func deletecarfromDB(id string) error {
	db := DBConnect()
	defer db.Close()
	query := `DELETE FROM carrental.cars WHERE car_id=$1`
	_, err := db.Exec(query, id)
	if err != nil {
		return err
	}
	fmt.Println("Car Deleted")
	return nil
}

func main() {
	router := gin.Default()
	router.GET("/Cars", getallCars)
	router.POST("/Cars", postCar)
	router.GET("/Cars/:id", getCar)
	router.DELETE("/Cars/:id", deleteCar)
	router.Run("localhost:8080")
}

func DBConnect() *sqlx.DB {
	db, err := sqlx.Connect("postgres", "user=postgres password=cactus1470 host=localhost port=5432 dbname=carrental sslmode=disable")
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return db
}
