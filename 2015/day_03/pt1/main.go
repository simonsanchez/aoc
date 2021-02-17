package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	if err := solvePart(); err != nil {
		log.Fatal(err)
	}
}

func solvePart() error {
	file, err := os.Open("input.txt")
	if err != nil {
		return err
	}
	defer file.Close()

	type p struct{ x, y int }

	curr := p{x: 0, y: 0}

	set := make(map[p]struct{})
	set[curr] = struct{}{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		// iterate over chars of string
		for _, r := range line {
			if r == '^' {
				curr.y++
				set[curr] = struct{}{}
			} else if r == 'v' {
				curr.y--
				set[curr] = struct{}{}
			} else if r == '>' {
				curr.x++
				set[curr] = struct{}{}
			} else if r == '<' {
				curr.x--
				set[curr] = struct{}{}
			}
		}
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	fmt.Printf("total: %v\n", len(set))
	fmt.Println("done")

	return nil
}
