package solvers

import (
	"log"
	"os"
	"strconv"
	"strings"
)

type SolverOne struct {
	content [][]string
}

func (s *SolverOne) SolveFirstProblem() int {
	numberRows := make([]int, 0)

	for _, row := range s.content {
		rowNumbers := make([]string, len(row))

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

	return sumSliceOfNumbers(numberRows)
}

func (s *SolverOne) SolveSecondProblem() int {
	numberRows := make([]int, 0)

	for _, row := range s.content {
		number := processRow(row)
		numberRows = append(numberRows, number)
	}

	return sumSliceOfNumbers(numberRows)
}

func processRow(row []string) int {
	leftNum := ""
	rightNum := ""

	for i := 0; i < len(row); i++ {
		char := row[i]
		word := strings.Join(row[i:], "")

		if charIsNumber(char) {
			if leftNum == "" {
				leftNum = char
			}
			rightNum = char
		} else if has, numberPrefix := wordStartsWithNumber(word); has {
			if leftNum == "" {
				leftNum = numberPrefix
			}
			rightNum = numberPrefix
		}

	}

	number := joinAndParse([]string{leftNum, rightNum})

	return number
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

func sumSliceOfNumbers(numbers []int) int {
	sum := 0
	for _, num := range numbers {
		sum += num
	}
	return sum
}

func wordStartsWithNumber(word string) (bool, string) {
	if strings.HasPrefix(word, "one") {
		return true, "1"
	} else if strings.HasPrefix(word, "two") {
		return true, "2"
	} else if strings.HasPrefix(word, "three") {
		return true, "3"
	} else if strings.HasPrefix(word, "four") {
		return true, "4"
	} else if strings.HasPrefix(word, "five") {
		return true, "5"
	} else if strings.HasPrefix(word, "six") {
		return true, "6"
	} else if strings.HasPrefix(word, "seven") {
		return true, "7"
	} else if strings.HasPrefix(word, "eight") {
		return true, "8"
	} else if strings.HasPrefix(word, "nine") {
		return true, "9"
	}

	return false, ""
}
