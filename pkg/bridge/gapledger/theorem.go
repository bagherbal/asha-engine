package gapledger

import (
	"fmt"

	"github.com/bagherbal/asha-engine/pkg/theorem"
)

func NJLGapKernelCriticalityLedgerTheorem() theorem.Theorem {
	const id = "BRIDGE-NJL-GAP-KERNEL-CRITICALITY-LEDGER"
	const name = "NJL gap-kernel / criticality ledger"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "construct NJL criticality ledger", Passed: false, Detail: err.Error()}}}
		}
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Checks: []theorem.Check{
			{Name: "finite criticality domain", Passed: a.GenerationCount == 3 && a.KindCount == 4, Detail: fmt.Sprintf("generations=%d, fermion kinds=%d, channel states=%d", a.GenerationCount, a.KindCount, len(a.ChannelCriticalities))},
			{Name: "three-color amplification retained", Passed: a.QuarkLeptonAmplification == 3, Detail: fmt.Sprintf("unit quark pressure=%.0f, unit lepton pressure=%.0f, amplification=%.1f", a.UnitQuarkPressureSkeleton, a.UnitLeptonPressureSkeleton, a.QuarkLeptonAmplification)},
			{Name: "strongest finite pressure skeleton", Passed: a.StrongestWeightedPressure > 0, Detail: fmt.Sprintf("strongest=%s", FormatChannel(a.StrongestWeightedChannel))},
			{Name: "unit-threshold diagnostic", Passed: a.UnitCriticalDiagnosticAvailable, Detail: fmt.Sprintf("if one artificially sets C_reg=1, required G_hat>%.10f; this is diagnostic, not physical", a.UnitThresholdRequiredCoupling)},
			{Name: "formal NJL condition exposed", Passed: len(a.FormalConditions) >= 4, Detail: "G_hat · K_channel > C_reg; finite K_channel known, G_hat and C_reg open"},
			{Name: "up/down tie remains", Passed: a.UpDownDegeneracyResidual == 0, Detail: "up-type and down-type quark channels still have equal unit-incidence pressure; top-only criticality is not selected"},
			{Name: "four-fermion attractive kernel", Passed: a.FourFermionKernelDerived && a.AttractiveInteractionDerived, Detail: "open; no native x∧p/u(4) four-fermion kernel has been derived"},
			{Name: "regulator and critical threshold", Passed: a.RegulatorDerived && a.CriticalThresholdDerived, Detail: "open; finite cutoff/spectral regulator and C_reg are not derived"},
			{Name: "gap equation solved", Passed: a.GapEquationSolved && a.NativeNJLComputationComplete, Detail: "open; no nonzero condensate solution is computed"},
			{Name: "condensate scale", Passed: a.CondensateScaleDerived, Detail: "open; no physical unit or v-scale is derived"},
			{Name: "hidden observed input", Passed: !a.HiddenObservedCouplingsUsed && !a.HiddenObservedMassScalesUsed, Detail: "no observed y_t, v, Higgs mass, or fermion mass was inserted"},
		}, Notes: []string{
			a.TruthStatement,
			fmt.Sprintf("all channel skeletons: %s", FormatChannels(a.ChannelCriticalities)),
			fmt.Sprintf("formal conditions: %s", FormatConditions(a.FormalConditions)),
			fmt.Sprintf("recommended next gate: %s", a.RecommendedNextGate),
			fmt.Sprintf("remaining unknowns: %s", FormatUnknowns(a.RemainingUnknowns)),
		}}
	}}
}
