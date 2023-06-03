package main

import (
	"cardgame/console"
	"cardgame/game"
)

func main() {
	game := game.GamePlay{}
	err := game.StartGame()
	if err != nil {
		console.Error("Couldn't start the game: %s", err)
	}
}
