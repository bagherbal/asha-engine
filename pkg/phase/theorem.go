package phase

import (
	"fmt"

	"github.com/bagherbal/asha-engine/pkg/theorem"
)

func CovariantPhaseSpaceTheorem(spaceTimeDimension int) theorem.Theorem {
	id := fmt.Sprintf("PHY-PHASE-COVARIANT-%dD", spaceTimeDimension)
	return theorem.Theorem{
		ID:     id,
		Name:   fmt.Sprintf("%dD covariant phase-space bookkeeping", spaceTimeDimension),
		Layer:  theorem.LayerAlgebra,
		Status: theorem.ExactFinite,
		Verify: func() theorem.Result {
			ps, err := NewCovariantPhaseSpace(spaceTimeDimension)
			if err != nil {
				return theorem.Result{
					ID: id, Name: "Covariant phase space", Layer: theorem.LayerAlgebra, Status: theorem.FailedRoute,
					Checks: []theorem.Check{{Name: "construct phase space", Passed: false, Detail: err.Error()}},
				}
			}

			sig := ps.CliffordSignature()
			checks := []theorem.Check{
				{
					Name:   "phase-space dimension",
					Passed: ps.Dimension() == 2*spaceTimeDimension,
					Detail: fmt.Sprintf("dim phase space = dim x + dim p = %d", ps.Dimension()),
				},
				{
					Name:   "coordinate split",
					Passed: len(ps.Coordinates) == 2*spaceTimeDimension,
					Detail: fmt.Sprintf("coordinates=%v", coordinateNames(ps.Coordinates)),
				},
				{
					Name:   "ASHA Cℓ(1,7) convention for 4D spacetime",
					Passed: spaceTimeDimension != 4 || (sig.Positive == 1 && sig.Negative == 7),
					Detail: fmt.Sprintf("signature=(%d,%d)", sig.Positive, sig.Negative),
				},
			}

			return theorem.Result{
				ID:     id,
				Name:   fmt.Sprintf("%dD covariant phase-space bookkeeping", spaceTimeDimension),
				Layer:  theorem.LayerAlgebra,
				Status: theorem.ExactFinite,
				Checks: checks,
				Notes: []string{
					"This package only builds the typed phase-space arena.",
					"Quantization, dynamics, and physical constants belong to later theorem gates.",
				},
			}
		},
	}
}

func coordinateNames(coords []Coordinate) []string {
	out := make([]string, len(coords))
	for i, c := range coords {
		out[i] = c.Name
	}
	return out
}
