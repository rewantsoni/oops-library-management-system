package library_management_system

import (
	"oops-library-management/book"
	"oops-library-management/invoice"
	"oops-library-management/librarian"
	"oops-library-management/library_card"
	"oops-library-management/payment"
)

type LibraryManagementSystem interface {
	BorrowBook(string, library_card.LibraryCard) (book.Book, error)
	Return(book.Book, library_card.LibraryCard) (invoice.Invoice, error)
}

type simpleLibraryManagementSystem struct {
	issuedBookStore  IssuedBookStore
	paymentSystem    payment.PaymentSystem
	librarian        librarian.Librarian
	invoiceGenerator invoice.InvoiceGenerator
}

func (s *simpleLibraryManagementSystem) BorrowBook(bookName string, card library_card.LibraryCard) (book.Book, error) {
	book, err := s.librarian.Search(bookName)
	if err != nil {
		return nil, err
	}

	err = s.issuedBookStore.Add(book, card)
	if err != nil {
		return nil, err
	}

	return book, err

}

func (s *simpleLibraryManagementSystem) Return(b book.Book, card library_card.LibraryCard) (invoice.Invoice, error) {
	issueDetails, err := s.issuedBookStore.Remove(b, card)
	if err != nil {
		return invoice.Invoice{}, err
	}

	err = s.librarian.Return(b)
	if err != nil {
		return invoice.Invoice{}, err
	}
	amount, err := s.paymentSystem.Compute(issueDetails.issuedDate)
	if err != nil {
		return invoice.Invoice{}, err
	}
	return s.invoiceGenerator.Generate(issueDetails.book, issueDetails.issuedDate, amount)
}

func NewLibraryManagementSystem(librarian librarian.Librarian, store IssuedBookStore, paymentSystem payment.PaymentSystem, invoiceGenerator invoice.InvoiceGenerator) (LibraryManagementSystem, error) {
	return &simpleLibraryManagementSystem{
		issuedBookStore:  store,
		paymentSystem:    paymentSystem,
		librarian:        librarian,
		invoiceGenerator: invoiceGenerator,
	}, nil
}
