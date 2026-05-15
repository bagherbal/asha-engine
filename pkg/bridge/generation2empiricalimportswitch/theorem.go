package generation2empiricalimportswitch

import "github.com/bagherbal/asha-engine/pkg/theorem"

func Generation2QuarkSectorEmpiricalImportSwitchCKMDataFirewallTheorem() theorem.Theorem {
	const id = AuditID
	const name = "Generation 2 quark-sector empirical import switch CKM data firewall"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "construct Gate465 audit", Passed: false, Detail: err.Error()}}}
		}
		checks := []theorem.Check{
			{Name: "inherits Gate464 CKM-null firewall", Passed: a.Inheritance.Executed && a.Inheritance.Gate464CKMNullResidualAdapter && a.Inheritance.Gate464RejectsObservedByDefault && a.Inheritance.Gate464RejectsNativePrediction && a.Inheritance.Gate464RejectsMatrixExport && a.Inheritance.Gate464DiagnosticOnly && a.Inheritance.NativeRegistryCleanBeforeAirlock, Detail: FormatInheritance(a.Inheritance)},
			{Name: "defines explicit empirical_import airlock", Passed: a.Policy.Executed && a.Policy.StateVariableName == "empirical_import" && !a.Policy.DefaultEmpiricalImport && a.Policy.RequiresExplicitTrue && a.Policy.RequiresSource && a.Policy.RequiresScale && a.Policy.RequiresScheme && a.Policy.RequiresUncertainty && a.Policy.RequiresBridgeOnlyQuarantine && a.Policy.AllowedLedger == ComparatorLedger && a.Policy.RequiredMetadataCount == RequiredMetadataCount, Detail: FormatPolicy(a.Policy)},
			{Name: "sieve accepts only quarantined imports", Passed: a.Sieve.Executed && a.Sieve.AcceptedCaseCount == 2 && a.Sieve.RejectedCaseCount == 12 && a.Sieve.QuarantinedQuarkMassImportAccepted && a.Sieve.QuarantinedCKMImportAccepted && a.Sieve.AllAcceptedQuarantined && a.Sieve.NoAcceptedNativeRegistryWrite, Detail: FormatSieve(a.Sieve)},
			{Name: "rejects unsafe empirical import routes", Passed: a.Sieve.ClosedSwitchRejected && a.Sieve.MissingMetadataRejected && a.Sieve.MissingUncertaintyRejected && a.Sieve.MissingBridgeOnlyRejected && a.Sieve.UnsupportedLedgerRejected && a.Sieve.NativePromotionRejected && a.Sieve.NativeRegistryWriteRejected && a.Sieve.CKMPMNSNativePredictionRejected && a.Sieve.ObservedDataAsTheoremRejected, Detail: FormatSieve(a.Sieve)},
			{Name: "preserves native 13-moduli firewall with airlock open", Passed: a.Firewall.Executed && a.Firewall.AirlockCanOpen && a.Firewall.EmpiricalRowsImported == 2 && a.Firewall.AllImportedRowsQuarantined && !a.Firewall.EmpiricalDataInNativeRegistry && !a.Firewall.NativePredictionFromEmpirical && !a.Firewall.NativeLawFromEmpirical && !a.Firewall.ObservedDataUsedAsTheoremInput && !a.Firewall.CKMMatrixNativePrediction && !a.Firewall.PMNSMatrixNativePrediction && !a.Firewall.CKMMatrixConstructed && !a.Firewall.CKMEntryComputed && !a.Firewall.QuarkMassNativePrediction && !a.Firewall.YukawaNativePrediction && a.Firewall.SectorCoefficientsStillSealed && a.Firewall.NativeFlavorDimAfter == NativeFlavorDim && a.Firewall.KXYCoeffDimAfter == KXYCoeffDim, Detail: FormatFirewall(a.Firewall)},
		}
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Checks: checks}
	}}
}
