package library

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"oops-library-management/book"
	"testing"
)

func TestCreateNewLibrary(t *testing.T) {

	testCases := []struct {
		name           string
		expectedResult Library
		expectedError  error
	}{
		{
			name:           "should create a new library",
			expectedResult: &listLibrary{},
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			res, err := NewLibrary()
			assert.Equal(t, err, testCase.expectedError)
			if err == nil {
				assert.Equal(t, res, testCase.expectedResult)
			}
		})
	}
}

func TestAddBooks(t *testing.T) {

	testCases := []struct {
		name          string
		params        book.Book
		expectedError error
	}{
		{
			name:          "should add books to the library",
			params:        book.CreateDummyBookOne(),
			expectedError: nil,
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			library, err := NewLibrary()
			assert.NoError(t, err)
			err = library.AddBook(testCase.params)
			assert.NoError(t, err)
		})
	}
}

func TestSearchBook(t *testing.T) {

	testCases := []struct {
		name           string
		param          string
		expectedResult book.Book
		expectedError  error
	}{
		{
			name:           "should return Book if book is present",
			param:          "two",
			expectedResult: book.CreateDummyBookTwo(),
		},
		{
			name:          "should return error if book is not present",
			param:         "five",
			expectedError: errors.New("no such book present"),
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			library, err := NewLibrary()
			assert.NoError(t, err)
			err = library.AddBook(book.CreateDummyBookOne())
			assert.NoError(t, err)
			err = library.AddBook(book.CreateDummyBookTwo())
			assert.NoError(t, err)
			err = library.AddBook(book.CreateDummyBookThree())
			assert.NoError(t, err)
			res, err := library.Search(testCase.param)
			assert.Equal(t, err, testCase.expectedError)
			if err == nil {
				assert.Equal(t, res, testCase.expectedResult)
			}
		})
	}
}
