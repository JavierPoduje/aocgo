package solvers

import (
	"testing"
)

func TestSolverSeven_SolveFirstProblem(t *testing.T) {
	solver := &SolverSeven{}
	solver.content = []string{
		"32T3K 765",
		"T55J5 684",
		"KK677 28",
		"KTJJT 220",
		"QQQJA 483",
	}
	ans := solver.SolveFirstProblem()
	if ans != 6440 {
		t.Errorf("Expected 6440, got %d", ans)
	}
}

func TestSolverSeven_rawCardsToCards(t *testing.T) {
	rawCards := "32T3K"
	cards := stringToSliceOfCards(rawCards)
	if len(cards) != 5 {
		t.Errorf("Expected 5, got %d", len(cards))
	}
	if cards[0].Value != "3" {
		t.Errorf("Expected 3, got %s", cards[0].Value)
	}
	if cards[1].Value != "2" {
		t.Errorf("Expected 2, got %s", cards[1].Value)
	}
	if cards[2].Value != "T" {
		t.Errorf("Expected T, got %s", cards[2].Value)
	}
	if cards[3].Value != "3" {
		t.Errorf("Expected 3, got %s", cards[3].Value)
	}
	if cards[4].Value != "K" {
		t.Errorf("Expected K, got %s", cards[4].Value)
	}
}

func TestSolverSeven_DetermineHandType(t *testing.T) {
	cards := stringToSliceOfCards("77777")
	handType := DetermineHandType(cards)
	stringifiedType := typeToString(handType)
	if stringifiedType != "FiveOfAKind" {
		t.Errorf("Expected FiveOfAKind, got %s", stringifiedType)
	}

	cards = stringToSliceOfCards("7777K")
	handType = DetermineHandType(cards)
	stringifiedType = typeToString(handType)
	if stringifiedType != "FourOfAKind" {
		t.Errorf("Expected FourOfAKind, got %s", stringifiedType)
	}

	cards = stringToSliceOfCards("KKK77")
	handType = DetermineHandType(cards)
	stringifiedType = typeToString(handType)
	if stringifiedType != "FullHouse" {
		t.Errorf("Expected FullHouse, got %s", stringifiedType)
	}

	cards = stringToSliceOfCards("KK777")
	handType = DetermineHandType(cards)
	stringifiedType = typeToString(handType)
	if stringifiedType != "FullHouse" {
		t.Errorf("Expected FullHouse, got %s", stringifiedType)
	}

	cards = stringToSliceOfCards("T55J5")
	handType = DetermineHandType(cards)
	stringifiedType = typeToString(handType)
	if stringifiedType != "ThreeOfAKind" {
		t.Errorf("Expected ThreeOfAKind, got %s", stringifiedType)
	}

	cards = stringToSliceOfCards("KK677")
	handType = DetermineHandType(cards)
	stringifiedType = typeToString(handType)
	if stringifiedType != "TwoPair" {
		t.Errorf("Expected TwoPair, got %s", stringifiedType)
	}

	cards = stringToSliceOfCards("32T3K")
	handType = DetermineHandType(cards)
	stringifiedType = typeToString(handType)
	if stringifiedType != "OnePair" {
		t.Errorf("Expected OnePair, got %s", stringifiedType)
	}

	cards = stringToSliceOfCards("A2345")
	handType = DetermineHandType(cards)
	stringifiedType = typeToString(handType)
	if stringifiedType != "HighCard" {
		t.Errorf("Expected HighCard, got %s", stringifiedType)
	}
}
