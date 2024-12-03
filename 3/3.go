package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	var total int = 0

	// Read input
	//data, _ := os.ReadFile("example.txt")
	data, _ := os.ReadFile("input.txt")
	input := string(data)

	// Split file on do() and don't()
	pattern := `do\(\)|don't\(\)`
	re := regexp.MustCompile(pattern)
	parts := re.Split(input, -1)
	enable := re.FindAllString(input, -1)

	// Parse instructions
	pattern = `mul\((\d{1,3}),(\d{1,3})\)`
	re = regexp.MustCompile(pattern)

	for i, instructions := range parts {
		if i == 0 || enable[i-1] == "do()" {
			matches := re.FindAllString(instructions, -1)

			fmt.Println(matches)

			// Verify reports
			for _, match := range matches {
				values := re.FindStringSubmatch(match)
				v1, _ := strconv.Atoi(values[1])
				v2, _ := strconv.Atoi(values[2])

				total += v1 * v2
			}
		}
	}

	fmt.Println(total)
}
