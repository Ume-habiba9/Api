package db

import "github.com/gin-gonic/gin"

func GetCarsfromDB(userID string, c gin.Context) ([]*Car, error) {
	db := DBConnect()
	cars := make([]*Car, 0)
	query := `SELECT car_id,user_id, car_name, car_type, capacity,price,gas_type FROM carrental.cars`
	role, _ := c.Get("role")
	if role= {
		
	}
	err := db.Select(&cars, userID)
	if err != nil {
		return nil, err
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
func GetcarfromDB(id, userid string) ([]*Car, error) {
	database := DBConnect()
	defer database.Close()
	car := make([]*Car, 0)
	query := `SELECT car_id,user_id, car_name, car_type, capacity,price,gas_type,steering FROM carrental.cars WHERE car_id= $1 AND user_id = $2`
	err := database.Select(&car, query, id, userid)
	if err != nil {
		return nil, err
	}
	return car, nil
}
func DeletecarfromDB(id, userid string) error {
	database := DBConnect()
	defer database.Close()
	query := `DELETE FROM carrental.cars WHERE car_id=$1 AND user_id = $2`
	_, err := database.Exec(query, id, userid)
	if err != nil {
		return err
	}
	return nil
}
func UpdatecarinDB(id, userid string, cardata Car) error {
	database := DBConnect()
	defer database.Close()
	query := `UPDATE carrental.cars SET car_id=:car_id,user_id=:user_id, car_name=:car_name,car_type=:car_type,capacity=:capacity,price=:price,gas_type=:gas_type, steering=:steering WHERE car_id=$1 AND user_id=$2`
	_, err := database.NamedExec(query, cardata)
	if err != nil {
		return err
	}
	return nil
}
