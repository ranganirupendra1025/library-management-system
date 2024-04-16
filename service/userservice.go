package service

import (
	"database/sql"
	"errors"
	"net/http"

	//"fmt"
	"lms/models"
	"lms/utils"
)

//I have added refernce code for getting user from databse  u guys can change as per the requirement

func GettingUsers(db *sql.DB) ([]models.User, error) {
	var users []models.User
	rows, err := db.Query("SELECT * FROM  users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Iterate through the result set
	for rows.Next() {
		var user models.User
		err := rows.Scan(&user.Id, &user.Username, &user.Age, &user.Email, &user.Password, &user.Isadmin, &user.Subid, &user.Subdate)
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
	query := "INSERT INTO users (username, age, email_address, password, is_admin) VALUES ($1, $2, $3, $4, $5)"
	password, err := utils.Hashpassword(newUser.Password)
	if err != nil {
		return err
	}
	_, err = db.Exec(query, newUser.Username, newUser.Age, newUser.Email, password, newUser.Isadmin)
	if err != nil {
		return err
	}
	return nil
}
func GetUser(db *sql.DB, id int) (*models.User, error) {
	var user models.User
	query := "SELECT  id,username,age,email_address,password,is_admin, coalesce(subscription_id, 0), subscription_end_date FROM users WHERE  id = $1"
	row := db.QueryRow(query, id)

	err := row.Scan(&user.Id, &user.Username, &user.Age, &user.Email, &user.Password, &user.Isadmin, &user.Subid, &user.Subdate)
	if err != nil {
		return nil, err
	}
	return &user, nil

}
func DeleteUser(db *sql.DB, user *models.User) error {
	query := "DELETE FROM users WHERE id=$1"
	_, err := db.Exec(query, user.Id)
	if err != nil {
		return err
	}
	return nil
}

func UpdateSubscription(db *sql.DB, user models.User) error {
	query := "UPDATE users set subscription_id=$1 ,subscription_end_date=$2 where id=$3"
	_, err := db.Exec(query, user.Subid, user.Subdate, user.Id)
	if err != nil {
		return err
	}
	//fmt.Println(rows)
	return nil
}
func Login(db *sql.DB, user *models.Login) error {
	query := "SELECT id,password from users where email_address=$1"
	row := db.QueryRow(query, &user.Email)
	var retrivedpass string
	err := row.Scan(&user.Id, &retrivedpass)
	if err != nil {
		return err
	}
	passwordValid := utils.Checkpassword(user.Password, retrivedpass)
	if !passwordValid {
		return errors.New("credentials invalid")
	}
	return nil

}

func ValidateAdmin(r *http.Request, w http.ResponseWriter, db *sql.DB, subDto models.SubscriptionDto) bool {

	user, err := GetUser(db, subDto.AdminId)
	if err != nil {
		http.Error(w, "No user with that name", http.StatusInternalServerError)
		return false
	}
	if !user.Isadmin {
		http.Error(w, "Only Admin user can access", http.StatusBadRequest)
		return false

	}
	return true
}
