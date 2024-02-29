package snake

import (
	"time"
)

func (g *Game) Run() {
	// run the game
	inputChan := g.GetInput()
	lastTimeIncrease := time.Now()
	for {
		// display the game
		g.DisplayGame()

		time.Sleep(time.Duration(g.gameSpeed) * time.Millisecond)
		if time.Since(lastTimeIncrease) > time.Duration(1) * time.Second &&
		!g.gameOver && !g.endGame && g.gameStarted && !g.settings.paused && g.IsMoving() {
			g.score += g.settings.survivalScoreIncrease
		}

		// get the user input
		var option string
		select {
			case option = <-inputChan:
			default:
				option = ""
		}

		if option == "exit" {
			g.tools.screen.Fini()
			break
		}
		g.HandleInput(option)
		// update the game
		g.UpdateGame()
	}
	// display the game over message
	// g.DisplayGameOver()
}