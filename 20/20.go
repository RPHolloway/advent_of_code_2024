package main

import (
	"advent_of_code/utils/grid"
	"fmt"
	"math"
	"os"
	"time"
)

type Node struct {
	Location grid.Point
	Steps    int
}

type Cheat struct {
	Start grid.Point
	End   grid.Point
	Delta int
}

const FILE_NAME = "input.txt"
const CHEAT_LENGTH = 20
const CHEAT_THRESHOLD = 100

var CheatDirections = []grid.Point{
	{X: 1, Y: 1},   // down right
	{X: -1, Y: 1},  // down left
	{X: 1, Y: -1},  // up right
	{X: -1, Y: -1}, // up left
}

var Track [][]rune
var Visited [][]int
var Path []grid.Point
var Shortcuts map[grid.Point]struct{}

func timeTrack(start time.Time) {
	elapsed := time.Since(start)
	fmt.Println(elapsed)
}

func findCheatsInDirection(start grid.Point, direction grid.Point) map[Cheat]struct{} {
	cheats := make(map[Cheat]struct{})
	startValue := grid.SafeGet(Visited, start)

	for y := 0; y <= CHEAT_LENGTH; y++ {
		for x := 0; x <= CHEAT_LENGTH-y; x++ {
			next := start.Add(grid.Point{X: direction.X * x, Y: direction.Y * y})
			endValue := grid.SafeGet(Visited, next)
			steps := x + y
			if (endValue - steps) > startValue {
				cheats[Cheat{
					Start: start,
					End:   next,
					Delta: int(math.Abs(float64(startValue-endValue))) - steps,
				}] = struct{}{}
			}
		}
	}

	return cheats
}

func cheat(start grid.Point) map[Cheat]struct{} {
	cheats := make(map[Cheat]struct{})

	for _, dir := range CheatDirections {
		for k := range findCheatsInDirection(start, dir) {
			cheats[k] = struct{}{}
		}
	}

	return cheats
}

func checkPaths(p grid.Point) []grid.Point {
	var unvisited []grid.Point

	for i, dir := range grid.Directions {
		if grid.CheckDirection(Track, p, i) == '.' || grid.CheckDirection(Track, p, i) == 'E' {
			unvisited = append(unvisited, p.Add(dir))
		}
	}

	return unvisited
}

func visitPaths(start grid.Point) {
	stack := []Node{{Location: start, Steps: 0}}

	for len(stack) > 0 {
		n := stack[0]
		stack = stack[1:]

		Path = append(Path, n.Location)
		unvisited := checkPaths(n.Location)
		for _, c := range unvisited {
			var next Node
			next.Location = c
			next.Steps = grid.SafeGet(Visited, n.Location) + 1

			// If the path is shorter add it to the visited maze
			steps := grid.SafeGet(Visited, next.Location)
			if steps == 0 || next.Steps < steps {
				stack = append(stack, next)
				grid.Set(Visited, next.Location, next.Steps)
			}
		}
	}
}

func run() {
	var start, end grid.Point
	for y, row := range Track {
		for x, c := range row {
			if c == 'S' {
				start = grid.Point{X: x, Y: y}
				grid.Set(Visited, start, 1)
			} else if c == 'E' {
				end = grid.Point{X: x, Y: y}
			}
		}
	}

	visitPaths(start)
	baseline := grid.SafeGet(Visited, end)
	fmt.Printf("Baseline: %d\n", baseline-1)
	grid.OutputInt(Visited)

	var cheats []Cheat
	for _, p := range Path {
		for k := range cheat(p) {
			cheats = append(cheats, k)
		}
	}

	var uniqueCheats = make(map[int]int)

	total := 0
	for _, cheat := range cheats {
		if cheat.Delta >= CHEAT_THRESHOLD {
			uniqueCheats[cheat.Delta]++
			total++
		}
	}

	//fmt.Println(uniqueCheats)
	fmt.Printf("Total: %d\n", total)
}

func main() {
	// Read input
	data, _ := os.ReadFile(FILE_NAME)
	input := string(data)

	// Parse input
	Track = grid.Parse(input, "\r\n")
	Visited = grid.Create[int](len(Track[0]), len(Track))
	Shortcuts = make(map[grid.Point]struct{})

	defer timeTrack(time.Now())

	run()
}
