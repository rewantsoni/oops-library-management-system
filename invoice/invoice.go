package invoice

import (
	"oops-library-management/amount"
	"oops-library-management/book"
	"time"
)

type Invoice struct {
	book       book.Book
	issuedDate time.Time
	amount     amount.Amount
}

func NewInvoice(b book.Book, date time.Time, a amount.Amount) (Invoice, error) {
	return Invoice{book: b, issuedDate: date, amount: a}, nil
}
