package types

type Solver interface {
	Parse(file string)
	SolveFirstProblem() int
	SolveSecondProblem() int
}
