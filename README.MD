Deck of Cards Builder, Shuffler, and Dealer in GO

A `method` (`class` in Java) that represents a deck of poker-style playing cards providing three operations:

<br>
buildDeck() 

- Builds a deck of 52 playing-cards with rank and suit values, returning that deck.

<br>
shuffle() 

- Shuffle returns no value, but results in the built deck being randomly permuted. 
- While not using a library-provided “shuffle” operation to implement this, in this case a source namespace overload function 'Shuffle' is used.

<br>
dealOneCard()

- Function returns one card from the deck to the caller. 
- A call to shuffle followed by 52 calls to dealOneCard() will result in the caller being provided all 52 cards of the deck in a random order. 
- If the caller then makes a 53rd call dealOneCard(), no card is dealt, but rather a prompt returned.

<br>
