package invoice

import (
	"github.com/stretchr/testify/assert"
	"oops-library-management/amount"
	"oops-library-management/book"
	"testing"
	"time"
)

func TestCreateNewInvoiceGenerator(t *testing.T) {
	testCases := []struct {
		name           string
		expectedResult InvoiceGenerator
		expectedError  error
	}{
		{
			name:           "should create a new invoice generator",
			expectedResult: SimpleInvoiceGenerator{},
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			res, err := NewInvoiceGenerator()
			assert.Equal(t, err, testCase.expectedError)
			if err == nil {
				assert.Equal(t, res, testCase.expectedResult)
			}
		})
	}
}

func TestGenerate(t *testing.T) {
	timeOne := time.Now()
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
			myInvoiceGenerator, err := NewInvoiceGenerator()
			assert.NoError(t, err)
			res, err := myInvoiceGenerator.Generate(testCase.book, testCase.issuedDate, testCase.amount)
			assert.Equal(t, err, testCase.expectedError)
			if err == nil {
				assert.Equal(t, res, testCase.expectedResult)
			}
		})
	}
}
