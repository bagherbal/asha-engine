package generation2rankcompleteexternalledger

import "github.com/bagherbal/asha-engine/pkg/theorem"

func Generation2RankCompleteExternalLedgerAcceptanceTestTheorem() theorem.Theorem {
	const id = AuditID
	const name = "Generation 2 rank-complete external ledger acceptance test"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "construct Gate471 audit", Passed: false, Detail: err.Error()}}}
		}
		checks := []theorem.Check{
			{Name: "inherits Gate470 non-smuggling boundary and Gate465 airlock", Passed: a.Inheritance.Executed && a.Inheritance.Gate470NonSmugglingValidated && a.Inheritance.Gate465AirlockAvailable && a.Inheritance.Gate464DUDSocketAvailable && a.Inheritance.NativeRegistryClean, Detail: FormatInheritance(a.Inheritance)},
			{Name: "loads explicit rank-complete ledger into bridge airlock", Passed: a.Import.Executed && a.Import.Loaded && a.Import.EmpiricalImport && a.Import.BridgeOnlyLedger && a.Import.AcceptedRows > 0 && a.Import.AllAcceptedBridgeOnly && a.Import.MetadataComplete && !a.Import.NativeRegistryWriteRequested, Detail: FormatImport(a.Import)},
			{Name: "computes cylinder coordinates from explicit I_spec/I_K/branch tags", Passed: a.Adapter.Executed && a.Adapter.CoordinatesComputed && a.Adapter.URay.Defined && a.Adapter.DRay.Defined && !a.Adapter.MissingISpecIKValues && !a.Adapter.MissingBranchTags, Detail: FormatAdapter(a.Adapter) + "\n" + FormatRay(a.Adapter.URay) + "\n" + FormatRay(a.Adapter.DRay)},
			{Name: "computes bridge-only d_ud and Cabibbo residual", Passed: a.Adapter.DUDComputed && a.Adapter.CabibboTargetAvailable && a.Adapter.CabibboResidualComputed && a.Adapter.AlignmentAchieved, Detail: FormatAdapter(a.Adapter)},
			{Name: "quarantines external I_K and branch tags from PDG/native law-space", Passed: a.Import.PDGIKClaimRejected && !a.Adapter.U.ClaimsPDGPublishesIK && !a.Adapter.D.ClaimsPDGPublishesIK, Detail: StatusExternalComparatorsNotPDGNative},
			{Name: "preserves native theorem firewall", Passed: a.Firewall.Executed && !a.Firewall.DataFileRowsNative && !a.Firewall.CoordinatesNative && !a.Firewall.DUDNativePrediction && !a.Firewall.CKMNativePrediction && !a.Firewall.CKMMatrixConstructed && !a.Firewall.CKMEntryComputed && !a.Firewall.CabibboUsedAsRayInput && !a.Firewall.NativeRegistryWritten && a.Firewall.KGenStillForced && a.Firewall.XTriangleStillForced && a.Firewall.YPhaseStillQuarantined && a.Firewall.SectorCoefficientsStillSealed && a.Firewall.NativeFlavorDimAfter == NativeFlavorDim && a.Firewall.KXYCoeffDimAfter == KXYCoeffDim, Detail: FormatFirewall(a.Firewall)},
		}
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Checks: checks, Notes: []string{StatusRankCompleteLedgerLoaded, StatusAirlockAcceptedRankCompleteRows, StatusCoordinatesComputed, StatusDUDComputed, StatusCabibboResidualComputed, StatusCKMGeometricAlignmentAchieved, StatusExternalComparatorsNotPDGNative, StatusFirewallPreserved, a.Truth}}
	}}
}
