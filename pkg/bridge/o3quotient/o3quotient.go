// Package o3quotient implements Gate 90: O(3) gauge quotient / physical
// orientation audit.
//
// Gates 87-89 left a three-dimensional protected contact carrier with an
// unresolved O(3) frame freedom.  This package distinguishes two questions:
//
//  1. Do the currently implemented intrinsic protected-contact observables
//     depend on the choice of protected O(3) frame?
//  2. Has a finite action selected a physical orientation inside that frame?
//
// The current answer is deliberately nuanced.  With only the abstract protected
// metric and flat protected curvature restriction, the protected-contact-only
// observables are O(3)-invariant, so the frame freedom should be quotiented in
// that sector.  However, a complete physical no-orientation theorem is not yet
// proven: a future protected BF/contact action or coupling to the broken scalar
// directions could select orientation data.  Therefore the engine treats the
// O(3) freedom as a gauge quotient for present diagnostics, while keeping the
// physical-orientation question open.
package o3quotient

import (
	"fmt"
	"strings"
	"sync"

	"github.com/bagherbal/asha-engine/pkg/bridge/protectedconnection"
)

type Analysis struct {
	ProtectedConnection protectedconnection.Analysis

	ProtectedDimension  int
	O3Dimension         int
	AbstractFrameFamily string

	IntrinsicMetricIsO3Invariant                bool
	ProtectedCurvatureIsO3Invariant             bool
	ProtectedContactObservablesFrameIndependent bool

	DiagonalSpurionExists              bool
	DiagonalSpurionIntrinsic           bool
	DiagonalSpurionPhysicalOrientation bool

	CurrentDataSupportsGaugeQuotient bool
	FullNoOrientationTheoremProven   bool
	PhysicalOrientationSelected      bool
	NewFiniteActionRequired          bool

	QuotientStatement           string
	TruthStatement              string
	RecommendedNextGate         string
	GaugeQuotientEvidence       []string
	PhysicalOrientationOpenings []string
	RejectedMoves               []string
	RemainingUnknowns           []string
}

var (
	defaultOnce  sync.Once
	defaultValue Analysis
	defaultErr   error
)

func BuildDefault() (Analysis, error) {
	defaultOnce.Do(func() {
		pc, err := protectedconnection.BuildDefault()
		if err != nil {
			defaultErr = err
			return
		}
		defaultValue, defaultErr = Build(pc)
	})
	return defaultValue, defaultErr
}

func Build(pc protectedconnection.Analysis) (Analysis, error) {
	if pc.ProtectedDimension != 3 {
		return Analysis{}, fmt.Errorf("O(3) quotient audit expects protected dimension 3, got %d", pc.ProtectedDimension)
	}

	metricInvariant := pc.AbstractSO3ConnectionExists && !pc.CanonicalProtectedFrameDerived
	curvatureInvariant := pc.ContactCurvatureFlatOnProtected && pc.ContactCurvatureSpanRank == 0
	frameIndependent := metricInvariant && curvatureInvariant && !pc.IntrinsicProtectedOperatorDerived

	gaugeEvidence := []string{
		"protected carrier has only the abstract Euclidean metric I3 at this stage",
		"implemented contact-side curvature restriction is zero on the protected carrier",
		"no intrinsic protected BF/contact operator or connection has been selected",
		"all protected-contact-only scalar diagnostics are invariant under protected-frame rotations",
	}
	openings := []string{
		"a future protected BF/contact action could select a nonzero so(3) connection",
		"a future protected-to-broken intertwiner could turn orientation into physical gauge-eating data",
		"a future coupling to generation texture could make a protected frame observable",
	}
	rejected := []string{
		"choosing an arbitrary O(3) frame to match broken generators",
		"using the diagonal generation spurion as intrinsic protected-frame data",
		"pulling back the broken-generator metric before deriving the protected-to-broken map",
		"declaring the O(3) freedom physical without an action term that observes it",
	}

	quotient := "For the currently implemented protected-contact sector, the unresolved O(3) frame freedom behaves as gauge: intrinsic protected observables are frame-independent."
	truth := "Gate 90 does not select a physical protected orientation. It sharpens the previous obstruction: with the current finite data, protected-frame rotations leave the intrinsic protected-contact sector unchanged, so the O(3) freedom should be quotiented for present diagnostics. A full no-physical-orientation theorem is still open because future protected BF/contact or coupling terms could make orientation observable."

	return Analysis{
		ProtectedConnection:                         pc,
		ProtectedDimension:                          3,
		O3Dimension:                                 3,
		AbstractFrameFamily:                         "O(3)",
		IntrinsicMetricIsO3Invariant:                metricInvariant,
		ProtectedCurvatureIsO3Invariant:             curvatureInvariant,
		ProtectedContactObservablesFrameIndependent: frameIndependent,
		DiagonalSpurionExists:                       pc.DiagonalSpurionAvailable,
		DiagonalSpurionIntrinsic:                    pc.DiagonalSpurionIntrinsicToProtected,
		DiagonalSpurionPhysicalOrientation:          false,
		CurrentDataSupportsGaugeQuotient:            frameIndependent,
		FullNoOrientationTheoremProven:              false,
		PhysicalOrientationSelected:                 false,
		NewFiniteActionRequired:                     true,
		QuotientStatement:                           quotient,
		TruthStatement:                              truth,
		RecommendedNextGate:                         "Gate 91 — Gauge-Quotiented Protected-to-Broken Correspondence Audit",
		GaugeQuotientEvidence:                       gaugeEvidence,
		PhysicalOrientationOpenings:                 openings,
		RejectedMoves:                               rejected,
		RemainingUnknowns: []string{
			"U-18C7B1-O3-GAUGE-QUOTIENT: formalize the quotient of protected-frame rotations in the intrinsic protected sector",
			"U-18C7B2-ORIENTATION-OBSERVABLE: determine whether any future finite action observes protected orientation",
			"U-18C7C1-PROTECTED-BROKEN-INTERTWINER: derive the protected-to-broken map only after quotienting arbitrary O(3) choices",
			"U-18C7D1-GAUGE-EATING-COMPLETE: combine scalar kinetic action, gauge Hessian, and quotient-safe intertwiner",
		},
	}, nil
}

func Join(xs []string) string { return strings.Join(xs, "; ") }
