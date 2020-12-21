package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/mbark/advent-of-code-2020/util"
)

var input = `
(4 + (2 * 7) * 4) + (6 * 9 + 8 * 4 + 7 * 3) * 3 * 5
5 * 9 + (5 * 9) * (6 + 2) + 3 + 7
8 + 5 * 9 * 9 + 9 + 5
(7 + 9 * 8 + 4 * 6) + 6 * ((9 + 9 + 5 * 7 + 4) * 8 * 2 + 5 * 6 + 2) + 2 * 6
(6 * (5 + 8 * 7 * 8 + 4) * (7 + 7 * 3 * 5)) + 5 * (8 + (8 + 3 + 5 + 5) + (3 + 2 + 7 * 2 * 9) + 6 * 5 + (2 * 6)) * ((4 * 3) + 3) + 9 * (3 + 6 * 2 + 3 * 8)
2 * 9 + (6 * 2 + 6 + (9 + 6 + 8 + 5 + 5) + (8 + 7 * 4 + 3) + 8) * (4 + 4 * 2 + 6)
5 + 4 + 5 + (7 + (9 + 7 + 4 + 4 + 4 + 5) + 2 * 4) + 4 * 9
9 + (7 * 8 + (6 + 3 + 4)) + (5 * 9)
2 * 4 + 8 + (5 + 8) * 9 * (9 + (5 + 3) + 3)
8 * 6 * 5 + 8 + (8 * (2 + 8 * 7 * 5 + 3 + 5)) * 2
6 * ((3 + 4 * 3 * 7 * 8) * 8 + (9 + 5 + 2 * 3) * 3 * 2) + 7 + 8 * 8 + 4
9 + 3 + 7 * 8
(9 + (4 + 9 + 3 + 5 * 5 + 5) * 5 + 9 * 6) * 3 * 2 * 6 * 5
7 * 9 * 6
(5 + 2 * 8) * 8 + 3 + 8
9 * 5 * (5 + (5 * 4 + 8 + 6 * 5) * (9 * 7 + 5)) * 3 * 9
(6 * (8 * 3 * 7 + 2) * (3 + 9 + 5 + 5) + 3) + 7 * 3
7 + 9 + (4 + 9 * (4 * 6 + 3 * 8 * 5) + 8 + 9 + 3) + 6 * (6 + 8 + 7 * 9 * 6 + 6) + 7
2 * (5 * 5 + 6 + (5 + 5 * 5 + 6)) * 4 + 5 + 2
8 + 2 * (8 * 7 * 4 + 8) * 9
9 + 5 * (4 * 8 + 9 * 6 + (3 * 3) + 2) + 5 + (4 * 5 + 2 * 2)
(5 * 7 + 2) + 3 + 9
7 * 9 + 8
3 * 7 * (7 + 5 + 8 * 5 * 8)
(7 + (3 * 6 + 6 * 9 + 2) * 8 * 6 + 6 + 2) * 6
4 * (7 * 5 * (6 + 8) * 7 * (5 * 3 * 2 * 9 + 2 * 4) * 6) * 2 + 2 * 5
((3 * 3 * 9 + 8) + 2 + 5 + 8) + 6 * (6 + 9 * 2) + 7 * 9 * ((9 + 4 + 4) * 4 + 4 + 9)
7 * (4 * 2 + 3 + 7 + 5) * (8 + 2 + (5 * 9) * 2 + 3 * 5)
7 + (8 + 3 + 8 * 4 * 2 * 2) + 7 * 2 + ((3 + 5 * 3 + 5) * 8 + 3) + 8
7 + 2 + 9 * 9 + 5
((6 + 4 + 2) * 2 + 9) + 3 * 6 * (6 + 4 + 9 + 2 + 6) * 9
6 + 6 * (9 * 9 + 6 + 6) + 2 * 6
2 + 5 * ((4 * 4) * 6) + (3 * 3 + 2) * 9 + 5
9 * 4 + 8 * 5 * ((6 * 5 * 4) + (3 + 8 * 7 + 6 * 7) * (6 * 6 + 7) * 8 + 4 + 7)
3 * 8 + (4 + (7 * 7 + 9 * 8 * 6 + 7) * (8 * 4 + 4 * 6 * 4 + 4) + 2 * 9)
3 + 6 * 2 * 3 * (4 + 7 + (3 * 6 + 7 * 8) * 7 + 9 * 8)
2 + 8 + (5 * 4 * 2 * 5 + 5) * 4 + 3 * 4
6 + 4 + (5 + 5)
6 * 3 + 8 + (8 + 2 + 9 * 3) * ((6 * 8 + 6) + 6 + (7 * 4 * 3) * 4 + 9 + 3)
((9 + 7 + 9 + 2 * 7 * 7) + 8 * 6) + 2 + 7
6 * (6 + 9 * 2 + 4 * 7 + 8) + 5 * 9
((2 + 5 + 6 * 3) + 7 + (7 * 6 * 8 * 3 + 8 * 3)) + (8 * 6 * 8 * 9 * 5)
6 * 7 * 5 + 5 * (2 * 9) * 4
(8 * 7) * (2 * (7 + 3) + 3) + 8
(7 * 9 + 7 + 7 + 7 + 7) + (7 * 5 * 8 + 4 * 5) + 6
5 * (6 * 7 * 8 * 5) * 6
5 * (2 + (6 * 7) * 8 * 3 + 6 * 2) + 2 * 4 + 8 + 3
3 * 5 + 6 + 6
((9 * 9) * 6) + (4 + 9 * 6 + 4 * 9 + 7)
(3 + 9 + 2) * (3 * 7 * 7 * 7 * (8 * 8 * 4 + 3 + 8)) + 2
6 * 6 + 8
((9 + 5 + 8) + 9) + 7 + ((9 * 5) * 3 * (3 * 4 * 4 + 3) * 7) + ((5 + 6 + 2) * (7 * 2 * 9 * 8) * 5)
6 + (8 * 7 * 6 + 5) + 9 * 7 * 7 * 7
2 + (4 * 9 * 3 * 9 + 8) + 2 * 7 * 3 + 4
7 + 3 + (5 + 7 * (4 + 4 * 8) + 2 + 8 * 6) * 2
(2 * 5 * (4 * 7) * 9 * (3 + 7)) + 3 * 4 * (2 * 3 * 8 * 4) * ((6 + 2 * 2) + (6 * 5 * 7 * 6) + 9 * (8 + 4))
((9 + 3 + 5) + 8) * 2
2 + 4 + 7 + 5 * (6 * 7 + 8 * (4 * 3 + 6 * 4 + 7) * (7 + 3 + 7 + 3 * 2) + 8)
5 + (9 * 8 + 5)
(2 + (9 + 7 * 8) + (6 + 7) + 6) * 5 + 3 * 2
7 + (9 + (8 * 7 * 9)) * 8 + 3
9 * (9 + 4 * 2 * 3 + (6 + 7 * 8 + 9)) + 6 * ((7 * 4 + 7 * 4) + 2 * 6 * 5 + 6 * 6)
((3 * 3 + 4 * 9) + 7 + 7 + 3) * 8
9 + 8 * (2 + 6 + 8 * 6 + 3 + (8 * 6 * 6 * 4 * 6))
3 * (9 + 2 * (9 + 2 * 8) + 3 + 8 * 5) + (9 + 8 + (6 * 9 + 2 + 8 * 5 + 7) * 7) * 2
((2 + 9) + 3) * 6 * 2 * 2 * 4 + 7
2 * 3 + 4 + (7 + (7 * 5 * 5 + 5) * 9 + (3 + 3 * 5) + 3 * 8)
4 * (8 + 2 + (3 * 7 + 4 + 2 * 6 + 2) * 8 + 7)
4 + (5 * 9 * (8 + 3 * 9 + 9 + 5) * (3 * 8 + 9 * 8)) * 8 + 9
4 * 5 * (2 + 2 + 7 + (3 * 7 * 4 * 2) + 3) * 6
((8 * 8 * 4 + 4 + 5) + 6) + 9 * (3 * 3) * 4
9 * (8 + (6 + 2 * 7 + 3 + 8 * 8) * 3 * (9 + 3) * (3 + 7 * 3) + 5)
9 * 3
6 * (7 + 8 * 8 * 4) * 6 * 9 + (4 * 7) * 4
4 * (3 * (4 * 9 * 6) + 5 + 5) * 4 * 6 * 6 * ((9 + 3 * 6 + 5 + 8) + 4)
5 * 2 * (4 * 3 + (8 * 4 + 2 + 2) * 6 * (8 * 8 * 2 + 7 + 9) + 3) * 5 + (4 + 6)
3 + (5 + 6 * 5 * 9) * 9 + 7
((3 + 6 + 8) * 7 * 3 + 4 + 4 * 8) * 2 * 2
2 + ((9 + 9 * 4 + 2 + 3) * 4 + 9 + 5 + (7 + 3 * 9)) + 8 + 7 * 3
5 * 5 * (5 * 6 * 6 + (3 + 4 * 8 + 7) * 2) * 6 + 4 + 3
4 + 3 * (5 + (6 * 5 + 6 * 2) * 4) + 9 + 7 * (5 * 5 + 9 * 5)
2 + (6 * 7) + 7 + 3 * 7
5 * 7 * (6 + (9 * 5)) * 4 + 2
((7 * 7 * 3) + 2 + 2 + (9 + 8 + 3)) * (3 * 8 + 6) + ((6 + 2 * 6) * (2 * 4 + 3 + 8 + 2)) * 8
((8 + 3 + 8 + 5 + 8 * 3) * 2 + 4 * (6 + 4)) * 8 * 2 * 5 + 9
7 + 6 * 3 + 3 + (5 * 7 + 3 + 3 + 8) * 6
8 + 7 + ((5 * 9 * 3 + 3 + 4 + 6) + (6 * 6 + 2 + 6) + 8)
5 + 5
6 + (6 * 5 * 7 + 8)
(8 + (9 + 2 + 8 + 2 + 2) * 4 + 7 * 6) * 9 * 6 * (6 * 6 * 5 + 5 * 4) + 7 + 9
3 + 8 + (3 + (6 * 9 + 9 * 2 * 3 * 4) * 3 + 9 * (4 * 2) + (3 * 2 * 9 + 2))
(5 * 8 + 7 * 8 + (2 + 3 * 8) * (8 + 4 * 7)) * 3
8 + 2 * (2 + (6 * 6 + 9 * 3 + 6) + (9 * 2 * 4) + 6 * 4 * (6 * 3 + 2 + 6 + 3 + 5)) + 8 + 4 * 9
8 + 8 + (4 * 9 + 9)
5 + 5 * (5 + 8 + (3 * 9 * 9 + 8 + 8 * 2) * 5) * (8 + 7 + 7 + 8 + (2 + 7 * 4 * 8))
6 * 3 + 6 + ((9 * 5 * 4 + 6 + 5 + 5) * 6) * 2 + (6 * 5 + 5 * 7 * 8 * (3 + 7 * 7 * 9))
4 + (3 + 4 + 5 * (4 * 2 + 7 * 9 + 6 + 6)) + (7 + 7)
2 * (6 + 2 + 6 * 3 + 4) + (6 + 5 * 4 * 8 + 3 + 7) + 8 + 3 * 5
((5 + 3 + 7 * 2 + 9 * 5) + 7 * 7 * (4 + 7 + 9) + 8 + 6) + 3
((9 * 9 + 2 * 7 * 5 * 6) + 8 * 7 * 4 * 2 + 6) * 5 * 9
((3 * 3 * 2 + 5 + 6) * 7 + (3 * 6 * 2) * 3 + 3) + 9 * 7 + 5
(7 + (4 + 7 + 6 * 8 + 9 * 9) + 7 * 2 * 2) * 7
(9 * 3 * 9) + ((2 * 6 * 6 + 9 * 7) * 3) + 3 + 8 + 8
2 + (3 * 5) * (2 + 8 + 8 * 3 * 6 + 6) * 6 * 9
(4 * 7 * 7 + 3 + 4) * 3 * 9 + 6 * (2 + 5 * 4)
(8 * 4 + 3 * 8) * 4 * 5
(7 + 5 + (4 + 2) + 4) + (2 * 5 * 9 * 2 * 6)
9 * 4 + 3 + 8 * (2 * 4 + 7 * 2 * 2) * (2 + 6 + 5 + 9 + 5)
8 + 3 * (6 * 8) + (2 * 4 + 7 * 8)
(9 + 6 + 3) * 5 + (9 * 8) * 5 + 6 * 5
((6 * 2 + 6 * 2) + 3 + 7 + 5 * 8) * 2 + 2
2 + (9 * 3 * 9) * (4 + (6 + 4 + 7) + 5) + 9
(3 + 4) + 5 + 7 + 7
9 + 8 + 8 + 7 + 5 + ((2 + 7 * 9 * 7 * 6 * 9) * 8 * 8 + 7 * 3 + 9)
(8 * 7 * (8 * 7)) * 9 + (4 * (4 + 9) * (7 * 9 + 9 + 6 * 5 + 5) + 7)
(6 + 7 * 9 + 5 + 5) * 5 * (2 + 6 + (3 * 6) + (7 + 3 + 5 * 5) + 9)
2 + ((7 + 8 * 2 + 8) * 4 + 2) * 9 * 9 + 8
8 * 8 * 7 * 2 + 2
7 * 2 + 9 + ((5 * 2 + 6 * 9 + 8 + 7) + 6 * 2 + 8 + 4 * (3 * 2 * 3 * 2))
5 * (7 * 5 + (7 * 9) * 4 * 6 * (8 + 7 + 3 + 9 + 9)) * 4
8 * ((3 + 7 + 5) + 7) + 5 + (6 + 4 + 6 + 4 * 4) + ((2 * 6) * 3) + 9
3 * 3
4 * (5 + 5 + 6 + 2 + 7) + 7 * 6
(7 + 4 + 5 * 9 + 5 + 8) * 8 * (7 * 9 * 9) + 4
3 + (9 * (7 * 7 + 3 * 4 + 8 + 5) * 7 * (2 * 4) + 9 + 5) * (4 + 7)
6 * (5 * (6 * 2) + 6 + (4 * 8 + 9) + 3 * (2 * 4 + 4)) * 6 * 5 + 6 + 3
(9 * 3 * 9) * 8 * 5 * (8 + 3 + 2 + 9)
(8 + 8 * (8 * 9) + 2 + 7) + 8 * (5 * 7 + 6 + 2 * 4 * 2) + 6 * (3 * 8 + 6) * 2
2 * (8 * 5 + 7) * 8 + 7 * 9
4 * 4 * (5 + (5 * 4 + 2 * 9 + 5 + 7)) * 8 * ((4 * 9 + 4) + 6 * (7 + 2 + 4 * 8 * 9))
4 * 6 + (4 * 8 * 3 * (9 + 8) * 7 * 9)
(6 + 9 + 3 * 2 + 8 * 8) * 6
(8 + (3 + 3 + 2 * 3) + (6 + 7 * 6) + 7 * 4 + (9 * 5 * 4 * 6 + 2)) * 2 + 5
4 * 6 + 2 * 5 * 2 + (5 + 3 + 2 + (5 * 3 * 3 + 7 * 6 + 2) + 8)
9 + 2 + (7 * 7 + 7 + 9)
5 + 8 + (8 + (5 + 9 + 6 * 2 * 4 * 4)) * 9
(2 + 4 + 8 + 7 * 8 * 5) + 3 + 3 + (9 + 9 + (9 * 2 + 6 + 4 + 4) + 6 * 2) + 2
7 + (6 + 5 * 2 * 9 + 7 + (9 + 7)) * (3 + 6 + 2 * 5 * 3)
8 + (9 + (4 * 3 * 7 * 7) * 7) + 9 * 6
6 * 9 * ((8 + 5) + 9) * 7 * 8
(2 * 9 + 2) + (6 * 3 * 3 + 6 * 7 + 9)
7 * 5 * (4 * 4 + 7 * 4 + 8 + 3)
3 * (6 + (5 + 3) + 5) * 3
(7 * (9 * 5 * 9 + 2 + 4) + 3 * 8 + (9 + 7) + (2 * 9 + 5)) * (6 + 3 * 9 + 3 + (7 + 4 * 9 + 5) + 9) * (8 * 2 + (2 * 9 + 7 + 5 * 4) * 4)
4 + 7 + 5 + 6 * 8
5 + 4 + 3
9 * (6 * 2 + 7 + 5) * 3 * 2 * 2
7 * 4 + ((6 * 7 * 6 * 8 + 6 * 9) * 7 * 3) + 8
9 * 5 + 4 + 8 + 8 * 8
(3 + 6) * 2 * 7 * (2 * 7 + 2 + 6 + 3)
5 * 2 + 8 * 7 + (4 + 8 + (3 * 6 + 7 * 2 + 4 + 7))
6 + (8 + 6 * (9 * 7 + 2 + 2 + 3 + 4) * (4 * 7 + 2 + 4 * 3) * 3)
(8 + 5 * 2) * (9 * 4) + 5
3 * (4 * 5 * (4 + 3 + 4) + 5 * (3 + 6 + 3 * 7 * 7)) * 3
2 + 7 + (8 + 3) + 8
8 * (9 * 8 * 3) + 4
(2 * 6) * 3
((9 * 6 + 3 + 3 * 9 + 7) + 2 * (2 + 9 * 8 + 2 * 4) + 9 * 8 + 6) + (5 * 9 * (2 + 7) + 7 * 9) + 7
4 + 7 * ((3 * 7 + 2) * 9 * 7 * 2 + 4 + 3) * 4 * (3 + 8 * 9) + 6
9 * (9 + (9 * 6 + 2 * 7 + 6) + 7 * 2 + 2)
(9 * 2 + 4 + 8 + 6 * 8) * (4 + 8 + 7) * 5 + (4 * 2 + 2) * ((4 * 5) * 8) * 2
(5 * 9 + 9 + 2) * 3 + ((7 * 5 + 9 * 3) + 2 + 3 + 6 + 7) + 4 * 6 * 6
9 * 8 + 3 * 4 * (9 + 5 + 7 + 5) * (6 * (7 + 4 * 6 * 7) + 6 * (6 * 5 * 3 + 9 + 4) + (2 * 3 + 3 + 6))
7 * 6 * 2 + (8 * 9 + (6 * 7 * 8 + 6 + 9 * 8) * 3 + 6)
5 + 7 * 8 + 2 + ((2 + 4) * 7)
2 * 5 + 6 * 6 * (3 * 6)
6 * 6 * 6 * (6 * (4 * 4 + 5) * 4 + 3 + 9) + 4
((3 * 4 * 7) * 5 * 9 + 7 * 6 + 5) + 2 * 2 * 6 + 4 * 8
7 * 2 * (3 + (6 * 9 * 9) * 4) * (5 * 8 * 7 * 6 + 2 * 9) * 7
9 + (6 + 6 + (8 + 8 * 3 + 6 + 6))
5 * 2 + 7 + 3 * (9 + 9 * 9 * 9 * (4 * 4)) * 2
4 + 6 + 3 + (2 + 2 * 7 * 2) + 5 + 7
7 * 9 * 9 * 2 * 4 + 7
6 + (6 + 5 + 5) + 7 + 4
9 + (9 * (7 * 3 + 9 + 7 + 8 * 8) * 6 * 5) + 2 * 5
8 + 6 * 8
8 * (9 * 6 + 4 + 8 + 8)
((8 + 3) + 8 * 2) + (3 * 6 * 7 + 7 + 4 + 2) * 9 * 5
6 * (4 * 5 + 6 * 7 * 4) + (2 * 8 + 4 + 7 + 7 + 6) + 9 + 5
9 * 8 * (2 + 7 + 9 + 4 + 5) + (8 * (3 * 8 + 3 * 6)) + ((7 * 2) * 7) + 2
9 + 4 + 9 * 2 * (5 * 2 + 6 + (9 * 5 + 8 + 2 * 5) + 6 + 6) + 7
9 + 9 * ((3 * 6) + (5 + 5 + 7 * 9 * 4 + 9) + 2 * 6 * 6) * 5
(9 * (3 * 2) + 8 * 2) * (6 * (8 + 3 * 8 + 5 * 5 * 9) + 4 + 9 * 6) * 5 + 9 * 6
7 * 9 * (2 * 9 + 4 * (3 * 2 * 6 * 9 + 5 + 6) + 2) + 7 * 6 + 3
9 + (5 * (3 * 9) + 4 * 6 + 6) + 9 + 8 * 5
8 + 8 * 8 + (9 + 4 * (5 + 3 * 8 * 3 * 4 * 4) + 2 * 8 * (9 * 6 * 4 * 6 + 9))
(6 + 6 * 3 + (5 + 8 * 4 + 4) + 3) * (3 + 7) * (6 + 5 * 4 * 6 + 6 + (5 * 7))
(9 + 7 + 6 + (8 * 8) * 4 * 6) + 3 * 5
5 * 2 * 3 + (3 + (6 + 9 * 2) * 3 * 3)
2 + (8 * 7 * (4 * 8 * 2 + 6 * 7)) * 5 + 5 + (8 * 5)
9 * 6 * 4 * 2 * ((6 * 3 + 9 * 8) * 5) * 3
((8 + 6) + 6 * (8 * 5 * 5 * 7) + (2 + 9 + 6 * 7) * 7 * (2 + 9 * 7 * 3)) * ((7 * 5 + 2 * 5) + (8 * 7))
(7 + 8) * 6 * 9 + 6 + 6 * 8
7 * 8 + (8 * 2) * 7 * 8 * 9
9 * ((8 + 6 * 9 + 2 + 3 * 2) * 2 + 4 + 3 + 7) * 2 * 6
2 + (8 + (2 + 9) * 7 * 4 + (8 * 6 + 8 * 6 * 4 + 5) + 6) + 4 * (9 + (3 + 4) * (3 * 9 * 7 * 6 + 9 + 7) + 9 * 6 * 3)
2 + (8 + (8 * 2 + 3 * 8 * 4 + 2)) + 3 * 6
(5 * 5 + (3 + 5 * 6 * 9 + 4 * 4)) * 8 * (7 * 7 + 3 * 7)
9 + 6 * (9 * 2 + 2 * (9 + 6 * 8 + 5 + 3))
5 * 5 + 5 * (7 * 3 * 3 * 5) + 6
7 * 5 * 7 * ((6 * 3 + 5 * 8 * 6) * 7 + 4 * 5) + 7
(5 * 4 + 8) * 3 + 6 + 9 + 6
(7 + 7 + 4 + 8) * (2 * 7 + 9) * 2 * 4
2 + ((7 + 2 + 9 * 4 + 2) * 6 + 9)
9 + 2 + (7 + 6 + 4) * 3 * 8
5 + 7 * 2 + 2 * (3 * 4 + 4 * 6) + 2
2 + (7 + 2 + (9 * 3 * 5 + 7 + 7) + 5 * (7 + 7 * 2 * 4 * 7 * 4) + 5) * (6 + 2 * 5 + (4 * 6 + 4 * 8 * 4) + 7) * 2 * 5
((5 * 9) * 9) * 2 + 4 + ((3 + 6) + (9 + 3 + 6 + 3) * 3 * 5 + (9 + 9 + 6 * 7) + 7) * 2 + 9
4 + (4 + 4 + 6) + 6 + 5 * 9 + 8
(3 * 8 * (7 * 5 * 7 * 5 * 7 + 2)) * 7 * 4 + 3 + 2 + 4
(4 * 8 + 7 + 8 * 5 * 9) + 2 * (5 * 2 * 3 * 9) + 7 + 7 * 2
6 * 3 * (4 + 6) + 7 + 2
(5 + 4 + 4 * 8 + (4 * 6 * 5 + 7 * 4 * 2)) * 9 * 6 * 6 * (3 + 5 + 9 * 7)
4 * (5 * 2 + 8 * 9 * 3) + 3
7 * 3 + 7 + 5 + ((6 * 8 * 8 + 8 + 2 * 4) * 8 * 2 * 9)
4 * 2 + 6 + 2
(2 * 2 + 2 * 9 * 6) + (5 + 9) * 3 * 9 * 5
(5 + 5 * (3 + 3 + 7) * (3 * 7 * 7)) + 4
9 + 7 + 9 + 5 * 6
(6 + 9 + (4 + 7 * 7 * 4) + (9 * 2 + 9 + 5 + 6)) + 7 * 8
5 + (5 * (6 * 5 + 8 + 2) + (4 * 7 * 4 * 7 * 7 * 8) * 8)
4 + (6 + 6 * 2 + 9) + 7 * ((3 + 8) * (3 * 5 + 5 * 3 * 3 * 6) + (4 + 9 * 3 + 3 * 2 + 9))
(7 * (7 + 7 + 5) * 8 * 2 * 8 + 3) + (2 * 4 + 8 * 7 + 3 + 2) + 7 * (7 + 2 + 2)
(3 * 7 * 4) * 5 + 4 * 7 * 8 * 4
7 * 3 * (2 * 7 * 9 * 7 * (8 + 4 + 2) * 3) + 7
7 * (9 + 2 + 5 + 5 + 7 * 5)
5 + 2 * 9 * (8 * 6 * 8 + 2 * 7)
4 + (2 + (2 * 5 * 7 * 5) + 3 + 6) + 8 * ((3 + 4 * 4 + 5 * 3 + 6) + 2 + (5 * 8)) + 8 + 3
2 * (9 * 7 * 3) + (8 + 7 + 8 + 2 * 3 + 9) + 8 + 9
8 * (7 + 5 * 3 + (2 * 6 * 6 * 4) + 9) + 3 * 4
3 + 4 * 8 * 9 + (6 + (5 + 7 * 7 * 4)) + 7
(5 + 6 + 8 + 5 * 7) * 2 + 5 * (9 * (9 * 9 + 8 + 5) + 7 * 4)
8 + (8 * 2 * 4 * 2 + 3 * 2) * (3 * 3) * 6 + 8
(6 * 2 * (6 + 5 * 9 * 3) * 6 * 2 + 7) + 7 * 8 + 2 * 4
5 + (2 + 4 + 5 * 4 + 6) + 9 * (5 * (5 + 9) * 9 * 9 * 2 + 8) * 5
7 * 4 + 9 * (9 + (5 * 7 + 8 * 7 * 5) + 9 + 2 + 2 + (2 * 4 * 7 + 5))
7 + 2 + (4 + (8 + 8 + 5 * 9)) + 9
(8 * 3 * 9 + 6) * 8 + 8 + 7
(2 + 3 * 9 * 4) * 7 + (8 * 6 * 9) * (8 + 6 + 9 + 4)
((7 * 6 + 4) + 6 + 6) + 5 * 4 + 3 * 4
5 * (2 + (3 * 9 + 9 * 6) + (7 + 3 + 2 * 4 + 9)) + 7 + 8 * 5 * 5
7 * 9 * (5 + 6 + (7 * 6) + 4 + 4 + 3) * 2 + 9
(3 * 2 + 5) * 8 * (2 + 4 + 2 + 8 + 7 * (8 * 9))
((2 + 7 * 8 * 9) * 4 * 3 + 6 * 7) + 9 * 8 + 5 + 9
5 + 4 + 9 + 8
7 + 7 + 9 + 2 + 5 + ((6 * 3 + 3 + 6) * 4 * (4 + 4))
(9 * 4) + 8 * 4
3 * 9 * (9 + 5 * 4) + 8 + 9 * 3
3 + (9 * 8 * (6 + 4 + 6 + 5 * 2) + 9 + (4 * 7 * 7) + (3 + 4 + 6 * 8 + 9)) * (2 + 8) + 7 * 5 * 6
9 + 6
6 * 8 * 8 + 6 * 6 * ((9 + 3) + (4 + 2 * 8 + 8 * 2 + 5) * 2)
4 + (4 * (2 + 2)) + 7 + 4 + (4 * (8 + 8 + 4 + 4 + 2) + (3 + 2) * 9 + 8) + (4 + 9 * 9)
(2 * (3 * 7 + 6 + 2)) + (4 * 6) + 6 + (6 + 9 + (2 * 2 * 2 + 6 + 4) * 4) * 5 + (4 + 2 * 5 * 8 + 6 * 2)
4 * ((5 + 5 + 9 + 3 * 5) * 7 * (7 + 9 * 5 * 7 + 2) + 5 * (4 + 7)) + 7 * 9 * 6 * 4
4 * 7 + 4 + ((4 + 8 * 6 * 3 + 2) + 3 + 8 * (6 * 7 * 2 * 8) * 7) * 9 + 6
4 * 6 * ((3 * 6) + 4 + (2 * 2 * 2 * 7) + 9 + 8) * (3 + 8 * 2 * 4 * 2 * 3) + (4 + 3 * (7 + 5) + 6 + 9) + 5
(7 + 2 + (8 + 6 * 4)) + 8 + (7 + 5 + (2 * 5 + 9) * 5)
(7 * 8 + 5 + 8 + 8 * 2) * 6 + (5 + (8 + 3 + 4 * 5) * 9)
8 * (4 + 5 * 9 + 6 + 8 * 7) * 3 * 6 + 4 * 5
6 + ((6 * 3 * 4) + (8 * 8 * 2 * 5 * 8 + 2)) + 3
9 * 3 + 6 + 5 + ((5 * 2 + 9 * 5 + 6) + (4 * 9)) + (4 * 2 * 3 + 7)
6 + 8 + (6 + 2 + 8 + 2) * (5 + 8 + 5 * 5 + (4 + 9 + 3))
(2 * 7) + (6 + 4 + (8 + 6 + 5 + 9 * 4 + 7))
8 * 3 + 3 * (3 * 5) * 6 + 2
9 + 4 + ((5 * 4 + 2 + 6) + 5) * 5 + 6 + 4
(9 * 9 * 3 + 5 * 3 * 4) * 3 + (2 + 3 * 5 + (3 + 6 * 6 + 2) * 8 + 3)
(6 + 9 * 9 + 7 * 6) + 7
4 + 6 * (3 * 3 * 7 * 7 * 5 + 2) * (7 + 7 * 2 * 5 + 4 + 9) + (8 * 2 + (8 + 8 * 4 + 7 + 3 + 5) * 7) + 9
4 + 8 + 6 * 7 * 7 + (2 * 8 + 3 + 5)
4 * 8 * 4 * (9 * 9 * 5) + 9
8 * 5 + (2 + 3 + 2 + 8 * (3 + 5 + 8 + 7 * 9))
3 * ((8 * 6 + 9 * 8 * 5) + 9 * 6 + 6) * 6
5 * 8 + 6 * (5 * 2 * 8)
8 + (6 + 7 + 7 * (6 + 7 + 3) * (5 + 2 * 7) + 7) + (2 * 8)
6 + 5 + 4 + 2 + ((3 + 8) * (5 + 4 * 9) * (4 + 9 + 3 + 9 + 9 * 7) + (5 + 2 + 9 + 3) + (6 + 6) + 3)
3 * (2 + 5 * (5 * 6 * 5 + 9) + 9 * 5) * 8 * 5 + 4
2 * 4 + 6 + (9 * 4 * 3 * 4)
(4 * 6 * (3 + 9) + (3 * 5 + 8 * 5 * 9) + 6) + 9 + 8
(3 * 5 + 8) + 8 * 8 * 7 + 8 * (9 + 6)
(5 * 4 * 6 + 3 + 3) * (3 * 8 * (7 + 2 * 9 * 6)) + ((9 * 9 * 4 + 8 * 7 * 7) + 4 * 4) + 8 * 6
3 + 5 * 5 * 6 * ((3 + 3) * 8 + (6 + 9 + 2 + 6 * 7)) * 4
9 + ((4 * 7 + 2 + 2 * 6) + 8 * (4 * 6 + 7 * 5) * 7 + 4) + 6 + 6
(9 + 9 * (8 * 2 + 9 + 5) * 3 * 4 * 4) + 5 * ((3 + 6 * 3 * 7) + 6 * 4) * (3 + 8 * 9 + 6 + (4 + 3 * 4 + 2) + 6)
5 * 5 * (9 + 8 * 3 * 3 + 3) * 4 + 8
3 + (6 * 9 * 8) + 9 + (3 + 3 + 2 + 5 + 2)
3 + 4 * (4 * 8 + 2 * (3 + 3 + 4 + 9 * 7) + 4 + (6 + 9 + 4 * 6 + 9))
6 * 5 * 9 * ((6 + 9 * 6 + 7) + (6 * 4 + 2 * 4 * 6 * 6) + 9 + 6 + 8)
7 + 7 * 2 * 4 + 7 * 2
2 + 8 * (7 + 3 * 3 + 2) + 6 + 3 * 7
2 * 2 * 9 + 8 + 8
9 + 5 + 4 * (6 * 7 * 5 * 7 * 4 * 3) * 3
7 * 5 * 7 * 6 + 4 * 9
((9 + 7 + 9 + 7 * 7 * 4) * 6 * 5) + 5 + (8 + (2 + 6 * 9 * 7 * 9 + 9) * (9 * 3 * 8 + 9 * 6 * 2) + 6 * 7) * (2 + 6 + (3 * 4 + 4) * 6 * 5) * 9 * 5
(7 * 5 * (3 + 2 + 6 * 5 + 9) + 5 + (2 + 9) + 2) + 2 * 6 + (5 * (7 * 2 + 3 + 8 * 9) * 7 + 6 + 6 * (5 + 3 + 9)) * (4 * 6 + (9 + 7 + 4))
((3 + 5) + 7 + 3) + 6
3 * 4 + 4 * 6 + (5 + 8 * 2)
7 + (7 * (4 + 4)) * 3 + 7 * 4
(5 + 4 * 6 + 3 * 9 + 7) * (7 * 3 * 5) + 6
3 + 4 * (8 + 2) + ((6 * 7) + 8 + 8 * (9 * 3 + 2) * 9)
7 + (7 + 9 * 2 * 9 * 2 * (2 + 5 + 9 * 6)) * 4
(7 + 6 + 3 + 5 * 2) + 2 * 3 + 4 * (3 + 4 * 5 * (3 + 7) + (7 + 2 * 7 * 6) + 8) * 8
4 + (9 * 2 * 6 * (6 + 6 * 7 * 7 + 5) * 9) * 9 + (4 + 8 * 7) * 4 + 4
6 * 6 * 2 + (4 * 2 + 7 + 4 + (2 + 2 * 5)) + 2
8 + 2 + 3 + 3 + 8 + (6 + 4)
8 + 9 + 2 + ((3 * 9 * 3 * 5) * 3 + (2 + 2 + 8 + 6 + 9 + 7)) + 9
6 * (3 * (5 * 4 + 4 * 4 * 4 + 4)) * 8 * (5 + (5 + 2 * 7 + 8) * 3)
3 + ((3 + 7) + 2 + (5 * 3 + 4 * 9))
4 + 5 * 6 * (7 * 5 + 6 * 5 * 5 + (4 * 6 * 9 * 6)) + 7 + 3
7 * 4 * 7 * 5 * (3 + (5 + 6 + 9 * 9 * 3 * 7) * 2 * 8 + 2 + 6)
5 * 4 + (4 * 7 * (3 + 8 + 9)) + 4
(7 + (3 * 5 * 7 + 5 * 8) * (6 + 9 * 6 * 7) * (3 + 9 * 8 + 4) + 8 + 7) * 2 + 4
8 + 7 + (9 + 6 + 7 + 9 + 6 + 4) + (7 * 8 * 6 * 5 * 3) + (3 * 4 * 3 * 8)
7 * ((4 + 4 + 9 + 3 * 3) + (9 * 6 * 8 * 7 * 3 + 6) * 7 + (7 + 2 * 3 + 6 * 2)) + 3 + (5 * 8) + (4 + 9 * 6 * (2 * 5 + 7 * 3 * 5) * 3 * 8)
8 + (3 + (8 + 3 * 4 + 4)) * 2 * 4 * (3 + 4)
(4 + 9 + 4 + (4 + 7 * 3 * 2 * 5 + 7)) * 2
2 * 8
6 * (7 * 5 + (4 * 5 * 2))
(3 + 3 * 4 * 3 + 7 + 4) + 8 + 2 + 4 * 4 * 5
7 * 2 * 5 + (2 * (9 + 6) * (3 * 8 + 4) + 3) * 9
(6 * 9 * (6 + 6 + 7 + 7) * (5 * 7 + 8 * 5 + 2 * 8) + 5 + 8) * 9 * 9 + 3 * 2 + 7
(2 * 3 + 8 + (4 * 2 + 7 * 2 + 7 + 7) + 3 + 6) * 3 * (6 + (9 + 7 * 2 + 2) + 6) + (4 + 3) * 3
4 * (6 * (8 * 8 + 7) + 7) * 6 * 9 * (4 * (6 * 4 * 3 * 2 + 6) + (8 * 9) + (9 + 3) * 4)
5 * 7 * 2 * (2 * 2 * 5 * (8 + 6))
(8 * (7 + 7 + 8)) + 2 + 9 * 9 + 9 * 4
6 + 6 + 5 + (2 * (8 * 7 + 3 * 6) + 9) * 7 + 6
8 + 6 + (8 * 4 * 8 * 2 + (6 * 6 + 9 * 3 + 6)) + 3
9 * 5 * 4 * (5 + (3 + 7) + (4 + 4 * 5 + 5 * 8 * 9) * 5)
5 + 4 * 9 * ((2 * 4 + 8 * 2 * 7 + 2) + 4) * 9 * 3
3 + 6 * (6 + 2 + 3 * 2) + (3 + 5) + ((3 * 6 + 5 + 7 + 8 * 8) * 9 + (5 + 8 * 7 + 6 * 8 + 7) + 7 + 8 + 5)
7 * ((7 + 5 * 3 * 5 + 4) + 7 + (6 * 3 * 6 * 6 * 7 * 7))
((5 + 7 * 9 * 2 + 3) * 2 + 9 * 3) * (4 + 6) * (4 * 7 * (9 + 6 * 8 + 5) * (3 + 8 + 9) + 9 * (8 + 6 * 9 + 6)) * 6 + 4
2 * (6 * 3 + 4 + 4 * 2) + 6 + 7 + 5
2 + 7 + 5 + 5 * ((7 * 6 + 2) * 8 * 8) + (2 * 5 * 4 * 5)
3 + (7 * 8 * 5) + 5 + (5 + (4 * 4 + 7 + 2) + 6 + (8 + 3) * 9 + 4)
6 * (9 + 5 + 6) * 5 + 5 * (7 * (3 * 5 + 6 + 2 * 5) * 2)
(9 + 6 + (9 + 7 * 5 + 6 * 7) * 7 * 3) + 3 + 4 + 6
7 * ((5 + 4 + 6) + 7) + 2 * 3 + (6 + 9 * (6 * 4 * 8 * 3 + 4 * 2) + 7) + 3
7 * (8 + 9 + 5 * (9 * 7 * 9) + 2) + (7 + 7 + (7 + 3 + 2 * 2) + 4 * 2)
3 * (2 * (4 + 6) * 8)
8 + ((6 * 4 * 8 + 4) + 7) * 7
(7 + (2 + 8 * 8 * 6 * 2)) + 7 * 2
(4 * 2 + 4 * 8 + 4) + (3 * 9 + (2 * 4 + 2 * 7) * 9 + 7) * (5 + 9) * (7 + (3 * 4 * 4 * 6) + 8 * 7) + 8 + 8
8 + (2 + 2 * 8 * 7 + 7) + 2 + 4 * 3 + (8 + 9 * 9)
3 + 2 + 8 + (9 + 3 * (3 + 2 + 9 + 7) + 4 * 4 * 7) * 3
5 * 4 * 7 * 3 + (9 * (6 * 6 + 2 * 4) * 8 + 2 * 9)
4 * 7 * (3 + 6 * 6 + 4 + 8) + 5
(8 + (8 + 2 * 9 + 3 * 3) * 8 * 8 * (9 * 6 + 2 * 6)) * 9 + 5 * 2 + 5
5 + 6 + 3 + 4 * 9 + (4 + (4 + 3 + 9 * 2) + 4)
3 * (5 + 3) * 7 + 6
((3 * 2 + 2 * 6 + 5 + 3) + 4) + (4 * (8 + 2 + 9 * 2) * 8 + (8 * 3 + 6 + 2 + 6) * 2) + 7 + 2 * 9
6 * 8 * ((8 * 8 + 3 * 7 + 8 * 4) * 4 + 9) + (8 * 7 * (6 * 8 * 3 * 2) * 3) * 2 * 8
(2 * 7 + 4) * 2 * (5 * 4 + 2 + 5 + (7 + 5) * (4 * 3 + 6 * 8 * 5 + 9)) * 5
(6 * 8) * (7 * 7 + 8 * 3 * 3) * (5 * 8 + 2)
(6 + 6 * 2) * 5 + ((7 * 6 * 8) * 8 * 9 * 4)
(8 * 4 * 9) + ((5 * 3 + 6 * 3) + (9 + 9 + 9 * 2 * 2) * (4 + 9 * 7 + 2 * 2 + 3) + 6) * 7 * 6 * 6
((4 + 3) * (2 * 5 + 8 * 2 * 4) * 9) + 6 + (6 + (5 * 4 * 5 * 3 + 3) * 6 * (5 * 3 + 7 * 2 + 4 + 8)) + 4 + 7
(6 * 9 + 9 + 7) * (6 * 2 * (4 * 7 * 4 + 4 * 5 + 7) + (7 + 5 * 8 * 4 + 9 + 6)) * 7 + 8 * 5 + (9 * 5 * 2 + 7)
6 * 6 + ((2 + 9) * 5 + 2 * 3 + 3 * 3) + 9 * ((4 * 7) + 5 + 8 + 4)
5 + (4 * 5 + 5 * 3 + 7) * 6 * 5 + (2 + (9 + 7 + 8 + 4))
(2 * 4 + 5 + 8 + 6 * 9) + 3 + 6
9 + (9 + 8 * 3) * (3 + 8 + 6 + 5) * 5
(3 + 9 * 5 + 7) + 2
((9 + 5 * 2) * 5 + (4 + 2 * 7) + 9) * 8
(4 + 4) * 9 + 3 + 2 * 8
(3 + 4 * 8 + 9 * 6 + 6) * 5 + 8 + 9
3 * 8 + 9 * (3 * 4)
6 + 7 * 6 * 7
3 * 3 + (8 * 8 * 9 + 3 * 8 + 2) * 3 + 8 * 2
5 + ((6 + 2 * 6) * 8 + 6 * 5) + ((7 + 3) + 2 * 5 + 7 + 6 + 5) + (8 + 2 * 7 + 4 + 2 * 3) * 2 * 5
(6 * (6 * 7) * 7 + 3) * 9 + 6 * 3
(5 * 9 + 7) + 8
(3 + 9 * 3 * 9 * 2) * 9 * 5
5 + (6 * 3)
`

var testInput = `
1 + 2 * 3 + 4 * 5 + 6
1 + (2 * 3) + (4 * (5 + 6))
2 * 3 + (4 * 5)
5 + (8 * 3 + 9 + 3 * 4 * 3)
5 * 9 * (7 * 3 * 3 + 9 * 3 + (8 + 6 * 4))
((2 + 4 * 9) * (6 + 9 * 8 + 6) + 6) + 2 + 4 * 2
`

func main() {
	lines := util.ReadInput(input, "\n")

	fmt.Printf("first %d\n", first(lines))
	fmt.Printf("second %d\n", second(lines))
}

func first(lines []string) int {
	total := 0
	for _, l := range lines {
		s := strings.ReplaceAll(l, " ", "")
		sum, _ := calc(s)
		total += sum
	}

	return total
}

func second(lines []string) int {
	total := 0
	for _, l := range lines {
		s := strings.ReplaceAll(l, " ", "")
		sum, _ := calcAdd(s)
		total += sum
	}

	return total
}

func calc(s string) (int, string) {
	first, rest := getNum(s, calc)
	if strings.HasPrefix(rest, ")") {
		return first, rest[1:]
	}

	operator, rest := getOp(rest)
	second, rest := getNum(rest, calc)

	sum := operator.Apply(first, second)
	if rest == "" {
		return sum, ""
	}

	return calc(strconv.Itoa(sum) + rest)
}

func calcAdd(s string) (int, string) {
	first, rest := getNum(s, calcAdd)
	if strings.HasPrefix(rest, ")") {
		return first, rest[1:]
	}

	if rest == "" {
		return first, ""
	}

	operator, rest := getOp(rest)
	second, rest := getNum(rest, calcAdd)

	var sum int
	switch operator.(type) {
	case Add:
		sum = operator.Apply(first, second)
	case Multiply:
		second, rest = calcAdd(strconv.Itoa(second) + rest)
		sum = operator.Apply(first, second)
		return sum, rest
	}

	if rest == "" {
		return sum, ""
	}

	return calcAdd(strconv.Itoa(sum) + rest)
}

var rgNum *regexp.Regexp = regexp.MustCompile("^(?P<Num>[0-9]+)(?P<Rest>.*)$")

func getNum(s string, calculate func(s string) (int, string)) (int, string) {
	var num int
	var rest string

	if strings.HasPrefix(s, "(") {
		return calculate(s[1:])
	}

	if rgNum.MatchString(s) {
		matches := rgNum.FindStringSubmatch(s)
		for i, name := range rgNum.SubexpNames() {
			switch name {
			case "Num":
				num, _ = strconv.Atoi(matches[i])
			case "Rest":
				rest = matches[i]
			}
		}
	}

	return num, rest
}

type Operator interface {
	Apply(int, int) int
}

type Multiply struct{}
type Add struct{}

func (o Multiply) Apply(a, b int) int {
	return a * b
}

func (o Add) Apply(a, b int) int {
	return a + b
}

var rgOp *regexp.Regexp = regexp.MustCompile("^(?P<Op>[+|*])(?P<Rest>.*)$")

func getOp(s string) (Operator, string) {
	var op Operator
	var rest string

	matches := rgOp.FindStringSubmatch(s)
	for i, name := range rgOp.SubexpNames() {
		switch name {
		case "Op":
			if matches[i] == "*" {
				op = Multiply{}
			} else if matches[i] == "+" {
				op = Add{}
			}
		case "Rest":
			rest = matches[i]
		}
	}

	return op, rest
}