package solvers

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type SolverSix struct {
	content   []string
	times     []int
	distances []int
}

func (s *SolverSix) build(join bool) {
	times := make([]int, 0)
	distances := make([]int, 0)

	rowOne := strings.TrimPrefix(s.content[0], "Time:")
	rowTwo := strings.TrimPrefix(s.content[1], "Distance:")

	rowOne = strings.TrimSpace(rowOne)
	rowTwo = strings.TrimSpace(rowTwo)

	if join {
		rowOne = strings.Join(strings.Split(rowOne, " "), "")
		rowTwo = strings.Join(strings.Split(rowTwo, " "), "")
	}

	for _, time := range strings.Split(rowOne, " ") {
		if time == "" {
			continue
		}
		parsedTime, err := strconv.Atoi(time)
		if err != nil {
			fmt.Println("Error parsing time")
			log.Fatal(err)
		}
		times = append(times, parsedTime)
	}

	for _, distance := range strings.Split(rowTwo, " ") {
		if distance == "" {
			continue
		}
		parsedTime, err := strconv.Atoi(distance)
		if err != nil {
			fmt.Println("Error parsing time")
			log.Fatal(err)
		}
		distances = append(distances, parsedTime)
	}

	s.times = times
	s.distances = distances
}

func (s *SolverSix) SolveFirstProblem() int {
	s.build(false)

	solutions := make([]int, 0)
	for i := 0; i < len(s.times); i++ {
		solution := possibleRaceWins(s.times[i], s.distances[i])
		solutions = append(solutions, solution)
	}

	product := 1
	for _, s := range solutions {
		product *= s
	}

	return product
}

func (s *SolverSix) SolveSecondProblem() int {
	s.build(true)

	solutions := make([]int, 0)
	for i := 0; i < len(s.times); i++ {
		solution := possibleRaceWins(s.times[i], s.distances[i])
		solutions = append(solutions, solution)
	}

	product := 1
	for _, s := range solutions {
		product *= s
	}

	return product

}

func possibleRaceWins(time int, thresholdDistance int) int {
	wins := 0
	for i := 1; i < time; i++ {
		traveledDistance := i * (time - i)
		if traveledDistance > thresholdDistance {
			wins++
		}
	}
	return wins
}

func (s *SolverSix) Parse(file string) {
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
