package librarian

import (
	"errors"
	"oops-library-management/book"
	"oops-library-management/library"
)

type Librarian interface {
	Search(bookName string) (book.Book, error)
	Return(book book.Book) error
}

type NoviceLibrarian struct {
	library library.Library
}

func (n *NoviceLibrarian) Return(book book.Book) error {
	return n.library.AddBook(book)
}

func (n *NoviceLibrarian) Search(bookName string) (book.Book, error) {
	return n.library.Search(bookName)
}

func NewLibrarian(l library.Library) (Librarian, error) {
	if l == nil {
		return nil, errors.New("library cannot be nil")
	}
	return &NoviceLibrarian{library: l}, nil
}
