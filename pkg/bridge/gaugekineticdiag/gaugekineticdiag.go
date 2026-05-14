// Package gaugekineticdiag implements Gate 94: gauge-kinetic Hessian
// diag(1,1,4) action-selection audit.
//
// Gate 93 exposed a strong diagnostic: the quotient-safe broken-image metric
// becomes isotropic if the neutral broken generator is normalized by 1/2.  In
// the original raw gauge-field coordinates this corresponds to the candidate
// kinetic Hessian diag(1,1,4) for the three broken directions.
//
// Gate 94 asks the next, stricter question: is diag(1,1,4) selected by a
// finite scalar/gauge action, or is it only the basis metric required to compare
// broken-generator images fairly?  The answer at this stage is disciplined: it
// is a coherent action-candidate and it exactly removes the diagnostic
// anisotropy, but the finite action second variation has not selected it.
package gaugekineticdiag

import (
	"fmt"
	"math"
	"strings"
	"sync"

	"github.com/bagherbal/asha-engine/pkg/bridge/brokengeneratornorm"
)

type Analysis struct {
	BrokenGeneratorNorm brokengeneratornorm.Analysis

	CandidateName        string
	BrokenDirections     []string
	CandidateDiagonal    []float64
	CandidateDeterminant float64
	CandidatePositive    bool

	RawMetricDiagonal        []float64
	NormalizedMetricDiagonal []float64
	WhitenedMetricDiagonal   []float64
	WhitenedCondition        float64
	WhitenedExact            bool

	NeutralFactor                  float64
	NeutralFactorDerivedFromMetric bool

	CompatibleWithMassDiagnostic bool
	SelectedByFiniteAction       bool
	ScalarKineticActionSelected  bool
	GaugeHessianSelected         bool
	PhysicalCouplingsDerived     bool
	PhysicalMassesDerived        bool

	TruthStatement      string
	RejectedClaims      []string
	RemainingUnknowns   []string
	RecommendedNextGate string
}

var (
	defaultOnce  sync.Once
	defaultValue Analysis
	defaultErr   error
)

func BuildDefault() (Analysis, error) {
	defaultOnce.Do(func() {
		bg, err := brokengeneratornorm.BuildDefault()
		if err != nil {
			defaultErr = err
			return
		}
		defaultValue, defaultErr = Build(bg, 1e-10)
	})
	return defaultValue, defaultErr
}

func Build(bg brokengeneratornorm.Analysis, eps float64) (Analysis, error) {
	if eps <= 0 {
		eps = 1e-10
	}
	if len(bg.RawCoordinateKineticCandidate) != 3 || len(bg.RawBrokenMetricDiagonal) != 3 {
		return Analysis{}, fmt.Errorf("Gate 94 requires Gate 93 3D broken metric and kinetic candidate")
	}
	k := append([]float64(nil), bg.RawCoordinateKineticCandidate...)
	raw := append([]float64(nil), bg.RawBrokenMetricDiagonal...)
	normed := append([]float64(nil), bg.NormalizedMetricDiagonal...)

	det := k[0] * k[1] * k[2]
	positive := k[0] > eps && k[1] > eps && k[2] > eps

	// In raw coordinates the quadratic form K=diag(1,1,4) exactly whitens the
	// raw image metric diag(c,c,4c) when the neutral coordinate is rescaled by
	// 1/sqrt(K_33)=1/2.  This is equivalent to Gate 93's normalized basis.
	white := make([]float64, 3)
	for i := range raw {
		white[i] = raw[i] / k[i]
	}
	mn, mx := minMax(white)
	cond := math.Inf(1)
	if mn > eps {
		cond = mx / mn
	}
	exact := math.Abs(cond-1) < 1e-8 && math.Abs(k[2]-4) < 1e-8

	truth := "Gate 94 audits the raw-coordinate Hessian candidate diag(1,1,4). It is exactly the metric that whitens the Gate 92/93 broken-generator image anisotropy: raw diag(c,c,4c) becomes diag(c,c,c). This makes diag(1,1,4) a coherent gauge-kinetic candidate for the broken-coordinate basis. However, the engine still has not derived a finite scalar/gauge action whose second variation selects this Hessian. Therefore it remains an action-candidate, not a physical coupling theorem."

	return Analysis{
		BrokenGeneratorNorm:            bg,
		CandidateName:                  "K_broken_raw = diag(1,1,4)",
		BrokenDirections:               []string{"T1", "T2", "Z_raw=T3-Y_phi"},
		CandidateDiagonal:              k,
		CandidateDeterminant:           det,
		CandidatePositive:              positive,
		RawMetricDiagonal:              raw,
		NormalizedMetricDiagonal:       normed,
		WhitenedMetricDiagonal:         white,
		WhitenedCondition:              cond,
		WhitenedExact:                  exact,
		NeutralFactor:                  bg.NeutralNormalization,
		NeutralFactorDerivedFromMetric: bg.MetricSelectsDiagnosticBasis,
		CompatibleWithMassDiagnostic:   exact && positive,
		SelectedByFiniteAction:         false,
		ScalarKineticActionSelected:    false,
		GaugeHessianSelected:           false,
		PhysicalCouplingsDerived:       false,
		PhysicalMassesDerived:          false,
		TruthStatement:                 truth,
		RejectedClaims: []string{
			"diag(1,1,4) is automatically the physical gauge kinetic Hessian",
			"metric whitening alone derives g2, gY, thetaW, or alpha",
			"dimensionless W/Z diagnostic masses are physical masses",
			"neutral normalization 1/2 is an observed input",
		},
		RemainingUnknowns: []string{
			"U-18C8B-FINITE-GAUGE-ACTION: construct scalar/gauge kinetic action and compute second variation",
			"U-18C6-SCALAR-KINETIC-ACTION: derive scalar/contact kinetic normalization",
			"U-18C8C-COUPLING-RATIO: determine whether diag(1,1,4) fixes a coupling ratio after action selection",
			"U-18C9-RG-PHYSICAL-COUPLINGS: connect action-selected kinetic terms to running couplings and alpha",
		},
		RecommendedNextGate: "Gate 95 — Broken-Sector Action Second Variation / Kinetic Hessian Search",
	}, nil
}

func minMax(xs []float64) (float64, float64) {
	mn := math.Inf(1)
	mx := math.Inf(-1)
	for _, x := range xs {
		if x < mn {
			mn = x
		}
		if x > mx {
			mx = x
		}
	}
	return mn, mx
}

func Join(xs []string) string { return strings.Join(xs, "; ") }
