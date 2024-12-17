package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"time"
)

var Registers []int
var Instructions []int

const (
	ADV = iota // A = A / 2**CO
	BXL        // B = B ^ O
	BST        // B = CO % 8
	JNZ        // if A != 0 JMP to O
	BXC        // B = B ^ C
	OUT        // CO % 8
	BDV        // B = A / 2**CO
	CDV        // C = A / 2**CO
)

func timeTrack(start time.Time) {
	elapsed := time.Since(start)
	fmt.Println(elapsed)
}

func comboOperator(co int) int {
	switch co {
	case 0:
		fallthrough
	case 1:
		fallthrough
	case 2:
		fallthrough
	case 3:
		return co
	case 4:
		fallthrough
	case 5:
		fallthrough
	case 6:
		return Registers[co-4]
	case 7:
		fallthrough
	default:
		// Invalid
	}

	return -1
}

func dv(co int) int {
	o := comboOperator(co)
	return int(float64(Registers[0]) / math.Pow(2, float64(o)))
}

func adv(co int) {
	Registers[0] = dv(co)
}

func bxl(o int) {
	Registers[1] = Registers[1] ^ o
}

func bst(co int) {
	o := comboOperator(co)
	Registers[1] = o % 8
}

func jnz(o int) int {
	return o - 2

}

func bxc() {
	Registers[1] = Registers[1] ^ Registers[2]
}

func out(co int) int {
	o := comboOperator(co)
	return o % 8
}

func bdv(co int) {
	Registers[1] = dv(co)
}

func cdv(co int) {
	Registers[2] = dv(co)
}

func toCSV(slice []int) string {
	var b strings.Builder
	for _, v := range slice {
		b.WriteString(fmt.Sprintf("%d,", v))
	}

	return b.String()
}

func run() {
	var output []int
	for pc := 0; pc < len(Instructions); pc += 2 {
		o := Instructions[pc+1]

		switch Instructions[pc] {
		case ADV:
			adv(o)
		case BXL:
			bxl(o)
		case BST:
			bst(o)
		case JNZ:
			if Registers[0] != 0 {
				pc = jnz(o)
			}
		case BXC:
			bxc()
		case OUT:
			output = append(output, out(o))
		case BDV:
			bdv(o)
		case CDV:
			cdv(o)
		}
	}

	fmt.Println(toCSV(output))
}

func main() {
	// Read input
	//data, _ := os.ReadFile("example_2.txt")
	data, _ := os.ReadFile("input.txt")
	input := string(data)

	// Parse input
	sections := strings.Split(input, "\r\n\r\n")
	registers := strings.Split(sections[0], "\r\n")
	for _, register := range registers {
		x := strings.Split(register, ": ")
		v, _ := strconv.Atoi(x[1])
		Registers = append(Registers, v)
	}

	program := strings.Split(sections[1], ": ")
	instructions := strings.Split(program[1], ",")
	for _, instruction := range instructions {
		v, _ := strconv.Atoi(instruction)
		Instructions = append(Instructions, v)
	}

	defer timeTrack(time.Now())

	run()
}
