package day2

import (
	"github.com/recrtl/aoc23"
	"github.com/samber/lo"
	"strconv"
	"strings"
)

type gameset struct {
	Red   int
	Green int
	Blue  int
}

func (s gameset) Possible() bool {
	return s.Red <= 12 &&
		s.Green <= 13 &&
		s.Blue <= 14
}

func (s gameset) Power() int {
	return s.Red * s.Green * s.Blue
}

type game struct {
	Game int
	Sets []gameset
}

func (g *game) NecessarySet() gameset {
	necessary := g.Sets[0]
	for _, set := range g.Sets {
		necessary.Red = aoc23.Max(necessary.Red, set.Red)
		necessary.Green = aoc23.Max(necessary.Green, set.Green)
		necessary.Blue = aoc23.Max(necessary.Blue, set.Blue)
	}
	return necessary
}

func parse(input string) []game {
	lines := aoc23.SplitLines(input)
	games := make([]game, len(lines))
	for i, line := range lines {
		game := game{}

		lineSplits := strings.Split(line, ":")
		gameSplits := strings.Split(lineSplits[0], " ")
		game.Game = aoc23.Atoi(gameSplits[1])

		sets := strings.Split(lineSplits[1], ";")
		game.Sets = make([]gameset, len(sets))
		for j, set := range sets {
			splits := strings.Split(set, ",")
			set := gameset{}

			for _, split := range splits {
				splits := strings.Split(split, " ")
				switch splits[2] {
				case "red":
					set.Red = aoc23.Atoi(splits[1])
				case "green":
					set.Green = aoc23.Atoi(splits[1])
				case "blue":
					set.Blue = aoc23.Atoi(splits[1])
				}
			}

			game.Sets[j] = set
		}

		games[i] = game
	}
	return games
}

func Part1(input string) string {
	games := parse(input)

	games = lo.Filter(games, func(game game, _ int) bool {
		for _, set := range game.Sets {
			if !set.Possible() {
				return false
			}
		}
		return true
	})

	sum := lo.SumBy(games, func(game game) int {
		return game.Game
	})

	return strconv.Itoa(sum)
}

func Part2(input string) string {
	games := parse(input)

	powers := lo.Map(games, func(game game, index int) int {
		return game.NecessarySet().Power()
	})

	return strconv.Itoa(lo.Sum(powers))
}
