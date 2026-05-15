package generation2syntheticinversion

import "github.com/bagherbal/asha-engine/pkg/theorem"

func Generation2CommonScaleSyntheticInversionUncertaintyPropagationHarnessTheorem() theorem.Theorem {
	const id = AuditID
	const name = "Generation 2 common-scale synthetic inversion uncertainty propagation harness"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "construct Gate468 audit", Passed: false, Detail: err.Error()}}}
		}
		checks := []theorem.Check{
			{Name: "inherits Gate467 rank-complete ledger contract", Passed: a.Inheritance.Executed && a.Inheritance.Gate467CommonScaleLedger && a.Inheritance.Gate467RequiresISpecIK && a.Inheritance.Gate467RequiresBranchTags && a.Inheritance.Gate467RequiresUncertainty && a.Inheritance.Gate467DUDComputableIfNumeric && a.Inheritance.Gate467DidNotComputeDUD && a.Inheritance.NativeRegistryClean, Detail: FormatInheritance(a.Inheritance)},
			{Name: "executes synthetic inverse and d_ud socket", Passed: a.Harness.Executed && a.Harness.AcceptedSyntheticCases == 1 && a.Harness.ValidSyntheticDUDComputed && a.Harness.UncertaintyPropagationExecuted && a.Harness.AllAcceptedBridgeOnlySynthetic, Detail: FormatHarness(a.Harness)},
			{Name: "rejects unsafe synthetic inversion routes", Passed: a.Harness.ObservedDataRejected && a.Harness.MissingRankLedgerRejected && a.Harness.ProjectiveDomainRejected && a.Harness.PhaseDomainRejected && a.Harness.CausticRejected && a.Harness.BranchTagRejected && a.Harness.MissingUncertaintyRejected && a.Harness.CabibboRayInputRejected && a.Harness.NativePromotionRejected, Detail: FormatHarness(a.Harness)},
			{Name: "does not construct CKM matrix or CKM entry", Passed: a.Harness.NoCKMMatrixConstructed && a.Harness.NoCKMEntryComputed && a.Harness.NoNativePredictionExported, Detail: FormatHarness(a.Harness)},
			{Name: "preserves native firewall", Passed: a.Firewall.Executed && !a.Firewall.SyntheticCoordinatesNative && !a.Firewall.SyntheticDUDNative && !a.Firewall.CKMNativePrediction && !a.Firewall.CKMMatrixConstructed && !a.Firewall.CKMEntryComputed && !a.Firewall.ObservedMassesImported && !a.Firewall.ObservedCKMImported && !a.Firewall.CabibboUsedAsRayInput && !a.Firewall.NativeRegistryWritten && a.Firewall.KGenStillForced && a.Firewall.XTriangleStillForced && a.Firewall.YPhaseStillQuarantined && a.Firewall.SectorCoefficientsStillSealed && a.Firewall.NativeFlavorDimAfter == NativeFlavorDim && a.Firewall.KXYCoeffDimAfter == KXYCoeffDim, Detail: FormatFirewall(a.Firewall)},
			{Name: "next gate observed complete comparator dry-run", Passed: a.Next.Gate == 469, Detail: FormatNext(a.Next)},
		}
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Checks: checks, Notes: []string{StatusSyntheticLedgerAccepted, StatusSymbolicInverseExecuted, StatusIntervalPropagationExecuted, StatusDUDSyntheticComputed, StatusSyntheticInversionValidated, StatusFailedObservedDataRejected, StatusFailedCKMNativePrediction, StatusFirewallPreserved, a.Truth}}
	}}
}
