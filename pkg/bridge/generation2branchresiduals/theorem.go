package generation2branchresiduals

import "github.com/bagherbal/asha-engine/pkg/theorem"

func Generation2BranchResolvedTextureResidualHarnessSyntheticNullMapTheorem() theorem.Theorem {
	const id = AuditID
	const name = "Generation 2 branch-resolved texture residual harness synthetic null map"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "construct Gate460 audit", Passed: false, Detail: err.Error()}}}
		}
		checks := []theorem.Check{
			{Name: "inherits inverse, provenance, evaluator, and branch tag ledgers", Passed: a.Inheritance.Executed && a.Inheritance.Gate444KGenForced && a.Inheritance.Gate445TriangleForced && a.Inheritance.Gate456InverseDerived && a.Inheritance.Gate457ProvenanceContractDefined && a.Inheritance.Gate458SyntheticHarnessDefined && a.Inheritance.Gate459BranchTagLedgerDefined && a.Inheritance.Gate459RequiresCPOddSign && a.Inheritance.Gate459RequiresC3Sheet && a.Inheritance.Gate459CompleteTagUnique && a.Inheritance.NativeCPSelectorAbsent && a.Inheritance.NativeC3SheetSelectorAbsent && a.Inheritance.NoObservedValuesImported, Detail: FormatInheritance(a.Inheritance)},
			{Name: "defines branch-resolved synthetic/null residual harness", Passed: a.Harness.Executed && a.Harness.ComposesGate456Inverse && a.Harness.ComposesGate458Evaluator && a.Harness.ComposesGate459BranchTags && a.Harness.RequiresCompleteBranchTag && a.Harness.ComputesProjectiveRay && a.Harness.ComputesTextureZeroResidual && a.Harness.ComputesComparatorResiduals && a.Harness.ComputesPhaseTagResiduals && a.Harness.SyntheticOnly && a.Harness.RedactedAllowedUnevaluated && a.Harness.ObservedDataRejected && a.Harness.BridgeOnlyExport, Detail: FormatHarness(a.Harness)},
			{Name: "residual ledger is diagnostic and bridge-only", Passed: a.Ledger.Executed && a.Ledger.ResidualsBridgeOnly && !a.Ledger.ResidualsNativeObservable, Detail: FormatLedger(a.Ledger)},
			{Name: "accepts only redacted and complete synthetic branch records", Passed: a.Sieve.Executed && a.Sieve.AcceptedCount == 2 && a.Sieve.RedactedPreserved && a.Sieve.SyntheticInteriorAccepted && a.Sieve.AllAcceptedBridgeOnly && a.Sieve.AllResidualsDiagnosticOnly && a.Sieve.NoNativeFlavorObservableOut, Detail: FormatSieve(a.Sieve)},
			{Name: "unsafe residual records fail closed", Passed: a.Sieve.RejectedCount == 6 && a.Sieve.IncompleteTagRejected && a.Sieve.CausticRejected && a.Sieve.ObservedDataRejected && a.Sieve.NativePromotionRejected && a.Sieve.ProjectiveDomainRejected && a.Sieve.PhaseCosDomainRejected, Detail: FormatSieve(a.Sieve)},
			{Name: "preserves flavor firewall", Passed: a.Firewall.Executed && a.Firewall.NoObservedMuonMassImported && a.Firewall.NoObservedCharmMassImported && a.Firewall.NoObservedYukawaImported && a.Firewall.NoCKMImported && a.Firewall.NoPMNSImported && a.Firewall.NoGSTPromotion && a.Firewall.NoCoefficientRayPromotion && a.Firewall.NoPhaseBranchPromotion && a.Firewall.NoCurveFitPromoted && a.Firewall.KGenStillForced && a.Firewall.XTriangleStillForced && a.Firewall.YPhaseStillQuarantined && a.Firewall.SectorCoefficientsStillSealed && a.Firewall.NativeFlavorDimAfter == NativeFlavorDim && a.Firewall.KXYCoeffDimAfter == KXYCoeffDim, Detail: FormatFirewall(a.Firewall)},
			{Name: "next gate audits sector multiplexing", Passed: a.Next.Gate == 461, Detail: FormatNext(a.Next)},
		}
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Checks: checks, Notes: append(statuses(), a.Truth)}
	}}
}
