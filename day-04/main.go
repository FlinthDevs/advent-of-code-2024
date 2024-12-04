package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type Letter struct {
	value string
	valid bool
}

func getLines(filename string) []string {
	data, err := os.ReadFile(filename)
	check(err)
	return strings.Split(string(data), "\n")
}

func secondPart(lines []string) int {
	xmasCount := 0
	linesCount := len(lines)
	lineLength := 0

	for n, line := range lines {
		if line == "" || n < 1 || n > linesCount-3 {
			continue
		}

		lineLength = len(line)

		for i, c := range line {
			if c != 'A' || i < 1 || i > lineLength-2 {
				continue
			}

			topLeft := lines[n-1][i-1]
			topRight := lines[n-1][i+1]
			bottomLeft := lines[n+1][i-1]
			bottomRight := lines[n+1][i+1]

			if (topLeft == 'S' && bottomRight == 'M' || topLeft == 'M' && bottomRight == 'S') &&
				(topRight == 'S' && bottomLeft == 'M' || topRight == 'M' && bottomLeft == 'S') {
				xmasCount++
			}
		}
	}

	return xmasCount
}

func firstPart(lines []string, debug bool) int {
	xmasCount := 0
	linesCount := len(lines)
	lineLength := 0
	niceLines := make([][]Letter, linesCount)

	for n := range lines {
		newLine := make([]Letter, len(lines[n]))

		for i := range lines[n] {
			newLine[i] = Letter{value: string(lines[n][i]), valid: false}
		}
		niceLines[n] = newLine
	}

	for n, line := range lines {
		if line == "" {
			continue
		}

		lineLength = len(line)

		for i, c := range line {
			// Not intersting, skip
			if c != 'X' {
				continue
			}

			if i < lineLength-3 {
				// Right
				if line[i+1] == 'M' && line[i+2] == 'A' && line[i+3] == 'S' {
					xmasCount++
					niceLines[n][i].valid = true
					niceLines[n][i+1].valid = true
					niceLines[n][i+2].valid = true
					niceLines[n][i+3].valid = true
				}

				// Bottom Right
				if n < linesCount-4 {
					if lines[n+1][i+1] == 'M' && lines[n+2][i+2] == 'A' && lines[n+3][i+3] == 'S' {
						xmasCount++
						niceLines[n][i].valid = true
						niceLines[n+1][i+1].valid = true
						niceLines[n+2][i+2].valid = true
						niceLines[n+3][i+3].valid = true
					}
				}

				// Top right
				if n > 2 {
					if lines[n-1][i+1] == 'M' && lines[n-2][i+2] == 'A' && lines[n-3][i+3] == 'S' {
						xmasCount++
						niceLines[n][i].valid = true
						niceLines[n-1][i+1].valid = true
						niceLines[n-2][i+2].valid = true
						niceLines[n-3][i+3].valid = true
					}
				}
			}

			if i > 2 {
				// Left
				if line[i-1] == 'M' && line[i-2] == 'A' && line[i-3] == 'S' {
					xmasCount++
					niceLines[n][i].valid = true
					niceLines[n][i-1].valid = true
					niceLines[n][i-2].valid = true
					niceLines[n][i-3].valid = true
				}

				// Bottom Left
				if n < linesCount-4 {
					if lines[n+1][i-1] == 'M' && lines[n+2][i-2] == 'A' && lines[n+3][i-3] == 'S' {
						xmasCount++
						niceLines[n][i].valid = true
						niceLines[n+1][i-1].valid = true
						niceLines[n+2][i-2].valid = true
						niceLines[n+3][i-3].valid = true
					}
				}

				// Top Left
				if n > 2 {
					if lines[n-1][i-1] == 'M' && lines[n-2][i-2] == 'A' && lines[n-3][i-3] == 'S' {
						xmasCount++
						niceLines[n][i].valid = true
						niceLines[n-1][i-1].valid = true
						niceLines[n-2][i-2].valid = true
						niceLines[n-3][i-3].valid = true
					}
				}
			}

			// Bottom
			if n < linesCount-4 {
				if lines[n+1][i] == 'M' && lines[n+2][i] == 'A' && lines[n+3][i] == 'S' {
					xmasCount++
					niceLines[n][i].valid = true
					niceLines[n+1][i].valid = true
					niceLines[n+2][i].valid = true
					niceLines[n+3][i].valid = true
				}
			}

			// Top
			if n > 2 {
				if lines[n-1][i] == 'M' && lines[n-2][i] == 'A' && lines[n-3][i] == 'S' {
					xmasCount++
					niceLines[n][i].valid = true
					niceLines[n-1][i].valid = true
					niceLines[n-2][i].valid = true
					niceLines[n-3][i].valid = true
				}
			}
		}
	}

	if debug {
		for _, l := range niceLines {
			for _, c := range l {
				if c.valid {
					fmt.Printf("%s", c.value)
				} else {
					fmt.Printf(".")
				}
			}
			fmt.Println()
		}
	}

	return xmasCount
}

func main() {
	debug := false
	lines := getLines("./input.txt")
	start := time.Now()
	fmt.Printf("First part result: %v (%v)\n", firstPart(lines, debug), time.Since(start))
	start = time.Now()
	fmt.Printf("Second part result: %v (%v)\n", secondPart(lines), time.Since(start))
}
