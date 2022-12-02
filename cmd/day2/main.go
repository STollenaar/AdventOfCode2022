package main

import (
	"strings"

	internal "github.com/STollenaar/AdventOfCode2022/internal"
)

type Strat struct {
	winMove  string
	drawMove string
	loseMove string
}

var (
	stratMap1 map[string]Strat
	stratMap2 map[string]Strat
	valueMap  map[string]int

	problem1Score int
	problem2Score int
)

func init() {
	stratMap1 = map[string]Strat{
		"A": {
			winMove:  "Y",
			drawMove: "X",
		},
		"B": {
			winMove:  "Z",
			drawMove: "Y",
		},
		"C": {
			winMove:  "X",
			drawMove: "Z",
		},
	}

	stratMap2 = map[string]Strat{
		"A": {
			winMove:  "B",
			loseMove: "C",
		},
		"B": {
			winMove:  "C",
			loseMove: "A",
		},
		"C": {
			winMove:  "A",
			loseMove: "B",
		},
	}

	valueMap = map[string]int{
		"X": 1,
		"Y": 2,
		"Z": 3,
		"A": 1,
		"B": 2,
		"C": 3,
	}
}

func main() {
	lines := internal.Reader()

	for _, line := range lines {
		problem1(line)
		problem2(line)
	}

	fmt.Printf("Solution problem1: %d\n", problem1Score)
	fmt.Printf("Solution problem2: %d\n", problem2Score)
}

func problem1(line string) {
	moves := strings.Split(line, " ")

	strat := stratMap1[moves[0]]

	problem1Score += valueMap[moves[1]]
	if strat.winMove == moves[1] {
		problem1Score += 6
	} else if strat.drawMove == moves[1] {
		problem1Score += 3
	}
}

func problem2(line string) {
	moves := strings.Split(line, " ")

	strat := stratMap2[moves[0]]

	switch moves[1] {
	case "X":
		problem2Score += valueMap[strat.loseMove]
	case "Y":
		problem2Score += valueMap[moves[0]]
		problem2Score += 3
	case "Z":
		problem2Score += valueMap[strat.winMove]
		problem2Score += 6
	}
}
