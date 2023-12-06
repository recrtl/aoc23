package day

import (
	"github.com/recrtl/aoc23"
	"strconv"
	"strings"
)

type race struct {
	Duration      int
	CurrentRecord int
}

func (r *race) WinsCount() int {
	count := 0
	for i := 0; i <= r.Duration; i++ {
		d := (r.Duration - i) * i
		if d > r.CurrentRecord {
			count++
		}
	}
	return count
}

func parse(input string) []race {
	lines := aoc23.SplitLines(input)
	durations := aoc23.SplitNoEmpty(lines[0], " ")
	records := aoc23.SplitNoEmpty(lines[1], " ")

	races := make([]race, len(durations)-1)
	for i := range races {
		races[i] = race{
			Duration:      aoc23.Atoi(durations[i+1]),
			CurrentRecord: aoc23.Atoi(records[i+1]),
		}
	}
	return races
}

func parse2(input string) race {
	lines := aoc23.SplitLines(input)
	durations := aoc23.SplitNoEmpty(lines[0], " ")
	records := aoc23.SplitNoEmpty(lines[1], " ")

	r := race{
		Duration:      aoc23.Atoi(strings.Join(durations[1:], "")),
		CurrentRecord: aoc23.Atoi(strings.Join(records[1:], "")),
	}
	return r
}

func Part1(input string) string {
	races := parse(input)
	product := 1
	for _, r := range races {
		product *= r.WinsCount()
	}
	return strconv.Itoa(product)
}

func Part2(input string) string {
	r := parse2(input)
	c := r.WinsCount()
	return strconv.Itoa(c)
}
