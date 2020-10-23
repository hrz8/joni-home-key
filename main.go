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

	// starting position of Joni
	joniPos := Position{
		Row: 4,
		Col: 1,
	}

	fmt.Println(floor[joniPos.Row][joniPos.Col])
}
