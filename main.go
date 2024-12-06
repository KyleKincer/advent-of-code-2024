package main

import (
	"advent2024/solutions/day01"
	"advent2024/solutions/day02"
	"fmt"
)

func main() {
	// Day 1
	part1, part2 := day01.Solve()
	fmt.Printf("Day 1:\n  Part 1: %v\n  Part 2: %v\n", part1, part2)

	// Day 2
	part1, part2 = day02.Solve()
	fmt.Printf("Day 2:\n  Part 1: %v\n  Part 2: %v\n", part1, part2)

}
