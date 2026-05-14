// Package gaugehessian implements Gate 83: gauge kinetic Hessian from finite
// action second variation.
//
// Gate 82 showed that several positive U(1) kinetic diagnostics exist, but none
// is selected as the physical gauge kinetic Hessian. This gate makes the next
// mathematical requirement explicit: a real finite action must contain U(1)
// gauge-field variables and a quadratic term whose second variation produces
// the Hessian.
//
// The current engine has charge tables, representation metrics, anomaly
// cancellations, and scalar/contact structure. It does not yet have a native
// finite U(1) gauge-field action with a computable second variation. Therefore
// this gate exposes the Hessian search space and rejects promotion of any
// diagnostic metric to physical coupling data.
package gaugehessian

import (
	"fmt"
	"sync"

	"github.com/bagherbal/asha-engine/pkg/bridge/gaugeaction"
)

type ActionSlot struct {
	Name    string
	Present bool
	Derived bool
	Detail  string
}

type HessianCandidate struct {
	Name                string
	Positive            bool
	FromSecondVariation bool
	FreeParameters      int
	Detail              string
}

type Analysis struct {
	GaugeAction gaugeaction.Analysis

	Fields              []string
	SymmetricHessianDim int
	DiagonalNoMixingDim int
	OffDiagonalDim      int

	ActionSlots []ActionSlot
	Candidates  []HessianCandidate

	FiniteActionVariablesTyped bool
	U1GaugeFieldActionPresent  bool
	SecondVariationComputed    bool
	HessianSelected            bool
	BoundaryCouplingFixed      bool
	PhysicalAlphaDerived       bool
	HiddenObservedInputUsed    bool

	TruthStatement      string
	RecommendedNextGate string
	RemainingUnknowns   []string
}

var (
	defaultOnce  sync.Once
	defaultValue Analysis
	defaultErr   error
)

func BuildDefault() (Analysis, error) {
	defaultOnce.Do(func() {
		ga, err := gaugeaction.BuildDefault()
		if err != nil {
			defaultErr = err
			return
		}
		defaultValue, defaultErr = Build(ga)
	})
	return defaultValue, defaultErr
}

func Build(ga gaugeaction.Analysis) (Analysis, error) {
	if ga.GaugeKineticActionSelected {
		return Analysis{}, fmt.Errorf("Gate 83 expects Gate 82 to leave the gauge kinetic action unselected")
	}

	fields := []string{"central-u1", "B-L", "contact-u1"}
	slots := []ActionSlot{
		{Name: "matter charge representation", Present: true, Derived: true, Detail: "finite charge tables and trace-Gram representation metrics are present"},
		{Name: "contact scalar representation", Present: true, Derived: true, Detail: "contact-u1/T_phi scalar-side representation metric is present"},
		{Name: "anomaly/no-mixing ledger", Present: true, Derived: true, Detail: "known off-diagonal source candidates cancel"},
		{Name: "finite U(1) gauge-field variables", Present: false, Derived: false, Detail: "typed A_c, A_BL, A_phi gauge fields are not yet part of a finite action"},
		{Name: "finite gauge curvature term", Present: false, Derived: false, Detail: "no F_A^2 or BF/Plebanski U(1) kinetic term has been constructed"},
		{Name: "scalar covariant derivative", Present: false, Derived: false, Detail: "D_phi is not yet derived on the finite scalar/contact frame"},
		{Name: "second variation operator", Present: false, Derived: false, Detail: "δ²S/δA_iδA_j cannot be computed before the action exists"},
	}

	candidates := []HessianCandidate{
		{Name: "trace-Gram Hessian", Positive: true, FromSecondVariation: false, FreeParameters: 0, Detail: "representation metric; not an action Hessian"},
		{Name: "inverse trace-Gram Hessian", Positive: true, FromSecondVariation: false, FreeParameters: 0, Detail: "diagnostic inverse metric; not selected by variation"},
		{Name: "unit Hessian", Positive: true, FromSecondVariation: false, FreeParameters: 0, Detail: "normalization convention; erases representation metric"},
		{Name: "general anomaly-constrained diagonal Hessian", Positive: true, FromSecondVariation: false, FreeParameters: 3, Detail: "K=diag(k_c,k_BL,k_phi), k_i>0; most honest open family"},
		{Name: "general symmetric U(1) Hessian", Positive: true, FromSecondVariation: false, FreeParameters: 6, Detail: "allowed algebraically, but known off-diagonal sources currently cancel"},
	}

	truth := "Gate 83 shows that the U(1) gauge kinetic Hessian cannot yet be derived from an action second variation. The engine has representation metrics and anomaly/no-mixing diagnostics, but not finite U(1) gauge-field variables, a finite curvature/kinetic term, or a scalar covariant derivative. Therefore all U(1) couplings remain boundary-family data rather than physical constants."

	return Analysis{
		GaugeAction:                ga,
		Fields:                     fields,
		SymmetricHessianDim:        6,
		DiagonalNoMixingDim:        3,
		OffDiagonalDim:             3,
		ActionSlots:                slots,
		Candidates:                 candidates,
		FiniteActionVariablesTyped: false,
		U1GaugeFieldActionPresent:  false,
		SecondVariationComputed:    false,
		HessianSelected:            false,
		BoundaryCouplingFixed:      false,
		PhysicalAlphaDerived:       false,
		HiddenObservedInputUsed:    false,
		TruthStatement:             truth,
		RecommendedNextGate:        "Gate 84 — Finite Scalar Covariant Derivative / Gauge-Boson Mass Matrix Search",
		RemainingUnknowns: []string{
			"U-20D3F5-ACTION-SELECTS-U1-HESSIAN: derive the U(1) kinetic Hessian as δ²S from a finite action",
			"U-18C-COVARIANT-DERIVATIVE: construct DΦ on the scalar/contact frame",
			"U-18D-GAUGE-MASS-MATRIX: compute the dimensionless W/Z/photon mass matrix from DΦ",
			"U-20D3F6-BOUNDARY-COUPLING: fix g_* only after the Hessian and action normalization are selected",
		},
	}, nil
}

func DerivedActionSlotCount(xs []ActionSlot) int {
	n := 0
	for _, x := range xs {
		if x.Derived {
			n++
		}
	}
	return n
}

func SelectedCandidateCount(xs []HessianCandidate) int {
	n := 0
	for _, x := range xs {
		if x.FromSecondVariation {
			n++
		}
	}
	return n
}
