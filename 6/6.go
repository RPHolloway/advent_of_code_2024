package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

type Point struct {
	X int
	Y int
}

var direction_idx int = 0
var directions = []Point{
	{0, -1}, // up
	{1, 0},  // right
	{0, 1},  // down
	{-1, 0}, // left
}

func (p1 Point) Add(p2 Point) Point {
	return Point{
		X: p1.X + p2.X,
		Y: p1.Y + p2.Y,
	}
}

func Rotate() Point {
	direction_idx++
	return directions[direction_idx%len(directions)]
}

func timeTrack(start time.Time) {
	elapsed := time.Since(start)
	fmt.Println(elapsed)
}

func safeAccess(arr [][]rune, p Point) rune {
	defer func() {
		if r := recover(); r != nil {
		}
	}()

	return arr[p.Y][p.X]
}

func test(rows [][]rune) int {
	total := 0

	// Find the gaurd
	var gaurd Point
	for y, row := range rows {
		for x, c := range row {
			if c == '^' {
				gaurd = Point{x, y}
			}
		}
	}

	direction := directions[0]
	for {
		if rows[gaurd.Y][gaurd.X] != 'X' {
			total++
			rows[gaurd.Y][gaurd.X] = 'X'
		}

		next := gaurd.Add(direction)
		if safeAccess(rows, next) == '#' {
			// hit an object
			direction = Rotate()
		} else if safeAccess(rows, next) == 0 {
			// left the room
			break
		} else {
			// move
			gaurd = next
		}
	}

	return total
}

func main() {
	// Read input
	//data, _ := os.ReadFile("example.txt")
	data, _ := os.ReadFile("input.txt")
	input := string(data)

	// Parse sections
	var rows [][]rune
	for _, line := range strings.Split(input, "\n") {
		rows = append(rows, []rune(line))
	}

	defer timeTrack(time.Now())

	total := test(rows)
	fmt.Printf("Total: %d\r\n", total)
}
