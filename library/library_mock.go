package library

import (
	"github.com/stretchr/testify/mock"
	"oops-library-management/book"
)

type MockLibrary struct {
	mock.Mock
}

func (m *MockLibrary) AddBook(book book.Book) error {
	args := m.Called(book)
	return args.Error(0)
}

func (m *MockLibrary) Search(bookName string) (book.Book, error) {
	args := m.Called(bookName)
	return args.Get(0).(book.Book), args.Error(1)
}
