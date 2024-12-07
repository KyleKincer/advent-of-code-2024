package day04

import (
	"advent2024/utils"
)

func Solve() (interface{}, interface{}) {
	input := utils.ReadInputFile("solutions/day04/input.txt")
	return solvePart1(input), solvePart2(input)
}

type Crossword struct {
	puzzle       Puzzle
	totalRows    int
	totalColumns int
}

type Puzzle struct {
	board [][]Cell
}

type Cell struct {
	letter string
}

func (p *Puzzle) countMatches(t string) int {
	var matches int
	var s []string
	for i := range t {
		s = append(s, string(t[i]))
	}

	for i, row := range p.board {
		for j, cell := range row {
			if cell.letter == s[0] {
				t := s[1:]

				// Up
				if i >= len(t) {
					for k, targetLetter := range t {
						next := p.board[i-k-1][j]
						if next.letter != targetLetter {
							break
						}

						if k == len(t)-1 {
							matches++
						}
					}
				}

				// Up + Right
				if i >= len(t) && j < len(row)-len(t) {
					for k, targetLetter := range t {
						next := p.board[i-k-1][j+k+1]

						if next.letter != targetLetter {
							break
						}

						if k == len(t)-1 {
							matches++
						}
					}
				}

				// Right
				if j < len(row)-len(t) {
					for k, targetLetter := range t {
						next := p.board[i][j+k+1]

						if next.letter != targetLetter {
							break
						}

						if k == len(t)-1 {
							matches++
						}
					}
				}

				// Down + Right
				if i <= len(p.board)-len(t) && j < len(row)-len(t) {
					for k, targetLetter := range t {
						next := p.board[i+k+1][j+k+1]

						if next.letter != targetLetter {
							break
						}

						if k == len(t)-1 {
							matches++
						}
					}
				}

				// Down
				if i <= len(p.board)-len(t) {
					for k, targetLetter := range t {
						next := p.board[i+k+1][j]
						if next.letter != targetLetter {
							break
						}

						if k == len(t)-1 {
							matches++
						}
					}
				}

				// Down + Left
				if i < len(p.board)-len(t) && j >= len(t) {
					for k, targetLetter := range t {
						next := p.board[i+k+1][j-(k+1)]

						if next.letter != targetLetter {
							break
						}

						if k == len(t)-1 {
							matches++
						}
					}
				}

				// Left
				if j >= len(t) {
					for k, targetLetter := range t {
						next := p.board[i][j-(k+1)]

						if next.letter != targetLetter {
							break
						}

						if k == len(t)-1 {
							matches++
						}
					}
				}

				// Up + Left
				if i >= len(t) && j >= len(t) {
					for k, targetLetter := range t {
						next := p.board[i-(k+1)][j-(k+1)]

						if next.letter != targetLetter {
							break
						}

						if k == len(t)-1 {
							matches++
						}
					}
				}

			}
		}
	}
	return matches
}

func NewCrossword(lines []string, target string) *Crossword {
	var crossword Crossword
	crossword.puzzle = Puzzle{}
	for _, line := range lines {
		row := make([]Cell, 0, len(line))
		for i := range line {
			cell := &Cell{
				letter: string(line[i]),
			}
			row = append(row, *cell)
		}
		crossword.puzzle.board = append(crossword.puzzle.board, row)
	}
	crossword.totalRows = len(crossword.puzzle.board)
	crossword.totalColumns = len(crossword.puzzle.board[0])

	return &crossword
}

func solvePart1(input []string) interface{} {
	c := NewCrossword(input, "XMAS")
	matches := c.puzzle.countMatches("XMAS")
	return matches
}

func solvePart2(input []string) interface{} {
	return nil
}
