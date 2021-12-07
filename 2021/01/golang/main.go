package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	file, err := os.Open("/data/input.txt")
	if err != nil {
		panic(fmt.Errorf("failed to open file, err: %w", err))
	}

	fmt.Println("Number of increase: ", CountNumberOfIncrease(file))
}

func CountNumberOfIncrease(r io.Reader) (count int) {
	return 0
}
