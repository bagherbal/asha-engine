// Package quotientedcorrespondence implements Gate 91: gauge-quotiented
// protected-to-broken correspondence audit.
//
// Gate 90 established that, for the currently implemented protected-contact
// diagnostics, the unresolved O(3) protected-frame freedom behaves as gauge.
// This gate therefore refuses to compare protected and broken frames component
// by component.  It compares only quotient-safe data: dimensions, ranks,
// spectrum/condition of the broken image metric, and whether any invariant map
// survives after quotienting arbitrary O(3) frame choices.
//
// The result is deliberately not a completed gauge-eating theorem.  The count
// match survives the quotient, and the broken-generator image metric is a real
// finite invariant, but no canonical protected-to-broken intertwiner is derived.
package quotientedcorrespondence

import (
	"fmt"
	"math"
	"strings"
	"sync"

	"github.com/bagherbal/asha-engine/pkg/bridge/o3quotient"
	"github.com/bagherbal/asha-engine/pkg/bridge/protectedintertwiner"
)

type Analysis struct {
	O3Quotient           o3quotient.Analysis
	ProtectedIntertwiner protectedintertwiner.Analysis

	ProtectedDimension int
	BrokenDimension    int
	BrokenRank         int
	QuotientGroup      string
	QuotientDimension  int

	CountCorrespondenceSurvivesQuotient bool
	FrameComponentComparisonRejected    bool
	BrokenMetricSpectrumInvariant       bool
	ProtectedMetricSpectrum             []float64
	BrokenMetricMinEigen                float64
	BrokenMetricMaxEigen                float64
	BrokenMetricCondition               float64
	BrokenMetricIsIsotropic             bool
	MetricIsometryDerived               bool
	QuotientSafeIntertwinerDerived      bool
	GaugeEatingBridgeCompleted          bool

	QuotientSafeInvariants []string
	RejectedMoves          []string
	RemainingUnknowns      []string
	TruthStatement         string
	RecommendedNextGate    string
}

var (
	defaultOnce  sync.Once
	defaultValue Analysis
	defaultErr   error
)

func BuildDefault() (Analysis, error) {
	defaultOnce.Do(func() {
		q, err := o3quotient.BuildDefault()
		if err != nil {
			defaultErr = err
			return
		}
		pi, err := protectedintertwiner.BuildDefault()
		if err != nil {
			defaultErr = err
			return
		}
		defaultValue, defaultErr = Build(q, pi, 1e-10)
	})
	return defaultValue, defaultErr
}

func Build(q o3quotient.Analysis, pi protectedintertwiner.Analysis, eps float64) (Analysis, error) {
	if eps <= 0 {
		eps = 1e-10
	}
	if q.ProtectedDimension != 3 || pi.ProtectedContactDirections != 3 || pi.BrokenImageRank != 3 {
		return Analysis{}, fmt.Errorf("quotiented correspondence expects protected=3 and broken-image rank=3; got protected=%d/%d brokenRank=%d", q.ProtectedDimension, pi.ProtectedContactDirections, pi.BrokenImageRank)
	}

	cond := pi.BrokenImageCondition
	if math.IsInf(cond, 1) || math.IsNaN(cond) {
		cond = math.Inf(1)
	}
	isotropic := math.Abs(cond-1) < eps
	countSurvives := q.CurrentDataSupportsGaugeQuotient && pi.CountLevelResonance
	brokenMetricInvariant := pi.BrokenImageMetricAvailable && pi.BrokenImageMetricPositive

	truth := "After quotienting the arbitrary protected O(3) frame, the protected-to-broken bridge retains only quotient-safe data: a 3↔3 dimension/rank correspondence and the broken-generator image metric spectrum. Component-wise frame matching is rejected. The broken metric is anisotropic, so the current data does not provide a canonical metric isometry from the abstract protected I3 carrier to the broken image carrier. The gauge-eating bridge remains a quotient-safe resonance, not a completed theorem."

	return Analysis{
		O3Quotient:           q,
		ProtectedIntertwiner: pi,
		ProtectedDimension:   3,
		BrokenDimension:      pi.BrokenGeneratorDirections,
		BrokenRank:           pi.BrokenImageRank,
		QuotientGroup:        q.AbstractFrameFamily,
		QuotientDimension:    q.O3Dimension,

		CountCorrespondenceSurvivesQuotient: countSurvives,
		FrameComponentComparisonRejected:    true,
		BrokenMetricSpectrumInvariant:       brokenMetricInvariant,
		ProtectedMetricSpectrum:             []float64{1, 1, 1},
		BrokenMetricMinEigen:                pi.BrokenImageMinEigen,
		BrokenMetricMaxEigen:                pi.BrokenImageMaxEigen,
		BrokenMetricCondition:               cond,
		BrokenMetricIsIsotropic:             isotropic,
		MetricIsometryDerived:               false,
		QuotientSafeIntertwinerDerived:      false,
		GaugeEatingBridgeCompleted:          false,

		QuotientSafeInvariants: []string{
			"protected dimension = 3",
			"broken-generator image rank = 3",
			"electromagnetic null direction inherited from scalar covariant derivative gate",
			"broken-image metric eigenvalue range/condition is invariant under protected O(3) relabelling",
		},
		RejectedMoves: []string{
			"choosing an arbitrary protected O(3) frame to align with broken-generator coordinates",
			"calling the dimension match a canonical map",
			"pulling back the broken metric before deriving an intertwiner",
			"treating the anisotropic broken-image metric as equal to the abstract protected I3 metric",
		},
		RemainingUnknowns: []string{
			"U-18C7C2-QUOTIENTED-INTERTWINER: derive an O(3)-invariant protected-to-broken map or prove none exists",
			"U-18C7C3-BROKEN-METRIC-NORMALIZATION: derive whether broken-image anisotropy is physical, kinetic, or gauge-normalization data",
			"U-18C8-SCALAR-GAUGE-ACTION: combine scalar kinetic action and gauge Hessian before declaring gauge eating",
			"U-18C9-PROTECTED-ACTION-OBSERVABLE: determine whether any future finite action observes protected orientation after quotienting",
		},
		TruthStatement:      truth,
		RecommendedNextGate: "Gate 92 — Broken-Image Metric / Kinetic Normalization Audit",
	}, nil
}

func Join(xs []string) string { return strings.Join(xs, "; ") }
