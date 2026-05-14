// Package ewcurvature implements Gate 97: full electroweak connection
// curvature / field-strength audit.
//
// Gate 96 showed that the broken variables {T1,T2,Z=T3-Y_phi} are not closed:
// [T1,T2] contains both Z and the electromagnetic Q direction.  Gate 97 keeps
// the full connection {T1,T2,Z,Q}, builds its structure constants, audits the
// curvature carrier, and tests whether the curvature quadratic form selects the
// earlier diag(1,1,4) broken-sector Hessian candidate.
//
// The truth is subtle: the full connection closes, so a finite field-strength
// carrier exists at the Lie-algebra level.  But the adjoint/Killing diagnostic
// only sees the semisimple SU(2) direction T3=(Z+Q)/2 and leaves the abelian
// U(1) direction Y_phi=(Q-Z)/2 null.  Therefore the curvature algebra alone does
// not select the broken-sector Hessian diag(1,1,4), nor the physical U(1)
// kinetic normalization.
package ewcurvature

import (
	"fmt"
	"math"
	"strings"
	"sync"

	"github.com/bagherbal/asha-engine/pkg/bridge/brokengaugefields"
)

type FieldVariable struct {
	Name   string
	Basis  string
	Broken bool
}

type Bracket struct {
	Left       string
	Right      string
	Components map[string]float64
}

type MatrixSummary struct {
	Name        string
	Rows        int
	Cols        int
	Rank        int
	Trace       float64
	Determinant float64
	Detail      string
}

type Analysis struct {
	Gate96 brokengaugefields.Analysis

	Variables []FieldVariable
	Dimension int

	Brackets        []Bracket
	ClosureResidual float64
	Closed          bool

	AdjointRank         int
	AdjointMetricTrace  float64
	AdjointMetric       [][]float64
	AbelianNullVector   []float64
	SemisimpleDirection []float64

	CurvatureCarrierTyped       bool
	FullFieldStrengthTyped      bool
	CurvatureQuadraticCandidate bool
	AdjointMetricPositive       bool
	U1KineticSelected           bool
	Diag114SelectedByCurvature  bool
	SecondVariationComputed     bool
	PhysicalCouplings           bool

	MatrixSummaries []MatrixSummary

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
		g96, err := brokengaugefields.BuildDefault()
		if err != nil {
			defaultErr = err
			return
		}
		defaultValue, defaultErr = Build(g96)
	})
	return defaultValue, defaultErr
}

func Build(g96 brokengaugefields.Analysis) (Analysis, error) {
	if !g96.RequiresPhoton || !g96.FullEWClosed {
		return Analysis{}, fmt.Errorf("Gate 97 requires Gate 96 full electroweak closure variables")
	}
	vars := []FieldVariable{
		{Name: "W1", Basis: "T1", Broken: true},
		{Name: "W2", Basis: "T2", Broken: true},
		{Name: "Z_raw", Basis: "Z=T3-Y_phi", Broken: true},
		{Name: "A_em", Basis: "Q=T3+Y_phi", Broken: false},
	}
	brackets := []Bracket{
		{Left: "T1", Right: "T2", Components: map[string]float64{"Z": 0.5, "Q": 0.5}},
		{Left: "T2", Right: "Z", Components: map[string]float64{"T1": 1}},
		{Left: "Z", Right: "T1", Components: map[string]float64{"T2": 1}},
		{Left: "T2", Right: "Q", Components: map[string]float64{"T1": 1}},
		{Left: "Q", Right: "T1", Components: map[string]float64{"T2": 1}},
		{Left: "Z", Right: "Q", Components: map[string]float64{}},
	}

	// Adjoint/Killing diagnostic in the basis [T1,T2,Z,Q].  It is the matrix of
	// Tr(ad_X ad_Y) for the above structure constants.  The sign is conventional;
	// only rank/nullspace/neutral mixing are used as invariant diagnostics here.
	adj := [][]float64{
		{-2, 0, 0, 0},
		{0, -2, 0, 0},
		{0, 0, -2, -2},
		{0, 0, -2, -2},
	}
	trace := adj[0][0] + adj[1][1] + adj[2][2] + adj[3][3]
	// The neutral block [[-2,-2],[-2,-2]] has rank one; the full metric has rank
	// three.  Its null vector is Q-Z, i.e. the pure abelian Y_phi direction.
	rank := 3
	null := []float64{0, 0, -1, 1}
	semisimple := []float64{0, 0, 0.5, 0.5} // T3=(Z+Q)/2

	// diag(1,1,4) cannot be read from the adjoint diagnostic: the adjoint metric
	// contains off-diagonal Z/Q mixing and a null abelian direction, not a positive
	// broken-coordinate Hessian.
	diag114Selected := false
	positive := false
	secondVariation := false

	truth := "Gate 97 upgrades Gate 96 from typed variables to a closed full electroweak connection curvature carrier.  The full basis {T1,T2,Z,Q} closes and supports a formal field-strength audit, but its adjoint/Killing diagnostic is rank three with the pure abelian Y_phi=(Q-Z)/2 direction null.  The semisimple curvature sees T3=(Z+Q)/2, not a positive physical U(1) kinetic Hessian.  Therefore the full curvature algebra does not by itself select diag(1,1,4), g2, gY, thetaW, alpha, or physical W/Z masses."

	return Analysis{
		Gate96:                      g96,
		Variables:                   vars,
		Dimension:                   len(vars),
		Brackets:                    brackets,
		ClosureResidual:             0,
		Closed:                      true,
		AdjointRank:                 rank,
		AdjointMetricTrace:          trace,
		AdjointMetric:               adj,
		AbelianNullVector:           null,
		SemisimpleDirection:         semisimple,
		CurvatureCarrierTyped:       true,
		FullFieldStrengthTyped:      true,
		CurvatureQuadraticCandidate: true,
		AdjointMetricPositive:       positive,
		U1KineticSelected:           false,
		Diag114SelectedByCurvature:  diag114Selected,
		SecondVariationComputed:     secondVariation,
		PhysicalCouplings:           false,
		MatrixSummaries: []MatrixSummary{
			{Name: "adjoint/Killing diagnostic", Rows: 4, Cols: 4, Rank: rank, Trace: trace, Determinant: 0, Detail: "rank-3 semisimple diagnostic; null direction Q-Z is pure abelian Y_phi"},
			{Name: "broken Hessian candidate from Gate 94", Rows: 3, Cols: 3, Rank: 3, Trace: 6, Determinant: 4, Detail: "diag(1,1,4) remains metric-whitening candidate, not curvature-selected"},
		},
		TruthStatement: truth,
		RejectedClaims: []string{
			"the broken-only connection is sufficient for curvature",
			"the full electroweak adjoint metric is a positive physical gauge Hessian",
			"diag(1,1,4) is selected by curvature closure alone",
			"the abelian U(1) kinetic normalization follows from the semisimple adjoint metric",
			"g2/gY/thetaW/alpha or W/Z masses are derived at this gate",
		},
		RemainingUnknowns: []string{
			"U-18C8G-FINITE-FIELD-STRENGTH-ACTION: derive an action functional from the full curvature carrier",
			"U-18C8F-SECOND-VARIATION: compute δ²S/δA_iδA_j from that action",
			"U-18C10-U1-KINETIC: supply the abelian field-strength/kinetic term missing from the adjoint metric",
			"U-18C11-SCALAR-GAUGE-ACTION: combine scalar covariant derivative and full field-strength action",
		},
		RecommendedNextGate: "Gate 98 — Full Electroweak Quadratic Action / Abelian Completion Search",
	}, nil
}

func FormatMatrix(m [][]float64) string {
	rows := make([]string, 0, len(m))
	for _, r := range m {
		parts := make([]string, 0, len(r))
		for _, x := range r {
			if math.Abs(x) < 1e-12 {
				x = 0
			}
			parts = append(parts, fmt.Sprintf("%.3g", x))
		}
		rows = append(rows, "["+strings.Join(parts, ",")+"]")
	}
	return strings.Join(rows, " ")
}

func Join(xs []string) string { return strings.Join(xs, "; ") }
