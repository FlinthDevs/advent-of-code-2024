package main

import (
	"fmt"
	"os"
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

func secondPart(lines []string) int {
	middlesCount := 0
	startSorting := false
	rules := make([][][2]int, 100)
	numbers := make([]int, 0, 99)

	for _, line := range lines {
		if line == "" {
			continue
		}

		if line[0] == 'X' {
			startSorting = true
			continue
		}

		// Building sorting rules.
		if !startSorting {
			a, errA := strconv.Atoi(string(line[0]) + string(line[1]))
			b, errB := strconv.Atoi(string(line[3]) + string(line[4]))
			check(errA)
			check(errB)

			if len(rules[a]) == 0 {
				rules[a] = make([][2]int, 0, 1000)
			}
			if len(rules[b]) == 0 {
				rules[b] = make([][2]int, 0, 1000)
			}

			rules[a] = append(rules[a], [2]int{a, b})
			rules[b] = append(rules[b], [2]int{a, b})

			continue
		}

		// Converting line in numbers.
		splits := strings.Split(line, ",")
		numbers = numbers[:len(splits)]

		for i, s := range splits {
			n, errN := strconv.Atoi(s)
			check(errN)
			numbers[i] = n
		}

		// Is line in correct order ?
		checked, matched := checkLine(numbers, rules)
		ogLineInError := matched != -1

		// Naive approeach.. switching invalid values and checking again until we got a valid line.
		for matched != -1 {
			tmp := numbers[checked]
			numbers[checked] = numbers[matched]
			numbers[matched] = tmp
			checked, matched = checkLine(numbers, rules)
		}

		// If line was originally in error.
		if ogLineInError {
			middlesCount += numbers[len(numbers)/2]
		}
	}

	return middlesCount
}

func firstPart(lines []string) int {
	middlesCount := 0
	startSorting := false
	rules := make([][][2]int, 100)
	numbers := make([]int, 0, 99)

	for _, line := range lines {
		if line == "" {
			continue
		}

		if line[0] == 'X' {
			startSorting = true
			continue
		}

		// Building sorting rules.
		if !startSorting {
			a, errA := strconv.Atoi(string(line[0]) + string(line[1]))
			b, errB := strconv.Atoi(string(line[3]) + string(line[4]))
			check(errA)
			check(errB)

			if len(rules[a]) == 0 {
				rules[a] = make([][2]int, 0, 100)
			}
			if len(rules[b]) == 0 {
				rules[b] = make([][2]int, 0, 100)
			}

			rules[a] = append(rules[a], [2]int{a, b})
			rules[b] = append(rules[b], [2]int{a, b})

			continue
		}

		// Converting line in numbers.
		splits := strings.Split(line, ",")
		numbers = numbers[:len(splits)]

		for i, s := range splits {
			n, errN := strconv.Atoi(s)
			check(errN)
			numbers[i] = n
		}

		// Is line in correct order ?
		if _, matched := checkLine(numbers, rules); matched == -1 {
			middlesCount += numbers[len(numbers)/2]
		}
	}

	return middlesCount
}
func checkLine(numbers []int, rules [][][2]int) (int, int) {
	// Check each number on the line
	for i, n := range numbers {
		// Check each rule for this number
		for _, r := range rules[n] {
			// If number checked is in the left part (lower than)
			if r[0] == n {
				// Check for upper values
				for j := i; j >= 0; j-- {
					if numbers[j] == r[1] {
						return i, j
					}
				}
			} else {
				// Else if number checked is in the right part (greater than)
				// Check all numbers on the line before
				for j := i; j < len(numbers); j++ {
					if numbers[j] == r[0] {
						return i, j
					}
				}
			}
		}
	}

	return -1, -1
}

func main() {
	lines := getLines("./input.txt")
	start := time.Now()
	fmt.Printf("First part result: %v (%v)\n", firstPart(lines), time.Since(start))
	start = time.Now()
	fmt.Printf("Second part result: %v (%v)\n", secondPart(lines), time.Since(start))
}
