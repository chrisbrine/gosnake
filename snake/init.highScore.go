package snake

func (g *Game) InitHighScore() {
	// initialize the high score
	score, name := g.GetMaxScore()
	g.highScore.score = score
	g.highScore.name = name
}