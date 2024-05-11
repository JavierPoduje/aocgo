package solvers

import (
	"testing"
)

func TestSolverOne_SolveFirstProblem(t *testing.T) {
	solver := &SolverOne{}
	solver.content = [][]string{
		{"1", "a", "b", "c", "2"},
		{"p", "q", "r", "3", "s", "t", "u", "8", "v", "w", "x"},
		{"a", "1", "b", "2", "c", "3", "d", "4", "e", "5", "f"},
		{"t", "r", "e", "b", "7", "u", "c", "h", "e", "t"},
	}
	firstProblemAns := solver.SolveFirstProblem()
	if firstProblemAns != 142 {
		t.Errorf("Expected 142, got %d", firstProblemAns)
	}
}
