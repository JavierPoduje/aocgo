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

func TestSolverEight_SolveSecondProblem(t *testing.T) {
	solver := &SolverEight{}
	solver.content = []string{
		"LR",
		"",
		"11A = (11B, XXX)",
		"11B = (XXX, 11Z)",
		"11Z = (11B, XXX)",
		"22A = (22B, XXX)",
		"22B = (22C, 22C)",
		"22C = (22Z, 22Z)",
		"22Z = (22B, 22B)",
		"XXX = (XXX, XXX)",
	}
	ans := solver.SolveSecondProblem()
	if ans != 6 {
		t.Errorf("Expected 6, got %d", ans)
	}
}
