package protectedmetric

import (
	"fmt"

	"github.com/bagherbal/asha-engine/pkg/theorem"
)

func ProtectedContactMetricConnectionSearchTheorem() theorem.Theorem {
	const id = "BRIDGE-PROTECTED-CONTACT-METRIC-CONNECTION-SEARCH"
	const name = "protected-contact metric and connection form search"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "build protected metric audit", Passed: false, Detail: err.Error()}}}
		}
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Checks: []theorem.Check{
			{Name: "Gate 87 protected/broken input", Passed: a.Intertwiner.CountLevelResonance, Detail: fmt.Sprintf("protected=%d, broken-image-rank=%d, abstract isometry=%v", a.ProtectedDimension, a.BrokenImageDimension, a.Intertwiner.AbstractIsometryExists)},
			{Name: "abstract protected Euclidean metric", Passed: a.AbstractEuclideanMetricAvailable, Detail: fmt.Sprintf("I_3 diagnostic: trace=%.1f, det=%.1f", a.AbstractEuclideanMetricTrace, a.AbstractEuclideanMetricDeterminant)},
			{Name: "abstract metric is not contact dynamics", Passed: !a.FiniteProtectedMetricDerived, Detail: "I_3 is a coordinate metric on a 3D carrier, not a finite BF/contact-derived metric"},
			{Name: "broken image metric available", Passed: a.BrokenImageMetricAvailable, Detail: fmt.Sprintf("positive broken-image metric exists with condition=%s", FormatFloat(a.BrokenImageMetricCondition))},
			{Name: "broken metric pullback", Passed: a.BrokenMetricPullbackDerived, Detail: "not derived; pulling it back requires the protected-to-broken intertwiner and would be circular before that map is selected"},
			{Name: "pullback circularity detected", Passed: a.PullbackCircularityDetected, Detail: "the metric-pullback route cannot be used to derive the same intertwiner it assumes"},
			{Name: "protected contact connection form", Passed: a.FiniteProtectedConnectionDerived, Detail: "not derived; no intrinsic protected-carrier connection/curvature form is available yet"},
			{Name: "O(3) freedom reduced", Passed: a.O3FreedomReduced, Detail: "not reduced by current data; it may be pure gauge, but the quotient theorem is still open"},
			{Name: "canonical protected frame", Passed: a.CanonicalProtectedFrameDerived, Detail: "not derived; arbitrary frame fixing is rejected"},
		}, Notes: []string{a.TruthStatement, "candidate metric sources: " + Join(a.CandidateMetricSources), "rejected shortcuts: " + Join(a.RejectedMetricSources), "remaining unknowns: " + Join(a.RemainingUnknowns), "Next: " + a.RecommendedNextGate}}
	}}
}
