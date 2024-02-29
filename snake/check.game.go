package snake

func (g *Game) IsPointEmpty(x int, y int) bool {
	// check food, obstacles and snakes if any are at that point
	if g.IsPointInFood(x, y) || g.IsPointInObstacle(x, y) || g.IsPointInSnake(-1, x, y) {
		return false
	}
	return true
}

func (g *Game) UpdateGame() {
	// update the game
	if !g.gameOver && !g.settings.paused {
		g.UpdateFood()
		g.UpdateObstacles()
		g.UpdateSnakes()
		g.gameOver = g.gameOver || g.CheckCollision()
	}
}
