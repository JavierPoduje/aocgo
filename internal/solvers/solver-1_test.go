package solvers

import (
	"testing"
)

func Test_SolveFirstProblem(t *testing.T) {
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

func Test_SolveSecondProblem(t *testing.T) {
	solver := &SolverOne{}
	solver.content = [][]string{
		{"t", "w", "o", "1", "n", "i", "n", "e"},
		{"e", "i", "g", "h", "t", "w", "o", "t", "h", "r", "e", "e"},
		{"a", "b", "c", "o", "n", "e", "2", "t", "h", "r", "e", "e", "x", "y", "z"},
		{"x", "t", "w", "o", "n", "e", "3", "f", "o", "u", "r"},
		{"4", "n", "i", "n", "e", "e", "i", "g", "h", "t", "s", "e", "v", "e", "n", "2"},
		{"z", "o", "n", "e", "i", "g", "h", "t", "2", "3", "4"},
		{"7", "p", "q", "r", "s", "t", "s", "i", "x", "t", "e", "e", "n"},
	}
	scdProblemAns := solver.SolveSecondProblem()
	if scdProblemAns != 281 {
		t.Errorf("Expected 281, got %d", scdProblemAns)
	}
}

func Test_processRow(t *testing.T) {
	numbersTrie := CreateNumbersTrie()
	rowNum := processRow([]string{"t", "w", "o", "1", "n", "i", "n", "e"}, numbersTrie)
	if rowNum != 29 {
		t.Errorf("Expected 29, got %d", rowNum)
	}
	rowNum = processRow([]string{"e", "i", "g", "h", "t", "w", "o", "t", "h", "r", "e", "e"}, numbersTrie)
	if rowNum != 83 {
		t.Errorf("Expected 83, got %d", rowNum)
	}
	rowNum = processRow([]string{"a", "b", "c", "o", "n", "e", "2", "t", "h", "r", "e", "e", "x", "y", "z"}, numbersTrie)
	if rowNum != 13 {
		t.Errorf("Expected 13, got %d", rowNum)
	}
	rowNum = processRow([]string{"x", "t", "w", "o", "n", "e", "3", "f", "o", "u", "r"}, numbersTrie)
	if rowNum != 24 {
		t.Errorf("Expected 24, got %d", rowNum)
	}
	rowNum = processRow([]string{"4", "n", "i", "n", "e", "e", "i", "g", "h", "t", "s", "e", "v", "e", "n", "2"}, numbersTrie)
	if rowNum != 42 {
		t.Errorf("Expected 42, got %d", rowNum)
	}
	rowNum = processRow([]string{"z", "o", "n", "e", "i", "g", "h", "t", "2", "3", "4"}, numbersTrie)
	if rowNum != 14 {
		t.Errorf("Expected 14, got %d", rowNum)
	}
	rowNum = processRow([]string{"7", "p", "q", "r", "s", "t", "s", "i", "x", "t", "e", "e", "n"}, numbersTrie)
	if rowNum != 76 {
		t.Errorf("Expected 76, got %d", rowNum)
	}
}
