package main

import (
	"io"
	"reflect"
	"strings"
	"testing"
)

const testData = `0,9 -> 5,9
8,0 -> 0,8
9,4 -> 3,4
2,2 -> 2,1
7,0 -> 7,4
6,4 -> 2,0
0,9 -> 2,9
3,4 -> 1,4
0,0 -> 8,8
5,5 -> 8,2`

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
			want: 5,
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
		{
			args: args{r: strings.NewReader(testData)},
			want: 12,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SolvePartTwo(tt.args.r); got != tt.want {
				t.Errorf("SolvePartTwo() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParseInput(t *testing.T) {
	type args struct {
		r io.Reader
	}
	tests := []struct {
		name      string
		args      args
		wantLines []Line
	}{
		// TODO: Add test cases.
		{
			args: args{r: strings.NewReader(testData)},
			wantLines: []Line{
				{From: Point{0, 9}, To: Point{5, 9}},
				{From: Point{8, 0}, To: Point{0, 8}},
				{From: Point{9, 4}, To: Point{3, 4}},
				{From: Point{2, 2}, To: Point{2, 1}},
				{From: Point{7, 0}, To: Point{7, 4}},
				{From: Point{6, 4}, To: Point{2, 0}},
				{From: Point{0, 9}, To: Point{2, 9}},
				{From: Point{3, 4}, To: Point{1, 4}},
				{From: Point{0, 0}, To: Point{8, 8}},
				{From: Point{5, 5}, To: Point{8, 2}},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotLines := ParseInput(tt.args.r); !reflect.DeepEqual(gotLines, tt.wantLines) {
				t.Errorf("ParseInput() = %v, want %v", gotLines, tt.wantLines)
			}
		})
	}
}

func TestLine_Points(t *testing.T) {
	type fields struct {
		From Point
		To   Point
	}
	tests := []struct {
		name       string
		fields     fields
		wantPoints []Point
	}{
		// TODO: Add test cases.
		{
			fields: fields{
				From: Point{0, 0},
				To:   Point{1, 0},
			},
			wantPoints: []Point{
				{0, 0}, {1, 0},
			},
		},
		{
			fields: fields{
				From: Point{0, 0},
				To:   Point{0, 1},
			},
			wantPoints: []Point{
				{0, 0}, {0, 1},
			},
		},
		{
			fields: fields{
				From: Point{0, 0},
				To:   Point{1, 1},
			},
			wantPoints: []Point{
				{0, 0}, {1, 1},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := Line{
				From: tt.fields.From,
				To:   tt.fields.To,
			}
			if gotPoints := l.Points(); !reflect.DeepEqual(gotPoints, tt.wantPoints) {
				t.Errorf("Points() = %v, want %v", gotPoints, tt.wantPoints)
			}
		})
	}
}
