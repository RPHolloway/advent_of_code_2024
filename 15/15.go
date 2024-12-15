package main

import (
	"advent_of_code/utils/grid"
	"fmt"
	"os"
	"strings"
	"time"
)

var Warehouse [][]rune
var Instructions []rune

var Robot grid.Point

func timeTrack(start time.Time) {
	elapsed := time.Since(start)
	fmt.Println(elapsed)
}

func move(dir int) {
	boxes := 0

	next := Robot
	for {
		if grid.CheckDirection(Warehouse, next, dir) == '.' {
			// Move robot
			grid.Set(Warehouse, Robot, '.')
			Robot = Robot.Add(grid.Directions[dir])
			grid.Set(Warehouse, Robot, '@')

			// Move boxes
			Box := Robot
			for i := 0; i < boxes; i++ {
				Box = Box.Add(grid.Directions[dir])
				grid.Set(Warehouse, Box, 'O')
			}

			//grid.Output(Warehouse)
			break
		} else if grid.CheckDirection(Warehouse, next, dir) == '#' {
			return
		} else {
			boxes++
		}

		next = next.Add(grid.Directions[dir])
	}
}

func run() {
	// Find robot
	for y, row := range Warehouse {
		for x, c := range row {
			if c == '@' {
				Robot = grid.Point{X: x, Y: y}
			}
		}
	}

	// Follow instructions
	for _, i := range Instructions {
		switch i {
		case '^':
			move(grid.DIR_UP)
		case '>':
			move(grid.DIR_RIGHT)
		case 'v':
			move(grid.DIR_DOWN)
		case '<':
			move(grid.DIR_LEFT)
		}
	}

	grid.Output(Warehouse)

	// Calculate result
	total := 0
	for y, row := range Warehouse {
		for x, c := range row {
			if c == 'O' {
				total += y*100 + x
			}
		}
	}

	fmt.Println(total)
}

func main() {
	// Read input
	//data, _ := os.ReadFile("example.txt")
	data, _ := os.ReadFile("input.txt")
	input := string(data)

	// Parse input
	sections := strings.Split(input, "\r\n\r\n")

	Warehouse = grid.Parse(sections[0], "\r\n")
	Instructions = []rune(strings.ReplaceAll(sections[1], "\r\n", ""))

	defer timeTrack(time.Now())

	run()
}
