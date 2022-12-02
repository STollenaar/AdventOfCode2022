package internal

import (
	"bufio"
	"log"
	"os"
)

func Reader() (lines []string) {
	f, err := os.Open("input.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		line := scanner.Text()

		if err != nil {
			log.Fatal(err)
		}
		lines = append(lines, line)
	}
	return lines
}
