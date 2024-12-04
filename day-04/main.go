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

func firstPart(lines []string) int {
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
		lineLength = len(line)

		for i, c := range line {
			if c != 'X' {
				continue
			}

			if i < lineLength-4 {
				if line[i+1] == 'M' && line[i+2] == 'A' && line[i+3] == 'S' {
					xmasCount++
					niceLines[n][i].valid = true
					niceLines[n][i+1].valid = true
					niceLines[n][i+2].valid = true
					niceLines[n][i+3].valid = true
				}

				if n < linesCount-4 {
					if lines[n+1][i+1] == 'M' && lines[n+2][i+2] == 'A' && lines[n+3][i+3] == 'S' {
						xmasCount++
						niceLines[n][i].valid = true
						niceLines[n+1][i+1].valid = true
						niceLines[n+2][i+2].valid = true
						niceLines[n+3][i+3].valid = true
					}
				}

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
				if line[i-1] == 'M' && line[i-2] == 'A' && line[i-3] == 'S' {
					xmasCount++
					niceLines[n][i].valid = true
					niceLines[n][i-1].valid = true
					niceLines[n][i-2].valid = true
					niceLines[n][i-3].valid = true
				}

				if n < linesCount-4 {
					if lines[n+1][i-1] == 'M' && lines[n+2][i-2] == 'A' && lines[n+3][i-3] == 'S' {
						xmasCount++
						niceLines[n][i].valid = true
						niceLines[n+1][i-1].valid = true
						niceLines[n+2][i-2].valid = true
						niceLines[n+3][i-3].valid = true
					}
				}

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

			if n < linesCount-4 {
				if lines[n+1][i] == 'M' && lines[n+2][i] == 'A' && lines[n+3][i] == 'S' {
					xmasCount++
					niceLines[n][i].valid = true
					niceLines[n+1][i].valid = true
					niceLines[n+2][i].valid = true
					niceLines[n+3][i].valid = true
				}
			}

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

	return xmasCount
}

func secondPart() int {
	return 0
}

func main() {
	// Not 2334 2324?
	lines := getLines("./test_input_2.txt")
	start := time.Now()
	fmt.Printf("First part result: %v (%v)\n", firstPart(lines), time.Since(start))
	start = time.Now()
	fmt.Printf("Second part result: %v (%v)\n", secondPart(), time.Since(start))
}
