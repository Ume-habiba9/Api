package db

import (
	"fmt"
)

func LogInCheck(id, email, passward string) ([]*User, error) {
	database := DBConnect()
	defer database.Close()
	user := make([]*User, 0)
	query := `SELECT user_id,email,passward FROM carrental.users WHERE user_id=$1 AND email=$2 AND passward=$3`
	err := database.Select(&user, query, id, email, passward)
	fmt.Println(err)
	if err != nil {
		return nil, err
	}
	return user, err
}


