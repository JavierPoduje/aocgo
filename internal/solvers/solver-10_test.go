package solvers

import (
	"testing"
)

func TestSolverTen_SolveFirstProblem(t *testing.T) {
	solver := &SolverTen{}
	solver.grid = [][]rune{
		{'.', '.', '.', '.', '.'},
		{'.', 'S', '-', '7', '.'},
		{'.', '|', '.', '|', '.'},
		{'.', 'L', '-', 'J', '.'},
		{'.', '.', '.', '.', '.'},
	}

	ans := solver.SolveFirstProblem()
	if ans != 4 {
		t.Errorf("Expected 4, got %d", ans)
	}

	solver.grid = [][]rune{
		{'7', '-', 'F', '7', '-'},
		{'.', 'F', 'J', '|', '7'},
		{'S', 'J', 'L', 'L', '7'},
		{'|', 'F', '-', '-', 'J'},
		{'L', 'J', '.', 'L', 'J'},
	}
	ans = solver.SolveFirstProblem()
	if ans != 8 {
		t.Errorf("Expected 8, got %d", ans)
	}
}

func TestSolverTen_startingPoint(t *testing.T) {
	solver := &SolverTen{}
	solver.grid = [][]rune{
		{'.', '.', '.', '.', '.'},
		{'.', 'S', '-', '7', '.'},
		{'.', '|', '.', '|', '.'},
		{'.', 'L', '-', 'J', '.'},
		{'.', '.', '.', '.', '.'},
	}

	rowIdx, colIdx := startingPoint(solver.grid)
	if rowIdx != 1 || colIdx != 1 {
		t.Errorf("Expected (1, 1), got (%d, %d)", rowIdx, colIdx)
	}
}

func TestSolverTen_findFarthestDistance(t *testing.T) {
	solver := &SolverTen{}
	solver.grid = [][]rune{
		{'.', '.', '.', '.', '.'},
		{'.', 'S', '-', '7', '.'},
		{'.', '|', '.', '|', '.'},
		{'.', 'L', '-', 'J', '.'},
		{'.', '.', '.', '.', '.'},
	}

	distance := solver.findFarthestDistance([]int{1, 2}, []int{2, 1}, []int{1, 1})
	if distance != 4 {
		t.Errorf("Expected 4, got %d", distance)
	}

	solver.grid = [][]rune{
		{'-', 'L', '|', 'F', '7'},
		{'7', 'S', '-', '7', '|'},
		{'L', '|', '7', '|', '|'},
		{'-', 'L', '-', 'J', '|'},
		{'L', '|', '-', 'J', 'F'},
	}

	distance = solver.findFarthestDistance([]int{1, 2}, []int{2, 1}, []int{1, 1})
	if distance != 4 {
		t.Errorf("Expected 4, got %d", distance)
	}
}

func TestSolverTen_getNextCoord(t *testing.T) {
	solver := &SolverTen{}
	solver.grid = [][]rune{
		{'.', '.', '.', '.', '.'},
		{'.', 'S', '-', '7', '.'},
		{'.', '|', '.', '|', '.'},
		{'.', 'L', '-', 'J', '.'},
		{'.', '.', '.', '.', '.'},
	}

	curr, prev := solver.getNextCoord([]int{1, 2}, []int{1, 1})

	if prev[0] != 1 || prev[1] != 2 {
		t.Errorf("Expected (1, 2), got (%d, %d)", prev[0], prev[1])
	}
	if curr[0] != 1 || curr[1] != 3 {
		t.Errorf("Expected (1, 3), got (%d, %d)", curr[0], curr[1])
	}

	curr, prev = solver.getNextCoord([]int{1, 3}, []int{1, 2})
	if prev[0] != 1 || prev[1] != 3 {
		t.Errorf("Expected (1, 3), got (%d, %d)", prev[0], prev[1])
	}
	if curr[0] != 2 || curr[1] != 3 {
		t.Errorf("Expected (2, 3), got (%d, %d)", curr[0], curr[1])
	}
}
