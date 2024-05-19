package solvers

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

type SolverSeven struct {
	content []string
}

const (
	HighCard = iota + 1
	OnePair
	TwoPair
	ThreeOfAKind
	FullHouse
	FourOfAKind
	FiveOfAKind
)

func typeToString(t int) string {
	switch t {
	case HighCard:
		return "HighCard"
	case OnePair:
		return "OnePair"
	case TwoPair:
		return "TwoPair"
	case ThreeOfAKind:
		return "ThreeOfAKind"
	case FullHouse:
		return "FullHouse"
	case FourOfAKind:
		return "FourOfAKind"
	case FiveOfAKind:
		return "FiveOfAKind"
	default:
		log.Fatal("Invalid type of hand")
		return ""
	}
}

type Hand struct {
	Cards      []Card
	bid        int
	typeOfHand int
}

type Hands []*Hand

// implement sort.Interface
func (h Hands) Len() int {
	return len(h)
}

func (h Hand) String() string {
	return fmt.Sprintf(
		"\nHand{Cards: %v, bid: %d, typeOfHand: %s}",
		h.Cards,
		h.bid,
		typeToString(h.typeOfHand),
	)
}

func (h Hands) Less(i, j int) bool {
	if h[i].Eq(h[j]) {
		return false
	}
	return h[i].Less(h[j])
}

func (h Hands) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *Hand) Less(other *Hand) bool {
	if h.typeOfHand == other.typeOfHand {
		return h.LessByStrength(other)
	}
	return h.typeOfHand < other.typeOfHand
}

func (h *Hand) LessByStrength(other *Hand) bool {
	for idx, hCard := range h.Cards {
		otherCard := other.Cards[idx]
		if hCard.Strength() < otherCard.Strength() {
			return true
		} else if hCard.Strength() > otherCard.Strength() {
			return false
		}
	}
	return false
}

func (h *Hand) Eq(other *Hand) bool {
	if h.typeOfHand == other.typeOfHand {
		return h.EqByStrength(other)
	}
	return false
}

func (h *Hand) EqByStrength(other *Hand) bool {
	for idx, hCard := range h.Cards {
		otherCard := other.Cards[idx]
		if hCard.Strength() != otherCard.Strength() {
			return false
		}
	}
	return true
}

func NewHand(cards []Card, bid int) Hand {
	return Hand{
		Cards:      cards,
		bid:        bid,
		typeOfHand: DetermineHandType(cards),
	}
}

func DetermineHandType(cards []Card) int {
	typeOfHand := HighCard

	mapOfCards := make(map[string]int)

	check := func(cards []Card, threshold int, exclusions map[string]bool) bool {
		localMap := make(map[string]int)
		for _, card := range cards {
			if _, ok := exclusions[card.Value]; ok {
				continue
			}
			localMap[card.Value] += 1
			if localMap[card.Value] >= threshold {
				return true
			}
		}
		return false
	}

	for _, card := range cards {
		mapOfCards[card.Value] += 1
		currCardValue := card.Value

		if mapOfCards[currCardValue] == 5 && typeOfHand < FiveOfAKind {
			typeOfHand = FiveOfAKind
		} else if mapOfCards[currCardValue] == 4 && typeOfHand < FourOfAKind {
			typeOfHand = FourOfAKind
		} else if mapOfCards[currCardValue] == 3 && typeOfHand < FullHouse && check(cards, 2, map[string]bool{currCardValue: true}) {
			typeOfHand = FullHouse
		} else if mapOfCards[currCardValue] == 2 && typeOfHand < FullHouse && check(cards, 3, map[string]bool{currCardValue: true}) {
			typeOfHand = FullHouse
		} else if mapOfCards[currCardValue] == 3 && typeOfHand < ThreeOfAKind {
			typeOfHand = ThreeOfAKind
		} else if mapOfCards[currCardValue] == 2 && typeOfHand < TwoPair && check(cards, 2, map[string]bool{currCardValue: true}) {
			typeOfHand = TwoPair
		} else if mapOfCards[currCardValue] == 2 && typeOfHand < OnePair {
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
	hands := buildHands(s.content)
	sort.Sort(hands)

	sumOfBiddMultipliedByIdx := 0
	for idx, hand := range hands {
		sumOfBiddMultipliedByIdx += hand.bid * (idx + 1)
	}

	return sumOfBiddMultipliedByIdx
}

func (s *SolverSeven) SolveSecondProblem() int {
	return 0
}

func buildHands(content []string) Hands {
	hands := make(Hands, 0)
	for _, row := range content {
		splittedRow := strings.Split(row, " ")
		rawCards := splittedRow[0]
		rawBid := splittedRow[1]

		cards := stringToSliceOfCards(rawCards)
		bid, err := strconv.Atoi(rawBid)
		if err != nil {
			fmt.Println("Error parsing bid")
			log.Fatal(err)
		}

		hand := NewHand(cards, bid)
		hands = append(hands, &hand)
	}
	return hands
}

func stringToSliceOfCards(rawCards string) []Card {
	cards := make([]Card, 0)
	for _, rawCard := range rawCards {
		cards = append(cards, Card{Value: string(rawCard)})
	}
	return cards
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
