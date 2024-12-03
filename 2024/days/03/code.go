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
	sum := 0
	i := 0
	for i < len(input) {
		if part2 {
			dont := strings.Index(input[i:], "don't()")
			beggining := strings.Index(input[i:], "mul(")

			if dont < beggining && dont != -1 && beggining != -1 {
				do := strings.Index(input[i:], "do()")
				if do == -1 {
					break
				}

				i += do
			}
		}

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
