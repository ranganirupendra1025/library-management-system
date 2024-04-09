package models

import "time"

type UserBookTransaction struct {
	Id               int
	UserId           int
	BookId           int
	IssueDate        time.Time
	ReturnDate       time.Time
	FineAmount       int
	ReturnStatus     bool
	ActualReturnDate time.Time
}
