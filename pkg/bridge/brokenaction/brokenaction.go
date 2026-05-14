// Package brokenaction implements Gate 95: broken-sector action second
// variation / kinetic Hessian search.
//
// Gate 94 identified K_broken_raw = diag(1,1,4) as the unique raw-coordinate
// metric that whitens the quotient-safe broken-generator image metric.  Gate 95
// asks the stricter variational question: can this Hessian be obtained as
// δ²S/δA_iδA_j from a finite scalar/gauge action, rather than as a diagnostic
// basis normalization?
package brokenaction

import (
	"fmt"
	"math"
	"strings"
	"sync"

	"github.com/bagherbal/asha-engine/pkg/bridge/gaugekineticdiag"
)

type ActionSlot struct {
	Name    string
	Present bool
	Detail  string
}

type CandidateSecondVariation struct {
	Name       string
	Diagonal   []float64
	Positive   bool
	Selected   bool
	Derivation string
}

type Analysis struct {
	Gate94 gaugekineticdiag.Analysis

	BrokenDirections  []string
	CandidateK        []float64
	CandidateDet      float64
	CandidatePositive bool
	WhiteningExact    bool

	ActionSlots []ActionSlot
	Candidates  []CandidateSecondVariation

	ScalarGaugeVariablesTyped  bool
	ScalarKineticActionDerived bool
	GaugeFieldVariablesDerived bool
	CurvatureTermDerived       bool
	CovariantDerivativeDerived bool
	SecondVariationComputed    bool
	Diag114SelectedByAction    bool
	GaugeHessianSelected       bool
	CouplingRatioDerived       bool
	PhysicalMassesDerived      bool

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
		g94, err := gaugekineticdiag.BuildDefault()
		if err != nil {
			defaultErr = err
			return
		}
		defaultValue, defaultErr = Build(g94, 1e-10)
	})
	return defaultValue, defaultErr
}

func Build(g94 gaugekineticdiag.Analysis, eps float64) (Analysis, error) {
	if eps <= 0 {
		eps = 1e-10
	}
	if len(g94.CandidateDiagonal) != 3 || len(g94.WhitenedMetricDiagonal) != 3 {
		return Analysis{}, fmt.Errorf("Gate 95 requires Gate 94 broken 3D candidate data")
	}
	k := append([]float64(nil), g94.CandidateDiagonal...)
	det := k[0] * k[1] * k[2]
	positive := k[0] > eps && k[1] > eps && k[2] > eps
	whiteCond := g94.WhitenedCondition
	whiteningExact := math.Abs(whiteCond-1) < 1e-8

	actionSlots := []ActionSlot{
		{Name: "scalar active frame Φ", Present: true, Detail: "4-real-dimensional scalar/contact active frame is constructed"},
		{Name: "broken generator image map X_a φ0", Present: true, Detail: "rank-3 broken image diagnostic is available"},
		{Name: "diagnostic scalar metric I4", Present: true, Detail: "Euclidean active-frame metric is available as a computational metric"},
		{Name: "finite scalar kinetic action", Present: false, Detail: "not derived; K_Φ=I4 is not yet an action-selected kinetic term"},
		{Name: "finite gauge field variables", Present: false, Detail: "broken-coordinate fields exist diagnostically but not as finite action variables"},
		{Name: "finite curvature/field-strength term", Present: false, Detail: "no F_A F_A term has been derived for the broken gauge fields"},
		{Name: "action second variation", Present: false, Detail: "δ²S/δA_iδA_j is not computed from a finite action"},
	}

	candidates := []CandidateSecondVariation{
		{Name: "metric-whitening Hessian", Diagonal: k, Positive: positive, Selected: false, Derivation: "selected by isotropizing broken-generator images, not by varying an action"},
		{Name: "raw image-metric Hessian", Diagonal: g94.RawMetricDiagonal, Positive: allPositive(g94.RawMetricDiagonal, eps), Selected: false, Derivation: "would follow from a naive ||Aφ0||² diagnostic with unit scalar/gauge metric"},
		{Name: "unit broken-coordinate Hessian", Diagonal: []float64{1, 1, 1}, Positive: true, Selected: false, Derivation: "basis convention only; ignores neutral/charged image metric"},
		{Name: "general positive diagonal Hessian", Diagonal: nil, Positive: true, Selected: false, Derivation: "three positive parameters remain unless the finite action fixes them"},
	}

	truth := "Gate 95 tests whether K_broken_raw=diag(1,1,4) is produced by a finite action second variation. The answer remains open: diag(1,1,4) is a mathematically precise metric-whitening candidate, but the engine still lacks finite gauge-field variables, a scalar kinetic action, a curvature/field-strength term, and an actual δ²S computation. Therefore the candidate is not yet a physical gauge kinetic Hessian."

	return Analysis{
		Gate94:                     g94,
		BrokenDirections:           []string{"T1", "T2", "Z_raw=T3-Y_phi"},
		CandidateK:                 k,
		CandidateDet:               det,
		CandidatePositive:          positive,
		WhiteningExact:             whiteningExact,
		ActionSlots:                actionSlots,
		Candidates:                 candidates,
		ScalarGaugeVariablesTyped:  true,
		ScalarKineticActionDerived: false,
		GaugeFieldVariablesDerived: false,
		CurvatureTermDerived:       false,
		CovariantDerivativeDerived: true,
		SecondVariationComputed:    false,
		Diag114SelectedByAction:    false,
		GaugeHessianSelected:       false,
		CouplingRatioDerived:       false,
		PhysicalMassesDerived:      false,
		TruthStatement:             truth,
		RejectedClaims: []string{
			"diag(1,1,4) is derived from δ²S without constructing S",
			"metric whitening is equivalent to physical gauge kinetic normalization",
			"the W/Z diagnostic mass matrix is a physical mass prediction",
			"the neutral factor 1/2 fixes theta_W or alpha by itself",
		},
		RemainingUnknowns: []string{
			"U-18C8D-FINITE-BROKEN-GAUGE-VARIABLES: define broken gauge fields as finite action variables",
			"U-18C6-SCALAR-KINETIC-ACTION: derive K_phi from scalar/contact dynamics",
			"U-18C8E-FIELD-STRENGTH: derive a finite F_A F_A or BF/Plebanski kinetic term for broken directions",
			"U-18C8F-SECOND-VARIATION: compute δ²S/δA_iδA_j and compare to diag(1,1,4)",
			"U-18C9-PHYSICAL-COUPLINGS: connect selected Hessian to g2, gY, thetaW, and RG flow",
		},
		RecommendedNextGate: "Gate 96 — Finite Broken Gauge Field Variables / Curvature Term Search",
	}, nil
}

func allPositive(xs []float64, eps float64) bool {
	if len(xs) == 0 {
		return false
	}
	for _, x := range xs {
		if x <= eps {
			return false
		}
	}
	return true
}

func Join(xs []string) string { return strings.Join(xs, "; ") }
