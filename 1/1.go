package main

import (
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	var left_list, right_list []int
	var total_distance, similarity_score int = 0, 0

	// read input
	//data, _ := os.ReadFile("example_1.txt")
	data, _ := os.ReadFile("input_1.txt")
	input := string(data)

	// parse input
	lines := strings.Split(input, "\n")
	for _, line := range lines {
		values := strings.Fields(line)

		l, _ := strconv.Atoi(values[0])
		left_list = append(left_list, l)

		r, _ := strconv.Atoi(values[1])
		right_list = append(right_list, r)
	}

	// sort lists from smallest to largests
	sort.Ints(left_list)
	sort.Ints(right_list)

	// calculate the difference
	for i, r := range right_list {
		total_distance += int(math.Abs(float64(r - left_list[i])))
	}

	fmt.Printf("%d\n", total_distance)

	counts := make(map[int]int)
	for _, x := range right_list {
		counts[x]++
	}

	for _, x := range left_list {
		count, exists := counts[x]
		if exists {
			similarity_score += x * count
		}
	}

	fmt.Printf("%d\n", similarity_score)
}
