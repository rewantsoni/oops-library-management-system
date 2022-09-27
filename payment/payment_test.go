package payment

import (
	"github.com/stretchr/testify/assert"
	"oops-library-management/amount"
	"testing"
	"time"
)

func TestCreateNewPaymentSystem(t *testing.T) {
	testCases := []struct {
		name           string
		expectedResult PaymentSystem
		expectedError  error
	}{
		{
			name:           "should create a new payment system",
			expectedResult: simplePaymentSystem{},
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			res, err := NewPaymentSystem()
			assert.Equal(t, err, testCase.expectedError)
			if err == nil {
				assert.Equal(t, res, testCase.expectedResult)
			}
		})
	}
}

func TestCompute(t *testing.T) {
	// TODO: Create a timeGenerator and to get time
	timeOne := time.Date(2022, 9, 24, 0, 0, 0, 0, &time.Location{})
	testCases := []struct {
		name           string
		issuedDate     time.Time
		expectedResult func() (amount.Amount, error)
	}{
		{
			name:       "should return the amount to pay",
			issuedDate: timeOne,
			expectedResult: func() (amount.Amount, error) {
				return amount.NewAmount(7)
			},
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			myPaymentSystem, err := NewPaymentSystem()
			assert.NoError(t, err)
			res, err := myPaymentSystem.Compute(testCase.issuedDate)
			expectedRes, expectedErr := testCase.expectedResult()
			assert.Equal(t, err, expectedErr)
			if err == nil {
				assert.Equal(t, res, expectedRes)
			}
		})
	}
}
