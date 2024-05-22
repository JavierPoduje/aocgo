package solvers

import (
	"reflect"
	"testing"
)

func TestSolverNine_SolveFirstProblem(t *testing.T) {
	solver := &SolverNine{}
	solver.content = []string{
		"0 3 6 9 12 15",
		"1 3 6 10 15 21",
		"10 13 16 21 30 45",
	}
	ans := solver.SolveFirstProblem()
	if ans != 114 {
		t.Errorf("Expected 114, got %d", ans)
	}
}

func TestSolverNine_SolveSecondProblem(t *testing.T) {
	solver := &SolverNine{}
	solver.content = []string{
		"0 3 6 9 12 15",
		"1 3 6 10 15 21",
		"10 13 16 21 30 45",
	}
	ans := solver.SolveSecondProblem()
	if ans != 2 {
		t.Errorf("Expected 2, got %d", ans)
	}
}

func TestSolverNine_extrapolateRow(t *testing.T) {
	extrapolatedRow, allZeroes := extrapolateRow([]int{0, 3, 6, 9, 12, 15})
	if !reflect.DeepEqual(extrapolatedRow, []int{3, 3, 3, 3, 3}) {
		t.Errorf("Expected [3,3,3,3,3], got %v", extrapolatedRow)
	}
	if allZeroes != false {
		t.Errorf("Expected false, got %v", allZeroes)
	}

	extrapolatedRow, allZeroes = extrapolateRow([]int{3, 3, 3, 3, 3})
	if !reflect.DeepEqual(extrapolatedRow, []int{0, 0, 0, 0}) {
		t.Errorf("Expected [0, 0, 0, 0], got %v", extrapolatedRow)
	}
	if allZeroes == true {
		t.Errorf("Expected true, got %v", allZeroes)
	}
}

func TestSolverNine_getPredictionByHistory(t *testing.T) {
	rowValue := getPredictionByHistory([]int{0, 3, 6, 9, 12, 15}, false)
	if rowValue != 18 {
		t.Errorf("Expected 18, got %d", rowValue)
	}
}

func TestSolverNine_getPredictionByExtrapolatedRow(t *testing.T) {
	rowValues := getPredictionByExtrapolatedRow([][]int{
		{0, 3, 6, 9, 12, 15},
		{3, 3, 3, 3, 3},
		{0, 0, 0, 0},
	}, false)
	if rowValues != 18 {
		t.Errorf("Expected 18, got %d", rowValues)
	}

	rowValues = getPredictionByExtrapolatedRow([][]int{
		{1, 3, 6, 10, 15, 21},
		{2, 3, 4, 5, 6},
		{1, 1, 1, 1},
		{0, 0, 0},
	}, false)
	if rowValues != 28 {
		t.Errorf("Expected 28, got %d", rowValues)
	}

	rowValues = getPredictionByExtrapolatedRow([][]int{
		{10, 13, 16, 21, 30, 45},
		{3, 3, 5, 9, 15},
		{0, 2, 4, 6},
		{2, 2, 2},
		{0, 0},
	}, false)
	if rowValues != 68 {
		t.Errorf("Expected 68, got %d", rowValues)
	}
}
