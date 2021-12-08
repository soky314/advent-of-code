package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(fmt.Errorf("failed to open file, err: %w", err))
	}

	fmt.Println("Number of increase: ", SolvePartOne(file))
	_, err = file.Seek(0, 0)
	if err != nil {
		panic(fmt.Errorf("failed to rewind opened file, err: %w", err))
	}
	fmt.Println("Number of window increase: ", SolvePartTwo(file))
}

func SolvePartOne(r io.Reader) int {
	return 0
}

func SolvePartTwo(r io.Reader) int {
	return 0
}
