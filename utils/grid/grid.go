package grid

import (
	"fmt"
	"os"
	"strings"
)

const (
	DIR_UP = iota
	DIR_RIGHT
	DIR_DOWN
	DIR_LEFT
	DIR_COUNT
)

var Directions = []Point{
	{0, -1}, // up
	{1, 0},  // right
	{0, 1},  // down
	{-1, 0}, // left
}

type Point struct {
	X int
	Y int
}

type Cursor struct {
	Location  Point
	Direction int
}

func Parse(input string, delimiter string) [][]rune {
	var arr [][]rune

	lines := strings.Split(input, delimiter)
	for _, line := range lines {
		arr = append(arr, []rune(line))
	}

	return arr
}

func ParseInt(input string, delimiter string) [][]int {
	var arr [][]int

	lines := strings.Split(input, delimiter)
	for _, line := range lines {
		var row []int
		for _, c := range line {
			row = append(row, int(c-'0'))
		}
		arr = append(arr, row)
	}

	return arr
}

func Create[T any](width int, height int) [][]T {
	slice := make([][]T, height)
	for i := range slice {
		slice[i] = make([]T, width)
	}

	return slice
}

func Output(arr [][]rune) {
	var builder strings.Builder
	for _, row := range arr {
		builder.WriteString(string(row))
		builder.WriteString("\n")
	}
	os.WriteFile("output.txt", []byte(builder.String()), 0644)
}

func OutputInt(arr [][]int) {
	var builder strings.Builder
	for _, row := range arr {
		for _, i := range row {
			fmt.Fprintf(&builder, "%06d ", i)
		}
		builder.WriteString("\n")
	}
	os.WriteFile("output.txt", []byte(builder.String()), 0644)
}

func GetSize[T any](arr [][]T) (int, int) {
	height := len(arr)
	width := len(arr[0])

	return width, height
}

func SafeGet[T any](arr [][]T, p Point) T {
	defer func() {
		if r := recover(); r != nil {
		}
	}()

	return arr[p.Y][p.X]
}

func Set[T any](arr [][]T, p Point, v T) {
	arr[p.Y][p.X] = v
}

func (p1 Point) Add(p2 Point) Point {
	return Point{
		X: p1.X + p2.X,
		Y: p1.Y + p2.Y,
	}
}

func CheckDirection[T any](arr [][]T, p Point, dir int) T {
	next := p.Add(Directions[dir])
	return SafeGet(arr, next)
}
