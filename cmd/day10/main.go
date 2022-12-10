package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/STollenaar/AdventOfCode2022/internal"
)

var (
	cycle    int
	register int
	signal   int

	screen [][]string
)

func init() {
	register = 1

	screen = make([][]string, 6)
	for i := range screen {
		screen[i] = make([]string, 41)
	}
}

func main() {
	lines := internal.Reader()

	for _, line := range lines {
		// fmt.Println(cycle, register, line)

		problem2CycleCheck()
		if line == "noop" {
			cycle++
			problem1CycleCheck()
		} else {
			cycle++
			problem1CycleCheck()
			v := strings.Split(line, " ")[1]
			value, _ := strconv.Atoi(v)
			problem2CycleCheck()
			cycle++
			problem1CycleCheck()
			register += value
		}
	}
	fmt.Printf("Problem 1: %d\n", signal)

	for _, r := range screen {
		for _, c := range r {
			fmt.Print(c)
		}
		fmt.Print("\n")
	}
}

func problem1CycleCheck() {
	if (cycle-20)%40 == 0 {
		signal += register * cycle
	}
}

func problem2CycleCheck() {
	rY := int(math.Floor(float64(cycle) / 40))
	rX := cycle - rY*40 + 1 // AN OFF BY 1 ERROR?? WHAT THE HELL
	fmt.Println(rY, rX, cycle, register)
	r := screen[rY]

	if register == rX || register+1 == rX || register+2 == rX {
		r[rX] = "#"
	} else {
		r[rX] = "."
	}
}
