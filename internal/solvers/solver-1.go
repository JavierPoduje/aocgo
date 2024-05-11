package solvers

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type SolverOne struct {
	content [][]string
}

func (s *SolverOne) SolveFirstProblem() int {
	fmt.Println(s.content)

	numberRows := make([]int, 0)

	for _, row := range s.content {
		rowNumbers := make([]string, 0)

		// from left to right
		for _, char := range row {
			if charIsNumber(char) {
				rowNumbers = append(rowNumbers, char)
				break
			}
		}
		// from right to left
		for i := len(row) - 1; i >= 0; i-- {
			char := row[i]
			if charIsNumber(char) {
				rowNumbers = append(rowNumbers, char)
				break
			}
		}

		numberRows = append(numberRows, joinAndParse(rowNumbers))
	}

	// sum all numbers
	sum := 0
	for _, num := range numberRows {
		sum += num
	}

	return sum
}

func (s *SolverOne) SolveSecondProblem() int {
	return 0
}

func (s *SolverOne) Parse(file string) {
	content, err := os.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}

	splittedByChar := make([][]string, 0)
	for _, line := range strings.Split(string(content), "\n") {
		if line == "" {
			continue
		}
		lineContent := strings.Split(line, "")
		splittedByChar = append(splittedByChar, lineContent)
	}

	s.content = splittedByChar
}

func joinAndParse(numbers []string) int {
	joinedNumbers := strings.Join(numbers, "")
	parsedNumber, err := strconv.Atoi(joinedNumbers)
	if err != nil {
		log.Fatalf("Joined numbers are not really 'numbers': %s", err)
	}
	return parsedNumber
}

func charIsNumber(char string) bool {
	_, err := strconv.Atoi(char)
	return err == nil
}
