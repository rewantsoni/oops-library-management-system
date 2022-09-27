package library_card

func CreateDummyLibraryCardOne() LibraryCard {
	card, _ := NewLibraryCard(1234)
	return card
}

func CreateDummyLibraryCardTwo() LibraryCard {
	card, _ := NewLibraryCard(5678)
	return card
}
