package main

import (
	"fmt"
	"strconv"

	"github.com/mbark/advent-of-code-2020/util"
)

var input = `1,0,15,2,10,13`

var testInput = `0,3,6`
var testInput2 = `3,1,2`

func main() {
	in := util.ReadInput(input, ",")
	var startNrs []int
	for _, i := range in {
		n, _ := strconv.Atoi(i)
		startNrs = append(startNrs, n)
	}

	fmt.Printf("first %d\n", first(startNrs, 2020))
	fmt.Printf("second %d\n", first(startNrs, 30000000))
}

func first(start []int, tillTurn int) int {
	var lastNr int
	saidAt := make([]int, tillTurn)
	for i, nr := range start {
		lastNr = nr
		if i < len(start)-1 {
			saidAt[nr] = i + 1
		}
	}

	var turn = len(start) + 1
	for ; turn <= tillTurn; turn++ {
		at := saidAt[lastNr]

		var n int
		if at < 1 {
			n = 0
		} else {
			n = turn - 1 - at
		}

		saidAt[lastNr] = turn - 1
		lastNr = n
	}

	return lastNr
}
