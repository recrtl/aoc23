package aoc23

import (
	"github.com/samber/lo"
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

func SplitTrim(input string, sep string, cutset string) []string {
	splits := strings.Split(input, sep)
	for i, split := range splits {
		splits[i] = strings.Trim(split, cutset)
	}
	return splits
}

func SplitNoEmpty(input string, sep string) []string {
	splits := strings.Split(input, sep)
	return lo.WithoutEmpty(splits)
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

func Atois(inputs []string) []int {
	return lo.Map(inputs, func(item string, _ int) int {
		return Atoi(item)
	})
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
