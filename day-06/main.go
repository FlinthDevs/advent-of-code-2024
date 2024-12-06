package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

const debugEnabled = true

type Direction int

const (
	Up    Direction = 0
	Right Direction = 1
	Down  Direction = 2
	Left  Direction = 3
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func getLines(filename string) []string {
	data, err := os.ReadFile(filename)
	check(err)
	return strings.Split(string(data), "\n")
}

func secondPart(lines []string) int {
	position := [3]int{-1, -1, -1}
	direction := Up
	width := -1
	height := len(lines)

	for y, line := range lines {
		if line == "" {
			continue
		}

		if width == -1 {
			width = len(line)
		}

		for x, c := range line {
			if c == '^' {
				position[0] = x
				position[1] = y
				position[2] = int(Up)
				break
			}
		}

		if position[0] != -1 {
			break
		}
	}

	places := make([][3]int, 0, width*height)
	startingPosition := position
	startingDirection := direction

	debug("Starting at %v:%v\n", position[0], position[1])

	// Get all places we go through.
	for {
		updatePosition(lines, &position, &direction, &places, width, height, false)

		if position[0] == -1 {
			break
		}
	}

	loopsCount := 0
	for _, coords := range places {
		if lines[coords[0]][coords[1]] == '#' {
			continue
		}

		position = startingPosition
		direction = startingDirection
		newPlaces := make([][3]int, 0, width*height)

		debug("===========================\n")
		debug("Added trap at %v\n", coords)
		debug("===========================\n")
		placeTrap(&lines, coords, '#')

		hasLoop := false

		for {
			hasLoop := updatePosition(lines, &position, &direction, &newPlaces, width, height, true)

			if position[0] == -1 || hasLoop {
				debug("%v: %v\n", position, hasLoop)
				break
			}
		}

		if hasLoop {
			debug(" LOOOOOOOOOOOOP FOUND")
			loopsCount++
		} else {
			debug("    NO LOOP FOUND\n")
		}

		placeTrap(&lines, coords, '.')
	}

	return loopsCount
}

func placeTrap(lines *[]string, coords [3]int, trap rune) {
	lineChars := []rune((*lines)[coords[0]])
	lineChars[coords[1]] = trap
	(*lines)[coords[0]] = string(lineChars)
}

func firstPart(lines []string) int {
	position := [3]int{-1, -1, -1}
	direction := Up
	width := -1
	height := len(lines)

	for y, line := range lines {
		if line == "" {
			continue
		}

		if width == -1 {
			width = len(line)
		}

		for x, c := range line {
			if c == '^' {
				position[0] = x
				position[1] = y
				position[2] = int(Up)
				break
			}
		}

		if position[0] != -1 {
			break
		}
	}

	places := make([][3]int, 0, width*height)

	debug("Starting at %v:%v\n", position[0], position[1])
	for {
		updatePosition(lines, &position, &direction, &places, width, height, false)

		if position[0] == -1 {
			break
		}
	}

	return len(places)
}

func updatePosition(lines []string, pos *[3]int, d *Direction, places *[][3]int, w int, h int, handleLoop bool) bool {
	nextPosition := *pos
	prevPosition := *pos
	xFactor := 0
	yFactor := 0

	// Index starts from top-left
	if *d == Up {
		yFactor = -1
	} else if *d == Down {
		yFactor = 1
	}

	if *d == Right {
		xFactor = 1
	} else if *d == Left {
		xFactor = -1
	}

	for {
		debug("Moving %v\n", *d)
		nextPosition[0] += xFactor
		nextPosition[1] += yFactor
		debug("Direction: %v - New coords: %v\n", []int{xFactor, yFactor}, nextPosition)

		if nextPosition[0] > w-1 || nextPosition[0] < 0 || nextPosition[1] > h-1 || nextPosition[1] < 0 || lines[nextPosition[1]] == "" {
			debug("BREAK\n")
			*pos = [3]int{-1, -1, -1}
			*d = Up
			return false
		}

		if lines[nextPosition[1]][nextPosition[0]] == '#' {
			debug("STOP - Obstacle met at %v:%v\n", nextPosition[0], nextPosition[1])
			*pos = prevPosition
			turnRight(d)
			return false
		}

		hasLoop := addPlace(places, nextPosition, handleLoop)

		if handleLoop && hasLoop {
			return true
		}

		prevPosition = nextPosition
	}
}

func addPlace(places *[][3]int, newCoords [3]int, loopDetection bool) bool {
	found := false

	for _, coords := range *places {
		if coords[0] == newCoords[0] && coords[1] == newCoords[1] {
			if !loopDetection || coords[2] == newCoords[2] {
				debug("existing step step %v (%v)\n", newCoords, len(*places))
				found = true
			} else {
				debug("Found A loop !!!!")
				return true
			}
		}
	}

	if !found {
		debug("add new step %v (%v)\n", newCoords, len(*places))
		*places = append(*places, newCoords)
	}

	return false
}

func debug(format string, a ...any) {
	if debugEnabled {
		fmt.Printf(format, a...)
	}
}

func turnRight(d *Direction) {
	debug("Turned from %v to %v\n", *d, (*d+1)%4)
	*d = (*d + 1) % 4
}

func main() {
	lines := getLines("./test_input.txt")
	start := time.Now()
	// fmt.Printf("First part result: %v (%v)\n", firstPart(lines), time.Since(start))
	start = time.Now()
	fmt.Printf("Second part result: %v (%v)\n", secondPart(lines), time.Since(start))
}
