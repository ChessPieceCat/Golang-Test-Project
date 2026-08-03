package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {

	// Check that deck has 52 cards
	d := newDeck()
	if len(d) != 52 {
		t.Errorf("Expected deck length of 52, but got %v", len(d))
	}

	// Check that first card is Ace of Spades
	if d[0] != "Ace of Spades" {
		t.Errorf("Expected 'Ace of Spades' at index 0 but got %v", d[0])
	}

	// Check that last card is Seven of Clubs
	if d[len(d)-1] != "King of Hearts" {
		t.Errorf("Expected 'King of Hearts' at index 51 but got %v", d[len(d)-1])
	}
}

func TestSaveDeckToFileAndLoadDeckFromFile(t *testing.T) {

	os.Remove("_testdeck")

	// Check that save creates a file with specified name
	d := newDeck()
	err := d.saveDeckToFile("_testdeck")
	if err != nil {
		t.Errorf("Failed to save deck due to %v", err)
	}

	// Check that load loads the specified file
	d, err = loadDeckFromFile("_testdeck")
	if err != nil {
		t.Errorf("Loading failed with error: %v", err)
	}

	// Check that loaded deck has correct length
	if len(d) != 52 {
		t.Errorf("Expected deck length of 52, but got %v", len(d))
	}

	// Check that load logs errors when applicable
	_, err = loadDeckFromFile("does_not_exist")
	if err == nil {
		t.Errorf("Expected load error but got none")
	}

	os.Remove("_testdeck")
}
