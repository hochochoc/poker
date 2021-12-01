package main

import (
	"flag"
	"fmt"
	"log"
	"sort"
	"strings"
)

const (
	letters = "23456789TJQKA"
	tie     = "Tie"
	hand1   = "Hand 1"
	hand2   = "Hand 2"
)

var cardLetters = map[string]int{
	"2": 2,
	"3": 3,
	"4": 4,
	"5": 5,
	"6": 6,
	"7": 7,
	"8": 8,
	"9": 9,
	"t": 10,
	"j": 11,
	"q": 12,
	"k": 13,
	"a": 14,
}

func main() {
	flag.Parse()
	hands := flag.Args()
	if len(hands) != 2 {
		log.Fatalf("wrong number of args, expected 2, got %d", len(hands))
	}
	winner, err := evaluate(hands[0], hands[1])
	if err != nil {
		log.Fatal(err)
	}
	println(winner)
}

type cardFrequency struct {
	Value int
	Freq  int
}

func evaluate(hand1, hand2 string) (string, error) {
	rankedHand1, err := rankHand(hand1)
	if err != nil {
		return "", err
	}
	rankedHand2, err := rankHand(hand2)
	if err != nil {
		return "", err
	}
	return compareRankedHands(rankedHand1, rankedHand2), nil
}

func rankHand(hand string) ([]*cardFrequency, error) {
	normHand := strings.ToLower(hand)
	rankedHand := make([]*cardFrequency, 0)
	cardMap := cardCount(strings.Split(normHand, ""))
	for card, freq := range cardMap {
		if value, ok := cardLetters[card]; !ok {
			return nil, fmt.Errorf("input hand's letters are not in %s", letters)
		} else {
			rankedHand = append(rankedHand, &cardFrequency{
				Value: value,
				Freq:  freq,
			})
		}
	}
	sort.Slice(rankedHand, func(i, j int) bool {
		if rankedHand[i].Freq == rankedHand[j].Freq {
			return rankedHand[i].Value > rankedHand[j].Value
		}
		return rankedHand[i].Freq > rankedHand[j].Freq
	})
	return rankedHand, nil
}

func compareRankedHands(rankedHand1, rankedHand2 []*cardFrequency) string {
	if len(rankedHand1) < len(rankedHand2) {
		return hand1
	} else if len(rankedHand1) > len(rankedHand2) {
		return hand2
	}
	for i, c1 := range rankedHand1 {
		c2 := rankedHand2[i]
		if c1.Freq > c2.Freq {
			return hand1
		} else if c1.Freq < c2.Freq {
			return hand2
		}
		if c1.Value > c2.Value {
			return hand1
		} else if c1.Value < c2.Value {
			return hand2
		}
	}
	return tie
}

func cardCount(cards []string) map[string]int {
	cardMap := make(map[string]int)
	for _, c := range cards {
		if _, ok := cardMap[c]; ok {
			cardMap[c] += 1
		} else {
			cardMap[c] = 1
		}
	}
	return cardMap
}

