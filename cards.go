package main

import (
	"fmt"
	"math/rand"
)

// Card - Collection of attributes defining a card
type Card struct {
	Rank string
	Suit string
}

// Deck - Array to hold Cards
type Deck []Card

// Build deck using defined Card into Deck
func buildDeck() (deck Deck) {
	ranks := []string{"Ace", "1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "Jack", "Queen", "King"}
	suits := []string{"Heart", "Diamond", "Club", "Spade"}

}

// Randomize array of Cards
func shuffleDeck(deck Deck) Deck {
	rand.Shuffle(len(deck), func(i, j int) {
		deck[i], deck[j] = deck[j], deck[i]
	})
	return deck
}

// Return and remove one value of array of Cards
func dealOne(deck Deck) Deck {

	if len(deck) == empty {
		fmt.Println("No cards left to deal. Deck is empty!")
	}

	fmt.Println(deck[0])

	copy(deck[0:], deck[1:])
	deck[len(deck)-1] = nil
	deck = deck[:len(deck)-1]

	return deck
}

func main() {

}
