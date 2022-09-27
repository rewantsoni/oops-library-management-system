package library_management_system

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"oops-library-management/book"
	"oops-library-management/library_card"
	"testing"
)

func TestCreateNewStore(t *testing.T) {
	testCases := []struct {
		name           string
		expectedResult IssuedBookStore
		expectedError  error
	}{
		{
			name:           "should create a new issued book store",
			expectedResult: &inMemoryIssuedBookStore{},
			expectedError:  nil,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			res, err := NewIssuedBookStore()
			assert.Equal(t, err, testCase.expectedError)
			if err == nil {
				assert.Equal(t, res, testCase.expectedResult)
			}
		})
	}
}

func TestAddToStore(t *testing.T) {
	testCases := []struct {
		name          string
		book          book.Book
		libraryCard   library_card.LibraryCard
		expectedError error
	}{
		{
			name:          "should create a new entry in the store",
			book:          book.CreateDummyBookOne(),
			libraryCard:   library_card.CreateDummyLibraryCardOne(),
			expectedError: nil,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			store, err := NewIssuedBookStore()
			assert.NoError(t, err)
			err = store.Add(testCase.book, testCase.libraryCard)
			assert.Equal(t, err, testCase.expectedError)
		})
	}
}

func TestNumberOfBooksIssuedOn(t *testing.T) {
	testCases := []struct {
		name           string
		libraryCard    library_card.LibraryCard
		expectedResult int
	}{
		{
			name:           "should return number of books on the libraryCard",
			libraryCard:    library_card.CreateDummyLibraryCardOne(),
			expectedResult: 1,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			store, err := NewIssuedBookStore()
			assert.NoError(t, err)
			err = store.Add(book.CreateDummyBookOne(), testCase.libraryCard)
			assert.NoError(t, err)
			res := store.NumberOfBooksIssuedOnLibraryCard(testCase.libraryCard)
			assert.Equal(t, res, testCase.expectedResult)
		})
	}
}

func TestRemoveFromStore(t *testing.T) {
	testCases := []struct {
		name          string
		book          book.Book
		libraryCard   library_card.LibraryCard
		expectedError error
	}{
		{
			name:          "should remove entry in the store",
			book:          book.CreateDummyBookOne(),
			libraryCard:   library_card.CreateDummyLibraryCardOne(),
			expectedError: nil,
		},
		{
			name:          "should remove entry in the store",
			book:          book.CreateDummyBookOne(),
			libraryCard:   library_card.CreateDummyLibraryCardOne(),
			expectedError: errors.New("the library card doesn't have this book issued"),
		},
		{
			name:          "should remove entry in the store",
			book:          book.CreateDummyBookTwo(),
			libraryCard:   library_card.CreateDummyLibraryCardOne(),
			expectedError: errors.New("the library card doesn't have this book issued"),
		},
	}
	store, err := NewIssuedBookStore()
	assert.NoError(t, err)
	err = store.Add(book.CreateDummyBookOne(), library_card.CreateDummyLibraryCardOne())
	assert.NoError(t, err)
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			err, _ := store.Remove(testCase.book, testCase.libraryCard)
			assert.Equal(t, err, testCase.expectedError)
		})
	}
}
