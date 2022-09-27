package librarian

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"oops-library-management/book"
	"oops-library-management/library"
	"testing"
)

func TestCreateNewLibrarian(t *testing.T) {
	testCases := []struct {
		name           string
		param          library.Library
		expectedResult Librarian
		expectedError  error
	}{
		{
			name:           "should create a new librarian",
			param:          library.CreateDummyLibrary(),
			expectedResult: &NoviceLibrarian{library: library.CreateDummyLibrary()},
		},
		{
			name:          "should create a new librarian",
			param:         nil,
			expectedError: errors.New("library cannot be nil"),
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			res, err := NewLibrarian(testCase.param)
			assert.Equal(t, err, testCase.expectedError)
			if err == nil {
				assert.Equal(t, res, testCase.expectedResult)
			}
		})
	}
}

func TestSearch(t *testing.T) {
	testCases := []struct {
		name           string
		param          string
		expectedResult book.Book
		expectedError  error
	}{
		{
			name:           "should return the book when present",
			param:          "one",
			expectedResult: book.CreateDummyBookOne(),
		},
		{
			name:          "should return error when book is not present",
			param:         "five",
			expectedError: errors.New("no such book present"),
		},
	}

	mockLibrary := &library.MockLibrary{}
	mockLibrary.On("Search", "one").Return(book.CreateDummyBookOne(), nil)
	//TODO: create an empty book and return it here
	mockLibrary.On("Search", "five").Return(book.CreateDummyBookOne(), errors.New("no such book present"))
	librarian, err := NewLibrarian(mockLibrary)
	assert.NoError(t, err)
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			res, err := librarian.Search(testCase.param)
			assert.Equal(t, err, testCase.expectedError)
			if err == nil {
				assert.Equal(t, res, testCase.expectedResult)
			}
		})
	}
}

func TestReturn(t *testing.T) {
	testCases := []struct {
		name          string
		param         book.Book
		expectedError error
	}{
		{
			name:  "should put the book back in the library",
			param: book.CreateDummyBookOne(),
		},
	}

	mockLibrary := &library.MockLibrary{}
	mockLibrary.On("AddBook", book.CreateDummyBookOne()).Return(nil)
	librarian, err := NewLibrarian(mockLibrary)
	assert.NoError(t, err)

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			err := librarian.Return(testCase.param)
			assert.Equal(t, err, testCase.expectedError)
		})
	}
}
