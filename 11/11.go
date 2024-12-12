package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

var stone_map = make(map[int][]int)

func timeTrack(start time.Time) {
	elapsed := time.Since(start)
	fmt.Println(elapsed)
}

func blink(stones map[int]int) map[int]int {
	result := make(map[int]int)

	for stone, count := range stones {
		if count == 0 {
			continue
		}

		if r, ok := stone_map[stone]; ok {
			for _, s := range r {
				result[s] += count
			}
		} else {
			stone_str := strconv.Itoa(stone)
			stone_length := len(stone_str)

			if stone == 0 {
				stone_map[stone] = append(stone_map[stone], 1)
				result[1] += count
			} else if stone_length%2 == 0 {
				v, _ := strconv.Atoi(stone_str[0 : stone_length/2])
				stone_map[stone] = append(stone_map[stone], v)
				result[v] += count

				v, _ = strconv.Atoi(stone_str[stone_length/2:])
				stone_map[stone] = append(stone_map[stone], v)
				result[v] += count
			} else {
				r := stone * 2024
				stone_map[stone] = append(stone_map[stone], r)
				result[r] += count
			}
		}
	}

	return result
}

func test(stones []int) int {
	total := 0

	stone_count := make(map[int]int)
	for _, stone := range stones {
		stone_count[stone]++
	}

	for i := 0; i < 75; i++ {
		stone_count = blink(stone_count)
		fmt.Println(i)
	}

	for _, count := range stone_count {
		total += count
	}

	return total
}

func main() {
	var stones []int

	// Read input
	//data, _ := os.ReadFile("example.txt")
	data, _ := os.ReadFile("input.txt")
	input := string(data)

	// Parse input
	for _, v := range strings.Fields(input) {
		x, _ := strconv.Atoi(v)
		stones = append(stones, x)
	}

	defer timeTrack(time.Now())

	total := test(stones)

	fmt.Printf("Total: %d\r\n", total)
}
