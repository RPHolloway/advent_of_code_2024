package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

var WORD []rune = []rune("XMAS")

func timeTrack(start time.Time) {
	elapsed := time.Since(start)
	fmt.Println(elapsed)
}

func safeAccess(arr [][]rune, x int, y int) rune {
	defer func() {
		if r := recover(); r != nil {
		}
	}()

	return arr[y][x]
}

func checkDirection(rows [][]rune, word_idx int, x int, y int, dir_x int, dir_y int) bool {
	for word_idx < len(WORD) {
		x += dir_x
		y += dir_y

		c := safeAccess(rows, x, y)
		if c == WORD[word_idx] {
			word_idx++
		} else {
			return false
		}
	}

	return true
}

func test(rows [][]rune) int {
	total := 0

	for y, row := range rows {
		for x, c := range row {
			word_idx := 0
			if c == WORD[word_idx] {
				word_idx++

				for dir_y := -1; dir_y <= 1; dir_y++ {
					for dir_x := -1; dir_x <= 1; dir_x++ {
						if checkDirection(rows, word_idx, x, y, dir_x, dir_y) {
							total++
						}
					}
				}
			}
		}
	}

	return total
}

func main() {
	// Read input
	//data, _ := os.ReadFile("example.txt")
	data, _ := os.ReadFile("input.txt")
	input := string(data)

	// Parse input
	var rows [][]rune
	lines := strings.Split(input, "\n")
	for _, line := range lines {
		rows = append(rows, []rune(line))
	}

	defer timeTrack(time.Now())

	total := test(rows)
	fmt.Printf("Total: %d\r\n", total)
}
