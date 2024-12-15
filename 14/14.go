package main

import (
	"advent_of_code/utils/grid"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

const ROOM_HEIGHT int = 103
const ROOM_WIDTH int = 101

type Robot struct {
	Location grid.Point
	Velocity grid.Point
}

var Robots []Robot

var re = regexp.MustCompile(`p=(-?\d+),(-?\d+)\s+v=(-?\d+),(-?\d+)`)

func timeTrack(start time.Time) {
	elapsed := time.Since(start)
	fmt.Println(elapsed)
}

func (robot Robot) move() Robot {
	next := robot.Location.Add(robot.Velocity)
	robot.Location.X = next.X
	robot.Location.Y = next.Y

	if robot.Location.X >= ROOM_WIDTH {
		robot.Location.X = robot.Location.X - ROOM_WIDTH
	} else if robot.Location.X < 0 {
		robot.Location.X = ROOM_WIDTH + robot.Location.X
	}

	if robot.Location.Y >= ROOM_HEIGHT {
		robot.Location.Y = robot.Location.Y - ROOM_HEIGHT
	} else if robot.Location.Y < 0 {
		robot.Location.Y = ROOM_HEIGHT + robot.Location.Y
	}

	return robot
}

func computeSafetyFactor() int {
	var nw, ne, sw, se int

	for _, robot := range Robots {
		if robot.Location.X < ROOM_WIDTH/2 && robot.Location.Y < ROOM_HEIGHT/2 {
			nw++
		} else if robot.Location.X > ROOM_WIDTH/2 && robot.Location.Y < ROOM_HEIGHT/2 {
			ne++
		} else if robot.Location.X < ROOM_WIDTH/2 && robot.Location.Y > ROOM_HEIGHT/2 {
			sw++
		} else if robot.Location.X > ROOM_WIDTH/2 && robot.Location.Y > ROOM_HEIGHT/2 {
			se++
		}
	}

	return nw * ne * sw * se
}

func run() {
	safetyFactor := 0

	for r, robot := range Robots {
		for i := 0; i < 100; i++ {
			robot = robot.move()
		}
		//fmt.Println(robot.Location)
		Robots[r] = robot
	}

	safetyFactor = computeSafetyFactor()

	fmt.Println(safetyFactor)
}

func main() {
	// Read input
	//data, _ := os.ReadFile("example.txt")
	data, _ := os.ReadFile("input.txt")
	input := string(data)

	// Parse input
	lines := strings.Split(input, "\r\n")
	for _, line := range lines {
		matches := re.FindStringSubmatch(line)
		if matches != nil {
			var robot Robot
			v, _ := strconv.Atoi(matches[1])
			robot.Location.X = v
			v, _ = strconv.Atoi(matches[2])
			robot.Location.Y = v
			v, _ = strconv.Atoi(matches[3])
			robot.Velocity.X = v
			v, _ = strconv.Atoi(matches[4])
			robot.Velocity.Y = v

			Robots = append(Robots, robot)
		}
	}

	defer timeTrack(time.Now())

	run()
}
