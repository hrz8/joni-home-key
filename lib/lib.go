package lib

import "math"

type (
	// PositionInterface is represent methods of Position struct
	PositionInterface interface {
		GoingNorth() bool
		GoingEast() bool
		GoingSouth() bool
		FindKey(f *[][]byte) []Step
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
func (p *Position) GoingNorth() bool {
	x := &p.Row
	incrementNorth := -1
	maxNorth := 0
	*x += incrementNorth
	return *x > maxNorth
}

// GoingEast is method to let Position walking through east untill reach the edge
func (p *Position) GoingEast() bool {
	y := &p.Col
	incrementEast := 1
	maxEast := 7
	*y += incrementEast
	return *y < maxEast
}

// GoingSouth is method to let Position walking through south untill reach the edge
func (p *Position) GoingSouth() bool {
	x := &p.Row
	incrementSouth := 1
	maxSouth := 5
	*x += incrementSouth
	return *x < maxSouth
}

// FindKey is method to let Position find the key with North, East, South move in the given floor
func (p Position) FindKey(f *[][]byte) []Step {
	result := make([]Step, 0)

	pn := p
	for pn.GoingNorth() {
		if (*f)[pn.Row][pn.Col] == 0 {
			break
		}

		pe := pn
		for pe.GoingEast() {
			if (*f)[pe.Row][pe.Col] == 0 {
				break
			}

			ps := pe
			for ps.GoingSouth() {
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

	return result
}

// NewPositionInterface is creating new instance of Position struct
func NewPositionInterface(row int, col int) PositionInterface {
	return &Position{
		Row: row,
		Col: col,
	}
}
