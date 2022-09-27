package library_card

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateLibraryCard(t *testing.T) {

	testCases := []struct {
		name           string
		param          int
		expectedResult LibraryCard
		expectedError  error
	}{
		{
			name:           "should create a new LibraryCard successfully",
			param:          1234,
			expectedResult: libraryCardWithNumber{number: 1234, limit: 1},
		},
		{
			name:          "should not create a new library card if the parameter is empty",
			expectedError: errors.New("library card number has to of at least 4 digits"),
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			res, err := NewLibraryCard(testCase.param)
			assert.Equal(t, err, testCase.expectedError)
			if err == nil {
				assert.Equal(t, res, testCase.expectedResult)
			}
		})
	}
}

func TestGetCardLimit(t *testing.T) {

	testCases := []struct {
		name           string
		expectedResult int
	}{
		{
			name:           "should return the right limit",
			expectedResult: 1,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			card, err := NewLibraryCard(1234)
			assert.NoError(t, err)
			res := card.GetCardLimit()
			assert.Equal(t, res, testCase.expectedResult)
		})
	}
}
