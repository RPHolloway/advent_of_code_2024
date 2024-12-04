package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"time"
)

var enable_re = regexp.MustCompile(`do\(\)|don't\(\)`)
var instruction_re = regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)

func timeTrack(start time.Time) {
	elapsed := time.Since(start)
	fmt.Println(elapsed)
}

func test(input string, iter int) {
	var total int = 0

	// Split file on do() and don't()
	parts := enable_re.Split(input, -1)
	enable := enable_re.FindAllString(input, -1)

	// Parse instructions
	for i, instructions := range parts {
		if i == 0 || enable[i-1] == "do()" {
			matches := instruction_re.FindAllString(instructions, -1)

			// Verify reports
			for _, match := range matches {
				values := instruction_re.FindStringSubmatch(match)
				v1, _ := strconv.Atoi(values[1])
				v2, _ := strconv.Atoi(values[2])

				total += v1 * v2
			}
		}
	}
}

func main() {
	// Read input
	//data, _ := os.ReadFile("example.txt")
	data, _ := os.ReadFile("input.txt")
	input := string(data)

	defer timeTrack(time.Now())

	for i := range 100000 {
		test(input, i)
	}
}
