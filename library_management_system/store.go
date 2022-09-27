package library_management_system

import (
	"errors"
	"oops-library-management/book"
	"oops-library-management/library_card"
	"time"
)

type IssuedBookStore interface {
	Add(book.Book, library_card.LibraryCard) error
	NumberOfBooksIssuedOnLibraryCard(card library_card.LibraryCard) int
	Remove(book.Book, library_card.LibraryCard) (issueDetails, error)
}

type issueDetails struct {
	issuedDate  time.Time
	book        book.Book
	libraryCard library_card.LibraryCard
}

type inMemoryIssuedBookStore struct {
	issueDetails []issueDetails
}

func (i *inMemoryIssuedBookStore) Remove(b book.Book, card library_card.LibraryCard) (issueDetails, error) {
	var newIssueDetails []issueDetails
	var found bool
	var issuedBookDetails issueDetails
	for _, issueDetail := range i.issueDetails {
		if issueDetail.book == b && issueDetail.libraryCard == card {
			found = true
			issuedBookDetails = issueDetail
		} else {
			newIssueDetails = append(newIssueDetails, issueDetail)
		}
	}
	if !found {
		return issueDetails{}, errors.New("the library card doesn't have this book issued")
	}
	i.issueDetails = newIssueDetails
	return issuedBookDetails, nil
}

func (i *inMemoryIssuedBookStore) NumberOfBooksIssuedOnLibraryCard(card library_card.LibraryCard) int {
	count := 0
	for _, issueDetail := range i.issueDetails {
		if issueDetail.libraryCard == card {
			count = count + 1
		}
	}
	return count
}

func (i *inMemoryIssuedBookStore) Add(b book.Book, card library_card.LibraryCard) error {

	if card.GetCardLimit() <= i.NumberOfBooksIssuedOnLibraryCard(card) {
		return errors.New("cannot issue more books on this Card")
	}

	i.issueDetails = append(i.issueDetails, issueDetails{
		issuedDate:  time.Now(),
		book:        b,
		libraryCard: card,
	})
	return nil
}

func NewIssuedBookStore() (IssuedBookStore, error) {
	return &inMemoryIssuedBookStore{}, nil
}
