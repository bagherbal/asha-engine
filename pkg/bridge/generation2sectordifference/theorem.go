package generation2sectordifference

import "github.com/bagherbal/asha-engine/pkg/theorem"

func Generation2SectorDifferenceInvariantCKMInterfaceFirewallAuditTheorem() theorem.Theorem {
	const id = AuditID
	const name = "Generation 2 sector-difference invariant CKM interface firewall audit"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "construct Gate462 audit", Passed: false, Detail: err.Error()}}}
		}
		checks := []theorem.Check{
			{Name: "inherits Gate461 sector-multiplex firewall", Passed: a.Inheritance.Executed && a.Inheritance.Gate461SectorMultiplex && a.Inheritance.Gate461IndependentSectorRaysAccepted && a.Inheritance.Gate461NativeUniversalityRejected && a.Inheritance.Gate461SectorContaminationRejected && a.Inheritance.NoObservedValuesImported, Detail: FormatInheritance(a.Inheritance)},
			{Name: "defines u-d relative-ray interface contract", Passed: a.Contract.Executed && a.Contract.RequiresUSector && a.Contract.RequiresDSector && a.Contract.RequiresProvenancePerSector && a.Contract.RequiresCompleteBranchTags && a.Contract.RequiresEigenbasisConvention && a.Contract.RelativeRayDimension == RelativeRayDOF && a.Contract.ExportsRelativeDiagnosticsOnly && a.Contract.RejectsCKMAsNativePrediction && a.Contract.RejectsPMNSInChargedCKMLedger && a.Contract.RejectsObservedMixingByDefault && a.Contract.RejectsNativeRelativeRayPromotion, Detail: FormatContract(a.Contract)},
			{Name: "sieve accepts only bridge-only synthetic u-d difference", Passed: a.Sieve.Executed && a.Sieve.AcceptedCaseCount == 1 && a.Sieve.RejectedCaseCount == 8 && a.Sieve.ValidUDDifferenceAccepted && a.Sieve.AllAcceptedBridgeOnly && a.Sieve.NoNativeMixingObservableExport, Detail: FormatSieve(a.Sieve)},
			{Name: "rejects unsafe CKM/PMNS and promotion routes", Passed: a.Sieve.MissingSectorRejected && a.Sieve.MissingProvenanceRejected && a.Sieve.MissingEigenbasisRejected && a.Sieve.ObservedCKMPMNSRejected && a.Sieve.NativePredictionRejected && a.Sieve.NativeRelativePromotionRejected && a.Sieve.LeptonPMNSMisrouteRejected && a.Sieve.UniversalityNativeRejected, Detail: FormatSieve(a.Sieve)},
			{Name: "preserves 13-moduli firewall", Passed: a.Firewall.Executed && a.Firewall.RelativeRayMayFeedCKMAdapter && !a.Firewall.CKMMatrixEntryComputed && !a.Firewall.CKMMatrixEntryNative && !a.Firewall.PMNSMatrixEntryComputed && !a.Firewall.PMNSMatrixEntryNative && a.Firewall.NoObservedMassesImported && a.Firewall.NoObservedYukawasImported && a.Firewall.NoObservedCKMImported && a.Firewall.NoObservedPMNSImported && a.Firewall.SectorCoefficientsStillSealed && a.Firewall.NativeFlavorDimAfter == NativeFlavorDim && a.Firewall.KXYCoeffDimAfter == KXYCoeffDim, Detail: FormatFirewall(a.Firewall)},
		}
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Checks: checks}
	}}
}
