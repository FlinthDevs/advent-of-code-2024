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

func firstPart(lines []string) int {
	safeCount := 0
	prev := -1
	increasing := false

	for _, line := range lines {
		numbers := getLinesAsInts(line)

		// Skip invalid lines
		if len(numbers) < 4 || numbers[0] == numbers[4] {
			continue
		}

		increasing = numbers[0]-numbers[4] < 0

		for i, n := range numbers {
			if i == 0 {
				prev = n
				continue
			}

			if increasing && n <= prev || n > prev+3 {
				prev = -1
				break
			} else if !increasing && n >= prev || n < prev-3 {
				prev = -1
				break
			}

			prev = n
		}

		if prev != -1 {
			safeCount++
		}

		prev = -1
	}

	return safeCount
}

func secondPart(lines []string) int {
	safeCount := 0
	prev := -1
	incErrs := 0
	increasing := 0
	decErrs := 0
	countInc := 0
	countDec := 0

	for _, line := range lines {
		if line == "" {
			continue
		}

		numbers := getLinesAsInts(line)
		// fmt.Println("Now it's ", numbers)

		for i, n := range numbers {
			if i == 0 {
				prev = n
				continue
			}

			// fmt.Printf("v: %v ; prev: %v; inc %v\n", n, prev, increasing)

			if n == prev {
				decErrs++
				incErrs++
			} else if n < prev-3 {
				decErrs++
			} else if n < prev && increasing == 1 {
				incErrs++
			} else if n > prev+3 {
				incErrs++
			} else if n > prev && increasing == -1 {
				decErrs++
			}

			// fmt.Printf("was %v (%v > %v+3 with %v)\n", incErrs, n, prev, increasing)
			if n > prev {
				increasing = 1
			} else {
				increasing = -1
			}

			// Increasing count of processed elements. Equalities count as both.
			if increasing == 0 {
				countInc++
				countDec++
			} else if increasing > 0 {
				countInc++
			} else {
				countDec++
			}

			// fmt.Println("current errors ", incErrs)
			prev = n
		}

		// If increase errors safe and moving up OR decrease errors safe and moving down.
		if incErrs < 2 && decErrs < 2 { //&& countInc > countDec || decErrs < 2 && countInc < countDec {
			safeCount++
		}

		fmt.Print(line, " - ", incErrs, decErrs, " / ", countInc, countDec)
		if incErrs < 2 && decErrs < 2 { //&& countInc > countDec || decErrs < 2 && countInc < countDec {
			fmt.Println(" (count)")
		} else {
			fmt.Println(" (NOT)")
		}

		incErrs = 0
		decErrs = 0
		countInc = 0
		countDec = 0
		increasing = 0
	}

	return safeCount
}

func main() {
	lines := getLines("./input.txt")

	start := time.Now()
	fmt.Printf("First part result: %v (%v)\n", firstPart(lines), time.Since(start))

	// Not 765
	// Not 571
	// Not 557
	// Try 566
	start = time.Now()
	fmt.Printf("Second part result: %v (%v)\n", secondPart(lines), time.Since(start))
}
