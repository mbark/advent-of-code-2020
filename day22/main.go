package main

import (
	"fmt"
	"strconv"

	"github.com/mbark/advent-of-code-2020/util"
)

const input = `
Player 1:
7
1
9
10
12
4
38
22
18
3
27
31
43
33
47
42
21
24
50
39
8
6
16
46
11

Player 2:
49
41
40
35
44
29
30
19
14
2
34
17
25
5
15
32
20
48
45
26
37
28
36
23
13
`

const testInput = `
Player 1:
9
2
6
3
1

Player 2:
5
8
4
7
10
`

const testInput2 = `
Player 1:
43
19

Player 2:
2
29
14
`

func main() {
	d := util.ReadInput(input, "\n\n")

	var deck1 []int
	var deck2 []int

	for _, s := range util.ReadInput(d[0], "\n")[1:] {
		i, _ := strconv.Atoi(string(s))
		deck1 = append(deck1, i)
	}

	for _, s := range util.ReadInput(d[1], "\n")[1:] {
		i, _ := strconv.Atoi(string(s))
		deck2 = append(deck2, i)
	}

	fmt.Printf("first %d\n", first(deck1, deck2))
	fmt.Printf("first %d\n", second(deck1, deck2))
}

func first(deck1, deck2 []int) int {
	for len(deck1) > 0 && len(deck2) > 0 {
		c1 := deck1[0]
		deck1 = deck1[1:]
		c2 := deck2[0]
		deck2 = deck2[1:]

		if c1 > c2 {
			deck1 = append(deck1, c1, c2)
		} else {
			deck2 = append(deck2, c2, c1)
		}
	}

	return calculateScore(deck1, deck2)
}

func second(deck1, deck2 []int) int {
	result := playRecursive(deck1, deck2, make(map[string]bool), gameCount)

	return calculateScore(result.Deck1, result.Deck2)
}

type Result struct {
	Player1Win bool
	Deck1      []int
	Deck2      []int
}

type State struct {
	Deck1 []int
	Deck2 []int
}

var gameCount int = 1

func playRecursive(deck1, deck2 []int, prev map[string]bool, game int) Result {
	for len(deck1) > 0 && len(deck2) > 0 {
		h := fmt.Sprintf("%v;%v", deck1, deck2)
		if _, ok := prev[h]; ok {
			return Result{Player1Win: true, Deck1: deck1, Deck2: deck2}
		} else {
			prev[h] = true
		}

		c1 := deck1[0]
		deck1 = deck1[1:]
		c2 := deck2[0]
		deck2 = deck2[1:]

		winner1 := true
		if c1 <= len(deck1) && c2 <= len(deck2) {
			d1 := make([]int, c1)
			copy(d1, deck1)
			d2 := make([]int, c2)
			copy(d2, deck2)

			gameCount += 1
			result := playRecursive(d1, d2, make(map[string]bool), gameCount)
			winner1 = result.Player1Win
		} else if c1 > c2 {
			winner1 = true
		} else {
			winner1 = false
		}

		if winner1 {
			deck1 = append(deck1, c1, c2)
		} else {
			deck2 = append(deck2, c2, c1)
		}
	}

	return Result{Player1Win: len(deck1) > 0, Deck1: deck1, Deck2: deck2}
}

func calculateScore(deck1, deck2 []int) int {
	if len(deck1) > 0 {
		return deckScore(deck1)
	}

	return deckScore(deck2)
}

func deckScore(deck []int) int {
	score := 0
	multiplier := 1
	for i := len(deck) - 1; i >= 0; i -= 1 {
		score += multiplier * deck[i]
		multiplier += 1
	}

	return score
}
