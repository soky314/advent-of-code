package main

import (
	"io"
	"reflect"
	"strings"
	"testing"
)

func TestSolvePartOne(t *testing.T) {
	type args struct {
		r io.Reader
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			args: args{r: strings.NewReader(testData)},
			want: 4512,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SolvePartOne(tt.args.r); got != tt.want {
				t.Errorf("SolvePartOne() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSolvePartTwo(t *testing.T) {
	type args struct {
		r io.Reader
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SolvePartTwo(tt.args.r); got != tt.want {
				t.Errorf("SolvePartTwo() = %v, want %v", got, tt.want)
			}
		})
	}
}

const testData = `7,4,9,5,11,17,23,2,0,14,21,24,10,16,13,6,15,25,12,22,18,20,8,19,3,26,1

22 13 17 11  0
 8  2 23  4 24
21  9 14 16  7
 6 10  3 18  5
 1 12 20 15 19

 3 15  0  2 22
 9 18 13 17  5
19  8  7 25 23
20 11 10 24  4
14 21 16 12  6

14 21 17 24  4
10 16 15  9 19
18  8 23 26 20
22 11 13  6  5
 2  0 12  3  7`

func TestParserBoard(t *testing.T) {
	type args struct {
		r io.Reader
	}
	tests := []struct {
		name  string
		args  args
		wantD Data
	}{
		// TODO: Add test cases.
		{
			args: args{r: strings.NewReader(testData)},
			wantD: Data{
				DrawNumbers: []int{7, 4, 9, 5, 11, 17, 23, 2, 0, 14, 21, 24, 10, 16, 13, 6, 15, 25, 12, 22, 18, 20, 8, 19, 3, 26, 1},
				Boards: []Board{
					{
						Cells: [][]int{
							{22, 13, 17, 11, 0},
							{8, 2, 23, 4, 24},
							{21, 9, 14, 16, 7},
							{6, 10, 3, 18, 5},
							{1, 12, 20, 15, 19},
						},
					},
					{
						Cells: [][]int{
							{3, 15, 0, 2, 22},
							{9, 18, 13, 17, 5},
							{19, 8, 7, 25, 23},
							{20, 11, 10, 24, 4},
							{14, 21, 16, 12, 6},
						},
					},
					{
						Cells: [][]int{
							{14, 21, 17, 24, 4},
							{10, 16, 15, 9, 19},
							{18, 8, 23, 26, 20},
							{22, 11, 13, 6, 5},
							{2, 0, 12, 3, 7},
						},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotD := ParserBoard(tt.args.r); !reflect.DeepEqual(gotD, tt.wantD) {
				t.Errorf("ParserBoard() = %v, want %v", gotD, tt.wantD)
			}
		})
	}
}

func TestBoard_IsComplete(t *testing.T) {
	type fields struct {
		Cells    [][]int
		Selected []Selected
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		// TODO: Add test cases.
		{
			name: "empty",
			fields: fields{
				Cells: [][]int{
					{22, 13, 17, 11, 0},
					{8, 2, 23, 4, 24},
					{21, 9, 14, 16, 7},
					{6, 10, 3, 18, 5},
					{1, 12, 20, 15, 19},
				},
				Selected: []Selected{},
			},
			want: false,
		},
		{
			name: "one",
			fields: fields{
				Cells: [][]int{
					{22, 13, 17, 11, 0},
					{8, 2, 23, 4, 24},
					{21, 9, 14, 16, 7},
					{6, 10, 3, 18, 5},
					{1, 12, 20, 15, 19},
				},
				Selected: []Selected{
					{0, 0},
				},
			},
			want: false,
		},
		{
			name: "five - diagonal",
			fields: fields{
				Cells: [][]int{
					{22, 13, 17, 11, 0},
					{8, 2, 23, 4, 24},
					{21, 9, 14, 16, 7},
					{6, 10, 3, 18, 5},
					{1, 12, 20, 15, 19},
				},
				Selected: []Selected{
					{0, 0}, {1, 1}, {2, 2}, {3, 3}, {4, 4},
				},
			},
			want: false,
		},
		{
			name: "five - column",
			fields: fields{
				Cells: [][]int{
					{22, 13, 17, 11, 0},
					{8, 2, 23, 4, 24},
					{21, 9, 14, 16, 7},
					{6, 10, 3, 18, 5},
					{1, 12, 20, 15, 19},
				},
				Selected: []Selected{
					{0, 0}, {0, 1}, {0, 2}, {0, 3}, {0, 4},
				},
			},
			want: true,
		},
		{
			name: "five - row",
			fields: fields{
				Cells: [][]int{
					{22, 13, 17, 11, 0},
					{8, 2, 23, 4, 24},
					{21, 9, 14, 16, 7},
					{6, 10, 3, 18, 5},
					{1, 12, 20, 15, 19},
				},
				Selected: []Selected{
					{0, 0}, {1, 0}, {2, 0}, {3, 0}, {4, 0},
				},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &Board{
				Cells:    tt.fields.Cells,
				Selected: tt.fields.Selected,
			}
			if got := b.IsComplete(); got != tt.want {
				t.Errorf("IsComplete() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBoard_Score(t *testing.T) {
	type fields struct {
		Cells    [][]int
		Selected []Selected
	}
	type args struct {
		number int
	}
	tests := []struct {
		name      string
		fields    fields
		args      args
		wantScore int
	}{
		// TODO: Add test cases.
		{
			name: "sum",
			fields: fields{
				Cells: [][]int{
					{14, 21, 17, 24, 4},
					{10, 16, 15, 9, 19},
					{18, 8, 23, 26, 20},
					{22, 11, 13, 6, 5},
					{2, 0, 12, 3, 7},
				},
				Selected: nil,
			},
			args: args{
				number: 1,
			},
			wantScore: 325,
		},
		{
			name: "sum",
			fields: fields{
				Cells: [][]int{
					{14, 21, 17, 24, 4},
					{10, 16, 15, 9, 19},
					{18, 8, 23, 26, 20},
					{22, 11, 13, 6, 5},
					{2, 0, 12, 3, 7},
				},
				Selected: nil,
			},
			args: args{
				number: 2,
			},
			wantScore: 650,
		},
		{
			name: "sum",
			fields: fields{
				Cells: [][]int{
					{14, 21, 17, 24, 4},
					{10, 16, 15, 9, 19},
					{18, 8, 23, 26, 20},
					{22, 11, 13, 6, 5},
					{2, 0, 12, 3, 7},
				},
				Selected: []Selected{
					{0, 0}, {1, 0}, {2, 0}, {3, 0}, {4, 0},
					{3, 1},
					{2, 2},
					{1, 3}, {4, 3},
					{0, 4}, {1, 4}, {4, 4},
				},
			},
			args: args{
				number: 1,
			},
			wantScore: 188,
		},
		{
			name: "sum",
			fields: fields{
				Cells: [][]int{
					{14, 21, 17, 24, 4},
					{10, 16, 15, 9, 19},
					{18, 8, 23, 26, 20},
					{22, 11, 13, 6, 5},
					{2, 0, 12, 3, 7},
				},
				Selected: []Selected{
					{0, 0}, {1, 0}, {2, 0}, {3, 0}, {4, 0},
					{3, 1},
					{2, 2},
					{1, 3}, {4, 3},
					{0, 4}, {1, 4}, {4, 4},
				},
			},
			args: args{
				number: 24,
			},
			wantScore: 4512,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &Board{
				Cells:    tt.fields.Cells,
				Selected: tt.fields.Selected,
			}
			if gotScore := b.Score(tt.args.number); gotScore != tt.wantScore {
				t.Errorf("Score() = %v, want %v", gotScore, tt.wantScore)
			}
		})
	}
}
