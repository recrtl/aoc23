package day1_test

import (
	_ "embed"
	"github.com/recrtl/aoc23/day1"
	"github.com/stretchr/testify/assert"
	"testing"
)

//go:embed input.txt
var input string

func Test_day11(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{
			name:  "example",
			input: "1abc2\npqr3stu8vwx\na1b2c3d4e5f\ntreb7uchet",
			want:  "142",
		},
		{
			name:  "example",
			input: input,
			want:  "56397",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := day1.Part1(tt.input)
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_day12(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{
			name:  "example",
			input: "two1nine\neightwothree\nabcone2threexyz\nxtwone3four\n4nineeightseven2\nzoneight234\n7pqrstsixteen",
			want:  "281",
		},
		{
			name:  "three",
			input: "three",
			want:  "33",
		},
		{
			name:  "oneight",
			input: "oneight",
			want:  "18",
		},
		{
			name:  "eightwo",
			input: "eightwo",
			want:  "82",
		},
		{
			name:  "1eightwo",
			input: "1eightwo",
			want:  "12",
		},
		{
			name:  "twone",
			input: "twone",
			want:  "21",
		},
		{
			name:  "eighthree",
			input: "eighthree",
			want:  "83",
		},
		{
			name:  "example",
			input: input,
			want:  "55701",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := day1.Part2(tt.input)
			assert.Equal(t, tt.want, got)
		})
	}
}
