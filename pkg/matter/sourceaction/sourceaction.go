// Package sourceaction formulates the variational problem for the missing
// active-to-generation source tensor.
//
// Gate 34 proved that the abstract source-tensor space exists,
//
//	M : H_active -> H_generation,    M in Hom(R^4,R^3),
//
// but the already-computed finite connection, BF curvature, and BF source
// contractions do not select a nonzero tensor. This package turns that
// diagnostic into an action principle.
//
// The minimal stable source action is
//
//	S[M] = 1/2 ||M||_F^2 - <J,M>,
//
// where J is the finite source term selected by geometry. Its stationary
// equation is M = J. Therefore, when the selected geometric source J is zero,
// the unique stable stationary tensor is M=0. This is a useful theorem: it says
// that a nonzero source tensor cannot be obtained by merely declaring the
// Hom(R^4,R^3) space available. A nonzero M requires a genuine nonzero source,
// a derived constraint, or a new finite interaction term.
package sourceaction

import (
	"fmt"
	"sync"

	"github.com/bagherbal/asha-engine/pkg/matter/sourcemap"
)

type ActionTerm string

const (
	TermQuadraticStiffness ActionTerm = "1/2 ||M||_F^2"
	TermLinearSource       ActionTerm = "- <J,M>"
	TermQuarticConstraint  ActionTerm = "quartic / fixed-norm constraint"
	TermTachyonicMass      ActionTerm = "negative quadratic instability"
)

type SourceAction struct {
	Name               string
	DomainDimension    int
	CodomainDimension  int
	TensorDimension    int
	QuadraticStiffness float64
	SourceNorm         float64
	StationaryNorm     float64
	MinimumAction      float64
	HessianPositive    bool
	UniqueMinimum      bool
	SelectedNonzeroMap bool
	Canonical          bool
	Derived            bool
	Detail             string
}

type Analysis struct {
	SourceMap sourcemap.Analysis

	GenerationDimension int
	ActiveDimension     int
	TensorDimension     int

	NaturalSourceNorm       float64
	NaturalStationaryNorm   float64
	NaturalMinimumAction    float64
	NaturalHessianPositive  bool
	NaturalUniqueMinimum    bool
	NaturalSelectsZero      bool
	NonzeroStationaryFound  bool
	ArbitraryMapRejected    bool
	QuarticConstraintNeeded bool
	TachyonicTermDerived    bool

	Actions []SourceAction
	Best    SourceAction

	TruthStatement    string
	RemainingUnknowns []string
}

var (
	defaultOnce  sync.Once
	defaultValue Analysis
	defaultErr   error
)

func BuildDefault() (Analysis, error) {
	defaultOnce.Do(func() {
		sm, err := sourcemap.BuildDefault()
		if err != nil {
			defaultErr = err
			return
		}
		defaultValue, defaultErr = Build(sm, 1e-8)
	})
	return defaultValue, defaultErr
}

func Build(sm sourcemap.Analysis, eps float64) (Analysis, error) {
	if eps <= 0 {
		eps = 1e-8
	}
	if sm.GenerationDimension <= 0 || sm.ActiveDimension <= 0 {
		return Analysis{}, fmt.Errorf("invalid source tensor dimensions %dx%d", sm.GenerationDimension, sm.ActiveDimension)
	}
	if sm.MapSpaceDimension != sm.GenerationDimension*sm.ActiveDimension {
		return Analysis{}, fmt.Errorf("source map dimension mismatch: got %d, expected %d", sm.MapSpaceDimension, sm.GenerationDimension*sm.ActiveDimension)
	}

	// In the current finite data every canonical active-to-generation source
	// candidate is zero. The best canonical force J is therefore zero. We keep
	// this explicit rather than silently constructing an arbitrary tensor.
	naturalSource := maxCanonicalSourceNorm(sm)
	stationary := naturalSource // For S=1/2||M||^2-<J,M>, stationarity gives M=J.
	minAction := -0.5 * naturalSource * naturalSource
	selectsZero := stationary <= eps

	natural := SourceAction{
		Name:               "minimal positive source action",
		DomainDimension:    sm.ActiveDimension,
		CodomainDimension:  sm.GenerationDimension,
		TensorDimension:    sm.MapSpaceDimension,
		QuadraticStiffness: 1,
		SourceNorm:         naturalSource,
		StationaryNorm:     stationary,
		MinimumAction:      minAction,
		HessianPositive:    true,
		UniqueMinimum:      true,
		SelectedNonzeroMap: stationary > eps,
		Canonical:          true,
		Derived:            true,
		Detail:             "Uses only the canonical geometric source candidates from Gate 34. Because they vanish, the stable stationary tensor is M=0.",
	}
	arbitrary := SourceAction{
		Name:               "arbitrary fitted source tensor",
		DomainDimension:    sm.ActiveDimension,
		CodomainDimension:  sm.GenerationDimension,
		TensorDimension:    sm.MapSpaceDimension,
		QuadraticStiffness: 1,
		SourceNorm:         1,
		StationaryNorm:     1,
		MinimumAction:      -0.5,
		HessianPositive:    true,
		UniqueMinimum:      true,
		SelectedNonzeroMap: true,
		Canonical:          false,
		Derived:            false,
		Detail:             "A nonzero source could be chosen by hand in the 12D map space, but that would be fitting rather than derivation.",
	}
	quartic := SourceAction{
		Name:               "symmetry-breaking fixed-norm action",
		DomainDimension:    sm.ActiveDimension,
		CodomainDimension:  sm.GenerationDimension,
		TensorDimension:    sm.MapSpaceDimension,
		QuadraticStiffness: -1,
		SourceNorm:         0,
		StationaryNorm:     0,
		MinimumAction:      0,
		HessianPositive:    false,
		UniqueMinimum:      false,
		SelectedNonzeroMap: false,
		Canonical:          false,
		Derived:            false,
		Detail:             "A Mexican-hat or fixed-norm term could force nonzero M, but no finite theorem has derived its sign, radius, or orientation.",
	}

	actions := []SourceAction{natural, arbitrary, quartic}

	return Analysis{
		SourceMap:               sm,
		GenerationDimension:     sm.GenerationDimension,
		ActiveDimension:         sm.ActiveDimension,
		TensorDimension:         sm.MapSpaceDimension,
		NaturalSourceNorm:       naturalSource,
		NaturalStationaryNorm:   stationary,
		NaturalMinimumAction:    minAction,
		NaturalHessianPositive:  true,
		NaturalUniqueMinimum:    true,
		NaturalSelectsZero:      selectsZero,
		NonzeroStationaryFound:  natural.SelectedNonzeroMap,
		ArbitraryMapRejected:    sm.ArbitraryMapsExist && !sm.CanonicalSourceTensorFound,
		QuarticConstraintNeeded: true,
		TachyonicTermDerived:    false,
		Actions:                 actions,
		Best:                    natural,
		TruthStatement:          truth(selectsZero, naturalSource, sm.MapSpaceDimension),
		RemainingUnknowns: []string{
			"U-17E-SOURCE-TENSOR-ACTION: derive a nonzero geometric source J or a finite symmetry-breaking action for M:H_active→H_generation",
			"U-17F-SOURCE-TENSOR-STABILITY: prove the Hessian and constraints select a unique nonzero M rather than an arbitrary point in Hom(R4,R3)",
			"U-16C-NONCOMMUTING-TEXTURES: obtain at least two non-commuting 3x3 texture operators before CKM/PMNS claims",
			"U-16D-YUKAWA-SCALE-BRIDGE: normalize any future source-action spectrum into physical Yukawa strengths without observed-mass fitting",
		},
	}, nil
}

func maxCanonicalSourceNorm(sm sourcemap.Analysis) float64 {
	max := 0.0
	for _, c := range sm.Candidates {
		if !c.Canonical {
			continue
		}
		if c.Kind == sourcemap.CandidateHiggsAnisotropy {
			continue // diagonal spurion, not an active-to-generation force J.
		}
		if c.Norm > max {
			max = c.Norm
		}
	}
	return max
}

func truth(selectsZero bool, sourceNorm float64, dim int) string {
	if selectsZero {
		return fmt.Sprintf("The minimal positive variational principle over Hom(H_active,H_generation) has zero geometric source, so it uniquely selects M=0. A nonzero texture source still requires a derived finite interaction; the %d-dimensional tensor space alone is not a derivation.", dim)
	}
	return fmt.Sprintf("A nonzero geometric source with norm %.6e appears in the source-action equation M=J. It is a candidate bridge, pending symmetry and normalization tests.", sourceNorm)
}

func FormatAction(a SourceAction) string {
	flags := "derived"
	if !a.Derived {
		flags = "not-derived"
	}
	canon := "canonical"
	if !a.Canonical {
		canon = "noncanonical"
	}
	return fmt.Sprintf("%s: %dx%d sourceNorm=%.6e stationaryNorm=%.6e minS=%.6e [%s,%s]", a.Name, a.CodomainDimension, a.DomainDimension, a.SourceNorm, a.StationaryNorm, a.MinimumAction, canon, flags)
}
