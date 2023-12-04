package aoc23_test

import (
	"github.com/recrtl/aoc23/day4"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_day41(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{
			name:  "sample",
			input: mustInput("sample4"),
			want:  "13",
		},
		{
			name:  "input",
			input: mustInput("input4"),
			want:  "25651",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := day4.Part1(tt.input)
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_day42(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{
			name:  "sample",
			input: mustInput("sample4"),
			want:  "30",
		},
		{
			name:  "input",
			input: mustInput("input4"),
			want:  "19499881",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := day4.Part2(tt.input)
			assert.Equal(t, tt.want, got)
		})
	}
}
