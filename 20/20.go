package main

import (
	"advent_of_code/utils/grid"
	"fmt"
	"math"
	"os"
	"time"
)

type Cheat struct {
	Start grid.Point
	End   grid.Point
	Delta int
}

const FILE_NAME = "example.txt"
const CHEAT_LENGTH = 2

var Track [][]rune
var Visited [][]int
var Path []grid.Point
var Shortcuts map[grid.Point]struct{}

func timeTrack(start time.Time) {
	elapsed := time.Since(start)
	fmt.Println(elapsed)
}

func checkShortcut(p grid.Point) int {
	var delta int

	var neighbors []int
	for _, dir := range grid.Directions {
		v := grid.SafeGet(Visited, p.Add(dir))
		neighbors = append(neighbors, v)
	}

	if neighbors[grid.DIR_UP] != 0 && neighbors[grid.DIR_DOWN] != 0 {
		delta = neighbors[grid.DIR_UP] - neighbors[grid.DIR_DOWN]
	} else if neighbors[grid.DIR_LEFT] != 0 && neighbors[grid.DIR_RIGHT] != 0 {
		delta = neighbors[grid.DIR_LEFT] - neighbors[grid.DIR_RIGHT]
	}

	return int(math.Abs(float64(delta)))
}

func cheat(start grid.Point) []Cheat {
	var cheats []Cheat
	var next grid.Point

	startValue := grid.SafeGet(Visited, start)
	for y := range CHEAT_LENGTH {
		next = start
		for x := range CHEAT_LENGTH - y {
			next = next.Add(grid.Point{X: x, Y: y})
			endValue := grid.SafeGet(Visited, next)
			if endValue > startValue {
				cheats = append(cheats, Cheat{
					Start: start,
					End:   next,
					Delta: startValue - endValue,
				})
			}
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

	for _, p := range Path {

	}

	cheats := make(map[int]int)
	for y, row := range Track {
		for x, c := range row {
			if c == '#' {
				delta := checkShortcut(grid.Point{X: x, Y: y}) - 2
				if delta > 0 {
					cheats[delta]++
				}
			}
		}
	}

	total := 0
	for score, count := range cheats {
		if score >= 100 {
			total += count
		}
	}

	//fmt.Println(cheats)
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
