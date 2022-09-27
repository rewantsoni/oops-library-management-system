package library

import (
	"errors"
	"oops-library-management/book"
)

type Library interface {
	AddBook(book.Book) error
	Search(bookName string) (book.Book, error)
}

type listLibrary struct {
	books []book.Book
}

func (l *listLibrary) Search(bookName string) (book.Book, error) {
	var newBooks []book.Book

	var found book.Book
	for _, myBook := range l.books {
		if myBook.GetName() == bookName {
			found = myBook
		} else {
			newBooks = append(newBooks, myBook)
		}
	}
	if found == nil {
		return nil, errors.New("no such book present")
	}
	l.books = newBooks
	return found, nil
}

func (l *listLibrary) AddBook(books book.Book) error {
	l.books = append(l.books, books)
	return nil
}

func NewLibrary() (Library, error) {
	return &listLibrary{}, nil
}
