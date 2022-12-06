package main

import (
	"fmt"
	"strings"

	"github.com/STollenaar/AdventOfCode2022/internal"
)

func main() {
	line := internal.Reader()[0]

	p1S := findMarker(line, 4)
	fmt.Printf("Problem 1 solution: %d\n", p1S)

	p2S := findMarker(line, 14)
	fmt.Printf("Problem 2 solution: %d\n", p2S)

}

func findMarker(line string, amount int) int {
	for i := 0; i+amount < len(line); i++ {
		subs := line[i : i+amount]
		if !containsDup(subs) {
			return i + amount
		}
	}
	return -1
}

func containsDup(sub string) bool {
	for _, c := range sub {
		if strings.Count(sub, string(c)) > 1 {
			return true
		}
	}
	return false
}
