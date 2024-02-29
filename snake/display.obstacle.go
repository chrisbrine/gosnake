package snake

func (g *Game) DisplayObstacles() {
	// display the food
	for _, obstacle := range g.obstacles {
		g.PrintCell(obstacle.location.x, obstacle.location.y, obstacle.char, obstacle.color, g.settings.backgroundColor)
	}
}
