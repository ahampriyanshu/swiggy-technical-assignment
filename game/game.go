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
	var NumberOfPlayers int
	IsValidInput := false

	for !IsValidInput {
		console.Prompt("Enter the number of players [2-4]")
		_, err := fmt.Scanln(&NumberOfPlayers)
		if err != nil {
			console.Error("Invalid input. Please try again.")
			continue
		}

		// Validating the number of users to be in the defined limits
		if NumberOfPlayers < constants.PLAYER_MIN_LIMIT || NumberOfPlayers > constants.PLAYER_MAX_LIMIT {
			console.Error("Invalid Input. Number of players cannot be less than 2 and more than 4. Please try again...")
			continue
		}

		IsValidInput = true
	}

	// Initializing a randomly shuffled Deck
	Deck := NewDeck()

	// Initializing the list of Players
	Players := make([]Player, NumberOfPlayers)

	// Distributing 5 cards to each player
	for i := 1; i <= NumberOfPlayers; i++ {
		player := Player{id: i}

		for j := 1; j <= 5; j++ {
			card := Deck.DrawCard()
			player.AddCard(card)
		}
		Players[i-1] = player
	}

	// Initializing the discard pile with top card for the first turn
	DiscardPile := make([]Card, 0)
	DiscardPile = append(DiscardPile, Deck.DrawCard())

	// Creating the draw pile
	DrawPile := make([]Card, len(Deck.deck))
	copy(DrawPile, Deck.deck)

	// Starting the game
	CurrentPlayerTurn := 0
	CurrentDirection := 1
	RemaingCardsToDraw := 1

	// Creating an infinte loop
	for {

		// Breaking the loop when the drawpile is empty
		if len(DrawPile) < RemaingCardsToDraw {
			console.Warn("Game ended in a Draw as no more cards were left in the Draw Pile")
			break
		}

		// Continuing the player's turn
		CurrentPlayerTurn %= NumberOfPlayers
		if CurrentPlayerTurn < 0 {
			CurrentPlayerTurn += NumberOfPlayers
		}
		CurrentPlayerTurn %= NumberOfPlayers

		Matched := false
		MatchedNumber := -1

		// Fetching top card from the Discard Pile
		TopDiscardCard := DiscardPile[len(DiscardPile)-1]
		console.Warn("Top card of the Discard pile is %v", DiscardPile[len(DiscardPile)-1])

		// Iterating over the Deck of cards of the current player
		for _, CurrentPlayerCard := range Players[CurrentPlayerTurn].GetCards() {

			// Validating the current card to be discarded card
			if CurrentPlayerCard.number == TopDiscardCard.number || CurrentPlayerCard.suit == TopDiscardCard.suit {
				// Validating the discarded card to be one of the action cards
				if TopDiscardCard.number == constants.ACE || TopDiscardCard.number == constants.JACK || TopDiscardCard.number == constants.QUEEN || TopDiscardCard.number == constants.KING {
					if CurrentPlayerCard.number == TopDiscardCard.number {
						continue
					}
				}

				console.Info("Cards matched for player %d. Card and new Discard Deck top card = %v\n", Players[CurrentPlayerTurn].GetID(), CurrentPlayerCard)

				// Validating if the current player needs to draw more than one card
				if RemaingCardsToDraw > 1 {
					for RemaingCardsToDraw > 0 {
						console.Info("Player %d has drawn %v", Players[CurrentPlayerTurn].GetID(), DrawPile[len(DrawPile)-1])
						Players[CurrentPlayerTurn].AddCard(DrawPile[len(DrawPile)-1])
						DrawPile = DrawPile[:len(DrawPile)-1]
						RemaingCardsToDraw--
					}
					RemaingCardsToDraw = 1
				}

				// Shifting the matched to the discard pile
				Players[CurrentPlayerTurn].RemoveCard(CurrentPlayerCard)
				DiscardPile = append(DiscardPile, CurrentPlayerCard)
				Matched = true
				MatchedNumber = CurrentPlayerCard.number
				break
			}
		}

		// If none of cards matched then drawing from Draw pile
		if Matched == false {
			console.Warn("No match found for Player %d, drawing %d number of cards", Players[CurrentPlayerTurn].GetID(), RemaingCardsToDraw)
			for RemaingCardsToDraw > 0 {
				console.Info("Player %d has drawn %v", Players[CurrentPlayerTurn].GetID(), DrawPile[len(DrawPile)-1])
				Players[CurrentPlayerTurn].AddCard(DrawPile[len(DrawPile)-1])
				DrawPile = DrawPile[:len(DrawPile)-1]
				RemaingCardsToDraw--
			}
			RemaingCardsToDraw = 1
		}

		// Validating if Player was won the match (no more cards left after matching his/her card)
		if Matched == true && len(Players[CurrentPlayerTurn].GetCards()) == 0 {
			console.Success("Player %d has won the match. Game Ended.", Players[CurrentPlayerTurn].GetID())
			break
		}

		// Bonus conditions

		// Case Ace Card: Skipping the next player's turn
		if Matched == true && MatchedNumber == constants.ACE {
			CurrentPlayerTurn += CurrentDirection
		}

		// Case Jack Card: Updating the number of cards to be drawn from draw pile to 4
		if Matched == true && MatchedNumber == 11 {
			RemaingCardsToDraw = 4
		}

		// Case Queen Card: Updating the number of cards to be drawn from draw pile to 2
		if Matched == true && MatchedNumber == 12 {
			RemaingCardsToDraw = 2
		}

		// Case King Card: Reversing the direction in which the player's would play
		if Matched == true && MatchedNumber == 13 {
			CurrentDirection *= -1
		}

		directionString := "Counter-clockwise direction"
		if CurrentDirection == 1 {
			directionString = "Clockwise direction"
		}

		console.Prompt("########################################################")
		console.Prompt("Next player will draw %d number of cards", RemaingCardsToDraw)
		console.Prompt("%d cards are remaining in the draw pile", len(DrawPile))
		console.Prompt("The game will continue in " + directionString)
		console.Prompt("########################################################")
		fmt.Println()

		// Updating the next player
		CurrentPlayerTurn += CurrentDirection

	}

	return nil
}
