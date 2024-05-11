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
	firstProblemAnswer := solver.SolveFirstProblem()
	fmt.Println(firstProblemAnswer)
}
