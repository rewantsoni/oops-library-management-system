package library_card

import "errors"

type LibraryCard interface {
	GetCardLimit() int
}

type libraryCardWithNumber struct {
	number int
	limit  int
}

func (l libraryCardWithNumber) GetCardLimit() int {
	return l.limit
}

func NewLibraryCard(number int) (LibraryCard, error) {
	if number < 1000 {
		return nil, errors.New("library card number has to of at least 4 digits")
	}
	return libraryCardWithNumber{number: number, limit: 1}, nil
}
