package main

import (
	"slices"
	"strconv"
	"strings"

	"github.com/jpillora/puzzler/harness/aoc"
)

func main() {
	aoc.Harness(run)
}

type Pair struct {
	X string
	Y string
}

func run(part2 bool, input string) any {
	split := strings.Split(strings.TrimSpace(input), "\n\n")
	hashSet := make(map[Pair]struct{})
	sum := 0

	cmp := func(a, b string) int {
		if _, exists := hashSet[Pair{X: a, Y: b}]; exists {
			return -1
		}

		return 0
	}

	for _, line := range strings.Split(split[0], "\n") {
		values := strings.Split(line, "|")

		hashSet[Pair{X: values[0], Y: values[1]}] = struct{}{}
	}

	for _, line := range strings.Split(split[1], "\n") {
		numbers := strings.Split(line, ",")

		correct := slices.IsSortedFunc(numbers, cmp)

		midValue, _ := strconv.Atoi(numbers[len(numbers)/2])

		if correct && !part2 {
			sum += midValue
		} else if !correct && part2 {
			slices.SortFunc(numbers, cmp)

			sum += midValue
		}
	}

	return sum
}

func isUpdateCorrect(hashSet *map[Pair]struct{}, numbers []string) bool {
	set := *hashSet

	for i, outer := range numbers {
		for j, inner := range numbers {
			if j <= i {
				continue
			}

			if _, exists := set[Pair{X: inner, Y: outer}]; exists {
				return false
			}
		}
	}
	return true
}
