// Package scalarpotential derives the finite scalar/Higgs effective-potential
// normal form from the active 4D contact sector.
//
// The package is intentionally strict about epistemic status. The finite engine
// has derived a positive vacuum-mixing operator M_K with a four-dimensional
// active spectrum and three protected contact directions. From those finite
// invariants one can build a dimensionless shifted normal form for a scalar
// potential, but one cannot yet infer the observed electroweak vev, the Higgs
// boson mass, or a physical tachyonic mass term in GeV units.
//
// This gate therefore separates three things:
//
//  1. exact finite scalar data: rank, spectrum, trace, quartic trace;
//  2. bridge-level normal form: V(r)=lambda_shape(r^2-r0^2)^2;
//  3. forbidden claims: electroweak scale, Higgs mass, gauge eating theorem.
package scalarpotential

import (
	"fmt"
	"math"
	"sync"

	"github.com/bagherbal/asha-engine/pkg/dynamics/higgspotential"
)

type PotentialForm string

const (
	ShiftedRadialNormalForm PotentialForm = "V(r)=lambda_shape*(r^2-r0^2)^2"
)

type Analysis struct {
	Input higgspotential.Analysis

	ActiveRealDimension     int
	ComplexDoubletDimension int
	ProtectedDirectionCount int
	PotentialForm           PotentialForm
	ActiveSpectrum          []float64
	PairDegenerate          bool
	PairDegeneracyResidual  float64
	HighPairEigenvalue      float64
	LowPairEigenvalue       float64
	PairSplitting           float64

	VacuumRadiusSquared       float64
	VacuumRadius              float64
	QuarticTrace              float64
	LambdaShape               float64
	NormalFormQuadraticCoeff  float64
	NormalFormConstant        float64
	DimensionlessRadialMassSq float64

	FiniteTachyonicMassDerived bool
	ShiftedNormalFormAvailable bool
	GoldstoneCountResonance    bool
	GaugeEatingTheoremDerived  bool
	ElectroweakScaleDerived    bool
	HiggsMassDerived           bool

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
		hp, err := higgspotential.BuildDefault()
		if err != nil {
			defaultErr = err
			return
		}
		defaultValue, defaultErr = Build(hp, 1e-8)
	})
	return defaultValue, defaultErr
}

func Build(hp higgspotential.Analysis, eps float64) (Analysis, error) {
	if eps <= 0 {
		eps = 1e-8
	}
	if hp.ActiveContactDimension != 4 {
		return Analysis{}, fmt.Errorf("scalar potential expects four active real directions, got %d", hp.ActiveContactDimension)
	}
	if hp.ProtectedContactDimension != 3 {
		return Analysis{}, fmt.Errorf("scalar potential expects three protected contact directions, got %d", hp.ProtectedContactDimension)
	}
	if len(hp.ActiveContactSpectrum) != 4 {
		return Analysis{}, fmt.Errorf("scalar potential expects four active eigenvalues, got %d", len(hp.ActiveContactSpectrum))
	}
	if hp.OrderParameterNormSquared <= eps {
		return Analysis{}, fmt.Errorf("scalar potential requires a positive finite order parameter trace")
	}
	if hp.QuarticTrace <= eps {
		return Analysis{}, fmt.Errorf("scalar potential requires a positive quartic trace")
	}

	active := append([]float64(nil), hp.ActiveContactSpectrum...)
	high, low, split := pairSummary(active)
	radiusSq := hp.OrderParameterNormSquared
	radius := math.Sqrt(radiusSq)
	lambdaShape := hp.NormalizedQuarticShape
	if lambdaShape <= eps {
		lambdaShape = hp.QuarticTrace / (radiusSq * radiusSq)
	}

	// The shifted normal form is a bridge-level parametrization of the already
	// nonzero finite scalar radius. Expanding V=lambda(r^2-r0^2)^2 gives a
	// negative quadratic coefficient, but that sign is not an independently
	// derived tachyonic mass in the exact finite action.
	quadraticCoeff := -2 * lambdaShape * radiusSq
	constant := lambdaShape * radiusSq * radiusSq
	radialMassSq := 8 * lambdaShape * radiusSq

	goldstoneResonance := hp.ActiveContactDimension == 4 && hp.ProtectedContactDimension == 3
	shiftedAvailable := hp.MexicanHatKinematics && lambdaShape > eps && radiusSq > eps

	return Analysis{
		Input:                      hp,
		ActiveRealDimension:        hp.ActiveContactDimension,
		ComplexDoubletDimension:    hp.ActiveContactDimension / 2,
		ProtectedDirectionCount:    hp.ProtectedContactDimension,
		PotentialForm:              ShiftedRadialNormalForm,
		ActiveSpectrum:             active,
		PairDegenerate:             hp.PairDegenerateSpectrum,
		PairDegeneracyResidual:     hp.PairDegeneracyResidual,
		HighPairEigenvalue:         high,
		LowPairEigenvalue:          low,
		PairSplitting:              split,
		VacuumRadiusSquared:        radiusSq,
		VacuumRadius:               radius,
		QuarticTrace:               hp.QuarticTrace,
		LambdaShape:                lambdaShape,
		NormalFormQuadraticCoeff:   quadraticCoeff,
		NormalFormConstant:         constant,
		DimensionlessRadialMassSq:  radialMassSq,
		FiniteTachyonicMassDerived: false,
		ShiftedNormalFormAvailable: shiftedAvailable,
		GoldstoneCountResonance:    goldstoneResonance,
		GaugeEatingTheoremDerived:  false,
		ElectroweakScaleDerived:    false,
		HiggsMassDerived:           false,
		TruthStatement:             truth(shiftedAvailable, goldstoneResonance),
		RemainingUnknowns: []string{
			"U-18A-SCALAR-SCALE-BRIDGE: derive the dimensional conversion from finite radius r0^2 to the electroweak vev without inserting v=246 GeV",
			"U-18B-HIGGS-MASS-BRIDGE: derive a physical radial curvature normalization before comparing to the observed Higgs mass",
			"U-18C-GAUGE-EATING-THEOREM: show that the three protected contact directions are precisely Goldstone/gauge directions, not merely a 3-count resonance",
			"U-18D-SCALAR-ACTION-SIGN: derive the negative quadratic sign from a finite action, rather than only from a shifted normal-form parametrization",
		},
	}, nil
}

func pairSummary(values []float64) (high float64, low float64, split float64) {
	if len(values) == 0 {
		return 0, 0, 0
	}
	high = values[0]
	low = values[0]
	for _, v := range values {
		if v > high {
			high = v
		}
		if v < low {
			low = v
		}
	}
	return high, low, high - low
}

func truth(shifted, goldstone bool) string {
	if shifted && goldstone {
		return "The finite scalar sector derives a four-real-dimensional active contact doublet, a positive finite radius, a quartic shape invariant, and a three-direction protected resonance. This supports a dimensionless scalar-potential normal form, but not yet the electroweak vev, Higgs mass, or a gauge-eating theorem."
	}
	return "The finite scalar sector does not yet supply enough kinematic data for a scalar-potential normal form."
}
