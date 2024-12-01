package main

import (
	"errors"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
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

func parseLine(line string) (int, int, error) {
	numbers := strings.Split(line, " ")

	if len(numbers) < 4 {
		return 0, 0, errors.New("not enough fields")
	}

	first, errF := strconv.Atoi(numbers[0])
	second, errS := strconv.Atoi(numbers[3])

	check(errF)
	check(errS)

	return first, second, nil
}

func firstPart(lines []string) int {
	linesCount := len(lines)
	leftCol := make([]int, linesCount, linesCount)
	rightCol := make([]int, linesCount, linesCount)

	// Storing each number in a col
	for i, line := range lines {
		first, second, err := parseLine(line)

		if err != nil {
			continue
		}

		leftCol[i] = first
		rightCol[i] = second
	}

	// Sorting
	sort.Ints(leftCol)
	sort.Ints(rightCol)

	// Doing the actual differences
	sum := 0
	diff := 0

	for i := range len(leftCol) {
		diff = leftCol[i] - rightCol[i]

		// math.Abs requires float64, so avoiding types boring stuff here to get abs.
		if diff < 0 {
			diff = -diff
		}

		sum += diff
	}

	return sum
}

func secondPart(lines []string) int {
	sum := 0
	counts := map[int]int{}

	// First loop: count occurences of each number in the 2nd column in a map.
	for _, line := range lines {
		_, second, err := parseLine(line)
		counts[second]++

		if err != nil {
			continue
		}
	}

	// Second loop: For each first column number, multiply by occurences in the second.
	for _, line := range lines {
		first, _, err := parseLine(line)

		if err != nil {
			continue
		}

		firstOccurencesCount, ok := counts[first]

		if !ok {
			continue
		}

		sum += first * firstOccurencesCount
	}

	return sum
}

func main() {
	lines := getLines("./input.txt")

	start := time.Now()
	fmt.Printf("First part result: %v (%v)\n", firstPart(lines), time.Since(start))

	start = time.Now()
	fmt.Printf("Second part result: %v, (%v)\n", secondPart(lines), time.Since(start))
}
