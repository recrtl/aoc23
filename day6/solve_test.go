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

func Test_Part1(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{
			name:  "sample",
			input: sample,
			want:  "288",
		},
		{
			name:  "input",
			input: input,
			want:  "275724",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Part1(tt.input)
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_Part2(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{
			name:  "sample",
			input: sample,
			want:  "71503",
		},
		{
			name:  "input",
			input: input,
			want:  "37286485",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Part2(tt.input)
			assert.Equal(t, tt.want, got)
		})
	}
}
