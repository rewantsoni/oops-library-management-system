package main

import (
	"fmt"
	"oops-library-management/book"
	"oops-library-management/invoice"
	"oops-library-management/librarian"
	"oops-library-management/library"
	"oops-library-management/library_card"
	"oops-library-management/library_management_system"
	"oops-library-management/payment"
)

func main() {

	myLibrary, _ := library.NewLibrary()
	myLibrarian, _ := librarian.NewLibrarian(myLibrary)
	store, _ := library_management_system.NewIssuedBookStore()
	paymentSystem, _ := payment.NewPaymentSystem()
	invoiceGenerator, _ := invoice.NewInvoiceGenerator()
	mylms, _ := library_management_system.NewLibraryManagementSystem(myLibrarian, store, paymentSystem, invoiceGenerator)

	bookOne, _ := book.NewBook("one")
	bookTwo, _ := book.NewBook("two")
	bookThree, _ := book.NewBook("three")
	bookFour, _ := book.NewBook("four")

	myLibrary.AddBook(bookOne)
	myLibrary.AddBook(bookTwo)
	myLibrary.AddBook(bookThree)

	myCardOne, _ := library_card.NewLibraryCard(1234)
	myCardTwo, _ := library_card.NewLibraryCard(5678)

	bookCardOne, err := mylms.BorrowBook("one", myCardOne)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(bookCardOne)
	}

	//bookCardOne, err = mylms.BorrowBook("three", myCardOne)
	//if err != nil {
	//	fmt.Println(err)
	//} else {
	//	fmt.Println(bookCardOne)
	//}

	myLibrary.AddBook(bookFour)

	bookCardTwo, err := mylms.BorrowBook("one", myCardTwo)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(bookCardTwo)
	}

	bookCardTwo, err = mylms.BorrowBook("four", myCardTwo)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(bookCardTwo)
	}

	invoice, err := mylms.Return(bookCardOne, myCardOne)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Invoice %v", invoice)
	}

}
