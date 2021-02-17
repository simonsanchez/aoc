package main

import (
	"bufio"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"unicode"
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
		line := scanner.Text()

		count := 0
		password := make([]string, 8, 8)
		assigned := make([]bool, 8, 8)

		for i := 0; ; i++ {
			input := line + strconv.Itoa(i)
			hash := md5.Sum([]byte(input))

			// Need to convert [16]byte to []byte so we can pass
			// this into hex.EncodeToString below
			var temp []byte
			for _, v := range hash {
				temp = append(temp, v)
			}

			hs := hex.EncodeToString(temp)

			if !startsWithFiveZeroes(hs) {
				continue
			}

			if !unicode.IsDigit(rune(hs[5])) {
				continue
			}

			pos, _ := strconv.Atoi(string(hs[5]))
			if pos > 7 {
				continue
			}

			if assigned[pos] {
				continue
			}

			password[pos] = string(hs[6])
			assigned[pos] = true

			count++
			if count == 8 {
				break
			}
		}

		fmt.Printf("the password: %v\n", strings.Join(password, ""))

	}

	if err := scanner.Err(); err != nil {
		return err
	}

	fmt.Println("\n\t\t*** done ***")

	return nil
}

func startsWithFiveZeroes(s string) bool {
	for i := 0; i < 5; i++ {
		if s[i] != '0' {
			return false
		}
	}
	return true
}
