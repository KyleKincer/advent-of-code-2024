package day03

import (
	"advent2024/utils"
	"strconv"
	"strings"
)

func Solve() (interface{}, interface{}) {
	input := utils.ReadInputFile("solutions/day03/input.txt")
	return solvePart1(input), solvePart2(input)
}

func solvePart1(input []string) interface{} {
	var sum int
	for _, l := range input {
		parts := strings.Split(l, "mul(")

		for _, p := range parts {
			if !strings.Contains(p, ")") {
				continue
			}

			numStr := strings.Split(p, ")")
			if !strings.Contains(numStr[0], ",") {
				continue
			}

			numsStr := strings.SplitN(numStr[0], ",", 2)

			if len(numsStr) != 2 {
				continue
			}

			num1, err := strconv.Atoi(numsStr[0])
			if err != nil {
				continue
			}

			num2, err := strconv.Atoi(numsStr[1])
			if err != nil {
				continue
			}

			sum += num1 * num2
		}
	}
	return sum
}

func solvePart2(input []string) interface{} {
	var sum int
	var file string
	for _, l := range input {
		file += l
	}
	for strings.Contains(file, "don't()") {
		start := strings.Index(file, "don't()")
		if strings.Contains(file[start:], "do()") {
			end := strings.Index(file[start:], "do()") + len("do()")
			file = file[:start] + file[start:][end:]
		} else {
			file = file[:start]
		}
	}
	parts := strings.Split(file, "mul(")

	for _, p := range parts {
		if !strings.Contains(p, ")") {
			continue
		}

		numStr := strings.Split(p, ")")
		if !strings.Contains(numStr[0], ",") {
			continue
		}

		numsStr := strings.SplitN(numStr[0], ",", 2)

		if len(numsStr) != 2 {
			continue
		}

		num1, err := strconv.Atoi(numsStr[0])
		if err != nil {
			continue
		}

		num2, err := strconv.Atoi(numsStr[1])
		if err != nil {
			continue
		}

		sum += num1 * num2
	}
	return sum
}
