package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

const FILE_NAME = "input.txt"

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
	for range 2000 {
		for i := range Secrets {
			Secrets[i] = nextSecret(Secrets[i])
		}
	}

	total := 0
	for _, secret := range Secrets {
		total += secret
	}
	fmt.Println(total)
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
