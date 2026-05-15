package generation2observedelectroweakfileadapter

import "github.com/bagherbal/asha-engine/pkg/theorem"

func Generation2ObservedElectroweakComparatorFileAdapterFirewallTheorem() theorem.Theorem {
	const id = AuditID
	const name = "Generation 2 observed electroweak comparator file adapter firewall"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "construct Gate507 electroweak file adapter", Passed: false, Detail: err.Error()}}}
		}
		checks := []theorem.Check{
			{Name: "inherit Gate506 preflight and redirect", Passed: a.Inheritance.Executed && a.Inheritance.Gate506PreflightValidated && !a.Inheritance.Gate506NumericalAdapterExecuted && !a.Inheritance.Gate506ObservedNumbersImported && a.Inheritance.Gate506NativeRegistryWriteBlocked && a.Inheritance.Gate507RedirectDefined, Detail: FormatInheritance(a.Inheritance)},
			{Name: "load explicit bridge-only electroweak comparator file", Passed: a.Import.Executed && a.Import.Loaded && a.Import.Rows == 6 && a.Import.AcceptedRows == 6 && a.Import.RejectedRows == 0 && a.Import.InputRows == 3 && a.Import.ComparatorRows == 3 && a.Import.EmpiricalImport && a.Import.BridgeOnlyLedger && a.Import.SyntheticFixture && !a.Import.ObservedValuesLoaded && !a.Import.NativeRegistryWriteRequested && a.Import.MetadataComplete, Detail: FormatImport(a.Import)},
			{Name: "execute tree-level adapter from file inputs", Passed: a.Output.Executed && a.Output.Attempted && a.Output.Ready && a.Output.UsedTreeLevelFormula && nearly(a.Output.Sin2ThetaW, 16.0/25.0, 1e-12) && nearly(a.Output.MW, 3, 1e-12) && nearly(a.Output.MZ, 5, 1e-12) && a.Output.PhotonZeroPreserved && a.Output.RhoIdentityConfirmed, Detail: FormatOutput(a.Output)},
			{Name: "compute bridge-only comparator residuals", Passed: a.Residuals.Executed && a.Residuals.ComparatorRowsAvailable && a.Residuals.WeakAngleResidualComputed && a.Residuals.MWResidualComputed && a.Residuals.MZResidualComputed && a.Residuals.AllResidualsZero && a.Residuals.BridgeOnly && !a.Residuals.NativePrediction, Detail: FormatResiduals(a.Residuals)},
			{Name: "preserve native electroweak firewall", Passed: a.Firewall.Executed && !a.Firewall.ObservedValuesImported && a.Firewall.SyntheticFixtureOnly && !a.Firewall.FileRowsNative && !a.Firewall.AdapterOutputsNative && !a.Firewall.WeakAngleNativePrediction && !a.Firewall.WZMassNativePrediction && !a.Firewall.GaugeCouplingsNativePrediction && !a.Firewall.VEVNativePrediction && !a.Firewall.KappaNativePromotion && !a.Firewall.NativeRegistryWritten && !a.Firewall.PhysicalElectroweakPredictionMade, Detail: FormatFirewall(a.Firewall)},
			{Name: "Gate508 residual-geometry redirect is defined", Passed: a.Next.Gate == 508, Detail: a.Next.Title + ": " + a.Next.PrimaryTask},
		}
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Checks: checks, Notes: append(statuses(), a.Truth)}
	}}
}
