package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
)

const MaxInt = int((^uint(0)) >> 1)

func main() {
	file, err := os.Open("data/input.txt")
	if err != nil {
		panic(fmt.Errorf("failed to open file, err: %w", err))
	}

	fmt.Println("Number of increase: ", CountNumberOfIncrease(file))
}

func CountNumberOfIncrease(r io.Reader) (count int) {
	scan := bufio.NewScanner(r)

	var (
		index int
		last  = MaxInt
	)
	for scan.Scan() {
		index++
		if err := scan.Err(); err != nil {
			if !errors.Is(err, io.EOF) {
				panic(fmt.Errorf("in reading line(%d), err: %w", index, err))
			}
			break
		}
		actual, err := strconv.Atoi(scan.Text())
		if err != nil {
			panic(fmt.Errorf("in parsing line(%d), err: %w", index, err))
		}
		if actual > last {
			count++
		}
		last = actual
	}

	return count
}

func CountNumberOfWindowIncrease(r io.Reader) (count int) {
	return 0
}
