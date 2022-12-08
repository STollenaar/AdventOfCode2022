package main

import (
	"fmt"
	"strconv"

	"github.com/STollenaar/AdventOfCode2022/internal"
)

type Grid struct {
	internal.Grid[int]
}

var (
	treeMap Grid
)

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
			if visibleFromEdge(x, y, -1, 0) ||
				visibleFromEdge(x, y, 1, 0) ||
				visibleFromEdge(x, y, 0, -1) ||
				visibleFromEdge(x, y, 0, 1) {
				visible++
			}
		}
	}
	return visible
}

func visibleFromEdge(x, y, modX, modY int) bool {
	r := treeMap.Rows[y]
	c := r[x]

	if modX != 0 {
		if modX > 0 {
			for i := x + 1; i < len(r); i++ {
				if c <= r[i] {
					return false
				}
			}
		} else {
			for i := x - 1; i >= 0; i-- {
				if c <= r[i] {
					return false
				}
			}
		}
	} else {
		if modY > 0 {
			for i := y + 1; i < len(treeMap.Rows); i++ {
				if c <= treeMap.Rows[i][x] {
					return false
				}
			}
		} else {
			for i := y - 1; i >= 0; i-- {
				if c <= treeMap.Rows[i][x] {
					return false
				}
			}
		}
	}

	return true
}

func maxVisibleTreeSpot() (amount int) {
	for y, r := range treeMap.Rows {
		for x := range r {
			aV := visibleTreesFromPoint(x, y)
			if aV > amount {
				amount = aV
			}
		}
	}
	return amount
}

func visibleTreesFromPoint(x, y int) (amount int) {
	r := treeMap.Rows[y]
	c := r[x]

	var n, e, w, s int
	for i := x - 1; i >= 0; i-- {
		e++
		if c <= r[i] {
			break
		}
	}
	for i := x + 1; i < len(r); i++ {
		w++
		if c <= r[i] {
			break
		}
	}
	for i := y - 1; i >= 0; i-- {
		n++
		if c <= treeMap.Rows[i][x] {
			break
		}
	}
	for i := y + 1; i < len(treeMap.Rows); i++ {
		s++
		if c <= treeMap.Rows[i][x] {
			break
		}
	}
	return n * e * w * s
}
