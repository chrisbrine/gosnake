package snake

func (g *Game) DisplaySnake(snake Snake) {
	// display the snake
	if (snake.tail.x != snake.head.x || snake.tail.y != snake.head.y) {
		g.PrintCell(snake.tail.x, snake.tail.y, snake.tailChar, snake.tailColor, g.settings.backgroundColor)
	}
	g.PrintCell(snake.head.x, snake.head.y, snake.headChar, snake.headColor, g.settings.backgroundColor)
	for _, point := range snake.body {
		g.PrintCell(point.x, point.y, snake.char, snake.color, g.settings.backgroundColor)
	}
}

func (g *Game) DisplaySnakes() {
	// display the snakes
	for _, snake := range g.snakes {
		g.DisplaySnake(snake)
	}
}
