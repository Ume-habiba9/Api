package db

import (
	"fmt"
)

func PostUserinDB(user User) error {
	database := DBConnect()
	defer database.Close()
	query := `INSERT INTO carrental.users (user_id, user_name, email, passward,role) VALUES (:user_id, :user_name, :email, :passward,:role)`
	_, err := database.NamedExec(query, user)
	if err != nil {
		return err
	}
	return nil
}
func GetUsersfromDB() ([]*User, error) {
	db := DBConnect()
	users := make([]*User, 0)
	err := db.Select(&users, `SELECT user_id, user_name, email FROM carrental.users`)
	if err != nil {
		return nil, err
	}
	return users, nil
}
func GetUserfromDB(id string) ([]*User, error) {
	database := DBConnect()
	defer database.Close()
	user := make([]*User, 0)
	query := `SELECT user_id , user_name,email FROM carrental.users WHERE user_id= $1`
	err := database.Select(&user, query, id)
	if err != nil {
		return nil, err
	}
	return user, nil
}
func DeleteUserfromDB(id string) error {
	database := DBConnect()
	defer database.Close()
	query := `DELETE FROM carrental.users WHERE user_id=$1`
	_, err := database.Exec(query, id)
	if err != nil {
		return err
	}
	fmt.Println("User Deleted")
	return nil
}
func UpdateUserinDB(id string, userdata User) error {
	database := DBConnect()
	defer database.Close()
	query := `UPDATE carrental.users SET user_id=:user_id,user_name=:user_name,email=:email,passward=:passward WHERE user_id=$1`
	_, err := database.NamedExec(query, userdata)
	if err != nil {
		return err
	}
	return nil
}
