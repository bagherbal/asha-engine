// Package protectedconnection implements Gate 89: protected-carrier operator /
// BF contact connection search.
//
// Gates 87-88 exposed a real 3+3+3 resonance but left an O(3) identification
// freedom on the protected contact carrier.  This package asks whether the
// current finite BF/contact data contains an intrinsic operator or connection
// form on the three protected directions that can reduce that freedom.
//
// The result is deliberately conservative.  Abstract linear operators on a 3D
// carrier always exist, and the Higgs/contact anisotropy provides a useful
// bridge-level diagonal generation spurion.  However, the implemented
// contact-side curvature is flat on the protected carrier, and the diagonal
// spurion is not yet an intrinsic protected-contact connection.  Therefore the
// O(3) freedom is not reduced by a finite BF/contact theorem at this stage.
package protectedconnection

import (
	"fmt"
	"strings"
	"sync"

	"github.com/bagherbal/asha-engine/pkg/bridge/protectedmetric"
	"github.com/bagherbal/asha-engine/pkg/matter/gencurvature"
	"github.com/bagherbal/asha-engine/pkg/matter/generationbreak"
)

type Analysis struct {
	ProtectedMetric     protectedmetric.Analysis
	GenerationBreak     generationbreak.Analysis
	GenerationCurvature gencurvature.Analysis

	ProtectedDimension                  int
	EndomorphismDimension               int
	SkewConnectionDimension             int
	SymmetricMetricDeformationDimension int

	AbstractOperatorSpaceExists    bool
	AbstractSO3ConnectionExists    bool
	AbstractOperatorSpaceCanonical bool

	DiagonalSpurionAvailable            bool
	DiagonalSpurionEigenvalues          []float64
	DiagonalSpurionIntrinsicToProtected bool
	DiagonalSpurionReducesO3            bool

	ContactCurvatureOperators       int
	ContactCurvatureFlatOnProtected bool
	ContactCurvatureMaxNorm         float64
	ContactCurvatureSpanRank        int

	ActiveCurvatureNonzero  bool
	ActiveCurvatureMaxNorm  float64
	ActiveCurvatureSpanRank int

	BFContactConnectionDerived        bool
	IntrinsicProtectedOperatorDerived bool
	CanonicalProtectedFrameDerived    bool
	O3FreedomReduced                  bool
	O3FreedomProvenGauge              bool

	CandidateOperatorSources []string
	RejectedShortcuts        []string
	TruthStatement           string
	RecommendedNextGate      string
	RemainingUnknowns        []string
}

var (
	defaultOnce  sync.Once
	defaultValue Analysis
	defaultErr   error
)

func BuildDefault() (Analysis, error) {
	defaultOnce.Do(func() {
		pm, err := protectedmetric.BuildDefault()
		if err != nil {
			defaultErr = err
			return
		}
		gb, err := generationbreak.BuildDefault()
		if err != nil {
			defaultErr = err
			return
		}
		gc, err := gencurvature.BuildDefault()
		if err != nil {
			defaultErr = err
			return
		}
		defaultValue, defaultErr = Build(pm, gb, gc)
	})
	return defaultValue, defaultErr
}

func Build(pm protectedmetric.Analysis, gb generationbreak.Analysis, gc gencurvature.Analysis) (Analysis, error) {
	if pm.ProtectedDimension != 3 {
		return Analysis{}, fmt.Errorf("protected connection search expects protected dimension 3, got %d", pm.ProtectedDimension)
	}
	if gc.CarrierDimension != 3 {
		return Analysis{}, fmt.Errorf("generation curvature carrier dimension = %d, want 3", gc.CarrierDimension)
	}

	diagonalEigen := []float64(nil)
	if gb.DiagonalSpurionFound {
		diagonalEigen = append(diagonalEigen, gb.BestCandidate.Eigenvalues...)
	}

	candidates := []string{
		"abstract End(R^3) operators on the protected carrier",
		"abstract so(3) skew connection forms on the protected carrier",
		"Higgs/contact anisotropy diagonal spurion on the triality-generation bridge",
		"contact-side second-fundamental curvature R^K_AB restricted to the protected carrier",
		"future finite BF/contact connection from an action-level protected carrier",
	}
	rejected := []string{
		"choosing an arbitrary so(3) frame: coordinate fixing, not finite dynamics",
		"using the bridge-level diagonal spurion as a connection: it splits labels but does not supply a protected BF connection",
		"using active-sector curvature directly: it is nonzero on active Higgs/contact directions but does not couple into the protected carrier",
		"pulling back the broken-generator metric: circular before the protected-to-broken map is derived",
	}

	truth := "Gate 89 finds no intrinsic protected BF/contact connection yet. The protected carrier is three-dimensional and admits abstract End(R^3) and so(3) operator spaces, and the Higgs/contact anisotropy supplies a useful diagonal generation spurion. But the implemented contact-side second-fundamental curvature is flat on the protected carrier, while the nonzero curvature lives on the active Higgs/contact carrier. Therefore the remaining O(3) freedom is not reduced by current finite action data."

	return Analysis{
		ProtectedMetric:                     pm,
		GenerationBreak:                     gb,
		GenerationCurvature:                 gc,
		ProtectedDimension:                  3,
		EndomorphismDimension:               9,
		SkewConnectionDimension:             3,
		SymmetricMetricDeformationDimension: 6,
		AbstractOperatorSpaceExists:         true,
		AbstractSO3ConnectionExists:         true,
		AbstractOperatorSpaceCanonical:      false,
		DiagonalSpurionAvailable:            gb.DiagonalSpurionFound,
		DiagonalSpurionEigenvalues:          diagonalEigen,
		DiagonalSpurionIntrinsicToProtected: false,
		DiagonalSpurionReducesO3:            false,
		ContactCurvatureOperators:           len(gc.Operators),
		ContactCurvatureFlatOnProtected:     gc.NonzeroOperators == 0 && gc.MaxCurvatureNorm < 1e-8 && gc.OperatorSpanRank == 0,
		ContactCurvatureMaxNorm:             gc.MaxCurvatureNorm,
		ContactCurvatureSpanRank:            gc.OperatorSpanRank,
		ActiveCurvatureNonzero:              gc.ActiveNonzeroOperators > 0 && gc.ActiveMaxCurvatureNorm > 1e-8,
		ActiveCurvatureMaxNorm:              gc.ActiveMaxCurvatureNorm,
		ActiveCurvatureSpanRank:             gc.ActiveOperatorSpanRank,
		BFContactConnectionDerived:          false,
		IntrinsicProtectedOperatorDerived:   false,
		CanonicalProtectedFrameDerived:      false,
		O3FreedomReduced:                    false,
		O3FreedomProvenGauge:                false,
		CandidateOperatorSources:            candidates,
		RejectedShortcuts:                   rejected,
		TruthStatement:                      truth,
		RecommendedNextGate:                 "Gate 90 — O(3) Gauge Quotient / Physical Orientation Audit",
		RemainingUnknowns: []string{
			"U-18C7A1-PROTECTED-BF-CONNECTION: derive a nonzero finite connection/curvature form on the protected 3D carrier",
			"U-18C7A2-INTRINSIC-PROTECTED-OPERATOR: derive an operator from contact/BF action data rather than from arbitrary End(R^3)",
			"U-18C7B1-O3-GAUGE-QUOTIENT: determine whether the unresolved O(3) frame freedom is pure gauge",
			"U-18C7C1-PROTECTED-BROKEN-INTERTWINER: derive the protected-to-broken map after the protected carrier is equipped with real geometry",
		},
	}, nil
}

func FormatFloat(x float64) string      { return fmt.Sprintf("%.10f", x) }
func FormatScientific(x float64) string { return fmt.Sprintf("%.3e", x) }
func FormatSlice(xs []float64) string {
	if len(xs) == 0 {
		return "[]"
	}
	parts := make([]string, len(xs))
	for i, x := range xs {
		parts[i] = FormatFloat(x)
	}
	return "[" + strings.Join(parts, ", ") + "]"
}
func Join(xs []string) string { return strings.Join(xs, "; ") }
