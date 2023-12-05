package day

import (
	_ "embed"
	"github.com/stretchr/testify/assert"
	"testing"
)

//go:embed sample.txt
var sample string

//go:embed input.txt
var input string

func Test_day41(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{
			name:  "sample",
			input: sample,
			want:  "13",
		},
		{
			name:  "input",
			input: input,
			want:  "25651",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Part1(tt.input)
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
			input: sample,
			want:  "30",
		},
		{
			name:  "input",
			input: input,
			want:  "19499881",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Part2(tt.input)
			assert.Equal(t, tt.want, got)
		})
	}
}
