package snake

func (g *Game) HandleInput(input string) {
	if input != "" {
		g.settings.paused = false;
		if !g.gameStarted {
			g.gameStarted = true
			input = ""
		}
	}

	switch input {
		case "up":
			g.SnakeDirection(0, -1)
		case "down":
			g.SnakeDirection(0, 1)
		case "left":
			g.SnakeDirection(-1, 0)
		case "right":
			g.SnakeDirection(1, 0)
		case "pause":
			g.settings.paused = true
		case "restart":
			if g.gameOver {
				g.RestartGame()
			}
		case "speedUp":
			g.gameSpeed -= 20
			if g.gameSpeed < 0 {
				g.gameSpeed = 0
			}
		case "speedDown":
			g.gameSpeed += 20
			if g.gameSpeed > 200 {
				g.gameSpeed = 200
			}
		default:
			if g.gameOver && g.score > g.highScore.score {
				if input == "enter" {
					g.highScore.score = g.score
					g.SetMaxScore(g.highScore.score, g.highScore.name)
				} else if input == "backspace" {
					g.highScore.name = g.highScore.name[:len(g.highScore.name)-1]
				} else {
					g.highScore.name += input
				}
			}
	}
}