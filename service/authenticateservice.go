package service

import "database/sql"

func AuthenticateUser(db *sql.DB, username, password string) (bool, error) {
	var count int
	err := db.QueryRow("SELECT COUNT(*) FROM users WHERE username = $1 AND password = $2", username, password).Scan(&count)
	if err != nil {
		return false, err
	}
	return count > 0, nil
}
