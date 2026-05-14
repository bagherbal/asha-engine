// Package scalarvacuum implements Gate 86: scalar vacuum orientation / finite
// minimizer search.
//
// Gate 84-85 used the usual lower-component real unitary-gauge vector to build
// the dimensionless W/Z/photon and Goldstone-image diagnostics.  That choice is
// physically natural, but it must not be confused with a finite theorem that
// the exact vector has been selected by the scalar/contact action.
//
// This package separates three layers:
//
//  1. the shifted radial normal form V(r)=lambda(r^2-r0^2)^2 selects a radius,
//     not an orientation; its vacuum manifold is S^3 at fixed radius;
//  2. the finite scalar/contact response S_Phi=(a,a,b,b), with a>b, selects the
//     lower two-dimensional active pair plane as a constrained minimizer;
//  3. inside that lower plane there remains an S^1 orientation freedom, so the
//     particular real unitary-gauge vector used by Gate 84 is still a gauge /
//     complex-phase convention rather than a uniquely derived finite vector.
package scalarvacuum

import (
	"fmt"
	"math"
	"sync"

	"github.com/bagherbal/asha-engine/pkg/bridge/gaugeeating"
	"github.com/bagherbal/asha-engine/pkg/dynamics/scalarpotential"
)

type Analysis struct {
	ScalarPotential scalarpotential.Analysis
	GaugeEating     gaugeeating.Analysis

	ActiveRealDimension int
	VacuumRadiusSquared float64
	VacuumRadius        float64

	RadialNormalFormSelectsRadius bool
	RadialNormalFormSelectsVector bool
	RadialVacuumManifold          string
	RadialVacuumManifoldDimension int

	ActiveSpectrum     []float64
	HighPairEigenvalue float64
	LowPairEigenvalue  float64
	PairSplit          float64
	LowPairDimension   int
	HighPairDimension  int
	LowPairSelected    bool
	HighPairSuppressed bool
	MinResponseEnergy  float64
	MaxResponseEnergy  float64
	EnergyGapAtRadius  float64

	DiagnosticVacuumVector         []float64
	DiagnosticVacuumEnergy         float64
	DiagnosticVacuumIsMinimizer    bool
	ResidualPhaseFreedomDimension  int
	UnitaryGaugeVectorSelected     bool
	CanonicalPhaseSelected         bool
	FiniteVacuumOrientationDerived bool

	GaugeEatingDiagnosticAvailable bool
	FullGaugeEatingTheoremDerived  bool

	TruthStatement      string
	RecommendedNextGate string
	RemainingUnknowns   []string
}

var (
	defaultOnce  sync.Once
	defaultValue Analysis
	defaultErr   error
)

func BuildDefault() (Analysis, error) {
	defaultOnce.Do(func() {
		sp, err := scalarpotential.BuildDefault()
		if err != nil {
			defaultErr = err
			return
		}
		ge, err := gaugeeating.BuildDefault()
		if err != nil {
			defaultErr = err
			return
		}
		defaultValue, defaultErr = Build(sp, ge, 1e-10)
	})
	return defaultValue, defaultErr
}

func Build(sp scalarpotential.Analysis, ge gaugeeating.Analysis, eps float64) (Analysis, error) {
	if eps <= 0 {
		eps = 1e-10
	}
	if sp.ActiveRealDimension != 4 {
		return Analysis{}, fmt.Errorf("scalar vacuum orientation audit expects 4 active real directions, got %d", sp.ActiveRealDimension)
	}
	if len(sp.ActiveSpectrum) != 4 {
		return Analysis{}, fmt.Errorf("scalar vacuum orientation audit expects 4 active spectrum entries, got %d", len(sp.ActiveSpectrum))
	}
	if len(ge.ScalarCovariant.VacuumVector) != 4 {
		return Analysis{}, fmt.Errorf("scalar vacuum orientation audit expects Gate 84 vacuum vector of length 4")
	}

	spectrum := append([]float64(nil), sp.ActiveSpectrum...)
	high, low := extrema(spectrum)
	lowDim, highDim := countNear(spectrum, low, eps), countNear(spectrum, high, eps)
	pairSplit := high - low
	r2 := sp.VacuumRadiusSquared
	r := sp.VacuumRadius
	minEnergy := low * r2
	maxEnergy := high * r2
	gap := (high - low) * r2
	diagnostic := append([]float64(nil), ge.ScalarCovariant.VacuumVector...)
	diagnosticEnergy := quadraticDiagonal(spectrum, diagnostic)
	diagnosticIsMin := math.Abs(diagnosticEnergy-minEnergy) < 1e-8

	lowPairSelected := lowDim == 2 && highDim == 2 && pairSplit > eps
	truth := "The finite scalar/contact response partially selects the vacuum orientation: at fixed finite radius it selects the lower active pair plane. The radial normal form itself only selects a radius, and the residual S^1 phase inside the lower pair remains unselected. Therefore the lower-component real vector used in the W/Z diagnostic is a valid minimizer representative, but not yet a uniquely derived finite vacuum orientation."

	return Analysis{
		ScalarPotential: sp,
		GaugeEating:     ge,

		ActiveRealDimension: sp.ActiveRealDimension,
		VacuumRadiusSquared: r2,
		VacuumRadius:        r,

		RadialNormalFormSelectsRadius: sp.ShiftedNormalFormAvailable,
		RadialNormalFormSelectsVector: false,
		RadialVacuumManifold:          "S^3 at |phi|=r0 before quotient/orientation data",
		RadialVacuumManifoldDimension: 3,

		ActiveSpectrum:     spectrum,
		HighPairEigenvalue: high,
		LowPairEigenvalue:  low,
		PairSplit:          pairSplit,
		LowPairDimension:   lowDim,
		HighPairDimension:  highDim,
		LowPairSelected:    lowPairSelected,
		HighPairSuppressed: lowPairSelected,
		MinResponseEnergy:  minEnergy,
		MaxResponseEnergy:  maxEnergy,
		EnergyGapAtRadius:  gap,

		DiagnosticVacuumVector:         diagnostic,
		DiagnosticVacuumEnergy:         diagnosticEnergy,
		DiagnosticVacuumIsMinimizer:    diagnosticIsMin,
		ResidualPhaseFreedomDimension:  1,
		UnitaryGaugeVectorSelected:     false,
		CanonicalPhaseSelected:         false,
		FiniteVacuumOrientationDerived: false,

		GaugeEatingDiagnosticAvailable: ge.GoldstoneImageTheoremDiagnostic,
		FullGaugeEatingTheoremDerived:  ge.FiniteGaugeEatingTheoremDerived,

		TruthStatement:      truth,
		RecommendedNextGate: "Gate 87 — Protected-Contact / Broken-Generator Intertwiner Search",
		RemainingUnknowns: []string{
			"U-19A1A-LOW-PAIR-TO-COMPLEX-PHASE: derive the phase/orientation inside the lower active scalar pair, not only the pair plane",
			"U-18C7-GAUGE-EATING-INTERTWINER: derive a canonical isometry from protected contact directions to broken generator images",
			"U-19C1-SCALAR-KINETIC-ACTION: derive scalar kinetic normalization from finite contact/BF action",
			"U-19C2-VACUUM-ORIENTATION-ACTION: derive a finite action term that chooses a representative on the low-pair S^1 or prove it is pure gauge",
		},
	}, nil
}

func extrema(values []float64) (high, low float64) {
	high, low = values[0], values[0]
	for _, v := range values[1:] {
		if v > high {
			high = v
		}
		if v < low {
			low = v
		}
	}
	return high, low
}

func countNear(values []float64, target, eps float64) int {
	count := 0
	for _, v := range values {
		if math.Abs(v-target) < eps {
			count++
		}
	}
	return count
}

func quadraticDiagonal(diag []float64, v []float64) float64 {
	s := 0.0
	for i := range diag {
		s += diag[i] * v[i] * v[i]
	}
	return s
}

func FormatSpectrum(values []float64) string {
	out := "["
	for i, v := range values {
		if i > 0 {
			out += ", "
		}
		out += fmt.Sprintf("%.10f", v)
	}
	return out + "]"
}

func FormatVector(values []float64) string {
	out := "["
	for i, v := range values {
		if i > 0 {
			out += ", "
		}
		out += fmt.Sprintf("%.10f", v)
	}
	return out + "]"
}
