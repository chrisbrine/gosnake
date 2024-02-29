package snake

func (g *Game) IsPointInObstacle(x int, y int) bool {
	// check if the point is an obstacle
	for _, obstacle := range g.obstacles {
		if obstacle.location.x == x && obstacle.location.y == y {
			return true
		}
	}
	return false
}

func (g *Game) UpdateObstacles() {
	var snakeSize int = 0
	width, height := g.tools.screen.Size()
	sizeCalc := (width * height) / 1000
	for _, snake := range g.snakes {
		snakeSize += len(snake.body)
	}
	obstaclesNeeded := (snakeSize / 6) + sizeCalc
	if len(g.obstacles) < obstaclesNeeded {
		g.AddObstacle()
	}
}
