package snake

func (g *Game) DisplayInstructions() {
	// display the game instructions
	g.DisplayCenterBlock([]string{
		"===========================",
		"|                         |",
		"|      SNAKE GAME         |",
		"|                         |",
		"|   Use the arrow keys    |",
		"|   to move the snake     |",
		"|   Press 'p' to pause    |",
		"|   Press 'q' to quit     |",
		"|                         |",
		"===========================",
	}, g.settings.backgroundColor, g.settings.foregroundColor)
}