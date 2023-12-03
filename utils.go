package aoc23

import (
	"golang.org/x/exp/constraints"
	"log"
	"strconv"
	"strings"
)

func splitLines(input string) []string {
	lines := strings.Split(input, "\n")
	for i, line := range lines {
		lines[i] = strings.TrimSuffix(line, "\r")
	}
	return lines
}

func must(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func atoi(input string) int {
	n, err := strconv.Atoi(input)
	must(err)
	return n
}

func max[T constraints.Ordered](x, y T) T {
	if x < y {
		return y
	}
	return x
}
