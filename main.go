package main

import (
	"go-snake/snake"
)

func main() {
	// create a new game
	game := snake.Game{}
	// initialize the game
	game.InitGame()
	// start the game
	game.Run()
}