package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 52 {
		t.Errorf("Expected deck length of %d", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected last card of Ace of Spades but got %v", d[0])
	}

	if d[len(d)-1] != "King of Clubs" {
		t.Errorf("Expected last card of King of Clubs but got %v", d[len(d)-1])
	}
}

func TestSaveAndReadDeckFromFile(t *testing.T) {
	os.Remove("_decktesting.txt")
	d := newDeck()
	d.saveToFile("_decktesting.txt")

	loadedDeck := newDeckFromFile("_decktesting.txt")

	if len(loadedDeck) != 52 {
		t.Errorf("Expected deck length of %d", len(d))
	}

	if loadedDeck[0] != "Ace of Spades" {
		t.Errorf("Expected last card of Ace of Spades but got %v", loadedDeck[0])
	}

	if loadedDeck[len(loadedDeck)-1] != "King of Clubs" {
		t.Errorf("Expected last card of King of Clubs but got %v", loadedDeck[len(loadedDeck)-1])
	}
	os.Remove("_decktesting.txt")
}
