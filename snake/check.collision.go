package snake

func (g *Game) CheckCollision() bool {
	// check if the snake has collided with itself or an obstacle
	for i, snake := range g.snakes {
		if g.IsPointInSnake(i, snake.head.x, snake.head.y) {
			return true
		}
		if g.IsPointInObstacle(snake.head.x, snake.head.y) {
			return true
		}
	}
	return false
}
