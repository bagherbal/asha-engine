// Package goldstone audits the finite Goldstone / gauge-eating correspondence.
//
// The finite engine has already derived three separate pieces:
//
//  1. a four-real-dimensional scalar/contact active sector;
//  2. a shifted scalar normal form with a nonzero finite radius;
//  3. a finite SU(2)_L doublet representation and hypercharge identity.
//
// This package tests the next bridge question: do the three protected contact
// directions canonically become the three Goldstone directions eaten by the
// broken electroweak generators?
//
// The answer at this gate is deliberately conservative.  The counts match, and
// that is highly structured.  But a true gauge-eating theorem also needs a
// canonical map from protected contact directions into the broken generator
// directions, plus a kinetic/covariant-derivative normalization.  Those are not
// yet derived by earlier gates.
package goldstone

import (
	"fmt"
	"sync"

	"github.com/bagherbal/asha-engine/pkg/dynamics/scalarpotential"
	"github.com/bagherbal/asha-engine/pkg/matter/su2lgauge"
)

type Classification string

const (
	CountLevelResonance Classification = "count-level-resonance"
	CanonicalTheorem    Classification = "canonical-gauge-eating-theorem"
)

type Analysis struct {
	Scalar scalarpotential.Analysis
	Gauge  su2lgauge.Analysis

	ActiveRealDirections       int
	RadialDirections           int
	ScalarAngularDirections    int
	ProtectedContactDirections int

	GaugeGroupDimension       int
	UnbrokenEMDimension       int
	BrokenGaugeDirections     int
	SU2LGeneratorCount        int
	HyperchargeGeneratorCount int

	ScalarGoldstoneCountMatchesProtected bool
	BrokenGaugeCountMatchesGoldstone     bool
	GoldstoneCountResonance              bool

	CanonicalProtectedToBrokenMapDerived bool
	SU2LActionOnContactScalarDerived     bool
	CovariantDerivativeDerived           bool
	GaugeBosonMassMatrixDerived          bool
	GaugeEatingTheoremDerived            bool

	Classification Classification
	TruthStatement string
	MissingData    []string
}

var (
	defaultOnce  sync.Once
	defaultValue Analysis
	defaultErr   error
)

func BuildDefault() (Analysis, error) {
	defaultOnce.Do(func() {
		s, err := scalarpotential.BuildDefault()
		if err != nil {
			defaultErr = err
			return
		}
		g, err := su2lgauge.BuildDefault()
		if err != nil {
			defaultErr = err
			return
		}
		defaultValue, defaultErr = Build(s, g)
	})
	return defaultValue, defaultErr
}

func Build(s scalarpotential.Analysis, g su2lgauge.Analysis) (Analysis, error) {
	if s.ActiveRealDimension != 4 {
		return Analysis{}, fmt.Errorf("Goldstone audit expects a four-real-dimensional scalar sector, got %d", s.ActiveRealDimension)
	}
	if !s.ShiftedNormalFormAvailable || s.VacuumRadiusSquared <= 0 {
		return Analysis{}, fmt.Errorf("Goldstone audit requires a nonzero shifted scalar normal form")
	}
	if g.Dimension != 8 || !g.NonabelianSU2LGeneratorsDerived {
		return Analysis{}, fmt.Errorf("Goldstone audit requires the finite SU(2)_L doublet generator audit")
	}

	active := s.ActiveRealDimension
	radial := 1
	angular := active - radial
	protected := s.ProtectedDirectionCount

	// Electroweak count at the representation level:
	// dim(SU(2)_L) + dim(U(1)_Y) - dim(U(1)_em) = 3 + 1 - 1 = 3.
	// These are group dimensions, not observed constants.
	su2 := 3
	u1y := 1
	em := 1
	gaugeDim := su2 + u1y
	broken := gaugeDim - em

	goldstoneProtected := angular == protected
	brokenGoldstone := broken == angular
	resonance := goldstoneProtected && brokenGoldstone

	class := CountLevelResonance
	if false {
		class = CanonicalTheorem
	}

	return Analysis{
		Scalar:                               s,
		Gauge:                                g,
		ActiveRealDirections:                 active,
		RadialDirections:                     radial,
		ScalarAngularDirections:              angular,
		ProtectedContactDirections:           protected,
		GaugeGroupDimension:                  gaugeDim,
		UnbrokenEMDimension:                  em,
		BrokenGaugeDirections:                broken,
		SU2LGeneratorCount:                   su2,
		HyperchargeGeneratorCount:            u1y,
		ScalarGoldstoneCountMatchesProtected: goldstoneProtected,
		BrokenGaugeCountMatchesGoldstone:     brokenGoldstone,
		GoldstoneCountResonance:              resonance,
		CanonicalProtectedToBrokenMapDerived: false,
		SU2LActionOnContactScalarDerived:     false,
		CovariantDerivativeDerived:           false,
		GaugeBosonMassMatrixDerived:          false,
		GaugeEatingTheoremDerived:            false,
		Classification:                       class,
		TruthStatement:                       truth(resonance),
		MissingData: []string{
			"U-18C1-PROTECTED-BROKEN-MAP: derive a canonical isometry or intertwiner from the three protected contact directions to the three broken electroweak generator directions",
			"U-18C2-SCALAR-CONTACT-SU2L-ACTION: derive the SU(2)_L action directly on the finite scalar/contact active frame, not only on the audited matter doublet table",
			"U-18C3-COVARIANT-DERIVATIVE: construct the finite covariant derivative DΦ and its kinetic normalization",
			"U-18C4-GAUGE-BOSON-MASS-MATRIX: derive the dimensionless W/Z mass matrix before any physical mass-scale bridge",
		},
	}, nil
}

func truth(resonance bool) string {
	if resonance {
		return "The finite data now has the correct Goldstone/electroweak count: four real scalar directions split into one radial plus three angular directions, and SU(2)_L×U(1)_Y→U(1)_em also has three broken directions. The three protected contact directions match this count, but no canonical gauge-eating map has yet been derived."
	}
	return "The finite data does not yet even satisfy the count-level Goldstone/electroweak resonance."
}

func FormatMissing(values []string) string {
	out := ""
	for i, v := range values {
		if i > 0 {
			out += "; "
		}
		out += v
	}
	return out
}
