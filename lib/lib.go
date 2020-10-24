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
		GoingWest(f *[][]byte) bool
		Walking(r *[]Step, f *[][]byte, d *[]string, i int, n *int, e *int, s *int, w *int) error
		FindPosibilities(f *[][]byte, d *[]string) ([]Step, error)
		FindPosibilitiesV1(f *[][]byte) ([]Step, error)
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
		West  int
	}
)

// GoingNorth is method to let Position walking through north untill reach the edge
func (p *Position) GoingNorth(f *[][]byte) bool {
	y := &p.Row
	incrementNorth := -1
	maxNorth := 0
	*y += incrementNorth
	return *y > maxNorth
}

// GoingEast is method to let Position walking through east untill reach the edge
func (p *Position) GoingEast(f *[][]byte) bool {
	x := &p.Col
	incrementEast := 1
	maxEast := len((*f)[0]) - 1
	*x += incrementEast
	return *x < maxEast
}

// GoingSouth is method to let Position walking through south untill reach the edge
func (p *Position) GoingSouth(f *[][]byte) bool {
	y := &p.Row
	incrementSouth := 1
	maxSouth := len(*f) - 1
	*y += incrementSouth
	return *y < maxSouth
}

// GoingWest is method to let Position walking through south untill reach the edge
func (p *Position) GoingWest(f *[][]byte) bool {
	x := &p.Col
	incrementWest := -1
	maxWest := 0
	*x += incrementWest
	return *x > maxWest
}

// Walking is method to let Position find the key from move in the given floor
func (p Position) Walking(r *[]Step, f *[][]byte, d *[]string, i int, n *int, e *int, s *int, w *int) error {
	var err error
	cp := p

	for {
		dc := (*d)[i]
		var going bool

		switch dc {
		case "north":
			going = cp.GoingNorth(f)
			*n = int(math.Abs(float64(p.Row) - float64(cp.Row)))
		case "east":
			going = cp.GoingEast(f)
			*e = int(math.Abs(float64(p.Col) - float64(cp.Col)))
		case "south":
			going = cp.GoingSouth(f)
			*s = int(math.Abs(float64(p.Row) - float64(cp.Row)))
		case "west":
			going = cp.GoingWest(f)
			*w = int(math.Abs(float64(p.Col) - float64(cp.Col)))
		default:
			err = errors.New("unknown direction clue")
			if err != nil {
				return err
			}
		}

		if going == false || (*f)[cp.Row][cp.Col] == 0 {
			break
		}

		if dc == (*d)[len(*d)-1] {
			*r = append(*r, Step{
				Position: Position{
					Row: cp.Row,
					Col: cp.Col,
				},
				North: *n,
				East:  *e,
				South: *s,
				West:  *w,
			})
		} else {
			np := cp
			ni := i + 1

			err = np.Walking(r, f, d, ni, n, e, s, w)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

// FindPosibilities is method to let Position find the key from anywhere in the given floor
func (p Position) FindPosibilities(f *[][]byte, d *[]string) ([]Step, error) {
	colSize := len((*f)[0])
	for _, col := range *f {
		if len(col) != colSize {
			err := errors.New("floor size not sync")
			return nil, err
		}
	}

	initalIndex := 0
	result := make([]Step, initalIndex)

	sn := 0 // step to north
	se := 0 // step to east
	ss := 0 // step to south
	sw := 0 // step to west

	if err := p.Walking(&result, f, d, initalIndex, &sn, &se, &ss, &sw); err != nil {
		return nil, err
	}

	return result, nil
}

// FindPosibilitiesV1 is method to let Position find the key from North -> East -> South move in the given floor
func (p Position) FindPosibilitiesV1(f *[][]byte) ([]Step, error) {
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
