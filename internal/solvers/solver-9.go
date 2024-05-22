package solvers

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type SolverNine struct {
	content []string
}

func (s *SolverNine) SolveFirstProblem() int {
	rows := buildHistory(s.content)

	predictions := make([]int, len(rows))
	for idx, history := range rows {
		predictions[idx] = getPredictionByHistory(history, false)
	}

	predictionsSum := 0
	for _, prediction := range predictions {
		predictionsSum += prediction
	}

	return predictionsSum
}

func (s *SolverNine) SolveSecondProblem() int {
	rows := buildHistory(s.content)

	predictions := make([]int, len(rows))
	for idx, history := range rows {
		predictions[idx] = getPredictionByHistory(history, true)
	}

	predictionsSum := 0
	for _, prediction := range predictions {
		predictionsSum += prediction
	}

	return predictionsSum

}

func getPredictionByHistory(row []int, predictFirstRow bool) int {
	rows := [][]int{row}

	for {
		newRow, allZeroes := extrapolateRow(row)
		rows = append(rows, newRow)

		if allZeroes {
			break
		}

		row = newRow
	}

	return getPredictionByExtrapolatedRow(rows, predictFirstRow)
}

func getPredictionByExtrapolatedRow(rows [][]int, predictFirstRow bool) int {
	valuesToPredictFrom := make([]int, len(rows))
	for idx, row := range rows {
		if predictFirstRow {
			valuesToPredictFrom[idx] = row[0]
		} else {
			valuesToPredictFrom[idx] = row[len(row)-1]
		}
	}

	var prediction int
	for i := len(valuesToPredictFrom) - 2; i >= 0; i-- {
		if i == len(valuesToPredictFrom)-2 {
			prediction = valuesToPredictFrom[i]
			continue
		}

		if predictFirstRow {
			prediction = valuesToPredictFrom[i] - prediction
		} else {
			prediction = valuesToPredictFrom[i] + prediction
		}
	}

	return prediction
}

func extrapolateRow(row []int) ([]int, bool) {
	if len(row) == 1 {
		return []int{0}, true
	}

	zeroes := 0
	processedRow := make([]int, len(row)-1)

	left := row[0]
	for idx, right := range row[1:] {
		diff := right - left
		if diff == 0 {
			zeroes++
		}
		processedRow[idx] = diff

		left = right
	}

	return processedRow, zeroes == len(row)-1
}

func (s *SolverNine) Parse(file string) {
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

func buildHistory(content []string) [][]int {
	rows := make([][]int, 0)

	for _, row := range content {
		if row == "" {
			continue
		}

		numbers := make([]int, 0)
		for _, rawNum := range strings.Split(row, " ") {
			trimmedNum := strings.TrimSpace(rawNum)
			num, err := strconv.Atoi(trimmedNum)
			if err != nil {
				fmt.Printf("Error converting %s to int\n", rawNum)
				log.Fatal(err)
			}
			numbers = append(numbers, num)
		}

		rows = append(rows, numbers)
	}

	return rows
}
