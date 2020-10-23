package main

import (
	"fmt"

	"github.com/hrz8/joni-home-key/lib"
)

func main() {
	joniPosition := lib.NewPositionInterface(4, 1)
	posibleKeyPositions := joniPosition.FindKey(&[][]byte{
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 1, 1, 1, 1, 1, 1, 0},
		{0, 1, 0, 0, 0, 1, 1, 0},
		{0, 1, 1, 1, 0, 1, 0, 0},
		{0, 1, 0, 1, 1, 1, 1, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
	})

	fmt.Println("These are posible key position coordinates (x, y) with moves (north -> east -> south):")
	if len(posibleKeyPositions) > 0 {
		for i, keyPos := range posibleKeyPositions {
			fmt.Printf("%d. at (%d, %d) with move (%d -> %d -> %d)\n", i+1, keyPos.Position.Row, keyPos.Position.Col, keyPos.North, keyPos.East, keyPos.South)
		}
	} else {
		fmt.Println("0. no posible position!")
	}

}
