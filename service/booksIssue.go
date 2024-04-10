package service

import (
	"database/sql"
	"lms/models"
	"time"
)

//IssueBook function is declared to issue a book to user and update transaction table
func IssueBook(userID, bookID int, db *sql.DB) error {

	//Update Stock Count
	_, err := db.Exec("UPDATE book SET stock_count=stock_count-1 WHERE id=$1", bookID)
	if err != nil {
		return err
	}
	//Insert entry into transaction table
	_, err = db.Exec("INSERT INTO user_book_transaction(user_id,book_id,issued_date,return_date) VALUES($1,$2,$3,$4)", userID, bookID, time.Now(), time.Now().AddDate(0, 0, 15))

	if err != nil {
		return err
	}
	return nil
}

func GetUserBooks(userId int, db *sql.DB) (*models.UserBookTransaction, error) {
	//todo
	return nil, nil
}

func ReturnBook(userID, bookID int, db *sql.DB) error {

	//Update Stock Count
	_, err := db.Exec("UPDATE books SET stock_count=stock_count+1 WHERE id=$1", bookID)
	if err != nil {
		return err
	}
	//Update return status in transaction table
	_, err = db.Exec("UPDATE transactions SET return_status=true, actual_return_date=NOW() where user_id=$1 and book_id=$2", userID, bookID)
	if err != nil {
		return err
	}
	return nil
}
