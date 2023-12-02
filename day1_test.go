package aoc23_test

import (
	"embed"
	"github.com/recrtl/aoc23"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

//go:embed inputs/*
var content embed.FS

func mustInput(name string) string {
	f, err := content.ReadFile("inputs/" + name + ".txt")
	if err != nil {
		log.Fatal(err)
	}
	return string(f)
}

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
			input: mustInput("1"),
			want:  "56397",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := aoc23.Day11(tt.input)
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
			input: mustInput("1"),
			want:  "54571",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := aoc23.Day12(tt.input)
			assert.Equal(t, tt.want, got)
		})
	}
}
