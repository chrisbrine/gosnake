package snake

import "github.com/gdamore/tcell"

func (g *Game) PrintCell(x int, y int, char rune, color tcell.Color, bgColor tcell.Color) {
	// print a cell
	g.tools.screen.SetContent(x, y, char, nil, tcell.StyleDefault.Foreground(color).Background(bgColor))
}

func (g *Game) PrintTextColor(x int, y int, text string, color tcell.Color, bgColor tcell.Color) {
	// print text
	for i, char := range text {
		g.PrintCell(x+i, y, char, color, bgColor)
	}
}

func (g *Game) PrintText(x int, y int, text string) {
	// print text
	g.PrintTextColor(x, y, text, g.settings.foregroundColor, g.settings.backgroundColor)
}

func (g * Game) PrintTextTop(x int, text string) {
	// print text at the top
	g.PrintTextColor(x, 0, text, g.settings.backgroundColor, g.settings.foregroundColor)
}

func (g *Game) DisplayCenter(y int, text string, color tcell.Color, bgColor tcell.Color) {
	// display text in the center
	width, _ := g.tools.screen.Size()
	g.PrintTextColor((width-len(text))/2, y, text, color, bgColor)
}

func (g *Game) DisplayCenterBlock(texts []string, color tcell.Color, bgColor tcell.Color) {
	// display a block of text in the center
	_, height := g.tools.screen.Size()
	pos := (height/2)-(len(texts)/2)
	for i, text := range texts {
		g.DisplayCenter(pos+i, text, color, bgColor)
	}
}

func (g *Game) TopBar() {
	width, _ := g.tools.screen.Size()
	for i := 0; i < width; i++ {
		g.PrintCell(i, 0, ' ', g.settings.backgroundColor, g.settings.foregroundColor)
	}
}

func (g *Game) DisplayGame() {
	// display the game
	g.tools.screen.Clear()
	g.TopBar()
	g.DisplayScore()
	g.DisplayTitle()
	g.DisplaySnakes()
	g.DisplayFood()
	g.DisplayObstacles()
	if g.debugString != "" {
		g.PrintText(5, 5, g.debugString)
	}
	if g.gameOver {
		g.DisplayGameOver()
	} else if g.settings.paused {
		g.DisplayPaused()
	} else if !g.gameStarted {
		g.DisplayInstructions()
	}
	g.tools.screen.Show()
}