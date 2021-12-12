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
		total, _ := strconv.Atoi(scanner.Text())
		first := total

		if scanner.Scan() {
			second, _ := strconv.Atoi(scanner.Text())
			total += second

			if scanner.Scan() {
				third, _ := strconv.Atoi(scanner.Text())
				total += third

				for scanner.Scan() {
					next, _ := strconv.Atoi(scanner.Text())
					new := total - first + next

					if new > total {
						increments++
					}

					total = new
					first = second
					second = third
					third = next
				}
			}
		}

	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}

	fmt.Println(increments)
}
