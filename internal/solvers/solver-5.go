package solvers

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Node struct {
	source      string
	destination string
	maps        map[int]int
}

type SolverFive struct {
	content []string
	seeds   []int
	nodes   []Node
}

func (s *SolverFive) Parse(file string) {
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

func (s *SolverFive) SolveFirstProblem() int {
	s.buildNodes()

	var location int
	for _, seed := range s.seeds {
		location = getLocationBySeed(seed, s.nodes)
	}

	return location
}

func (s *SolverFive) SolveSecondProblem() int {
	return 0
}

func getLocationBySeed(seed int, nodes []Node) int {
	sourcesDestinations := [][]string{
		{"seed", "soil"},
		{"soil", "fertilizer"},
		{"fertilizer", "water"},
		{"water", "light"},
		{"light", "temperature"},
		{"temperature", "humidity"},
		{"humidity", "location"},
	}

	valueToMapFrom := seed
	var mappedValue int
	for _, sourceDestination := range sourcesDestinations {
		source := sourceDestination[0]
		destination := sourceDestination[1]

		node := getBySourceAndDestination(source, destination, nodes)
		mappedValue = node.maps[valueToMapFrom]
		if mappedValue == 0 {
			mappedValue = valueToMapFrom
		}

		valueToMapFrom = mappedValue
	}

	return mappedValue
}

func (s *SolverFive) buildNodes() {
	s.seeds = buildSeedsRow(s.content[0])
	nodes := make([]Node, 0)

	for _, row := range s.content[1:] {
		if isHeaderRow(row) {
			source, destination := getSourceAndDestination(row)
			nodes = append(nodes, Node{
				source:      source,
				destination: destination,
				maps:        make(map[int]int),
			})
		} else if row != "" {
			sourceDestinationMap := buildMapRow(row)
			mergedMaps := mergeSourceDestinationMaps(nodes[len(nodes)-1].maps, sourceDestinationMap)
			nodes[len(nodes)-1].maps = mergedMaps
		}
	}

	s.nodes = nodes
}

func buildSeedsRow(row string) []int {
	seeds := make([]int, 0)

	row = strings.TrimPrefix(row, "seeds: ")
	row = strings.TrimSpace(row)

	for _, seed := range strings.Split(row, " ") {
		seed = strings.TrimSpace(seed)
		intSeed, err := strconv.Atoi(seed)
		if err != nil {
			fmt.Println("Error converting seed to int")
			log.Fatal(err)
		}
		seeds = append(seeds, intSeed)
	}

	return seeds
}

func getSourceAndDestination(row string) (string, string) {
	row = strings.TrimSuffix(row, "map:")
	row = strings.TrimSpace(row)

	sourcesAndDestinations := strings.Split(row, "-to-")
	source := sourcesAndDestinations[0]
	destination := sourcesAndDestinations[1]

	return source, destination
}

func isHeaderRow(row string) bool {
	return strings.Contains(row, "map:")
}

func buildMapRow(row string) map[int]int {
	mappedValues := make(map[int]int)

	splittedRow := strings.Split(row, " ")
	destinatioNumber, err := strconv.Atoi(splittedRow[0])
	if err != nil {
		fmt.Println("Error converting destination number to int")
		log.Fatal(err)
	}
	sourceNumber, err := strconv.Atoi(splittedRow[1])
	if err != nil {
		fmt.Println("Error converting source number to int")
		log.Fatal(err)
	}
	mapRange, err := strconv.Atoi(splittedRow[2])
	if err != nil {
		fmt.Println("Error converting map range to int")
		log.Fatal(err)
	}

	for i := 0; i < mapRange; i++ {
		mappedValues[sourceNumber] = destinatioNumber
		sourceNumber++
		destinatioNumber++
	}

	return mappedValues
}

func mergeSourceDestinationMaps(map1 map[int]int, map2 map[int]int) map[int]int {
	for key, value := range map2 {
		map1[key] = value
	}

	return map1
}

func getBySourceAndDestination(source string, destination string, nodes []Node) Node {
	var ans Node
	for _, node := range nodes {
		if node.source == source && node.destination == destination {
			ans = node
		}
	}
	return ans
}
