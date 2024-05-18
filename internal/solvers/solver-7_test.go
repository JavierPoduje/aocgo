package solvers

import (
	"testing"
)

func TestSolverSeven_SolveFirstProblem(t *testing.T) {
	solver := &SolverSeven{}
	solver.content = []string{
		"32T3K 765",
		"T55J5 684",
		"KK677 28",
		"KTJJT 220",
		"QQQJA 483",
	}
	ans := solver.SolveFirstProblem()
	if ans != 6440 {
		t.Errorf("Expected 6440, got %d", ans)
	}
}
