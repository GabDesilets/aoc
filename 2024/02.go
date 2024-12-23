package main

// 7 7 6 4 2 1
// 2 5 4 3 2
// 1 2 7 4 5
// 7 6 4 2 1
// 1 2 7 8 9
// 9 7 6 2 1
// 1 3 2 4 5
// 8 6 4 4 1
// 1 3 6 7 9
// 10 1 2 3 4

import (
	"fmt"
	"strconv"
	"strings"
)

func Day02(content string) {
	reports := strings.Split(content, "\n")
	safeCount := 0
	safeReports := []string{}
	unsafeReports := []string{}
	unsafeLevels := [][]int{}
	for _, report := range reports {
		lvls := strings.Split(report, " ")
		levels := levelsToInt(lvls)

		if isSafe(levels) {
			safeReports = append(safeReports, report)
			safeCount++
		} else {
			unsafeReports = append(unsafeReports, report)
			unsafeLevels = append(unsafeLevels, levels)
		}

	}

	fmt.Println("Safe reports:", safeCount)
	part2(unsafeLevels)
}

func isSafe(levels []int) bool {
	decrease := false
	increase := false
	safe := true

	for i, lvl := range levels {
		if i == len(levels)-1 {
			break
		}
		if lvl < levels[i+1] {
			increase = true
		} else if lvl > levels[i+1] {
			decrease = true
		} else {
			safe = false // same level
			break
		}

		if increase && decrease {
			safe = false
			break
		}

		if increase {
			change := levels[i+1] - lvl
			if !(change <= 3 && change > 0) {
				safe = false
				break
			}
		}

		if decrease {
			change := lvl - levels[i+1]
			if !(change <= 3 && change > 0) {
				safe = false
				break
			}
		}
	}
	return safe
}

func levelDirection(levels []int) string {
	directions := []string{}
	for i, lvl := range levels {
		if i == len(levels)-1 {
			break
		}

		if lvl < levels[i+1] {
			directions = append(directions, "i")
		} else if lvl > levels[i+1] {
			directions = append(directions, "d")
		} else {
			directions = append(directions, "s")
		}
	}
	return getMostFrequentValue(directions)

}

func getMostFrequentValue(arr []string) string {
	counts := make(map[string]int)
	for _, val := range arr {
		counts[val]++
	}

	if counts["i"] > counts["d"] {
		return "i"
	}
	return "d"
}

func part2(unsafeLevels [][]int) {
	//fmt.Println("Unsafe levels:", unsafeLevels)
	safeCount := 0
	for _, levels := range unsafeLevels {

		for i := 0; i < len(levels); i++ {
			newLevels := removeIndex(levels, i)

			if isSafe(newLevels) {
				safeCount++
				fmt.Println(newLevels)
				break
			} else {
				//fmt.Println("Unsafe levels:", levels)
			}
		}

	}
	fmt.Println("Safe reports:", safeCount)
}

func removeIndex(levels []int, index int) []int {
	newLevels := make([]int, 0, len(levels)-1)
	for i, lvl := range levels {
		if i != index {
			newLevels = append(newLevels, lvl)
		}
	}
	return newLevels
}

func levelsToInt(lvls []string) []int {
	levels := make([]int, len(lvls))
	for i, lvl := range lvls {
		num, _ := strconv.Atoi(lvl)
		levels[i] = num
	}
	return levels
}
