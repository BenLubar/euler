// Poker hands
// Problem 54
//
// In the card game poker, a hand consists of five cards and are ranked, from
// lowest to highest, in the following way:
//
// High Card: Highest value card.
// One Pair: Two cards of the same value.
// Two Pairs: Two different pairs.
// Three of a Kind: Three cards of the same value.
// Straight: All cards are consecutive values.
// Flush: All cards of the same suit.
// Full House: Three of a kind and a pair.
// Four of a Kind: Four cards of the same value.
// Straight Flush: All cards are consecutive values of same suit.
// Royal Flush: Ten, Jack, Queen, King, Ace, in same suit.
// The cards are valued in the order:
// 2, 3, 4, 5, 6, 7, 8, 9, 10, Jack, Queen, King, Ace.
//
// If two players have the same ranked hands then the rank made up of the
// highest value wins; for example, a pair of eights beats a pair of fives
// (see example 1 below). But if two ranks tie, for example, both players have
// a pair of queens, then highest cards in each hand are compared (see example
// 4 below); if the highest cards tie then the next highest cards are compared,
// and so on.
//
// Consider the following five hands dealt to two players:
//
// Hand	 	Player 1	 	Player 2	 	Winner
// 1	 	5H 5C 6S 7S KD		2C 3S 8S 8D TD		Player 2
//		Pair of Fives		Pair of Eights
// 2	 	5D 8C 9S JS AC		2C 5C 7D 8S QH		Player 1
//		Highest card Ace	Highest card Queen
// 3	 	2D 9C AS AH AC		3D 6D 7D TD QD		Player 2
//		Three Aces		Flush with Diamonds
// 4	 	4D 6S 9H QH QC		3D 6D 7H QD QS		Player 1
//		Pair of Queens		Pair of Queens
//		Highest card Nine	Highest card Seven
// 5	 	2H 2D 4C 4D 4S		3C 3D 3S 9S 9D		Player 1
//		Full House		Full House
//		With Three Fours	with Three Threes
//
// The file, poker.txt, contains one-thousand random hands dealt to two
// players. Each line of the file contains ten cards (separated by a single
// space): the first five are Player 1's cards and the last five are Player 2's
// cards. You can assume that all hands are valid (no invalid characters or
// repeated cards), each player's hand is in no specific order, and in each
// hand there is a clear winner.
//
// How many hands does Player 1 win?
package main

import (
	"bufio"
	"fmt"
	"net/http"
	"sort"
	"strings"
)

var types = []func(hand) []int{
	royalFlush,
	straightFlush,
	fourOfAKind,
	fullHouse,
	flush,
	straight,
	threeOfAKind,
	twoPairs,
	onePair,
	highCard,
}

func main() {
	ch := make(chan round)
	go gen(ch)

	var count int

	for r := range ch {
		for _, t := range types {
			if s := compare(t(r.p1), t(r.p2)); s != 0 {
				if s == 1 {
					count++
				}
				break
			}
		}
	}

	fmt.Println(count)
}

type round struct {
	p1, p2 hand
}

type hand [5]card

type card struct {
	rank
	suit
}

var ranks = map[byte]rank{
	'2': two,
	'3': three,
	'4': four,
	'5': five,
	'6': six,
	'7': seven,
	'8': eight,
	'9': nine,
	'T': ten,
	'J': jack,
	'Q': queen,
	'K': king,
	'A': ace,
}

type rank int

const (
	two rank = iota
	three
	four
	five
	six
	seven
	eight
	nine
	ten
	jack
	queen
	king
	ace
	rankCount
)

var suits = map[byte]suit{
	'H': hearts,
	'D': diamonds,
	'S': spades,
	'C': clubs,
}

type suit int

const (
	hearts suit = iota
	diamonds
	spades
	clubs
)

func gen(ch chan<- round) {
	defer close(ch)

	resp, err := http.Get("https://projecteuler.net/project/resources/p054_poker.txt")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		panic(resp.Status)
	}

	s := bufio.NewScanner(resp.Body)
	for s.Scan() {
		ch <- parse(s.Text())
	}
	if err := s.Err(); err != nil {
		panic(err)
	}
}

func parse(s string) round {
	f := strings.Fields(s)

	return round{
		p1: parseHand(f[:5]),
		p2: parseHand(f[5:]),
	}
}

func parseHand(f []string) hand {
	return hand{
		parseCard(f[0]),
		parseCard(f[1]),
		parseCard(f[2]),
		parseCard(f[3]),
		parseCard(f[4]),
	}
}

func parseCard(s string) card {
	return card{
		rank: ranks[s[0]],
		suit: suits[s[1]],
	}
}

func compare(s1, s2 []int) int {
	if s1 == nil {
		if s2 == nil {
			return 0
		}

		return 2
	}
	if s2 == nil {
		return 1
	}

	for i := range s1 {
		if s1[i] < s2[i] {
			return 2
		} else if s1[i] > s2[i] {
			return 1
		}
	}

	return 0
}

func isFlush(h hand) bool {
	s := h[0].suit
	for _, c := range h {
		if c.suit != s {
			return false
		}
	}
	return true
}

func royalFlush(h hand) []int {
	if !isFlush(h) {
		return nil
	}

	var have [rankCount]bool
	for _, c := range h {
		have[c.rank] = true
	}

	if have[ten] && have[jack] && have[queen] && have[king] && have[ace] {
		return []int{0}
	}

	return nil
}

func straightFlush(h hand) []int {
	if !isFlush(h) {
		return nil
	}

	return straight(h)
}

func getRanks(h hand) [rankCount]int {
	var r [rankCount]int

	for _, c := range h {
		r[c.rank]++
	}

	return r
}

func kinds(r [rankCount]int, counts ...int) []int {
	var score []int

search:
	for _, c := range counts {
		for j := len(r) - 1; j >= 0; j-- {
			if r[j] == c {
				r[j] = 0
				score = append(score, j)
				continue search
			}
		}
		return nil
	}

	for j := len(r) - 1; j >= 0; j-- {
		for i := 0; i < r[j]; i++ {
			score = append(score, j)
		}
	}

	return score
}

func fourOfAKind(h hand) []int {
	return kinds(getRanks(h), 4)
}

func fullHouse(h hand) []int {
	return kinds(getRanks(h), 3, 2)
}

func flush(h hand) []int {
	if isFlush(h) {
		return highCard(h)
	}

	return nil
}

func straight(h hand) []int {
	r := getRanks(h)
	for i := range r[:len(r)-4] {
		if r[i] == 1 && r[i+1] == 1 && r[i+2] == 1 && r[i+3] == 1 && r[i+4] == 1 {
			return []int{i}
		}
	}
	return nil
}

func threeOfAKind(h hand) []int {
	return kinds(getRanks(h), 3)
}

func twoPairs(h hand) []int {
	return kinds(getRanks(h), 2, 2)
}

func onePair(h hand) []int {
	return kinds(getRanks(h), 2)
}

func highCard(h hand) []int {
	var score [5]int
	for i, c := range h {
		score[i] = int(c.rank)
	}
	sort.Ints(score[:])
	score[0], score[4] = score[4], score[0]
	score[1], score[3] = score[3], score[1]
	return score[:]
}
