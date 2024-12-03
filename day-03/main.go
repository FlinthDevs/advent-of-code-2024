package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"time"
)

var mulRegex = regexp.MustCompile(`mul\([0-9][0-9]?[0-9]?,[0-9][0-9]?[0-9]?\)`)
var numberRegex = regexp.MustCompile(`[0-9][0-9]?[0-9]?`)
var doRegex = regexp.MustCompile(`do\(\)`)
var dontRegex = regexp.MustCompile(`don\'t\(\)`)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func getInputData(filename string) string {
	data, err := os.ReadFile(filename)
	check(err)
	return string(data)
}

func compileMul(word string) int {
	// If this "mul(X[XX],X[XX])" ?
	mulMatch := mulRegex.FindStringSubmatch(word)

	if len(mulMatch) != 1 {
		return 0
	}

	// Extracts numbers for calculus.
	// Don't double check regex nMatches indexes because we know they're right from the previous one.
	nMatches := numberRegex.FindAllStringSubmatch(mulMatch[0], -1)

	x, xErr := strconv.Atoi(nMatches[0][0])
	y, yErr := strconv.Atoi(nMatches[1][0])
	check(xErr)
	check(yErr)

	return x * y
}

func compileDo(word string) (bool, bool) {
	// Do we have a "do()" ?
	if doMatch := doRegex.FindStringSubmatch(word); len(doMatch) == 1 {
		return true, true
	}

	// Do we have a "don't()" ?
	if dontMatch := dontRegex.FindStringSubmatch(word); len(dontMatch) == 1 {
		return false, true
	}

	return false, false
}

func firstPart(inputData string) int {
	result := 0
	startRegistering := false
	currentWord := ""

	for _, c := range inputData {
		// Found the start of 'mul', start storing chars
		// and resetting if we stumble upon another one.
		if c == 'm' {
			currentWord = ""
			startRegistering = true
		}

		if startRegistering {
			currentWord += string(c)
		}

		// Found the end of 'mul(XXX,XXX)', do the calculus.
		if c == ')' {
			result += compileMul(currentWord)
			startRegistering = false
			currentWord = ""
		}
	}

	return result
}

func secondPart(inputData string) int {
	result := 0
	startRegistering := false
	isDo := true
	currentWord := ""

	for _, c := range inputData {
		// Found the start of 'mul' or 'do'
		if c == 'm' || c == 'd' {
			currentWord = ""
			startRegistering = true
		}

		if startRegistering {
			currentWord += string(c)
		}

		// Found a closing one, check what we have in store.
		if c == ')' {
			startRegistering = false

			// If we should do calculus start by that.
			if isDo {
				mulValue := compileMul(currentWord)

				// We got a value from 'mul' parser !
				// Add it, reset and move to next char.
				if mulValue > 0 {
					result += mulValue
					currentWord = ""
					continue
				}
			}

			// Check parsing results:
			// - shouldDo : Theorical value of isDo
			// - doFound: If false, means no actual match was done.
			shouldDo, doFound := compileDo(currentWord)

			// Only update isDo if we go a meaningful match
			if doFound {
				isDo = shouldDo
			}

			// Start over
			currentWord = ""
		}
	}

	return result
}

func main() {
	lines := getInputData("./input.txt")

	start := time.Now()
	fmt.Printf("First part result: %v (%v)\n", firstPart(lines), time.Since(start))
	start = time.Now()
	fmt.Printf("Second part result: %v (%v)\n", secondPart(lines), time.Since(start))
}
