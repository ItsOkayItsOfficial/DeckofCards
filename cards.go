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

	// String arrays defining Card values
	ranks := []string{"Ace", "1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "Jack", "Queen", "King"}
	suits := []string{"Heart", "Diamond", "Club", "Spade"}

	// Populate our Deck array of Card values
	for i := 0; i < len(ranks); i++ {
		for n := 0; n < len(suits); n++ {
			card := Card{
				Rank: ranks[i],
				Suit: suits[n],
			}
			deck = append(deck, card)
		}
	}

	return
}

// Randomize array of Cards
func shuffle(deck Deck) Deck {

	// "pseudo"-randomize order of elements
	rand.Shuffle(len(deck), func(i, j int) {
		deck[i], deck[j] = deck[j], deck[i]
	})

	return deck
}

// Return and remove last value of array of Cards
func dealOneCard(deck Deck) Deck {

	// Print Card and then remove
	if len(deck) > 0 {
		fmt.Println(deck[len(deck)])
		deck = deck[:len(deck)-1]
	} else {
		// No Card left in Deck array
		fmt.Println("No cards left to deal. Deck is empty!")
	}

	return deck
}

// Ideally we'll make this on recall
func main() {

	// Get a deck of cards
	deck := buildDeck()

	// Shuffle deck of cards
	shuffle(deck)

	// Deal one card
	dealOneCard(deck)
}
