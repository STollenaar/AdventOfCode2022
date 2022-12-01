package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

type Inventory struct {
	calories []int
	total    int
}

var (
	inventories []Inventory
)

func main() {

	f, err := os.Open("input.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	current := Inventory{}
	for scanner.Scan() {
		line := scanner.Text()

		if err != nil {
			log.Fatal(err)
		}
		if line == "" {
			inventories = append(inventories, current)
			current = Inventory{}
		} else {
			amount, err := strconv.ParseInt(line, 10, 64)
			if err != nil {
				log.Fatal(err)
			}
			current.calories = append(current.calories, int(amount))
			current.total += int(amount)
		}
	}
	inventories = append(inventories, current)
	maxI := maxInventory()

	fmt.Printf("Maximum amount of calories carried by 1 elf (problem 1): %d\n", maxI.total)
	top3 := top3Inventories()

	var top3Total int
	for _, inv := range top3 {
		top3Total += inv.total
	}
	fmt.Printf("Maximum amount of calories carried by top 3 elves (problem 2): %d\n", top3Total)
}

func maxInventory() (max Inventory) {
	max = inventories[0]
	for _, inv := range inventories {
		if inv.total > max.total {
			max = inv
		}
	}
	return max
}

func top3Inventories() (top []Inventory) {
	for _, inv := range inventories {
		if len(top) < 3 {
			top = append(top, inv)
			sort.Slice(top, func(i, j int) bool {
				return top[i].total < top[j].total
			})
		} else {
			for i, t := range top {
				if inv.total > t.total {
					top[i] = inv
					sort.Slice(top, func(i, j int) bool {
						return top[i].total < top[j].total
					})
					break
				}
			}
		}
	}
	return top
}
