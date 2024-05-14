package solvers

import (
	"log"
	"os"
	"strconv"
	"strings"
)

type SolverTwo struct {
	content []string
}

type Game struct {
	blue  int
	green int
	red   int
}

type Record struct {
	id    int
	games []Game
}

func (s *SolverTwo) Parse(file string) {
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

func (s *SolverTwo) SolveFirstProblem() int {
	records := make([]Record, len(s.content))
	for _, row := range s.content {
		id, err := strconv.Atoi(gameId(row))
		if err != nil {
			log.Fatal(err)
		}
		games := gamesFromRow(row)
		records = append(records, Record{id: id, games: games})
	}

	filteredRecords := make([]Record, 0)
	for _, record := range records {
		if isValidRecord(record) {
			filteredRecords = append(filteredRecords, record)
		}
	}

	recordIdsSum := 0
	for _, record := range filteredRecords {
		recordIdsSum += record.id
	}

	return recordIdsSum
}

func (s *SolverTwo) SolveSecondProblem() int {
	return 0
}

func isValidRecord(record Record) bool {
	for _, game := range record.games {
		if !isValidGame(game) {
			return false
		}
	}
	return true
}

func isValidGame(game Game) bool {
	const (
		MaxAmountOfBlueCubes  = 14
		MaxAmountOfGreenCubes = 13
		MaxAmountOfRedCubes   = 12
	)
	return game.blue <= MaxAmountOfBlueCubes &&
		game.green <= MaxAmountOfGreenCubes &&
		game.red <= MaxAmountOfRedCubes
}

func gameId(row string) string {
	left := strings.Split(row, ":")[0]
	number := strings.Split(left, " ")[1]
	return number
}

func gamesFromRow(row string) []Game {
	right := strings.Split(row, ":")[1]
	rawGames := strings.Split(right, ";")

	games := make([]Game, 0)
	for _, rawGame := range rawGames {
		game := parseGame(rawGame)
		games = append(games, game)
	}

	return games
}

func parseGame(rawGame string) Game {
	red := 0
	blue := 0
	green := 0

	trimmedRawGame := strings.TrimSpace(rawGame)
	colors := strings.Split(trimmedRawGame, ",")

	for _, color := range colors {
		trimmed := strings.TrimSpace(color)
		numberAndColor := strings.Split(trimmed, " ")

		number, err := strconv.Atoi(numberAndColor[0])
		if err != nil {
			log.Fatal(err)
		}
		color := numberAndColor[1]

		if color == "red" {
			red = number
		} else if color == "blue" {
			blue = number
		} else {
			green = number
		}
	}

	return Game{
		blue:  blue,
		green: green,
		red:   red,
	}
}
