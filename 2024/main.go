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
	//Day02(string(content))
	Day03(string(content))
}
