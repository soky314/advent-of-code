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
	scan := bufio.NewScanner(r)

	var (
		index   int
		counts  = map[int][2]int{}
		numbers = [2][]string{
			make([]string, 0, 1000),
			make([]string, 0, 1000),
		}
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
				if i == 0 {
					numbers[0] = append(numbers[0], scan.Text())
				}
			case '1':
				positionCount[1]++
				if i == 0 {
					numbers[1] = append(numbers[1], scan.Text())
				}
			default:
				panic(fmt.Errorf("in parsing line(%d), got: %s ", index, string(value)))
			}
			counts[i] = positionCount
		}
	}
	var (
		oxygenRatingBinary, co2RatingBinary string
	)

	if len(numbers[0]) > len(numbers[1]) {
		oxygenRatingBinary = find(1, numbers[0], true)
		co2RatingBinary = find(1, numbers[1], false)
	} else {
		oxygenRatingBinary = find(1, numbers[1], true)
		co2RatingBinary = find(1, numbers[0], false)
	}

	oxygen, err := strconv.ParseInt(string(oxygenRatingBinary), 2, 64)
	if err != nil {
		panic(fmt.Errorf("in parsing oxygen scrubber rating (%s), err: %w ", string(oxygenRatingBinary), err))

	}
	co2, err := strconv.ParseInt(string(co2RatingBinary), 2, 64)
	if err != nil {
		panic(fmt.Errorf("in parsing co2scrubber rating(%s), err: %w ", string(co2RatingBinary), err))

	}
	return int(oxygen * co2)
}

func find(index int, numbers []string, fromBiggerGroup bool) string {
	var (
		subset = [2][]string{
			make([]string, 0, len(numbers)),
			make([]string, 0, len(numbers)),
		}
		counts = map[int][2]int{}
	)
	for _, value := range numbers {
		for i, val := range value[index:] {
			positionCount := counts[i]
			switch val {
			case '0':
				positionCount[0]++
				if i == 0 {
					subset[0] = append(subset[0], value)
				}
			case '1':
				positionCount[1]++
				if i == 0 {
					subset[1] = append(subset[1], value)
				}
			default:
				panic(fmt.Errorf("in parsing line(%d), got: %s ", index, string(value)))
			}
			counts[i] = positionCount
		}
	}

	switch {
	case fromBiggerGroup && (len(subset[0]) > len(subset[1])):
		numbers = subset[0]
	case fromBiggerGroup && (len(subset[0]) <= len(subset[0])):
		numbers = subset[1]
	case !fromBiggerGroup && (len(subset[0]) > len(subset[1])):
		numbers = subset[1]
	case !fromBiggerGroup && (len(subset[0]) <= len(subset[1])):
		numbers = subset[0]
	}

	if len(numbers) == 1 {
		return numbers[0]
	}

	return find(index+1, numbers, fromBiggerGroup)
}
