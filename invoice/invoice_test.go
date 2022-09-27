package invoice

import (
	"github.com/stretchr/testify/assert"
	"oops-library-management/amount"
	"oops-library-management/book"
	"testing"
	"time"
)

func TestCreateNewInvoice(t *testing.T) {
	timeOne := time.Time{}
	testCases := []struct {
		name           string
		book           book.Book
		issuedDate     time.Time
		amount         amount.Amount
		expectedResult Invoice
		expectedError  error
	}{
		{
			name:           "should create a new invoice",
			book:           book.CreateDummyBookOne(),
			issuedDate:     timeOne,
			amount:         amount.CreateDummyAmount(),
			expectedResult: Invoice{book: book.CreateDummyBookOne(), issuedDate: timeOne, amount: amount.CreateDummyAmount()},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			res, err := NewInvoice(testCase.book, testCase.issuedDate, testCase.amount)
			assert.Equal(t, err, testCase.expectedError)
			if err == nil {
				assert.Equal(t, res, testCase.expectedResult)
			}
		})
	}
}
