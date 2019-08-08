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
func shuffleDeck(deck Deck) Deck {
	// "pseudo"-randomize order of elements
	rand.Shuffle(len(deck), func(i, j int) {
		deck[i], deck[j] = deck[j], deck[i]
	})
	return deck
}

// Return and remove one value of array of Cards
func dealOne(deck Deck) Deck {

	if len(deck) == 0 {
		fmt.Println("No cards left to deal. Deck is empty!")
	}

	fmt.Println(deck[0])

	copy(deck[0:], deck[1:])
	deck[len(deck)-1] = empty
	deck = deck[:len(deck)-1]

	return deck
}

func main() {

}
