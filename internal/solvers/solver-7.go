package solvers

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type SolverSeven struct {
	content []string
}

const (
	FiveOfAKind = iota + 1
	FourOfAKind
	FullHouse
	ThreeOfAKind
	TwoPair
	OnePair
	HighCard
)

type Hand struct {
	Cards      []Card
	bid        int
	typeOfHand int
}

func (h *Hand) New(cards []Card, bid int) Hand {
	return Hand{
		Cards:      cards,
		bid:        bid,
		typeOfHand: h.determineHandType(cards),
	}
}

func (h *Hand) determineHandType(cards []Card) int {
	typeOfHand := HighCard

	mapOfCards := make(map[string]int)
	pairWasFound := false
	threeOfAKindWasFound := false

	for _, card := range cards {
		mapOfCards[card.Value] += 1

		if mapOfCards[card.Value] == 5 && typeOfHand < FiveOfAKind {
			typeOfHand = FiveOfAKind
		} else if mapOfCards[card.Value] == 4 && typeOfHand < FourOfAKind {
			typeOfHand = FourOfAKind
		} else if mapOfCards[card.Value] == 3 && pairWasFound && typeOfHand < FullHouse {
			threeOfAKindWasFound = true
			typeOfHand = FullHouse
		} else if mapOfCards[card.Value] == 2 && threeOfAKindWasFound && typeOfHand < FullHouse {
			typeOfHand = FullHouse
		} else if mapOfCards[card.Value] == 3 && typeOfHand < ThreeOfAKind {
			threeOfAKindWasFound = true
			typeOfHand = ThreeOfAKind
		} else if mapOfCards[card.Value] == 2 && pairWasFound && typeOfHand < TwoPair {
			typeOfHand = TwoPair
		} else if mapOfCards[card.Value] == 2 && typeOfHand < OnePair {
			pairWasFound = true
			typeOfHand = OnePair
		}
	}

	return typeOfHand
}

// A, K, Q, J, T, 9, 8, 7, 6, 5, 4, 3, or 2.
type Card struct {
	Value string
}

func (c *Card) Strength() int {
	switch c.Value {
	case "A":
		return 14
	case "K":
		return 13
	case "Q":
		return 12
	case "J":
		return 11
	case "T":
		return 10
	case "9", "8", "7", "6", "5", "4", "3", "2":
		strengthAsInt, err := strconv.Atoi(c.Value)
		if err != nil {
			fmt.Println("Error parsing card value")
			log.Fatal(err)
		}
		return strengthAsInt
	default:
		log.Fatal("Invalid card value")
		return 0
	}
}

func (s *SolverSeven) SolveFirstProblem() int {
	return 0
}

func (s *SolverSeven) SolveSecondProblem() int {
	return 0
}

func (s *SolverSeven) Parse(file string) {
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
