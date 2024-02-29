package snake

func (g *Game) NewObstacle(x int, y int) Obstacle {
	// create a new obstacle
	obstacle := Obstacle{
		location: Point{x, y},
		color: g.RandomColor(g.settings.obstacleColors),
		char: g.RandomChar(g.settings.obstacleChars),
	}
	return obstacle
}

func (g *Game) AddObstacle() {
	// add obstacle to the game
	x, y := g.FindRandomEmptyPoint()
	g.obstacles = append(g.obstacles, g.NewObstacle(x, y))
}

func (g *Game) InitObstacles() {
	// initialize obstacles
	g.obstacles = []Obstacle{}
	for i := 0; i < g.settings.initialObstacles; i++ {
		g.AddObstacle()
	}
}
