package solvers

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	ds "github.com/javierpoduje/aocgo/internal/data-structures"
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
	numbersTrie := CreateNumbersTrie()
	numberRows := make([]int, 0)

	for _, row := range s.content {
		number := processRow(row, numbersTrie)
		numberRows = append(numberRows, number)
	}

	return sumSliceOfNumbers(numberRows)
}

func processRow(row []string, numbersTrie *ds.Trie) int {
	leftNum := ""
	rightNum := ""

	left := 0
	right := 0

	fmt.Printf("Row: %v\n", row)

	for {
		if left >= len(row) || right >= len(row) {
			break
		}

		if right == len(row) && left < right {
			numberAsWord := strings.Join(row[left:right], "")
			if numbersTrie.Search(numberAsWord) {
				if leftNum == "" {
					leftNum = numberAsWordToChar(numberAsWord)
					rightNum = numberAsWordToChar(numberAsWord)
				} else {
					rightNum = numberAsWordToChar(numberAsWord)
				}
			}
			break
		}

		if left == right {
			char := row[left]

			if charIsNumber(char) {
				if leftNum == "" {
					leftNum = char
					rightNum = char
				} else {
					rightNum = char
				}
				left++
				right++
				continue
			} else if numbersTrie.StartsWith(char) {
				right++
				continue
			} else {
				left++
				right++
				continue
			}
		} else {
			numberAsWord := strings.Join(row[left:right], "")
			if numbersTrie.Search(numberAsWord) {
				if leftNum == "" {
					leftNum = numberAsWordToChar(numberAsWord)
					rightNum = numberAsWordToChar(numberAsWord)
				} else {
					rightNum = numberAsWordToChar(numberAsWord)
				}
				right++
				left = right
				continue
			} else if numbersTrie.StartsWith(strings.Join(row[left:right], "")) {
				right++
				continue
			} else {
				left++
				right++
				continue
			}
		}
	}

	number := joinAndParse([]string{leftNum, rightNum})

	fmt.Printf("number: %v\n", number)

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

func CreateNumbersTrie() *ds.Trie {
	trie := ds.NewTrie()
	trie.InsertMany([]string{
		"one",
		"two",
		"three",
		"four",
		"five",
		"six",
		"seven",
		"eight",
		"nine",
	})
	return trie
}

func numberAsWordToChar(word string) string {
	switch word {
	case "one":
		return "1"
	case "two":
		return "2"
	case "three":
		return "3"
	case "four":
		return "4"
	case "five":
		return "5"
	case "six":
		return "6"
	case "seven":
		return "7"
	case "eight":
		return "8"
	case "nine":
		return "9"
	default:
		return ""
	}
}
