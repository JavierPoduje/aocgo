package solvers

import (
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

type SolverFour struct {
	content []string
}

type Scratchcard struct {
	id    int
	wins  []int
	goten []int
}

func (s *Scratchcard) CalculateValue() int {
	numberOfWins := s.CalculateNumberOfWins()
	zero := float64(0)
	pow := math.Pow(float64(2), float64(numberOfWins-1))
	maxValue := math.Max(zero, pow)
	return int(maxValue)
}

func (s *Scratchcard) CalculateNumberOfWins() int {
	numberOfWins := 0
	parsedWins := make(map[int]bool)
	for _, win := range s.wins {
		parsedWins[win] = true
	}
	for _, goten := range s.goten {
		if _, ok := parsedWins[goten]; ok {
			numberOfWins++
		}
	}
	return numberOfWins
}

func (s *SolverFour) Parse(file string) {
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

func (s *SolverFour) SolveFirstProblem() int {
	scratchcards := parseScratchcards(s.content)
	valuesSum := 0
	for _, scratchcard := range scratchcards {
		cardValue := scratchcard.CalculateValue()
		valuesSum += cardValue
	}
	return valuesSum
}

func (s *SolverFour) SolveSecondProblem() int {
	return 0
}

func parseScratchcards(content []string) []Scratchcard {
	scratchcards := make([]Scratchcard, 0)
	for _, row := range content {
		splittedRow := strings.Split(row, ":")
		numbers := strings.Split(splittedRow[1], "|")

		splittedRowLeft := strings.Split(splittedRow[0], " ")
		rawId := splittedRowLeft[len(splittedRowLeft)-1]
		id, err := strconv.Atoi(rawId)
		if err != nil {
			fmt.Println("Error parsing number: ", id)
			log.Fatal(err)
		}

		scratchcards = append(scratchcards, Scratchcard{
			id:    id,
			wins:  parseNumbers(strings.Split(numbers[0], " ")),
			goten: parseNumbers(strings.Split(numbers[1], " ")),
		})
	}
	return scratchcards
}

func parseNumbers(rawNumbers []string) []int {
	numbers := make([]int, 0)
	for _, rawNumber := range rawNumbers {
		number := strings.TrimSpace(rawNumber)
		if number == "" {
			continue
		}
		parsedNumber, err := strconv.Atoi(number)
		if err != nil {
			fmt.Println("Error parsing number: ", number)
		}
		numbers = append(numbers, parsedNumber)
	}
	return numbers
}
