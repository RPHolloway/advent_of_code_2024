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
	Turned   bool
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
			grid.Set(VisitedMaze, p.Add(dir), 777777)
		}
	}

	return unvisited
}

func checkShortest(p grid.Point) []grid.Point {
	var shortest []grid.Point

	score := grid.SafeGet(VisitedMaze, p)
	for i, dir := range grid.Directions {
		if grid.CheckDirection(VisitedMaze, p, i) < score {
			shortest = append(shortest, p.Add(dir))
		}
	}

	return shortest
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
				if !n.Turned {
					next.Score = n.Score + 1000
				} else {
					next.Score = n.Score
				}
			} else {
				next.Score = n.Score

				if grid.CheckDirection(Maze, c.Location, c.Direction) == '#' &&
					grid.SafeGet(Maze, c.Location) != 'E' {
					next.Score += 1000
					next.Turned = true
				}
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

	seats := make(map[grid.Point]struct{})
	seats[end] = struct{}{}

	pathStack := []grid.Point{end}
	for len(pathStack) > 0 {
		p := pathStack[0]
		pathStack = pathStack[1:]

		paths := checkShortest(p)

		score := grid.SafeGet(VisitedMaze, p)
		grid.Set(VisitedMaze, p, score%100)

		for _, path := range paths {
			seats[path] = struct{}{}
			pathStack = append(pathStack, path)
		}
	}

	grid.OutputInt(VisitedMaze)

	fmt.Println(grid.SafeGet(VisitedMaze, end))
	fmt.Println(len(seats))
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
