package game

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

type Suit int

// Enum to maintain index of Suit
const (
	Spades Suit = iota
	Hearts
	Clubs
	Diamonds
)

// Map of sting to mantain name of Suit
var suitNames = map[Suit]string{
	Spades:   "Spades",
	Hearts:   "Hearts",
	Clubs:    "Clubs",
	Diamonds: "Diamonds",
}

type Card struct {
	number int
	suit   Suit
}

// Custom method for formatting the Suit and Number of Card before logging out
func (c Card) String() string {
	var numberStr string

	switch c.number {
	case 1:
		numberStr = "Ace"
	case 11:
		numberStr = "Jack"
	case 12:
		numberStr = "Queen"
	case 13:
		numberStr = "King"
	default:
		numberStr = strconv.Itoa(c.number)
	}

	return fmt.Sprintf("%s of %s", numberStr, suitNames[c.suit])
}

type Deck struct {
	deck []Card
}

func NewDeck() *Deck {

	// new list of struct Card
	deck := []Card{}

	// Iterating for number of cards
	for i := 1; i <= 13; i++ {

		// Iterating for number of suits
		for j := 0; j < 4; j++ {

			// appending the card to the deck
			deck = append(deck, Card{
				number: i,
				suit:   Suit(j),
			})
		}
	}

	// Shuffling the deck at random using timestamp as seed
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(deck), func(i, j int) {
		deck[i], deck[j] = deck[j], deck[i]
	})

	return &Deck{deck: deck}
}

func (d *Deck) DrawCard() Card {

	// Drawing a card out of the deck
	card := d.deck[len(d.deck)-1]
	d.deck = d.deck[:len(d.deck)-1]
	return card
}
