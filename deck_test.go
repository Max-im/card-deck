package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	var expectedLen = 16

	if len(d) != expectedLen {
		t.Errorf("Expected deck length %v but got %v", expectedLen, len(d))
	}

	var firstCard = "Ace of Spades"

	if d[0] != firstCard {
		t.Errorf("Expected first card %v but got %v", firstCard, d[0])
	}

	var lastCard = "Four of Clubs"

	if d[len(d)-1] != lastCard {
		t.Errorf("Expected first card %v but got %v", lastCard, d[len(d)-1])
	}
}

func TestSaveAndReadToDeck(t *testing.T) {
	var fileName = "_decktesting"
	os.Remove(fileName)

	deck := newDeck()
	deck.saveToFile(fileName)

	loadedDeck := deckFromFile(fileName)

	var expectedLen = 16

	if len(loadedDeck) != expectedLen {
		t.Errorf("Expected deck length %v but got %v", expectedLen, len(loadedDeck))
	}

	os.Remove(fileName)
}
