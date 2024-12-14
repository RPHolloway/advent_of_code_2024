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
	Corners   int
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

func check_diagonals(garden [][]int, p grid.Point, plot_id int) int {
	corners := 0

	if grid.SafeGet(garden, p.Add(grid.Directions[grid.DIR_UP])) != plot_id {

		if grid.CheckDirection(garden, p, grid.DIR_LEFT) == plot_id && grid.CheckDirection(garden, p.Add(grid.Directions[grid.DIR_UP]), grid.DIR_LEFT) == plot_id {
			corners++
		}

		if grid.CheckDirection(garden, p, grid.DIR_RIGHT) == plot_id && grid.CheckDirection(garden, p.Add(grid.Directions[grid.DIR_UP]), grid.DIR_RIGHT) == plot_id {
			corners++
		}
	}

	if grid.SafeGet(garden, p.Add(grid.Directions[grid.DIR_DOWN])) != plot_id {
		if grid.CheckDirection(garden, p, grid.DIR_LEFT) == plot_id && grid.CheckDirection(garden, p.Add(grid.Directions[grid.DIR_DOWN]), grid.DIR_LEFT) == plot_id {
			corners++
		}

		if grid.CheckDirection(garden, p, grid.DIR_RIGHT) == plot_id && grid.CheckDirection(garden, p.Add(grid.Directions[grid.DIR_DOWN]), grid.DIR_RIGHT) == plot_id {
			corners++
		}
	}

	return corners
}

func fill_garden(garden [][]rune, start grid.Point, plot_id int, labeled_garden [][]int) {
	stack := []grid.Point{start}
	blank := grid.Point{}
	var next grid.Point

	for len(stack) > 0 {
		seed := stack[0]
		stack = stack[1:]

		plant := grid.SafeGet(garden, seed)

		up_path, down_path := false, false
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

		up_path, down_path = false, false
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

			neighbors := 0
			for dir := range grid.DIR_COUNT {
				if grid.CheckDirection(garden, location, dir) == plot_id {
					neighbors++
				}

			}

			if neighbors == 0 {
				plot.Corners += 4
			} else if neighbors >= 1 {
				if neighbors == 1 {
					plot.Corners += 2
				}

				plot.Corners += check_diagonals(garden, location, plot_id)

				if neighbors >= 2 {
					if (grid.CheckDirection(garden, location, grid.DIR_UP) != grid.CheckDirection(garden, location, grid.DIR_DOWN)) &&
						(grid.CheckDirection(garden, location, grid.DIR_LEFT) != grid.CheckDirection(garden, location, grid.DIR_RIGHT)) {
						plot.Corners += 1
					}
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

func calculate_bulk_price(plot Plot) int {
	return plot.Area * plot.Corners
}

func test(garden [][]rune) int {
	total := 0
	bulkTotal := 0

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
		bulkTotal += calculate_bulk_price(plot)
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
