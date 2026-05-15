package generation2branchtags

import "github.com/bagherbal/asha-engine/pkg/theorem"

func Generation2OrientedComparatorBranchTagSieveCPSignLedgerTheorem() theorem.Theorem {
	const id = AuditID
	const name = "Generation 2 oriented comparator branch tag sieve CP-sign ledger"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "construct Gate459 audit", Passed: false, Detail: err.Error()}}}
		}
		checks := []theorem.Check{
			{Name: "inherits Gate456 inverse and Gate458 redacted harness", Passed: a.Inheritance.Executed && a.Inheritance.Gate444KGenForced && a.Inheritance.Gate445TriangleForced && a.Inheritance.Gate456InverseDerived && a.Inheritance.Gate456GenericBranchCount == CosineOnlyBranchCount && a.Inheritance.Gate457ProvenanceContractDefined && a.Inheritance.Gate458RedactedHarnessDefined && a.Inheritance.Gate458ObservedValuesRejected && a.Inheritance.Gate458BridgeOnly && a.Inheritance.NativeCPSelectorAbsent && a.Inheritance.NativeC3SheetSelectorAbsent && a.Inheritance.NoObservedValuesImported, Detail: FormatInheritance(a.Inheritance)},
			{Name: "defines necessary and sufficient branch-tag ledger", Passed: a.Ledger.Executed && a.Ledger.RequiresCosineInvariant && a.Ledger.RequiresCPOddSign && a.Ledger.RequiresC3Sheet && a.Ledger.CosineOnlyBranchCount == CosineOnlyBranchCount && a.Ledger.CPOddSignOnlyBranchCount == CPSignOnlyBranchCount && a.Ledger.CompleteBranchTagCount == CompleteBranchCount && a.Ledger.RejectsCKMOrPMNSAsSelector && a.Ledger.RejectsNativePromotion && a.Ledger.BridgeOnly && a.Ledger.NativeCPOddSignSelectorAbsent && a.Ledger.NativeC3SheetSelectorAbsent, Detail: FormatLedger(a.Ledger)},
			{Name: "complete bridge branch tags select unique synthetic phase", Passed: a.Sieve.Executed && a.Sieve.AcceptedCount == 2 && a.Sieve.CompletePositiveAccepted && a.Sieve.CompleteNegativeAccepted && a.Sieve.AllAcceptedBridgeOnly && a.Sieve.NoNativePhaseExport, Detail: FormatSieve(a.Sieve)},
			{Name: "incomplete and unsafe branch selectors fail closed", Passed: a.Sieve.RejectedCount == 5 && a.Sieve.CosineOnlyFlagged && a.Sieve.CPOddOnlyFlagged && a.Sieve.CKMPMNSSelectorRejected && a.Sieve.NativePromotionRejected && a.Sieve.InvalidTagRejected, Detail: FormatSieve(a.Sieve)},
			{Name: "preserves flavor firewall", Passed: a.Firewall.Executed && a.Firewall.NoObservedMuonMassImported && a.Firewall.NoObservedCharmMassImported && a.Firewall.NoObservedYukawaImported && a.Firewall.NoCKMImported && a.Firewall.NoPMNSImported && a.Firewall.NoGSTPromotion && a.Firewall.NoCoefficientRayPromotion && a.Firewall.NoCPPhasePromotion && a.Firewall.NoCurveFitPromoted && a.Firewall.KGenStillForced && a.Firewall.XTriangleStillForced && a.Firewall.YPhaseStillQuarantined && a.Firewall.SectorCoefficientsStillSealed && a.Firewall.NativeFlavorDimAfter == NativeFlavorDim && a.Firewall.KXYCoeffDimAfter == KXYCoeffDim, Detail: FormatFirewall(a.Firewall)},
			{Name: "next gate composes branch-resolved residual harness", Passed: a.Next.Gate == 460, Detail: FormatNext(a.Next)},
		}
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Checks: checks, Notes: append(statuses(), a.Truth)}
	}}
}
