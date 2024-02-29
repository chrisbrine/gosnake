package snake

func (g *Game) DisplayPaused() {
	// pause the game
	g.DisplayCenterBlock([]string{
		"===========================",
		"|                         |",
		"|          PAUSED         |",
		"|                         |",
		"|   Press any key to go   |",
		"|    Press 'q' to quit    |",
		"|                         |",
		"===========================",
	}, g.settings.backgroundColor, g.settings.foregroundColor)
}