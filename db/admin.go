package db

func PostAdmininDB(admin User) error {
	database := DBConnect()
	defer database.Close()
	query := `INSERT INTO carrental.users (user_id, user_name, email, passward,role) VALUES (:user_id, :user_name, :email, :passward,:role)`
	_, err := database.NamedExec(query, admin)
	if err != nil {
		return err
	}
	return nil
}

func GetCarsbyAdmin() ([]*Car, error) {
	db := DBConnect()
	cars := make([]*Car, 0)
	err := db.Select(&cars, `SELECT car_id,user_id, car_name, car_type, capacity,price,gas_type FROM carrental.cars`)
	if err != nil {
		return nil, err
	}
	return cars, nil
}
func GetcarbyAdmin(id string) ([]*Car, error) {
	database := DBConnect()
	defer database.Close()
	car := make([]*Car, 0)
	query := `SELECT car_id,user_id, car_name, car_type, capacity,price,gas_type,steering FROM carrental.cars WHERE car_id= $1`
	err := database.Select(&car, query, id)
	if err != nil {
		return nil, err
	}
	return car, nil
}
func DeleteCarbyAdmin(id string) error {
	database := DBConnect()
	defer database.Close()
	query := `DELETE FROM carrental.cars WHERE car_id=$1`
	_, err := database.Exec(query, id)
	if err != nil {
		return err
	}
	return nil
}
func UpdateCarbyAdmin(id string, cardata Car) error {
	database := DBConnect()
	defer database.Close()
	query := `UPDATE carrental.cars SET car_id=:car_id,user_id=:user_id, car_name=:car_name,car_type=:car_type,capacity=:capacity,price=:price,gas_type=:gas_type, steering=:steering WHERE car_id=$1`
	_, err := database.NamedExec(query, cardata)
	if err != nil {
		return err
	}
	return nil
}
