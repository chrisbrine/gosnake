package snake

import (
	"math/rand"
	"time"

	"github.com/gdamore/tcell"
)

func (g *Game) InitTools() {
	g.tools.rand = rand.New(rand.NewSource(time.Now().UnixNano()))
	screen, err := tcell.NewScreen()
	if err != nil {
		panic(err)
	}
	if err := screen.Init(); err != nil {
		panic(err)
	}
	g.tools.screen = screen
}

func (g *Game) RandomColor(colors []tcell.Color) tcell.Color {
	return colors[g.tools.rand.Intn(len(colors))]
}

func (g *Game) RandomChar(chars []rune) rune {
	return chars[g.tools.rand.Intn(len(chars))]
}

func (g *Game) FindRandomEmptyPoint() (int, int) {
	width, height := g.tools.screen.Size()
	for {
		x := g.tools.rand.Intn(width - 3) + 3
		y := g.tools.rand.Intn(height - 3) + 3
		if g.IsPointEmpty(x, y) {
			return x, y
		}
	}
}

func (g *Game) RandomFoodValue() int {
	return g.tools.rand.Intn(g.settings.foodMaxValue-g.settings.foodMinValue) + g.settings.foodMinValue
}
