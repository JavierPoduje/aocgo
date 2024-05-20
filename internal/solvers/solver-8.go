package solvers

import (
	"fmt"
	"log"
	"os"
	"strings"
)

type SolverEight struct {
	content      []string
	instructions []string
}

func (s *SolverEight) SolveFirstProblem() int {
	s.instructions = buildInstructions(s.content[0])
	nodesByName := buildNodes(s.content[2:])

	node := nodesByName["AAA"]
	instructionsIdx := 0
	steps := 0
	for node.Name != "ZZZ" {
		if instructionsIdx == len(s.instructions) {
			instructionsIdx = 0
		}

		instruction := s.instructions[instructionsIdx]

		if instruction == "L" {
			node = node.Left
		} else if instruction == "R" {
			node = node.Right
		} else {
			fmt.Println("Invalid instruction")
			log.Fatal()
		}

		steps++
		instructionsIdx++
	}

	return steps
}

func (s *SolverEight) SolveSecondProblem() int {
	return 0
}

func (s *SolverEight) Parse(file string) {
	content, err := os.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}

	rows := make([]string, 0)
	for _, line := range strings.Split(string(content), "\n") {
		trimmedLine := strings.TrimSpace(line)
		if trimmedLine == "" {
			continue
		}
		rows = append(rows, trimmedLine)
	}

	s.content = rows
}

func buildInstructions(row string) []string {
	intructions := make([]string, 0)
	for _, char := range row {
		intructions = append(intructions, string(char))
	}
	return intructions
}

func buildNodes(content []string) map[string]*Node {
	nodesByName := make(map[string]*Node)
	leftAndRightByNodeName := make(map[string][]string)

	for _, row := range content {
		splittedRow := strings.Split(row, " = ")
		branches := strings.Replace(splittedRow[1], "(", "", -1)
		branches = strings.Replace(branches, ")", "", -1)
		branches = strings.Replace(branches, ",", "", -1)
		parsedBranches := strings.Split(branches, " ")

		nodeName := splittedRow[0]
		left := parsedBranches[0]
		right := parsedBranches[1]

		nodesByName[nodeName] = NewNode(nodeName)
		leftAndRightByNodeName[nodeName] = []string{left, right}
	}

	for nodeName, leftAndRight := range leftAndRightByNodeName {
		currNode := nodesByName[nodeName]
		leftNode := nodesByName[leftAndRight[0]]
		rightNode := nodesByName[leftAndRight[1]]
		currNode.SetLeft(leftNode)
		currNode.SetRight(rightNode)
	}

	return nodesByName
}

// *** *** ***
// ** Node **
// *** *** ***

type Node struct {
	Name  string
	Left  *Node
	Right *Node
}

func NewNode(name string) *Node {
	return &Node{Name: name}
}

func (n *Node) SetLeft(node *Node) {
	n.Left = node
}

func (n *Node) SetRight(node *Node) {
	n.Right = node
}
