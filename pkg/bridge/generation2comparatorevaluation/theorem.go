package generation2comparatorevaluation

import "github.com/bagherbal/asha-engine/pkg/theorem"

func Generation2ComparatorLedgerEvaluationHarnessRedactedPhenomenologySlotTheorem() theorem.Theorem {
	const id = AuditID
	const name = "Generation 2 comparator ledger evaluation harness redacted phenomenology slot"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "construct Gate458 audit", Passed: false, Detail: err.Error()}}}
		}
		checks := []theorem.Check{
			{Name: "inherits Gate456 inverse and Gate457 provenance firewall", Passed: a.Inheritance.Executed && a.Inheritance.Gate456InverseDerived && a.Inheritance.Gate456BridgeOnly && a.Inheritance.Gate456GenericBranchCount == GenericPhaseBranchCount && a.Inheritance.Gate457ProvenanceContractDefined && a.Inheritance.Gate457RequiredFields == Gate457RequiredFields && a.Inheritance.Gate457BridgeOnly && a.Inheritance.Gate457ObservedImportExplicitOnly && a.Inheritance.NativeCoefficientRaySelectorAbsent && a.Inheritance.NoObservedValuesImported, Detail: FormatInheritance(a.Inheritance)},
			{Name: "defines redacted/synthetic evaluation harness", Passed: a.Harness.Executed && a.Harness.AcceptsOnlyGate457ValidInput && a.Harness.SyntheticModeAllowed && a.Harness.RedactedModeAllowed && a.Harness.ObservedNumericModeRejected && a.Harness.UsesGate456Inverse && a.Harness.ComputesAlpha && a.Harness.ComputesCos3Phi && a.Harness.ComputesBranchDiagnostics && a.Harness.ComputesDomainGuards && a.Harness.BridgeOnlyOutput, Detail: FormatHarness(a.Harness)},
			{Name: "evaluates only safe bridge comparator records", Passed: a.Sieve.Executed && a.Sieve.AcceptedCount == 2 && a.Sieve.RedactedAccepted && a.Sieve.SyntheticInteriorAccepted && a.Sieve.AllAcceptedBridgeOnly && a.Sieve.NoNativeRayExport, Detail: FormatSieve(a.Sieve)},
			{Name: "rejects observed/domain/native unsafe routes", Passed: a.Sieve.RejectedCount == 5 && a.Sieve.SyntheticCausticFlagged && a.Sieve.ObservedValueRejected && a.Sieve.IKDomainRejected && a.Sieve.PhaseCosDomainRejected && a.Sieve.NativePromotionRejected, Detail: FormatSieve(a.Sieve)},
			{Name: "preserves flavor firewall", Passed: a.Firewall.Executed && a.Firewall.NoObservedMuonMassImported && a.Firewall.NoObservedCharmMassImported && a.Firewall.NoObservedYukawaImported && a.Firewall.NoCKMImported && a.Firewall.NoPMNSImported && a.Firewall.NoGSTPromotion && a.Firewall.NoCoefficientRayPromotion && a.Firewall.NoCurveFitPromoted && a.Firewall.KGenStillForced && a.Firewall.XTriangleStillForced && a.Firewall.YPhaseStillQuarantined && a.Firewall.SectorCoefficientsStillSealed && a.Firewall.NativeFlavorDimAfter == NativeFlavorDim && a.Firewall.KXYCoeffDimAfter == KXYCoeffDim, Detail: FormatFirewall(a.Firewall)},
			{Name: "next gate targets oriented branch tags", Passed: a.Next.Gate == 459, Detail: FormatNext(a.Next)},
		}
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Checks: checks, Notes: append(statuses(), a.Truth)}
	}}
}
