package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

const day = "05" // todo: add day of advent calendar

func main() {
	file, err := os.Open(fmt.Sprintf("2021/%s/input.txt", day))
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
	lines := ParseInput(r)
	_ = lines
	ventsField := make(map[int]int, 999*999)
	dangerMap := map[int]struct{}{}
	wantDanger := 2
	for i, line := range lines {
		if !line.IsStraight() {
			fmt.Printf("Line: %d: From: (%d, %d), To:(%d,%d) is not straight\n", i, line.From.X, line.From.Y, line.To.X, line.To.Y)
			continue
		}
		for _, point := range line.Points() {
			index := point.Index()
			danger := ventsField[index] + 1
			ventsField[index] = danger
			if danger >= wantDanger {
				dangerMap[index] = struct{}{}
			}
		}
	}

	return len(dangerMap)
}
func SolvePartTwo(r io.Reader) int {
	lines := ParseInput(r)
	_ = lines
	ventsField := make(map[int]int, 999*999)
	dangerMap := map[int]struct{}{}
	wantDanger := 2
	for _, line := range lines {
		for _, point := range line.Points() {
			index := point.Index()
			danger := ventsField[index] + 1
			ventsField[index] = danger
			if danger >= wantDanger {
				dangerMap[index] = struct{}{}
			}
		}
	}

	return len(dangerMap)
}

func ParseInput(r io.Reader) (lines []Line) {
	s := bufio.NewScanner(r)
	for s.Scan() {
		line := NewLine(s.Text())
		lines = append(lines, line)
	}
	return lines
}

type Line struct {
	From, To Point
}

func NewLine(s string) (l Line) {
	fields := strings.Split(s, " -> ")
	l.From = NewPoint(fields[0])
	l.To = NewPoint(fields[1])
	return l
}
func (l Line) IsStraight() bool {
	return l.From.X == l.To.X || l.From.Y == l.To.Y
}
func (l Line) DirectionX() int {
	if l.From.X-l.To.X > 0 {
		return -1
	}
	if l.From.X-l.To.X < 0 {
		return 1
	}
	return 0
}
func (l Line) DirectionY() int {
	if l.From.Y-l.To.Y > 0 {
		return -1
	}
	if l.From.Y-l.To.Y < 0 {
		return 1
	}
	return 0
}
func (l Line) Points() (points []Point) {
	p := l.From
	dirX := l.DirectionX()
	dirY := l.DirectionY()
	points = append(points, p)
	for p.X != l.To.X || p.Y != l.To.Y {
		p.X += dirX
		p.Y += dirY
		points = append(points, p)
	}

	return points
}

type Point struct {
	X, Y int
}

func NewPoint(s string) (p Point) {
	fields := strings.Split(s, ",")
	var err error
	p.X, err = strconv.Atoi(fields[0])
	if err != nil {
		panic(fmt.Errorf("in parsing: %s, err: %w", fields[0], err))
	}
	p.Y, err = strconv.Atoi(fields[1])
	if err != nil {
		panic(fmt.Errorf("in parsing: %s, err: %w", fields[0], err))
	}
	return p
}
func (p Point) Index() int {
	return 999*p.Y + p.X
}
