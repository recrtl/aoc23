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
			want:  "35",
		},
		{
			name:  "input",
			input: input,
			want:  "457535844",
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
			want:  "46",
		},
		{
			name:  "input",
			input: input,
			want:  "41222968",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Part2(tt.input)
			assert.Equal(t, tt.want, got)
		})
	}
}
