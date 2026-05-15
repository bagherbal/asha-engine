package generation2observednumericaladapter

import "github.com/bagherbal/asha-engine/pkg/theorem"

func Generation2ObservedNumericalDUDAdapterExplicitDataFileRunTheorem() theorem.Theorem {
	const id = AuditID
	const name = "Generation 2 observed numerical d_ud adapter explicit data-file run"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "construct Gate470 audit", Passed: false, Detail: err.Error()}}}
		}
		checks := []theorem.Check{
			{Name: "inherits Gate469 preflight and Gate465 airlock", Passed: a.Inheritance.Executed && a.Inheritance.Gate469PreflightValidated && a.Inheritance.Gate465AirlockAvailable && a.Inheritance.Gate464DUDSocketAvailable && a.Inheritance.NativeRegistryClean, Detail: FormatInheritance(a.Inheritance)},
			{Name: "loads explicit observed data file into bridge airlock", Passed: a.Import.Executed && a.Import.Loaded && a.Import.EmpiricalImport && a.Import.BridgeOnlyLedger && a.Import.AcceptedRows > 0 && a.Import.AllAcceptedBridgeOnly && a.Import.MetadataComplete && !a.Import.NativeRegistryWriteRequested, Detail: FormatImport(a.Import)},
			{Name: "refuses to fabricate missing ASHA rank-complete comparators", Passed: a.Adapter.Executed && a.Adapter.Attempted && !a.Adapter.DUDComputed && a.Adapter.MissingISpecIKValues && a.Adapter.MissingBranchTags && a.Adapter.PDGNoIK && a.Adapter.Verdict == StatusFailedDUDNotComputableFromFile, Detail: FormatAdapter(a.Adapter)},
			{Name: "keeps Cabibbo as residual target only", Passed: a.Adapter.CabibboTargetAvailable && !a.Adapter.CabibboResidualComputed && !a.Adapter.AlignmentAchieved, Detail: FormatAdapter(a.Adapter)},
			{Name: "preserves native theorem firewall", Passed: a.Firewall.Executed && !a.Firewall.DataFileRowsNative && !a.Firewall.CoordinatesNative && !a.Firewall.DUDNativePrediction && !a.Firewall.CKMNativePrediction && !a.Firewall.CKMMatrixConstructed && !a.Firewall.CKMEntryComputed && !a.Firewall.CabibboUsedAsRayInput && !a.Firewall.NativeRegistryWritten && a.Firewall.KGenStillForced && a.Firewall.XTriangleStillForced && a.Firewall.YPhaseStillQuarantined && a.Firewall.SectorCoefficientsStillSealed && a.Firewall.NativeFlavorDimAfter == NativeFlavorDim && a.Firewall.KXYCoeffDimAfter == KXYCoeffDim, Detail: FormatFirewall(a.Firewall)},
		}
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Checks: checks, Notes: []string{StatusDataFileLoaded, StatusAirlockAcceptedBridgeRows, StatusFailedPDGNoIK, StatusFailedDUDNotComputableFromFile, StatusFailedCabibboResidualUndefined, StatusFirewallPreserved, a.Truth}}
	}}
}
