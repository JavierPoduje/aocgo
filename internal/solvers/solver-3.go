package solvers

import (
	"log"
	"os"
	"strconv"
	"strings"
)

type SolverThree struct {
	content [][]string
}

func (s *SolverThree) Parse(file string) {
	content, err := os.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}

	rows := make([][]string, 0)
	for _, line := range strings.Split(string(content), "\n") {
		trimmedLine := strings.TrimSpace(line)
		if trimmedLine == "" {
			continue
		}
		splittedRow := strings.Split(trimmedLine, "")
		rows = append(rows, splittedRow)
	}

	s.content = rows
}

func (s *SolverThree) SolveFirstProblem() int {
	numberParts := findNumberParts(s.content)
	sum := 0
	for _, numberPart := range numberParts {
		sum += numberPart
	}
	return sum
}

func (s *SolverThree) SolveSecondProblem() int {
	return 0
}

func findNumberParts(content [][]string) []int {
	numberParts := make([]int, 0)
	usedIdxs := map[string]bool{}

	for rowIdx, row := range content {
		for colIdx, char := range row {
			if charIsNumber(char) && hasAdjacentSign(content, rowIdx, colIdx) && !isUsed(usedIdxs, rowIdx, colIdx) {
				numberPart, numberPartIdxs := buildNumberPart(content, rowIdx, colIdx)
				numberParts = append(numberParts, numberPart)

				for _, idx := range numberPartIdxs {
					row := idx[0]
					col := idx[1]
					usedIdxs[idxKey(row, col)] = true
				}
			}
		}
	}

	return numberParts
}

func isUsed(usedIdxs map[string]bool, rowIdx int, colIdx int) bool {
	return usedIdxs[idxKey(rowIdx, colIdx)]
}

func buildNumberPart(content [][]string, rowIdx int, colIdx int) (int, [][]int) {
	pivot := content[rowIdx][colIdx]
	numberSlice := []string{pivot}
	usedIdxs := [][]int{
		{rowIdx, colIdx},
	}

	left := colIdx - 1
	for {
		if left < 0 || !charIsNumber(content[rowIdx][left]) {
			break
		}

		numberChar := content[rowIdx][left]
		usedIdxs = append(usedIdxs, []int{rowIdx, left})
		numberSlice = append([]string{numberChar}, numberSlice...)
		left--
	}

	right := colIdx + 1
	for {
		if right >= len(content[0]) || !charIsNumber(content[rowIdx][right]) {
			break
		}

		numberChar := content[rowIdx][right]
		usedIdxs = append(usedIdxs, []int{rowIdx, right})
		numberSlice = append(numberSlice, numberChar)
		right++
	}

	joinedNumber := strings.Join(numberSlice, "")
	number, err := strconv.Atoi(joinedNumber)
	if err != nil {
		log.Fatalf("buildNumberPart is not building the numbers correctly: %v", err)
	}

	return number, usedIdxs
}

func hasAdjacentSign(content [][]string, rowIdx int, colIdx int) bool {
	dirs := getDirections()
	for _, dir := range dirs {
		row := rowIdx + dir[0]
		col := colIdx + dir[1]

		if inBounds(content, row, col) && charIsSign(content[row][col]) {
			return true
		}
	}
	return false
}

func inBounds(content [][]string, row int, col int) bool {
	return row >= 0 && row < len(content) && col >= 0 && col < len(content[0])
}

func charIsSign(char string) bool {
	return !charIsNumber(char) && char != "."
}

func getDirections() [][]int {
	return [][]int{
		{-1, 0},
		{-1, 1},
		{0, 1},
		{1, 1},
		{1, 0},
		{1, -1},
		{0, -1},
		{-1, -1},
	}
}

func idxKey(row int, col int) string {
	return strings.Join([]string{strconv.Itoa(row), strconv.Itoa(col)}, "-")
}
