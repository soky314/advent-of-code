package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

const day = "04" // todo: add day of advent calendar

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
	data := ParserBoard(r)

	for _, number := range data.DrawNumbers {
		for index := range data.Boards {
			marked := data.Boards[index].Mark(number)
			if marked {
				if data.Boards[index].IsComplete() {
					return data.Boards[index].Score(number)
				}
			}
		}
	}
	return 0
}

func SolvePartTwo(r io.Reader) int {
	data := ParserBoard(r)
	var (
		winners = map[int]int{}
		order   = []int{}
	)

	for _, number := range data.DrawNumbers {
		for index := range data.Boards {
			marked := data.Boards[index].Mark(number)
			if marked {
				if data.Boards[index].IsComplete() {
					if _, exist := winners[index]; !exist {
						score := data.Boards[index].Score(number)
						winners[index] = score
						order = append(order, index)
					}
				}
			}
		}
	}
	return winners[order[len(order)-1]]
}
func ParserBoard(r io.Reader) (d Data) {
	s := bufio.NewScanner(r)
	s.Scan()
	fields := strings.Split(s.Text(), ",")
	d.DrawNumbers = make([]int, 0, len(fields))
	for _, field := range fields {
		n, err := strconv.Atoi(field)
		if err != nil {
			panic(fmt.Errorf("failed to parse number: %s, err: %w", field, err))
		}
		d.DrawNumbers = append(d.DrawNumbers, n)
	}
	s.Scan()
	var board Board
	for s.Scan() {
		rowText := s.Text()
		if len(rowText) == 0 {
			d.Boards = append(d.Boards, board)
			board = Board{Cells: make([][]int, 0, 5)}
			continue
		}
		fields = strings.Fields(rowText)
		row := make([]int, 0, 5)
		for _, field := range fields {
			n, err := strconv.Atoi(field)
			if err != nil {
				panic(fmt.Errorf("failed to parse number: %s, err: %w", field, err))
			}
			row = append(row, n)
		}
		board.Cells = append(board.Cells, row)
	}
	d.Boards = append(d.Boards, board)
	return d
}

type Data struct {
	DrawNumbers []int
	Boards      []Board
}
type Board struct {
	Cells    [][]int
	Selected []Selected
}

func (b *Board) Mark(number int) (marked bool) {
	for y := 0; y < 5; y++ {
		for x := 0; x < 5; x++ {
			if marked = number == b.Cells[y][x]; marked {
				b.Selected = append(b.Selected, Selected{
					X: x,
					Y: y,
				})
				return marked
			}
		}
	}
	return false
}

func (b *Board) IsComplete() bool {
	rowCol := map[int]int{}
	for _, value := range b.Selected {
		sum := rowCol[value.X] + 1
		if sum == 5 {
			return true
		}
		rowCol[value.X] = sum
		sum = rowCol[10+value.Y] + 1
		if sum == 5 {
			return true
		}
		rowCol[10+value.Y] = sum
	}
	return false
}

func (b *Board) Score(number int) (score int) {
	for y := 0; y < 5; y++ {
		for x := 0; x < 5; x++ {
			score += b.Cells[y][x]
		}
	}
	for _, value := range b.Selected {
		score -= b.Cells[value.Y][value.X]
	}
	return score * number
}

type Selected struct {
	X, Y int
}
