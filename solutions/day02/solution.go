package day02

import (
	"advent2024/utils"
	"math"
	"strconv"
	"strings"
)

func Solve() (interface{}, interface{}) {
	input := utils.ReadInputFile("solutions/day02/input.txt")
	return solvePart1(input), solvePart2(input)
}

func solvePart1(input []string) interface{} {
	var safe int
	for _, v := range input {
		reportStrs := strings.Split(v, " ")
		report := make([]int, len(reportStrs))
		for j, s := range reportStrs {
			i, err := strconv.Atoi(s)
			if err != nil {
				panic(err)
			}
			report[j] = i
		}

		inc := report[1] > report[0]
		for k := range report {
			if k == len(report)-1 {
				safe++
				break
			}
			current := report[k]
			next := report[k+1]

			if inc && next < current {
				break
			}

			if !inc && next > current {
				break
			}

			diff := math.Abs(float64(current - next))
			if diff < 1 || diff > 3 {
				break
			}
		}
	}

	return safe
}

func solvePart2(input []string) interface{} {
	var safe int
	for _, v := range input {
		reportStrs := strings.Split(v, " ")
		report := make([]int, len(reportStrs))
		for j, s := range reportStrs {
			i, err := strconv.Atoi(s)
			if err != nil {
				panic(err)
			}
			report[j] = i
		}

		if reportIsSafe(report) {
			safe++
		} else {
			for k := range report {
				var reportCpy []int
				for l := range report {
					if l == k {
						continue
					}
					reportCpy = append(reportCpy, report[l])
				}
				if reportIsSafe(reportCpy) {
					safe++
					break
				}
			}
		}
	}

	return safe
}

func reportIsSafe(report []int) bool {
	inc := report[1] > report[0]
	for i := range report {
		if i == len(report)-1 {
			return true
		}
		current := report[i]
		next := report[i+1]

		if inc && next < current {
			return false
		}

		if !inc && next > current {
			return false
		}

		diff := math.Abs(float64(current - next))
		if diff < 1 || diff > 3 {
			return false
		}
	}
	return false
}
