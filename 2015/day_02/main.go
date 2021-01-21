package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	inputs, err := getInputs()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("total (no goroutines): %v\n", noGoroutines(inputs))
	fmt.Printf("total (with goroutines): %v\n", withGoroutines(inputs))
}

func surfaceArea(nums []int) int {
	l, w, h := nums[0], nums[1], nums[2]
	return 2 * (l*w + w*h + h*l)
}

func minSmallestSide(nums []int) int {
	l, w, h := nums[0], nums[1], nums[2]
	sides := []int{l * w, w * h, l * h}
	return min(sides)
}

func min(nums []int) int {
	var m int

	for i, e := range nums {
		if i == 0 || e < m {
			m = e
		}
	}

	return m
}

func noGoroutines(inputs [][]int) int {
	sum := 0

	for _, nums := range inputs {
		sum += surfaceArea(nums) + minSmallestSide(nums)
	}

	return sum
}

func withGoroutines(inputs [][]int) int {
	sum := 0

	results := make(chan int, 1000)
	defer close(results)

	for _, nums := range inputs {
		nums := nums // need to re-declare to properly scope
		go func() {
			results <- surfaceArea(nums) + minSmallestSide(nums)
		}()
	}

	count := 0
	for res := range results {
		sum += res

		count++
		if count == 1000 {
			break
		}
	}

	return sum
}

func getInputs() ([][]int, error) {
	file, err := os.Open("input.txt")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var inputs [][]int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		numsStr := strings.Split(line, "x")

		// convert to slice of ints
		var nums []int
		for _, v := range numsStr {
			n, _ := strconv.Atoi(v)
			nums = append(nums, n)
		}

		inputs = append(inputs, nums)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return inputs, nil
}
