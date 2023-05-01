package main

import "testing"

func TestNewDeck(t *testing.T) {
	d := newDeck()
	var expectedLen = 16

	if len(d) != expectedLen {
		t.Errorf("Expected deck length %v but got %v", expectedLen, len(d))
	}
}
