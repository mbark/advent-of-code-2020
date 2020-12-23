package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/mbark/advent-of-code-2020/util"
)

const input = `
643719258
`

const testInput = `
389125467
`

type CircularList struct {
	Start int
	Nodes []int
	Max   int
	Min   int
}

func main() {
	stop := util.WithProfiling()
	defer stop()

	var ns []int
	for _, i := range strings.TrimSpace(testInput) {
		n, _ := strconv.Atoi(string(i))
		ns = append(ns, n)
	}

	list := Build(ns)
	fmt.Printf("first %s\n", first(list, 10))
	fmt.Printf("second %d\n", second(ns))
}

func first(list CircularList, iterations int) string {
	list = play(list, iterations)

	values := list.Values(list.Next(1))
	var result strings.Builder
	for _, v := range values {
		// ignore the 1
		if v == 1 {
			continue
		}
		result.WriteString(strconv.Itoa(v))
	}

	return result.String()
}

func second(ns []int) int {
	onem := 1000000
	for n := 1; n <= onem; n++ {
		if n <= len(ns) {
			continue
		}

		ns = append(ns, n)
	}

	list := Build(ns)

	list = play(list, 10*onem)
	next1 := list.Next(1)
	next2 := list.Next(next1)
	return next1 * next2
}

func play(list CircularList, iterations int) CircularList {
	curr := list.Start
	for i := 0; i < iterations; i++ {
		next1 := list.Next(curr)
		next2 := list.Next(next1)
		next3 := list.Next(next2)

		list.Nodes[curr] = list.Next(next3)

		dest := curr
		for {
			dest -= 1
			if dest < list.Min {
				dest = list.Max
			}

			if dest != next1 && dest != next2 && dest != next3 {
				break
			}
		}

		d := list.Next(dest)
		list.Nodes[dest] = next1
		list.Nodes[next3] = d
		curr = list.Next(curr)
	}

	return list
}

func Build(list []int) CircularList {
	var start int
	var current int
	var max, min int
	nodes := make([]int, len(list)+1)

	for _, n := range list {
		if start == 0 {
			start = n
			current = start
			continue
		}

		nodes[current] = n
		current = n
		if min == 0 || n < min {
			min = n
		}
		if max == 0 || n > max {
			max = n
		}
	}

	nodes[current] = start
	return CircularList{Start: start, Nodes: nodes, Max: max, Min: min}
}

func (c CircularList) Values(val int) []int {
	start := val
	curr := start

	var values []int
	for {
		values = append(values, curr)
		curr = c.Next(curr)
		if curr == start {
			break
		}
	}

	return values
}

func (c CircularList) Next(val int) int {
	return c.Nodes[val]
}
