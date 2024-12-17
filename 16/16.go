package main

import (
	"advent_of_code/utils/grid"
	"fmt"
	"os"
	"time"
)

type Node struct {
	Reindeer grid.Cursor
	Score    int
}

var Maze [][]rune
var VisitedMaze [][]int

func timeTrack(start time.Time) {
	elapsed := time.Since(start)
	fmt.Println(elapsed)
}

func checkPaths(p grid.Point) []grid.Cursor {
	var unvisited []grid.Cursor

	for i, dir := range grid.Directions {
		if grid.CheckDirection(Maze, p, i) == '.' || grid.CheckDirection(Maze, p, i) == 'E' {
			unvisited = append(unvisited, grid.Cursor{Location: p.Add(dir), Direction: i})
		} else if grid.CheckDirection(Maze, p, i) == '#' {
			grid.Set(VisitedMaze, p.Add(dir), 77777)
		}
	}

	return unvisited
}

func run() {
	// Find the start
	var start grid.Cursor
	var end grid.Point
	for y, row := range Maze {
		for x, c := range row {
			if c == 'S' {
				start.Location = grid.Point{X: x, Y: y}
				start.Direction = grid.DIR_RIGHT
			} else if c == 'E' {
				end = grid.Point{X: x, Y: y}
			}
		}
	}

	stack := []Node{{Reindeer: start, Score: 0}}

	for len(stack) > 0 {
		n := stack[0]
		stack = stack[1:]

		unvisted := checkPaths(n.Reindeer.Location)
		for _, c := range unvisted {
			var next Node

			// turn
			if n.Reindeer.Direction != c.Direction {
				next.Score = n.Score + 1000
			} else {
				next.Score = n.Score
			}

			// move
			next.Score++
			next.Reindeer = c

			// If the path is shorter add it to the visited maze
			score := grid.SafeGet(VisitedMaze, next.Reindeer.Location)
			if score == 0 || next.Score < score {
				stack = append(stack, next)
				grid.Set(VisitedMaze, next.Reindeer.Location, next.Score)
			}
		}
	}

	grid.OutputInt(VisitedMaze)

	fmt.Println(grid.SafeGet(VisitedMaze, end))
}

func main() {
	// Read input
	//data, _ := os.ReadFile("example.txt")
	//data, _ := os.ReadFile("example_1.txt")
	data, _ := os.ReadFile("input.txt")
	input := string(data)

	// Parse input
	Maze = grid.Parse(input, "\r\n")
	VisitedMaze = grid.Create[int](len(Maze[0]), len(Maze))

	defer timeTrack(time.Now())

	run()
}
