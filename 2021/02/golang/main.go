package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
)

func main() {
	file, err := os.Open("2021/02/golang/input.txt")
	if err != nil {
		panic(fmt.Errorf("failed to open file, err: %w", err))
	}

	fmt.Println("Solution od Part ONE: ", SolvePartOne(file))
	_, err = file.Seek(0, 0)
	if err != nil {
		panic(fmt.Errorf("failed to rewind opened file, err: %w", err))
	}
	fmt.Println("Solution od Part TWO: ", SolvePartTwo(file))
}

func SolvePartOne(r io.Reader) int {
	var (
		index    int
		position Position
	)

	scan := bufio.NewScanner(r)
	for scan.Scan() {
		if err := scan.Err(); err != nil {
			if !errors.Is(err, io.EOF) {
				panic(fmt.Errorf("in reading line(%d), err: %w", index, err))
			}
			break
		}
		position.Move(scan.Text())
	}

	return position.X * position.Y
}

func SolvePartTwo(r io.Reader) int {
	return 0
}

type Position struct {
	X, Y int
}

var commandRegex = regexp.MustCompile("([a-zA-Z]+) ([0-9]+)")

func (p *Position) Move(cmd string) {
	parts := commandRegex.FindStringSubmatch(cmd)
	if len(parts) != 3 {
		panic(fmt.Errorf("failed parse command: '%s', got: %v", cmd, parts))
	}
	steps, err := strconv.Atoi(parts[2])
	if err != nil {
		panic(fmt.Errorf("failed parse steps: '%s', err: %w", parts[2], err))
	}

	switch parts[1] {
	case "forward":
		p.X += steps
	case "down":
		p.Y += steps
	case "up":
		p.Y -= steps
	default:
		panic(fmt.Errorf("unknown direction: '%s'", cmd))
	}
}
