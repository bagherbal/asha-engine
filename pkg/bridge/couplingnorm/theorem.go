package couplingnorm

import (
	"fmt"

	"github.com/bagherbal/asha-engine/pkg/theorem"
)

func CouplingNormalizationBridgeTheorem() theorem.Theorem {
	const id = "BRIDGE-COUPLING-NORMALIZATION"
	const name = "coupling-normalization bridge audit"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "construct coupling-normalization audit", Passed: false, Detail: err.Error()}}}
		}
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Checks: []theorem.Check{
			{Name: "topological action input", Passed: a.TopologicalActionSeal > 0, Detail: fmt.Sprintf("I_BG=%.10f, S_top=%.10f, exp(-S_top)=%.3e", a.ContactIndex, a.TopologicalActionSeal, a.InstantonWeight)},
			{Name: "Yang-Mills normalization form", Passed: a.ContinuumFormula != "" && a.FiniteFormula != "", Detail: fmt.Sprintf("%s; %s", a.ContinuumFormula, a.FiniteFormula)},
			{Name: "coupling family exposed", Passed: a.CouplingFamily != "", Detail: a.CouplingFamily},
			{Name: "unit-trace diagnostic", Passed: a.UnitTraceGaugeCouplingSq > 0 && a.UnitTraceInverseAlpha > 0, Detail: fmt.Sprintf("κ=%.1f ⇒ g²=%.10f, g=%.10f, α_unit=%.10f, α_unit^{-1}=%.10f", a.UnitTraceNormalization, a.UnitTraceGaugeCouplingSq, a.UnitTraceGaugeCoupling, a.UnitTraceAlpha, a.UnitTraceInverseAlpha)},
			{Name: "continuum index bridge", Passed: a.ContinuumIndexBridgeDerived, Detail: "not yet derived; finite I_BG=1 is not automatically the continuum instanton charge in the kinetic action"},
			{Name: "trace/kinetic normalization", Passed: a.TraceNormalizationDerived, Detail: "not yet derived; κ remains open, so g²=1/κ is not fixed"},
			{Name: "gauge coupling derived", Passed: a.GaugeCouplingDerived, Detail: "not derived; unit-trace g²=1 is only a convention diagnostic"},
			{Name: "RG boundary condition", Passed: a.RGBoundaryDerived, Detail: "not derived; no running coupling or threshold scale has been fixed"},
			{Name: "fine-structure constant", Passed: a.FineStructureDerived, Detail: "not derived; α_em requires electroweak projection, mixing angle, RG scale, and normalization"},
			{Name: "dimensionful scale", Passed: a.DimensionfulScaleDerived, Detail: "not derived; coupling normalization still does not supply a mass unit μ"},
			{Name: "hidden observed coupling insertion", Passed: !a.HiddenObservedCouplingUsed, Detail: "no measured coupling or α value was inserted"},
		}, Notes: []string{
			a.TruthStatement,
			fmt.Sprintf("remaining unknowns: %v", a.RemainingUnknowns),
		}}
	}}
}
