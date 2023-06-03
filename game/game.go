package game

import "fmt"

type GamePlay struct {
	deck        []Card
	players     []Player
	drawPile    []Card
	discardPile []Card
}

func (gp *GamePlay) PlayGame() error {
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
			card := deck.DrawCard()
			player.AddCard(card)
		}
		players[i-1] = player
	}

	discardPile := make([]Card, 0)
	discardPile = append(discardPile, deck.DrawCard())

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

		for _, currentPlayerCard := range players[playerTurn].GiveCards() {
			if currentPlayerCard.number == topDiscardCard.number || currentPlayerCard.suit == topDiscardCard.suit {
				if topDiscardCard.number == 1 || topDiscardCard.number == 11 || topDiscardCard.number == 12 || topDiscardCard.number == 13 {
					if currentPlayerCard.number == topDiscardCard.number {
						continue
					}
				}

				fmt.Printf("Cards matched for player %d. Card and new Discard Deck top card = %v\n", players[playerTurn].GiveID(), currentPlayerCard)

				if numCardsTake > 1 {
					for numCardsTake > 0 {
						fmt.Println("Drawing", drawPile[len(drawPile)-1], "Card")
						players[playerTurn].AddCard(drawPile[len(drawPile)-1])
						drawPile = drawPile[:len(drawPile)-1]
						numCardsTake--
					}
					numCardsTake = 1
				}

				players[playerTurn].RemoveCard(currentPlayerCard)
				discardPile = append(discardPile, currentPlayerCard)
				matched = true
				matchedNumber = currentPlayerCard.number
				break
			}
		}

		if matched == false {
			fmt.Printf("No cards match for player %d. Taking %d Card(s)\n", players[playerTurn].GiveID(), numCardsTake)
			for numCardsTake > 0 {
				fmt.Println("Drawing", drawPile[len(drawPile)-1], "Card")
				players[playerTurn].AddCard(drawPile[len(drawPile)-1])
				drawPile = drawPile[:len(drawPile)-1]
				numCardsTake--
			}
			numCardsTake = 1
		}

		if matched == true && len(players[playerTurn].GiveCards()) == 0 {
			fmt.Printf("YAY! Player %d won the match!!!\n", players[playerTurn].GiveID())
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
