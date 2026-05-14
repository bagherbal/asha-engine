package o3quotient

import (
	"fmt"

	"github.com/bagherbal/asha-engine/pkg/theorem"
)

func O3GaugeQuotientPhysicalOrientationAuditTheorem() theorem.Theorem {
	const id = "BRIDGE-O3-GAUGE-QUOTIENT-PHYSICAL-ORIENTATION"
	const name = "O(3) gauge quotient and physical orientation audit"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "build O(3) quotient audit", Passed: false, Detail: err.Error()}}}
		}
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Checks: []theorem.Check{
			{Name: "Gate 89 protected-connection input", Passed: !a.ProtectedConnection.IntrinsicProtectedOperatorDerived && !a.ProtectedConnection.O3FreedomReduced, Detail: "no intrinsic protected operator/connection was selected by Gate 89"},
			{Name: "protected frame family", Passed: a.ProtectedDimension == 3 && a.O3Dimension == 3, Detail: fmt.Sprintf("protected carrier dim=%d; frame freedom=%s with dim=%d", a.ProtectedDimension, a.AbstractFrameFamily, a.O3Dimension)},
			{Name: "intrinsic metric is O(3)-invariant", Passed: a.IntrinsicMetricIsO3Invariant, Detail: "current protected metric data is only the abstract Euclidean I3 metric"},
			{Name: "protected curvature is O(3)-invariant", Passed: a.ProtectedCurvatureIsO3Invariant, Detail: fmt.Sprintf("protected curvature span rank=%d and max norm inherited as %.3e", a.ProtectedConnection.ContactCurvatureSpanRank, a.ProtectedConnection.ContactCurvatureMaxNorm)},
			{Name: "protected-contact observables are frame-independent", Passed: a.ProtectedContactObservablesFrameIndependent, Detail: "current intrinsic protected-contact diagnostics do not observe the O(3) frame"},
			{Name: "diagonal spurion does not select protected orientation", Passed: a.DiagonalSpurionExists && !a.DiagonalSpurionIntrinsic && !a.DiagonalSpurionPhysicalOrientation, Detail: "bridge-level generation spurion exists but is not intrinsic protected-frame data"},
			{Name: "current data supports gauge quotient", Passed: a.CurrentDataSupportsGaugeQuotient, Detail: a.QuotientStatement},
			{Name: "full no-orientation theorem", Passed: a.FullNoOrientationTheoremProven, Detail: "not proven; a future finite action could still make orientation observable"},
			{Name: "physical protected orientation selected", Passed: a.PhysicalOrientationSelected, Detail: "not selected; choosing a protected frame remains a gauge/bridge choice at this stage"},
		}, Notes: []string{a.TruthStatement, "gauge-quotient evidence: " + Join(a.GaugeQuotientEvidence), "physical-orientation openings: " + Join(a.PhysicalOrientationOpenings), "rejected moves: " + Join(a.RejectedMoves), "remaining unknowns: " + Join(a.RemainingUnknowns), "Next: " + a.RecommendedNextGate}}
	}}
}
