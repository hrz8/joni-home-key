package main

import (
	"fmt"
	"math"
)

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
		Position
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

func (p Position) findKey(f *[][]byte) []Step {
	result := make([]Step, 0)

	posNorth := p
	for posNorth.goingNorth() {
		if (*f)[posNorth.Row][posNorth.Col] == 0 {
			break
		}

		posEast := posNorth
		for posEast.goingEast() {
			if (*f)[posEast.Row][posEast.Col] == 0 {
				break
			}

			posSouth := posEast
			for posSouth.goingSouth() {
				if (*f)[posSouth.Row][posSouth.Col] == 0 {
					break
				}
				result = append(result, Step{
					Position: Position{
						Row: posSouth.Row,
						Col: posSouth.Col,
					},
					North: int(math.Abs(float64(p.Row) - float64(posNorth.Row))),
					East:  int(math.Abs(float64(posNorth.Col) - float64(posEast.Col))),
					South: int(math.Abs(float64(posEast.Row) - float64(posSouth.Row))),
				})
			}
		}
	}

	return result
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

	posibleKeyPos := joniPos.findKey(&floor)

	fmt.Println("These are posible key position coordinates (x, y) with moves (north, east, south):")
	for i, keyPos := range posibleKeyPos {
		fmt.Printf("%d. at (%d, %d) with move (%d, %d, %d)\n", i+1, keyPos.Position.Row, keyPos.Position.Col, keyPos.North, keyPos.East, keyPos.South)
	}
}
