package day

import (
	"fmt"
	"github.com/recrtl/aoc23"
	"github.com/samber/lo"
	"math"
	"strconv"
	"strings"
)

type mapEntry struct {
	Src, Dest, Len int
}

type rangeMap struct {
	Name   string
	Values []mapEntry
}

func (m *rangeMap) GetValue(src int) int {
	for _, value := range m.Values {
		if !aoc23.InRange(src, value.Src, value.Src+value.Len-1) {
			continue
		}
		return src + value.Dest - value.Src
	}

	return src
}

type exercise struct {
	Seeds []int
	Maps  []rangeMap
}

func (e *exercise) Locations() []int {
	res := make([]int, len(e.Seeds))

	for i, seed := range e.Seeds {
		value := seed
		for _, m := range e.Maps {
			value = m.GetValue(value)
		}
		res[i] = value
	}
	return res
}

func (e *exercise) MinLocation2() int {
	min := math.MaxInt

	for i := 0; i < len(e.Seeds); i += 2 {
		startSeed := e.Seeds[i]
		seedLen := e.Seeds[i+1]
		fmt.Println("seed", startSeed, "len", seedLen)

		for seed := startSeed; seed < startSeed+seedLen; seed++ {
			value := seed
			for _, m := range e.Maps {
				value = m.GetValue(value)
			}
			min = aoc23.Min(min, value)
		}
	}

	return min
}

func parse(input string) exercise {
	var res exercise

	input = strings.ReplaceAll(input, "\r\n", "\n")
	blocks := strings.Split(input, "\n\n")

	seedsSplits := strings.Split(blocks[0], ":")
	res.Seeds = aoc23.Atois(aoc23.SplitNoEmpty(seedsSplits[1], " "))

	for _, block := range blocks[1:] {
		lines := aoc23.SplitLines(block)
		m := rangeMap{
			Name:   lines[0],
			Values: make([]mapEntry, 0, len(lines)-1),
		}

		for _, line := range lines[1:] {
			parts := aoc23.Atois(aoc23.SplitNoEmpty(line, " "))
			m.Values = append(m.Values, mapEntry{
				Src:  parts[1],
				Dest: parts[0],
				Len:  parts[2],
			})
		}

		res.Maps = append(res.Maps, m)
	}

	return res
}

func Part1(input string) string {
	e := parse(input)
	l := e.Locations()
	return strconv.Itoa(lo.Min(l))
}

func Part2(input string) string {
	e := parse(input)
	min := e.MinLocation2()
	return strconv.Itoa(min)
}
