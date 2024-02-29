package snake

import (
	"github.com/gdamore/tcell"
)

func (g *Game) GetInput() chan string {
	inputChan := make(chan string)
	go func() {
		for {
			switch ev := g.tools.screen.PollEvent().(type) {
				case *tcell.EventKey:
					switch ev.Key() {
						case tcell.KeyEscape:
							inputChan <- "exit"
						case tcell.KeyUp:
							inputChan <- "up"
						case tcell.KeyDown:
							inputChan <- "down"
						case tcell.KeyLeft:
							inputChan <- "left"
						case tcell.KeyRight:
							inputChan <- "right"
						case tcell.KeyCtrlC:
							inputChan <- "exit"
						case tcell.KeyEnter:
							// confirm if entering high score name and if enter then return
							if g.gameOver && g.score > g.highScore.score {
								inputChan <- "enter"
							}
						case tcell.KeyBackspace, tcell.KeyBackspace2:
							// confirm if entering high score name and if backspace then return
							if g.gameOver && g.score > g.highScore.score {
								inputChan <- "backspace"
							}
						case tcell.KeyRune:
							// confirm if entering high score name and if a character then return
							if g.gameOver && g.score > g.highScore.score {
								if ev.Rune() != 0 {
									inputChan <- string(ev.Rune())
								}
							} else {
								if ev.Rune() == 'p' || ev.Rune() == 'P' {
									inputChan <- "pause"
								}
								if ev.Rune() == 'r' || ev.Rune() == 'R' {
									inputChan <- "restart"
								}
								if ev.Rune() == '+' || ev.Rune() == '=' {
									inputChan <- "speedUp"
								}
								if ev.Rune() == '-' || ev.Rune() == '_' {
									inputChan <- "speedDown"
								}
								if ev.Rune() == 'q' || ev.Rune() == 'Q' {
									inputChan <- "exit"
								}
							}
					}
			}
		}
	}()

	return inputChan
}

