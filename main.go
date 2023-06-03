package main

import (
	"fmt"
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

type Player struct {
	id    int
	cards []Card
}

type GamePlay struct {
	deck        []Card
	players     []Player
	drawPile    []Card
	discardPile []Card
}

func (p *Player) addCard(card Card) {
	p.cards = append(p.cards, card)
}

func (p *Player) removeCard(card Card) {
	for i, c := range p.cards {
		if c == card {
			p.cards = append(p.cards[:i], p.cards[i+1:]...)
			break
		}
	}
}

func (p *Player) giveId() int {
	return p.id
}

func (p *Player) giveCards() []Card {
	return p.cards
}

func (gp *GamePlay) playGame() error {
	var numOfPlayers int
	validInput := false

	for !validInput {
		fmt.Println("Enter number of players (2-4)")
		_, err := fmt.Scanln(&numOfPlayers)
		if err != nil {
			fmt.Println("Invalid input. Please try again.")
			continue
		}

		if numOfPlayers < 2 || numOfPlayers > 4 {
			fmt.Println("Invalid number of players. Please try again.")
			continue
		}

		validInput = true
	}

	deck := NewDeck()

	players := make([]Player, numOfPlayers)
	for i := 1; i <= numOfPlayers; i++ {
		player := Player{id: i}

		for j := 1; j <= 5; j++ {
			card := deck.drawCard()
			player.addCard(card)
		}
		players[i-1] = player
	}

	discardPile := make([]Card, 0)
	discardPile = append(discardPile, deck.drawCard())

	drawPile := make([]Card, len(deck.deck))
	copy(drawPile, deck.deck)

	playerTurn := 0
	direction := 1
	numCardsTake := 1

	for {
		if len(drawPile) < numCardsTake {
			fmt.Println("Game drawn!! Cards are less.")
			break
		}

		playerTurn %= numOfPlayers
		if playerTurn < 0 {
			playerTurn += numOfPlayers
		}
		playerTurn %= numOfPlayers

		matched := false
		matchedNumber := -1
		topDiscardCard := discardPile[len(discardPile)-1]

		fmt.Println("Discard deck top card =", discardPile[len(discardPile)-1])

		for _, currentPlayerCard := range players[playerTurn].giveCards() {
			if currentPlayerCard.number == topDiscardCard.number || currentPlayerCard.suit == topDiscardCard.suit {
				if topDiscardCard.number == 1 || topDiscardCard.number == 11 || topDiscardCard.number == 12 || topDiscardCard.number == 13 {
					if currentPlayerCard.number == topDiscardCard.number {
						continue
					}
				}

				fmt.Printf("Cards matched for player %d. Card and new Discard Deck top card = %v\n", players[playerTurn].giveId(), currentPlayerCard)

				if numCardsTake > 1 {
					for numCardsTake > 0 {
						fmt.Println("Drawing", drawPile[len(drawPile)-1], "Card")
						players[playerTurn].addCard(drawPile[len(drawPile)-1])
						drawPile = drawPile[:len(drawPile)-1]
						numCardsTake--
					}
					numCardsTake = 1
				}

				players[playerTurn].removeCard(currentPlayerCard)
				discardPile = append(discardPile, currentPlayerCard)
				matched = true
				matchedNumber = currentPlayerCard.number
				break
			}
		}

		if matched == false {
			fmt.Printf("No cards match for player %d. Taking %d Card(s)\n", players[playerTurn].giveId(), numCardsTake)
			for numCardsTake > 0 {
				fmt.Println("Drawing", drawPile[len(drawPile)-1], "Card")
				players[playerTurn].addCard(drawPile[len(drawPile)-1])
				drawPile = drawPile[:len(drawPile)-1]
				numCardsTake--
			}
			numCardsTake = 1
		}

		if matched == true && len(players[playerTurn].giveCards()) == 0 {
			fmt.Printf("YAY! Player %d won the match!!!\n", players[playerTurn].giveId())
			break
		}

		if matched == true && matchedNumber == 1 {
			playerTurn += direction
		}

		if matched == true && matchedNumber == 13 {
			direction *= -1
		}

		if matched == true && matchedNumber == 11 {
			numCardsTake = 4
		}

		if matched == true && matchedNumber == 12 {
			numCardsTake = 2
		}

		playerTurn += direction

		fmt.Println("--- Next Turn ---")
	}

	return nil
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

func (d *Deck) drawCard() Card {
	card := d.deck[len(d.deck)-1]
	d.deck = d.deck[:len(d.deck)-1]
	return card
}

func main() {
	gp := GamePlay{}
	err := gp.playGame()
	if err != nil {
		fmt.Println("Error:", err)
	}
}
