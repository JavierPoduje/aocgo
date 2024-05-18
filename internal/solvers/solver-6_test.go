package solvers

import (
	"testing"
)

func TestSolverSix_SolveFirstProblem(t *testing.T) {
	solver := &SolverSix{}
	solver.content = []string{
		"Time:      7  15   30",
		"Distance:  9  40  200",
	}
	ans := solver.SolveFirstProblem()
	if ans != 288 {
		t.Errorf("Expected 288, got %d", ans)
	}
}

func TestSolverSix_SolveSecondProblem(t *testing.T) {
	solver := &SolverSix{}
	solver.content = []string{
		"Time:      7  15   30",
		"Distance:  9  40  200",
	}
	ans := solver.SolveSecondProblem()
	if ans != 71503 {
		t.Errorf("Expected 71503, got %d", ans)
	}
}

func TestSolverSix_possibleRaceWins(t *testing.T) {
	ans := possibleRaceWins(7, 9)
	if ans != 4 {
		t.Errorf("Expected 4, got %d", ans)
	}
}
