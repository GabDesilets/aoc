package main

import (
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
)

func Day01(content string) {

	temp := strings.Split(content, "\n")

	left := []int{}
	right := []int{}
	rightOccurences := make(map[int]int)
	for _, v := range temp {
		if v != "" {
			parts := strings.Split(v, "   ")
			leftNum, _ := strconv.Atoi(parts[0])
			rightNum, _ := strconv.Atoi(parts[1])

			left = append(left, leftNum)
			right = append(right, rightNum)

			rightOccurences[rightNum]++
		}
	}

	sort.Ints(left)
	sort.Ints(right)

	distance := 0
	for i := 0; i < len(left); i++ {
		distance += int(math.Abs(float64(left[i] - right[i])))
	}

	similarity := 0
	for i := 0; i < len(left); i++ {
		if count, exists := rightOccurences[left[i]]; exists {
			similarity += left[i] * count
		}
	}

	fmt.Println("Distance:", distance)
	fmt.Println("similarity:", similarity)
}
