package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/STollenaar/AdventOfCode2022/internal"
	"golang.org/x/exp/slices"
)

type Grid struct {
	internal.Grid[string]
}

var (
	grid1 *Grid
	grid2 *Grid
)

type Move struct {
	amount, fromX, toX int
}

func init() {
	grid1 = new(Grid)
	grid2 = new(Grid)
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
				grid1.AddRow(row)
				grid2.AddRow(slices.Clone(row))
			} else {
				continue
			}
		} else {
			move := createMove(line)
			makeMove(move, grid1, false)
			makeMove(move, grid2, true)
		}
	}
	fmt.Println(getTopCrates(*grid1))
	fmt.Println(getTopCrates(*grid2))
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

func makeMove(move Move, grid *Grid, fromTop bool) {
	startRowH := grid.GetHeight(move.fromX)
	endRowH := grid.GetHeight(move.toX)

	var j int
	var shifted int
	for i := move.amount; i > 0; i-- {
		startY := startRowH + shifted
		if fromTop {
			startY += (i - 1)
		} else {
			startY += j
		}
		crate := grid.GetSafeColumn(move.fromX, startY)
		grid.SetSafeColumn("", move.fromX, startY)

		endY := endRowH - j - 1

		// Shifting things to fit into the grid
		if endY < 0 {
			endY = 0
			t := make([]string, 9)
			shifted++
			grid.ShiftRow(t)
		}

		grid.SetSafeColumn(crate, move.toX, endY)
		j++
	}
}

func getTopCrates(grid Grid) (top string) {
	var ignoreX []int
	crates := make(map[int]string)
	for _, row := range grid.Rows {
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
