package game

import (
	"cardgame/console"
	"cardgame/constants"
	"fmt"
)

type CardGame struct {
	deck        []Card
	players     []Player
	drawPile    []Card
	discardPile []Card
}

func (gp *CardGame) StartGame() error {
	var numOfPlayers int
	validInput := false

	for !validInput {
		console.Prompt("Enter the number of players [2-4]")
		_, err := fmt.Scanln(&numOfPlayers)
		if err != nil {
			console.Error("Invalid input. Please try again.")
			continue
		}

		// Validating the number of users to be in the defined limits
		if numOfPlayers < constants.PLAYER_MIN_LIMIT || numOfPlayers > constants.PLAYER_MAX_LIMIT {
			console.Error("Invalid Input. Number of players cannot be less than 2 and more than 4. Please try again...")
			continue
		}

		validInput = true
	}

	// Initializing a randomly shuffled deck
	deck := NewDeck()

	// Initializing the list of players
	players := make([]Player, numOfPlayers)

	// Distributing 5 cards to each player
	for i := 1; i <= numOfPlayers; i++ {
		player := Player{id: i}

		for j := 1; j <= 5; j++ {
			card := deck.DrawCard()
			player.AddCard(card)
		}
		players[i-1] = player
	}

	// Initializing the discard pile with top card for the first turn
	discardPile := make([]Card, 0)
	discardPile = append(discardPile, deck.DrawCard())

	// Creating the draw pile
	drawPile := make([]Card, len(deck.deck))
	copy(drawPile, deck.deck)

	// Starting the game
	playerTurn := 0
	direction := 1
	numCardsTake := 1

	// Creating an infinte loop
	for {

		// Breaking the loop when the drawpile is empty
		if len(drawPile) < numCardsTake {
			console.Warn("Game ended in a Draw as no more cards were left in the Draw Pile")
			break
		}

		// Continuing the player's turn
		playerTurn %= numOfPlayers
		if playerTurn < 0 {
			playerTurn += numOfPlayers
		}
		playerTurn %= numOfPlayers

		matched := false
		matchedNumber := -1

		// Fetching top card from the Discard Pile
		topDiscardCard := discardPile[len(discardPile)-1]
		console.Warn("Top card of the Discard pile is %v", discardPile[len(discardPile)-1])

		// Iterating over the Deck of cards of the current player
		for _, currentPlayerCard := range players[playerTurn].GetCards() {

			// Validating the current card to be discarded card
			if currentPlayerCard.number == topDiscardCard.number || currentPlayerCard.suit == topDiscardCard.suit {
				// Validating the discarded card to be one of the action cards
				if topDiscardCard.number == constants.ACE || topDiscardCard.number == constants.JACK || topDiscardCard.number == constants.QUEEN || topDiscardCard.number == constants.KING {
					if currentPlayerCard.number == topDiscardCard.number {
						continue
					}
				}

				console.Info("Cards matched for player %d. Card and new Discard Deck top card = %v\n", players[playerTurn].GetID(), currentPlayerCard)

				// Validating if the current player needs to draw more than one card
				if numCardsTake > 1 {
					for numCardsTake > 0 {
						console.Info("Player %d has drawn %v", players[playerTurn].GetID(), drawPile[len(drawPile)-1])
						players[playerTurn].AddCard(drawPile[len(drawPile)-1])
						drawPile = drawPile[:len(drawPile)-1]
						numCardsTake--
					}
					numCardsTake = 1
				}

				// Shifting the matched to the discard pile
				players[playerTurn].RemoveCard(currentPlayerCard)
				discardPile = append(discardPile, currentPlayerCard)
				matched = true
				matchedNumber = currentPlayerCard.number
				break
			}
		}

		// If none of cards matched then drawing from Draw pile
		if matched == false {
			console.Warn("No match found for Player %d, drawing %d number of cards", players[playerTurn].GetID(), numCardsTake)
			for numCardsTake > 0 {
				console.Info("Player %d has drawn %v", players[playerTurn].GetID(), drawPile[len(drawPile)-1])
				players[playerTurn].AddCard(drawPile[len(drawPile)-1])
				drawPile = drawPile[:len(drawPile)-1]
				numCardsTake--
			}
			numCardsTake = 1
		}

		// Validating if Player was won the match (no more cards left after matching his/her card)
		if matched == true && len(players[playerTurn].GetCards()) == 0 {
			console.Success("Player %d has won the match. Game Ended.", players[playerTurn].GetID())
			break
		}

		// Bonus conditions

		// Case Ace Card: Skipping the next player's turn
		if matched == true && matchedNumber == constants.ACE {
			playerTurn += direction
		}

		// Case Jack Card: Updating the number of cards to be drawn from draw pile to 4
		if matched == true && matchedNumber == 11 {
			numCardsTake = 4
		}

		// Case Queen Card: Updating the number of cards to be drawn from draw pile to 2
		if matched == true && matchedNumber == 12 {
			numCardsTake = 2
		}

		// Case King Card: Reversing the direction in which the player's would play
		if matched == true && matchedNumber == 13 {
			direction *= -1
		}

		directionString := "Counter-clockwise direction"
		if direction == 1 {
			directionString = "Clockwise direction"
		}

		console.Prompt("########################################################")
		console.Prompt("Next player will draw %d number of cards", numCardsTake)
		console.Prompt("%d cards are remaining in the draw pile", len(drawPile))
		console.Prompt("The game will continue in " + directionString)
		console.Prompt("########################################################")
		fmt.Println()

		// Updating the next player
		playerTurn += direction

	}

	return nil
}
