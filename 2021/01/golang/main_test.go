package main

import (
	"io"
	"strings"
	"testing"
)

func TestCountNumberOfIncrease(t *testing.T) {
	type args struct {
		r io.Reader
	}
	tests := []struct {
		name      string
		args      args
		wantCount int
	}{
		// TODO: Add test cases.
		{
			name: "origin example",
			args: args{
				r: strings.NewReader("199\n200\n208\n210\n200\n207\n240\n269\n260\n263"),
			},
			wantCount: 7,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotCount := CountNumberOfIncrease(tt.args.r); gotCount != tt.wantCount {
				t.Errorf("CountNumberOfIncrease() = %v, want %v", gotCount, tt.wantCount)
			}
		})
	}
}

func TestCountNumberOfWindowIncrease(t *testing.T) {
	type args struct {
		r io.Reader
	}
	tests := []struct {
		name      string
		args      args
		wantCount int
	}{
		// TODO: Add test cases.
		{
			name: "origin example",
			args: args{
				r: strings.NewReader("199\n200\n208\n210\n200\n207\n240\n269\n260\n263"),
			},
			wantCount: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotCount := CountNumberOfWindowIncrease(tt.args.r); gotCount != tt.wantCount {
				t.Errorf("CountNumberOfWindowIncrease() = %v, want %v", gotCount, tt.wantCount)
			}
		})
	}
}
