// Package brokengaugefields implements Gate 96: finite broken gauge-field
// variables / curvature term search.
//
// Gate 95 established K_broken_raw = diag(1,1,4) as a metric-whitening
// candidate, but it could not be selected by an action because finite broken
// gauge-field variables and a curvature/field-strength term were still absent.
//
// Gate 96 therefore types the broken gauge variables and audits their closure.
// The key result is structural: the broken directions {T1,T2,Z=T3-Y} are not a
// closed Lie algebra by themselves because [T1,T2] has an electromagnetic Q
// component.  A finite curvature term must therefore be written for the full
// electroweak connection, not for the broken sector alone.
package brokengaugefields

import (
	"fmt"
	"math"
	"strings"
	"sync"

	"github.com/bagherbal/asha-engine/pkg/bridge/brokenaction"
)

type GaugeVariable struct {
	Name       string
	Sector     string
	Broken     bool
	Normalized bool
}

type BracketDecomposition struct {
	Left           string
	Right          string
	Components     map[string]float64
	ClosedInBroken bool
}

type CurvatureSlot struct {
	Name    string
	Present bool
	Detail  string
}

type Analysis struct {
	Gate95 brokenaction.Analysis

	BrokenVariables []GaugeVariable
	FullVariables   []GaugeVariable
	BrokenDimension int
	FullDimension   int

	CandidateHessian []float64
	CandidateDet     float64

	Brackets         []BracketDecomposition
	BrokenOnlyClosed bool
	FullEWClosed     bool
	RequiresPhoton   bool

	CurvatureSlots []CurvatureSlot

	BrokenFieldsTyped       bool
	FullEWFieldsRequired    bool
	FieldStrengthTermTyped  bool
	CurvatureTermDerived    bool
	FiniteActionVariables   bool
	SecondVariationPossible bool
	Diag114ActionSelected   bool
	PhysicalCouplings       bool

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
		g95, err := brokenaction.BuildDefault()
		if err != nil {
			defaultErr = err
			return
		}
		defaultValue, defaultErr = Build(g95)
	})
	return defaultValue, defaultErr
}

func Build(g95 brokenaction.Analysis) (Analysis, error) {
	if len(g95.CandidateK) != 3 {
		return Analysis{}, fmt.Errorf("Gate 96 requires Gate 95 broken-sector candidate Hessian")
	}
	k := append([]float64(nil), g95.CandidateK...)
	det := k[0] * k[1] * k[2]
	if math.Abs(det-4) > 1e-9 {
		return Analysis{}, fmt.Errorf("unexpected Gate 95 Hessian determinant %.12f", det)
	}

	broken := []GaugeVariable{
		{Name: "W1", Sector: "T1", Broken: true, Normalized: true},
		{Name: "W2", Sector: "T2", Broken: true, Normalized: true},
		{Name: "Z_raw", Sector: "Z=T3-Y_phi", Broken: true, Normalized: false},
	}
	full := []GaugeVariable{
		broken[0], broken[1], broken[2],
		{Name: "A_em", Sector: "Q=T3+Y_phi", Broken: false, Normalized: false},
	}

	brackets := []BracketDecomposition{
		{Left: "T1", Right: "T2", Components: map[string]float64{"Z": 0.5, "Q": 0.5}, ClosedInBroken: false},
		{Left: "T2", Right: "Z", Components: map[string]float64{"T1": 1.0}, ClosedInBroken: true},
		{Left: "Z", Right: "T1", Components: map[string]float64{"T2": 1.0}, ClosedInBroken: true},
	}
	brokenOnlyClosed := true
	requiresPhoton := false
	for _, b := range brackets {
		if !b.ClosedInBroken {
			brokenOnlyClosed = false
		}
		if _, ok := b.Components["Q"]; ok {
			requiresPhoton = true
		}
	}

	slots := []CurvatureSlot{
		{Name: "broken gauge variables", Present: true, Detail: "W1, W2, and Z_raw are typed as diagnostic broken-coordinate fields"},
		{Name: "unbroken electromagnetic variable", Present: true, Detail: "A_em is required by Lie closure of the full electroweak connection"},
		{Name: "broken-only curvature", Present: false, Detail: "not closed: [T1,T2] contains an electromagnetic Q component"},
		{Name: "full electroweak curvature variables", Present: true, Detail: "{T1,T2,Z,Q} is the minimal closure carrier in the broken basis"},
		{Name: "finite field-strength action", Present: false, Detail: "no F_A F_A or BF/Plebanski kinetic term is derived for these variables"},
		{Name: "second variation", Present: false, Detail: "cannot compute δ²S until the finite curvature/action term exists"},
	}

	truth := "Gate 96 types the broken gauge fields but finds that the broken sector alone is not a closed curvature carrier.  The bracket [T1,T2] contains both Z and the electromagnetic Q direction, so any finite curvature term must use the full electroweak connection {T1,T2,Z,Q}.  This is progress from untyped diagnostics to typed variables, but diag(1,1,4) remains a candidate until a full finite field-strength action and second variation are derived."

	return Analysis{
		Gate95:                  g95,
		BrokenVariables:         broken,
		FullVariables:           full,
		BrokenDimension:         len(broken),
		FullDimension:           len(full),
		CandidateHessian:        k,
		CandidateDet:            det,
		Brackets:                brackets,
		BrokenOnlyClosed:        brokenOnlyClosed,
		FullEWClosed:            true,
		RequiresPhoton:          requiresPhoton,
		CurvatureSlots:          slots,
		BrokenFieldsTyped:       true,
		FullEWFieldsRequired:    true,
		FieldStrengthTermTyped:  true,
		CurvatureTermDerived:    false,
		FiniteActionVariables:   false,
		SecondVariationPossible: false,
		Diag114ActionSelected:   false,
		PhysicalCouplings:       false,
		TruthStatement:          truth,
		RejectedClaims: []string{
			"the broken sector alone defines a closed finite gauge algebra",
			"diag(1,1,4) is action-selected before a finite curvature term exists",
			"the W/Z/photon diagnostic is already a physical mass theorem",
			"the electromagnetic direction can be discarded when constructing broken curvature",
		},
		RemainingUnknowns: []string{
			"U-18C8E-FULL-EW-CURVATURE: build F_A for the full {T1,T2,Z,Q} connection",
			"U-18C8G-FINITE-FIELD-STRENGTH-ACTION: derive F_A F_A or BF/Plebanski kinetic action",
			"U-18C8F-SECOND-VARIATION: compute δ²S/δA_iδA_j and compare to diag(1,1,4)",
			"U-18C9-COUPLINGS: derive g2/gY/thetaW/alpha only after the Hessian is selected",
		},
		RecommendedNextGate: "Gate 97 — Full Electroweak Connection Curvature / Field-Strength Audit",
	}, nil
}

func Join(xs []string) string { return strings.Join(xs, "; ") }
