package snake

import (
	"math/rand"

	"github.com/gdamore/tcell"
)

type Point struct {
	// The x and y coordinates of the point
	x, y int
}

type Snake struct {
	// The snake's body
	body []Point
	// The direction the snake is moving
	direction Point
	// The size of the snake
	size int
	// The speed of the snake
	speed int
	// The snake's score
	score int
	// The snake's colour
	color tcell.Color
	// the snake's head colour
	headColor tcell.Color
	// the snake's tail colour
	tailColor tcell.Color
	// The snake's head
	head Point
	// The snake's tail
	tail Point
	// The snake's character
	char rune
	// the snake's head character
	headChar rune
	// the snake's tail character
	tailChar rune
}

type Food struct {
	// The food's location
	location Point
	// The food's color
	color tcell.Color
	// The food's character
	char rune
	// The food's score value
	value int
}

type Obstacle struct {
	// The obstacle's location
	location Point
	// The obstacle's color
	color tcell.Color
	// the obstacle's character
	char rune
}

type Settings struct {
	// The initial number of snakes
	initialSnakes int
	// The initial number of food
	initialFood int
	// The initial number of obstacles
	initialObstacles int
	// The game's initial speed
	initialSpeed int
	// The game's initial size
	initialSize int
	// The game's amount to add to the size
	sizeIncrease int
	// The amount to increase the score by for eating food
	scoreIncrease int
	// The amount to increase the score by just surviving each moment
	survivalScoreIncrease int
	// The game's paused status
	paused bool
	// The game's default foreground colour
	foregroundColor tcell.Color
	// The game's default background colour
	backgroundColor tcell.Color
	// The game's default snake head character
	snakeHeadChars []rune
	// The game's default snake character
	snakeChars []rune
	// The game's default snake tail character
	snakeTailChars []rune
	// The game's set of possible colours for the snake
	snakeColors []tcell.Color
	// The game's set of possible runes for food
	foodChars []rune
	// The game's set of possible colours for food
	foodColors []tcell.Color
	// The game's set of possible runes for obstacles
	obstacleChars []rune
	// The game's set of possible colours for obstacles
	obstacleColors []tcell.Color
	// Food's max score value
	foodMaxValue int
	// Food's min score value
	foodMinValue int
}

type Tools struct {
	// the game's random number generator
	rand *rand.Rand
	// The game's screen
	screen tcell.Screen
}

type HighScore struct {
	// The high score's name
	name string
	// The high score's score
	score int
}

type Game struct {
	// The game's snakes
	snakes []Snake
	// The game's food
	food []Food
	// The game's obstacles
	obstacles []Obstacle
	// the game's settings
	settings Settings
	// the game's tools
	tools Tools
	// if the game is over
	gameOver bool
	// food count
	foodCount int
	// game score
	score int
	// debug string
	debugString string
	// end game boolean
	endGame bool
	// Game Speed
	gameSpeed int
	// Game started boolean
	gameStarted bool
	// The high score
	highScore HighScore
}

func NewGame() Game {
	// create a new game
	game := Game{}
	game.InitGame()
	return game
}
