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

	var firstCardName = "Ace of Spades"
	var firstCardVal = 1

	if d[0].name != firstCardName {
		t.Errorf("Expected first card name %v but got %v", firstCardName, d[0].name)
	}

	if d[0].value != firstCardVal {
		t.Errorf("Expected first card value %v but got %v", firstCardVal, d[0].value)
	}

	var lastCardName = "Four of Clubs"
	var lastCardValue = 4

	if d[len(d)-1].name != lastCardName {
		t.Errorf("Expected first card name %v but got %v", lastCardName, d[len(d)-1].name)
	}

	if d[len(d)-1].value != lastCardValue {
		t.Errorf("Expected first card value %v but got %v", lastCardValue, d[len(d)-1].value)
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
