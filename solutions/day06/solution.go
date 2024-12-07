package day06

import "advent2024/utils"

func Solve() (interface{}, interface{}) {
	input := utils.ReadInputFile("solutions/day06/input.txt")
	return solvePart1(input), solvePart2(input)
}

const (
	Empty    = "."
	Obstacle = "#"
	Guard    = "^"
)

type Map struct {
	spaces        [][]Space
	guard         GuardSpace
	height, width int
}

type GuardSpace struct {
	position Position
	heading  Heading
	space    *Space
}

type Position struct {
	col, row int
}

type Heading int

const (
	North Heading = iota
	East
	South
	West
)

type Space struct {
	visited  bool
	obstacle bool
	guard    bool
}

func NewMap(input []string) *Map {
	newMap := &Map{}
	for i, row := range input {
		newRow := make([]Space, 0, len(row))
		for j, space := range row {
			var newSpace *Space
			switch string(space) {
			case Empty:
				newSpace = &Space{}
			case Obstacle:
				newSpace = &Space{obstacle: true}
			case Guard:
				newSpace = &Space{guard: true}
				newMap.guard = GuardSpace{
					position: Position{
						col: j,
						row: i,
					},
					heading: North,
					space:   newSpace,
				}
			}
			newRow = append(newRow, *newSpace)
		}
		newMap.spaces = append(newMap.spaces, newRow)
	}

	newMap.width = len(newMap.spaces[0])
	newMap.height = len(newMap.spaces)
	return newMap
}

func (s *Space) SetVisited() {
	s.visited = true
}

func (s Space) IsGuard() bool {
	return s.guard
}

func (s Space) IsObstacle() bool {
	return s.obstacle
}

func (m Map) GuardHasValidMove() bool {
	nextPosition := m.GetNextGuardPosition()
	if nextPosition.col < 0 || nextPosition.col > m.width-1 {
		return false
	}

	if nextPosition.row < 0 || nextPosition.row > m.height-1 {
		return false
	}

	return true
}

func (m Map) GetNextGuardPosition() *Position {
	var nextPosition Position
	switch m.guard.heading {
	case North:
		nextPosition = Position{
			col: m.guard.position.col,
			row: m.guard.position.row - 1,
		}
	case South:
		nextPosition = Position{
			col: m.guard.position.col,
			row: m.guard.position.row + 1,
		}
	case East:
		nextPosition = Position{
			col: m.guard.position.col + 1,
			row: m.guard.position.row,
		}
	case West:
		nextPosition = Position{
			col: m.guard.position.col - 1,
			row: m.guard.position.row,
		}
	}
	return &nextPosition
}

func (m Map) GetSpaceAtPosition(p Position) *Space {
	return &m.spaces[p.row][p.col]
}

func (m *Map) TurnGuardRight() {
	m.guard.heading = (m.guard.heading + 1) % 4
}

func (m *Map) MoveGuard(p Position) {
	m.GetSpaceAtPosition(m.guard.position).SetVisited()
	m.guard.position = p
}

func (m Map) CountVisitedSpaces() int {
	var count int
	for _, row := range m.spaces {
		for _, space := range row {
			if space.visited {
				count++
			}
		}
	}
	return count
}

func (m *Map) RunGuard() {
	for m.GuardHasValidMove() {
		nextPosition := m.GetNextGuardPosition()
		if m.spaces[nextPosition.row][nextPosition.col].IsObstacle() {
			m.TurnGuardRight()
			continue
		}

		m.MoveGuard(*nextPosition)
	}
	m.GetSpaceAtPosition(m.guard.position).SetVisited()
}

func solvePart1(input []string) interface{} {
	puzzleMap := NewMap(input)
	puzzleMap.RunGuard()
	visited := puzzleMap.CountVisitedSpaces()
	return visited
}

func solvePart2(input []string) interface{} {
	return nil
}
