package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/STollenaar/AdventOfCode2022/internal"
)

type file struct {
	size int
}

type dir struct {
	parent    *dir
	totalSize int
	dirs      map[string]*dir
	files     map[string]*file
}

var (
	current *dir
)

func init() {
	current = &dir{
		parent: nil,
		dirs:   make(map[string]*dir),
		files:  make(map[string]*file),
	}
}

func main() {
	lines := internal.Reader()

	for _, line := range lines {
		fields := strings.Split(line, " ")

		switch fields[0] {
		case "$":
			if fields[1] == "cd" {
				switch fields[2] {
				case "..":
					current = current.parent
				case "/":
					for current.parent != nil {
						current = current.parent
					}
				default:
					current = current.dirs[fields[2]]
				}
			}
		case "dir":
			current.dirs[fields[1]] = &dir{
				parent: current,
				dirs:   make(map[string]*dir),
				files:  make(map[string]*file),
			}
		default:
			size, _ := strconv.Atoi(fields[0])
			current.files[fields[1]] = &file{
				size: size,
			}
			tmp := current
			for tmp != nil {
				tmp.totalSize += size
				tmp = tmp.parent
			}
		}
	}
	for current.parent != nil {
		current = current.parent
	}

	fmt.Printf("Problem 1: %d\n", problem1Traverse(current))
	amount := int(math.Abs(float64(30000000 - (70000000 - current.totalSize))))
	p2Dirs := problem2Traverse(current, amount)
	min := findSmallest(p2Dirs)
	fmt.Printf("Problem 2: %d\n", min)
}

func problem1Traverse(dir *dir) (total int) {
	if dir.totalSize < 100000 {
		total += dir.totalSize
	}
	for _, d := range dir.dirs {
		total += problem1Traverse(d)
	}
	return total
}

func problem2Traverse(dir *dir, amount int) (dirs []*dir) {
	if dir.totalSize >= amount {
		dirs = append(dirs, dir)
	}
	for _, d := range dir.dirs {
		dirs = append(dirs, problem2Traverse(d, amount)...)
	}
	return dirs
}

func findSmallest(dirs []*dir) (min int) {
	min = dirs[0].totalSize
	for _, d := range dirs {
		if d.totalSize < min {
			min = d.totalSize
		}
	}
	return min
}
