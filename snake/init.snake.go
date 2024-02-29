package snake

func (g *Game) NewSnake(x int, y int) Snake {
	mainColor := g.RandomColor(g.settings.snakeColors)
	return Snake{
		size: g.settings.initialSize,
		speed: g.settings.initialSpeed,
		score: 0,
		color: mainColor,
		headColor: g.RandomColor(g.settings.snakeColors),
		direction: Point{0, 0},
		tailColor: mainColor,
		head: Point{x, y},
		tail: Point{x, y},
		char: g.RandomChar(g.settings.snakeChars),
		headChar: g.RandomChar(g.settings.snakeHeadChars),
		tailChar: g.RandomChar(g.settings.snakeTailChars),
	}

}

func (g *Game) AddSnake() {
	x, y := g.FindRandomEmptyPoint()
	g.snakes = append(g.snakes, g.NewSnake(x, y))
}

func (g *Game) InitSnake() {
	g.snakes = []Snake{}
	for i := 0; i < g.settings.initialSnakes; i++ {
		g.AddSnake()
	}
}