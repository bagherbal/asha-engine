package generation2spectralcutoffrenormalizationairlock

import "github.com/bagherbal/asha-engine/pkg/theorem"

func Generation2SpectralCutoffRenormalizationAirlockComparatorTheorem() theorem.Theorem {
	const id = AuditID
	const name = "Generation 2 spectral cutoff and renormalization airlock comparator"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "construct Gate514 cutoff-renormalization airlock", Passed: false, Detail: err.Error()}}}
		}
		checks := []theorem.Check{
			{Name: "inherits Gate513 stripped hierarchy and firewall", Passed: a.Inheritance.Executed && a.Inheritance.Gate513Inherited && a.Inheritance.StrippedHierarchyNative && nearly(a.Inheritance.A2OverA0Ratio, 1.0/12.0, 1e-12) && nearly(a.Inheritance.A4OverA0Ratio, 1.0/360.0, 1e-12) && nearly(a.Inheritance.A4OverA2Ratio, 1.0/30.0, 1e-12) && !a.Inheritance.Gate513F2Selected && a.Inheritance.Gate513NativeNormalizationBlocked, Detail: FormatInheritance(a.Inheritance)},
			{Name: "defines redacted bridge comparator schema", Passed: a.Schema.Executed && a.Schema.RequiredRowCount == 10 && a.Schema.AcceptedRedactedRows == 10 && a.Schema.NumericalRows == 0 && a.Schema.EmpiricalRows == 0 && a.Schema.RowsBridgeOnly && a.Schema.RowsRejectNativePromotion && a.Schema.AllMetadataComplete, Detail: FormatSchema(a.Schema)},
			{Name: "rejects invalid cutoff and renormalization cases", Passed: a.Preflight.Executed && a.Preflight.AcceptedCases == 1 && a.Preflight.RejectedCases == 8 && a.Preflight.RejectedNativePromotionCases >= 2 && a.Preflight.RejectedNumericalCases >= 2 && a.Preflight.RejectedMissingMetadataCases == 1 && a.Preflight.RejectedExecutionCases == 1, Detail: FormatPreflight(a.Preflight)},
			{Name: "quarantines cutoff, moments, Planck matching, and vacuum subtraction", Passed: a.Airlock.Executed && !a.Airlock.LambdaCutoffSelected && !a.Airlock.F2MomentSelected && !a.Airlock.F4MomentSelected && !a.Airlock.F2LambdaProductSeparated && !a.Airlock.F4LambdaProductSeparated && !a.Airlock.PlanckMatchingNative && !a.Airlock.NewtonConstantDerived && !a.Airlock.CosmologicalConstantDerived && !a.Airlock.VacuumSubtractionSelectedNative && !a.Airlock.RenormalizationSchemeNative && !a.Airlock.NumericalAdapterExecuted && !a.Airlock.NativeNormalizationWrite, Detail: FormatAirlock(a.Airlock)},
			{Name: "preserves gravity/cosmology firewall", Passed: a.Firewall.Executed && !a.Firewall.NewtonConstantImported && !a.Firewall.PlanckScaleImported && !a.Firewall.CutoffLambdaImported && !a.Firewall.F2MomentImported && !a.Firewall.F4MomentImported && !a.Firewall.F2LambdaProductImported && !a.Firewall.F4LambdaProductImported && !a.Firewall.CosmologicalConstantImported && !a.Firewall.DarkEnergyImported && !a.Firewall.VacuumSubtractionImported && !a.Firewall.ObservedComparatorImported && !a.Firewall.NativeCutoffRenormalizationWrite, Detail: FormatFirewall(a.Firewall)},
		}
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Checks: checks, Notes: append(statuses(), a.Truth)}
	}}
}
