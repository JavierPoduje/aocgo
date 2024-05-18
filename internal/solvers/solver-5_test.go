package solvers

import (
	"testing"
)

func TestSolverFive_SolveFirstProblem(t *testing.T) {
	solver := &SolverFive{}
	solver.content = []string{
		"seeds: 79 14 55 13",
		"",
		"seed-to-soil map:",
		"50 98 2",
		"52 50 48",
		"",
		"soil-to-fertilizer map:",
		"0 15 37",
		"37 52 2",
		"39 0 15",
		"",
		"fertilizer-to-water map:",
		"49 53 8",
		"0 11 42",
		"42 0 7",
		"57 7 4",
		"",
		"water-to-light map:",
		"88 18 7",
		"18 25 70",
		"",
		"light-to-temperature map:",
		"45 77 23",
		"81 45 19",
		"68 64 13",
		"",
		"temperature-to-humidity map:",
		"0 69 1",
		"1 0 69",
		"",
		"humidity-to-location map:",
		"60 56 37",
		"56 93 4",
	}
	firstProblemAns := solver.SolveFirstProblem()
	if firstProblemAns != 35 {
		t.Errorf("Expected 35, got %d", firstProblemAns)
	}
}

func TestFromToRange_Get(t *testing.T) {
	fromToRange := FromToRange{from: 10, to: 20, size: 5}

	mappedValue, found := fromToRange.Get(12)
	if mappedValue != 22 {
		t.Errorf("Expected 22, got %d", mappedValue)
	}
	if !found {
		t.Errorf("expected found, got %v", found)
	}

	mappedValue, found = fromToRange.Get(9)
	if mappedValue != 9 {
		t.Errorf("Expected 9, got %d", mappedValue)
	}
	if found {
		t.Errorf("expected !found, got %v", found)
	}

	mappedValue, found = fromToRange.Get(22)
	if mappedValue != 22 {
		t.Errorf("Expected 22, got %d", mappedValue)
	}
	if !found {
		t.Errorf("expected found, got %v", found)
	}

	mappedValue, found = fromToRange.Get(15)
	if mappedValue != 25 {
		t.Errorf("Expected 25, got %d", mappedValue)
	}
	if !found {
		t.Errorf("expected found, got %v", found)
	}

	mappedValue, found = fromToRange.Get(20)
	if mappedValue != 20 {
		t.Errorf("Expected 20, got %d", mappedValue)
	}
	if found {
		t.Errorf("expected !found, got %v", found)
	}

	fromToRange = FromToRange{from: 50, to: 52, size: 48}
	mappedValue, found = fromToRange.Get(79)
	if mappedValue != 81 {
		t.Errorf("Expected 81, got %d", mappedValue)
	}
	if !found {
		t.Errorf("expected found, got %v", found)
	}
}
