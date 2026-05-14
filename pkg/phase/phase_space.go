package phase

import (
	"fmt"

	"github.com/bagherbal/asha-engine/pkg/clifford"
)

type CoordinateKind string

const (
	Position CoordinateKind = "POSITION"
	Momentum CoordinateKind = "MOMENTUM"
)

type Coordinate struct {
	Kind  CoordinateKind
	Index int
	Name  string
}

type CovariantPhaseSpace struct {
	SpaceTimeDimension int
	Coordinates        []Coordinate
}

func NewCovariantPhaseSpace(spaceTimeDimension int) (CovariantPhaseSpace, error) {
	if spaceTimeDimension <= 0 {
		return CovariantPhaseSpace{}, fmt.Errorf("space-time dimension must be positive")
	}

	coords := make([]Coordinate, 0, 2*spaceTimeDimension)
	for i := 0; i < spaceTimeDimension; i++ {
		coords = append(coords, Coordinate{Kind: Position, Index: i, Name: fmt.Sprintf("x%d", i)})
	}
	for i := 0; i < spaceTimeDimension; i++ {
		coords = append(coords, Coordinate{Kind: Momentum, Index: i, Name: fmt.Sprintf("p%d", i)})
	}

	return CovariantPhaseSpace{
		SpaceTimeDimension: spaceTimeDimension,
		Coordinates:        coords,
	}, nil
}

func (p CovariantPhaseSpace) Dimension() int {
	return len(p.Coordinates)
}

// CliffordSignature returns the current ASHA phase-space Clifford signature convention.
func (p CovariantPhaseSpace) CliffordSignature() clifford.Signature {
	return clifford.Signature{Positive: 1, Negative: p.Dimension() - 1}
}
