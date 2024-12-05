package main

import (
	"strconv"
	"strings"

	"github.com/jpillora/puzzler/harness/aoc"
)

func main() {
	aoc.Harness(run)
}

type Pair struct {
	X int
	Y int
}

func run(part2 bool, input string) any {
	if part2 {
		return "not implemented"
	}

	split := strings.Split(strings.TrimSpace(input), "\n\n")
	hashSet := make(map[Pair]struct{})
	sum := 0

	for _, line := range strings.Split(split[0], "\n") {
		values := strings.Split(line, "|")
		key, err1 := strconv.Atoi(values[0])
		value, err2 := strconv.Atoi(values[1])
		if err1 == nil && err2 == nil {
			hashSet[Pair{X: key, Y: value}] = struct{}{}
		}
	}

	for _, line := range strings.Split(split[1], "\n") {
		numbers := strings.Split(line, ",")
		if isUpdateCorrect(&hashSet, numbers) {
			midIndex := len(numbers) / 2
			midValue, err := strconv.Atoi(numbers[midIndex])
			if err == nil {
				sum += midValue
			}
		}
	}

	return sum
}

func isUpdateCorrect(hashSet *map[Pair]struct{}, numbers []string) bool {
	for i, outer := range numbers {
		for j, inner := range numbers {
			if j <= i {
				continue
			}
			valOuter, err1 := strconv.Atoi(outer)
			valInner, err2 := strconv.Atoi(inner)
			if err1 == nil && err2 == nil {
				if _, exists := (*hashSet)[Pair{X: valInner, Y: valOuter}]; exists {
					return false
				}
			}
		}
	}
	return true
}
