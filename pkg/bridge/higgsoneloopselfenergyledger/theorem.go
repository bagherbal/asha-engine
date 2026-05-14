package higgsoneloopselfenergyledger

import "github.com/bagherbal/asha-engine/pkg/theorem"

func HiggsOneLoopSelfEnergyComponentLedgerRenormalizedPoleKernelAuditTheorem() theorem.Theorem {
	const id = "BRIDGE-HIGGS-ONE-LOOP-SELF-ENERGY-COMPONENT-LEDGER-RENORMALIZED-POLE-KERNEL-AUDIT"
	const name = "Higgs One-Loop Self-Energy Component Ledger / Renormalized Pole Kernel Audit"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "build Gate 333 one-loop self-energy ledger", Passed: false, Detail: err.Error()}}}
		}
		checks := []theorem.Check{
			{Name: "Gate 332 pole target inherited", Passed: a.Inputs.HighestInheritedGate == inheritedHighestGate && a.Inputs.RequiredRePiGeV2 > 40 && a.Inputs.NativeProxyMassGeV > a.Inputs.ObservedPoleGeV, Detail: FormatInputs(a.Inputs)},
			{Name: "one-loop component ledger formalized", Passed: len(a.Ledger.Components) == 4 && a.Ledger.Components[0].ContributionGeV2 < 0 && a.Ledger.Components[1].ContributionGeV2 > 0, Detail: FormatLedger(a.Ledger)},
			{Name: "raw one-loop kernel is top dominated and not the finite target", Passed: a.Kernel.RawKernelGeV2 < -900 && !a.Kernel.MatchesTarget, Detail: FormatKernel(a.Kernel)},
			{Name: "renormalized counterterm ledger is mandatory", Passed: a.Counterterms.CountertermMandatory && a.Counterterms.RequiredFiniteCountertermGeV2 > 1000, Detail: FormatCounterterms(a.Counterterms)},
			{Name: "scheme ledger requires PV functions and finite counterterms", Passed: a.Scheme.NeedsPVIntegrals && a.Scheme.NeedsFiniteCounterterms && a.Scheme.NeedsMassInputScheme && !a.Scheme.ExactPoleMassComputed, Detail: FormatScheme(a.Scheme)},
			{Name: "firewall preserves no exact collider pole-mass claim", Passed: a.Audit.NoExactPVFunctions && a.Audit.NoCountertermDerivation && a.Audit.NoExactColliderClaim && a.Audit.NoNativeSMInputClaim, Detail: FormatAudit(a.Audit)},
		}
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Checks: checks, Notes: []string{a.Truth, FormatSummary(a.Summary), FormatStatuses(Statuses(a))}}
	}}
}
