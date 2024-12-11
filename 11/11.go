package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func timeTrack(start time.Time) {
	elapsed := time.Since(start)
	fmt.Println(elapsed)
}

func blink(stones []int) []int {
	var result []int

	for _, stone := range stones {
		stone_str := strconv.Itoa(stone)
		stone_length := len(stone_str)

		if stone == 0 {
			result = append(result, 1)
		} else if stone_length%2 == 0 {
			v, _ := strconv.Atoi(stone_str[0 : stone_length/2])
			result = append(result, v)
			v, _ = strconv.Atoi(stone_str[stone_length/2:])
			result = append(result, v)
		} else {
			result = append(result, stone*2024)
		}
	}

	return result
}

func test(stones []int) int {
	for i := 0; i < 25; i++ {
		stones = blink(stones)
	}

	return len(stones)
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
