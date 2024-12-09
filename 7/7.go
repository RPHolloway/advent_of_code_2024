package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

const MAX_THREADS int = 100000

type Problem struct {
	Solution int
	Operands []int
}

func timeTrack(start time.Time) {
	elapsed := time.Since(start)
	fmt.Println(elapsed)
}

func concat(v1 int, v2 int) int {
	r, _ := strconv.Atoi(strconv.Itoa(v1) + strconv.Itoa(v2))
	return r
}

func check_next(answer int, operands []int, i int, solution int) int {
	add := answer + operands[i]
	mul := answer * operands[i]
	cc := concat(answer, operands[i])
	i++

	// check to see if we are at the end
	if i >= len(operands) {
		if add == solution || mul == solution || cc == solution {
			return solution
		} else {
			return 0
		}
	}

	if add <= solution {
		answer = check_next(add, operands, i, solution)
		if answer == solution {
			return answer
		}
	}

	if mul <= solution {
		answer = check_next(mul, operands, i, solution)
		if answer == solution {
			return answer
		}
	}

	if cc <= solution {
		answer = check_next(cc, operands, i, solution)
		if answer == solution {
			return answer
		}
	}

	// both solutions were too high
	return 0
}

func test(problems []Problem) int {
	total := 0

	for _, problem := range problems {
		answer := check_next(problem.Operands[0], problem.Operands, 1, problem.Solution)
		if answer == problem.Solution {
			total += problem.Solution
		}
	}

	return total
}

func main() {
	var problems []Problem

	// Read input
	//data, _ := os.ReadFile("example.txt")
	data, _ := os.ReadFile("input.txt")
	input := string(data)

	// Parse input
	lines := strings.Split(input, "\r\n")
	for _, line := range lines {
		x := strings.Split(line, ":")
		solution, _ := strconv.Atoi(x[0])

		var operands []int
		for _, v := range strings.Fields(x[1]) {
			x, _ := strconv.Atoi(v)
			operands = append(operands, x)
		}

		problems = append(problems, Problem{
			solution,
			operands,
		})
	}

	defer timeTrack(time.Now())

	total := test(problems)
	fmt.Printf("Total: %d\r\n", total)
}
