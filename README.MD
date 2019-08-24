# Deck of Cards Builder, Shuffler, and Dealer in GO
[![Go Report Card](https://goreportcard.com/badge/github.com/ItsOkayItsOfficial/DeckofCards?style=flat-square)](https://goreportcard.com/report/github.com/ItsOkayItsOfficial/DeckofCards)
[![Go Doc](https://img.shields.io/badge/godoc-reference-blue.svg?style=flat-square)](https://godoc.org/-/go)
[![Release](https://img.shields.io/github/release/golang-standards/project-layout.svg?style=flat-square)](https://github.com/ItsOkayItsOfficial/DeckofCards/releases/latest)

A `method` (`class` in Java) that represents a deck of poker-style playing cards providing three operations.

## Operations
<br>

`buildDeck()` 

- Builds a deck of 52 playing-cards with rank and suit values, returning that deck.

<br>

`shuffle()`

- Shuffle returns no value, but results in the built deck being randomly permuted. 
- While not using a library-provided “shuffle” operation to implement this, in this case a source namespace overload function 'Shuffle' is used.

<br>

`dealOneCard()`

- Function returns one card from the deck to the caller. 
- A call to shuffle followed by 52 calls to dealOneCard() will result in the caller being provided all 52 cards of the deck in a random order. 
- If the caller then makes a 53rd call dealOneCard(), no card is dealt, but rather a prompt returned.

<br>

## Notes
