package main

import (
	"fmt"

	"github.com/mbark/advent-of-code-2020/util"
)

var input = `
L.LLLLLLLLLL.LLLLLLLLL.LLLLLLLLL.LLL.LL.LLLLLLLLL.LLLLLL.LLLL.LL..LLLL.LLLL.LLL.LL.LLLLLLL
.LLLL.L.LLLLLLL.LLLLLLLLLLLLLLL..LLLLLL.LLLLLLLLL.LLLLLL.LLLLLLL.LLLLLLLLLL.LLLLLLLLLLLLLL
LLLLLLLL.LLL.LLLL.LLLL.LLLLLLLLL.LLLLLLLLLLLLLLLL.LLLLLL.LLLLLLLLL.LLL.LLLL.LLLLLLLLLLLLLL
LLLLL..LL.LL.LLLLLLLLLLLLLLLLLLLLLLLLLL.LLLLLLLLL.LLLLLLLLLLLLLL.LLLLLLL.LLLLLLLLL.LLLLLLL
LLLLL.L.LLLLLLLL.LLLLL.LLLLLLLLLLLLLLLL.LLLLLLLLLLLLLLLL.LLLLLLL.LLLLL.LLLLLLLLLLLLLLLLLLL
LLLLL..L..L.L........L....L.L..LLL...L....L.....L...L..L.....LLL.L..L..LLL...L.LLL...L..L.
LLLLL.LLLLLL.LLLLLLLLLLLLLLLLLLLLL.LLLL.LLLLLL.LL.LLLLLL.LLLLLLL.LLLLL.LLLL.LLLLLLLLLLLLLL
LLLLL.LLLLLL.LLLLLLLLLLLLLLLLLL..LLLLLL.L.LLL.LLLLLLLLLLLLLLLLLL..LLLL.LLLL.LLLLLL.LLLLLLL
LLLLL.LLL.LLLLLLLLLLLL.LLLLLLLLLLLLLLLL.LLLLLLLLLLLLLLLL.LLLLLLL.LLLLLLL.LL.LLLLLLLLLLLLLL
LLLLLLLLLLLLLLLLLLLLLL.LLLLLLLL..LLLLLL.LLLLLLLLLL.LLLLL.LLLLLLLLLLLLL.LLLL.LLLLLL.L.LLLLL
LLLLLLLLLLLL.LLLLLLLLL.LLLLLLLLLLLLLLLL.LLL..LL.LL.LLLLLLLLLLLLL.LLLLL.LLLL.LLLLLL.LLLLLLL
LLLLL.LLLLL..LLLLLLLLL.LLLLLLLLLLLLLLLLLLLLLLL.LLLLLLLLL..LLLLLL.LLLLL.LLLL.LLLLLLLLLLLLLL
LLL.LLLLLLLL.LLLLLLLLLLLLLLLLLLL.LLLLLL.LLLLLL.LL.LLLLLL.LL.LLLLLLLLLLLLLLL.LLLLLL.LLLLLLL
.L.LL..LL.LLL...L....L..L....LLL...LL.L.L......L.....................LLLL.....L..LLL.L..L.
LLLLLL.LLLLL.LLLLL.LLL.LLLLLLLLL.LLLLLL.LLLLLL.LL.LLL.LL.LLLL.LLLLLLLL.LLLLLLLLLLL.LLLLLLL
LLLLLLLLLLLL..L.LLLLLL.LLLLLLLLL.LLLLLLLLLLLLLLLL.LLLLLL.LLLLLLLLLLLLLLLL.LLLLLLLL.LLLLLLL
LLLLLLLLLLLL.LLL.LLLLL.LLLLLLLLL.L.LLL..LLLLLLLLL.LLL.LL.LLLLLLL.LLLL..LLLL..LLLLLLLLLL.LL
LLLLL.LLLLLL.LL.LLL.LL.LLLLLL.LL.LLLLLL.LLLLLLLLL.LLLLLL.LLLLLL..LLLLL.LLLL.LLLLLLLLLLLLLL
..L...LLL..L.......L..LL.LLL........L..L..LLL..L.L.LLL..LL....LL.L.L...LLL...LLL..LL.L...L
LLLLL.LLLLLL.LLLLLLLLLLLLLLLLLLL.LLLLLLLLLLLLLLLLLLLLLLL.L.LLLLLLLLLLL.LLLL.LLLLLLLLLLLLLL
LLLLLLLLLLL..LLLL.LLLLLLLLLLLLLL.LLLLLLLLL.LLLLL..LLLLLL.LLLLLLL.LLLLL.LLLL.L.LLLLLLLLLLLL
LLLLL.LLL.LLLLLLLLLLLL.LLLLLL.L..LLLLL..LLLLLLLLL.LLLLLL.LLLLLLL.LLLLL.LLLLLLLLLLL.LLLLLLL
LLLLL.LLLLLLLLL.LLLLLL.LLLLLLLLLLL.LLLLLLLLLLLLLL.LLLLLL.LLL.LLL.LLLLL.LLLL.LLL.LL.LL.LLLL
LLLLL.LLLLLL.LLLLLLLLL.LLLLLLLLL.LLLLLLLLLLLLLLLL.L.LLLL.LLLLLLL.LLLLLLLLLLLLLLLLL.LLLLLLL
LLLLL.LLLLLL.LLLLLLLLL.LLLLLLLLL.LLLLLL.LLLLLLLLLLL.LLLL.LLLLLLL.LLLLL.LLLLLLLLLLLLLL.LLLL
LLLLL.LLLLLL.LLLLLLLL..LLL.LLLLL.LLLLLL.LLLLLLLLL.LLLLLL.LLLLLLL.LLLLL.LLLLLLLLLLL.LLLLLLL
LLLLL.LLLLLL..LLLLLLLLLLLLLLLLLLLL.LLLLLLLL.LLLLL.LLLLLL.LLLLLLLLLLLLLLLLLL.LLLLLL.LLLLLLL
..L..L......L.L...L.....L....L..L.L..L.LL.LLLLL....LL...L.L.LL....L..L..L..............LL.
LLLLLLLLLLL..LLLLLLLLL.LL.LLLLLL.LLLLLL.LLLLLLLLL.LLLLLL.LLLLLLLLLLLLLLLLLL.LLLLLLLLLLLLL.
LLLLLL.LLLLL.LLLLLLLLLLLLL.LLLLL.LLLLLL.LLLLLLLLL.LLLLLLLLLLLL.LLLLLLL.LLLLLLLLLLL.LLLL.LL
LLLLL.LLLLLLLLLLLLLLL..LLLLLLL.L.LLLLLL.LLLLLLLLLLLLLLLLLL.LLLLLL.LLLL..LLL.LLL.LL.LLLLLLL
LLLLL.LLLLLL.LLLLLLLLL.LLLLLLLLLLLLLLLLLLLLLLLLLL.LLLLLLLLLLLLLLLLLLLL.LLLL.LLLLLL.LLLLLLL
LLLLL.LL.LLL.LLLLLLLLL.LLLLLL.LLLLLLLLL.LLLLLLLLL.L.LLLLLL.LLLLLLLLLLLLLLLL.LLLLLL.LLLLLLL
L.......L...........L..L.....L.L.L.....L...L.....LL...........L..L.....L.....LL......L....
LLLLL.LLLLLL.LLLLLLLLL.LLLLLLLLLLLLLLLL.LLLLLLLLL.LLLLLLLLLLLLLLLLLLLL.LLLLLLLLLLL.LLLLLLL
LLLLL.LLLLLL.LLLLLLLLL.LLL.LLLLLLLLLLLL.LL.L.L.LLLLLLLLLL.L.LLLL.LLL.LLLLLLLLLLLLL.LLLLLLL
LLLLL.LLLLLL.LLLLLLLLL.LLLLLLLLL.LLLLLL.LLLLLLLLLLLLLLLL.LLLLLLL.LLLLL.LLLL.LLLLLL.LLLLLLL
LLLL..LLLLLL.LLLLLLLLLLLLLLLLLLL.LLLLLL.LLLLLLLL.LLLLLLL.LLLLLLL.LLLLL.LLLL.LLLLLLLLLLLLLL
LLLLL.LLLLLL.LLLLLLLLL.LLLLLLLLLLLLLLLL.LLLL.LLLLLLLLLLL.LLLLLLL.LLLLLLLLLL.LLLLLL.LLLLLLL
.L..L...L.......L....L.....L...........LLLLLL.L...L..L.L....L.LLL...LL......LL.L..L.LL..L.
LLLLL.LLLLLL.LLLLLLLLL.LLLLLLLLLLLLLLLL.LLLLLLLLLLLL.LLL.LLLLLLLLLLLLL.LLLL.LLLLLL.LLLLLLL
LLLLL.LLLL.L.LLLLLLLLL.LLLLLLLLLLLLLLLL.LLLLLL.LL.LLLLLL..LLLLLLLLLLLL.LLLLLLLLLLLLLLLLLLL
LLLLL.LLLLLLLLLLLLLLLLLLLLLLLLLL.LLLLLL.LLLLLL.LLLLLLLLLLLLLLLLLLLLLLL.L.LLL.LLLLLLL.LLLLL
LLLL.LL.LLLL.LLLLLLLLLLLLLLLLLLL.LLLLLL.LLLLLLLLLLLLL.LL.LLLLLLL.LLLLLLLLLL.LLLLLLL.LLLLLL
LLLLL.LLLLLL.LLLLLLLLLLLLLLLLLLL.LLLLLL.LLLLLLLLL.LLLLLL.LLLLLLLLLLL.LLL..L.LLLLLL.LLLLLLL
LLLLLLLLLLLLLLLLLLLL.L.LLLLLLLLLLLLLLLLLLLLLLLLLL.LLLLL..LLLLLLLLLLLLL.LLLL.LLLLLL.LLLLLLL
L.LLL.LLLLLL.LLLLLLLLL.L.LLLLLLL..LLLLL..LLLLLLLLLLLLLLL.LLLL.LL.LLLLL.LLLL.LLLLLL.LLLLLLL
LLLLL.L.LLLL.LLLLLLLLLLLL.LLLLLL.LLLLLL.LLLLLLLLL..LLLLLLLLLLLLL.LLLLL.LL.L.LLLLLL.LL.LLLL
LLLLL.LLLLLL.LLLLLLLLL.LLLLLLLLLLLLLLLL.LLLLLLLLL.LLLLLL.LLLLLLL.LLLLL.LLLL.LLLLLL.LL.LLLL
L..L.L.LL.LLL..LLL.L.L.......L...L.LL...L..L.L......L....L.....L..L....L....LL.L..........
LLLLLLLLLLLLLLLLLLLLLL.LLLLLLLLL.LLLLLL.LLLLLLLLL.LLLLLL.LLLLLLL.LLLLLLLLLL.LLLLLL.LLLLLLL
LLLLL.LLLLLL.LLLLLLLLLLLLLLLLLLL.LLLLLL.LLL.LLLLLLLLLLLLLLLLLLLLLLLLLLLLLLL.LLLL.L.LLLLLLL
LLLLLLLLLLLL.LL.LLLLLL.LLLLL.LLL.LLLLLLLLLL.LLLLLLLLLLLL.LLLL.LL.LLLLL.LLLL.LLLLLL.LLL.LLL
LLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLL.LLLLLL.LLLLLLLLL.LLL.LLLLLLLLLLLLLLL..LLLL.LLLLLLLLLLLLLL
LLLLL.LLLLLL.LLLLLLLLL.LLLLLLLLLLLLLLLLLLLLLLLLLL.LLLLLL.LLLLLLLLLLLLL.LLLL.LLLLLL..LLLLLL
LLLLLLLLLLLL.LLLLLLLLL.LLLLLLLLLLLLLLLL.LLLLLLLLL.L.LLLL.LLLLLL..LLLLL.LLLL.LLLLLLLLL.LLLL
.......L....L....L.L....L.LLLL......L.L............L....L..L..L..L.LLL..L.LLLL.LL..L....L.
LLLLL.LLLLLL.L.LLLLLLL.LLLLLLLLL.LLLLLL..LLLLLLLL.LLLLLL.LL.L.LLLLLLLL.LLLL.LLLLLL.LLLLLLL
LLLLL.LLLLLL..LLLLLLLLLLLLLLLLLL.LLLLLL.LLLLLLLLL.LLLLLL.LLLLLLLLLLLLL.LLL..LLLLLLLLLLLLLL
LLLLL.LLL.LL.LLLLLLLL.LLLLLLLLLL.LLL.LLLLLLLLLLLL.LLLLLLLLLLLLLL..LLLLLLLLLLLLLLLLLLLLLLLL
L.LLL.LLLLLL.LLLLLLLLL.LLLLLL.LL.LLLL.L.LLLLLLLLL.LLL.LLLLLLLLLL.LLLLL.LLL.LLLLLLLLLLLLLL.
LLLLLLL.LLLL.LLLLLLLLL.LLLLLLLLL.LLLLLL.LLLLLLLLL.LL.LLL.LLLLLLL.LLLL.LLLLLLLLLLLL.LLLLLLL
LLLLL.LLLLLLLLLLLLLLLL.LLLLLLLLL.LLLLLL.LLLLLLLLL.LLLLLL.LLLLLLLLLLLLL.LLL.LLLL.LLLLL.LLLL
LL...L..LL........L...LL....L.LL..L.........LL..L...L..L.L...L.L.LL..L.L....LL....LL.....L
LLLLL.LLLLLL.LLLLLLLL..LLLLLLLLL.LLLLLL.LLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLL.L.LLLLLLLLLLLL
LLLLL.LLLLLL.LLLLLLLLL.LLLLLLLLLLLLLLLL.LLLLLLLLLLLLLLLLLLLLLLLL.LLLLL.LLLLLLLLLLL.LLLL.LL
LLLLL.LLLLLL.LLLLLLLLL.LLLLL.LLL.LLLL.L.LLLLLLLLLLL.LLLL.LLLLLLL.L.LLL.LLLLLLLLLLL.LLLLLLL
LL.LL.LLLLLLLLL.LLLLLLLLLLLLLLLL.LLLLLL.LL.LLLLLL.LLLLLL.LLLLLLLLLLLLLLL.LL.LLLLLL.LLL.LLL
LLLLLLLLLLLL.LLLLLLLLL.LL.LLLLLL.LLLLLLLLLLLLLLLLLLLLLLLLLLLLLLL.LLLLLL.LLL.LLLLLL.LLLLLLL
LLL.L.L.LLLL.LLLLLLLLLLLLLLLLLLL.L.LLLLLLLLLLLLLL.LLLLLL.LLLLLLLLLLLLL..LLL.LLLLLL.LLL.LLL
L.LL..L.L.L...L......L...L.......L.L.LL.....L.........L....L.......L.LL....L..L.L...L....L
LLLLL.LLLLLLLLLLLLLLLL.LLLLLLLLL.LLLLLL.LLLLLLLLLLLLL..LLLLLLLLLLLLLL.LLLL.LLLLLLLLLLLLLLL
LLLLL..LLLLL.LLLLLLLLL.LLLLLLLLL.LLLLLL.LLLLLLLLLLLLLLLL.LLL.LLL.LLLLLLLLLL.LLLLLL.LLLLLLL
LLLLLLLLLLLL.LLLLLLLLL.LLLL.LLLL.LLLLLL.LLLLLLLLL.LLLL.L.LLLLLLL.L.LLL.LLLL.LLLLLL.LLL.LLL
LLLLLLLLLLLL.LLLL.LLLLLLLLLLLLL..LLLLLLLLLLLLLLLL.LLLLLL.LLLLLLL.LLLLL.LLLL.LLLLLL.LLLLLLL
LLLLL.LLLLL..LLLLLLLLLLLLLLLLLLL.LLLLLLL.LLLLLLLLLLLLLLL.L.LLLLLLLLLLLLLLLL.LLLLLL.LLLLLLL
LLLLL.LLLLLL.LLLLLL.LL.LLLLLLLLLLLLLLLL..LLLLLLLL.LLLLLLL.LL.LLLLLLLLL.LLLL.LLL.LLLLLLLLLL
LLLLL.LLLLLL.LLLLLLLLL.LLLLLLLLLLLLLLLL.LLLLL.LLL.LLLLLL.LLLLLL..LLLLL.LLLL.L.LLLL.LLLLLL.
LLLLL.LLLLLLLLLLLLL..L.LLLLLLLLL.LLLLLL.LLL.LLLLL.LLLLLL.LLLLLLL..LLLL.LLLLLLLLLLL.LLLLLLL
LLLLL.LLLLLL.LLLLLLLLL.LLLLLLLLL.LLLLLLLLLLLLLLLL.LLLLLL.LLL.LLL.LLLLLLLLLL.LLLLLLLLLLLLLL
.LL...L...........LL...LLL.L.L.L.L...L.L.....LLLLLL..........L.L.L.L.L...L.L.L.L........L.
LLLLLLLLLLLL.LLLLLLLLL..LLLLLLLL.L.LLLLLLLLLLLLLL.LLLLLL.LLLLLLLLLLLLLLLLLL.LLLLL..LLLLLLL
L.LLL.LLLLLL.LLLLLLLLL.LLL.LLLLL.LLLLLL.LLLLLLLLLLLLLLLL.LLLLLLLLLLLLL.LLLLLLLLLLL.LLLLLL.
LLL.LLLLLLLL.LL.LLLLLL.L.LLLLLLL.LLLLLLLLLL.LLLLLLLLLLLL.LLLLLLL.LLLLL.LLLL.LLLLLLLLLLLLLL
LLLLLLLLL.L.LLLLLLLLLL.L.LLLLLLL.LLLL.L.LLLLLLLLLLLLLLLLLLLLL.LL.LLLLLLLLLL.LLLL.LLLLLLLLL
LLLLL.LLLLLLL.LLLLLLL..LLLL.LLLL.LLLLLL.LLLLLLLLL.LLLLLL.LLLLLLL.LLLLL.LLLL.L.LLLLLLLLLLL.
LLLLL.LLLLLLLLLLLLL.LLL.LLLLLLLLLLLLLLLLLLLLLLLLL.LLL.LL.LLLLLLL.LLLLLLLLLL.LLLLLL.LLLLLLL
.LLLL.LLLLLL.LLLLLLLLLLLLLLL.LLL.LLLLLL.LLLLLLL.L.LLLLLLLLLLLLLL.LLLL..L.L..LLLLLL.LLLLLLL
LLLLL.LLLLLLLLLLLLLLLL.LLLLLLLLL.LLLLLL.LLLLLLL.L.LLLLLL.LLLLLLLLLLLLL.LLLL.LLLLLL.LLLLLLL
LLLLL.LLLLLL..LLLLLLLL.LLLL.LLL.LLLLLLL.LLLLLLLLL.LL.LLL.LLLLLLLLLLLLL.LLLLLLLLLLL.LLLLLL.
LLLLLLLLLLLLLLLLLLLLLL.LLLLLLLLL.LLLLLLLLLLLLLLLLLLLLLLL.LLLLLLL.LLLLL.LLLL.LLLLLL..LLLLLL
LLLLL.LLLLLL.LLLLLLLLL.LLLLLLLLL.LLLLLLLLLLLLLLLL..LLLLL.LLLLLL.LLLLLLLLLLLL.LLLLLLLLLLLLL
L.LLL.LLLLLL.LLLLLLLLL.LLLLLLLLL.L.LLLLLLLLLLLLLLLLLL.LL.LLLLLLLLLLLLLL.LL..LLLLLL.LLLLLLL
`

var testInput = `
L.LL.LL.LL
LLLLLLL.LL
L.L.L..L..
LLLL.LL.LL
L.LL.LL.LL
L.LLLLL.LL
..L.L.....
LLLLLLLLLL
L.LLLLLL.L
L.LLLLL.LL
`

// type Seats = map[int]*Seat
type Seats struct {
	Seats   []*Seat
	Indexes []int
}

type Seat struct {
	y     int
	x     int
	state State
}

type State int

const (
	Free State = iota
	Occupied
)

func main() {
	stop := util.WithProfiling()
	defer stop()

	lines := util.ReadInput(input, "\n")

	maxhash := Seat{y: len(lines), x: len(lines[0])}.Hash()

	var indexes []int
	fSeats := make([]*Seat, maxhash+1)
	sSeats := make([]*Seat, maxhash+1)
	for y, row := range lines {
		for x, l := range row {
			if l == 'L' {
				hash := Seat{y: y, x: x}.Hash()
				indexes = append(indexes, hash)
				fSeats[hash] = &Seat{y: y, x: x, state: Free}
				sSeats[hash] = &Seat{y: y, x: x, state: Free}
			}
		}
	}

	fmt.Printf("first %d\n", first(Seats{Seats: fSeats, Indexes: indexes}))
	fmt.Printf("second %d\n", second(Seats{Seats: sSeats, Indexes: indexes}))
}

func first(seats Seats) int {
	return emptySeats(seats, 4, adjacentLocal(seats))
}

func second(seats Seats) int {
	return emptySeats(seats, 5, adjacentView(seats))
}

func emptySeats(seats Seats, leaveIf int, adjacent map[int][]*Seat) int {
	for {
		updates := make(map[int]State, len(seats.Indexes))

		for _, hash := range seats.Indexes {
			s := seats.Seats[hash]
			adj := adjacent[hash]
			switch s.state {
			case Free:
				hasOccupied := false
				for _, a := range adj {
					if a.state == Occupied {
						hasOccupied = true
						break
					}
				}

				if !hasOccupied {
					updates[hash] = Occupied
				}
			case Occupied:
				occoupied := 0
				for _, a := range adj {
					if a.state == Occupied {
						occoupied += 1
					}
					if occoupied >= leaveIf {
						updates[hash] = Free
					}
				}
			}
		}

		if len(updates) == 0 {
			break
		}

		for hash, to := range updates {
			s := seats.Seats[hash]
			s.state = to
			seats.Seats[hash] = s
		}
	}

	occupied := 0
	for _, hash := range seats.Indexes {
		s := seats.Seats[hash]
		if s == nil {
			continue
		}

		if s.state == Occupied {
			occupied += 1
		}
	}

	return occupied
}

func printSeats(seats [][]rune) {
	for _, row := range seats {
		for _, r := range row {
			fmt.Printf("%c", r)
		}

		fmt.Println()
	}

	fmt.Println()
}

func adjacentLocal(seats Seats) map[int][]*Seat {
	adj := func(ri, ci int) []*Seat {
		var adjSeat []*Seat
		for i := ri - 1; i <= ri+1; i++ {
			if i < 0 {
				continue
			}

			for j := ci - 1; j <= ci+1; j++ {
				if i == ri && j == ci {
					continue
				}

				if j < 0 {
					continue
				}

				seat := seats.Seats[Seat{y: i, x: j}.Hash()]
				if seat != nil {
					adjSeat = append(adjSeat, seat)
				}
			}
		}

		return adjSeat
	}

	adjacent := make(map[int][]*Seat, len(seats.Indexes))
	for _, hash := range seats.Indexes {
		seat := seats.Seats[hash]
		adjacent[hash] = adj(seat.y, seat.x)
	}

	return adjacent
}

func adjacentView(seats Seats) map[int][]*Seat {
	maxx := 0
	maxy := 0
	for _, hash := range seats.Indexes {
		s := seats.Seats[hash]
		if s.y > maxy {
			maxy = s.y
		}
		if s.x > maxx {
			maxx = s.x
		}
	}

	adj := func(ri, ci int) []*Seat {
		var adjSeats []*Seat
		for i := -1; i <= 1; i++ {
			for j := -1; j <= 1; j++ {
				if i == 0 && j == 0 {
					continue
				}

				m := 0
				for {
					m += 1

					y := ri - i*m
					x := ci - j*m

					if y < 0 || y > maxy || x < 0 || x > maxx {
						break
					}

					seat := seats.Seats[Seat{y: y, x: x}.Hash()]
					if seat != nil {
						adjSeats = append(adjSeats, seat)
						break
					}
				}
			}
		}

		return adjSeats
	}

	adjacent := make(map[int][]*Seat, len(seats.Indexes))
	for _, hash := range seats.Indexes {
		seat := seats.Seats[hash]
		adjacent[hash] = adj(seat.y, seat.x)
	}

	return adjacent
}

func (seat Seat) Hash() int {
	return seat.y + 1000*seat.x
}
