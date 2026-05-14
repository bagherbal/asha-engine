package looppotential

import (
	"fmt"

	"github.com/bagherbal/asha-engine/pkg/theorem"
)

func NativeOneLoopPotentialLedgerTheorem() theorem.Theorem {
	const id = "BRIDGE-NATIVE-ONE-LOOP-POTENTIAL-LEDGER"
	const name = "native one-loop effective-potential ledger"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "construct one-loop ledger", Passed: false, Detail: err.Error()}}}
		}
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Checks: []theorem.Check{
			{Name: "anticommuting Fock substrate", Passed: a.AnticommutationSubstrateAvailable && a.Fock.ModeCount() == 4 && a.Fock.StateCount() == 16, Detail: fmt.Sprintf("Fock modes=%d, states=%d; fermion-loop sign ledger available=%v", a.Fock.ModeCount(), a.Fock.StateCount(), a.FermionLoopSignAvailable)},
			{Name: "three-color amplification", Passed: a.ColorAmplificationFactor == 3, Detail: fmt.Sprintf("spatial/color channels=%d from gauge-compatible up/down Yukawa channels", a.ColorAmplificationFactor)},
			{Name: "top-like negative skeleton", Passed: a.TopLikeCoefficientSkeleton == -6, Detail: fmt.Sprintf("coefficient skeleton=%d·y_top-like²; sign/multiplicity available, strength not derived", a.TopLikeCoefficientSkeleton)},
			{Name: "positive bosonic sectors exposed", Passed: a.GaugePositiveSectorCount == 4 && a.ScalarSelfSectorAvailable, Detail: fmt.Sprintf("gauge positive sector count=%d; scalar quartic shape=%.10f", a.GaugePositiveSectorCount, a.ScalarQuarticShapeInvariant)},
			{Name: "native loop operator", Passed: a.NativeLoopOperatorDerived, Detail: "open; finite Fock/Yukawa loop trace has not been constructed"},
			{Name: "top-like Yukawa strength", Passed: a.TopLikeYukawaStrengthDerived, Detail: "open; no observed top Yukawa or fitted coupling is inserted"},
			{Name: "bosonic counterweights", Passed: a.GaugeCouplingsDerived && a.ScalarSelfCouplingScaleDerived, Detail: "open; gauge/scalar couplings require kinetic normalization and finite trace rules"},
			{Name: "regulator and renormalization prescription", Passed: a.RegulatorOrCutoffDerived && a.RenormalizationPrescriptionDerived, Detail: "open; finite spectral cutoff/renormalization rule not derived"},
			{Name: "native instability sign", Passed: a.MuSquaredSignDerived && a.SymmetricOriginInstabilityDerived, Detail: "open; μ²_eff<0 is not derived by this ledger"},
			{Name: "imported Standard Model RGE", Passed: !a.ImportedSMRGE, Detail: "not imported; this gate is a native-computation ledger, not an RGE lookup"},
			{Name: "hidden observed couplings", Passed: !a.HiddenObservedCouplingsUsed, Detail: "no observed y_t, g, g′, λ, v, or Higgs mass was inserted"},
		}, Notes: []string{
			a.TruthStatement,
			fmt.Sprintf("ledger terms: %s", FormatTerms(a.Terms)),
			fmt.Sprintf("recommended next gate: %s", a.RecommendedNextGate),
			fmt.Sprintf("remaining unknowns: %v", a.RemainingUnknowns),
		}}
	}}
}
