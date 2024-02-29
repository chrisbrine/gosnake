package snake

import (
	"os"
	"strconv"
	"strings"
)

func fileName() string {
	path, err := os.Getwd()
	if err != nil {
		return "score.txt"
	}
	return path + "/score.txt"
}

func (g *Game) GetMaxScore() (int, string) {
	// open the filename and get the max score and name of person that got it from it
	// if the file does not exist, then just return 0 and a blank name
	maxScore := 0
	maxName := ""
	file := fileName()
	state, err := os.ReadFile(file)
	if err != nil {
		return maxScore, maxName
	}
	// split the string into score and name
	parts := strings.Split(string(state), ",")
	if len(parts) != 2 {
		return maxScore, maxName
	}
	maxScore, err = strconv.Atoi(parts[0])
	if err != nil {
		return 0, ""
	}
	maxName = parts[1]
	return maxScore, maxName
}

func (g *Game) SetMaxScore(score int, name string) {
	// write the max score and name to the file
	state := strconv.Itoa(score) + "," + name
	file := fileName()
	f, fErr := os.Create(file)
	if fErr != nil {
		return
	}
	_, err := f.WriteString(state)
	if err != nil {
		return
	}
	err = f.Close()
	if err != nil {
		return
	}
}
