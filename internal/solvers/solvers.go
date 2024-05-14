package solvers

import (
	"github.com/javierpoduje/aocgo/internal/types"
)

func GetSolverByNumber(solverNumber int) (solver types.Solver, filename string) {
	const (
		fileOne = "./internal/text-files/1.txt"
		fileTwo = "./internal/text-files/2.txt"
	)

	switch solverNumber {
	case 1:
		return &SolverOne{}, fileOne
	case 2:
		return &SolverTwo{}, fileTwo
	default:
		return &SolverOne{}, fileOne
	}
}
