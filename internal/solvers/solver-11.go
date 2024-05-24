package solvers

import (
	"log"
	"math"
	"os"
	"strings"
)

type SolverEleven struct {
	grid  Grid
	adder int
}

type Grid [][]rune

const (
	Hash = '#'
	Dot  = '.'
)

func (s *SolverEleven) SolveFirstProblem() int {
	expandedGrid := s.expandGrid()
	pairs := s.findPairs(expandedGrid)

	distances := make([]int, 0)
	for _, pair := range pairs {
		distance := s.distance(pair[0], pair[1])
		distances = append(distances, distance)
	}

	sumOfDistances := 0
	for _, distance := range distances {
		sumOfDistances += distance
	}

	return sumOfDistances
}

func (s *SolverEleven) distance(from []int, to []int) int {
	left := math.Abs(float64(from[0] - to[0]))
	right := math.Abs(float64(from[1] - to[1]))

	return int(left + right)
}

func (s *SolverEleven) distanceWithAdders(from []int, to []int, rowAdder []int, colAdder []int) int {
	addedFrom := []int{from[0] + rowAdder[from[0]], from[1] + colAdder[from[1]]}
	addedTo := []int{to[0] + rowAdder[to[0]], to[1] + colAdder[to[1]]}

	left := math.Abs(float64(addedFrom[0] - addedTo[0]))
	right := math.Abs(float64(addedFrom[1] - addedTo[1]))

	return int(left + right)
}

func (s *SolverEleven) SolveSecondProblem() int {
	pairs := s.findPairs(s.grid)

	// for easier testing
	if s.adder == 0 {
		s.adder = 999999
	}

	rowAdder := s.getRowAdder()
	colAdder := s.getColAdder()

	distances := make([]int, 0)
	for _, pair := range pairs {
		distance := s.distanceWithAdders(pair[0], pair[1], rowAdder, colAdder)
		distances = append(distances, distance)
	}

	sumOfDistances := 0
	for _, distance := range distances {
		sumOfDistances += distance
	}

	return sumOfDistances
}

func (s *SolverEleven) getRowAdder() []int {
	adder := make([]int, len(s.grid))
	addBy := 0
	for idx := range s.grid {
		if rowShouldBeExpanded(s.grid, idx) {
			addBy += s.adder
		}
		adder[idx] = addBy
	}
	return adder
}

func (s *SolverEleven) getColAdder() []int {
	adder := make([]int, len(s.grid[0]))
	addBy := 0
	for i := range s.grid[0] {
		if colShouldBeExpanded(s.grid, i) {
			addBy += s.adder

		}
		adder[i] = addBy
	}
	return adder
}

func (s *SolverEleven) Parse(file string) {
	content, err := os.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}

	rows := make(Grid, 0)
	for _, line := range strings.Split(string(content), "\n") {
		trimmedLine := strings.TrimSpace(line)
		if trimmedLine == "" {
			continue
		}

		rows = append(rows, []rune(trimmedLine))
	}

	s.grid = rows
}

func (s *SolverEleven) expandGrid() Grid {
	rowsToExpand := make(map[int]bool, 0)
	colsToExpand := make(map[int]bool, 0)

	for idx := range s.grid {
		if rowShouldBeExpanded(s.grid, idx) {
			rowsToExpand[idx] = true
		}
	}

	for idx := range s.grid[0] {
		if colShouldBeExpanded(s.grid, idx) {
			colsToExpand[idx] = true
		}
	}

	expandedGrid := make(Grid, 0)
	for rowIdx, row := range s.grid {
		currRow := make([]rune, 0)
		for colIdx, cell := range row {
			currRow = append(currRow, cell)
			if _, ok := colsToExpand[colIdx]; ok {
				currRow = append(currRow, cell)
			}
		}

		expandedGrid = append(expandedGrid, currRow)
		if _, ok := rowsToExpand[rowIdx]; ok {
			expandedGrid = append(expandedGrid, currRow)
		}
	}

	return expandedGrid
}

func (s *SolverEleven) findPairs(expandedGrid Grid) [][][]int {
	hashes := s.findHashes(expandedGrid)
	pairs := make([][][]int, 0)

	for i := 0; i < len(hashes); i++ {
		for j := i + 1; j < len(hashes); j++ {
			pairs = append(
				pairs,
				[][]int{
					{hashes[i][0], hashes[i][1]},
					{hashes[j][0], hashes[j][1]},
				},
			)
		}
	}

	return pairs
}

func (s *SolverEleven) findHashes(expandedGrid Grid) [][]int {
	hashes := make([][]int, 0)
	for rowIdx, row := range expandedGrid {
		for colIdx, cell := range row {
			if cell == Hash {
				hashes = append(hashes, []int{rowIdx, colIdx})
			}
		}
	}
	return hashes
}

func (g Grid) String() string {
	var sb strings.Builder
	for _, row := range g {
		sb.WriteString(string(row))
		sb.WriteString("\n")
	}
	return sb.String()
}

func rowShouldBeExpanded(grid Grid, rowIdx int) bool {
	for _, cell := range grid[rowIdx] {
		if cell == Hash {
			return false
		}
	}
	return true
}

func colShouldBeExpanded(grid Grid, colIdx int) bool {
	for idx := 0; idx < len(grid); idx++ {
		if grid[idx][colIdx] == Hash {
			return false
		}
	}
	return true
}
