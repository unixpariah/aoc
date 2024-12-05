package main

import (
	"image"
	"strings"

	"github.com/jpillora/puzzler/harness/aoc"
)

func main() {
	aoc.Harness(run)
}

func run(part2 bool, input string) any {
	grid := map[image.Point]rune{}
	for row, line := range strings.Split(strings.TrimSpace(input), "\n") {
		for col, char := range line {
			grid[image.Point{col, row}] = char
		}
	}

	findSequences := func(start image.Point, length int) []string {
		directions := []image.Point{
			{0, -1},
			{1, 0},
			{0, 1},
			{-1, 0},
			{-1, -1},
			{1, -1},
			{1, 1},
			{-1, 1},
		}

		sequences := make([]string, len(directions))
		for dirIndex, direction := range directions {
			for step := 0; step < length; step++ {
				sequences[dirIndex] += string(grid[start.Add(direction.Mul(step))])
			}
		}

		return sequences
	}

	sum := 0
	for point := range grid {
		if part2 {
			sum += strings.Count("AMAMASASAMAMAS", strings.Join(findSequences(point, 2)[4:], ""))
		} else {
			sum += strings.Count(strings.Join(findSequences(point, 4), " "), "XMAS")
		}
	}

	return sum
}
