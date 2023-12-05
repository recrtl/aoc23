package day3_test

import (
	_ "embed"
	"github.com/recrtl/aoc23/day3"
	"github.com/stretchr/testify/assert"
	"testing"
)

//go:embed sample.txt
var sample string

//go:embed input.txt
var input string

func Test_day31(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{
			name:  "sample",
			input: sample,
			want:  "4361",
		},
		{
			name:  "input",
			input: input,
			want:  "530849",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := day3.Part1(tt.input)
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_day32(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{
			name:  "sample",
			input: sample,
			want:  "467835",
		},
		{
			name:  "input",
			input: input,
			want:  "84900879",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := day3.Part2(tt.input)
			assert.Equal(t, tt.want, got)
		})
	}
}
