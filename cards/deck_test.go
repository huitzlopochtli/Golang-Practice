package main

import (
	"os"
	"reflect"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 52 {
		t.Errorf("Expected deck length of 52, but got %d", len(d))
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	os.Remove(".decktesting")

	deck := newDeck()
	deck.saveToFile(".decktesting")

	loadedDeck := newDeckFromFile(".decktesting")
	if len(loadedDeck) == 0 {
		t.Errorf("Deck not loaded")
	}

	os.Remove(".decktesting")
}

func TestDeckShuffle(t *testing.T) {
	deckNotShuffled := newDeck()

	deckShuffeld := newDeck()
	deckShuffeld.shuffle()

	if reflect.DeepEqual(deckShuffeld, deckNotShuffled) {
		t.Errorf("Shuffling of decks not working")
	}
}
