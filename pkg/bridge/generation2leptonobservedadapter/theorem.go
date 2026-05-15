package generation2leptonobservedadapter

import "github.com/bagherbal/asha-engine/pkg/theorem"

func Generation2LeptonObservedComparatorAdapterPMNSAirlockNonComputationTheorem() theorem.Theorem {
	const id = AuditID
	const name = "Generation 2 lepton observed comparator adapter PMNS airlock non-computation"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "construct Gate478 audit", Passed: false, Detail: err.Error()}}}
		}
		checks := []theorem.Check{
			{Name: "inherits Gate477 lepton airlock and Gate476 socket", Passed: a.Inheritance.Executed && a.Inheritance.Gate477LeptonAirlockAvailable && a.Inheritance.Gate476DENuSocketAvailable && a.Inheritance.Gate475LeptonPreflightValidated && a.Inheritance.Gate456InverseAvailable && a.Inheritance.Gate459BranchTagsRequired && a.Inheritance.NativeRegistryClean, Detail: FormatInheritance(a.Inheritance)},
			{Name: "loads explicit observed lepton data file into bridge airlock", Passed: a.Import.Executed && a.Import.Loaded && a.Import.EmpiricalImport && a.Import.BridgeOnlyLedger && a.Import.AcceptedRows > 0 && a.Import.AllAcceptedBridgeOnly && a.Import.MetadataComplete && a.Import.LeptonPoliciesComplete && !a.Import.NativeRegistryWriteRequested, Detail: FormatImport(a.Import)},
			{Name: "refuses to fabricate missing ASHA lepton rank-complete comparators", Passed: a.Adapter.Executed && a.Adapter.Attempted && !a.Adapter.DENuComputed && a.Adapter.MissingISpecIKValues && a.Adapter.MissingBranchTags && a.Adapter.LeptonDataNoIK && a.Adapter.Verdict == StatusFailedDENuNotComputableFromFile, Detail: FormatAdapter(a.Adapter)},
			{Name: "keeps PMNS as residual target only", Passed: a.Adapter.PMNSTargetAvailable && !a.Adapter.PMNSResidualComputed && !a.Adapter.AlignmentAchieved, Detail: FormatAdapter(a.Adapter)},
			{Name: "preserves native theorem firewall", Passed: a.Firewall.Executed && !a.Firewall.DataFileRowsNative && !a.Firewall.CoordinatesNative && !a.Firewall.DENuNativePrediction && !a.Firewall.PMNSNativePrediction && !a.Firewall.PMNSMatrixConstructed && !a.Firewall.PMNSEntryComputed && !a.Firewall.PMNSUsedAsRayInput && !a.Firewall.NativeRegistryWritten && a.Firewall.KGenStillForced && a.Firewall.XTriangleStillForced && a.Firewall.YPhaseStillQuarantined && a.Firewall.SectorCoefficientsStillSealed && a.Firewall.NativeFlavorDimAfter == NativeFlavorDim && a.Firewall.KXYCoeffDimAfter == KXYCoeffDim, Detail: FormatFirewall(a.Firewall)},
		}
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Checks: checks, Notes: []string{StatusDataFileLoaded, StatusAirlockAcceptedBridgeRows, StatusFailedLeptonDataNoIK, StatusFailedDENuNotComputableFromFile, StatusFailedPMNSResidualUndefined, StatusFailedPMNSAsRayInput, StatusFailedPMNSNativePrediction, StatusFailedPMNSMatrixExport, StatusFirewallPreserved, a.Truth}}
	}}
}
