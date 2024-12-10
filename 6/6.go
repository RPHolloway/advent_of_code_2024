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

func Rotate(current_dir int) int {
	return (current_dir + 1) % len(directions)
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

func test_loop(rows [][]rune, origin Point, gaurd Point, current_dir int) Point {
	rows_copy := make([][]rune, len(rows))
	copy(rows_copy, rows)
	for i := range rows {
		rows_copy[i] = append([]rune(nil), rows[i]...)
	}

	block := gaurd.Add(directions[current_dir])
	rows_copy[block.Y][block.X] = '#'
	if block == origin {
		// Can't put the object at the origin
		return Point{}
	}

	gaurd = origin
	current_dir = 0
	next := gaurd.Add(directions[current_dir])

	for {
		if rows_copy[gaurd.Y][gaurd.X] == '.' {
			rows_copy[gaurd.Y][gaurd.X] = rune(current_dir + 10)
		}
		next_obj := safeAccess(rows_copy, next)

		if next_obj == '#' {
			// hit an object
			current_dir = Rotate(current_dir)
		} else if next_obj == 0 {
			// left the room
			break
		} else if next_obj == rune(current_dir+10) {
			return block
		} else {
			// move
			gaurd = next
		}

		next = gaurd.Add(directions[current_dir])
	}

	return Point{}
}

func test(rows [][]rune) int {
	steps := 1
	loop_objs := make(map[Point]struct{})

	// Find the gaurd
	var origin Point
	for y, row := range rows {
		for x, c := range row {
			if c == '^' {
				origin = Point{x, y}
			}
		}
	}

	gaurd := origin
	current_dir := 0
	for {
		if rows[gaurd.Y][gaurd.X] == rune('.') {
			//rows[gaurd.Y][gaurd.X] = rune(current_dir + 10)
			steps++
		}

		next := gaurd.Add(directions[current_dir])
		next_obj := safeAccess(rows, next)

		if next_obj == '#' {
			// hit an object
			current_dir = Rotate(current_dir)
		} else if next_obj == 0 {
			// left the room
			break
		} else {
			loop_objs[test_loop(rows, origin, gaurd, current_dir)] = struct{}{}

			// move
			gaurd = next
		}
	}

	fmt.Printf("Steps: %d\n", steps)
	fmt.Printf("Loops: %d\n", len(loop_objs)-1)

	return steps
}

func main() {
	// Read input
	//data, _ := os.ReadFile("corner.txt")
	//data, _ := os.ReadFile("example.txt")
	data, _ := os.ReadFile("input.txt")
	input := string(data)

	// Parse sections
	var rows [][]rune
	for _, line := range strings.Split(input, "\r\n") {
		rows = append(rows, []rune(line))
	}

	defer timeTrack(time.Now())

	total := test(rows)
	fmt.Printf("Total: %d\r\n", total)
}
