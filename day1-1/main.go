package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	increments := 0
	if scanner.Scan() {
		val, _ := strconv.Atoi(scanner.Text())

		for scanner.Scan() {
			next, _ := strconv.Atoi(scanner.Text())

			if next > val {
				increments++
			}

			val = next
		}

	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}

	fmt.Println(increments)
}
