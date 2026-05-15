package generation2nullconeikselector

import "github.com/bagherbal/asha-engine/pkg/theorem"

func Generation2AlgebraicNullConeBridgeIKSelectionTheorem() theorem.Theorem {
	const id = AuditID
	const name = "Generation 2 algebraic null-cone bridge I_K selection"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "construct Gate480 audit", Passed: false, Detail: err.Error()}}}
		}
		checks := []theorem.Check{
			{Name: "inherits post-Gate478 firewall", Passed: a.Inheritance.Executed && a.Inheritance.Gate444KGenForced && a.Inheritance.Gate445TriangleForced && a.Inheritance.Gate454IKFormulaAvailable && a.Inheritance.Gate473RawMassIKFailed && a.Inheritance.Gate474ElectroweakIKFailed && a.Inheritance.Gate478LeptonFirewallClean && a.Inheritance.NativeRegistryClean, Detail: FormatInheritance(a.Inheritance)},
			{Name: "maps K/X/Y to declared null-cone ledger", Passed: a.Map.Executed && a.Map.NativeNullConeExists && a.Map.KTimelikeAssigned && a.Map.XYSpacelikeAssigned && a.Map.AssignmentDeclaredForGate480 && a.Map.NullBoundaryDeclaredForGate480 && !a.Map.AssignmentPreviouslyForced && !a.Map.NullBoundaryPreviouslyForced, Detail: FormatMap(a.Map)},
			{Name: "derives alpha_vac=1 and I_K=1/2", Passed: a.Sieve.Executed && a.Sieve.NullForcesAEqualsR && a.Sieve.UniqueModuloScaleOrientation && a.Sieve.AcceptedNullCases == 2 && a.Sieve.TimelikeRejected && a.Sieve.SpacelikeRejected && a.Sieve.AlphaVac == 1 && a.Sieve.IKVac == 0.5, Detail: FormatSieve(a.Sieve)},
			{Name: "keeps sector coordinates unresolved", Passed: a.Gap.Executed && a.Gap.VacuumBaselineIKDefined && !a.Gap.PhysicalSectorIKDefined && !a.Gap.QuarkSectorCoordinatesSolved && !a.Gap.LeptonSectorCoordinatesSolved && !a.Gap.DUDComputed && !a.Gap.DENuComputed && !a.Gap.CKMPrediction && !a.Gap.PMNSPrediction && a.Gap.NeedsSectorRayPerturbationLedger, Detail: FormatGap(a.Gap)},
			{Name: "preserves flavor firewall", Passed: a.Firewall.Executed && !a.Firewall.ObservedMassImported && !a.Firewall.CKMImported && !a.Firewall.PMNSImported && a.Firewall.VacuumIKNativeBaseline && !a.Firewall.VacuumIKPhysicalSectorCoordinate && !a.Firewall.DUDNativePrediction && !a.Firewall.DENuNativePrediction && !a.Firewall.CKMMatrixConstructed && !a.Firewall.PMNSMatrixConstructed && !a.Firewall.NativeRegistryWritten && a.Firewall.KGenStillForced && a.Firewall.XTriangleStillForced && a.Firewall.YPhaseStillQuarantined && a.Firewall.SectorCoefficientsStillSealed && a.Firewall.NativeFlavorDimAfter == NativeFlavorDim && a.Firewall.KXYCoeffDimAfter == KXYCoeffDim, Detail: FormatFirewall(a.Firewall)},
		}
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Checks: checks, Notes: []string{StatusNullConeNative, StatusNullBoundaryApplied, StatusEquipartitionDerived, StatusIKDerived, StatusVacuumBaselineExported, StatusFailedNullBoundaryNotPriorForced, StatusFailedKXYSignatureAssumption, StatusFailedSectorCoordinatesUnsolved, StatusFailedCKMPMNSPredictionRejected, StatusFailedPhysicalIKPromotion, StatusFirewallPreserved, a.Truth}}
	}}
}
