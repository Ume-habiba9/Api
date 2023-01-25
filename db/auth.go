package db

func LogInCheck(email, passward string) (User, error) {
	database := DBConnect()
	defer database.Close()
	var userdata User
	query := `SELECT * FROM carrental.users WHERE email=$1 AND passward=$2`
	err := database.Select(&userdata, query, email, passward)
	if err != nil {
		return userdata, err
	}
	return userdata, err
}
