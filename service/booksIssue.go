package service

import (
	"database/sql"
	"time"
)

type BookService struct {
	DB *sql.DB
}

//IssueBook function is declared to issue a book to user and update transaction table
func (s *BookService) IssueBook(userID, bookID int) error {
	timestamp := time.Now()
	tx, err := s.DB.Begin()
	if err != nil {
		return err
	}

	defer func() {
		if err != nil {
			tx.Rollback()
			return
		}
		err = tx.Commit()
	}()
	//Update Stock Count
	_, err = tx.Exec("UPDATE book SET stock=stock-1 WHERE id=$1", bookID)
	if err != nil {
		return err
	}
	//Insert entry into transaction table
	_, err = tx.Exec("INSERT INTO user_book_transaction(user_id,book_id,issued_date,return_date,fineamt,retun_status,actual_return_date) VALUES($1,$2,$3,false)", userID, bookID, timestamp)
	if err != nil {
		return err
	}
	return nil
}
func (s *BookService) ReturnBook(userID, bookID int) error {
	//start a transaction
	tx, err := s.DB.Begin()
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			tx.Rollback()
			return
		}
		err = tx.Commit()
	}()
	//Update Stock Count
	_, err = tx.Exec("UPDATE books SET stock=stock-1 WHERE id=$1", bookID)
	if err != nil {
		return err
	}
	//Update return status in transaction table
	_, err = tx.Exec("UPDATE transactions SET return_status=true where user_id=$1 and book_id=$1 AND return_status=false", userID, bookID)
	if err != nil {
		return err
	}
	return nil
}
