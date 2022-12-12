package main

import (
	"fmt"
	"math"
	"sort"
	"strings"

	"github.com/STollenaar/AdventOfCode2022/internal"
)

type Queue struct {
	internal.Queue[*Point]
}

type Point struct {
	x        int
	y        int
	value    int
	weight   int
	distance int

	previous *Point
}

func (q *Queue) inSlice(currentPoint *Point) bool {
	for _, point := range q.Elements {
		if point.x == currentPoint.x && currentPoint.y == point.y {
			return true
		}
	}
	return false
}

type Node struct {
	height, x, y        int
	start, end, visited bool
}

type Grid map[int]map[int]*Node

const letterValue = "abcdefghijklmnopqrstuvwxyz"

var (
	grid  Grid
	queue Queue
)

func init() {
	grid = make(Grid)

	queue.SortFunction = func(i, j int) bool {
		if queue.Elements[i].weight == queue.Elements[j].weight {
			return queue.Elements[i].value < queue.Elements[j].value
		}
		return queue.Elements[i].weight < queue.Elements[j].weight
	}
}

func main() {
	lines := internal.Reader()

	var startX, startY int
	for r, line := range lines {
		for x, c := range line {
			h := strings.Index(letterValue, string(c))

			if h == -1 {
				if string(c) == "S" {
					h = 0
				} else {
					h = strings.Index(letterValue, "z")
				}
			}

			if _, ok := grid[r]; !ok {
				grid[r] = make(map[int]*Node)
			}
			grid[r][x] = &Node{
				height: h,
				start:  string(c) == "S",
				end:    string(c) == "E",
				x:      x,
				y:      r,
			}
			if string(c) == "S" {
				startX, startY = x, r
			}
		}
	}

	path := findPath(startX, startY)
	steps := findPathLength(path)
	fmt.Printf("Problem 1: %d\n", steps)

	startingLocations := findStartingLocations()

	for _, sl := range startingLocations {
		p := findPath(sl.x, sl.y)
		if p == nil {
			sl.value = math.MaxInt
		} else {
			s := findPathLength(p)
			sl.value = s
		}
	}
	sort.Slice(startingLocations, func(i, j int) bool {
		return startingLocations[i].value < startingLocations[j].value
	})
	fmt.Printf("Problem 2: %d\n", startingLocations[0].value)
}

func resetGrid() {
	for _, r := range grid {
		for _, n := range r {
			n.visited = false
		}
	}
}

func findStartingLocations() (points []*Point) {
	for y, r := range grid {
		for x, c := range r {
			if string(letterValue[c.height]) == "a" {
				points = append(points, &Point{x: x, y: y})
			}
		}
	}
	return points
}

func findPath(startX, startY int) *Point {
	maxY := len(grid)
	maxX := len(grid[maxY-1])

	queue.Push(&Point{
		x:        startX,
		y:        startY,
		value:    0,
		distance: maxX + maxY,
	})

	var iterations int
	for len(queue.Elements) > 0 {
		node := queue.Shift()

		currentNode := grid[node.y][node.x]

		x, y := node.x, node.y
		currentNode.visited = true

		if currentNode.end {
			queue.Empty()
			resetGrid()
			return node
		}

		neighbours := getNeighbours(x, y)

		for _, n := range neighbours {
			nN := grid[n.y][n.x]
			hDiff := nN.height - currentNode.height
			if (hDiff <= 0 || hDiff == 1) && !queue.inSlice(n) {
				n.previous = node
				n.distance = n.getDistance(maxX, maxY)
				n.weight = node.weight + n.value
				n.value = iterations
				queue.Enqueue(n)
			}
		}
		iterations++
		queue.Sort()
	}
	queue.Empty()
	resetGrid()
	return nil
}

func findPathLength(p *Point) (steps int) {
	for p != nil {
		steps++
		p = p.previous
	}
	steps--
	return steps
}

func (p *Point) getDistance(x, y int) int {
	return (x - p.x) + (y - p.y)
}

func getNeighbours(x, y int) (neighbours []*Point) {
	if n, ok := grid[y+1][x]; ok && !n.visited {
		neighbours = append(neighbours, &Point{x: x, y: y + 1})
	}
	if n, ok := grid[y-1][x]; ok && !n.visited {
		neighbours = append(neighbours, &Point{x: x, y: y - 1})
	}
	if n, ok := grid[y][x+1]; ok && !n.visited {
		neighbours = append(neighbours, &Point{x: x + 1, y: y})
	}
	if n, ok := grid[y][x-1]; ok && !n.visited {
		neighbours = append(neighbours, &Point{x: x - 1, y: y})
	}
	return neighbours
}
