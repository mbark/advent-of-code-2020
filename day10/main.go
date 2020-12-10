package main

import (
	"fmt"
	"sort"
	"strconv"

	"github.com/mbark/advent-of-code-2020/util"
)

var input = `38
23
31
16
141
2
124
25
37
147
86
150
99
75
81
121
93
120
96
55
48
58
108
22
132
62
107
54
69
51
7
134
143
122
28
60
123
82
95
14
6
106
41
131
109
90
112
1
103
44
127
9
83
59
117
8
140
151
89
35
148
76
100
114
130
19
72
36
133
12
34
46
15
45
87
144
80
13
142
149
88
94
61
154
24
66
113
5
73
79
74
65
137
47
`

var testInput1 = `
16
10
15
5
1
11
7
19
6
12
4
`

var testInput2 = `
28
33
18
42
31
14
46
20
48
47
24
23
49
45
19
38
39
11
1
32
25
35
8
17
7
9
4
2
34
10
3`

func main() {
	lines := util.ReadInput(input, "\n")
	var adapters []int
	for _, l := range lines {
		n, _ := strconv.Atoi(l)
		adapters = append(adapters, n)
	}

	sort.SliceStable(adapters, func(i, j int) bool {
		return adapters[i] < adapters[j]
	})

	nrOnes, nrThrees := first(adapters)
	fmt.Printf("first %d (%d * %d)\n", nrOnes*nrThrees, nrOnes, nrThrees)
	fmt.Printf("second %.0f\n", second(adapters))
}

func first(adapters []int) (int, int) {
	nrOnes := 0
	nrThrees := 0

	start := 0
	for _, a := range adapters {
		diff := a - start
		if diff == 1 {
			nrOnes += 1
		} else if diff == 3 {
			nrThrees += 1
		}
		start = a
	}

	nrThrees += 1

	return nrOnes, nrThrees
}

func second(adapters []int) float64 {
	var chains [][]int

	chain := []int{0}
	for _, a := range adapters {
		curr := chain[len(chain)-1]
		if a == curr+1 {
			chain = append(chain, a)
		} else {
			chains = append(chains, chain)
			chain = []int{a}
		}
	}

	if len(chain) > 0 {
		chains = append(chains, chain)
	}

	var variations float64 = 1
	for _, chain := range chains {
		l := len(chain)
		if l < 3 {
			continue
		}

		var multiplyBy float64 = 2
		if l == 4 {
			multiplyBy = 4
		} else if l == 5 {
			multiplyBy = 7
		}

		variations *= multiplyBy
	}

	fmt.Println(chains)
	return variations
}
