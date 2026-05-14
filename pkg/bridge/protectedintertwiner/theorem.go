package protectedintertwiner

import (
	"fmt"
	"strings"

	"github.com/bagherbal/asha-engine/pkg/theorem"
)

func ProtectedContactBrokenGeneratorIntertwinerSearchTheorem() theorem.Theorem {
	const id = "BRIDGE-PROTECTED-CONTACT-BROKEN-GENERATOR-INTERTWINER"
	const name = "protected-contact to broken-generator intertwiner search"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "build protected-contact/broken-generator audit", Passed: false, Detail: err.Error()}}}
		}
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Checks: []theorem.Check{
			{Name: "Gate 86 scalar-vacuum input", Passed: a.ScalarVacuum.LowPairSelected, Detail: "finite scalar response selects the lower active pair but leaves residual phase freedom"},
			{Name: "protected/goldstone/broken dimension match", Passed: a.CountLevelResonance, Detail: fmt.Sprintf("protected=%d, scalar-angular=%d, broken-generators=%d, broken-image-rank=%d", a.ProtectedContactDirections, a.ScalarAngularDirections, a.BrokenGeneratorDirections, a.BrokenImageRank)},
			{Name: "broken-generator image metric", Passed: a.BrokenImageMetricPositive, Detail: fmt.Sprintf("min eigen=%s, max eigen=%s, condition=%s", FormatFloat(a.BrokenImageMinEigen), FormatFloat(a.BrokenImageMaxEigen), FormatFloat(a.BrokenImageCondition))},
			{Name: "abstract isometry exists", Passed: a.AbstractIsometryExists, Detail: fmt.Sprintf("any two 3D Euclidean frames admit an %s family; freedom dimension=%d", a.AbstractIsometryGroup, a.AbstractIsometryFreedomDimension)},
			{Name: "protected contact metric/connection", Passed: a.ProtectedContactMetricAvailable, Detail: "not derived; the engine has protected-direction count but not a canonical protected-contact frame metric/connection"},
			{Name: "canonical protected-to-broken map", Passed: a.CanonicalIntertwinerDerived, Detail: "not derived; count equality leaves an O(3) identification freedom"},
			{Name: "finite gauge-eating theorem", Passed: a.GaugeEatingTheoremDerived, Detail: "not derived; still requires selected intertwiner plus scalar/gauge kinetic actions"},
		}, Notes: []string{a.TruthStatement, "remaining unknowns: " + strings.Join(a.RemainingUnknowns, "; "), "Next: " + a.RecommendedNextGate}}
	}}
}
