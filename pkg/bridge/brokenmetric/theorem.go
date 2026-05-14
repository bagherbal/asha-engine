package brokenmetric

import (
	"fmt"

	"github.com/bagherbal/asha-engine/pkg/theorem"
)

func BrokenImageMetricKineticNormalizationAuditTheorem() theorem.Theorem {
	const id = "BRIDGE-BROKEN-IMAGE-METRIC-KINETIC-NORMALIZATION"
	const name = "broken-image metric / kinetic normalization audit"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "build broken-image metric audit", Passed: false, Detail: err.Error()}}}
		}
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Checks: []theorem.Check{
			{Name: "Gate 91 quotient-safe input", Passed: a.QuotientedCorrespondence.CountCorrespondenceSurvivesQuotient, Detail: "protected/broken count correspondence survives O(3) quotient; component-wise frame matching remains rejected"},
			{Name: "raw broken-image metric", Passed: a.RawAnisotropic, Detail: fmt.Sprintf("charged eigen=%.10f, neutral eigen=%.10f, condition=%.10f", a.ChargedEigenvalue, a.NeutralEigenvalue, a.RawCondition)},
			{Name: "neutral-to-charged ratio", Passed: abs(a.NeutralToChargedRatio-4) < 1e-8, Detail: fmt.Sprintf("neutral/charged=%.10f", a.NeutralToChargedRatio)},
			{Name: "normalization that removes anisotropy", Passed: a.IsotropizationPossible, Detail: fmt.Sprintf("neutral generator factor=%.10f; normalized neutral eigen=%.10f", a.NeutralNormFactor, a.NormalizedNeutralEigen)},
			{Name: "isotropized broken metric diagnostic", Passed: a.IsotropizationExact, Detail: fmt.Sprintf("eigen=[%.10f, %.10f, %.10f], condition=%.10f", a.IsotropizedEigenvalues[0], a.IsotropizedEigenvalues[1], a.IsotropizedEigenvalues[2], a.IsotropizedCondition)},
			{Name: "physical anisotropy derived", Passed: a.PhysicalAnisotropyDerived, Detail: "not derived; raw anisotropy can be removed by generator normalization"},
			{Name: "gauge-normalization artifact possible", Passed: a.GaugeNormalizationArtifactPossible, Detail: "yes; broken-image anisotropy is currently normalization data, not a physical mass prediction"},
			{Name: "scalar kinetic normalization selected", Passed: a.ScalarKineticNormalizationSelected, Detail: "not selected by current finite action"},
			{Name: "gauge kinetic normalization selected", Passed: a.GaugeKineticNormalizationSelected, Detail: "not selected by current finite action"},
			{Name: "gauge-eating theorem completed", Passed: a.GaugeEatingTheoremCompleted, Detail: "not completed; needs action-selected scalar/gauge kinetic normalization and quotient-safe intertwiner"},
		}, Notes: []string{a.TruthStatement, "quotient-safe conclusion: " + a.QuotientSafeConclusion, "rejected claims: " + Join(a.RejectedClaims), "remaining unknowns: " + Join(a.RemainingUnknowns), "Next: " + a.RecommendedNextGate}}
	}}
}

func abs(x float64) float64 {
	if x < 0 {
		return -x
	}
	return x
}
