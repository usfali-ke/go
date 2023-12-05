package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 20 {
		// t.Errorf("Expected deck length of 20, but go", len(d))
		t.Errorf("Expected deck length of 16, but got %v", len(d))
	}
	if d[0] != "One of Spades" {
		t.Errorf("Expected first card, One of Spades, but got %v", d[0])
	}
	if d[len(d)-1] != "Four of Hearts" {
		t.Errorf("Expected last card, Four of Hearts, but got %v", d[len(d)-1])
	}

}

func TestSaveToFileAndNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")
	deck := newDeck()
	deck.saveToFile("_decktesting") // save to file with name _decktesting

	loadedDeck := newDeckFromFile("_decktesting")

	if len(loadedDeck) != 16 {
		t.Errorf("Expected deck length of 16, but got %v", len(loadedDeck))
	}
	os.Remove("_decktesting")
}
