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
	Start *Node
	Nodes []*Node
	Max   int
	Min   int
}

type Node struct {
	Val  int
	Next *Node
}

func main() {
	stop := util.WithProfiling()
	defer stop()

	var ns []int
	for _, i := range strings.TrimSpace(input) {
		n, _ := strconv.Atoi(string(i))
		ns = append(ns, n)
	}

	list := Build(ns)
	fmt.Printf("first %s\n", first(list, 100))
	fmt.Printf("second %d\n", second(ns))

}

func first(list CircularList, iterations int) string {
	list = play(list, iterations)

	values := list.Nodes[1].Next.Values()
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
	fth := list.Nodes[1]
	return fth.Next.Val * fth.Next.Next.Val
}

func play(list CircularList, iterations int) CircularList {
	curr := list.Start
	for i := 0; i < iterations; i++ {
		next1 := curr.Next
		next2 := next1.Next
		next3 := next2.Next

		curr.Next = next3.Next

		dest := curr.Val
		var destCup *Node
		for {
			dest -= 1
			if dest < list.Min {
				dest = list.Max
			}

			if dest != next1.Val && dest != next2.Val && dest != next3.Val {
				break
			}
		}

		destCup = list.Nodes[dest]
		destCup.Next, next3.Next = next1, destCup.Next
		curr = curr.Next
	}

	return list
}

func Build(list []int) CircularList {
	var start *Node
	var current *Node
	var max, min int
	nodes := make([]*Node, len(list)+1)

	for _, n := range list {
		node := &Node{Val: n}
		nodes[node.Val] = node
		if current != nil {
			current.Next = node
		}
		if start == nil {
			start = node
		}

		current = node
		if min == 0 || n < min {
			min = n
		}
		if max == 0 || n > max {
			max = n
		}
	}

	current.Next = start
	return CircularList{Start: start, Nodes: nodes, Max: max, Min: min}
}

func (n *Node) Values() []int {
	start := n.Val
	curr := n

	var values []int

	for {
		values = append(values, curr.Val)
		curr = curr.Next
		if curr.Val == start {
			break
		}
	}

	return values
}
