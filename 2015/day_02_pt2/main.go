package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
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

	sum := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var a, b, c int
		_, err := fmt.Sscanf(scanner.Text(), "%dx%dx%d", &a, &b, &c)
		if err != nil {
			return err
		}

		nums := []int{a, b, c}
		sort.Ints(nums)

		sum += nums[0] + nums[0] + nums[1] + nums[1] + nums[0]*nums[1]*nums[2]
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	fmt.Println("ans:", sum)

	return nil
}
