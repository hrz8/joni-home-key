package lib

import (
	"errors"
	"math"
)

type (
	// PositionInterface is represent methods of Position struct
	PositionInterface interface {
		GoingNorth(f *[][]byte) bool
		GoingEast(f *[][]byte) bool
		GoingSouth(f *[][]byte) bool
		FindKey(f *[][]byte) ([]Step, error)
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

// GoingNorth is method to let Position walking through north untill reach the edge
func (p *Position) GoingNorth(f *[][]byte) bool {
	x := &p.Row
	incrementNorth := -1
	maxNorth := 0
	*x += incrementNorth
	return *x > maxNorth
}

// GoingEast is method to let Position walking through east untill reach the edge
func (p *Position) GoingEast(f *[][]byte) bool {
	y := &p.Col
	incrementEast := 1
	maxEast := len((*f)[0]) - 1
	*y += incrementEast
	return *y < maxEast
}

// GoingSouth is method to let Position walking through south untill reach the edge
func (p *Position) GoingSouth(f *[][]byte) bool {
	x := &p.Row
	incrementSouth := 1
	maxSouth := len(*f) - 1
	*x += incrementSouth
	return *x < maxSouth
}

// FindKey is method to let Position find the key from North -> East -> South move in the given floor
func (p Position) FindKey(f *[][]byte) ([]Step, error) {
	// TODO: make sure each element of each floor row has the same size
	colSize := len((*f)[0])
	for _, col := range *f {
		if len(col) != colSize {
			err := errors.New("floor size not sync")
			return nil, err
		}
	}

	result := make([]Step, 0)

	pn := p
	for pn.GoingNorth(f) {
		if (*f)[pn.Row][pn.Col] == 0 {
			break
		}

		pe := pn
		for pe.GoingEast(f) {
			if (*f)[pe.Row][pe.Col] == 0 {
				break
			}

			ps := pe
			for ps.GoingSouth(f) {
				if (*f)[ps.Row][ps.Col] == 0 {
					break
				}
				result = append(result, Step{
					Position: Position{
						Row: ps.Row,
						Col: ps.Col,
					},
					North: int(math.Abs(float64(p.Row) - float64(pn.Row))),
					East:  int(math.Abs(float64(pn.Col) - float64(pe.Col))),
					South: int(math.Abs(float64(pe.Row) - float64(ps.Row))),
				})
			}
		}
	}

	return result, nil
}

// NewPositionInterface is creating new instance of Position struct
func NewPositionInterface(row int, col int) PositionInterface {
	return &Position{
		Row: row,
		Col: col,
	}
}
