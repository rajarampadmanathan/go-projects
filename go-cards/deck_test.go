package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	deck := newDeck()
	if len(deck) != 52 {
		t.Errorf("Expected 52 but got %v", len(deck))
	}
	if deck[0] != "ace of heart" {
		t.Errorf("Expected ace of heart but got %v", deck[0])
	}

	if deck[len(deck)-1] != "queen of clubs" {
		t.Errorf("Expected ace of heart but got %v", deck[len(deck)-1])
	}
}

func TestSaveTODeckAndNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")
	deck := newDeck()
	deck.saveToFile("_decktesting")
	loadedDeck := readDeckFromFile("_decktesting")
	if len(loadedDeck) != 52 {
		t.Errorf("Expected 52 but got %v", len(deck))
	}
	os.Remove("_decktesting")
}
