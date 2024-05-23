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

func TestSolverTen_SolveSecondProblem(t *testing.T) {
	solver := &SolverTen{}
	solver.grid = [][]rune{
		{'.', '.', '.', '.', '.'},
		{'.', 'S', '-', '7', '.'},
		{'.', '|', '.', '|', '.'},
		{'.', 'L', '-', 'J', '.'},
		{'.', '.', '.', '.', '.'},
	}

	ans := solver.SolveSecondProblem()
	if ans != 4 {
		t.Errorf("Expected 4, got %d", ans)
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

	distance, _ := solver.traverseAndFindFarthestDistance([]int{1, 2}, []int{2, 1}, []int{1, 1})
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

	distance, _ = solver.traverseAndFindFarthestDistance([]int{1, 2}, []int{2, 1}, []int{1, 1})
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

func TestSolverTen_countRowInversions(t *testing.T) {
	solver := &SolverTen{}
	solver.grid = [][]rune{
		{'.', '.', '.', '.', '.'},
		{'.', 'S', '-', '7', '.'},
		{'.', '|', '.', '|', '.'},
		{'.', 'L', '-', 'J', '.'},
		{'.', '.', '.', '.', '.'},
	}
	rowIdx, colIdx := startingPoint(solver.grid)
	head, tail := solver.findHeadAndTail(rowIdx, colIdx)
	_, visited := solver.traverseAndFindFarthestDistance(head, tail, []int{rowIdx, colIdx})

	inversions := solver.countRowInversions(2, 4, visited)
	if inversions != 2 {
		t.Errorf("Expected 2, got %d", inversions)
	}

	inversions = solver.countRowInversions(2, 1, visited)
	if inversions != 0 {
		t.Errorf("Expected 0, got %d", inversions)
	}

	inversions = solver.countRowInversions(3, 2, visited)
	if inversions != 0 {
		t.Errorf("Expected 0, got %d", inversions)
	}

	inversions = solver.countRowInversions(3, 4, visited)
	if inversions != 2 {
		t.Errorf("Expected 0, got %d", inversions)
	}
}

func TestSolverTen_countInnerTiles(t *testing.T) {
	solver := &SolverTen{}
	solver.grid = [][]rune{
		{'.', '.', '.', '.', '.'},
		{'.', 'S', '-', '7', '.'},
		{'.', '|', '.', '|', '.'},
		{'.', 'L', '-', 'J', '.'},
		{'.', '.', '.', '.', '.'},
	}
	rowIdx, colIdx := startingPoint(solver.grid)
	head, tail := solver.findHeadAndTail(rowIdx, colIdx)
	_, visited := solver.traverseAndFindFarthestDistance(head, tail, []int{rowIdx, colIdx})
	tiles := solver.countInnerTiles(visited)
	if tiles != 1 {
		t.Errorf("Expected 1, got %d", tiles)
	}

	solver.grid = [][]rune{
		{'.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.'},
		{'.', 'S', '-', '-', '-', '-', '-', '-', '-', '7', '.'},
		{'.', '|', 'F', '-', '-', '-', '-', '-', '7', '|', '.'},
		{'.', '|', '|', '.', '.', '.', '.', '.', '|', '|', '.'},
		{'.', '|', '|', '.', '.', '.', '.', '.', '|', '|', '.'},
		{'.', '|', 'L', '-', '7', '.', 'F', '-', 'J', '|', '.'},
		{'.', '|', '.', '.', '|', '.', '|', '.', '.', '|', '.'},
		{'.', 'L', '-', '-', 'J', '.', 'L', '-', '-', 'J', '.'},
		{'.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.'},
	}
	rowIdx, colIdx = startingPoint(solver.grid)
	head, tail = solver.findHeadAndTail(rowIdx, colIdx)
	_, visited = solver.traverseAndFindFarthestDistance(head, tail, []int{rowIdx, colIdx})
	tiles = solver.countInnerTiles(visited)
	if tiles != 4 {
		t.Errorf("Expected 4, got %d", tiles)
	}

	solver.grid = [][]rune{
		{'.', '.', '.', '.', '.', '.', '.', '.', '.', '.'},
		{'.', 'S', '-', '-', '-', '-', '-', '-', '7', '.'},
		{'.', '|', 'F', '-', '-', '-', '-', '7', '|', '.'},
		{'.', '|', '|', 'O', 'O', 'O', 'O', '|', '|', '.'},
		{'.', '|', '|', 'O', 'O', 'O', 'O', '|', '|', '.'},
		{'.', '|', 'L', '-', '7', 'F', '-', 'J', '|', '.'},
		{'.', '|', 'I', 'I', '|', '|', 'I', 'I', '|', '.'},
		{'.', 'L', '-', '-', 'J', 'L', '-', '-', 'J', '.'},
		{'.', '.', '.', '.', '.', '.', '.', '.', '.', '.'},
	}
	rowIdx, colIdx = startingPoint(solver.grid)
	head, tail = solver.findHeadAndTail(rowIdx, colIdx)
	_, visited = solver.traverseAndFindFarthestDistance(head, tail, []int{rowIdx, colIdx})
	tiles = solver.countInnerTiles(visited)
	if tiles != 4 {
		t.Errorf("Expected 4, got %d", tiles)
	}

	solver.grid = [][]rune{
		{'F', 'F', '7', 'F', 'S', 'F', '7', 'F', '7', 'F', '7', 'F', '7', 'F', '7', 'F', '-', '-', '-', '7'},
		{'L', '|', 'L', 'J', '|', '|', '|', '|', '|', '|', '|', '|', '|', '|', '|', '|', 'F', '-', '-', 'J'},
		{'F', 'L', '-', '7', 'L', 'J', 'L', 'J', '|', '|', '|', '|', '|', '|', 'L', 'J', 'L', '-', '7', '7'},
		{'F', '-', '-', 'J', 'F', '-', '-', '7', '|', '|', 'L', 'J', 'L', 'J', '7', 'F', '7', 'F', 'J', '-'},
		{'L', '-', '-', '-', 'J', 'F', '-', 'J', 'L', 'J', '.', '|', '|', '-', 'F', 'J', 'L', 'J', 'J', '7'},
		{'|', 'F', '|', 'F', '-', 'J', 'F', '-', '-', '-', '7', 'F', '7', '-', 'L', '7', 'L', '|', '7', '|'},
		{'|', 'F', 'F', 'J', 'F', '7', 'L', '7', 'F', '-', 'J', 'F', '7', '|', 'J', 'L', '-', '-', '-', '7'},
		{'7', '-', 'L', '-', 'J', 'L', '7', '|', '|', 'F', '7', '|', 'L', '7', 'F', '-', '7', 'F', '7', '|'},
		{'L', '.', 'L', '7', 'L', 'F', 'J', '|', '|', '|', '|', '|', 'F', 'J', 'L', '7', '|', '|', 'L', 'J'},
		{'L', '7', 'J', 'L', 'J', 'L', '-', 'J', 'L', 'J', 'L', 'J', 'L', '-', '-', 'J', 'L', 'J', '.', 'L'},
	}
	rowIdx, colIdx = startingPoint(solver.grid)
	head, tail = solver.findHeadAndTail(rowIdx, colIdx)
	_, visited = solver.traverseAndFindFarthestDistance(head, tail, []int{rowIdx, colIdx})
	tiles = solver.countInnerTiles(visited)
	if tiles != 10 {
		t.Errorf("Expected 10, got %d", tiles)
	}
}
