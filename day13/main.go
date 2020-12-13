package main

import (
	"fmt"
	"math"
	"strconv"

	"github.com/mbark/advent-of-code-2020/util"
)

var input = `
1000507
29,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,37,x,x,x,x,x,631,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,13,19,x,x,x,23,x,x,x,x,x,x,x,383,x,x,x,x,x,x,x,x,x,41,x,x,x,x,x,x,17
`

var testInput = `
939
7,13,x,x,59,x,31,19
`

var testInput2 = `
0
1789,37,47,1889`

func main() {
	lines := util.ReadInput(input, "\n")

	ts, _ := strconv.Atoi(lines[0])
	s := lines[1]

	times := util.ReadInput(s, ",")

	wait, bus := first(ts, times)
	fmt.Printf("first %d (%d * %d)\n", wait*bus, wait, bus)
	fmt.Printf("second %s\n", second(times))
	fmt.Printf("second take two: %d\n", secondTakeTwo(times))
}

func first(timestamp int, times []string) (int, int) {
	waitTime := math.MaxInt64
	busID := 0

	for _, b := range times {
		if b == "x" {
			continue
		}

		i, _ := strconv.Atoi(b)
		wait := i - (timestamp % i)
		if wait < waitTime {
			waitTime = wait
			busID = i
		}
	}

	return waitTime, busID
}

type Remainder struct {
	Num int
	Mod int
}

func second(buses []string) string {
	var remainder []Remainder
	for i, b := range buses {
		if b == "x" {
			continue
		}

		id, _ := strconv.Atoi(b)
		remainder = append(remainder, Remainder{Num: id - i, Mod: id})
	}

	s := "ChineseRemainder[{"
	for _, r := range remainder {
		s += strconv.Itoa(r.Num) + ","
	}

	//remove last comma
	s = s[:len(s)-1]
	s += "},{"
	for _, r := range remainder {
		s += strconv.Itoa(r.Mod) + ","
	}

	s = s[:len(s)-1]
	s += "}]"

	return s
}

func secondTakeTwo(buses []string) int {
	nr := 1
	step := 1

	for i, b := range buses {
		if b == "x" {
			continue
		}

		id, _ := strconv.Atoi(b)
		for (nr+i)%id != 0 {
			nr += step
		}

		step *= id
	}

	return nr
}
