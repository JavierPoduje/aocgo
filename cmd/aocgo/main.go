package main

import (
	"fmt"
	"github.com/javierpoduje/aocgo/internal/solvers"
)

const (
	fileOne = "./internal/text-files/1.txt"
)

func main() {
	solver := solvers.GetSolverByNumber(1)

	solver.Parse(fileOne)

	fstProblemAnswer := solver.SolveFirstProblem()
	scdProblemAnswer := solver.SolveSecondProblem()

	fmt.Printf("First problem answer: %d\n", fstProblemAnswer)
	fmt.Printf("Second problem answer: %d\n", scdProblemAnswer)
}
