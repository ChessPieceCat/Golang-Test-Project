package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
)

// Create a new type called 'deck' which is a slice of strings
type deck []string

// Create a function to create a new deck
func newDeck() deck {
	cards := deck{}
	cardSuits := []string{"Spades", "Clubs", "Diamonds", "Hearts"}
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards
}

// Create a new function called 'print' which is a receiver function for the 'deck' type
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

// Create a function to deal a hand of cards
func (d deck) deal(handSize int) (deck, deck) {
	// Create hand of first three cards
	hand := d[0:handSize]

	// Update deck
	cards := d[handSize:]

	return hand, cards
}

// Create a function to shuffle the deck
func (d deck) shuffle() {
	for i := len(d) - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		d[i], d[j] = d[j], d[i]
	}
}

// Create a function to convert deck to string
func (d deck) toString() string {
	deckString := strings.Join([]string(d), ",")

	return deckString
}

// Create a function to convert byte slice to string slice
func deckFromBytes(bs []byte) deck {
	// Convert to string
	s := string(bs)
	// Split to string slice
	deckSlice := strings.Split(s, ",")

	return deckSlice
}

// Create a function to load deck from file
func loadDeckFromFile(fileName string) (deck, error) {
	// Read from file
	bs, err := os.ReadFile(fileName)
	// Log error and exit if it exists
	if err != nil {
		return nil, err
	}
	loadedDeck := deckFromBytes(bs)

	return loadedDeck, nil
}

// Create a function to save deck to file
func (d deck) saveDeckToFile(fileName string) error {
	cards := d.toString()
	cardsByte := []byte(cards)

	// Return error if it exists
	return os.WriteFile(fileName, cardsByte, 0666)
}
