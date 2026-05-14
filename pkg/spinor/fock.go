// Package spinor contains the finite Witt/Fock bookkeeping used by the matter
// bridge layer.
//
// The package deliberately implements the representation as a finite typed basis
// first: four creation modes generate 2^4 occupation states. Physical labels
// such as lepton/color seeds are metadata attached to the modes, not fitted
// constants.
package spinor

import (
	"fmt"
	"math"
)

type ModeKind string

const (
	TemporalMode ModeKind = "temporal"
	SpatialMode  ModeKind = "spatial"
)

type Mode struct {
	Index int
	Name  string
	Kind  ModeKind
}

type FockState struct {
	Occupation []bool
}

type FockSpace struct {
	Modes  []Mode
	States []FockState
}

func NewCovariantPhaseFockSpace(spacetimeDimension int) (FockSpace, error) {
	if spacetimeDimension <= 0 {
		return FockSpace{}, fmt.Errorf("spacetime dimension must be positive")
	}
	if spacetimeDimension != 4 {
		return FockSpace{}, fmt.Errorf("current matter bridge expects 4 covariant modes, got %d", spacetimeDimension)
	}

	modes := []Mode{
		{Index: 0, Name: "a†_0", Kind: TemporalMode},
		{Index: 1, Name: "a†_1", Kind: SpatialMode},
		{Index: 2, Name: "a†_2", Kind: SpatialMode},
		{Index: 3, Name: "a†_3", Kind: SpatialMode},
	}
	stateCount := 1 << len(modes)
	states := make([]FockState, 0, stateCount)
	for mask := 0; mask < stateCount; mask++ {
		occ := make([]bool, len(modes))
		for i := range modes {
			occ[i] = (mask & (1 << i)) != 0
		}
		states = append(states, FockState{Occupation: occ})
	}
	return FockSpace{Modes: modes, States: states}, nil
}

func (f FockSpace) ModeCount() int  { return len(f.Modes) }
func (f FockSpace) StateCount() int { return len(f.States) }

func (f FockSpace) ExpectedStateCount() int {
	if len(f.Modes) >= 63 {
		return 0
	}
	return 1 << len(f.Modes)
}

func (f FockSpace) Vacuum() (FockState, error) {
	for _, s := range f.States {
		if s.ExcitationNumber() == 0 {
			return s, nil
		}
	}
	return FockState{}, fmt.Errorf("Fock space contains no vacuum state")
}

func (f FockSpace) SpatialModeCount() int {
	count := 0
	for _, m := range f.Modes {
		if m.Kind == SpatialMode {
			count++
		}
	}
	return count
}

func (f FockSpace) TemporalModeCount() int {
	count := 0
	for _, m := range f.Modes {
		if m.Kind == TemporalMode {
			count++
		}
	}
	return count
}

func (s FockState) ExcitationNumber() int {
	count := 0
	for _, occupied := range s.Occupation {
		if occupied {
			count++
		}
	}
	return count
}

func (s FockState) BMinusL() float64 {
	if len(s.Occupation) != 4 {
		return math.NaN()
	}
	charge := 0.0
	if s.Occupation[0] {
		charge -= 1.0
	}
	for i := 1; i < 4; i++ {
		if s.Occupation[i] {
			charge += 1.0 / 3.0
		}
	}
	return charge
}

func (s FockState) IsVacuum() bool {
	return s.ExcitationNumber() == 0
}

func (s FockState) IsSterileVacuumCandidate(eps float64) bool {
	return s.IsVacuum() && math.Abs(s.BMinusL()) < eps
}

func (f FockSpace) OneParticleStates() []FockState {
	out := make([]FockState, 0, len(f.Modes))
	for _, s := range f.States {
		if s.ExcitationNumber() == 1 {
			out = append(out, s)
		}
	}
	return out
}

func (f FockSpace) ChargeSpectrum() map[string]int {
	buckets := map[string]int{}
	for _, s := range f.States {
		key := fmt.Sprintf("%.6f", s.BMinusL())
		buckets[key]++
	}
	return buckets
}
