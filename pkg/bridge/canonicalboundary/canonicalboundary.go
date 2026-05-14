// Package canonicalboundary implements Gate 101: canonical finite RG/scale
// boundary seed from the Gate 100 variational Hessian.
//
// Gate 100 selected a full finite electroweak Hessian in the closed basis
// [T1,T2,Z,Q].  This package performs the next strictly allowed operation: it
// changes basis to the physical generator coordinates
// [T1,T2,T3=(Z+Q)/2,Y_phi=(Q-Z)/2], reads the action-selected dimensionless
// kinetic coefficients, and exposes the remaining bridge data required before
// any physical alpha, theta_W, W/Z mass, or RG running claim is allowed.
package canonicalboundary

import (
	"fmt"
	"math"
	"strings"
	"sync"

	"github.com/bagherbal/asha-engine/pkg/bridge/canonicalaction"
	"github.com/bagherbal/asha-engine/pkg/bridge/ewprojection"
	"github.com/bagherbal/asha-engine/pkg/linear"
)

type BoundaryCandidate struct {
	Name     string
	Value    float64
	Formula  string
	Physical bool
	Detail   string
}

type BasisChange struct {
	FromBasis []string
	ToBasis   []string
	// Columns are the ToBasis generators written in FromBasis coordinates.
	Matrix linear.Matrix
}

type Analysis struct {
	CanonicalAction canonicalaction.Analysis
	EWProjection    ewprojection.Analysis

	ClosedBasis      []string
	GeneratorBasis   []string
	BasisTransform   BasisChange
	GeneratorHessian linear.Matrix

	SU2KineticEntries     []float64
	SU2KineticIsotropic   bool
	SU2KineticCoefficient float64

	ScalarContactU1KineticCoefficient float64
	ScalarContactU1ToSU2Ratio         float64
	InverseKineticCouplingRatio       float64
	ScalarContactBoundarySin2         float64

	MatterHyperchargeKY      float64
	MatterBoundarySin2       float64
	ContactMatterMismatch    float64
	RequiredEmbeddingScaleSq float64
	RequiredEmbeddingScale   float64
	EmbeddingMapSelected     bool

	BoundaryCandidates []BoundaryCandidate

	ActionSelectedDimensionlessBoundarySeed bool
	GaugeKineticNormalizationDerived        bool
	RGFlowDetermined                        bool
	BoundaryScaleDerived                    bool
	PhysicalWeakAngleDerived                bool
	FineStructureDerived                    bool
	PhysicalMassesDerived                   bool
	HiddenObservedInputUsed                 bool

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
		ca, err := canonicalaction.BuildDefault()
		if err != nil {
			defaultErr = err
			return
		}
		ew, err := ewprojection.BuildDefault()
		if err != nil {
			defaultErr = err
			return
		}
		defaultValue, defaultErr = Build(ca, ew, 1e-10)
	})
	return defaultValue, defaultErr
}

func Build(ca canonicalaction.Analysis, ew ewprojection.Analysis, eps float64) (Analysis, error) {
	if eps <= 0 {
		eps = 1e-10
	}
	if !ca.CanonicalActionSelected || !ca.FullGaugeHessianSelected {
		return Analysis{}, fmt.Errorf("Gate 101 requires Gate 100 selected canonical action and full gauge Hessian")
	}
	if ca.FullGaugeHessian.Rows() != 4 || ca.FullGaugeHessian.Cols() != 4 {
		return Analysis{}, fmt.Errorf("Gate 101 expects a 4x4 electroweak Hessian")
	}
	if !ew.EqualNormalizedCouplingBoundaryCandidate || ew.HyperchargeNormalizationKY <= eps {
		return Analysis{}, fmt.Errorf("Gate 101 requires the finite matter hypercharge normalization candidate")
	}

	closed := []string{"T1", "T2", "Z=T3-Y_phi", "Q=T3+Y_phi"}
	physical := []string{"T1", "T2", "T3=(Z+Q)/2", "Y_phi=(Q-Z)/2"}
	transform := linear.NewMatrix(4, 4)
	transform.Set(0, 0, 1)
	transform.Set(1, 1, 1)
	transform.Set(2, 2, 0.5)
	transform.Set(3, 2, 0.5)
	transform.Set(2, 3, -0.5)
	transform.Set(3, 3, 0.5)

	kt, err := ca.FullGaugeHessian.Mul(transform)
	if err != nil {
		return Analysis{}, err
	}
	genH, err := transform.Transpose().Mul(kt)
	if err != nil {
		return Analysis{}, err
	}
	if !genH.IsSymmetric(100 * eps) {
		return Analysis{}, fmt.Errorf("generator-basis Hessian is not symmetric")
	}

	su2 := []float64{genH.At(0, 0), genH.At(1, 1), genH.At(2, 2)}
	su2Iso := closeAll(su2, su2[0], eps) && maxOffDiagonal(genH, 0, 3) <= eps
	su2Coeff := mean(su2)
	u1Coeff := genH.At(3, 3)
	ratio := u1Coeff / su2Coeff
	invRatio := su2Coeff / u1Coeff
	contactSin2 := invRatio / (1 + invRatio)

	kY := ew.HyperchargeNormalizationKY
	matterSin2 := ew.EqualNormalizedCouplingBoundarySin2
	mismatch := u1Coeff - kY
	requiredScaleSq := kY / u1Coeff
	requiredScale := math.Sqrt(requiredScaleSq)

	candidates := []BoundaryCandidate{
		{
			Name:     "action-selected SU(2)_L kinetic unit",
			Value:    su2Coeff,
			Formula:  "K(T1)=K(T2)=K(T3)=1 after T3=(Z+Q)/2",
			Physical: false,
			Detail:   "dimensionless finite Hessian normalization; not a measured g2",
		},
		{
			Name:     "action-selected scalar/contact U(1) coefficient",
			Value:    u1Coeff,
			Formula:  "K(Y_phi)=3 after Y_phi=(Q-Z)/2",
			Physical: false,
			Detail:   "finite scalar-contact abelian kinetic seed; not yet the matter hypercharge coupling",
		},
		{
			Name:     "contact-sector no-running angle diagnostic",
			Value:    contactSin2,
			Formula:  "sin²_contact = (1/K_Yphi)/(1/K_SU2+1/K_Yphi)=1/4",
			Physical: false,
			Detail:   "diagnostic only; it uses the finite contact U(1) seed before matter embedding, RG flow, and thresholds",
		},
		{
			Name:     "matter-table equal-normalized boundary diagnostic",
			Value:    matterSin2,
			Formula:  "sin²_matter = 1/(1+k_Y), k_Y=5/3, hence 3/8",
			Physical: false,
			Detail:   "representation-level hypercharge normalization from the finite one-generation table; not low-energy thetaW",
		},
		{
			Name:     "required contact-to-matter embedding scale squared",
			Value:    requiredScaleSq,
			Formula:  "lambda² = k_Y/K(Y_phi)=5/9 if Y=lambda·Y_phi is demanded",
			Physical: false,
			Detail:   "exposes the missing embedding theorem; the scale is not selected by Gate 101",
		},
	}

	truth := "Gate 101 transforms the Gate 100 Hessian from the closed carrier [T1,T2,Z,Q] to the generator basis [T1,T2,T3,Y_phi].  The selected Hessian becomes diag(1,1,1,3): SU(2)_L is isotropic, and the scalar/contact U(1) seed has coefficient 3.  This gives a dimensionless RG-boundary seed, not a physical electroweak prediction.  The finite matter table separately gives k_Y=5/3 and the usual equal-normalized diagnostic sin²=3/8.  The mismatch 3 versus 5/3 identifies the next missing theorem: a contact-to-matter hypercharge embedding/threshold map.  No observed alpha, thetaW, mass, or scale is used."

	return Analysis{
		CanonicalAction: ca,
		EWProjection:    ew,
		ClosedBasis:     closed,
		GeneratorBasis:  physical,
		BasisTransform: BasisChange{
			FromBasis: closed,
			ToBasis:   physical,
			Matrix:    transform,
		},
		GeneratorHessian: genH,

		SU2KineticEntries:                 su2,
		SU2KineticIsotropic:               su2Iso,
		SU2KineticCoefficient:             su2Coeff,
		ScalarContactU1KineticCoefficient: u1Coeff,
		ScalarContactU1ToSU2Ratio:         ratio,
		InverseKineticCouplingRatio:       invRatio,
		ScalarContactBoundarySin2:         contactSin2,

		MatterHyperchargeKY:      kY,
		MatterBoundarySin2:       matterSin2,
		ContactMatterMismatch:    mismatch,
		RequiredEmbeddingScaleSq: requiredScaleSq,
		RequiredEmbeddingScale:   requiredScale,
		EmbeddingMapSelected:     false,

		BoundaryCandidates: candidates,

		ActionSelectedDimensionlessBoundarySeed: su2Iso && close(u1Coeff, 3, eps) && close(contactSin2, 0.25, eps),
		GaugeKineticNormalizationDerived:        true,
		RGFlowDetermined:                        false,
		BoundaryScaleDerived:                    false,
		PhysicalWeakAngleDerived:                false,
		FineStructureDerived:                    false,
		PhysicalMassesDerived:                   false,
		HiddenObservedInputUsed:                 false,

		TruthStatement: truth,
		RejectedClaims: []string{
			"diag(1,1,1,3) is the observed electroweak coupling vector",
			"the contact-sector diagnostic sin²=1/4 is the physical weak mixing angle",
			"the matter-table diagnostic sin²=3/8 is the physical low-energy weak mixing angle",
			"the contact U(1) coefficient 3 is already the Standard Model hypercharge normalization 5/3",
			"alpha, W/Z masses, Higgs vev, or RG running follow from the finite Hessian alone",
		},
		RemainingUnknowns: []string{
			"U-21A-CONTACT-MATTER-HYPERCHARGE-EMBEDDING: derive or reject the finite map Y=lambda·Y_phi with lambda²=5/9, without fitting to observed couplings",
			"U-21B-BOUNDARY-SCALE: derive the scale M* at which the dimensionless Hessian is to be interpreted as boundary data",
			"U-21C-BETA-THRESHOLD-MAP: derive beta coefficients and threshold activations from the finite spectrum rather than importing them",
			"U-21D-PHYSICAL-COUPLINGS: only after embedding, scale, and RG data exist may thetaW, alpha, g2, and gY be computed",
			"U-21E-MASS-UNIT: W/Z and Higgs masses still require the scalar radius-to-physical-unit bridge",
		},
		RecommendedNextGate: "Gate 102 — contact-to-matter hypercharge embedding / threshold-map theorem",
	}, nil
}

func close(a, b, eps float64) bool { return math.Abs(a-b) <= eps }

func closeAll(xs []float64, target float64, eps float64) bool {
	for _, x := range xs {
		if !close(x, target, eps) {
			return false
		}
	}
	return true
}

// maxOffDiagonal tests the SU(2) 3x3 block only.
func maxOffDiagonal(m linear.Matrix, start, size int) float64 {
	max := 0.0
	for r := start; r < start+size; r++ {
		for c := start; c < start+size; c++ {
			if r == c {
				continue
			}
			if v := math.Abs(m.At(r, c)); v > max {
				max = v
			}
		}
	}
	return max
}

func mean(xs []float64) float64 {
	sum := 0.0
	for _, x := range xs {
		sum += x
	}
	return sum / float64(len(xs))
}

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

func FormatCandidates(xs []BoundaryCandidate) string {
	parts := make([]string, 0, len(xs))
	for _, x := range xs {
		state := "diagnostic"
		if x.Physical {
			state = "physical"
		}
		parts = append(parts, fmt.Sprintf("%s[%s]=%.10f via %s", x.Name, state, x.Value, x.Formula))
	}
	return strings.Join(parts, "; ")
}

func Join(xs []string) string { return strings.Join(xs, "; ") }
