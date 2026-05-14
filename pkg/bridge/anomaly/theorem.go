package anomaly

import (
	"fmt"
	"strings"

	"github.com/bagherbal/asha-engine/pkg/theorem"
)

func AbelianAnomalyCancellationLedgerTheorem() theorem.Theorem {
	const id = "BRIDGE-ABELIAN-ANOMALY-CANCELLATION-LEDGER"
	const name = "anomaly / cancellation ledger for abelian sources"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "build anomaly cancellation ledger", Passed: false, Detail: err.Error()}}}
		}
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Checks: []theorem.Check{
			{Name: "Gate 78 cancellation input", Passed: a.YukawaIncidenceCancellation, Detail: fmt.Sprintf("natural orientational sources cancel; non-factorized signed source=%.3e", a.Previous.Previous.SignedBLContactMoment)},
			{Name: "left-handed anomaly table", Passed: len(a.States) == 16, Detail: "right-handed states are converted to left-handed conjugates; one-generation Weyl table has 16 states"},
			{Name: "hypercharge anomaly ledger", Passed: a.YAnomalyCancels, Detail: formatMoments(a.HyperchargeMoments)},
			{Name: "B-L anomaly ledger", Passed: a.BMinusLAnomalyCancels, Detail: formatMoments(a.BMinusLMoments)},
			{Name: "mixed abelian anomaly ledger", Passed: a.MixedAbelianCancels, Detail: formatMoments(a.MixedMoments)},
			{Name: "anomaly-shadow interpretation", Passed: a.AnomalyShadowSupported, Detail: "Yukawa-incidence cancellation is consistent with anomaly-balanced finite charge bookkeeping"},
			{Name: "stricter no-mixing theorem", Passed: a.StricterNoMixingTheorem, Detail: "not derived; anomaly cancellation does not by itself prove all kinetic mixing terms vanish"},
			{Name: "kinetic mixing source", Passed: a.KineticMixingDerived, Detail: "not derived; physical U(1)_Y coupling remains bridge-gated"},
			{Name: "hidden observed input", Passed: !a.HiddenObservedInputUsed, Detail: "no observed alpha, theta_W, v, or measured coupling was inserted"},
		}, Notes: []string{a.TruthStatement, "Next: " + a.RecommendedNextGate}}
	}}
}

func formatMoments(xs []Moment) string {
	parts := make([]string, 0, len(xs))
	for _, x := range xs {
		parts = append(parts, fmt.Sprintf("%s=%.3e(cancel=%v)", x.Name, x.Value, x.Cancels))
	}
	return "[" + strings.Join(parts, "; ") + "]"
}
