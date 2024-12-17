package main

import (
	"advent_of_code/utils/grid"
	"fmt"
	"os"
	"strings"
	"time"
)

var Warehouse [][]rune
var BigWarehouse [][]rune
var Instructions []rune

var Robot grid.Point

func timeTrack(start time.Time) {
	elapsed := time.Since(start)
	fmt.Println(elapsed)
}

func expand() {
	for _, row := range Warehouse {
		var bigRow []rune
		for _, c := range row {
			switch c {
			case 'O':
				bigRow = append(bigRow, '[')
				bigRow = append(bigRow, ']')
			case '@':
				bigRow = append(bigRow, '@')
				bigRow = append(bigRow, '.')
			default:
				bigRow = append(bigRow, c)
				bigRow = append(bigRow, c)
			}
		}
		BigWarehouse = append(BigWarehouse, bigRow)
	}
}

func moveBigLeft() {
	dir := grid.DIR_LEFT
	boxes := 0

	next := Robot
	for {
		if grid.CheckDirection(BigWarehouse, next, dir) == '.' {
			// Move robot
			grid.Set(BigWarehouse, Robot, '.')
			Robot = Robot.Add(grid.Directions[dir])
			grid.Set(BigWarehouse, Robot, '@')

			// Move boxes
			Box := Robot
			for i := 0; i < boxes; i += 2 {
				Box = Box.Add(grid.Directions[dir])
				grid.Set(BigWarehouse, Box, ']')
				Box = Box.Add(grid.Directions[dir])
				grid.Set(BigWarehouse, Box, '[')
			}
			break
		} else if grid.CheckDirection(BigWarehouse, next, dir) == '#' {
			return
		} else {
			boxes++
		}

		next = next.Add(grid.Directions[dir])
	}
}

func moveBigRight() {
	dir := grid.DIR_RIGHT
	boxes := 0

	next := Robot
	for {
		if grid.CheckDirection(BigWarehouse, next, dir) == '.' {
			// Move robot
			grid.Set(BigWarehouse, Robot, '.')
			Robot = Robot.Add(grid.Directions[dir])
			grid.Set(BigWarehouse, Robot, '@')

			// Move boxes
			Box := Robot
			for i := 0; i < boxes; i += 2 {
				Box = Box.Add(grid.Directions[dir])
				grid.Set(BigWarehouse, Box, '[')
				Box = Box.Add(grid.Directions[dir])
				grid.Set(BigWarehouse, Box, ']')
			}
			break
		} else if grid.CheckDirection(BigWarehouse, next, dir) == '#' {
			return
		} else {
			boxes++
		}

		next = next.Add(grid.Directions[dir])
	}
}

func checkBigVerticalMove(next grid.Point, dir int) bool {
	canMove := false

	if grid.CheckDirection(BigWarehouse, next, dir) == '.' {
		canMove = true
	} else if grid.CheckDirection(BigWarehouse, next, dir) == '#' {
		canMove = false
	} else if grid.CheckDirection(BigWarehouse, next, dir) == '[' {
		if checkBigVerticalMove(next.Add(grid.Directions[dir]), dir) {
			next = next.Add(grid.Directions[grid.DIR_RIGHT])
			canMove = checkBigVerticalMove(next.Add(grid.Directions[dir]), dir)
		}
	} else if grid.CheckDirection(BigWarehouse, next, dir) == ']' {
		if checkBigVerticalMove(next.Add(grid.Directions[dir]), dir) {
			next = next.Add(grid.Directions[grid.DIR_LEFT])
			canMove = checkBigVerticalMove(next.Add(grid.Directions[dir]), dir)
		}
	}

	return canMove
}

func bigVerticalMove(next grid.Point, dir int) {
	if grid.CheckDirection(BigWarehouse, next, dir) == '.' {
		v := grid.SafeGet(BigWarehouse, next)
		grid.Set(BigWarehouse, next.Add(grid.Directions[dir]), v)
		grid.Set(BigWarehouse, next, '.')
	} else if grid.CheckDirection(BigWarehouse, next, dir) == '[' {
		bigVerticalMove(next.Add(grid.Directions[dir]), dir)
		v := grid.SafeGet(BigWarehouse, next)
		grid.Set(BigWarehouse, next.Add(grid.Directions[dir]), v)
		grid.Set(BigWarehouse, next, '.')

		next = next.Add(grid.Directions[grid.DIR_RIGHT])
		bigVerticalMove(next.Add(grid.Directions[dir]), dir)
		v = grid.SafeGet(BigWarehouse, next)
	} else if grid.CheckDirection(BigWarehouse, next, dir) == ']' {
		bigVerticalMove(next.Add(grid.Directions[dir]), dir)
		v := grid.SafeGet(BigWarehouse, next)
		grid.Set(BigWarehouse, next.Add(grid.Directions[dir]), v)
		grid.Set(BigWarehouse, next, '.')

		next = next.Add(grid.Directions[grid.DIR_LEFT])
		bigVerticalMove(next.Add(grid.Directions[dir]), dir)
		v = grid.SafeGet(BigWarehouse, next)
	}
}

func moveBigDown() {
	dir := grid.DIR_DOWN

	next := Robot
	if grid.CheckDirection(BigWarehouse, next, dir) == '.' {
		// Move robot
		grid.Set(BigWarehouse, Robot, '.')
		Robot = Robot.Add(grid.Directions[dir])
		grid.Set(BigWarehouse, Robot, '@')
	} else if grid.CheckDirection(BigWarehouse, next, dir) == '#' {
		return
	} else {
		if checkBigVerticalMove(next, dir) {
			bigVerticalMove(next, dir)

			// Move robot
			grid.Set(BigWarehouse, Robot, '.')
			Robot = Robot.Add(grid.Directions[dir])
			grid.Set(BigWarehouse, Robot, '@')
		}
	}
}

func moveBigUp() {
	dir := grid.DIR_UP

	next := Robot
	if grid.CheckDirection(BigWarehouse, next, dir) == '.' {
		// Move robot
		grid.Set(BigWarehouse, Robot, '.')
		Robot = Robot.Add(grid.Directions[dir])
		grid.Set(BigWarehouse, Robot, '@')
	} else if grid.CheckDirection(BigWarehouse, next, dir) == '#' {
		return
	} else {
		if checkBigVerticalMove(next, dir) {
			bigVerticalMove(next, dir)

			// Move robot
			grid.Set(BigWarehouse, Robot, '.')
			Robot = Robot.Add(grid.Directions[dir])
			grid.Set(BigWarehouse, Robot, '@')
		}
	}
}

func runBig() {
	expand()
	grid.Output(BigWarehouse)

	// Find robot
	for y, row := range BigWarehouse {
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
			moveBigUp()
		case '>':
			moveBigRight()
		case 'v':
			moveBigDown()
		case '<':
			moveBigLeft()
		}

		//grid.Output(BigWarehouse)
	}

	grid.Output(BigWarehouse)

	// Calculate result
	total := 0
	for y, row := range BigWarehouse {
		for x, c := range row {
			if c == '[' {
				total += y*100 + x
			}
		}
	}
	fmt.Println(total)
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

	runBig()
	run()
}
