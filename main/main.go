package main

func main() {
	cards := newDeck()
	cards.saveToFile("my_cards")

	imported_cards := newDeckFromFile("my_cards")
	imported_cards.print()
	imported_cards.shuffle()
	imported_cards.print()
}
