package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	if err := solve(); err != nil {
		log.Fatal(err)
	}
}

func solve() error {
	file, err := os.Open("input.txt")
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		//line := scanner.Text()
		// TODO
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	fmt.Println("done")

	return nil
}
