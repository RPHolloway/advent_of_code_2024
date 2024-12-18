package main

import (
	"advent_of_code/utils/grid"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type Node struct {
	Location grid.Point
	Steps    int
}

//const FILE_NAME = "example.txt"
//const REGION_SIZE = 7
//const BYTES = 12

const FILE_NAME = "input.txt"
const REGION_SIZE = 71
const BYTES = 1024

var Walls []grid.Point
var Region [][]rune
var VisitedRegion [][]int

func timeTrack(start time.Time) {
	elapsed := time.Since(start)
	fmt.Println(elapsed)
}

func checkPaths(p grid.Point) []grid.Point {
	var unvisited []grid.Point

	for i, dir := range grid.Directions {
		if grid.CheckDirection(Region, p, i) == '.' {
			unvisited = append(unvisited, p.Add(dir))
		}
	}

	return unvisited
}

func visitPaths() {
	stack := []Node{{Location: grid.Point{X: 0, Y: 0}, Steps: 0}}

	for len(stack) > 0 {
		n := stack[0]
		stack = stack[1:]

		unvisted := checkPaths(n.Location)
		for _, c := range unvisted {
			var next Node
			next.Location = c
			next.Steps = grid.SafeGet(VisitedRegion, n.Location) + 1

			// If the path is shorter add it to the visited maze
			steps := grid.SafeGet(VisitedRegion, next.Location)
			if steps == 0 || next.Steps < steps {
				stack = append(stack, next)
				grid.Set(VisitedRegion, next.Location, next.Steps)
			}
		}
	}
}

func run() {
	visitPaths()
	grid.OutputInt(VisitedRegion)

	steps := grid.SafeGet(VisitedRegion, grid.Point{X: REGION_SIZE - 1, Y: REGION_SIZE - 1})
	fmt.Println(steps)

	for i := BYTES; i < len(Walls); i++ {
		VisitedRegion = grid.Create[int](REGION_SIZE, REGION_SIZE)
		wall := Walls[i]
		Region[wall.Y][wall.X] = '#'

		visitPaths()
		steps := grid.SafeGet(VisitedRegion, grid.Point{X: REGION_SIZE - 1, Y: REGION_SIZE - 1})
		if steps == 0 {
			fmt.Println(Walls[i])
			break
		}
	}
}

func main() {
	// Read input
	data, _ := os.ReadFile(FILE_NAME)
	input := string(data)

	// Parse input
	lines := strings.Split(input, "\r\n")
	for _, line := range lines {
		p := strings.Split(line, ",")
		x, _ := strconv.Atoi(p[0])
		y, _ := strconv.Atoi(p[1])

		Walls = append(Walls, grid.Point{X: x, Y: y})
	}

	// Populate the region
	VisitedRegion = grid.Create[int](REGION_SIZE, REGION_SIZE)
	Region = grid.Create[rune](REGION_SIZE, REGION_SIZE)
	grid.Fill(Region, '.')

	for i := 0; i < BYTES; i++ {
		wall := Walls[i]
		Region[wall.Y][wall.X] = '#'
	}

	grid.Output(Region)

	defer timeTrack(time.Now())

	run()
}
