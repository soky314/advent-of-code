package main

import (
	"fmt"
	"io"
	"os"
)

const day = "03" // todo: add day of advent calendar

func main() {
	file, err := os.Open(fmt.Sprintf("2021/%s/golang/input.txt", day))
	if err != nil {
		panic(fmt.Errorf("failed to open file, err: %w", err))
	}

	fmt.Printf("Solution of day: %s, part ONE: %d\n", day, SolvePartOne(file))
	_, err = file.Seek(0, 0)
	if err != nil {
		panic(fmt.Errorf("failed to rewind opened file, err: %w", err))
	}
	fmt.Printf("Solution of day: %s, part TWO: %d\n", day, SolvePartTwo(file))
}

func SolvePartOne(r io.Reader) int {
	return 0
}

func SolvePartTwo(r io.Reader) int {
	return 0
}
