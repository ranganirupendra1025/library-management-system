package service

import (
	"database/sql"
	"lms/models"
)

//I have added refernce code for getting user from databse  u guys can change as per the requirement

func GettingUsers(db *sql.DB) ([]models.User, error) {
	var users []models.User
	rows, err := db.Query("SELECT id,username,age,email_address,password,is_admin FROM  users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Iterate through the result set
	for rows.Next() {
		var user models.User
		err := rows.Scan(&user.Userid, &user.Username, &user.Age, &user.Email, &user.Password, &user.Isadmin)
		if err != nil {
			return nil, err
		}
		users = append(users, user)

		//fmt.Println(users)
		// fmt.Println("ID:%d,name:%s,descr:%s,location:%s",user.Id,user.Name,user.Description,user.Location )
	}
	return users, nil

}

func RegisterUser(db *sql.DB, newUser models.User) error {
	// Insert the new user into the database
	_, err := db.Exec("INSERT INTO users (username, age, email_address, password, is_admin,subscription_id, subscription_end_date) VALUES ($1, $2, $3, $4, $5, $6, $7)", newUser.Username, newUser.Age, newUser.Email, newUser.Password, newUser.Isadmin, newUser.Subid, newUser.Subdate)
	if err != nil {
		return err
	}
	return nil
}
func GetUser(db *sql.DB, id int) (*models.User, error) {
	var user models.User
	query := "SELECT  * FROM users WHERE  id = $1"
	row := db.QueryRow(query, id)

	err := row.Scan(&user.Userid, &user.Username, &user.Age, &user.Email, &user.Password, &user.Isadmin, &user.Subid, &user.Subdate)
	if err != nil {
		return nil, err
	}
	return &user, nil

}
func DeleteUser(db *sql.DB, user *models.User) error {
	query := "DELETE FROM users WHERE id=$1"
	_, err := db.Exec(query, user.Userid)
	if err != nil {
		return err
	}
	return nil
}

func UpdateSubscription(db *sql.DB, user models.User) error {
	query := "update users set subscription_id=$1, subscription_end_date=$2 where id=$1"
	_, err := db.Exec(query, user.Subid, user.Subdate, user.Userid)
	if err != nil {
		return err
	}
	return nil
}
