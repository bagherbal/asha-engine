package quotientedcorrespondence

import (
	"fmt"

	"github.com/bagherbal/asha-engine/pkg/theorem"
)

func GaugeQuotientedProtectedBrokenCorrespondenceAuditTheorem() theorem.Theorem {
	const id = "BRIDGE-GAUGE-QUOTIENTED-PROTECTED-BROKEN-CORRESPONDENCE"
	const name = "gauge-quotiented protected-to-broken correspondence audit"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "build quotient-safe correspondence audit", Passed: false, Detail: err.Error()}}}
		}
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Checks: []theorem.Check{
			{Name: "Gate 90 O(3) quotient input", Passed: a.O3Quotient.CurrentDataSupportsGaugeQuotient, Detail: "protected-frame O(3) freedom is quotiented for current protected-contact diagnostics"},
			{Name: "quotient-safe count correspondence", Passed: a.CountCorrespondenceSurvivesQuotient, Detail: fmt.Sprintf("protected dim=%d, broken-generator directions=%d, broken-image rank=%d", a.ProtectedDimension, a.BrokenDimension, a.BrokenRank)},
			{Name: "component-wise frame matching rejected", Passed: a.FrameComponentComparisonRejected, Detail: "after quotienting O(3), only invariant data may be compared; frame components are gauge choices"},
			{Name: "broken metric spectrum is quotient invariant", Passed: a.BrokenMetricSpectrumInvariant, Detail: fmt.Sprintf("eigen range=[%.10f, %.10f], condition=%.10f", a.BrokenMetricMinEigen, a.BrokenMetricMaxEigen, a.BrokenMetricCondition)},
			{Name: "protected metric reference spectrum", Passed: len(a.ProtectedMetricSpectrum) == 3, Detail: "abstract protected metric spectrum=[1,1,1] before any action-selected normalization"},
			{Name: "broken metric isotropy", Passed: a.BrokenMetricIsIsotropic, Detail: "not isotropic; anisotropic broken-image metric cannot be identified with protected I3 by quotient-safe data alone"},
			{Name: "metric isometry derived", Passed: a.MetricIsometryDerived, Detail: "not derived; an abstract O(3) family exists but no finite quotient-safe isometry is selected"},
			{Name: "quotient-safe protected-to-broken intertwiner", Passed: a.QuotientSafeIntertwinerDerived, Detail: "not derived; count/rank correspondence survives, canonical map remains open"},
			{Name: "gauge-eating bridge completed", Passed: a.GaugeEatingBridgeCompleted, Detail: "not completed; still requires scalar/gauge kinetic action and quotient-safe intertwiner"},
		}, Notes: []string{a.TruthStatement, "quotient-safe invariants: " + Join(a.QuotientSafeInvariants), "rejected moves: " + Join(a.RejectedMoves), "remaining unknowns: " + Join(a.RemainingUnknowns), "Next: " + a.RecommendedNextGate}}
	}}
}
