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

	lines := strings.Split(strings.TrimSpace(input), "\n")

	sum := 0

	for _, line := range lines {
		data, err := string_arr_to_int(strings.Split(line, " "))
		if err != nil {
			panic(err)
		}

		prev := data[0]
		increasing := data[0] < data[1]
		for i := 1; i < len(data); i++ {
			val := data[i]

			if increasing != (prev < val) || prev-val == 0 || abs(prev-val) > 3 {
				break
			}

			if i == len(data)-1 {
				sum++
			}

			prev = val
		}
	}
	return sum
}

func string_arr_to_int(arr []string) ([]int, error) {
	var res []int
	for _, val := range arr {
		v, err := strconv.Atoi(val)
		if err != nil {
			return res, err
		}

		res = append(res, v)
	}

	return res, nil
}

func abs(num int) int {
	if num < 0 {
		return -num
	}

	return num
}
