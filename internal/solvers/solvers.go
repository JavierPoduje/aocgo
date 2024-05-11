package solvers

import (
	"github.com/javierpoduje/aocgo/internal/types"
)

func GetSolverByNumber(solverNumber int) types.Solver {
	switch solverNumber {
	case 1:
		return &SolverOne{}
	default:
		return &SolverOne{}
	}
}
