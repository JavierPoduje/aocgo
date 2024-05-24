package solvers

import (
	"reflect"
	"testing"
)

func TestSolverEleven_SolveFirstProblem(t *testing.T) {
	solver := &SolverEleven{}
	solver.grid = Grid{
		{'.', '.', '.', '#', '.', '.', '.', '.', '.', '.'},
		{'.', '.', '.', '.', '.', '.', '.', '#', '.', '.'},
		{'#', '.', '.', '.', '.', '.', '.', '.', '.', '.'},
		{'.', '.', '.', '.', '.', '.', '.', '.', '.', '.'},
		{'.', '.', '.', '.', '.', '.', '#', '.', '.', '.'},
		{'.', '#', '.', '.', '.', '.', '.', '.', '.', '.'},
		{'.', '.', '.', '.', '.', '.', '.', '.', '.', '#'},
		{'.', '.', '.', '.', '.', '.', '.', '.', '.', '.'},
		{'.', '.', '.', '.', '.', '.', '.', '#', '.', '.'},
		{'#', '.', '.', '.', '#', '.', '.', '.', '.', '.'},
	}

	ans := solver.SolveFirstProblem()
	if ans != 374 {
		t.Errorf("Expected 374, got %d", ans)
	}
}

func TestSolverEleven_SolveSecondProblem(t *testing.T) {
	solver := &SolverEleven{
		adder: 99,
	}
	solver.grid = Grid{
		{'.', '.', '.', '#', '.', '.', '.', '.', '.', '.'},
		{'.', '.', '.', '.', '.', '.', '.', '#', '.', '.'},
		{'#', '.', '.', '.', '.', '.', '.', '.', '.', '.'},
		{'.', '.', '.', '.', '.', '.', '.', '.', '.', '.'},
		{'.', '.', '.', '.', '.', '.', '#', '.', '.', '.'},
		{'.', '#', '.', '.', '.', '.', '.', '.', '.', '.'},
		{'.', '.', '.', '.', '.', '.', '.', '.', '.', '#'},
		{'.', '.', '.', '.', '.', '.', '.', '.', '.', '.'},
		{'.', '.', '.', '.', '.', '.', '.', '#', '.', '.'},
		{'#', '.', '.', '.', '#', '.', '.', '.', '.', '.'},
	}

	ans := solver.SolveSecondProblem()
	if ans != 8410 {
		t.Errorf("Expected 8410, got %d", ans)
	}
}

func TestSolverEleven_expandGrid(t *testing.T) {
	solver := &SolverEleven{}
	solver.grid = Grid{
		{'.', '.', '#'},
		{'.', '.', '.'},
		{'#', '.', '.'},
	}

	expandedGrid := solver.expandGrid()
	expectedGrid := Grid{
		{'.', '.', '.', '#'},
		{'.', '.', '.', '.'},
		{'.', '.', '.', '.'},
		{'#', '.', '.', '.'},
	}

	if !reflect.DeepEqual(expandedGrid, expectedGrid) {
		t.Errorf("Expected another expanded-grid %d", expandedGrid)
	}
}

func TestSolverEleven_findPairs(t *testing.T) {
	s := &SolverEleven{}
	pairs := s.findPairs(Grid{
		{'#', '.', '.'},
		{'.', '#', '.'},
		{'.', '.', '#'},
	})
	expectedPairs := [][][]int{
		{{0, 0}, {1, 1}},
		{{0, 0}, {2, 2}},
		{{1, 1}, {2, 2}},
	}

	if !reflect.DeepEqual(pairs, expectedPairs) {
		t.Errorf(`Expected {
			{ {0,0},{1,1} }
			{ {0,0},{2,2} }
			{ {1,1},{2,2} }
		}, got %d`, pairs)
	}
}

func TestSolverEleven_findHashes(t *testing.T) {
	s := &SolverEleven{}
	hashes := s.findHashes(Grid{
		{'.', '.', '#'},
		{'.', '.', '.'},
		{'#', '.', '.'},
	})
	expectedHahes := [][]int{
		{0, 2},
		{2, 0},
	}

	if !reflect.DeepEqual(hashes, expectedHahes) {
		t.Errorf("Expected [][]{ {0,2}, {2,0} }, got %d", hashes)
	}
}

func TestSolverEleven_distance(t *testing.T) {
	s := &SolverEleven{}

	distance := s.distance([]int{0, 0}, []int{1, 2})
	if distance != 3 {
		t.Errorf("Expected 3, got %d", distance)
	}

	distance = s.distance([]int{0, 0}, []int{1, 0})
	if distance != 1 {
		t.Errorf("Expected 1, got %d", distance)
	}

	distance = s.distance([]int{0, 0}, []int{2, 2})
	if distance != 4 {
		t.Errorf("Expected 4, got %d", distance)
	}

	distance = s.distance([]int{6, 1}, []int{11, 5})
	if distance != 9 {
		t.Errorf("Expected 9, got %d", distance)
	}
}
