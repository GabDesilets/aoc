package main

import (
	"log"
	"os"
)

func main() {
	testMode := true
	if len(os.Args) < 2 {
		testMode = false
	}

	var content []byte
	var err error

	if testMode {
		content, err = os.ReadFile("_input.txt")
	} else {
		content, err = os.ReadFile("input.txt")
	}

	if err != nil {
		log.Fatalf("Failed to read file: %v", err)
	}
	//Day01(string(content))
	Day02(string(content))
}

// 7 7 6 4 2 1
// 2 5 4 3 2
// 1 2 7 4 5
// 7 6 4 2 1
// 1 2 7 8 9
// 9 7 6 2 1
// 1 3 2 4 5
// 8 6 4 4 1
// 1 3 6 7 9
// 10 1 2 3 4 5
