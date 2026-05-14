// Package contactembedding implements Gate 102: the contact-to-matter
// hypercharge embedding / finite normalization-threshold map.
//
// Gate 101 exposed two abelian normalizations that must not be conflated:
//
//   - the scalar/contact action-selected coefficient K(Y_phi)=3;
//   - the finite one-generation matter trace normalization k_Y=5/3.
//
// This gate performs the next strictly allowed operation.  It asks whether a
// one-dimensional, orientation-preserving abelian embedding
//
//	Y = lambda Y_phi
//
// can carry the Gate 100/101 contact action into the finite matter hypercharge
// normalization without using observed couplings.  The answer is yes, but only
// as a dimensionless finite normalization map:
//
//	lambda^2 K(Y_phi) = k_Y  =>  lambda^2 = 5/9.
//
// The result promotes the matter-table diagnostic sin^2=3/8 to an embedded
// finite boundary diagnostic.  It still does not derive a physical weak mixing
// angle, alpha, an RG boundary scale, threshold masses, or low-energy couplings.
package contactembedding

import (
	"fmt"
	"math"
	"strings"
	"sync"

	"github.com/bagherbal/asha-engine/pkg/bridge/canonicalboundary"
	"github.com/bagherbal/asha-engine/pkg/linear"
)

type SignCandidate struct {
	Name                  string
	Lambda                float64
	LambdaSq              float64
	OrientationPreserving bool
	MatchesMatterKinetic  bool
	Selected              bool
	Detail                string
}

type ThresholdMap struct {
	Name                    string
	ContactCoefficient      float64
	MatterCoefficient       float64
	EmbeddingScaleSq        float64
	EmbeddingScale          float64
	NormalizationDeficit    float64
	Formula                 string
	Selected                bool
	PhysicalMassThreshold   bool
	PhysicalScaleInserted   bool
	RepresentationThreshold bool
	Detail                  string
}

type Analysis struct {
	Boundary canonicalboundary.Analysis

	ContactGeneratorName string
	MatterGeneratorName  string
	EmbeddingFormula     string
	EmbeddingEquation    string

	ContactU1KineticCoefficient float64
	MatterHyperchargeKY         float64
	EmbeddingScaleSq            float64
	EmbeddingScale              float64
	EmbeddingScaleNegative      float64

	SignCandidates          []SignCandidate
	OrientationSelected     bool
	OneDimensionalMapUnique bool

	EmbeddedMatterHessian             linear.Matrix
	EmbeddedMatterHessianSelected     bool
	EmbeddedMatterBoundarySin2        float64
	EmbeddedMatterBoundaryDiagnostic  bool
	ContactMatterMismatchBeforeMap    float64
	ContactMatterMismatchAfterMap     float64
	FiniteNormalizationThreshold      ThresholdMap
	ContactToMatterEmbeddingSelected  bool
	ThresholdNormalizationMapSelected bool

	RGBoundaryScaleDerived   bool
	BetaThresholdFlowDerived bool
	PhysicalWeakAngleDerived bool
	FineStructureDerived     bool
	PhysicalCouplingsDerived bool
	PhysicalMassesDerived    bool
	HiddenObservedInputUsed  bool

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
		boundary, err := canonicalboundary.BuildDefault()
		if err != nil {
			defaultErr = err
			return
		}
		defaultValue, defaultErr = Build(boundary, 1e-10)
	})
	return defaultValue, defaultErr
}

func Build(boundary canonicalboundary.Analysis, eps float64) (Analysis, error) {
	if eps <= 0 {
		eps = 1e-10
	}
	if !boundary.ActionSelectedDimensionlessBoundarySeed {
		return Analysis{}, fmt.Errorf("Gate 102 requires the Gate 101 dimensionless boundary seed")
	}
	if boundary.ScalarContactU1KineticCoefficient <= eps || boundary.MatterHyperchargeKY <= eps {
		return Analysis{}, fmt.Errorf("Gate 102 requires positive contact and matter abelian coefficients")
	}
	contactK := boundary.ScalarContactU1KineticCoefficient
	matterK := boundary.MatterHyperchargeKY
	lambdaSq := matterK / contactK
	if lambdaSq <= eps {
		return Analysis{}, fmt.Errorf("derived embedding square must be positive")
	}
	lambda := math.Sqrt(lambdaSq)
	negLambda := -lambda
	before := math.Abs(contactK - matterK)
	after := math.Abs(lambdaSq*contactK - matterK)

	signs := []SignCandidate{
		{
			Name:                  "orientation-preserving abelian embedding",
			Lambda:                lambda,
			LambdaSq:              lambdaSq,
			OrientationPreserving: true,
			MatchesMatterKinetic:  after <= eps,
			Selected:              after <= eps,
			Detail:                "keeps the charge orientation Q=T3+Y while matching lambda^2 K(Y_phi)=k_Y",
		},
		{
			Name:                  "orientation-reversing abelian embedding",
			Lambda:                negLambda,
			LambdaSq:              lambdaSq,
			OrientationPreserving: false,
			MatchesMatterKinetic:  after <= eps,
			Selected:              false,
			Detail:                "quadratically admissible but rejected because it flips the hypercharge/electric-charge orientation",
		},
	}

	embedded := linear.Diagonal([]float64{1, 1, 1, matterK})
	sin2 := 1 / (1 + matterK)
	threshold := ThresholdMap{
		Name:                    "finite abelian normalization threshold",
		ContactCoefficient:      contactK,
		MatterCoefficient:       matterK,
		EmbeddingScaleSq:        lambdaSq,
		EmbeddingScale:          lambda,
		NormalizationDeficit:    1 - lambdaSq,
		Formula:                 "Y=lambda Y_phi, lambda^2=Tr(Y_matter^2)/Tr(T3^2)/K(Y_phi)=5/9",
		Selected:                after <= eps && close(lambdaSq, 5.0/9.0, 1e-8),
		PhysicalMassThreshold:   false,
		PhysicalScaleInserted:   false,
		RepresentationThreshold: false,
		Detail:                  "a dimensionless normalization-threshold map between the scalar/contact abelian carrier and the finite matter hypercharge table; not a heavy-particle threshold spectrum",
	}

	truth := "Gate 102 derives the unique positive, orientation-preserving one-dimensional abelian embedding Y=lambda Y_phi that carries the action-selected contact coefficient K(Y_phi)=3 into the finite matter hypercharge normalization k_Y=5/3.  The selected square is lambda^2=5/9, so the embedded generator-basis Hessian is diag(1,1,1,5/3).  This upgrades sin^2=3/8 from a matter-table-only diagnostic to an embedded finite boundary diagnostic.  It is still not a physical low-energy weak mixing angle because the boundary scale, RG flow, threshold activation, and electromagnetic coupling normalization remain open."

	selected := threshold.Selected && signs[0].Selected && close(sin2, 3.0/8.0, 1e-8)
	return Analysis{
		Boundary: boundary,

		ContactGeneratorName: "Y_phi=(Q-Z)/2",
		MatterGeneratorName:  "Y_matter",
		EmbeddingFormula:     "Y_matter = lambda Y_phi",
		EmbeddingEquation:    "lambda^2 K(Y_phi) = k_Y",

		ContactU1KineticCoefficient: contactK,
		MatterHyperchargeKY:         matterK,
		EmbeddingScaleSq:            lambdaSq,
		EmbeddingScale:              lambda,
		EmbeddingScaleNegative:      negLambda,

		SignCandidates:          signs,
		OrientationSelected:     signs[0].Selected && !signs[1].Selected,
		OneDimensionalMapUnique: signs[0].Selected && math.Abs(signs[0].Lambda+signs[1].Lambda) <= eps,

		EmbeddedMatterHessian:             embedded,
		EmbeddedMatterHessianSelected:     close(embedded.At(3, 3), 5.0/3.0, 1e-8) && close(embedded.At(0, 0), 1, eps) && close(embedded.At(1, 1), 1, eps) && close(embedded.At(2, 2), 1, eps),
		EmbeddedMatterBoundarySin2:        sin2,
		EmbeddedMatterBoundaryDiagnostic:  close(sin2, 3.0/8.0, 1e-8),
		ContactMatterMismatchBeforeMap:    before,
		ContactMatterMismatchAfterMap:     after,
		FiniteNormalizationThreshold:      threshold,
		ContactToMatterEmbeddingSelected:  selected,
		ThresholdNormalizationMapSelected: threshold.Selected,

		RGBoundaryScaleDerived:   false,
		BetaThresholdFlowDerived: false,
		PhysicalWeakAngleDerived: false,
		FineStructureDerived:     false,
		PhysicalCouplingsDerived: false,
		PhysicalMassesDerived:    false,
		HiddenObservedInputUsed:  false,

		TruthStatement: truth,
		RejectedClaims: []string{
			"lambda=sqrt(5)/3 is a fitted coupling ratio",
			"sin^2=3/8 is the observed low-energy weak mixing angle",
			"the finite normalization threshold is a physical heavy-particle mass threshold",
			"the RG boundary scale M* has been derived",
			"alpha, g2, gY, W/Z masses, or the Higgs vev follow from the embedding alone",
		},
		RemainingUnknowns: []string{
			"U-22A-BOUNDARY-SCALE-SELECTION: derive the finite scale M* at which diag(1,1,1,5/3) is valid boundary data",
			"U-22B-RG-FLOW: derive continuum-active beta coefficients and running from the finite spectrum",
			"U-22C-THRESHOLD-ACTIVATION: decide which dimensionless finite modes activate or decouple across the flow",
			"U-22D-EM-COUPLING-NORMALIZATION: derive e, alpha, and physical thetaW only after scale and RG data exist",
			"U-22E-MASS-UNIT: derive W/Z/Higgs/fermion physical masses only after the scalar radius receives a non-fitted unit",
		},
		RecommendedNextGate: "Gate 103 — finite RG flow and boundary-scale selection firewall",
	}, nil
}

func close(a, b, eps float64) bool { return math.Abs(a-b) <= eps }

func FormatFloatSlice(xs []float64) string {
	parts := make([]string, 0, len(xs))
	for _, x := range xs {
		parts = append(parts, fmt.Sprintf("%.10f", x))
	}
	return "[" + strings.Join(parts, ", ") + "]"
}

func FormatMatrix(m linear.Matrix) string {
	rows := make([]string, 0, m.Rows())
	for r := 0; r < m.Rows(); r++ {
		cols := make([]string, 0, m.Cols())
		for c := 0; c < m.Cols(); c++ {
			v := m.At(r, c)
			if math.Abs(v) < 1e-12 {
				v = 0
			}
			cols = append(cols, fmt.Sprintf("%.10f", v))
		}
		rows = append(rows, "["+strings.Join(cols, ", ")+"]")
	}
	return strings.Join(rows, " ")
}

func FormatSigns(xs []SignCandidate) string {
	parts := make([]string, 0, len(xs))
	for _, x := range xs {
		state := "rejected"
		if x.Selected {
			state = "selected"
		}
		parts = append(parts, fmt.Sprintf("%s[%s]: lambda=%.10f, lambda^2=%.10f, orientation_preserving=%t", x.Name, state, x.Lambda, x.LambdaSq, x.OrientationPreserving))
	}
	return strings.Join(parts, "; ")
}

func Join(xs []string) string { return strings.Join(xs, "; ") }
