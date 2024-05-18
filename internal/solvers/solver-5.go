package solvers

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type FromToRange struct {
	from int
	size int
	to   int
}

type Node struct {
	source      string
	destination string
	ranges      []FromToRange
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

	locations := make([]int, 0)
	for _, seed := range s.seeds {
		location := getLocationBySeed(seed, s.nodes)
		locations = append(locations, location)
	}

	var closestLocation int
	for _, location := range locations {
		if closestLocation == 0 || location < closestLocation {
			closestLocation = location
		}
	}

	return closestLocation
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
		for _, fromToRange := range node.ranges {
			localMappedValue, found := fromToRange.Get(valueToMapFrom)

			mappedValue = localMappedValue

			if found {
				break
			}
		}

		valueToMapFrom = mappedValue
	}

	return mappedValue
}

func (f FromToRange) Get(valueToMapFrom int) (int, bool) {
	valueInRange := f.from <= valueToMapFrom && valueToMapFrom <= f.from+f.size
	if !valueInRange {
		return valueToMapFrom, false
	}

	offset := valueToMapFrom - f.from

	return f.to + offset, true
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
				ranges:      make([]FromToRange, 0),
			})
		} else if row != "" {
			from, to, size := buildMapRow(row)
			lastNodeIdx := len(nodes) - 1
			nodes[lastNodeIdx].ranges = append(nodes[lastNodeIdx].ranges, FromToRange{
				from: from,
				to:   to,
				size: size,
			})
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

func buildMapRow(row string) (int, int, int) {
	splittedRow := strings.Split(row, " ")

	to, err := strconv.Atoi(splittedRow[0])
	if err != nil {
		fmt.Println("Error converting destination number to int")
		log.Fatal(err)
	}

	from, err := strconv.Atoi(splittedRow[1])
	if err != nil {
		fmt.Println("Error converting source number to int")
		log.Fatal(err)
	}

	size, err := strconv.Atoi(splittedRow[2])
	if err != nil {
		fmt.Println("Error converting map range to int")
		log.Fatal(err)
	}

	return from, to, size
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
