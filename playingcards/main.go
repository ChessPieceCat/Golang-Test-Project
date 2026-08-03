package main

import (
	"fmt"
	"os"
)

func main() {
	// Initialize cards variable
	var cards deck

	// Check for existence of deck file
	_, err := os.Stat("saved_cards")

	// Create and save deck if it does not exist, else read from file
	if os.IsNotExist(err) {
		fmt.Println(Magenta + "Creating New Deck..." + Reset)

		// Create new deck
		cards = newDeck()

		// Save deck to file
		cards.saveDeckToFile("saved_cards")

	} else {
		fmt.Println(Magenta + "Loading Deck..." + Reset)
		cards, err = loadDeckFromFile("saved_cards")
		if err != nil {
			fmt.Println("Error: ", err)
		}
	}

	// Call print function with variable cards
	fmt.Println(Magenta + "Full Deck: " + Reset)
	cards.print()

	// Shuffle cards and print
	cards.shuffle()
	fmt.Println(Magenta + "Shuffled Deck: " + Reset)
	cards.print()

	// Create hand from cards
	hand, remainingCards := cards.deal(5)

	// Print each card in hand
	fmt.Println(Magenta + "Current Hand: " + Reset)
	hand.print()

	// Print updated deck
	fmt.Println(Magenta + "Updated Deck: " + Reset)
	remainingCards.print()

}
