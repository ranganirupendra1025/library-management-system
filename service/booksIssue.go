package service

import (
	"database/sql"
)

//IssueBook function is declared to issue a book to user and update transaction table
func IssueBook(userID, bookID int, db *sql.DB) error {

	//Update Stock Count
	_, err := db.Exec("UPDATE book SET stock=stock-1 WHERE id=$1", bookID)
	if err != nil {
		return err
	}
	//Insert entry into transaction table
	_, err = db.Exec("INSERT INTO user_book_transaction(user_id,book_id,issued_date,return_date,fineamt,retun_status,actual_return_date) VALUES($1,$2,$3,false)", userID, bookID, timestamp)

	if err != nil {
		return err
	}
	return nil
}
func ReturnBook(userID, bookID int, db *sql.DB) error {

	//Update Stock Count
	_, err := db.Exec("UPDATE books SET stock=stock-1 WHERE id=$1", bookID)
	if err != nil {
		return err
	}
	//Update return status in transaction table
	_, err = db.Exec("UPDATE transactions SET return_status=true where user_id=$1 and book_id=$1 AND return_status=false", userID, bookID)
	if err != nil {
		return err
	}
	return nil
}
