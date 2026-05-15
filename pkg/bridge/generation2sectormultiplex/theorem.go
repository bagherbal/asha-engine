package generation2sectormultiplex

import "github.com/bagherbal/asha-engine/pkg/theorem"

func Generation2ThreeSectorComparatorMultiplexUniversalityAssumptionAuditTheorem() theorem.Theorem {
	const id = AuditID
	const name = "Generation 2 three-sector comparator multiplex universality assumption audit"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "construct Gate461 audit", Passed: false, Detail: err.Error()}}}
		}
		checks := []theorem.Check{
			{Name: "inherits Gate460 residual-harness firewall", Passed: a.Inheritance.Executed && a.Inheritance.Gate444KGenForced && a.Inheritance.Gate445TriangleForced && a.Inheritance.Gate456InverseDerived && a.Inheritance.Gate457ProvenanceContract && a.Inheritance.Gate458ComparatorHarness && a.Inheritance.Gate459BranchTags && a.Inheritance.Gate460ResidualHarness && a.Inheritance.Gate460ResidualsBridgeOnly && a.Inheritance.Gate460ObservedRejected && a.Inheritance.Gate460NativePromotionBlocked && a.Inheritance.NoObservedValuesImported, Detail: FormatInheritance(a.Inheritance)},
			{Name: "defines sector-indexed multiplex contract", Passed: a.Contract.Executed && a.Contract.SectorIndexed && a.Contract.IndependentRayPerSector && a.Contract.NoImplicitRaySharing && a.Contract.NoImplicitPhaseSharing && a.Contract.NoImplicitBranchTagSharing && a.Contract.RequiresProvenancePerRow && a.Contract.RequiresCompleteBranchTags && a.Contract.AllowsLabelledBridgeOnlyUniversality && a.Contract.RejectsNativeUniversality && a.Contract.BridgeOnlyExport, Detail: FormatContract(a.Contract)},
			{Name: "preserves 9 charged K/X/Y coefficient ledger", Passed: a.Dimensions.Executed && a.Dimensions.TotalKXYChargedCoefficients == KXYCoeffDim && a.Dimensions.CoefficientsPerSector == 3 && a.Dimensions.NativeChargedFlavorDimBefore == NativeFlavorDim && a.Dimensions.NativeChargedFlavorDimAfter == NativeFlavorDim && a.Dimensions.UniversalityWouldReduceBridgeDOF && !a.Dimensions.UniversalityReductionNative && !a.Dimensions.SectorRayUniversalityNative, Detail: FormatDimensions(a.Dimensions)},
			{Name: "accepts independent and labelled bridge-only multiplex cases", Passed: a.Sieve.Executed && a.Sieve.AcceptedCaseCount == 2 && a.Sieve.IndependentThreeSectorAccepted && a.Sieve.LabelledBridgeUniversalityAccepted && a.Sieve.AllAcceptedBridgeOnly && a.Sieve.NoNativeObservableExport, Detail: FormatSieve(a.Sieve)},
			{Name: "rejects unsafe sector multiplex routes", Passed: a.Sieve.RejectedCaseCount == 6 && a.Sieve.MissingSectorRejected && a.Sieve.NativeUniversalityRejected && a.Sieve.UnlabelledUniversalityRejected && a.Sieve.ObservedDataRejected && a.Sieve.NativePromotionRejected && a.Sieve.SectorContaminationRejected, Detail: FormatSieve(a.Sieve)},
			{Name: "preserves flavor firewall and rejects native sector universality", Passed: a.Firewall.Executed && a.Firewall.NoObservedMuonMassImported && a.Firewall.NoObservedCharmMassImported && a.Firewall.NoObservedTopBottomImported && a.Firewall.NoObservedYukawaImported && a.Firewall.NoCKMImported && a.Firewall.NoPMNSImported && a.Firewall.NoGSTPromotion && a.Firewall.NoCoefficientRayPromotion && a.Firewall.NoCrossSectorUniversalityLaw && a.Firewall.NoCurveFitPromoted && a.Firewall.KGenStillForced && a.Firewall.XTriangleStillForced && a.Firewall.YPhaseStillQuarantined && a.Firewall.SectorCoefficientsStillSealed && a.Firewall.NativeFlavorDimAfter == NativeFlavorDim && a.Firewall.KXYCoeffDimAfter == KXYCoeffDim, Detail: FormatFirewall(a.Firewall)},
			{Name: "next gate audits sector-difference CKM interface", Passed: a.Next.Gate == 462, Detail: FormatNext(a.Next)},
		}
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Checks: checks, Notes: append(statuses(), a.Truth)}
	}}
}
