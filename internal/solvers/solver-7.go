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

func DetermineHandType(cards []Card, withJokers bool) int {
	if withJokers {
		return DetermineHandTypeWithJokers(cards)
	}
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

func DetermineHandTypeWithJokers(cards []Card) int {
	typeOfHand := HighCard

	mapOfCards := make(map[string]int)

	check := func(cards []Card, threshold int, exclusions map[string]bool) bool {
		localMap := make(map[string]int)
		for _, card := range cards {
			if _, ok := exclusions[card.Value]; ok || card.Value == "J" {
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

		if card.Value == "J" {
			continue
		}

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

	if numberOfJokers, ok := mapOfCards["J"]; ok {
		typeOfHand = typeOfHandUpgradeWithJokers(typeOfHand, numberOfJokers)
	}

	return typeOfHand
}

func typeOfHandUpgradeWithJokers(typeOfHand int, numberOfJokers int) int {
	switch typeOfHand {
	case HighCard:
		if numberOfJokers == 1 {
			return OnePair
		}
		if numberOfJokers == 2 {
			return ThreeOfAKind
		}
		if numberOfJokers == 3 {
			return FourOfAKind
		}
		if numberOfJokers >= 4 {
			return FiveOfAKind
		}
	case OnePair:
		if numberOfJokers == 1 {
			return ThreeOfAKind
		}
		if numberOfJokers == 2 {
			return FourOfAKind
		}
		if numberOfJokers >= 3 {
			return FiveOfAKind
		}
	case TwoPair:
		if numberOfJokers == 1 {
			return FullHouse
		}
		if numberOfJokers == 2 {
			return FourOfAKind
		}
		if numberOfJokers >= 3 {
			return FiveOfAKind
		}
	case ThreeOfAKind:
		if numberOfJokers == 1 {
			return FourOfAKind
		}
		if numberOfJokers >= 2 {
			return FiveOfAKind
		}
	case FourOfAKind:
		if numberOfJokers == 1 {
			return FiveOfAKind
		}
	}
	return typeOfHand
}

func (s *SolverSeven) SolveFirstProblem() int {
	hands := buildHands(s.content, false)
	sort.Sort(hands)

	sumOfBiddMultipliedByIdx := 0
	for idx, hand := range hands {
		sumOfBiddMultipliedByIdx += hand.bid * (idx + 1)
	}

	return sumOfBiddMultipliedByIdx
}

func (s *SolverSeven) SolveSecondProblem() int {
	hands := buildHands(s.content, true)
	sort.Sort(hands)

	sumOfBiddMultipliedByIdx := 0
	for idx, hand := range hands {
		sumOfBiddMultipliedByIdx += hand.bid * (idx + 1)
	}

	return sumOfBiddMultipliedByIdx
}

func buildHands(content []string, withJokers bool) Hands {
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

		hand := NewHand(cards, bid, withJokers)
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

// *** *** ***
// ** CARD **
// *** *** ***

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

func (c *Card) StrengthWithJokers() int {
	switch c.Value {
	case "A":
		return 14
	case "K":
		return 13
	case "Q":
		return 12
	case "T":
		return 10
	case "9", "8", "7", "6", "5", "4", "3", "2":
		strengthAsInt, err := strconv.Atoi(c.Value)
		if err != nil {
			fmt.Println("Error parsing card value")
			log.Fatal(err)
		}
		return strengthAsInt
	case "J":
		return 1
	default:
		log.Fatal("Invalid card value")
		return 0
	}
}

// *** *** ***
// ** HAND **
// *** *** ***

type Hand struct {
	Cards      []Card
	bid        int
	typeOfHand int
	withJokers bool
}

func NewHand(cards []Card, bid int, withJokers bool) Hand {
	return Hand{
		Cards:      cards,
		bid:        bid,
		typeOfHand: DetermineHandType(cards, withJokers),
		withJokers: withJokers,
	}
}

func (h *Hand) Less(other *Hand) bool {
	if h.typeOfHand == other.typeOfHand {
		return h.LessByStrength(other)
	}
	return h.typeOfHand < other.typeOfHand
}

func (h *Hand) LessByStrength(other *Hand) bool {
	if h.withJokers && other.withJokers {
		return h.LessByStrengthWithJokers(other)
	}

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
	if h.withJokers && other.withJokers {
		return h.EqByStrengthWithJokers(other)
	}

	for idx, hCard := range h.Cards {
		otherCard := other.Cards[idx]
		if hCard.Strength() != otherCard.Strength() {
			return false
		}
	}
	return true
}

func (h *Hand) EqByStrengthWithJokers(other *Hand) bool {
	for idx, hCard := range h.Cards {
		otherCard := other.Cards[idx]
		if hCard.StrengthWithJokers() != otherCard.StrengthWithJokers() {
			return false
		}
	}
	return true
}

func (h *Hand) LessByStrengthWithJokers(other *Hand) bool {
	for idx, hCard := range h.Cards {
		otherCard := other.Cards[idx]
		if hCard.StrengthWithJokers() < otherCard.StrengthWithJokers() {
			return true
		} else if hCard.StrengthWithJokers() > otherCard.StrengthWithJokers() {
			return false
		}
	}
	return false
}

// *** *** ***
// ** HANDS **
// *** *** ***

type Hands []*Hand

func (h Hands) Len() int {
	return len(h)
}

func (h Hand) String() string {
	return fmt.Sprintf(
		"\nHand{Cards: %v, bid: %d, typeOfHand: %s, withJokers: %v}",
		h.Cards,
		h.bid,
		typeToString(h.typeOfHand),
		h.withJokers,
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
