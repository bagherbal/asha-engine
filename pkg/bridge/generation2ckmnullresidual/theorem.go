package generation2ckmnullresidual

import "github.com/bagherbal/asha-engine/pkg/theorem"

func Generation2CKMNullResidualAdapterConventionReadySymbolicMapTheorem() theorem.Theorem {
	const id = AuditID
	const name = "Generation 2 CKM null residual adapter convention-ready symbolic map"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "construct Gate464 audit", Passed: false, Detail: err.Error()}}}
		}
		checks := []theorem.Check{
			{Name: "inherits Gate462/Gate463 CKM interface firewall", Passed: a.Inheritance.Executed && a.Inheritance.Gate462RelativeRay && a.Inheritance.Gate463EigenbasisLedger && a.Inheritance.Gate463ReadyForResidualAdapter && a.Inheritance.Gate462RejectsObservedCKMPMNS && a.Inheritance.Gate462RejectsNativePrediction && a.Inheritance.NoObservedValuesImported, Detail: FormatInheritance(a.Inheritance)},
			{Name: "defines CKM-null symbolic residual map", Passed: a.Map.Executed && a.Map.RequiresRelativeRay && a.Map.RequiresEigenbasisConvention && a.Map.RequiresOrderingConvention && a.Map.RequiresPhaseGaugeConvention && a.Map.RequiresBranchTags && a.Map.RequiresProvenance && a.Map.RequiresBridgeOnlySyntheticMode && a.Map.RelativeRayDimension == RelativeRayDOF && a.Map.ExportsResidualDiagnosticsOnly && !a.Map.CKMMatrixConstructed && !a.Map.CKMMatrixElementExported, Detail: FormatMap(a.Map)},
			{Name: "sieve accepts only synthetic bridge residual", Passed: a.Sieve.Executed && a.Sieve.AcceptedCaseCount == 1 && a.Sieve.RejectedCaseCount == 11 && a.Sieve.ValidSyntheticResidualAccepted && a.Sieve.AllAcceptedBridgeOnly && a.Sieve.NoCKMMatrixConstructed && a.Sieve.NoNativeObservableExport, Detail: FormatSieve(a.Sieve)},
			{Name: "rejects unsafe CKM residual routes", Passed: a.Sieve.MissingRelativeRayRejected && a.Sieve.MissingEigenbasisRejected && a.Sieve.MissingBranchProvenanceRejected && a.Sieve.ObservedCKMPMNSRejected && a.Sieve.NativePredictionRejected && a.Sieve.MatrixExportRejected && a.Sieve.RawDiagonalizerRejected && a.Sieve.DegenerateSpectrumRejected && a.Sieve.KGenBasisRotationRejected && a.Sieve.NativeResidualPromotionRejected && a.Sieve.GSTSelectorRejected, Detail: FormatSieve(a.Sieve)},
			{Name: "preserves 13-moduli firewall", Passed: a.Firewall.Executed && a.Firewall.CKMNullResidualAdapterDefined && a.Firewall.CKMNullResidualMayRunBridgeOnly && !a.Firewall.CKMMatrixConstructed && !a.Firewall.CKMMatrixEntryComputed && !a.Firewall.CKMMatrixEntryNative && !a.Firewall.PMNSMatrixEntryComputed && !a.Firewall.PMNSMatrixEntryNative && !a.Firewall.ObservedMassesImported && !a.Firewall.ObservedYukawasImported && !a.Firewall.ObservedCKMImported && !a.Firewall.ObservedPMNSImported && !a.Firewall.GSTFritzschPromoted && !a.Firewall.RelativeRayPromotedNative && !a.Firewall.EigenbasisPromotedNative && a.Firewall.SectorCoefficientsStillSealed && a.Firewall.NativeFlavorDimAfter == NativeFlavorDim && a.Firewall.KXYCoeffDimAfter == KXYCoeffDim, Detail: FormatFirewall(a.Firewall)},
		}
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Checks: checks}
	}}
}
