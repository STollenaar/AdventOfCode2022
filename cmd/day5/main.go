package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/STollenaar/AdventOfCode2022/internal"
	"golang.org/x/exp/slices"
)

var (
	gridP1 [][]string
	gridP2 [][]string
)

type Move struct {
	amount, fromX, toX int
}

func main() {
	lines := internal.Reader()
	var moves bool

	for _, line := range lines {
		if line == "" {
			moves = true
			continue
		}
		if !moves {
			if strings.Contains(line, "[") {
				line = strings.ReplaceAll(line, "    ", " ") // WHY 4 SPACES, WHYYYYYY, DAMN CLEANUP
				// Removing boxes
				line = strings.ReplaceAll(line, "[", "")
				line = strings.ReplaceAll(line, "]", "")

				row := strings.Split(line, " ")
				gridP1 = append(gridP1, row)
				gridP2 = append(gridP2, slices.Clone(row))
			} else {
				continue
			}
		} else {
			move := createMove(line)
			gridP1 = makeMove(move, gridP1, false)
			gridP2 = makeMove(move, gridP2, true)
		}
	}
	fmt.Println(getTopCrates(gridP1))
	fmt.Println(getTopCrates(gridP2))
}

func getHeight(x int, grid [][]string) int {
	for y, l := range grid {
		if l[x] != "" {
			return y
		}
	}
	return len(grid) - 1
}

func createMove(move string) Move {
	m := strings.Split(move, " ")
	amount, _ := strconv.Atoi(m[1])
	fx, _ := strconv.Atoi(m[3])
	tx, _ := strconv.Atoi(m[5])

	return Move{
		amount: amount,
		fromX:  fx - 1,
		toX:    tx - 1,
	}
}

func makeMove(move Move, grid [][]string, fromTop bool) [][]string {
	startRowH := getHeight(move.fromX, grid)
	endRowH := getHeight(move.toX, grid)

	var j int
	var shifted int
	for i := move.amount; i > 0; i-- {
		startY := startRowH + shifted
		if fromTop {
			startY += (i - 1)
		} else {
			startY += j
		}
		crate := grid[startY][move.fromX]
		grid[startY][move.fromX] = ""

		endY := endRowH - j - 1

		// Shifting things to fit into the grid
		if endY < 0 {
			endY = 0
			t := make([]string, 9)
			shifted++
			grid = append([][]string{t}, grid...)
		}

		grid[endY][move.toX] = crate
		j++
	}
	return grid
}

func getTopCrates(grid [][]string) (top string) {
	var ignoreX []int
	crates := make(map[int]string)
	for _, row := range grid {
		for x, column := range row {
			if slices.Contains(ignoreX, x) {
				continue
			}
			if column != "" {
				ignoreX = append(ignoreX, x)
				crates[x] = column
			}
		}
	}
	sort.Ints(ignoreX)
	for _, x := range ignoreX {
		top += crates[x]
	}
	return top
}
