package day01

import (
	"advent2024/utils"
	"math"
	"sort"
	"strconv"
	"strings"
)

func Solve() (interface{}, interface{}) {
	input := utils.ReadInputFile("solutions/day01/input.txt")
	return solvePart1(input), solvePart2(input)
}

func solvePart1(input []string) interface{} {
	var list1, list2 []int
	for _, line := range input {
		numbers := strings.Split(line, "   ")
		number1, err := strconv.Atoi(numbers[0])
		if err != nil {
			return err
		}

		number2, err := strconv.Atoi(numbers[1])
		if err != nil {
			return err
		}

		list1 = append(list1, number1)
		list2 = append(list2, number2)
	}

	sort.Slice(list1, func(i, j int) bool {
		return list1[i] < list1[j]
	})

	sort.Slice(list2, func(i, j int) bool {
		return list2[i] < list2[j]
	})

	var totalDiff int
	for i := range list1 {
		diff := math.Abs(float64(list1[i] - list2[i]))
		totalDiff += int(diff)
	}

	return totalDiff
}

func solvePart2(input []string) interface{} {
	var list1, list2 []int
	for _, line := range input {
		numbers := strings.Split(line, "   ")
		number1, err := strconv.Atoi(numbers[0])
		if err != nil {
			return err
		}

		number2, err := strconv.Atoi(numbers[1])
		if err != nil {
			return err
		}

		list1 = append(list1, number1)
		list2 = append(list2, number2)
	}

	var ss int
	for _, v := range list1 {
		var n int
		for _, v2 := range list2 {
			if v == v2 {
				n++
			}
		}
		ss += v * n
	}

	return ss
}
