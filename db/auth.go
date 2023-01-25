package db

import "fmt"

func LogInCheck(email, passward string) (*User, error) {
	database := DBConnect()
	defer database.Close()
	var userdata User
	query := `SELECT * FROM carrental.users WHERE email=$1 AND passward=$2`
	err := database.Get(&userdata, query, email, passward)
	if err != nil {
		fmt.Println("here ", err)
		return nil, err
	}
	return &userdata, err
}
