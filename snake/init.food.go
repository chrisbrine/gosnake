package snake

func (g *Game) NewFood(x int, y int) Food {
	// create a new food
	food := Food{
		location: Point{x, y},
		color: g.RandomColor(g.settings.foodColors),
		char: g.RandomChar(g.settings.foodChars),
		value: g.RandomFoodValue(),
	}
	return food
}

func (g *Game) AddFood() {
	// add food to the game
	x, y := g.FindRandomEmptyPoint()
	g.food = append(g.food, g.NewFood(x, y))
}

func (g *Game) InitFood() {
	// initialize food
	g.food = []Food{}
	for i := 0; i < g.settings.initialFood; i++ {
		g.AddFood()
	}
}
