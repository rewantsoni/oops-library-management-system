package library

import "oops-library-management/book"

func CreateDummyLibrary() Library {
	library, _ := NewLibrary()
	library.AddBook(book.CreateDummyBookOne())
	library.AddBook(book.CreateDummyBookTwo())
	library.AddBook(book.CreateDummyBookThree())
	return library
}
