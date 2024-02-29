package snake

func (g *Game) DisplayFood() {
	// display the food
	for _, food := range g.food {
		g.PrintCell(food.location.x, food.location.y, food.char, food.color, g.settings.backgroundColor)
	}
}
