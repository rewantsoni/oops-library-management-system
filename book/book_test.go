package book

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateNewBook(t *testing.T) {

	testCases := []struct {
		name           string
		param          string
		expectedResult Book
		expectedError  error
	}{
		{
			name:           "should create a new Book successfully",
			param:          "Aladdin",
			expectedResult: paperBackBook{name: "Aladdin"},
		},
		{
			name:          "should create not create a book if name is empty",
			param:         "",
			expectedError: errors.New("book name cannot be empty"),
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			res, err := NewBook(testCase.param)
			assert.Equal(t, err, testCase.expectedError)
			if err == nil {
				assert.Equal(t, res, testCase.expectedResult)
			}
		})
	}
}

func TestGetBookName(t *testing.T) {

	testCases := []struct {
		name             string
		param            string
		expectedBookName string
	}{
		{
			name:             "should return the right book name",
			param:            "Aladdin",
			expectedBookName: "Aladdin",
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			book, err := NewBook(testCase.param)
			assert.NoError(t, err)
			res := book.GetName()
			assert.Equal(t, res, testCase.expectedBookName)
		})
	}
}
