package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

type Rule struct {
	X int
	Y int
}

type KV struct {
	K int
	V int
}

func timeTrack(start time.Time) {
	elapsed := time.Since(start)
	fmt.Println(elapsed)
}

func safeAccess(arr [][]rune, x int, y int) rune {
	defer func() {
		if r := recover(); r != nil {
		}
	}()

	return arr[y][x]
}

func test(order_rules []Rule, updates [][]int) int {
	total := 0
	var incorrect_updates [][]KV

	for _, update := range updates {
		correct := true

		update_set := make(map[int]int)
		for _, x := range update {
			update_set[x]++
		}

		set := make(map[int]int)
		for _, rule := range order_rules {
			if update_set[rule.X] > 0 && update_set[rule.Y] > 0 {
				set[rule.X]++
			}
		}

		count := len(set) + 1
		for _, x := range update {
			if set[x] > count {
				correct = false

				var incorrect_update []KV
				for k, v := range set {
					incorrect_update = append(incorrect_update, KV{k, v})
				}

				incorrect_updates = append(incorrect_updates, incorrect_update)
				break
			} else {
				count = set[x]
			}
		}

		if correct {
			mid := len(update) / 2
			total += update[mid]
		}
	}

	total = 0
	for _, update := range incorrect_updates {
		sort.Slice(update, func(i, j int) bool {
			return update[i].V > update[j].V
		})

		mid := len(update) / 2
		total += update[mid].K
	}

	return total
}

func main() {
	// Read input
	//data, _ := os.ReadFile("example.txt")
	data, _ := os.ReadFile("input.txt")
	input := string(data)

	// Parse sections
	sections := strings.Split(input, "\r\n\r\n")

	// Parse rules
	var order_rules []Rule
	for _, lines := range strings.Fields(sections[0]) {
		v := strings.Split(lines, "|")
		x, _ := strconv.Atoi(v[0])
		y, _ := strconv.Atoi(v[1])
		rule := Rule{x, y}
		order_rules = append(order_rules, rule)
	}

	// Parse updates
	var updates [][]int
	lines := strings.Split(sections[1], "\r\n")
	for _, line := range lines {
		var update []int
		for _, x := range strings.Split(line, ",") {
			v, _ := strconv.Atoi(x)
			update = append(update, v)
		}
		updates = append(updates, update)
	}

	defer timeTrack(time.Now())

	total := test(order_rules, updates)
	fmt.Printf("Total: %d\r\n", total)
}
