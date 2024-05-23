package solvers

import (
	"log"
	"math"
	"os"
	"strings"
)

type SolverTen struct {
	grid [][]rune
}

const (
	StartChar      = 'S'
	VerticalPipe   = '|'
	HorizontalPipe = '-'
	NorthEastBend  = 'L'
	NorthWestBend  = 'J'
	SouthWestBend  = '7'
	SouthEastBend  = 'F'
	Ground         = '.'
)

func (s *SolverTen) SolveFirstProblem() int {
	rowIdx, colIdx := startingPoint(s.grid)
	head, tail := s.findHeadAndTail(rowIdx, colIdx)
	farthestDistance, _ := s.traverseAndFindFarthestDistance(head, tail, []int{rowIdx, colIdx})
	return farthestDistance
}

func (s *SolverTen) SolveSecondProblem() int {
	rowIdx, colIdx := startingPoint(s.grid)
	head, tail := s.findHeadAndTail(rowIdx, colIdx)
	_, visited := s.traverseAndFindFarthestDistance(head, tail, []int{rowIdx, colIdx})
	tiles := s.countInnerTiles(visited)
	return tiles
}

func (s *SolverTen) countInnerTiles(visited map[string]bool) int {
	tiles := 0
	for rowIdx, row := range s.grid {
		for colIdx := range row {
			key := idxKey(rowIdx, colIdx)
			if visited[key] {
				continue
			}

			rowInversions := s.countRowInversions(rowIdx, colIdx, visited)
			if rowInversions%2 == 1 {
				tiles++
			}
		}
	}
	return tiles
}

func (s *SolverTen) countRowInversions(rowIdx, colIdx int, visited map[string]bool) int {
	if visited[idxKey(rowIdx, colIdx)] {
		return 0
	}

	inversions := 0

	for colIdx >= 0 {
		if !visited[idxKey(rowIdx, colIdx)] {
			colIdx--
			continue
		}

		cell := s.grid[rowIdx][colIdx]

		if cell == NorthWestBend || cell == NorthEastBend || cell == VerticalPipe {
			inversions++
		}

		colIdx--
	}

	return inversions
}

func (s *SolverTen) traverseAndFindFarthestDistance(head, tail, start []int) (int, map[string]bool) {
	prevHead := start
	prevTail := start

	headCount := 1
	tailCount := 1

	visited := map[string]bool{
		idxKey(start[0], start[1]): true,
		idxKey(head[0], head[1]):   true,
		idxKey(tail[0], tail[1]):   true,
	}

	for {
		if head[0] == tail[0] && head[1] == tail[1] {
			break
		}

		// move head
		head, prevHead = s.getNextCoord(head, prevHead)
		headCount++
		visited[idxKey(head[0], head[1])] = true

		if head[0] == tail[0] && head[1] == tail[1] {
			break
		}

		// move tail
		tail, prevTail = s.getNextCoord(tail, prevTail)
		tailCount++
		visited[idxKey(tail[0], tail[1])] = true

		if head[0] == tail[0] && head[1] == tail[1] {
			break
		}
	}

	return int(math.Max(float64(headCount), float64(tailCount))), visited
}

func (s *SolverTen) getNextCoord(coord, prev []int) ([]int, []int) {
	dirs := directions()

	for idx, dir := range dirs {
		rowIdx := coord[0] + dir[0]
		colIdx := coord[1] + dir[1]
		nextCoord := []int{rowIdx, colIdx}

		if !s.coordInBounds(rowIdx, colIdx) {
			continue
		}

		if rowIdx == prev[0] && colIdx == prev[1] {
			continue
		}

		lookingNorth := idx == 0
		lookingEast := idx == 1
		lookingSouth := idx == 2
		lookingWest := idx == 3

		if (lookingNorth && s.isValidMove(coord, nextCoord)) ||
			(lookingEast && s.isValidMove(coord, nextCoord)) ||
			(lookingSouth && s.isValidMove(coord, nextCoord)) ||
			(lookingWest && s.isValidMove(coord, nextCoord)) {
			return []int{rowIdx, colIdx}, coord
		}
	}

	log.Fatal("No valid move found")

	return []int{}, []int{}
}

func (s *SolverTen) findHeadAndTail(startRowIdx, startColIdx int) (head, tail []int) {
	headAndTail := make([][]int, 0)

	startCoord := []int{startRowIdx, startColIdx}

	for idx, dir := range directions() {
		rowIdx := startRowIdx + dir[0]
		colIdx := startColIdx + dir[1]
		nextCoord := []int{rowIdx, colIdx}

		if !s.coordInBounds(rowIdx, colIdx) {
			continue
		}

		lookingNorth := idx == 0
		lookingEast := idx == 1
		lookingSouth := idx == 2
		lookingWest := idx == 3

		if lookingNorth && s.isValidMove(startCoord, nextCoord) {
			headAndTail = append(headAndTail, nextCoord)
		} else if lookingEast && s.isValidMove(startCoord, nextCoord) {
			headAndTail = append(headAndTail, nextCoord)
		} else if lookingSouth && s.isValidMove(startCoord, nextCoord) {
			headAndTail = append(headAndTail, nextCoord)
		} else if lookingWest && s.isValidMove(startCoord, nextCoord) {
			headAndTail = append(headAndTail, nextCoord)
		}
	}

	if len(headAndTail) < 2 {
		log.Fatal("Head or tail not found")
	}

	return headAndTail[0], headAndTail[1]
}

func (s *SolverTen) isValidMove(curr []int, next []int) bool {
	isNorthMovement := curr[0]-1 == next[0] && curr[1] == next[1]
	isEastMovement := curr[0] == next[0] && curr[1]+1 == next[1]
	isSouthMovement := curr[0]+1 == next[0] && curr[1] == next[1]
	isWestMovement := curr[0] == next[0] && curr[1]-1 == next[1]

	currCell := s.grid[curr[0]][curr[1]]
	nextCell := s.grid[next[0]][next[1]]

	if isNorthMovement {
		isValidCurrCell := currCell == StartChar || currCell == VerticalPipe || currCell == NorthEastBend || currCell == NorthWestBend
		isValidNextCell := nextCell == VerticalPipe || nextCell == SouthEastBend || nextCell == SouthWestBend
		return isValidCurrCell && isValidNextCell
	} else if isEastMovement {
		isValidCurrCell := currCell == StartChar || currCell == HorizontalPipe || currCell == NorthEastBend || currCell == SouthEastBend
		isValidNextCell := nextCell == HorizontalPipe || nextCell == NorthWestBend || nextCell == SouthWestBend
		return isValidCurrCell && isValidNextCell
	} else if isSouthMovement {
		isValidCurrCell := currCell == StartChar || currCell == VerticalPipe || currCell == SouthEastBend || currCell == SouthWestBend
		isValidNextCell := nextCell == VerticalPipe || nextCell == NorthEastBend || nextCell == NorthWestBend
		return isValidCurrCell && isValidNextCell
	} else if isWestMovement {
		isValidCurrCell := currCell == StartChar || currCell == HorizontalPipe || currCell == SouthWestBend || currCell == NorthWestBend
		isValidNextCell := nextCell == HorizontalPipe || nextCell == SouthEastBend || nextCell == NorthEastBend
		return isValidCurrCell && isValidNextCell
	}

	log.Fatal("Invalid direction found")
	return false
}

func directions() [][]int {
	return [][]int{
		{-1, 0}, // north
		{0, 1},  // east
		{1, 0},  // south
		{0, -1}, // west
	}
}

func (s *SolverTen) coordInBounds(rowIdx, colIdx int) bool {
	return rowIdx >= 0 && rowIdx < len(s.grid) && colIdx >= 0 && colIdx < len(s.grid[0])
}

func startingPoint(content [][]rune) (int, int) {
	for rowIdx, row := range content {
		for colIdx, cell := range row {
			if cell == 'S' {
				return rowIdx, colIdx
			}
		}
	}

	log.Fatal("Starting point not found")
	return -1, -1
}

func (s *SolverTen) Parse(file string) {
	content, err := os.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}

	rows := make([][]rune, 0)
	for _, line := range strings.Split(string(content), "\n") {
		trimmedLine := strings.TrimSpace(line)
		if trimmedLine == "" {
			continue
		}

		rows = append(rows, []rune(trimmedLine))
	}

	s.grid = rows
}
