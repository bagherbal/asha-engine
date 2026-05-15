package generation2eigenbasisledger

import "github.com/bagherbal/asha-engine/pkg/theorem"

func Generation2EigenbasisConventionLedgerMixingMatrixGaugeAuditTheorem() theorem.Theorem {
	const id = AuditID
	const name = "Generation 2 eigenbasis convention ledger mixing-matrix gauge audit"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "construct Gate463 audit", Passed: false, Detail: err.Error()}}}
		}
		checks := []theorem.Check{
			{Name: "inherits Gate462 CKM interface firewall", Passed: a.Inheritance.Executed && a.Inheritance.Gate462SectorDifference && a.Inheritance.Gate462RequiresEigenbasis && a.Inheritance.Gate462RejectsObservedCKMPMNS && a.Inheritance.Gate462RejectsNativePrediction && a.Inheritance.Gate462NoMixingObservableExport && a.Inheritance.NoObservedValuesImported, Detail: FormatInheritance(a.Inheritance)},
			{Name: "audits raw diagonalizer gauge", Passed: a.Gauge.Executed && a.Gauge.RawDiagonalizerPhaseGaugePerSector == PerSectorPhaseGauge && a.Gauge.RawEigenvaluePermutationSheetsPerSector == PerSectorPermutation && a.Gauge.PairPhaseGaugeDimension == PairPhaseGaugeDim && a.Gauge.PairPermutationSheets == PairPermutationSheets && a.Gauge.KGenPreservingRephasingsOnly && a.Gauge.RawDiagonalizersAreNotObservables && a.Gauge.PermutationLabelsAreNotNative && a.Gauge.ConventionCannotCreatePrediction, Detail: FormatGaugeAudit(a.Gauge)},
			{Name: "defines bridge-only eigenbasis convention contract", Passed: a.Contract.Executed && a.Contract.RequiresUSector && a.Contract.RequiresDSector && a.Contract.RequiresKGenBasisDeclaration && a.Contract.RequiresEigenvalueOrdering && a.Contract.RequiresPhaseGauge && a.Contract.RequiresNormalization && a.Contract.RequiresDegeneracyPolicy && a.Contract.RequiresBranchTag && a.Contract.RequiresProvenance && a.Contract.BridgeOnly && a.Contract.ExportsConventionReadinessOnly, Detail: FormatContract(a.Contract)},
			{Name: "sieve accepts only convention-ready bridge pair", Passed: a.Sieve.Executed && a.Sieve.AcceptedCaseCount == 1 && a.Sieve.RejectedCaseCount == 10 && a.Sieve.ValidConventionAccepted && a.Sieve.AllAcceptedBridgeOnly && a.Sieve.NoMixingMatrixExported, Detail: FormatSieve(a.Sieve)},
			{Name: "rejects unsafe gauge and prediction routes", Passed: a.Sieve.MissingSectorRejected && a.Sieve.MissingConventionRejected && a.Sieve.RawPhaseGaugeRejected && a.Sieve.PermutationNativeRejected && a.Sieve.DegenerateSpectrumRejected && a.Sieve.KGenBasisRotationRejected && a.Sieve.ObservedCKMPMNSRejected && a.Sieve.NativePredictionRejected && a.Sieve.EigenbasisNativePromotionRejected && a.Sieve.MatrixExportRejected, Detail: FormatSieve(a.Sieve)},
			{Name: "preserves 13-moduli firewall", Passed: a.Firewall.Executed && a.Firewall.EigenbasisConventionDefined && a.Firewall.CKMNullAdapterMayProceed && !a.Firewall.CKMMatrixEntryComputed && !a.Firewall.CKMMatrixEntryNative && !a.Firewall.PMNSMatrixEntryComputed && !a.Firewall.PMNSMatrixEntryNative && !a.Firewall.ObservedMassesImported && !a.Firewall.ObservedYukawasImported && !a.Firewall.ObservedCKMImported && !a.Firewall.ObservedPMNSImported && a.Firewall.SectorCoefficientsStillSealed && a.Firewall.NativeFlavorDimAfter == NativeFlavorDim && a.Firewall.KXYCoeffDimAfter == KXYCoeffDim, Detail: FormatFirewall(a.Firewall)},
		}
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Checks: checks}
	}}
}
