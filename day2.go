package aoc23

import (
	"github.com/samber/lo"
	"strconv"
	"strings"
)

type Day21Set struct {
	Red   int
	Green int
	Blue  int
}

func (s Day21Set) Possible() bool {
	return s.Red <= 12 &&
		s.Green <= 13 &&
		s.Blue <= 14
}

func (s Day21Set) Power() int {
	return s.Red * s.Green * s.Blue
}

type Day21Game struct {
	Game int
	Sets []Day21Set
}

func (g *Day21Game) NecessarySet() Day21Set {
	necessary := g.Sets[0]
	for _, set := range g.Sets {
		necessary.Red = max(necessary.Red, set.Red)
		necessary.Green = max(necessary.Green, set.Green)
		necessary.Blue = max(necessary.Blue, set.Blue)
	}
	return necessary
}

func Day21ParseGames(input string) []Day21Game {
	lines := splitLines(input)
	games := make([]Day21Game, len(lines))
	for i, line := range lines {
		game := Day21Game{}

		lineSplits := strings.Split(line, ":")
		gameSplits := strings.Split(lineSplits[0], " ")
		game.Game = atoi(gameSplits[1])

		sets := strings.Split(lineSplits[1], ";")
		game.Sets = make([]Day21Set, len(sets))
		for j, set := range sets {
			splits := strings.Split(set, ",")
			set := Day21Set{}

			for _, split := range splits {
				splits := strings.Split(split, " ")
				switch splits[2] {
				case "red":
					set.Red = atoi(splits[1])
				case "green":
					set.Green = atoi(splits[1])
				case "blue":
					set.Blue = atoi(splits[1])
				}
			}

			game.Sets[j] = set
		}

		games[i] = game
	}
	return games
}

func Day21(input string) string {
	games := Day21ParseGames(input)

	games = lo.Filter(games, func(game Day21Game, _ int) bool {
		for _, set := range game.Sets {
			if !set.Possible() {
				return false
			}
		}
		return true
	})

	sum := lo.SumBy(games, func(game Day21Game) int {
		return game.Game
	})

	return strconv.Itoa(sum)
}

func Day22(input string) string {
	games := Day21ParseGames(input)

	powers := lo.Map(games, func(game Day21Game, index int) int {
		return game.NecessarySet().Power()
	})

	return strconv.Itoa(lo.Sum(powers))
}
