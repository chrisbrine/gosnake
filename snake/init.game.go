package snake

func (g *Game) InitGame() {
	g.gameOver = false
	g.endGame = false
	g.gameStarted = false
	g.foodCount = g.settings.initialFood
	g.score = 0
	g.debugString = ""
	g.gameSpeed = 100

	g.InitTools()
	g.InitSettings()
	g.InitSnake()
	g.InitFood()
	g.InitObstacles()
	g.InitHighScore()
}
