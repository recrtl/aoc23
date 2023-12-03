package aoc23

import (
	"golang.org/x/exp/constraints"
	"log"
	"strconv"
	"strings"
)

func SplitLines(input string) []string {
	lines := strings.Split(input, "\n")
	for i, line := range lines {
		lines[i] = strings.TrimSuffix(line, "\r")
	}
	return lines
}

func Must(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func Atoi(input string) int {
	n, err := strconv.Atoi(input)
	Must(err)
	return n
}

func Max[T constraints.Ordered](x, y T) T {
	if x < y {
		return y
	}
	return x
}

func InRange[T constraints.Ordered](x, min, max T) bool {
	return x >= min && x <= max
}
