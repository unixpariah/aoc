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
	if part2 {
		return "not implemented"
	}

	var res = strings.Split(input, "\n")

	var list_one []int
	var list_two []int

	for _, v := range res {
		var r = strings.Split(v, "   ")

		if len(r) < 2 {
			continue
		}

		i_o, err := strconv.Atoi(r[0])
		if err != nil {
			panic(err)
		}

		i_t, err := strconv.Atoi(r[1])
		if err != nil {
			panic(err)
		}

		list_one = append(list_one, i_o)
		list_two = append(list_two, i_t)
	}

	sort.Slice(list_one, func(i, j int) bool {
		return list_one[i] < list_one[j]
	})

	sort.Slice(list_two, func(i, j int) bool {
		return list_two[i] < list_two[j]
	})

	var sum = 0
	for i := 0; i < len(list_one); i++ {
		if list_two[i] < list_one[i] {
			sum += list_one[i] - list_two[i]
		} else {
			sum += list_two[i] - list_one[i]
		}
	}

	return sum
}
