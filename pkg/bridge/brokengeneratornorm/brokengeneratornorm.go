// Package brokengeneratornorm implements Gate 93: normalized broken-generator
// basis / gauge-kinetic candidate search.
//
// Gate 92 showed that the quotient-safe broken-image metric has a raw
// neutral-to-charged anisotropy of exactly 4, and that this anisotropy is
// removed by scaling the neutral broken generator by 1/2.  Gate 93 promotes
// that observation into a precise normalization audit:
//
//   - the normalization factor is determined by the diagnostic metric ratio,
//     not by an arbitrary convention;
//   - the normalized broken-generator image metric becomes isotropic;
//   - the equivalent raw-coordinate gauge-kinetic candidate is exposed;
//   - the candidate is still not selected by a finite gauge-field action.
//
// This gate therefore sharpens the distinction between a canonical diagnostic
// basis and a physical gauge-kinetic Hessian.  It still refuses to claim W/Z
// masses, physical electroweak mixing, or alpha.
package brokengeneratornorm

import (
	"fmt"
	"math"
	"strings"
	"sync"

	"github.com/bagherbal/asha-engine/pkg/bridge/brokenmetric"
)

type Analysis struct {
	BrokenMetric brokenmetric.Analysis

	RawGeneratorNames        []string
	RawBrokenMetricDiagonal  []float64
	RawCondition             float64
	NormalizationFactors     []float64
	NeutralNormalization     float64
	NormalizedMetricDiagonal []float64
	NormalizedCondition      float64
	NormalizationExact       bool

	// In the raw gauge-field coordinate basis, rescaling the neutral generator
	// by n=1/2 corresponds to the diagnostic inverse-square kinetic candidate
	// diag(1,1,1/n^2)=diag(1,1,4).  This is useful, but it is still not an
	// action-selected Hessian.
	RawCoordinateKineticCandidate []float64
	KineticCandidatePositive      bool
	KineticCandidateDeterminant   float64

	MetricSelectsDiagnosticBasis bool
	FiniteActionSelectsBasis     bool
	GaugeKineticHessianSelected  bool
	PhysicalAnisotropyDerived    bool
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
		bm, err := brokenmetric.BuildDefault()
		if err != nil {
			defaultErr = err
			return
		}
		defaultValue, defaultErr = Build(bm, 1e-10)
	})
	return defaultValue, defaultErr
}

func Build(bm brokenmetric.Analysis, eps float64) (Analysis, error) {
	if eps <= 0 {
		eps = 1e-10
	}
	if !bm.IsotropizationPossible || bm.NeutralNormFactor <= eps {
		return Analysis{}, fmt.Errorf("Gate 93 requires Gate 92 isotropization data; factor=%.12g", bm.NeutralNormFactor)
	}
	charged := bm.ChargedEigenvalue
	neutral := bm.NeutralEigenvalue
	n := bm.NeutralNormFactor
	raw := []float64{charged, charged, neutral}
	factors := []float64{1, 1, n}
	normed := []float64{raw[0] * factors[0] * factors[0], raw[1] * factors[1] * factors[1], raw[2] * factors[2] * factors[2]}
	mn, mx := minMax(normed)
	cond := math.Inf(1)
	if mn > eps {
		cond = mx / mn
	}
	kin := []float64{1, 1, 1 / (n * n)}
	det := kin[0] * kin[1] * kin[2]
	exact := math.Abs(cond-1) < 1e-8 && math.Abs(n-0.5) < 1e-8

	truth := "Gate 93 converts the Gate 92 metric observation into a normalized broken-generator basis diagnostic. The neutral broken generator is scaled by 1/2, which isotropizes the quotient-safe broken-image metric. Equivalently, in raw gauge-field coordinates this exposes a candidate diagonal kinetic metric diag(1,1,4). This is a canonical diagnostic extracted from the broken-image metric, but it is still not a finite action-selected gauge kinetic Hessian and therefore does not determine physical W/Z masses, gauge couplings, or alpha."

	return Analysis{
		BrokenMetric: bm,

		RawGeneratorNames:        []string{"T1", "T2", "Z_raw=T3-Y_phi"},
		RawBrokenMetricDiagonal:  raw,
		RawCondition:             bm.RawCondition,
		NormalizationFactors:     factors,
		NeutralNormalization:     n,
		NormalizedMetricDiagonal: normed,
		NormalizedCondition:      cond,
		NormalizationExact:       exact,

		RawCoordinateKineticCandidate: kin,
		KineticCandidatePositive:      kin[0] > eps && kin[1] > eps && kin[2] > eps,
		KineticCandidateDeterminant:   det,

		MetricSelectsDiagnosticBasis: true,
		FiniteActionSelectsBasis:     false,
		GaugeKineticHessianSelected:  false,
		PhysicalAnisotropyDerived:    false,
		PhysicalMassesDerived:        false,

		TruthStatement: truth,
		RejectedClaims: []string{
			"neutral factor 1/2 is a measured weak-mixing input",
			"diag(1,1,4) is already the physical gauge kinetic Hessian",
			"isotropized broken metric proves physical W/Z masses",
			"basis normalization alone determines alpha or electroweak couplings",
		},
		RemainingUnknowns: []string{
			"U-18C8A-GAUGE-ACTION-HESSIAN: derive whether diag(1,1,4) is selected by finite gauge-field second variation",
			"U-18C6-SCALAR-KINETIC-ACTION: derive scalar/contact kinetic normalization instead of using active-frame diagnostics",
			"U-18C7C2-QUOTIENTED-INTERTWINER: derive a quotient-safe protected-to-broken map after normalization",
			"U-18C9-PHYSICAL-COUPLINGS: derive g2, gY, thetaW, and alpha through action-selected kinetic terms and RG flow",
		},
		RecommendedNextGate: "Gate 94 — Gauge-Kinetic Hessian diag(1,1,4) Action-Selection Audit",
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
