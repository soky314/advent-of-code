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
	_, err = file.Seek(0, 0)
	if err != nil {
		panic(fmt.Errorf("failed to rewind opened file, err: %w", err))
	}
	fmt.Println("Number of window increase: ", CountNumberOfWindowIncrease(file))
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
	scan := bufio.NewScanner(r)

	var (
		index  int
		last   = MaxInt
		window = [3]int{0, 0, 0}
	)
	sumWindow := func() int {
		return window[0] + window[1] + window[2]
	}
	for scan.Scan() {
		index++
		if err := scan.Err(); err != nil {
			if !errors.Is(err, io.EOF) {
				panic(fmt.Errorf("in reading line(%d), err: %w", index, err))
			}
			break
		}
		tmp, err := strconv.Atoi(scan.Text())
		if err != nil {
			panic(fmt.Errorf("in parsing line(%d), err: %w", index, err))
		}
		window[(index-1)%3] = tmp
		if index < 3 {
			continue
		}
		actual := sumWindow()
		if actual > last {
			count++
		}
		last = actual
	}

	return count
}
