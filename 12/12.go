package main

import (
	"advent_of_code/utils/grid"
	"fmt"
	"os"
	"time"
)

type Plot struct {
	Area      int
	Perimeter int
}

func timeTrack(start time.Time) {
	elapsed := time.Since(start)
	fmt.Println(elapsed)
}

func check_neighbor(garden [][]rune, labeled_garden [][]int, p grid.Point, dir int, plant rune) grid.Point {
	p = p.Add(grid.Directions[dir])

	if grid.SafeGet(garden, p) == 0 {
		return grid.Point{}
	}

	if grid.SafeGet(garden, p) == plant && grid.SafeGet(labeled_garden, p) == 0 {
		return p
	}

	return grid.Point{}
}

func fill_garden(garden [][]rune, start grid.Point, plot_id int, labeled_garden [][]int) {
	stack := []grid.Point{start}
	blank := grid.Point{}
	var next grid.Point

	for len(stack) > 0 {
		up_path, down_path := false, false
		seed := stack[0]
		stack = stack[1:]

		plant := grid.SafeGet(garden, seed)

		left := seed
		for grid.SafeGet(garden, left) == plant {
			grid.Set(labeled_garden, left, plot_id)

			next = check_neighbor(garden, labeled_garden, left, grid.DIR_UP, plant)
			if next == blank {
				up_path = false
			} else if !up_path {
				stack = append(stack, next)
				up_path = true
			}
			next = check_neighbor(garden, labeled_garden, left, grid.DIR_DOWN, plant)
			if next == blank {
				down_path = false
			} else if !down_path {
				stack = append(stack, next)
				down_path = true
			}

			left = left.Add(grid.Directions[grid.DIR_LEFT])
		}

		right := seed.Add(grid.Directions[grid.DIR_RIGHT])
		for grid.SafeGet(garden, right) == plant {
			grid.Set(labeled_garden, right, plot_id)

			next = check_neighbor(garden, labeled_garden, right, grid.DIR_UP, plant)
			if next == blank {
				up_path = false
			} else if !up_path {
				stack = append(stack, next)
				up_path = true
			}
			next = check_neighbor(garden, labeled_garden, right, grid.DIR_DOWN, plant)
			if next == blank {
				down_path = false
			} else if !down_path {
				stack = append(stack, next)
				down_path = true
			}

			right = right.Add(grid.Directions[grid.DIR_RIGHT])
		}
	}
}

func measure_plots(garden [][]int) map[int]Plot {
	plots := make(map[int]Plot)

	for y, row := range garden {
		for x, plot_id := range row {
			location := grid.Point{X: x, Y: y}
			plot := plots[plot_id]
			plot.Area++

			for dir := range grid.DIR_COUNT {
				if grid.CheckDirection(garden, location, dir) != plot_id {
					plot.Perimeter++
				}
			}

			plots[plot_id] = plot
		}
	}

	return plots
}

func calculate_price(plot Plot) int {
	return plot.Area * plot.Perimeter
}

func test(garden [][]rune) int {
	total := 0

	width, height := grid.GetSize(garden)
	labeled_garden := grid.Create[int](width, height)

	plot_count := 1
	for y, row := range labeled_garden {
		for x, plot_id := range row {
			if plot_id == 0 {
				fill_garden(garden, grid.Point{X: x, Y: y}, plot_count, labeled_garden)
				plot_count++
			}
		}
	}

	plots := measure_plots(labeled_garden)
	for _, plot := range plots {
		total += calculate_price(plot)
	}

	return total
}

func main() {
	var garden [][]rune

	// Read input
	//data, _ := os.ReadFile("example.txt")
	data, _ := os.ReadFile("input.txt")
	input := string(data)

	// Parse input
	garden = grid.Parse(input, "\r\n")

	defer timeTrack(time.Now())

	total := test(garden)

	fmt.Printf("Total: %d\r\n", total)
}
