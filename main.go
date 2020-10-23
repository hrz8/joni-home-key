package main

import "fmt"

type (
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

	findKey(joniPos, &floor)

	fmt.Println(joniPos)
}

func goingNorth(p *Position) bool {
	x := &p.Row
	incrementNorth := -1
	maxNorth := 0
	*x += incrementNorth
	return *x > maxNorth
}

func goingEast(p *Position) bool {
	y := &p.Col
	incrementEast := 1
	maxEast := 7
	*y += incrementEast
	return *y < maxEast
}

func goingSouth(p *Position) bool {
	x := &p.Row
	incrementSouth := 1
	maxSouth := 5
	*x += incrementSouth
	return *x < maxSouth
}

func findKey(p Position, f *[][]byte) {
	posNorth := p
	for goingNorth(&posNorth) {
		if (*f)[posNorth.Row][posNorth.Col] == 0 {
			break
		}
		fmt.Println(&posNorth)

		posEast := posNorth
		for goingEast(&posEast) {
			if (*f)[posEast.Row][posEast.Col] == 0 {
				break
			}
			fmt.Println(&posEast)

			posSouth := posEast
			for goingSouth(&posSouth) {
				if (*f)[posSouth.Row][posSouth.Col] == 0 {
					break
				}
				fmt.Println(&posSouth)
				fmt.Println("gaca!")
			}
		}
	}
}
