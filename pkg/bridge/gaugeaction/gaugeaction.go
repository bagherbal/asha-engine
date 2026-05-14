// Package gaugeaction implements Gate 82: gauge kinetic action selection / RG
// boundary coupling audit.
//
// Gate 81 showed that the surviving diagonal U(1) trace-Gram data is a valid
// representation metric, but not automatically the physical gauge kinetic
// Hessian.  This gate tests the next question: can a finite action select that
// metric as the Hessian and thereby fix a boundary coupling?
//
// The result is deliberately conservative.  The code exposes several internally
// consistent diagnostic action choices, including the diagonal trace-Gram
// candidate, but none is selected by a finite variational theorem.  Therefore
// charge-level k_Y=5/3 and sin^2(theta)=3/8 remain valid boundary-geometry
// candidates, while physical g_Y, alpha_em, and RG boundary data remain open.
package gaugeaction

import (
	"fmt"
	"sync"

	"github.com/bagherbal/asha-engine/pkg/bridge/abeliancoupling"
)

type CandidateAction struct {
	Name       string
	MatrixRole string
	Positive   bool
	Selected   bool
	Reason     string
}

type DiagnosticCoupling struct {
	Name    string
	Norm2   float64
	Inverse float64
	Meaning string
}

type Analysis struct {
	AbelianCoupling abeliancoupling.Analysis

	Fields              []string
	SymmetricHessianDim int
	DiagonalHessianDim  int

	CandidateActions []CandidateAction
	Diagnostics      []DiagnosticCoupling

	TraceGramAsRepresentationMetric             bool
	TraceGramPromotedToGaugeKineticByAssumption bool
	GaugeKineticActionSelected                  bool

	ChargeTableKY                       float64
	ChargeTableSin2Boundary             float64
	TwoCarrierHyperchargeNormDiagnostic float64
	TwoCarrierInverseNormDiagnostic     float64

	BoundaryCouplingFamilyExposed bool
	BoundaryCouplingFixed         bool
	PhysicalU1CouplingDerived     bool
	FineStructureDerived          bool
	RGBoundaryScaleDerived        bool
	HiddenObservedInputUsed       bool

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
		ac, err := abeliancoupling.BuildDefault()
		if err != nil {
			defaultErr = err
			return
		}
		defaultValue, defaultErr = Build(ac)
	})
	return defaultValue, defaultErr
}

func Build(ac abeliancoupling.Analysis) (Analysis, error) {
	if !ac.DiagonalTraceGramSelectedAsRepresentationMetric {
		return Analysis{}, fmt.Errorf("Gate 82 requires Gate 81 representation-metric diagnostics")
	}
	if ac.PhysicalGaugeCouplingsDerived || ac.FineStructureDerived {
		return Analysis{}, fmt.Errorf("Gate 82 expects physical U(1) coupling and alpha to remain open")
	}

	diagnostics := make([]DiagnosticCoupling, 0, len(ac.Fields))
	fields := make([]string, 0, len(ac.Fields))
	for _, f := range ac.Fields {
		fields = append(fields, f.Name)
		diagnostics = append(diagnostics, DiagnosticCoupling{
			Name:    f.Name,
			Norm2:   f.TraceNormSquared,
			Inverse: f.InverseNormSquared,
			Meaning: "canonical-generator diagnostic only; not an action-selected gauge coupling",
		})
	}

	candidates := []CandidateAction{
		{
			Name:       "diagonal trace-Gram Hessian",
			MatrixRole: "K = diag(Tr(T_A^2)) on central, B-L, contact-u1",
			Positive:   true,
			Selected:   false,
			Reason:     "available representation metric, but no finite action has selected it as the gauge-field Hessian",
		},
		{
			Name:       "inverse trace-Gram Hessian",
			MatrixRole: "K = diag(1/Tr(T_A^2))",
			Positive:   true,
			Selected:   false,
			Reason:     "valid diagnostic but reverses sector dominance; not variationally selected",
		},
		{
			Name:       "unit abelian Hessian",
			MatrixRole: "K = I_3",
			Positive:   true,
			Selected:   false,
			Reason:     "erases representation metrics; no finite theorem selects unit weighting",
		},
		{
			Name:       "anomaly-constrained diagonal family",
			MatrixRole: "K = diag(k_c, k_BL, k_phi), k_A>0",
			Positive:   true,
			Selected:   false,
			Reason:     "most general no-mixing positive family still has three free kinetic coefficients",
		},
	}

	hyperNorm := ac.Hypercharge.CombinedBridgeNorm
	invHyper := 0.0
	if hyperNorm > 0 {
		invHyper = 1.0 / hyperNorm
	}

	truth := "Gate 82 shows that the diagonal U(1) trace-Gram data can be used as a representation-metric diagnostic and can parameterize a family of possible gauge kinetic actions. It is not selected by the finite action. Therefore k_Y=5/3 and sin²θ=3/8 remain charge-geometry boundary candidates, but g_Y, alpha_em, and RG boundary coupling remain undetermined."

	return Analysis{
		AbelianCoupling:                 ac,
		Fields:                          fields,
		SymmetricHessianDim:             6,
		DiagonalHessianDim:              3,
		CandidateActions:                candidates,
		Diagnostics:                     diagnostics,
		TraceGramAsRepresentationMetric: true,
		TraceGramPromotedToGaugeKineticByAssumption: false,
		GaugeKineticActionSelected:                  false,
		ChargeTableKY:                               ac.Hypercharge.ChargeTableKY,
		ChargeTableSin2Boundary:                     ac.Hypercharge.BoundarySin2,
		TwoCarrierHyperchargeNormDiagnostic:         hyperNorm,
		TwoCarrierInverseNormDiagnostic:             invHyper,
		BoundaryCouplingFamilyExposed:               true,
		BoundaryCouplingFixed:                       false,
		PhysicalU1CouplingDerived:                   false,
		FineStructureDerived:                        false,
		RGBoundaryScaleDerived:                      false,
		HiddenObservedInputUsed:                     false,
		TruthStatement:                              truth,
		RecommendedNextGate:                         "Gate 83 — Gauge Kinetic Hessian from Finite Action Second Variation",
		RemainingUnknowns: []string{
			"U-20D3F5-ACTION-SELECTS-U1-HESSIAN: derive the U(1) gauge kinetic Hessian from an action second variation",
			"U-20D3F6-BOUNDARY-COUPLING: derive the boundary coupling g_* instead of exposing only coupling families",
			"U-20D3F7-RG-BOUNDARY: derive a boundary scale and matching prescription before alpha_em can be computed",
			"U-20D3F8-CENTRAL-U1-FATE: decide whether central u(1) is projected, massive, global, or separately gauged",
		},
	}, nil
}

func SelectedActionCount(xs []CandidateAction) int {
	n := 0
	for _, x := range xs {
		if x.Selected {
			n++
		}
	}
	return n
}
