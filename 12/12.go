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

func label_garden(garden [][]rune) [][]int {
	width, height := grid.GetSize(garden)
	labeled_garden := grid.Create[int](width, height)
	plot_count := 10

	for y, row := range garden {
		for x, plant := range row {
			location := grid.Point{X: x, Y: y}
			plot_id := grid.SafeGet(labeled_garden, location)
			if plot_id == 0 {
				plot_count++
				plot_id = plot_count
				grid.Set(labeled_garden, location, plot_id)
			}

			for dir := range grid.DIR_COUNT {
				if grid.CheckDirection(garden, location, dir) == plant {
					next_id := grid.CheckDirection(labeled_garden, location, dir)
					if next_id == 0 {
						grid.Set(labeled_garden, location.Add(grid.Directions[dir]), plot_id)
					} else if next_id < plot_id {
						grid.Set(labeled_garden, location, next_id)
					}
				}
			}
		}
	}

	for y := height - 1; y >= 0; y-- {
		for x := width - 1; x >= 0; x-- {
			location := grid.Point{X: x, Y: y}
			plant := grid.SafeGet(garden, location)
			plot_id := grid.SafeGet(labeled_garden, location)

			for dir := range grid.DIR_COUNT {
				if grid.CheckDirection(garden, location, dir) == plant {
					next_id := grid.CheckDirection(labeled_garden, location, dir)
					if next_id > plot_id {
						grid.Set(labeled_garden, location.Add(grid.Directions[dir]), plot_id)
					} else if next_id < plot_id {
						grid.Set(labeled_garden, location, next_id)
					}
				}
			}
		}
	}

	for y := 0; y <= height-1; y++ {
		for x := 0; x <= width-1; x++ {
			location := grid.Point{X: x, Y: y}
			plant := grid.SafeGet(garden, location)
			plot_id := grid.SafeGet(labeled_garden, location)

			for dir := range grid.DIR_COUNT {
				if grid.CheckDirection(garden, location, dir) == plant {
					next_id := grid.CheckDirection(labeled_garden, location, dir)
					if next_id > plot_id {
						grid.Set(labeled_garden, location.Add(grid.Directions[dir]), plot_id)
					} else if next_id < plot_id {
						grid.Set(labeled_garden, location, next_id)
					}
				}
			}
		}
	}

	return labeled_garden
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

	labeled_garden := label_garden(garden)
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
