package book

import "errors"

type Book interface {
	GetName() string
}

type paperBackBook struct {
	name string
}

func (p paperBackBook) GetName() string {
	return p.name
}

func NewBook(name string) (Book, error) {
	if len(name) < 1 {
		return nil, errors.New("book name cannot be empty")
	}
	return paperBackBook{name: name}, nil
}
