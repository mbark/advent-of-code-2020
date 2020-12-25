package main

import (
	"fmt"
	"strconv"

	"github.com/mbark/advent-of-code-2020/util"
)

const input = `
12320657
9659666
`

const testInput = `
5764801
17807724
`

func main() {
	l := util.ReadInput(input, "\n")
	pbk1, _ := strconv.Atoi(l[0])
	pbk2, _ := strconv.Atoi(l[1])

	fmt.Printf("first %d\n", first(pbk1, pbk2))
}

const (
	subjectNumber = 7
	divideBy      = 20201227
)

func first(pbk1, pbk2 int) int {
	loop1 := find(pbk1)
	loop2 := find(pbk2)

	enckey1 := 1
	for i := 0; i < loop1; i++ {
		enckey1 = apply(enckey1, pbk2)
	}
	enckey2 := 1
	for i := 0; i < loop2; i++ {
		enckey2 = apply(enckey2, pbk1)
	}

	return enckey1
}

func find(target int) int {
	loop := 0
	key := 1
	for {
		loop += 1
		key = apply(key, subjectNumber)
		if key == target {
			return loop
		}
	}
}

func apply(key, subjectNumber int) int {
	return (key * subjectNumber) % divideBy
}
