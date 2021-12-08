package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
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
	scan := bufio.NewScanner(r)

	var (
		index  int
		counts = map[int][2]int{}
	)
	for scan.Scan() {
		index++
		if err := scan.Err(); err != nil {
			if !errors.Is(err, io.EOF) {
				panic(fmt.Errorf("in reading line(%d), err: %w", index, err))
			}
			break
		}
		for i, value := range scan.Text() {
			positionCount := counts[i]
			switch value {
			case '0':
				positionCount[0]++
			case '1':
				positionCount[1]++
			default:
				panic(fmt.Errorf("in parsing line(%d), got: %s ", index, string(value)))
			}
			counts[i] = positionCount
		}
	}
	var (
		gamaBinary    = make([]rune, len(counts))
		epsilonBinary = make([]rune, len(counts))
	)

	for i, value := range counts {
		if value[1] > value[0] {
			gamaBinary[i] = '1'
			epsilonBinary[i] = '0'
		} else {
			gamaBinary[i] = '0'
			epsilonBinary[i] = '1'
		}
	}

	gama, err := strconv.ParseInt(string(gamaBinary), 2, 64)
	if err != nil {
		panic(fmt.Errorf("in parsing gama(%s), err: %w ", string(gamaBinary), err))

	}
	epsilon, err := strconv.ParseInt(string(epsilonBinary), 2, 64)
	if err != nil {
		panic(fmt.Errorf("in parsing epsilon(%s), err: %w ", string(epsilonBinary), err))

	}
	return int(gama * epsilon)
}

func SolvePartTwo(r io.Reader) int {
	return 0
}
