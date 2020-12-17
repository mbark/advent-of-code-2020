package main

import (
	"fmt"
	"math"
	"reflect"
	"sort"

	"github.com/mbark/advent-of-code-2020/util"
)

var input = `
.#.##..#
....#.##
##.###..
.#.#.###
#.#.....
.#..###.
.#####.#
#..####.
`

var testInput = `
.#.
..#
###
`

type Region = []string

func main() {
	rows := util.ReadInput(input, "\n")

	fmt.Printf("first %d\n", first(rows, 6))
	fmt.Printf("second %d\n", second(rows, 6))
}

type Row = map[int]rune
type Grid = map[int]Row
type Cube = map[int]Grid
type Hypercube = map[int]Cube

type Update struct {
	x  int
	y  int
	z  int
	to rune
}

type HyperUpdate struct {
	x  int
	y  int
	z  int
	w  int
	to rune
}

func first(rows []string, cycles int) int {
	cube := make(Cube)
	cube[0] = make(Grid)

	maxz := 0
	minz := 0
	maxy := len(rows)
	miny := 0
	maxx := len(rows[0])
	minx := 0

	for i, row := range rows {
		cube[0][i] = make(map[int]rune)
		for j, r := range row {
			cube[0][i][j] = rune(r)
		}
	}

	for i := 0; i < cycles; i++ {
		var updates []Update

		for z := minz - 1; z <= maxz+1; z++ {
			for y := miny - 1; y <= maxy+1; y++ {
				for x := minx - 1; x <= maxx+1; x++ {
					activeNeighbors := 0
					for _, nc := range neighbors(cube, x, y, z) {
						if nc == '#' {
							activeNeighbors += 1
						}
					}

					c := getAt(cube, z, y, x)
					if c == '#' {
						if activeNeighbors < 2 || activeNeighbors > 3 {
							updates = append(updates, Update{x: x, y: y, z: z, to: '.'})
						}
					} else if activeNeighbors == 3 {
						updates = append(updates, Update{x: x, y: y, z: z, to: '#'})
					} else {
						updates = append(updates, Update{x: x, y: y, z: z, to: c})
					}
				}
			}
		}

		for _, u := range updates {
			grid, ok := cube[u.z]
			if !ok {
				grid = make(Grid)
			}

			row, ok := grid[u.y]
			if !ok {
				row = make(Row)
			}

			row[u.x] = u.to
			grid[u.y] = row
			cube[u.z] = grid

			if u.z < minz {
				minz = u.z
			}
			if u.z > maxz {
				maxz = u.z
			}
			if u.y < miny {
				miny = u.y
			}
			if u.y > maxy {
				maxy = u.y
			}
			if u.x < minx {
				minx = u.x
			}
			if u.x > maxx {
				maxx = u.x
			}
		}
	}

	active := 0
	for _, grid := range cube {
		for _, row := range grid {
			for _, p := range row {
				if p == '#' {
					active += 1
				}
			}
		}
	}

	return active
}

func second(rows []string, cycles int) int {
	hypercube := make(Hypercube)
	hypercube[0] = make(Cube)
	hypercube[0][0] = make(Grid)

	maxw := 0
	minw := 0
	maxz := 0
	minz := 0
	maxy := len(rows)
	miny := 0
	maxx := len(rows[0])
	minx := 0

	for i, row := range rows {
		hypercube[0][0][i] = make(map[int]rune)
		for j, r := range row {
			hypercube[0][0][i][j] = rune(r)
		}
	}

	for i := 0; i < cycles; i++ {
		var updates []HyperUpdate

		for w := minw - 1; w <= maxw+1; w++ {
			for z := minz - 1; z <= maxz+1; z++ {
				for y := miny - 1; y <= maxy+1; y++ {
					for x := minx - 1; x <= maxx+1; x++ {
						activeNeighbors := 0
						for _, nc := range neighbors4d(hypercube, x, y, z, w) {
							if nc == '#' {
								activeNeighbors += 1
							}
						}

						c := getAt4d(hypercube, w, z, y, x)
						if c == '#' {
							if activeNeighbors < 2 || activeNeighbors > 3 {
								updates = append(updates, HyperUpdate{x: x, y: y, z: z, w: w, to: '.'})
							}
						} else if activeNeighbors == 3 {
							updates = append(updates, HyperUpdate{x: x, y: y, z: z, w: w, to: '#'})
						} else {
							updates = append(updates, HyperUpdate{x: x, y: y, z: z, w: w, to: c})
						}
					}
				}
			}
		}

		for _, u := range updates {
			cube, ok := hypercube[u.w]
			if !ok {
				cube = make(Cube)
			}

			grid, ok := cube[u.z]
			if !ok {
				grid = make(Grid)
			}

			row, ok := grid[u.y]
			if !ok {
				row = make(Row)
			}

			row[u.x] = u.to
			grid[u.y] = row
			cube[u.z] = grid
			hypercube[u.w] = cube

			if u.w < minw {
				minw = u.w
			}
			if u.w > maxw {
				maxw = u.w
			}
			if u.z < minz {
				minz = u.z
			}
			if u.z > maxz {
				maxz = u.z
			}
			if u.y < miny {
				miny = u.y
			}
			if u.y > maxy {
				maxy = u.y
			}
			if u.x < minx {
				minx = u.x
			}
			if u.x > maxx {
				maxx = u.x
			}
		}
	}

	active := 0
	for _, cube := range hypercube {
		for _, grid := range cube {
			for _, row := range grid {
				for _, p := range row {
					if p == '#' {
						active += 1
					}
				}
			}
		}
	}

	return active
}

func printCube(cube Cube) {
	var zs []int
	for k := range cube {
		zs = append(zs, k)
	}
	sort.Ints(zs)

	for _, z := range zs {
		fmt.Printf("z = %d\n", z)

		grid := cube[z]
		var ys []int
		for k := range grid {
			ys = append(ys, k)
		}
		sort.Ints(ys)
		for _, y := range ys {
			row := grid[y]
			var xs []int
			for k := range row {
				xs = append(xs, k)
			}
			sort.Ints(xs)

			for _, x := range xs {
				fmt.Printf("%s", string(row[x]))
			}
			fmt.Println()
		}
	}

	fmt.Println("===============")
}

func rangeStartEnd(values []reflect.Value) (int, int) {
	var max int64 = math.MinInt64
	var min int64 = math.MaxInt64

	for _, v := range values {
		i := v.Int()
		if i > max {
			max = i
		}
		if i < min {
			min = i
		}
	}

	return int(min), int(max)
}

func getAt(cube Cube, z, y, x int) rune {
	empty := '.'
	grid, ok := cube[z]
	if !ok {
		return empty
	}

	row, ok := grid[y]
	if !ok {
		return empty
	}

	pos, ok := row[x]
	if !ok {
		return empty
	}

	return pos
}

func getAt4d(hcube Hypercube, w, z, y, x int) rune {
	empty := '.'
	cube, ok := hcube[w]
	if !ok {
		return empty
	}

	grid, ok := cube[z]
	if !ok {
		return empty
	}

	row, ok := grid[y]
	if !ok {
		return empty
	}

	pos, ok := row[x]
	if !ok {
		return empty
	}

	return pos
}

func neighbors(cube Cube, atx, aty, atz int) []rune {
	var positions []rune

	for z := atz - 1; z <= atz+1; z++ {
		for y := aty - 1; y <= aty+1; y++ {
			for x := atx - 1; x <= atx+1; x++ {
				if x == atx && y == aty && z == atz {
					continue
				}

				positions = append(positions, getAt(cube, z, y, x))
			}
		}
	}

	return positions
}

func neighbors4d(cube Hypercube, atx, aty, atz, atw int) []rune {
	var positions []rune

	for w := atw - 1; w <= atw+1; w++ {
		for z := atz - 1; z <= atz+1; z++ {
			for y := aty - 1; y <= aty+1; y++ {
				for x := atx - 1; x <= atx+1; x++ {
					if w == atw && x == atx && y == aty && z == atz {
						continue
					}

					positions = append(positions, getAt4d(cube, w, z, y, x))
				}
			}
		}
	}

	return positions
}
