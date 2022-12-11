package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/STollenaar/AdventOfCode2022/internal"
	"golang.org/x/exp/slices"
)

type Monke struct {
	inspectedAmount int
	monkeyThrow     map[bool]int

	modulo int

	operation string
	modifier  string
	items     []int
}

var (
	problem1Monkeys map[int]*Monke
	problem2Monkeys map[int]*Monke
	commonDiv       int
)

func init() {
	problem1Monkeys = make(map[int]*Monke)
	problem2Monkeys = make(map[int]*Monke)
	commonDiv = 1
}

func main() {
	lines := internal.Reader()

	monkeys := make(map[int]*Monke)
	current := &Monke{
		monkeyThrow: make(map[bool]int),
	}
	for _, line := range lines {
		if line == "" {
			monkeys[len(monkeys)] = current
			current = &Monke{
				monkeyThrow: make(map[bool]int),
			}
		} else {
			line = strings.TrimSpace(line)
			line = strings.ReplaceAll(line, ":", "")
			line = strings.ReplaceAll(line, ",", "")
			args := strings.Split(line, " ")
			switch args[0] {
			case "Starting":
				for i := 2; i < len(args); i++ {
					v, _ := strconv.Atoi(args[i])
					current.addItem(v)
				}
			case "Test":
				v, _ := strconv.Atoi(args[3])
				current.modulo = v
			case "If":
				v, _ := strconv.Atoi(args[5])
				current.monkeyThrow[args[1] == "true"] = v
			case "Operation":
				current.operation = args[4]
				current.modifier = args[5]
			}

		}
	}
	monkeys[len(monkeys)] = current
	problem1Monkeys = monkeys

	for k, v := range monkeys {
		problem2Monkeys[k] = &Monke{
			items:       slices.Clone(v.items),
			monkeyThrow: v.monkeyThrow,
			modulo:      v.modulo,
			operation:   v.operation,
			modifier:    v.modifier,
		}
		commonDiv = commonDiv * v.modulo
	}

	runSimulation(problem1Monkeys, 20, true)

	problem1MonkySlice := toSlice(problem1Monkeys)
	sort.Slice(problem1MonkySlice, func(i, j int) bool {
		return problem1MonkySlice[i].inspectedAmount > problem1MonkySlice[j].inspectedAmount
	})

	fmt.Printf("Problem 1: %d\n", problem1MonkySlice[0].inspectedAmount*problem1MonkySlice[1].inspectedAmount)

	runSimulation(problem2Monkeys, 10000, false)

	problem2MonkySlice := toSlice(problem2Monkeys)
	sort.Slice(problem2MonkySlice, func(i, j int) bool {
		return problem2MonkySlice[i].inspectedAmount > problem2MonkySlice[j].inspectedAmount
	})

	fmt.Printf("Problem 2: %d\n", problem2MonkySlice[0].inspectedAmount*problem2MonkySlice[1].inspectedAmount)
}

func toSlice(monkeys map[int]*Monke) (s []*Monke) {
	for _, v := range monkeys {
		s = append(s, v)
	}
	return s
}

// Do stress operation
func (m *Monke) doOperation(n int) int {
	v, e := strconv.Atoi(m.modifier)
	if e != nil {
		v = n
	}

	if m.operation == "*" {
		return n * v
	} else {
		return n + v
	}
}

// Give Monke item
func (m *Monke) addItem(i int) {
	m.items = append(m.items, i)
}

// Remove Item from Monke
func (m *Monke) popItem() int {
	currentItem := m.items[0]
	m.items = m.items[1:]
	return currentItem
}

func runSimulation(monkeys map[int]*Monke, rounds int, reduce bool) {
	for ; rounds > 0; rounds-- {
		for i := 0; i < len(monkeys); i++ {
			c := monkeys[i]
			for len(c.items) > 0 {
				c.inspectedAmount++
				currentItem := c.popItem()
				currentItem = c.doOperation(currentItem)

				if reduce {
					currentItem = currentItem / 3
				} else {
					currentItem = currentItem % commonDiv
				}
				monkeys[c.monkeyThrow[currentItem%c.modulo == 0]].addItem(currentItem)
			}
		}
	}
}
