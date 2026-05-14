package gaugekineticdiag

import (
	"fmt"

	"github.com/bagherbal/asha-engine/pkg/theorem"
)

func GaugeKineticDiagActionSelectionAuditTheorem() theorem.Theorem {
	const id = "BRIDGE-GAUGE-KINETIC-DIAG-114-ACTION-SELECTION"
	const name = "gauge-kinetic Hessian diag(1,1,4) action-selection audit"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "build diag(1,1,4) gauge-kinetic audit", Passed: false, Detail: err.Error()}}}
		}
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Checks: []theorem.Check{
			{Name: "Gate 93 normalized-basis input", Passed: a.BrokenGeneratorNorm.NormalizationExact, Detail: fmt.Sprintf("neutral factor=%.10f; normalized condition=%.10f", a.NeutralFactor, a.BrokenGeneratorNorm.NormalizedCondition)},
			{Name: "raw-coordinate Hessian candidate", Passed: a.CandidatePositive, Detail: fmt.Sprintf("%s with diag=[%.10f, %.10f, %.10f], det=%.10f", a.CandidateName, a.CandidateDiagonal[0], a.CandidateDiagonal[1], a.CandidateDiagonal[2], a.CandidateDeterminant)},
			{Name: "metric whitening diagnostic", Passed: a.WhitenedExact, Detail: fmt.Sprintf("raw metric=[%.10f, %.10f, %.10f]; K^-1 raw metric=[%.10f, %.10f, %.10f], condition=%.10f", a.RawMetricDiagonal[0], a.RawMetricDiagonal[1], a.RawMetricDiagonal[2], a.WhitenedMetricDiagonal[0], a.WhitenedMetricDiagonal[1], a.WhitenedMetricDiagonal[2], a.WhitenedCondition)},
			{Name: "compatible with W/Z diagnostic basis", Passed: a.CompatibleWithMassDiagnostic, Detail: "diag(1,1,4) exactly removes the raw neutral/chairged image-metric anisotropy in broken-coordinate space"},
			{Name: "finite action selects diag(1,1,4)", Passed: a.SelectedByFiniteAction, Detail: "not derived; the candidate is selected by metric whitening, not by a finite action second variation"},
			{Name: "scalar kinetic action selected", Passed: a.ScalarKineticActionSelected, Detail: "not derived; active-frame scalar metric is still diagnostic"},
			{Name: "gauge Hessian selected", Passed: a.GaugeHessianSelected, Detail: "not selected; physical gauge-field kinetic operator remains open"},
			{Name: "physical couplings or masses", Passed: a.PhysicalCouplingsDerived || a.PhysicalMassesDerived, Detail: "not derived; no claim of g2, gY, thetaW, alpha, W mass, or Z mass is allowed"},
		}, Notes: []string{a.TruthStatement, "rejected claims: " + Join(a.RejectedClaims), "remaining unknowns: " + Join(a.RemainingUnknowns), "Next: " + a.RecommendedNextGate}}
	}}
}
