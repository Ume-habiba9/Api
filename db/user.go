package db
import (
	"fmt"
)
func Postuserindb(user User) error {
	database := DBConnect()
	defer database.Close()
	fmt.Println(user)
	query := `INSERT INTO carrental.users (user_id, user_name, email, passward) VALUES (:user_id, :user_name, :email, :passward)`
	_, err := database.NamedExec(query, user)
	if err != nil {
		return err
	}
	return nil
}
func Getusersfromdb() ([]*User, error) {
	db := DBConnect()
	users := make([]*User, 0)
	err := db.Select(&users, `SELECT user_id, user_name, email, passward FROM carrental.users`)
	if err != nil {
		return nil, err
	}
	return users, nil
}
func Getuserfromdb(id string) ([]*User, error) {
	database := DBConnect()
	defer database.Close()
	user := make([]*User, 0)
	query := `SELECT user_id , user_name,email,passward FROM carrental.users WHERE user_id= $1`
	err := database.Select(&user, query, id)
	if err != nil {
		return nil, err
	}
	return user, nil
}
func Deleteuserfromdb(id string) error {
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
func Updateuserindb(id string, userdata User) error {
	database := DBConnect()
	defer database.Close()
	query := `UPDATE carrental.users SET user_id=:user_id,user_name=:user_name,email=:email,passward=:passward WHERE user_id=$1`
	_, err := database.NamedExec(query, userdata)
	if err != nil {
		return err
	}
	return nil
}