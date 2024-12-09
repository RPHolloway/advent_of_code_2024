package main

import (
	"fmt"
	"os"
	"strings"
	"time"
	"unicode"
)

type Point struct {
	X int
	Y int
}

func timeTrack(start time.Time) {
	elapsed := time.Since(start)
	fmt.Println(elapsed)
}

func is_antenna(x rune) bool {
	return unicode.IsLetter(x) || unicode.IsDigit(x)
}

func factorial(n int) int {
	result := 1
	for i := 2; i <= n; i++ {
		result *= i
	}
	return result
}

func combinations(n, k int) int {
	if k > n {
		return 0
	}
	return factorial(n) / (factorial(k) * factorial(n-k))
}

func test(city [][]rune) int {
	city_height := len(city)
	city_width := 0

	// Find antennas in the city
	antennas := make(map[rune][]Point)
	for y, street := range city {
		city_width = len(street)
		for x, roof := range street {
			if is_antenna(roof) {
				antennas[roof] = append(antennas[roof], Point{x, y})
			}
		}
	}

	// Calculate the antinodes
	antinodes := make(map[Point]struct{})
	for _, antennas := range antennas {
		antenna_pairs := combinations(len(antennas), 2)
		for i, a1 := range antennas {
			// check if all combinations have been tried
			if i >= antenna_pairs {
				break
			}

			// Calculate antinode locations
			for _, a2 := range antennas {
				// Don't compare against self
				if a1 == a2 {
					continue
				}

				delta_y := a1.Y - a2.Y
				delta_x := a1.X - a2.X

				antinode1 := Point{a1.X + delta_x, a1.Y + delta_y}
				antinode2 := Point{a2.X - delta_x, a2.Y - delta_y}

				antinodes[antinode1] = struct{}{}
				antinodes[antinode2] = struct{}{}
			}
		}
	}

	// Remove antinodes that are outside of the city
	for node := range antinodes {
		if node.X < 0 || node.Y < 0 || node.X >= city_width || node.Y >= city_height {
			delete(antinodes, node)
		}
	}

	return len(antinodes)
}

func main() {
	var city [][]rune

	// Read input
	//data, _ := os.ReadFile("example.txt")
	data, _ := os.ReadFile("input.txt")
	input := string(data)

	// Parse input
	lines := strings.Split(input, "\r\n")
	for _, line := range lines {
		city = append(city, []rune(line))
	}

	defer timeTrack(time.Now())

	total := test(city)
	fmt.Printf("Total: %d\r\n", total)
}
