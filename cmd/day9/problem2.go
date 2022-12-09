package main

var (
	gridP2 Grid
	headP2 *Pos
	knot1  *Pos
	knot2  *Pos
	knot3  *Pos
	knot4  *Pos
	knot5  *Pos
	knot6  *Pos
	knot7  *Pos
	knot8  *Pos
	tailP2 *Pos
)

// Here is another rage comment... Fuck this initialization
func init() {
	gridP2 = make(Grid)
	gridP2["0-0"] = "#"
	tailP2 = &Pos{x: 0, y: 0}
	knot8 = &Pos{x: 0, y: 0, child: tailP2}
	knot7 = &Pos{x: 0, y: 0, child: knot8}
	knot6 = &Pos{x: 0, y: 0, child: knot7}
	knot5 = &Pos{x: 0, y: 0, child: knot6}
	knot4 = &Pos{x: 0, y: 0, child: knot5}
	knot3 = &Pos{x: 0, y: 0, child: knot4}
	knot2 = &Pos{x: 0, y: 0, child: knot3}
	knot1 = &Pos{x: 0, y: 0, child: knot2}
	headP2 = &Pos{x: 0, y: 0, child: knot1}
}

func problem2Move(x, y int) {
	headP2.move(x, y, gridP2, false)

	// Moving the knots starting from the head
	c := headP2.child
	prev := headP2
	for c != nil {
		if !c.isTouching(*prev) {
			c.moveCloser(*prev, gridP2)
		}
		prev = c
		c = c.child
	}
}
