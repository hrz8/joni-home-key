package main

import (
	"fmt"
	"strings"

	"github.com/hrz8/joni-home-key/lib"
)

func main() {
	floor := &[][]byte{
		// #1 -> 0 is wall, unavailable route
		// #2 -> everything not 0 is available route

		/* col
		|0, 1, 2, 3, 4, 5, 6, 7| */
		{0, 0, 0, 0, 0, 0, 0, 0}, // row 0
		{0, 1, 1, 1, 1, 1, 1, 0}, // row 1
		{0, 1, 0, 0, 0, 1, 1, 0}, // row 2
		{0, 1, 1, 1, 0, 1, 0, 0}, // row 3
		{0, 1, 0, 1, 1, 1, 1, 0}, // row 4
		{0, 0, 0, 0, 0, 0, 0, 0}, // row 5
	}

	clue := &[]string{
		"north", // 1sr step
		"east",  // 2nd step
		"south", // 3rd step
	}

	joniPosition := lib.NewPositionInterface(4, 1)
	posibleKeyPositions, err := joniPosition.FindPosibilities(floor, clue)
	// posibleKeyPositions, err := joniPosition.FindPosibilitiesV1(floor)
	if err != nil {
		panic(fmt.Sprintf("[ERROR] %s", err.Error()))
	}

	fmt.Printf("these are posible key position coordinates (row, col) with moves %s:\n", strings.Join(*clue, "->"))
	if len(posibleKeyPositions) > 0 {
		for i, keyPos := range posibleKeyPositions {
			fmt.Printf("%d. at (%d, %d) with move (%d->%d->%d)\n", i+1, keyPos.Position.Row, keyPos.Position.Col, keyPos.North, keyPos.East, keyPos.South)
		}
	} else {
		fmt.Println("0. no posible position!")
	}
}
