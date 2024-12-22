package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

const FILE_NAME = "input.txt"

type IntSliceKey [4]int

var Secrets []int

func timeTrack(start time.Time) {
	elapsed := time.Since(start)
	fmt.Println(elapsed)
}

func nextSecret(secret int) int {
	v := secret * 64
	secret ^= v
	secret = secret % 0x1000000

	v = secret / 32
	secret ^= v
	secret = secret % 0x1000000

	v = secret * 2048
	secret ^= v
	secret = secret % 0x1000000

	return secret
}

func run() {
	profitSet := make(map[IntSliceKey]int)
	for _, secret := range Secrets {
		var change []int
		patternSet := make(map[IntSliceKey]int)

		for range 2000 {
			next := nextSecret(secret)

			s := next % 10
			s1 := secret % 10
			delta := s - s1
			change = append(change, delta)
			if len(change) > 4 {
				change = change[1:]
			}

			if len(change) == 4 {
				if _, ok := patternSet[IntSliceKey(change)]; !ok {
					patternSet[IntSliceKey(change)] = s
				}
			}

			secret = next
		}

		for k, v := range patternSet {
			profitSet[k] += v
		}
	}

	total := 0
	for _, secret := range Secrets {
		total += secret
	}
	fmt.Printf("Secret Total: %d\n", total)

	var maxProfitPattern IntSliceKey
	maxProfit := 0
	for k, v := range profitSet {
		if v > maxProfit {
			maxProfit = v
			maxProfitPattern = k
		}
	}
	fmt.Printf("Max Profit: %d %d\n", maxProfit, maxProfitPattern)
}

func main() {
	// Read input
	data, _ := os.ReadFile(FILE_NAME)
	input := string(data)

	// Parse input
	lines := strings.Split(input, "\r\n")
	for _, line := range lines {
		v, _ := strconv.Atoi(line)
		Secrets = append(Secrets, v)
	}

	defer timeTrack(time.Now())

	run()
}
