package day3

import (
	"github.com/recrtl/aoc23"
	"github.com/samber/lo"
	"strconv"
)

type part struct {
	Row, ColMin, ColMax int
	Value               int
}

func (p *part) IsPartNumber(symbols []symbol) bool {
	for _, s := range symbols {
		if s.IsNextTo(*p) {
			return true
		}
	}

	return false
}

type symbol struct {
	Row, Col int
	Value    string
}

func (s *symbol) IsNextTo(p part) bool {
	return aoc23.InRange(s.Row, p.Row-1, p.Row+1) &&
		aoc23.InRange(s.Col, p.ColMin-1, p.ColMax+1)
}

func (s *symbol) GetAdjacentParts(parts []part) []part {
	return lo.Filter(parts, func(p part, _ int) bool {
		return s.IsNextTo(p)
	})
}

type schematic struct {
	Parts   []part
	Symbols []symbol
}

func parse(input string) schematic {
	s := schematic{
		Parts:   make([]part, 0),
		Symbols: make([]symbol, 0),
	}

	for row, line := range aoc23.SplitLines(input) {
		var currentPart *part = nil
		for col, char := range line {
			if char >= '0' && char <= '9' {
				value := aoc23.Atoi(string(char))

				if currentPart == nil {
					currentPart = &part{
						Row:    row,
						ColMin: col,
						ColMax: col,
						Value:  aoc23.Atoi(string(char)),
					}
					continue
				}

				currentPart.ColMax++
				currentPart.Value = 10*currentPart.Value + value
				continue
			}

			if currentPart != nil {
				s.Parts = append(s.Parts, *currentPart)
				currentPart = nil
			}

			if char == '.' {
				continue
			}

			s.Symbols = append(s.Symbols, symbol{
				Row:   row,
				Col:   col,
				Value: string(char),
			})
		}

		if currentPart != nil {
			s.Parts = append(s.Parts, *currentPart)
		}
	}

	return s
}

func Part1(input string) string {
	s := parse(input)

	parts := lo.Filter(s.Parts, func(p part, _ int) bool {
		return p.IsPartNumber(s.Symbols)
	})

	return strconv.Itoa(lo.SumBy(parts, func(p part) int {
		return p.Value
	}))
}

func Part2(input string) string {
	schem := parse(input)

	gearSum := 0

	for _, s := range schem.Symbols {
		if s.Value != "*" {
			continue
		}

		parts := s.GetAdjacentParts(schem.Parts)
		if len(parts) != 2 {
			continue
		}

		gearSum += parts[0].Value * parts[1].Value
	}

	return strconv.Itoa(gearSum)
}
