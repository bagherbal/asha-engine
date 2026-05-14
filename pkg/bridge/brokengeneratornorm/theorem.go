package brokengeneratornorm

import (
	"fmt"

	"github.com/bagherbal/asha-engine/pkg/theorem"
)

func NormalizedBrokenGeneratorBasisGaugeKineticCandidateTheorem() theorem.Theorem {
	const id = "BRIDGE-NORMALIZED-BROKEN-GENERATOR-BASIS"
	const name = "normalized broken-generator basis / gauge-kinetic candidate search"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "build normalized broken-generator basis audit", Passed: false, Detail: err.Error()}}}
		}
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Checks: []theorem.Check{
			{Name: "Gate 92 broken-metric input", Passed: a.BrokenMetric.IsotropizationExact, Detail: fmt.Sprintf("raw condition=%.10f, neutral factor=%.10f", a.RawCondition, a.NeutralNormalization)},
			{Name: "metric-selected diagnostic normalization", Passed: a.MetricSelectsDiagnosticBasis, Detail: fmt.Sprintf("factors=[%.1f, %.1f, %.10f] for [%s]", a.NormalizationFactors[0], a.NormalizationFactors[1], a.NormalizationFactors[2], joinNames(a.RawGeneratorNames))},
			{Name: "normalized broken-image metric", Passed: a.NormalizationExact, Detail: fmt.Sprintf("diag=[%.10f, %.10f, %.10f], condition=%.10f", a.NormalizedMetricDiagonal[0], a.NormalizedMetricDiagonal[1], a.NormalizedMetricDiagonal[2], a.NormalizedCondition)},
			{Name: "raw-coordinate gauge-kinetic candidate", Passed: a.KineticCandidatePositive, Detail: fmt.Sprintf("K_diag=[%.10f, %.10f, %.10f], det=%.10f", a.RawCoordinateKineticCandidate[0], a.RawCoordinateKineticCandidate[1], a.RawCoordinateKineticCandidate[2], a.KineticCandidateDeterminant)},
			{Name: "finite action selects normalization", Passed: a.FiniteActionSelectsBasis, Detail: "not derived; the factor 1/2 is selected by diagnostic metric isotropization, not by a finite gauge-field action"},
			{Name: "gauge kinetic Hessian selected", Passed: a.GaugeKineticHessianSelected, Detail: "not selected; diag(1,1,4) remains a candidate Hessian, not a theorem"},
			{Name: "physical anisotropy derived", Passed: a.PhysicalAnisotropyDerived, Detail: "not derived; raw anisotropy is removable by generator normalization"},
			{Name: "physical W/Z masses or alpha", Passed: a.PhysicalMassesDerived, Detail: "not derived; couplings, scalar kinetic action, and RG bridge remain open"},
		}, Notes: []string{a.TruthStatement, "rejected claims: " + Join(a.RejectedClaims), "remaining unknowns: " + Join(a.RemainingUnknowns), "Next: " + a.RecommendedNextGate}}
	}}
}

func joinNames(xs []string) string {
	out := ""
	for i, x := range xs {
		if i > 0 {
			out += ", "
		}
		out += x
	}
	return out
}
