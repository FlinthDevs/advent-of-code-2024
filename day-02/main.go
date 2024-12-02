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

func getLinesAsInts(line string) []int {
	splits := strings.Split(line, " ")
	numbers := make([]int, len(splits), len(splits))

	for i, number := range splits {
		if number == "" {
			continue
		}

		n, err := strconv.Atoi(number)
		check(err)

		numbers[i] = n
	}

	return numbers
}

// Check for an int array if everything is going straight down or up without an error.
// Second param is for ignoring a specific index.
func checkLine(numbers []int, ignoredIndex int) bool {
	prev := -1
	increasing := 0

	for i, n := range numbers {
		// If ignoredIndex = -1 we ignore nothing, else we skip.
		if ignoredIndex >= 0 && i == ignoredIndex {
			continue
		}

		// First actual value, cannot compare to store and go next.
		if prev == -1 {
			prev = n
			continue
		}

		// All error cases.
		if n == prev ||
			n > prev+3 ||
			n < prev-3 ||
			n > prev && increasing < 0 ||
			n < prev && increasing > 0 {
			return false
		}

		// Pfew, no errors. Checking out if we're going up or down with this array.
		if n > prev {
			increasing = 1
		} else {
			increasing = -1
		}

		prev = n
	}

	return true

}

func firstPart(lines []string) int {
	safeCount := 0

	// Go through each line and check for all valid ones.
	for _, line := range lines {
		if line == "" {
			continue
		}

		if checkLine(getLinesAsInts(line), -1) {
			safeCount++
		}
	}

	return safeCount
}

func secondPart(lines []string) int {
	safeCount := 0
	failed := false

	// Go through each line and check for there's a version for each where by removing a value we are valid.
	for _, line := range lines {
		if line == "" {
			continue
		}

		numbers := getLinesAsInts(line)

		// First check: No ignored value.
		failed = !checkLine(numbers, -1)

		// If full line is wrong, ignore each value one by one to check if it can be valid.
		if failed {
			// Going through each array value and checking the line without it.
			for i := range numbers {
				failed = !checkLine(numbers, i)

				if !failed {
					break
				}
			}
		}

		// If it didn't fail with at least a combination, we can call that a success I guess.
		if !failed {
			safeCount++
		}
	}

	return safeCount
}

func main() {
	lines := getLines("./input.txt")

	start := time.Now()
	fmt.Printf("First part result: %v (%v)\n", firstPart(lines), time.Since(start))
	start = time.Now()
	fmt.Printf("Second part result: %v (%v)\n", secondPart(lines), time.Since(start))
}
