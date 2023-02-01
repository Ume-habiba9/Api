package db

import (
	"fmt"

	"github.com/Ume-habiba9/Api/constants"
	"github.com/gin-gonic/gin"
)

func GetCarsfromDB(userID string, c *gin.Context) ([]*Car, error) {
	db := DBConnect()
	cars := make([]*Car, 0)
	query := `SELECT car_id,user_id, car_name, car_type, capacity,price,gas_type FROM carrental.cars`
	role, _ := c.Get("role")
	fmt.Println("Role in db", role)
	if role == constants.ROLE_USER {
		err := db.Select(&cars, query+" WHERE user_id=$1", userID)
		if err != nil {
			return nil, err
		}
	} else {
		err := db.Select(&cars, query)
		if err != nil {
			return nil, err
		}
	}
	return cars, nil
}
func PostcarinDB(car Car) error {
	database := DBConnect()
	defer database.Close()
	query := `INSERT INTO carrental.cars (car_id,user_id, car_name, car_type, capacity,price,gas_type,steering) VALUES (:car_id,:user_id, :car_name, :car_type, :capacity,:price,:gas_type,:steering)`
	_, err := database.NamedExec(query, car)
	if err != nil {
		return err
	}
	return nil
}
func GetcarfromDB(id, userid string, c *gin.Context) ([]*Car, error) {
	database := DBConnect()
	defer database.Close()
	car := make([]*Car, 0)
	query := `SELECT car_id,user_id, car_name, car_type, capacity,price,gas_type,steering FROM carrental.cars WHERE car_id= $1`
	role, _ := c.Get("role")
	fmt.Println("ROLE IN GET CAR", role)
	if role == constants.ROLE_USER {
		err := database.Select(&car, query+"  AND user_id = $", id, userid)
		if err != nil {
			return nil, err
		}
	} else {
		err := database.Select(&car, query, id)
		if err != nil {
			return nil, err
		}

	}
	return car, nil
}
func DeletecarfromDB(id, userid string, c *gin.Context) error {
	database := DBConnect()
	defer database.Close()
	query := `DELETE FROM carrental.cars WHERE car_id=$1`
	role, _ := c.Get("role")
	if role == constants.ROLE_USER {
		_, err := database.Exec(query+"  AND user_id = $2", id, userid)
		if err != nil {
			return err
		}
	} else {
		_, err := database.Exec(query, id)
		if err != nil {
			return err
		}
	}
	return nil
}
func UpdatecarinDB(id, userid string, cardata Car, c *gin.Context) error {
	database := DBConnect()
	defer database.Close()
	role, _ := c.Get("role")
	query := `UPDATE carrental.cars SET car_id=:car_id,user_id=:user_id, car_name=:car_name,car_type=:car_type,capacity=:capacity,price=:price,gas_type=:gas_type, steering=:steering WHERE car_id=$1`
	if role == constants.ROLE_USER {
		_, err := database.NamedExec(query+" AND user_id=$2", cardata)
		if err != nil {
			return err
		}
	} else {
		_, err := database.NamedExec(query, cardata)
		if err != nil {
			return err
		}
	}
	return nil
}
