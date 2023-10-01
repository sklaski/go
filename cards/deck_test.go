package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()
	lengthDeck := 52
	if len(d) != lengthDeck {
		t.Errorf("Expected deck length of %v, but got %v.", lengthDeck, len(d))
	}

	firstCard := "Ace of Spades"
	if d[0] != firstCard {
		t.Errorf("Expected first card %v, but got %v", firstCard, d[0])
	}

	lastCard := "King of Clubs"
	if d[len(d)-1] != lastCard {
		t.Errorf("Expexted last card %v, but got %v", lastCard, d[len(d)-1])
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")

	deck := newDeck()
	_ = deck.saveToFile("_decktesting")

	loadedDeck := newDeckFromFile("_decktesting")

	lengthDeck := 52
	if len(loadedDeck) != lengthDeck {
		t.Errorf("Expected deck length of %v, but got %v.", lengthDeck, len(loadedDeck))
	}

	firstCard := "Ace of Spades"
	if loadedDeck[0] != firstCard {
		t.Errorf("Expected first card %v, but got %v", firstCard, loadedDeck[0])
	}

	lastCard := "King of Clubs"
	if loadedDeck[len(loadedDeck)-1] != lastCard {
		t.Errorf("Expexted last card %v, but got %v", lastCard, loadedDeck[len(loadedDeck)-1])
	}

	os.Remove("_decktesting")

}
