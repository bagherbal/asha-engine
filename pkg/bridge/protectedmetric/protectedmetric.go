// Package protectedmetric implements Gate 88: protected-contact metric /
// connection form search.
//
// Gate 87 found a real 3+3+3 resonance: three protected contact directions,
// three scalar angular directions, and three broken gauge-generator images.  It
// also found the precise obstruction: an abstract O(3) family of isometries
// remains because the protected contact side has not yet been equipped with a
// finite metric, connection, or action-selected frame.
//
// Gate 88 audits the protected side.  It distinguishes a harmless abstract
// Euclidean metric on a three-dimensional vector space from a derived finite
// protected-contact metric.  It also tests whether the broken-image metric can
// be pulled back to the protected side.  That pullback would require the very
// protected-to-broken intertwiner that Gate 87 left open, so it is recorded as
// circular rather than accepted.
package protectedmetric

import (
	"fmt"
	"strings"
	"sync"

	"github.com/bagherbal/asha-engine/pkg/bridge/protectedintertwiner"
)

type Analysis struct {
	Intertwiner protectedintertwiner.Analysis

	ProtectedDimension         int
	BrokenImageDimension       int
	BrokenImageMetricAvailable bool
	BrokenImageMetricCondition float64

	AbstractEuclideanMetricAvailable   bool
	AbstractEuclideanMetricCanonical   bool
	AbstractEuclideanMetricTrace       float64
	AbstractEuclideanMetricDeterminant float64

	BrokenMetricPullbackRequiresIntertwiner bool
	BrokenMetricPullbackDerived             bool
	PullbackCircularityDetected             bool

	FiniteProtectedMetricDerived     bool
	FiniteProtectedConnectionDerived bool
	CanonicalProtectedFrameDerived   bool
	O3FreedomReduced                 bool
	O3FreedomLikelyGauge             bool
	O3FreedomProvenGauge             bool

	CandidateMetricSources []string
	RejectedMetricSources  []string
	TruthStatement         string
	RecommendedNextGate    string
	RemainingUnknowns      []string
}

var (
	defaultOnce  sync.Once
	defaultValue Analysis
	defaultErr   error
)

func BuildDefault() (Analysis, error) {
	defaultOnce.Do(func() {
		pi, err := protectedintertwiner.BuildDefault()
		if err != nil {
			defaultErr = err
			return
		}
		defaultValue, defaultErr = Build(pi)
	})
	return defaultValue, defaultErr
}

func Build(pi protectedintertwiner.Analysis) (Analysis, error) {
	if pi.ProtectedContactDirections != 3 {
		return Analysis{}, fmt.Errorf("protected metric audit expects a 3D protected carrier, got %d", pi.ProtectedContactDirections)
	}
	if pi.BrokenImageRank != 3 {
		return Analysis{}, fmt.Errorf("protected metric audit expects a 3D broken-image rank, got %d", pi.BrokenImageRank)
	}

	candidates := []string{
		"abstract Euclidean metric I_3 on the protected carrier",
		"pullback of the broken-generator image metric through a protected-to-broken intertwiner",
		"future finite BF/contact connection metric from a protected-carrier action",
	}
	rejected := []string{
		"dimension count alone: 3 protected directions do not define a canonical frame",
		"pullback metric without an intertwiner: circular, because the map is the object being sought",
		"arbitrary O(3) gauge fixing: would choose coordinates, not derive contact geometry",
	}

	truth := "Gate 88 shows that the protected contact side has an abstract Euclidean metric, but not yet a derived finite protected-contact metric or connection form. Pulling back the broken-generator image metric would be circular until a protected-to-broken intertwiner is selected. Therefore the O(3) freedom from Gate 87 is not reduced by current finite data; it may ultimately be pure gauge, but that has not yet been proven."

	return Analysis{
		Intertwiner:                             pi,
		ProtectedDimension:                      pi.ProtectedContactDirections,
		BrokenImageDimension:                    pi.BrokenImageRank,
		BrokenImageMetricAvailable:              pi.BrokenImageMetricAvailable,
		BrokenImageMetricCondition:              pi.BrokenImageCondition,
		AbstractEuclideanMetricAvailable:        true,
		AbstractEuclideanMetricCanonical:        true,
		AbstractEuclideanMetricTrace:            3,
		AbstractEuclideanMetricDeterminant:      1,
		BrokenMetricPullbackRequiresIntertwiner: true,
		BrokenMetricPullbackDerived:             false,
		PullbackCircularityDetected:             true,
		FiniteProtectedMetricDerived:            false,
		FiniteProtectedConnectionDerived:        false,
		CanonicalProtectedFrameDerived:          false,
		O3FreedomReduced:                        false,
		O3FreedomLikelyGauge:                    true,
		O3FreedomProvenGauge:                    false,
		CandidateMetricSources:                  candidates,
		RejectedMetricSources:                   rejected,
		TruthStatement:                          truth,
		RecommendedNextGate:                     "Gate 89 — Protected-Carrier Operator / BF Contact Connection Search",
		RemainingUnknowns: []string{
			"U-18C7A1-PROTECTED-CARRIER-OPERATOR: derive an operator acting intrinsically on the three protected contact directions",
			"U-18C7A2-PROTECTED-CONNECTION-FORM: derive a finite connection/curvature form on the protected carrier",
			"U-18C7B1-O3-GAUGE-QUOTIENT: prove whether the O(3) freedom is pure gauge or carries physical orientation data",
			"U-18C7C1-PROTECTED-BROKEN-INTERTWINER: derive the protected-to-broken map after the protected metric/connection is known",
		},
	}, nil
}

func FormatFloat(x float64) string { return fmt.Sprintf("%.10f", x) }
func Join(xs []string) string      { return strings.Join(xs, "; ") }
