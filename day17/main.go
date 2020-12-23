package main

import (
	"fmt"

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
	stop := util.WithProfiling()
	defer stop()

	rows := util.ReadInput(testInput, "\n")

	fmt.Printf("first %d\n", run(build3D(rows), 6))
	fmt.Printf("second %d\n", run(build4D(rows), 6))
}

type Point interface {
	Neighbors() []Point
	Hash() int
}

var _ Point = Point3D{}
var _ Point = Point4D{}

type Point3D struct {
	z int
	y int
	x int
}

type Point4D struct {
	w int
	z int
	y int
	x int
}

func build3D(rows []string) map[int]Point {
	points := make(map[int]Point)

	for i, row := range rows {
		for j, r := range row {
			if r == '#' {
				p := Point3D{z: 0, y: i, x: j}
				points[p.Hash()] = p
			}
		}
	}

	return points
}

func build4D(rows []string) map[int]Point {
	points := make(map[int]Point)

	for i, row := range rows {
		for j, r := range row {
			if r == '#' {
				p := Point4D{w: 0, z: 0, y: i, x: j}
				points[p.Hash()] = p
			}
		}
	}

	return points
}

func run(points map[int]Point, cycles int) int {
	for i := 0; i < cycles; i++ {
		nextPoints := make(map[int]Point, len(points))

		inactiveNeighbors := make(map[Point]int)
		for _, p := range points {
			active := 0
			for _, n := range p.Neighbors() {
				if _, ok := points[n.Hash()]; ok {
					active += 1
					continue
				}

				inactiveNeighbors[n] += 1
			}

			if active == 2 || active == 3 {
				nextPoints[p.Hash()] = p
			}
		}

		for p, active := range inactiveNeighbors {
			if active == 3 {
				nextPoints[p.Hash()] = p
			}
		}

		points = nextPoints
	}

	return len(points)
}

func (p Point3D) Neighbors() []Point {
	var n []Point

	for z := p.z - 1; z <= p.z+1; z++ {
		for y := p.y - 1; y <= p.y+1; y++ {
			for x := p.x - 1; x <= p.x+1; x++ {
				if x == p.x && y == p.y && z == p.z {
					continue
				}

				n = append(n, Point3D{z: z, y: y, x: x})
			}
		}
	}

	return n
}

func (p Point3D) Hash() int {
	m := 1
	hash := 0

	for _, i := range []int{p.z, p.y, p.x} {
		hash += m * i
		m *= 100
	}

	return hash
}

func (p Point4D) Neighbors() []Point {
	var n []Point

	for w := p.w - 1; w <= p.w+1; w++ {
		for z := p.z - 1; z <= p.z+1; z++ {
			for y := p.y - 1; y <= p.y+1; y++ {
				for x := p.x - 1; x <= p.x+1; x++ {
					if w == p.w && x == p.x && y == p.y && z == p.z {
						continue
					}

					n = append(n, Point4D{w: w, z: z, y: y, x: x})
				}
			}
		}
	}

	return n
}

func (p Point4D) Hash() int {
	m := 1
	hash := 0

	for _, i := range []int{p.w, p.z, p.y, p.x} {
		hash += m * i
		m *= 100
	}

	return hash

}
