package main

import (
	"io"
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
		{
			name: "origin example",
			args: args{
				r: strings.NewReader("forward 5\ndown 5\nforward 8\nup 3\ndown 8\nforward 2"),
			},
			want: 150,
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
		{
			name: "origin example",
			args: args{
				r: strings.NewReader("forward 5\ndown 5\nforward 8\nup 3\ndown 8\nforward 2"), // TODO: add origin example input
			},
			want: 900, // TODO: add origin example input
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
