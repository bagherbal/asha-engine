package fierz

import (
	"fmt"

	"github.com/bagherbal/asha-engine/pkg/theorem"
)

func FiniteFierzProjectionAuditTheorem() theorem.Theorem {
	const id = "BRIDGE-FINITE-FIERZ-PROJECTION-AUDIT"
	const name = "finite Fierz projection / scalar-channel sign audit"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "construct finite Fierz audit", Passed: false, Detail: err.Error()}}}
		}
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Checks: []theorem.Check{
			{Name: "current-current domain", Passed: a.U4Dimension == 16 && a.CurrentSectorCount == 4, Detail: fmt.Sprintf("u(4) dim=%d split across %d current sectors", a.U4Dimension, a.CurrentSectorCount)},
			{Name: "scalar LR target available", Passed: a.ScalarLRTargetAvailable, Detail: fmt.Sprintf("target channels=%d = %d kinds × %d generations", a.ScalarChannelCount, a.ChannelKinds, a.GenerationCount)},
			{Name: "formal Fierz projection exposed", Passed: a.FormalProjectionExpression != "", Detail: a.FormalProjectionExpression},
			{Name: "chiral bilinear metric", Passed: a.ChiralBilinearMetricDerived, Detail: "open; Ψ̄_RΨ_L scalar projector is not yet constructed in the finite representation"},
			{Name: "Clifford trace rules", Passed: a.CliffordTraceRulesDerived, Detail: "open; native trace/Fierz identities are not computed"},
			{Name: "generator normalization", Passed: a.GeneratorNormalizationDerived, Detail: "open; relative current weights are not fixed by the current inventory alone"},
			{Name: "scalar-channel coefficients", Passed: a.ScalarProjectionCoefficientsKnown, Detail: "open; coefficients c_A for each current sector are still symbolic"},
			{Name: "attractive scalar-channel sign", Passed: a.AttractiveSignDerived, Detail: "open; attraction cannot be inferred from generator count or incidence alone"},
			{Name: "four-fermion strength G_hat", Passed: a.FourFermionStrengthDerived && a.NativeFierzProjectionComplete, Detail: "open; G_hat requires signed c_A plus propagator/kinetic normalization"},
			{Name: "up/down splitting", Passed: a.UpDownSplittingDerived, Detail: "open; Fierz audit has not broken the top/bottom tie"},
			{Name: "hidden observed input", Passed: !a.HiddenObservedInputUsed, Detail: "no observed Yukawa, v, Higgs mass, or fitted coupling was inserted"},
		}, Notes: []string{
			a.TruthStatement,
			fmt.Sprintf("projection slots: %s", FormatSlots(a.ProjectionSlots)),
			fmt.Sprintf("required tensors: %s", FormatRequirements(a.Requirements)),
			fmt.Sprintf("recommended next gate: %s", a.RecommendedNextGate),
			fmt.Sprintf("remaining unknowns: %s", FormatUnknowns(a.RemainingUnknowns)),
		}}
	}}
}
