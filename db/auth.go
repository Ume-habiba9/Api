package db

func LogInCheck(email, passward string) ([]*User, error) {
	database := DBConnect()
	defer database.Close()
	userData := make([]*User, 0)
	query := `SELECT user_id,email,passward FROM carrental.users WHERE email=$1 AND passward=$2`
	err := database.Select(&userData, query ,email, passward)
	if err != nil {
		return nil, err
	}
	return userData, err
}