package main

import (
	"cardgame/game"
	"fmt"
)

func main() {
	gp := game.GamePlay{}
	err := gp.PlayGame()
	if err != nil {
		fmt.Println("Error:", err)
	}
}
