package solvers

import (
	"testing"
)

func TestSolverEight_SolveFirstProblem(t *testing.T) {
	solver := &SolverEight{}
	solver.content = []string{
		"RL",
		"",
		"AAA = (BBB, CCC)",
		"BBB = (DDD, EEE)",
		"CCC = (ZZZ, GGG)",
		"DDD = (DDD, DDD)",
		"EEE = (EEE, EEE)",
		"GGG = (GGG, GGG)",
		"ZZZ = (ZZZ, ZZZ)",
	}
	ans := solver.SolveFirstProblem()
	if ans != 2 {
		t.Errorf("Expected 2, got %d", ans)
	}

	solver.content = []string{
		"LLR",
		"",
		"AAA = (BBB, BBB)",
		"BBB = (AAA, ZZZ)",
		"ZZZ = (ZZZ, ZZZ)",
	}

	ans = solver.SolveFirstProblem()
	if ans != 6 {
		t.Errorf("Expected 6, got %d", ans)
	}
}
