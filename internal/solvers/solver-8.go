package solvers

import (
	"fmt"
	"log"
	"os"
	"strings"
)

type SolverEight struct {
	content []string
	cmds    []string
}

func (s *SolverEight) SolveFirstProblem() int {
	s.cmds = buildCommands(s.content[0])
	nodesByName := buildNodes(s.content[2:])

	node := nodesByName["AAA"]
	instructionsIdx := 0
	steps := 0
	for node.Name != "ZZZ" {
		if instructionsIdx == len(s.cmds) {
			instructionsIdx = 0
		}

		instruction := s.cmds[instructionsIdx]

		if instruction == "L" {
			node = nodesByName[node.Left]
		} else if instruction == "R" {
			node = nodesByName[node.Right]
		}

		steps++
		instructionsIdx++
	}

	return steps
}

func (s *SolverEight) SolveSecondProblem() int {
	s.cmds = buildCommands(s.content[0])
	nodesByName := buildNodes(s.content[1:])

	nodesToTraverse := getStartNodes(nodesByName)

	stepsInBetweenZsPerNode := make([][]int, len(nodesToTraverse))
	for nodeIdx, nodeName := range nodesToTraverse {
		node := nodesByName[nodeName]
		steps := 0
		cmdIdx := 0
		firstZNodeName := ""
		stepsInBetweenZs := make([]int, 0)

		for {
			if cmdIdx == len(s.cmds) {
				cmdIdx = 0
			}

			cmd := s.cmds[cmdIdx]

			if cmd == "L" {
				node = nodesByName[node.Left]
			} else if cmd == "R" {
				node = nodesByName[node.Right]
			}

			steps++
			cmdIdx++

			if node.IsEndNode() {
				stepsInBetweenZs = append(stepsInBetweenZs, steps)
				steps = 0
				if firstZNodeName == "" {
					firstZNodeName = node.Name
				} else if node.Name == firstZNodeName {
					break
				}
			}
		}

		stepsInBetweenZsPerNode[nodeIdx] = stepsInBetweenZs
	}

	// take the first value of each `stepsInBetweenZsPerNode`
	nums := make([]int, len(stepsInBetweenZsPerNode))
	for i, stepsInBetweenZs := range stepsInBetweenZsPerNode {
		nums[i] = stepsInBetweenZs[0]
	}

	return lcm(nums)
}

func lcm(nums []int) int {
	gcd := func(a, b int) int {
		for b != 0 {
			a, b = b, a%b
		}
		return a

	}

	lcm := nums[0]
	for i := 1; i < len(nums); i++ {
		lcm = lcm * nums[i] / gcd(lcm, nums[i])
	}
	return lcm
}

func getStartNodes(nodesByName map[string]Node) []string {
	startNodes := make([]string, 0)
	for _, node := range nodesByName {
		if node.IsStartNode() {
			startNodes = append(startNodes, node.Name)
		}
	}
	return startNodes
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

func buildCommands(row string) []string {
	intructions := make([]string, 0)
	for _, char := range row {
		intructions = append(intructions, string(char))
	}
	return intructions
}

func buildNodes(content []string) map[string]Node {
	nodesByName := make(map[string]Node)

	for _, row := range content {
		if row == "" {
			continue
		}

		splittedRow := strings.Split(row, " = ")
		branches := strings.Replace(splittedRow[1], "(", "", -1)
		branches = strings.Replace(branches, ")", "", -1)
		branches = strings.Replace(branches, ",", "", -1)
		parsedBranches := strings.Split(branches, " ")

		nodeName := splittedRow[0]
		left := parsedBranches[0]
		right := parsedBranches[1]

		nodesByName[nodeName] = NewNode(nodeName, left, right)
	}

	return nodesByName
}

// *** *** ***
// ** Node **
// *** *** ***

type Node struct {
	Name  string
	Left  string
	Right string
}

func NewNode(name string, left string, right string) Node {
	return Node{
		Name:  name,
		Left:  left,
		Right: right,
	}
}

func (n *Node) SetLeft(nodeName string) {
	n.Left = nodeName
}

func (n *Node) SetRight(nodeName string) {
	n.Right = nodeName
}

func (n Node) IsStartNode() bool {
	return n.Name[2] == 'A'
}

func (n Node) IsEndNode() bool {
	return n.Name[2] == 'Z'
}

func (n Node) String() string {
	return fmt.Sprintf("\tName: %s, Left: %s, Right: %s\n", n.Name, n.Left, n.Right)
}
