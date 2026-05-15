package generation2observedcomparatoradapter

import "github.com/bagherbal/asha-engine/pkg/theorem"

func Generation2QuarkSectorObservedComparatorAdapterCKMDataFirewallTheorem() theorem.Theorem {
	const id = AuditID
	const name = "Generation 2 quark-sector observed comparator adapter CKM data firewall"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "construct Gate466 audit", Passed: false, Detail: err.Error()}}}
		}
		checks := []theorem.Check{
			{Name: "inherits Gate465 airlock and Gate464 socket", Passed: a.Inheritance.Executed && a.Inheritance.Gate465Airlock && a.Inheritance.Gate464CKMNullSocket && a.Inheritance.Gate454SpectrumOnlyRankOne && a.Inheritance.Gate459BranchTagsRequired && a.Inheritance.Gate465RejectsNativeWrites, Detail: FormatInheritance(a.Inheritance)},
			{Name: "imports observed rows through quarantined airlock", Passed: a.Airlock.Executed && a.Airlock.EmpiricalImport && a.Airlock.AcceptedRows == RequiredRows && a.Airlock.QuarkMassRowsImported == 6 && a.Airlock.CKMRowsImported == 1 && a.Airlock.AllAcceptedQuarantined && a.Airlock.NoNativeRegistryWrite, Detail: FormatAirlock(a.Airlock)},
			{Name: "rejects native promotion probes", Passed: a.Airlock.NativePromotionRejectedProbe && a.Airlock.NativeRegistryWriteRejected && a.Airlock.ObservedAsTheoremRejected, Detail: FormatAirlock(a.Airlock)},
			{Name: "blocks coordinate map from mass spectra alone", Passed: a.Coordinate.Executed && !a.Coordinate.DUDDefined && !a.Coordinate.AlignmentAchieved && !a.Coordinate.CommonScaleSchemeSatisfied && !a.Coordinate.TraceZeroSpectrumModelSupplied && !a.Coordinate.IKComparatorSupplied && !a.Coordinate.BranchTagsSupplied && !a.Coordinate.AlphaUDefined && !a.Coordinate.PhiUDefined && !a.Coordinate.AlphaDDefined && !a.Coordinate.PhiDDefined && a.Coordinate.ProjectiveRayDOF == 2 && a.Coordinate.SpectrumOnlyRank == 1, Detail: FormatCoordinate(a.Coordinate)},
			{Name: "preserves native theorem registry", Passed: a.Firewall.Executed && a.Firewall.ObservedRowsImported == RequiredRows && a.Firewall.AllObservedRowsQuarantined && !a.Firewall.EmpiricalDataInNativeRegistry && !a.Firewall.NativePredictionFromEmpirical && !a.Firewall.NativeLawFromEmpirical && !a.Firewall.ObservedDataUsedAsTheoremInput && !a.Firewall.QuarkMassNativePrediction && !a.Firewall.CKMNativePrediction && !a.Firewall.CKMMatrixConstructed && !a.Firewall.CKMEntryComputed && !a.Firewall.DUDPromotedNative && !a.Firewall.AlignmentPromotedNative && a.Firewall.SectorCoefficientsStillSealed && a.Firewall.NativeFlavorDimAfter == NativeFlavorDim && a.Firewall.KXYCoeffDimAfter == KXYCoeffDim, Detail: FormatFirewall(a.Firewall)},
			{Name: "next gate requires common-scale comparator design", Passed: a.Next.Gate == 467, Detail: FormatNext(a.Next)},
		}
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: checks, Notes: []string{StatusObservedRowsImported, StatusFailedCommonScaleSchemeRequired, StatusFailedMassSpectraDoNotDefineRay, StatusFailedMissingIKComparator, StatusFailedMissingBranchTags, StatusFailedDUDUndefined, StatusFailedAlignmentNotComputable, StatusFirewallPreserved, a.Truth}}
	}}
}
