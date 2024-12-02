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

func run(part2 bool, input string) any {
	sum := 0

	for _, line := range strings.Split(strings.TrimSpace(input), "\n") {
		report := string_arr_to_int(strings.Split(line, " "))

		safe := is_safe(report)
		if safe {
			sum++
			continue
		}

		if part2 {
			for i := 0; i < len(report); i++ { // RAAHHHHH BRUTE FORCE
				fixed_report := slices.Delete(slices.Clone(report), i, i+1)

				if res := is_safe(fixed_report); res {
					sum++
					break
				}
			}
		}
	}
	return sum
}

func is_safe(report []int) bool {
	prev := report[0]
	decrease, increase := false, false

	for i := 1; i < len(report); i++ {
		val := report[i]

		if prev < val {
			increase = true
		} else if prev > val {
			decrease = true
		} else {
			return false
		}

		if abs(prev-val) > 3 || (decrease && increase) {
			return false
		}

		prev = val
	}

	return true
}

func string_arr_to_int(arr []string) []int {
	var res []int
	for _, val := range arr {
		v, err := strconv.Atoi(val)
		if err != nil {
			panic(err)
		}

		res = append(res, v)
	}

	return res
}

func abs(num int) int {
	if num < 0 {
		return -num
	}

	return num
}
