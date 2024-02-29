package snake

import (
	"math"
)

func (g *Game) SnakeDirection(x int, y int) {
	// if gameOver then don't change the snake direction
	if g.gameOver || g.settings.paused || !g.gameStarted {
		return
	}
	// change the snake direction
	for i := range g.snakes {
		// check if looking to move in the opposite of the current direction, if so then don't change
		if (math.Abs(float64(g.snakes[i].direction.x)) == math.Abs(float64(x)) &&
		math.Abs(float64(g.snakes[i].direction.y)) == math.Abs(float64(y))) {
			continue
		}
		g.snakes[i].direction.x = x
		g.snakes[i].direction.y = y
	}
}

func (g *Game) IsPointInSnake(snakeId int, x int, y int) bool {
	// check if the point is in the snake
	for i, snake := range g.snakes {
		if snake.head.x == x && snake.head.y == y && i != snakeId {
			return true
		}
		for _, point := range snake.body {
			if point.x == x && point.y == y {
				return true
			}
		}
	}
	return false
}

func (g *Game) IsMoving() bool {
	// check if the snake is moving
	for _, snake := range g.snakes {
		if snake.direction.x != 0 || snake.direction.y != 0 {
			return true
		}
	}
	return false
}

func (s *Snake) MoveSnake(x int, y int) Snake {
	// move the snake
	s.body = append([]Point{{s.head.x, s.head.y}}, s.body...)
	s.head.x = x
	s.head.y = y
	if s.size < len(s.body) {
		s.body = s.body[:s.size]
	}
	// move tail
	s.tail = s.body[len(s.body)-1]
	return *s
}

func (g *Game) UpdateSnakes() {
	// update the snake position
	width, height := g.tools.screen.Size()
	for i, snake := range g.snakes {
		// move snake by direction, but check position for food and obstacles
		if (snake.direction.x == 0 && snake.direction.y == 0) {
			continue
		}
		newX := snake.head.x + snake.direction.x
		newY := snake.head.y + snake.direction.y
		// check if snake is out of bounds
		if newX < 0 || newX >= width || newY < 1 || newY >= height {
			g.gameOver = true
			return
		}
		if g.IsPointInFood(newX, newY) {
			// add to snake body
			g.snakes[i].size += g.settings.sizeIncrease
			snake = g.snakes[i]
			// remove food
			g.RemoveFood(newX, newY)
		}
		if g.IsPointEmpty(newX, newY) {
			// move snake
			g.snakes[i] = snake.MoveSnake(newX, newY)
		} else {
			// game over
			g.gameOver = true
		}
	}
}
