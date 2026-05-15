package generation2leptonempiricalimportswitch

import "github.com/bagherbal/asha-engine/pkg/theorem"

func Generation2LeptonSectorEmpiricalImportSwitchPMNSDataFirewallTheorem() theorem.Theorem {
	const id = AuditID
	const name = "Generation 2 lepton-sector empirical import switch PMNS data firewall"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "construct Gate477 audit", Passed: false, Detail: err.Error()}}}
		}
		checks := []theorem.Check{
			{Name: "inherits Gate476 PMNS-null firewall", Passed: a.Inheritance.Executed && a.Inheritance.Gate476PMNSNullResidualAdapter && a.Inheritance.Gate476SyntheticOnly && a.Inheritance.Gate476RejectsObservedPMNS && a.Inheritance.Gate476RejectsPMNSAsRayInput && a.Inheritance.Gate476RejectsNativePrediction && a.Inheritance.Gate476RejectsMatrixExport && a.Inheritance.Gate476DiagnosticOnly && a.Inheritance.NativeRegistryCleanBeforeAirlock, Detail: FormatInheritance(a.Inheritance)},
			{Name: "defines explicit lepton empirical_import airlock", Passed: a.Policy.Executed && a.Policy.StateVariableName == "empirical_import" && !a.Policy.DefaultEmpiricalImport && a.Policy.RequiresExplicitTrue && a.Policy.RequiresSource && a.Policy.RequiresScale && a.Policy.RequiresScheme && a.Policy.RequiresUncertainty && a.Policy.RequiresBridgeOnlyQuarantine && a.Policy.RequiresNeutrinoOrderingPolicy && a.Policy.RequiresAbsoluteNeutrinoScalePolicy && a.Policy.RequiresMajoranaDiracPhasePolicy && a.Policy.AllowsPMNSResidualTarget && !a.Policy.AllowsPMNSAsRayInput && a.Policy.AllowedLedger == ComparatorLedger && a.Policy.RequiredMetadataCount == RequiredMetadataCount, Detail: FormatPolicy(a.Policy)},
			{Name: "sieve accepts only quarantined lepton imports", Passed: a.Sieve.Executed && a.Sieve.AcceptedCaseCount == 3 && a.Sieve.RejectedCaseCount == 14 && a.Sieve.QuarantinedChargedLeptonImportAccepted && a.Sieve.QuarantinedNeutrinoImportAccepted && a.Sieve.QuarantinedPMNSResidualTargetAccepted && a.Sieve.AllAcceptedQuarantined && a.Sieve.NoAcceptedNativeRegistryWrite && a.Sieve.NoAcceptedPMNSAsRayInput, Detail: FormatSieve(a.Sieve)},
			{Name: "rejects unsafe lepton/PMNS import routes", Passed: a.Sieve.ClosedSwitchRejected && a.Sieve.MissingMetadataRejected && a.Sieve.MissingUncertaintyRejected && a.Sieve.MissingBridgeOnlyRejected && a.Sieve.UnsupportedLedgerRejected && a.Sieve.MissingLeptonPoliciesRejected && a.Sieve.PMNSAsRayInputRejected && a.Sieve.NativePromotionRejected && a.Sieve.NativeRegistryWriteRejected && a.Sieve.PMNSNativePredictionRejected && a.Sieve.ObservedDataAsTheoremRejected, Detail: FormatSieve(a.Sieve)},
			{Name: "preserves native 13-moduli firewall with lepton airlock open", Passed: a.Firewall.Executed && a.Firewall.AirlockCanOpen && a.Firewall.EmpiricalRowsImported == 3 && a.Firewall.AllImportedRowsQuarantined && !a.Firewall.EmpiricalDataInNativeRegistry && !a.Firewall.NativePredictionFromEmpirical && !a.Firewall.NativeLawFromEmpirical && !a.Firewall.ObservedDataUsedAsTheoremInput && !a.Firewall.PMNSMatrixNativePrediction && !a.Firewall.PMNSMatrixConstructed && !a.Firewall.PMNSEntryComputed && !a.Firewall.LeptonMassNativePrediction && !a.Firewall.NeutrinoMassNativePrediction && !a.Firewall.IKNativeSelectorFound && !a.Firewall.DENuComputedFromObserved && !a.Firewall.DENuNativePrediction && a.Firewall.SectorCoefficientsStillSealed && a.Firewall.NativeFlavorDimAfter == NativeFlavorDim && a.Firewall.KXYCoeffDimAfter == KXYCoeffDim, Detail: FormatFirewall(a.Firewall)},
		}
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Checks: checks, Notes: []string{StatusGate476Inherited, StatusLeptonAirlockDefined, StatusQuarantinedLeptonAccepted, StatusLeptonImportSwitchValid, StatusFailedPMNSAsRayInput, StatusFailedNativePromotion, StatusFailedPMNSNativePrediction, StatusFirewallPreserved, a.Truth}}
	}}
}
