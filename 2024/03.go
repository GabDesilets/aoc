package main

import (
	"fmt"
	"regexp"
	"strconv"
)

func Day03(content string) {
	// Define the regular expression to match valid mul instructions
	re := regexp.MustCompile(`mul\((\d+),(\d+)\)`)

	// Find all matches
	matches := re.FindAllStringSubmatch(content, -1)

	total := 0
	for _, match := range matches {
		// Convert the captured groups to integers
		a, _ := strconv.Atoi(match[1])
		b, _ := strconv.Atoi(match[2])

		// Add the result of the multiplication to the total
		total += a * b
	}

	fmt.Println(total)
	d3part2(content)
}

func d3part2(content string) {
	// Define the regular expressions to match valid instructions
	mulRe := regexp.MustCompile(`mul\((\d+),(\d+)\)`)
	doRe := regexp.MustCompile(`do\(\)`)
	dontRe := regexp.MustCompile(`don't\(\)`)

	// Split the content into individual instructions
	instructions := regexp.MustCompile(`mul\(\d+,\d+\)|do\(\)|don't\(\)`).FindAllString(content, -1)

	// Initialize the state to enable mul instructions
	mulEnabled := true
	total := 0

	for _, instruction := range instructions {
		if doRe.MatchString(instruction) {
			mulEnabled = true
		} else if dontRe.MatchString(instruction) {
			mulEnabled = false
		} else if mulEnabled && mulRe.MatchString(instruction) {
			matches := mulRe.FindStringSubmatch(instruction)
			if len(matches) == 3 {
				// Convert the captured groups to integers
				a, _ := strconv.Atoi(matches[1])
				b, _ := strconv.Atoi(matches[2])
				total += a * b
			}
		}
	}

	fmt.Println("Total:", total)
}
