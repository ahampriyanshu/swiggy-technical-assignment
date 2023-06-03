package game

import (
	"math/rand"
	"time"
)

type Suit int

const (
	Spades Suit = iota
	Hearts
	Clubs
	Diamonds
)

type Card struct {
	number int
	suit   Suit
}

type Deck struct {
	deck []Card
}

func NewDeck() *Deck {
	deck := []Card{}
	for i := 1; i <= 13; i++ {
		for j := 0; j < 4; j++ {
			deck = append(deck, Card{
				number: i,
				suit:   Suit(j),
			})
		}
	}

	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(deck), func(i, j int) {
		deck[i], deck[j] = deck[j], deck[i]
	})

	return &Deck{deck: deck}
}

func (d *Deck) DrawCard() Card {
	card := d.deck[len(d.deck)-1]
	d.deck = d.deck[:len(d.deck)-1]
	return card
}
