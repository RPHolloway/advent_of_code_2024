package main

import (
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

func test(report []int) bool {
	unsafe := 0
	safe := true

	increasing := report[0] < report[1]
	for i := 0; i < len(report)-1; i++ {
		delta := int(math.Abs(float64(report[i] - report[i+1])))
		if delta == 0 || delta > 3 {
			unsafe++
			report = append(report[:i+1], report[i+2:]...)
			i = -1
		} else if (report[i] < report[i+1]) != increasing {
			unsafe++
			report = append(report[:i+1], report[i+2:]...)
			i = -1
		}

		if unsafe > 1 {
			safe = false
			break
		}
	}

	return safe
}

func main() {
	var safe_count int = 0

	// Read input
	//data, _ := os.ReadFile("example.txt")
	data, _ := os.ReadFile("input.txt")
	input := string(data)

	// Parse input
	var reports [][]int
	lines := strings.Split(input, "\n")
	for _, line := range lines {
		var report []int
		for _, x := range strings.Fields((line)) {
			v, _ := strconv.Atoi(x)
			report = append(report, v)
		}
		reports = append(reports, report)
	}

	// Verify reports
	for _, report := range reports {
		fwd := make([]int, len(report))
		copy(fwd, report)

		rev := make([]int, len(report))
		copy(rev, report)
		slices.Reverse(rev)

		if test(fwd) {
			fmt.Println(report)
			safe_count++
		} else if test(rev) {
			fmt.Println(report)
			safe_count++
		}
	}

	fmt.Printf("%d\n", safe_count)
}
