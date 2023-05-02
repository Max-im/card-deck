package main

func main() {
	deck := newDeck()
	deck.print()
	deck.shuffle()
	deck.print()

	one, two := deal(deck, 5)
	one.print()
	two.print()

	one.saveToFile("one")

	newDeck := deckFromFile("one")
	newDeck.print()

}
