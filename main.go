package main

func main() {
	cards := newDeck()
	cards.saveToFile("my_cards")

	loadedCards := newDeckFromFile("my_cards")
	loadedCards.print()
}
