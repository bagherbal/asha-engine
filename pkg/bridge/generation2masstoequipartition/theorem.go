package generation2masstoequipartition

import "github.com/bagherbal/asha-engine/pkg/theorem"

func Generation2MassToEquipartitionInversionEpistemologicalLoopClosureTheorem() theorem.Theorem {
	const id = AuditID
	const name = "Generation 2 mass-to-equipartition inversion epistemological loop closure audit"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "construct Gate473 audit", Passed: false, Detail: err.Error()}}}
		}
		checks := []theorem.Check{
			{Name: "inherits Gate471 socket and Gate454 rank boundary", Passed: a.Inheritance.Executed && a.Inheritance.Gate454RankAuditAvailable && a.Inheritance.Gate465AirlockAvailable && a.Inheritance.Gate471ExternalSocketAvailable && a.Inheritance.NativeRegistryClean, Detail: FormatInheritance(a.Inheritance)},
			{Name: "loads raw masses through airlock without I_K or CKM imports", Passed: a.Import.Executed && a.Import.Loaded && a.Import.EmpiricalImport && a.Import.BridgeOnlyLedger && a.Import.AcceptedRows == 6 && a.Import.MetadataComplete && a.Import.AllAcceptedBridgeOnly && a.Import.IKImportRejected && a.Import.CKMImportRejected && !a.Import.NativeRegistryWriteRequested, Detail: FormatImport(a.Import)},
			{Name: "confirms extreme third-generation hierarchy", Passed: a.Up.ExtremeHierarchy && a.Down.ExtremeHierarchy && a.Up.ThirdGenerationSquareFraction > 0.999 && a.Down.ThirdGenerationSquareFraction > 0.999, Detail: FormatSpectrum(a.Up) + "\n" + FormatSpectrum(a.Down)},
			{Name: "rejects alpha=1 equipartition from raw mass spectra", Passed: !a.Up.AlphaOneForced && !a.Down.AlphaOneForced && !a.Up.IKHalfDerived && !a.Down.IKHalfDerived && !a.Up.AlphaOneCompatible && !a.Down.AlphaOneCompatible, Detail: FormatSpectrum(a.Up) + "\n" + FormatSpectrum(a.Down)},
			{Name: "refuses CKM loop closure without independent I_K and branch tags", Passed: a.Loop.Executed && a.Loop.RawMassesOnly && !a.Loop.AlphaDerived && !a.Loop.IKDerived && !a.Loop.DUDComputed && !a.Loop.CabibboResidualComputed && !a.Loop.AlignmentAchieved && a.Loop.Verdict == StatusFailedProjectNotAchieved, Detail: FormatLoop(a.Loop)},
			{Name: "preserves native theorem firewall", Passed: a.Firewall.Executed && !a.Firewall.RawMassRowsNative && !a.Firewall.IKNative && !a.Firewall.AlphaNative && !a.Firewall.DUDNativePrediction && !a.Firewall.CKMNativePrediction && !a.Firewall.CKMMatrixConstructed && !a.Firewall.CKMEntryComputed && !a.Firewall.NativeRegistryWritten && a.Firewall.KGenStillForced && a.Firewall.XTriangleStillForced && a.Firewall.YPhaseStillQuarantined && a.Firewall.SectorCoefficientsStillSealed && a.Firewall.NativeFlavorDimAfter == NativeFlavorDim && a.Firewall.KXYCoeffDimAfter == KXYCoeffDim, Detail: FormatFirewall(a.Firewall)},
		}
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Checks: checks, Notes: []string{StatusRawMassLedgerLoaded, StatusExtremeHierarchyConfirmed, StatusAsymptoticLimitDerived, StatusFailedMassHierarchyNoEquipartition, StatusFailedRawMassCannotDeriveIK, StatusFailedDUDUndefined, StatusFailedProjectNotAchieved, StatusFirewallPreserved, a.Truth}}
	}}
}
