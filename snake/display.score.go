package snake

import (
	"fmt"
)

func (g *Game) DisplayScore() {
	// display the score
	g.PrintTextTop(1, "Score: "+fmt.Sprint(g.score) + " High Score: " + fmt.Sprint(g.highScore.score) + " " + g.highScore.name)
}