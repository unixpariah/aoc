package main

import (
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

	sum := 0
	i := 0
	for i < len(input) {
		beggining := strings.Index(input[i:], "mul(")
		if beggining == -1 {
			break
		}

		beggining += i

		comma := strings.Index(input[beggining:], ",")
		if comma == -1 {
			break
		}

		comma += beggining

		num, err := strconv.Atoi(input[beggining+4 : comma])
		if err != nil {
			i = beggining + 4
			continue
		}

		end := strings.Index(input[comma:], ")") + comma
		if end == -1 {
			break
		}

		num_two, err := strconv.Atoi(input[comma+1 : end])
		if err != nil {
			i = comma
			continue
		}

		i = end
		sum += num * num_two
	}

	return sum
}
