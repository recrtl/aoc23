package day1

import (
	"github.com/recrtl/aoc23"
	"log"
	"regexp"
	"strconv"
	"strings"
)

func Part1(input string) string {
	r := regexp.MustCompile("([0-9])")

	sum := 0
	lines := aoc23.SplitLines(input)
	for _, line := range lines {
		matches := r.FindAllString(line, -1)
		n, _ := strconv.Atoi(matches[0] + matches[len(matches)-1])
		sum += n
	}

	return strconv.Itoa(sum)
}

func Part2(input string) string {
	matches := []string{"one", "1", "two", "2", "three", "3", "four", "4", "five", "5", "six", "6", "seven", "7", "eight", "8", "nine", "9"}
	replacer := strings.NewReplacer(matches...)

	sum := 0
	lines := aoc23.SplitLines(input)
	for _, line := range lines {
		first := ""
		last := ""

		for i := 0; i < len(line); i++ {
			prefix, ok := matchPrefix(line[i:], matches)
			if ok {
				first = replacer.Replace(prefix)
				break
			}
		}

		for i := len(line) - 1; i >= 0; i-- {
			prefix, ok := matchPrefix(line[i:], matches)
			if ok {
				last = replacer.Replace(prefix)
				break
			}
		}

		if first == "" || last == "" {
			log.Fatalln("no result", line, first, last)
		}

		n, _ := strconv.Atoi(first + last)
		sum += n
	}

	return strconv.Itoa(sum)
}

func matchPrefix(value string, prefixes []string) (string, bool) {
	for _, prefix := range prefixes {
		if strings.HasPrefix(value, prefix) {
			return prefix, true
		}
	}

	return "", false
}
