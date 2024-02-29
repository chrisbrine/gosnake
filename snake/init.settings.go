package snake

import (
	"github.com/gdamore/tcell"
)

func (g *Game) InitSettings() {
	g.settings = Settings{
		initialSpeed: 100,
		initialSize: 4,
		initialFood: 3,
		initialSnakes: 1,
		sizeIncrease: 8,
		scoreIncrease: 10,
		survivalScoreIncrease: 1,
		initialObstacles: 3,
		paused: false,
		foregroundColor: tcell.ColorWhite,
		backgroundColor: tcell.ColorBlack,
		// snakeHeadChars: []rune{'▀', '▄', '▐', '▌', '▚', '▞', '▙', '▟'},
		snakeHeadChars: []rune{'▟'},
		// snakeChars: []rune{'█', '▓', '▒', '░'},
		snakeChars: []rune{'█'},
		// snakeTailChars: []rune{'▀', '▄', '▐', '▌', '▚', '▞', '▙', '▟'},
		snakeTailChars: []rune{'▞'},
		snakeColors: []tcell.Color{
			tcell.ColorRed,
			tcell.ColorGreen,
			tcell.ColorBlue,
			tcell.ColorYellow,
			tcell.ColorOrange,
			tcell.ColorPurple,
			tcell.ColorAqua,
			tcell.ColorFuchsia,
		},
		// foodChars: []rune{'♥', '♦', '♣', '♠', '•', '◘', '○', '◙'},
		foodChars: []rune{'♥'},
		foodColors: []tcell.Color{
			tcell.ColorDarkRed,
			tcell.ColorDarkGreen,
			tcell.ColorDarkKhaki,
			tcell.ColorDarkOliveGreen,
		},
		// obstacleChars: []rune{'▓', '▒', '░', '█'},
		obstacleChars: []rune{'X', '█'},
		obstacleColors: []tcell.Color{
			tcell.ColorDarkSlateGray,
			tcell.ColorDimGray,
			tcell.ColorGray,
			tcell.ColorGhostWhite,
			tcell.ColorLightSlateGray,
		},
		foodMaxValue: 50,
		foodMinValue: 5,
	}
	style := tcell.StyleDefault.Background(g.settings.backgroundColor).Foreground(g.settings.foregroundColor)
	g.tools.screen.SetStyle(style)
}