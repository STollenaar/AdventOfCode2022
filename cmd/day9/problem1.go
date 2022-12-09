package main

var (
	gridP1 Grid
	headP1 Pos
	tailP1 Pos
)

func init() {
	gridP1 = make(Grid)
	gridP1["0-0"] = "#"
	headP1 = Pos{x: 0, y: 0}
	tailP1 = Pos{x: 0, y: 0}
}

func problem1Move(x, y int) {
	headP1.move(x, y, gridP1, false)
	if !tailP1.isTouching(headP1) {
		tailP1.moveCloser(headP1, gridP1)
	}
}
