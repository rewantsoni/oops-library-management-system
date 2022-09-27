package book

func CreateDummyBookOne() Book {
	book, _ := NewBook("one")
	return book
}

func CreateDummyBookTwo() Book {
	book, _ := NewBook("two")
	return book
}

func CreateDummyBookThree() Book {
	book, _ := NewBook("three")
	return book
}

func CreateListOfDummyBooks() []Book {
	var books []Book
	books = append(books, CreateDummyBookOne())
	books = append(books, CreateDummyBookTwo())
	books = append(books, CreateDummyBookThree())
	return books
}
