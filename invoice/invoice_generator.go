package invoice

import (
	"oops-library-management/amount"
	"oops-library-management/book"
	"time"
)

type InvoiceGenerator interface {
	Generate(book.Book, time.Time, amount.Amount) (Invoice, error)
}

type SimpleInvoiceGenerator struct {
}

func (s SimpleInvoiceGenerator) Generate(b book.Book, t time.Time, a amount.Amount) (Invoice, error) {
	return NewInvoice(b, t, a)
}

func NewInvoiceGenerator() (InvoiceGenerator, error) {
	return SimpleInvoiceGenerator{}, nil
}
