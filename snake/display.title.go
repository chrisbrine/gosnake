package snake

func (g *Game) DisplayTitle() {
	// display the title
	width, _ := g.tools.screen.Size()
	title := "SNAKE"
	author := "By Chris Brine"
	g.PrintTextTop((width-len(title))/2, title)
	g.PrintTextTop((width-len(author))-2, author)
}