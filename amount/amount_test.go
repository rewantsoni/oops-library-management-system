package amount

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateNewAmount(t *testing.T) {

	testCases := []struct {
		name           string
		param          float64
		expectedResult Amount
		expectedError  error
	}{
		{
			name:           "should create a new amount",
			param:          10,
			expectedResult: Amount{value: 10},
		},
		{
			name:          "should not create a new amount",
			param:         0,
			expectedError: errors.New("amount cannot be zero or negative"),
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			res, err := NewAmount(testCase.param)
			assert.Equal(t, err, testCase.expectedError)
			if err == nil {
				assert.Equal(t, res, testCase.expectedResult)
			}
		})
	}
}
