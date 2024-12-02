package main

import (
	"sort"
	"strconv"
	"strings"

	"github.com/jpillora/puzzler/harness/aoc"
)

func main() {
	aoc.Harness(run)
}

func run(part2 bool, input string) any {
	lines := strings.Split(strings.TrimSpace(input), "\n")

	var lists [2][]int

	for _, line := range lines {
		parts := strings.Fields(line)

		val_one, err := strconv.Atoi(parts[0])
		if err != nil {
			panic(err)
		}

		val_two, err := strconv.Atoi(parts[1])
		if err != nil {
			panic(err)
		}

		lists[0] = append(lists[0], val_one)
		lists[1] = append(lists[1], val_two)
	}

	if part2 {
		occurences := make(map[int]int)

		sum := 0

		for _, id := range lists[0] {
			if count, exists := occurences[id]; exists {
				sum += id * count
				continue
			}

			for _, item_a := range lists[1] {
				if id == item_a {
					occurences[id]++
				}
			}

			sum += id * occurences[id]
		}

		return sum
	}

	sort.Ints(lists[0])
	sort.Ints(lists[1])

	sum := 0
	for i := 0; i < len(lists[0]); i++ {
		sum += abs(lists[0][i] - lists[1][i])
	}

	return sum
}

func abs(num int) int {
	if num < 0 {
		return -num
	}

	return num
}
