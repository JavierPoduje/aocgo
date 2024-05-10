package main

import (
	"fmt"
	"github.com/javierpoduje/aocgo/internal/solvers"
)

func main() {
	fmt.Println("from main")
	fromSolve := solvers.Solve()
	fmt.Printf("%s", fromSolve)
}
