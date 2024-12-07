package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

const debugEnabled = false

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
	return 0
}

func placeTrap(lines *[]string, coords [3]int, trap rune) {
	lineChars := []rune((*lines)[coords[0]])
	lineChars[coords[1]] = trap
	(*lines)[coords[0]] = string(lineChars)
}

func firstPart(lines []string) int {
	position := [2]int{-1, -1}
	direction := Up
	width := -1
	height := len(lines)
	histories := make([][][]rune, 0, len(lines))

	for y, line := range lines {
		if line == "" {
			continue
		}

		histories = append(histories, make([][]rune, len(line)))

		if width == -1 {
			width = len(line)
		}

		for x, c := range line {
			if c == '^' {
				position[0] = x
				position[1] = y
			}

			histories[y][x] = make([]rune, 0, 10)
			histories[y][x] = append(histories[y][x], c)
		}
	}

	debug("Starting at %v:%v\n", position[0], position[1])
	for {
		updatePosition(&histories, &position, &direction, width, height)

		if position[0] == -1 {
			break
		}
	}

	count := 0

	for _, l := range histories {
		debug("%c\n", l)
		for _, c := range l {
			found := false
			for _, r := range c {
				if r != '#' && r != '.' {
					found = true
				}
			}
			if found {
				count++
			}
		}

	}

	return count
}

func updatePosition(lines *[][][]rune, pos *[2]int, d *Direction, w int, h int) {
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

		// Ouf of bounds
		if nextPosition[0] > w-1 ||
			nextPosition[0] < 0 ||
			nextPosition[1] > h-1 ||
			nextPosition[1] < 0 ||
			nextPosition[1] > len(*lines)-1 ||
			nextPosition[0] > len((*lines)[0])-1 {
			debug("BREAK\n")
			(*pos)[0] = -1
			(*pos)[1] = -1
			*d = Up
			return
		}

		if (*lines)[nextPosition[1]][nextPosition[0]][0] == '#' {
			debug("STOP - Obstacle met at %v:%v\n", nextPosition[0], nextPosition[1])
			*pos = prevPosition
			turnRight(d)
			(*lines)[prevPosition[1]][prevPosition[0]] = append((*lines)[prevPosition[1]][prevPosition[0]], getTurnChar(*d))
			return
		}

		(*lines)[nextPosition[1]][nextPosition[0]] = append((*lines)[nextPosition[1]][nextPosition[0]], getTurnChar(*d))
		prevPosition = nextPosition
	}
}

func getTurnChar(d Direction) rune {
	switch d {
	case Up:
		return 'N'
	case Right:
		return 'E'
	case Down:
		return 'S'
	case Left:
		return 'W'
	}

	return 'O'
}
func getDirectionChar(d Direction) rune {
	switch d {
	case Up:
		return 'U'
	case Right:
		return 'R'
	case Down:
		return 'D'
	case Left:
		return 'L'
	}

	return 'O'
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
	lines := getLines("./input.txt")
	start := time.Now()
	fmt.Printf("First part result: %v (%v)\n", firstPart(lines), time.Since(start))
	start = time.Now()
	// fmt.Printf("Second part result: %v (%v)\n", secondPart(lines), time.Since(start))
}
