package snake

func (g *Game) IsPointInFood(x int, y int) bool {
	// check if the point is in the food
	for _, food := range g.food {
		if food.location.x == x && food.location.y == y {
			return true
		}
	}
	return false
}

func (g *Game) RemoveFood(x int, y int) {
	// remove food from the game
	for i, food := range g.food {
		if food.location.x == x && food.location.y == y {
			g.score += food.value + g.settings.scoreIncrease
			g.food = append(g.food[:i], g.food[i+1:]...)
		}
	}
}

func (g *Game) UpdateFood() {
	// check if the game needs more food
	width, height := g.tools.screen.Size()
	sizeCalc := (width * height) / 1000
	if len(g.food) < (g.foodCount + sizeCalc) {
		g.AddFood()
	}
}
