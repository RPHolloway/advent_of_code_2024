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

func checkCorners(rows [][]rune, x int, y int) bool {
	top_left := '.'
	top_right := '.'
	bottom_left := '.'
	bottom_right := '.'

	// top left
	next_x := x - 1
	next_y := y - 1
	c := safeAccess(rows, next_x, next_y)
	if c == 'M' || c == 'S' {
		top_left = c
	} else {
		return false
	}

	// top right
	next_x = x + 1
	next_y = y - 1
	c = safeAccess(rows, next_x, next_y)
	if c == 'M' || c == 'S' {
		top_right = c
	} else {
		return false
	}

	// bottom left
	next_x = x - 1
	next_y = y + 1
	c = safeAccess(rows, next_x, next_y)
	if c == 'M' || c == 'S' {
		bottom_left = c
	} else {
		return false
	}

	// bottom right
	next_x = x + 1
	next_y = y + 1
	c = safeAccess(rows, next_x, next_y)
	if c == 'M' || c == 'S' {
		bottom_right = c
	} else {
		return false
	}

	if top_left == bottom_right || top_right == bottom_left {
		return false
	}

	if (top_left != bottom_left && top_right != bottom_right) && (top_left != top_right && bottom_left != bottom_right) {
		return false
	}

	return true
}

func test(rows [][]rune) int {
	total := 0

	for y, row := range rows {
		for x, c := range row {
			if c == 'A' {
				if checkCorners(rows, x, y) {
					total++
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
