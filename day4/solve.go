package day

import (
	"github.com/recrtl/aoc23"
	"github.com/samber/lo"
	"math"
	"strconv"
	"strings"
)

type card struct {
	Number         int
	WinningNumbers []int
	GotNumbers     []int
	Intersection   []int
}

func (c *card) Score() int {
	n := len(c.Intersection)
	if n == 0 {
		return 0
	}
	return int(math.Pow(2, float64(n-1)))
}

func parse(input string) []card {
	lines := aoc23.SplitLines(input)
	cards := make([]card, len(lines))

	for i, line := range lines {
		splits1 := strings.Split(line, ":")
		splits2 := aoc23.SplitTrim(splits1[1], "|", " ")
		c := card{
			Number:         aoc23.Atoi(aoc23.SplitNoEmpty(splits1[0], " ")[1]),
			WinningNumbers: aoc23.Atois(aoc23.SplitNoEmpty(splits2[0], " ")),
			GotNumbers:     aoc23.Atois(aoc23.SplitNoEmpty(splits2[1], " ")),
		}
		c.Intersection = lo.Intersect(c.WinningNumbers, c.GotNumbers)
		cards[i] = c
	}

	return cards
}

func Part1(input string) string {
	cards := parse(input)
	sum := lo.SumBy(cards, func(c card) int {
		return c.Score()
	})
	return strconv.Itoa(sum)
}

func Part2(input string) string {
	cards := parse(input)
	copies := make([]int, len(cards))

	count := 0
	for i, c := range cards {
		totalCopies := 1 + copies[i]
		count += totalCopies
		for j := i + 1; j < i+1+len(c.Intersection); j++ {
			copies[j] += totalCopies
		}
	}

	return strconv.Itoa(count)
}
