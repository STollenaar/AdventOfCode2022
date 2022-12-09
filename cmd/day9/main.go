package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/STollenaar/AdventOfCode2022/internal"
)

type Grid map[string]string

type Pos struct {
	child *Pos
	x, y  int
}

func main() {
	lines := internal.Reader()

	for _, line := range lines {
		direction, a := strings.Split(line, " ")[0], strings.Split(line, " ")[1]
		amount, _ := strconv.Atoi(a)
		for amount > 0 {
			var x, y int
			switch direction {
			case "R":
				x = 1
			case "U":
				y = 1
			case "D":
				y = -1
			case "L":
				x = -1
			}
			problem1Move(x, y)
			problem2Move(x, y)
			amount--
		}
	}
	fmt.Printf("Problem 1: %d\n", len(gridP1))
	fmt.Printf("Problem 2: %d\n", len(gridP2))
}

func (p *Pos) move(x, y int, grid Grid, isTail bool) {
	p.x += x
	p.y += y

	if isTail {
		pos := fmt.Sprintf("%v-%v", p.x, p.y)
		grid[pos] = "#"
	}
}

func (p *Pos) getDiff(o Pos) (difX, difY int) {
	difX = p.x - o.x
	difY = p.y - o.y

	return difX, difY
}

func (p *Pos) moveCloser(o Pos, grid Grid) {
	for !p.isTouching(o) {
		difX, difY := o.getDiff(*p)
		if difX > 1 {
			difX = 1
		}
		if difY > 1 {
			difY = 1
		}
		if difX < -1 {
			difX = -1
		}
		if difY < -1 {
			difY = -1
		}
		p.move(difX, difY, grid, (p == &tailP1 || p == tailP2))
	}
}

func (p *Pos) isTouching(o Pos) bool {
	difX, difY := p.getDiff(o)
	difX = int(math.Abs(float64(difX)))
	difY = int(math.Abs(float64(difY)))
	if difX <= 1 && difY <= 1 {
		return true
	}
	return false
}
