package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type Point struct {
	X int
	Y int
}

const (
	up = iota
	right
	down
	left
)

var directions = []Point{
	{0, -1}, // up
	{1, 0},  // right
	{0, 1},  // down
	{-1, 0}, // left
}

func timeTrack(start time.Time) {
	elapsed := time.Since(start)
	fmt.Println(elapsed)
}

func safeAccess(arr [][]int, p Point) int {
	defer func() {
		if r := recover(); r != nil {
		}
	}()

	return arr[p.Y][p.X]
}

func (p1 Point) Add(p2 Point) Point {
	return Point{
		X: p1.X + p2.X,
		Y: p1.Y + p2.Y,
	}
}

func check_next(topo_map [][]int, p Point, peaks map[Point]struct{}) map[Point]struct{} {
	for _, direction := range directions {
		z := safeAccess(topo_map, p)
		next_p := p.Add(direction)
		next_z := safeAccess(topo_map, next_p)

		if next_z == z+1 {
			if next_z == 9 {
				peaks[next_p] = struct{}{}
			} else {
				check_next(topo_map, next_p, peaks)
			}
		}
	}

	return peaks
}

func test(topo_map [][]int) int {
	total := 0

	var trailheads []Point
	for y, row := range topo_map {
		for x, z := range row {
			if z == 0 {
				trailheads = append(trailheads, Point{x, y})
			}
		}
	}

	for _, trail := range trailheads {
		peaks := make(map[Point]struct{})
		peaks = check_next(topo_map, trail, peaks)
		total += len(peaks)
	}

	return total
}

func main() {
	var topo_map [][]int

	// Read input
	//data, _ := os.ReadFile("example_1.txt")
	//data, _ := os.ReadFile("example_2.txt")
	data, _ := os.ReadFile("input.txt")
	input := string(data)

	// Parse input
	lines := strings.Split(input, "\r\n")
	for _, line := range lines {
		var row []int
		for _, c := range line {
			v, _ := strconv.Atoi(string(c))
			row = append(row, v)
		}
		topo_map = append(topo_map, row)
	}

	defer timeTrack(time.Now())

	total := test(topo_map)

	fmt.Printf("Total: %d\r\n", total)
}
