package solvers

import (
	"log"

	"github.com/javierpoduje/aocgo/internal/types"
)

func GetSolverByNumber(solverNumber int) (solver types.Solver, filename string) {
	const (
		fileOne   = "./internal/text-files/1.txt"
		fileTwo   = "./internal/text-files/2.txt"
		fileThree = "./internal/text-files/3.txt"
		fileFour  = "./internal/text-files/4.txt"
		fileFive  = "./internal/text-files/5.txt"
		fileSix   = "./internal/text-files/6.txt"
		fileSeven = "./internal/text-files/7.txt"
	)

	switch solverNumber {
	case 1:
		return &SolverOne{}, fileOne
	case 2:
		return &SolverTwo{}, fileTwo
	case 3:
		return &SolverThree{}, fileThree
	case 4:
		return &SolverFour{}, fileFour
	case 5:
		return &SolverFive{}, fileFive
	case 6:
		return &SolverSix{}, fileSix
	case 7:
		return &SolverSeven{}, fileSeven
	default:
		log.Fatal("Invalid solver number")
	}

	return nil, ""
}
