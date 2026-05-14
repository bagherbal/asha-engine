package abeliancoupling

import (
	"fmt"
	"strings"

	"github.com/bagherbal/asha-engine/pkg/theorem"
)

func AbelianCouplingNormalizationFromDiagonalHessianAuditTheorem() theorem.Theorem {
	const id = "BRIDGE-ABELIAN-COUPLING-DIAGONAL-HESSIAN-AUDIT"
	const name = "abelian coupling normalization from diagonal Hessian audit"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "build abelian coupling normalization audit", Passed: false, Detail: err.Error()}}}
		}
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Checks: []theorem.Check{
			{Name: "Gate 80 diagonal/no-mixing input", Passed: a.AnomalyKinetic.DiagonalPositive && a.AnomalyKinetic.AllKnownOffDiagonalSourcesCancel, Detail: "diagonal trace-Gram sector survives; all known off-diagonal U(1) sources cancel"},
			{Name: "canonical representation metrics", Passed: a.DiagonalTraceGramSelectedAsRepresentationMetric, Detail: formatDiagnostics(a.Fields)},
			{Name: "inverse-norm coupling diagnostics", Passed: a.CanonicalGeneratorDiagnosticsDerived, Detail: "candidate diagnostic g_A^2 ∝ 1/Tr(T_A^2) exists, but is not selected as a physical coupling"},
			{Name: "charge-level hypercharge inherited", Passed: a.Hypercharge.ChargeTableKY > 0, Detail: fmt.Sprintf("k_Y=%.10f, sin²θ_boundary=%.10f remain charge-table results", a.Hypercharge.ChargeTableKY, a.Hypercharge.BoundarySin2)},
			{Name: "two-carrier hypercharge norm diagnostic", Passed: a.Hypercharge.CombinedBridgeNorm > 0, Detail: fmt.Sprintf("(1/2)^2 Tr((B-L)^2)=%.10f, Tr(T_phi^2)=%.10f, bridge diagnostic sum=%.10f", a.Hypercharge.MatterBLContribution, a.Hypercharge.ScalarContactContribution, a.Hypercharge.CombinedBridgeNorm)},
			{Name: "diagonal trace-Gram is gauge kinetic Hessian", Passed: a.DiagonalTraceGramSelectedAsGaugeKineticHessian, Detail: "not derived; representation metric is not automatically the gauge-field kinetic Hessian"},
			{Name: "physical U(1) coupling and alpha", Passed: a.PhysicalGaugeCouplingsDerived || a.FineStructureDerived, Detail: "not derived; g_Y and alpha_em require action-selected kinetic terms plus RG boundary/matching"},
			{Name: "hidden observed input", Passed: !a.HiddenObservedInputUsed, Detail: "no observed alpha, theta_W, coupling, mass, or scale was inserted"},
		}, Notes: []string{a.TruthStatement, "Next: " + a.RecommendedNextGate}}
	}}
}

func formatDiagnostics(xs []CouplingDiagnostic) string {
	parts := make([]string, 0, len(xs))
	for _, x := range xs {
		parts = append(parts, fmt.Sprintf("%s: Tr(T²)=%.10f, ||T||=%.10f, inverse=%.10f", x.Name, x.TraceNormSquared, x.TraceNorm, x.InverseNormSquared))
	}
	return "[" + strings.Join(parts, "; ") + "]"
}
