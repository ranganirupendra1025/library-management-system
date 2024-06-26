package service

import (
	"database/sql"
	"errors"
	"lms/models"
	"lms/utils"
	"time"
)

//IssueBook function is declared to issue a book to user and update transaction table
func IssueBook(userID, bookID int, db *sql.DB) error {
	userOverDueBooks, err := GetUserOverdueBooks(userID, db)
	if err == nil && userOverDueBooks != nil {
		return errors.New("There are overdue books which are not yet returned by this user. Can't issue a new book.")
	}

	//Update Stock Count
	_, err = db.Exec("UPDATE book SET stock_count=stock_count-1 WHERE id=$1", bookID)

	if err != nil {
		return err
	}
	//Insert entry into transaction table
	_, err = db.Exec("INSERT INTO user_book_transaction(user_id,book_id,issued_date,return_date) VALUES($1,$2,$3,$4)", userID, bookID, time.Now(), time.Now().AddDate(0, 0, utils.BookBorrowPeriodInDays))
	if err != nil {
		return err
	}
	return nil
}

func GetUserBooks(userId int, db *sql.DB) (*models.UserBookTransaction, error) {
	//todo - Renewal and (Go Routine to calculate fineAmount periodically/ calculate fineamount and update in DB while returning userBook books)
	//todo - when adding and renewing subscription, find a way to update the subscription end date
	var userBook models.UserBookTransaction
	// It shows error if we get all records becuase actual return date was null,so i removed that..
	query := "SELECT id,user_id,book_id,issued_date,return_date,fine_amount,return_status FROM user_book_transaction WHERE  user_id = $1"
	//query := "SELECT * FROM user_book_transaction WHERE  user_id = $1"

	row := db.QueryRow(query, userId)

	err := row.Scan(&userBook.Id, &userBook.UserId, &userBook.BookId, &userBook.IssueDate, &userBook.ReturnDate, &userBook.FineAmount, &userBook.ReturnStatus)
	if err != nil {
		return nil, err
	}
	return &userBook, nil
}

func GetUserPendingBooks(userId int, db *sql.DB) (*models.UserBookTransaction, error) {
	//todo - Renewal and (Go Routine to calculate fineAmount periodically/ calculate fineamount and update in DB while returning userBook books)
	//todo - when adding and renewing subscription, find a way to update the subscription end date
	var userBook models.UserBookTransaction
	query := "SELECT id,user_id,book_id,issued_date,return_date,fine_amount,return_status  FROM user_book_transaction WHERE  user_id = $1 and return_status=false"
	row := db.QueryRow(query, userId)

	err := row.Scan(&userBook.Id, &userBook.UserId, &userBook.BookId, &userBook.IssueDate, &userBook.ReturnDate, &userBook.FineAmount, &userBook.ReturnStatus)
	if err != nil {
		return nil, err
	}
	return &userBook, nil
}

func GetUserOverdueBooks(userId int, db *sql.DB) (*models.UserBookTransaction, error) {
	//todo - Renewal and (Go Routine to calculate fineAmount periodically/ calculate fineamount and update in DB while returning userBook books)
	//todo - when adding and renewing subscription, find a way to update the subscription end date
	var userBook models.UserBookTransaction
	query := "SELECT id, user_id,book_id,issued_date,return_date,fine_amount,return_status  FROM user_book_transaction WHERE  user_id = $1 and return_status=false and return_date <= NOW()"
	row := db.QueryRow(query, userId)

	err := row.Scan(&userBook.Id, &userBook.UserId, &userBook.BookId, &userBook.IssueDate, &userBook.ReturnDate, &userBook.FineAmount, &userBook.ReturnStatus)
	if err != nil {
		return nil, err
	}
	return &userBook, nil
}

func ReturnBook(userID, bookID int, db *sql.DB) error {

	//Update Stock Count
	_, err := db.Exec("UPDATE book SET stock_count=stock_count+1 WHERE id=$1", bookID)
	if err != nil {
		return err
	}
	//Update return status in transaction table
	_, err = db.Exec("UPDATE user_book_transaction SET return_status=true, actual_return_date=NOW() where user_id=$1 and book_id=$2", userID, bookID)
	if err != nil {
		return err
	}
	return nil
}

func GetUserBook(userID, bookID int, db *sql.DB) (*models.UserBookTransaction, error) {

	var userBook models.UserBookTransaction
	query := "SELECT return_date FROM user_book_transaction WHERE user_id = $1 and book_id = $2 and return_status=false"
	row := db.QueryRow(query, userID, bookID)

	err := row.Scan(&userBook.ReturnDate)
	if err != nil {
		return nil, err
	}
	return &userBook, nil
}
