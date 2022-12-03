package main

import (
	"fmt"
	"strings"

	"github.com/STollenaar/AdventOfCode2022/internal"
	"golang.org/x/exp/slices"
)

const letterValue = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func main() {
	lines := internal.Reader()

	problem1(lines)
	problem2(lines)
}

func problem1(lines []string) {
	var score int
	for _, line := range lines {
		comp1 := line[:len(line)/2]
		comp2 := line[len(line)/2:]

		var bothFound []rune
		for _, i := range comp1 {
			if strings.ContainsRune(comp2, i) && !slices.Contains(bothFound, i) {
				index := strings.Index(letterValue, string(i)) + 1
				score += index
				bothFound = append(bothFound, i)
			}
		}
	}
	fmt.Printf("Problem 1 solution %d\n", score)
}

func problem2(lines []string) {
	var score int

	for i := 0; i < len(lines); i += 3 {
		r1 := lines[i]
		r2 := lines[i+1]
		r3 := lines[i+2]

		for _, i := range r1 {
			if strings.ContainsRune(r2, i) && strings.ContainsRune(r3, i) {
				index := strings.Index(letterValue, string(i)) + 1
				score += index
				break
			}
		}
	}
	fmt.Printf("Problem 2 solution %d\n", score)

}
