package main

import (
	"flag"
	"fmt"

	"github.com/javierpoduje/aocgo/internal/solvers"
)

func main() {
	solverArg := flag.Int("solver", 1, "The solver number")
	flag.Parse()

	solverNumber := *solverArg

	solver, filename := solvers.GetSolverByNumber(solverNumber)

	solver.Parse(filename)
	fstProblemAnswer := solver.SolveFirstProblem()
	scdProblemAnswer := solver.SolveSecondProblem()

	fmt.Printf("First problem answer: %d\n", fstProblemAnswer)
	fmt.Printf("Second problem answer: %d\n", scdProblemAnswer)
}
