package generation2observedtopologyboundaryfileadapter

import "github.com/bagherbal/asha-engine/pkg/theorem"

func Generation2ObservedTopologyBoundaryFileAdapterFirewallTheorem() theorem.Theorem {
	const id = AuditID
	const name = "Generation 2 observed topology and boundary file adapter firewall"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "construct Gate520 topology/boundary file adapter", Passed: false, Detail: err.Error()}}}
		}
		checks := []theorem.Check{
			{Name: "inherit Gate519 topology/boundary preflight", Passed: a.Inheritance.Executed && a.Inheritance.Gate519PreflightDefined && a.Inheritance.Gate519TopologyRows == 7 && a.Inheritance.Gate519BoundaryRows == 7 && a.Inheritance.Gate519RequiresBridgeOnly && a.Inheritance.Gate519RejectsNativePromotion && !a.Inheritance.Gate519ComparatorExecuted && !a.Inheritance.Gate519ObservedDataImported && a.Inheritance.Gate519NativeRegistryBlocked && a.Inheritance.Gate520FileAdapterRedirect, Detail: FormatInheritance(a.Inheritance)},
			{Name: "load explicit bridge-only topology/boundary file", Passed: a.Import.Executed && a.Import.Loaded && a.Import.Rows == 15 && a.Import.AcceptedRows == 15 && a.Import.RejectedRows == 0 && a.Import.TopologyRows == 7 && a.Import.BoundaryRows == 7 && a.Import.AdapterRows == 1 && a.Import.EmpiricalImport && a.Import.BridgeOnlyLedger && a.Import.SyntheticFixture && !a.Import.ObservedValuesLoaded && !a.Import.NativeRegistryWriteRequested && a.Import.MetadataComplete && a.Import.AllRowsBridgeOnly && a.Import.AllRowsComparatorOnly && a.Import.AllRowsNoTheoremInput, Detail: FormatImport(a.Import)},
			{Name: "compute APS and signature residuals bridge-only", Passed: a.Output.Executed && a.Output.Attempted && a.Output.Ready && a.Output.BridgeOnly && !a.Output.NativePrediction && a.Output.UsesAPSBoundaryCorrection && nearly(a.Output.BoundaryCorrection, 2, 1e-12) && nearly(a.Output.ComputedAPSIndex, 9, 1e-12) && nearly(a.Output.APSResidual, 0, 1e-12) && nearly(a.Output.ComputedSignatureFromP1, 1, 1e-12) && nearly(a.Output.SignatureResidual, 0, 1e-12) && a.Output.BoundaryMode && a.Output.AllResidualsZero, Detail: FormatOutput(a.Output)},
			{Name: "preserve topology/boundary firewall", Passed: a.Firewall.Executed && !a.Firewall.ObservedTopologyImported && !a.Firewall.ObservedBoundaryDataImported && !a.Firewall.ObservedBoundarySpectrumImported && a.Firewall.SyntheticFixtureOnly && !a.Firewall.FileRowsNative && !a.Firewall.AdapterOutputsNative && !a.Firewall.EulerNativePrediction && !a.Firewall.PontryaginNativePrediction && !a.Firewall.SignatureNativePrediction && !a.Firewall.GlobalAPSIndexNativePrediction && !a.Firewall.EtaNativePrediction && !a.Firewall.BoundarySpectrumNativePrediction && !a.Firewall.BoundaryConditionNativeSelected && !a.Firewall.NativeRegistryWritten && !a.Firewall.NewtonPlanckCosmologyImported, Detail: FormatFirewall(a.Firewall)},
			{Name: "Gate521 bordism redirect is defined", Passed: a.Next.Gate == 521, Detail: a.Next.Title + ": " + a.Next.PrimaryTask},
		}
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Checks: checks, Notes: append(statuses(), a.Truth)}
	}}
}
