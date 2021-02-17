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
	robo := p{x: 0, y: 0}

	set := make(map[p]struct{})
	roboSet := make(map[p]struct{})
	set[curr] = struct{}{}
	roboSet[robo] = struct{}{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		// iterate over chars of string
		for i, r := range line {
			if i%2 == 0 {
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
			} else {
				if r == '^' {
					robo.y++
					roboSet[robo] = struct{}{}
				} else if r == 'v' {
					robo.y--
					roboSet[robo] = struct{}{}
				} else if r == '>' {
					robo.x++
					roboSet[robo] = struct{}{}
				} else if r == '<' {
					robo.x--
					roboSet[robo] = struct{}{}
				}
			}
		}
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	res := make(map[p]struct{})
	for k, v := range set {
		res[k] = v
	}

	for k, v := range roboSet {
		res[k] = v
	}

	fmt.Printf("res: %v\n", len(res))
	fmt.Println("done")

	return nil
}
