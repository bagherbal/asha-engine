package gaugeaction

import (
	"fmt"
	"strings"

	"github.com/bagherbal/asha-engine/pkg/theorem"
)

func GaugeKineticActionSelectionRGBoundaryAuditTheorem() theorem.Theorem {
	const id = "BRIDGE-GAUGE-KINETIC-ACTION-SELECTION-RG-BOUNDARY"
	const name = "gauge kinetic action selection / RG boundary coupling audit"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "build gauge kinetic action audit", Passed: false, Detail: err.Error()}}}
		}
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Checks: []theorem.Check{
			{Name: "Gate 81 representation-metric input", Passed: a.TraceGramAsRepresentationMetric, Detail: "diagonal U(1) trace-Gram metrics are available as representation diagnostics"},
			{Name: "U(1) gauge kinetic Hessian domain", Passed: a.SymmetricHessianDim == 6 && a.DiagonalHessianDim == 3, Detail: fmt.Sprintf("fields=%v; symmetric Hessian dim=%d, diagonal no-mixing dim=%d", a.Fields, a.SymmetricHessianDim, a.DiagonalHessianDim)},
			{Name: "candidate kinetic actions exposed", Passed: len(a.CandidateActions) >= 4, Detail: formatActions(a.CandidateActions)},
			{Name: "candidate positivity", Passed: allPositive(a.CandidateActions), Detail: "all exposed Hessian diagnostics are positive, but positivity does not select the physical action"},
			{Name: "trace-Gram promoted to gauge kinetic Hessian", Passed: a.TraceGramPromotedToGaugeKineticByAssumption, Detail: "not derived; promoting representation metric to gauge kinetic Hessian would be an additional assumption"},
			{Name: "charge-level boundary geometry preserved", Passed: a.ChargeTableKY > 0 && a.ChargeTableSin2Boundary > 0, Detail: fmt.Sprintf("k_Y=%.10f, sin²θ_boundary=%.10f remain charge-table boundary candidates", a.ChargeTableKY, a.ChargeTableSin2Boundary)},
			{Name: "two-carrier coupling diagnostic", Passed: a.TwoCarrierHyperchargeNormDiagnostic > 0 && a.TwoCarrierInverseNormDiagnostic > 0, Detail: fmt.Sprintf("bridge norm diagnostic=%.10f, inverse=%.10f; diagnostic only", a.TwoCarrierHyperchargeNormDiagnostic, a.TwoCarrierInverseNormDiagnostic)},
			{Name: "boundary coupling family", Passed: a.BoundaryCouplingFamilyExposed && !a.BoundaryCouplingFixed, Detail: "families of couplings are exposed, but no finite action fixes g_*"},
			{Name: "physical U(1) coupling and alpha", Passed: a.PhysicalU1CouplingDerived || a.FineStructureDerived, Detail: "not derived; alpha_em requires action-selected kinetic terms, boundary scale, and RG matching"},
			{Name: "hidden observed input", Passed: !a.HiddenObservedInputUsed, Detail: "no observed alpha, weak angle, coupling, or mass scale was inserted"},
		}, Notes: []string{a.TruthStatement, "Next: " + a.RecommendedNextGate}}
	}}
}

func formatActions(xs []CandidateAction) string {
	parts := make([]string, 0, len(xs))
	for _, x := range xs {
		state := "open"
		if x.Selected {
			state = "selected"
		}
		parts = append(parts, fmt.Sprintf("%s[%s]: %s", x.Name, state, x.MatrixRole))
	}
	return "[" + strings.Join(parts, "; ") + "]"
}

func allPositive(xs []CandidateAction) bool {
	if len(xs) == 0 {
		return false
	}
	for _, x := range xs {
		if !x.Positive {
			return false
		}
	}
	return true
}
