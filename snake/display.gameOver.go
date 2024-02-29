package snake

import (
	"fmt"

	"github.com/gdamore/tcell"
)

func (g *Game) DisplayHighScore() {
	// display the game over splash to ask for a name for the high score
	g.DisplayCenterBlock([]string{
		"===========================",
		"|                         |",
		"|      NEW HIGH SCORE     |",
		"|                         |",
		"|  Enter your name:       |",
		"|												 |",
		"|												 |",
		"|												 |",
		"|        GAME OVER        |",
		"|                         |",
		"===========================",
		"",
		"Score: " + fmt.Sprint(g.score),
		"",
	}, tcell.ColorDarkRed, tcell.ColorLightYellow)
	g.DisplayCenterBlock([]string{
		g.highScore.name,
	}, tcell.ColorDarkRed, tcell.ColorLightYellow)
}

func (g *Game) DisplayGameOver() {
	// if entering high score show the high score display instead
	if g.score > g.highScore.score {
		g.DisplayHighScore()
		return
	}
	// display the game over message
	g.DisplayCenterBlock([]string{
		"===========================",
		"|                         |",
		"|        GAME OVER        |",
		"|                         |",
		"|  Press 'r' to restart!  |",
		"|    Press 'q' to quit    |",
		"|                         |",
		"===========================",
	}, tcell.ColorDarkRed, tcell.ColorLightYellow)
}