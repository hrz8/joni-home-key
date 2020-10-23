package main

import "fmt"

type (
	// PositionInterface is represent methods of Position struct
	PositionInterface interface {
		goingNorth() bool
		goingEast() bool
		goingSouth() bool
	}

	// Position is represent coordinate
	Position struct {
		Row int
		Col int
	}

	// Step is represent the direction step
	Step struct {
		North int
		East  int
		South int
	}
)

func (p *Position) goingNorth() bool {
	x := &p.Row
	incrementNorth := -1
	maxNorth := 0
	*x += incrementNorth
	return *x > maxNorth
}

func (p *Position) goingEast() bool {
	y := &p.Col
	incrementEast := 1
	maxEast := 7
	*y += incrementEast
	return *y < maxEast
}

func (p *Position) goingSouth() bool {
	x := &p.Row
	incrementSouth := 1
	maxSouth := 5
	*x += incrementSouth
	return *x < maxSouth
}

func (p Position) findKey(f *[][]byte) {
	posNorth := p
	for posNorth.goingNorth() {
		if (*f)[posNorth.Row][posNorth.Col] == 0 {
			break
		}
		fmt.Println(&posNorth)

		posEast := posNorth
		for posEast.goingEast() {
			if (*f)[posEast.Row][posEast.Col] == 0 {
				break
			}
			fmt.Println(&posEast)

			posSouth := posEast
			for posSouth.goingSouth() {
				if (*f)[posSouth.Row][posSouth.Col] == 0 {
					break
				}
				fmt.Println(&posSouth)
				fmt.Println("gaca!")
			}
		}
	}
}

func main() {
	floor := [][]byte{
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 1, 1, 1, 1, 1, 1, 0},
		{0, 1, 0, 0, 0, 1, 1, 0},
		{0, 1, 1, 1, 0, 1, 0, 0},
		{0, 1, 0, 1, 1, 1, 1, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
	}

	joniPos := Position{
		Row: 4,
		Col: 1,
	}

	joniPos.findKey(&floor)

	fmt.Println(joniPos)
}
