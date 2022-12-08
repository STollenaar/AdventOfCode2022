package main

import (
	"fmt"
	"strconv"

	"github.com/STollenaar/AdventOfCode2022/internal"
)

type Grid struct {
	internal.Grid[int]
}

var treeMap Grid

func main() {
	lines := internal.Reader()

	for i, line := range lines {
		for _, c := range line {
			h, _ := strconv.Atoi(string(c))
			treeMap.AddSafeToColumn(h, i)
		}
	}
	fmt.Printf("Problem 1: %d\n", calcVisible())
	fmt.Printf("Problem 2: %d\n", maxVisibleTreeSpot())
}

func calcVisible() (visible int) {
	for y, r := range treeMap.Rows {
		for x := range r {
			a, _ := visibleFromEdge(x, y, -1, 0)
			b, _ := visibleFromEdge(x, y, 1, 0)
			c, _ := visibleFromEdge(x, y, 0, -1)
			d, _ := visibleFromEdge(x, y, 0, 1)
			if a || b || c || d {
				visible++
			}
		}
	}
	return visible
}

func visibleFromEdge(x, y, modX, modY int) (b bool, it int) {
	r := treeMap.Rows[y]
	c := r[x]

	if modX != 0 {
		for i := x + modX; i < len(r) && i >= 0; i += modX {
			it++
			if c <= r[i] {
				return false, it
			}
		}
	} else {
		for i := y + modY; i < len(treeMap.Rows) && i >= 0; i += modY {
			it++
			if c <= treeMap.Rows[i][x] {
				return false, it
			}
		}
	}

	return true, it
}

func maxVisibleTreeSpot() (amount int) {
	for y, r := range treeMap.Rows {
		for x := range r {
			aV := areaVisible(x, y)
			if aV > amount {
				amount = aV
			}
		}
	}
	return amount
}

func areaVisible(x, y int) (amount int) {
	_, e := visibleFromEdge(x, y, -1, 0)
	_, w := visibleFromEdge(x, y, 1, 0)
	_, n := visibleFromEdge(x, y, 0, -1)
	_, s := visibleFromEdge(x, y, 0, 1)
	return n * e * w * s
}
