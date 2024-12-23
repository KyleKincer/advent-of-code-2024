package day05

import (
	"advent2024/utils"
	"slices"
	"sort"
	"strconv"
	"strings"
)

func Solve() (interface{}, interface{}) {
	input := utils.ReadInputFile("solutions/day05/input.txt")
	return solvePart1(input), solvePart2(input)
}

type Rule struct {
	first  int
	second int
}

type Update struct {
	pages []int
}

func NewRule(s string) *Rule {
	pagesStr := strings.Split(s, "|")
	first, err := strconv.Atoi(pagesStr[0])
	if err != nil {
		panic(err)
	}

	second, err := strconv.Atoi(pagesStr[1])
	if err != nil {
		panic(err)
	}
	return &Rule{
		first:  first,
		second: second,
	}
}

func NewUpdate(s string) *Update {
	var pages []int
	pagesStr := strings.Split(s, ",")
	for _, pageStr := range pagesStr {
		page, err := strconv.Atoi(pageStr)
		if err != nil {
			panic(err)
		}
		pages = append(pages, page)
	}
	return &Update{
		pages: pages,
	}
}

func (u Update) ruleApplies(r *Rule) bool {
	return slices.Contains(u.pages, r.first) && slices.Contains(u.pages, r.second)
}

func (u Update) rulePasses(r *Rule) bool {
	return slices.Index(u.pages, r.first) < slices.Index(u.pages, r.second)
}

func (u Update) getMiddlePage() int {
	return u.pages[int(len(u.pages)/2)]
}

func (u *Update) updateForRule(rule *Rule) {
	sort.Slice(u.pages, func(i, j int) bool {
		if u.pages[i] == rule.first && u.pages[j] == rule.second {
			return true
		}
		if u.pages[i] == rule.second && u.pages[j] == rule.first {
			return false
		}

		return i < j
	})
}

func (u *Update) updateForAllRules(rules []*Rule) {
	sort.Slice(u.pages, func(i, j int) bool {
		for _, rule := range rules {
			if u.pages[i] == rule.first && u.pages[j] == rule.second {
				return true
			}
			if u.pages[i] == rule.second && u.pages[j] == rule.first {
				return false
			}
		}
		return i < j
	})
}

func solvePart1(input []string) interface{} {
	var rules []*Rule
	var updates []*Update
	for _, l := range input {
		if strings.Contains(l, "|") {
			rules = append(rules, NewRule(l))
		}

		if strings.Contains(l, ",") {
			updates = append(updates, NewUpdate(l))
		}
	}

	var passes []*Update
	for _, update := range updates {
		pass := true
		for _, rule := range rules {
			if !update.ruleApplies(rule) {
				continue
			}
			if !update.rulePasses(rule) {
				pass = false
				break
			}

		}
		if pass {
			passes = append(passes, update)
		}
	}

	var sum int
	for _, update := range passes {
		sum += update.getMiddlePage()
	}
	return sum
}

func solvePart2(input []string) interface{} {
	var rules []*Rule
	var updates []*Update
	for _, l := range input {
		if strings.Contains(l, "|") {
			rules = append(rules, NewRule(l))
		}

		if strings.Contains(l, ",") {
			updates = append(updates, NewUpdate(l))
		}
	}

	var corrected []Update
	for _, update := range updates {
		fixed := false
		for _, rule := range rules {
			if !update.ruleApplies(rule) {
				continue
			}

			if !update.rulePasses(rule) {
				update.updateForAllRules(rules)
				fixed = true
				break
			}
		}
		if fixed {
			corrected = append(corrected, *update)
		}
	}

	var sum int
	for _, update := range corrected {
		sum += update.getMiddlePage()
	}
	return sum
}
