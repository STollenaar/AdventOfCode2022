package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/STollenaar/AdventOfCode2022/internal"
)

type Range struct {
	min, max int
}

func main() {
	lines := internal.Reader()

	var totalOverlaps, partialOverlaps int
	for _, line := range lines {
		pairs := strings.Split(line, ",")

		p1, p2 := createRange(pairs[0]), createRange(pairs[1])
		if totalOverlap(p1, p2) {
			totalOverlaps++
		}
		if partialOverlap(p1, p2) {
			partialOverlaps++
		}
	}
	fmt.Printf("Problem 1 solution: %d\n", totalOverlaps)
	fmt.Printf("Problem 2 solution: %d\n", partialOverlaps)
}

func createRange(r string) Range {
	ends := strings.Split(r, "-")
	min, _ := strconv.ParseInt(ends[0], 10, 64)
	max, _ := strconv.ParseInt(ends[1], 10, 64)

	return Range{
		min: int(min),
		max: int(max),
	}
}

// checking for totalOverlap
func totalOverlap(p1, p2 Range) bool {
	if (p1.min <= p2.min && p2.max <= p1.max) ||
		(p2.min <= p1.min && p1.max <= p2.max) {
		return true
	}
	return false
}

// checking for totalOverlap, or if a point would overlap
func partialOverlap(p1, p2 Range) bool {
	if totalOverlap(p1, p2) ||
		(p2.max >= p1.min && p2.min <= p1.max) ||
		(p1.max >= p2.min && p1.min <= p2.max) {
		return true
	}
	return false
}
