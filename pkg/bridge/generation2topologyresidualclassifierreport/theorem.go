package generation2topologyresidualclassifierreport

import "github.com/bagherbal/asha-engine/pkg/theorem"

func Generation2TopologyResidualClassifierReportNativeNonSelectionAuditTheorem() theorem.Theorem {
	const id = AuditID
	const name = "Generation 2 topology residual classifier report and native non-selection audit"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "construct Gate523 topology residual report", Passed: false, Detail: err.Error()}}}
		}
		checks := []theorem.Check{
			{Name: "inherit Gate520 and Gate522 topology adapters", Passed: a.Inheritance.Executed && a.Inheritance.Gate520FileAdapterDefined && a.Inheritance.Gate520FileLoaded && a.Inheritance.Gate520BridgeOnly && a.Inheritance.Gate520SyntheticOnly && a.Inheritance.Gate520APSResidualZero && a.Inheritance.Gate520SignatureResidualZero && a.Inheritance.Gate520BoundaryMode && a.Inheritance.Gate520NativeWriteBlocked && a.Inheritance.Gate522FileAdapterDefined && a.Inheritance.Gate522FileLoaded && a.Inheritance.Gate522BridgeOnly && a.Inheritance.Gate522SyntheticOnly && a.Inheritance.Gate522OrientedAdmissible && a.Inheritance.Gate522SpinAdmissible && a.Inheritance.Gate522SpinCAdmissible && a.Inheritance.Gate522CharacteristicResidualsZero && a.Inheritance.Gate522ClosedBoundary && a.Inheritance.Gate522NativeWriteBlocked, Detail: FormatInheritance(a.Inheritance)},
			{Name: "aggregate bridge-only residual classes", Passed: a.Report.Executed && a.Report.Rows == 4 && a.Report.ZeroResidualRows == 4 && a.Report.APSBoundaryRows == 2 && a.Report.ClosedBordismRows == 2 && a.Report.BridgeOnly && a.Report.SyntheticOnly && !a.Report.ObservedImported && !a.Report.NativePrediction && a.Report.ReportReady && a.Report.ClassifiesButDoesNotSelect, Detail: FormatReport(a.Report)},
			{Name: "enforce heterogeneous fixture identity guard", Passed: a.Guard.Executed && !a.Guard.CrossLedgerIdentityAsserted && !a.Guard.CrossLedgerIdentityAllowed && a.Guard.CrossLedgerMergeRejected && a.Guard.DifferentSyntheticContexts && a.Guard.BoundaryStatusCompatibleOnlyIfSeparated && a.Guard.Gate520BoundaryMode && a.Guard.Gate522ClosedBoundary && nearly(a.Guard.MergedSignatureResidual, 17, 1e-12) && nearly(a.Guard.BoundaryComponentResidualIfMerged, 1, 1e-12) && !a.Guard.NativeManifoldSelected, Detail: FormatGuard(a.Guard)},
			{Name: "preserve residual-report native-write firewall", Passed: a.Firewall.Executed && !a.Firewall.ObservedTopologyImported && !a.Firewall.ObservedBoundaryImported && !a.Firewall.ObservedBordismImported && !a.Firewall.ObservedTangentBundleImported && !a.Firewall.FileResidualsNative && !a.Firewall.ReportNative && !a.Firewall.ZeroResidualsNativeSelector && !a.Firewall.CrossLedgerMergeNative && !a.Firewall.BoundaryConditionNativeSelected && !a.Firewall.EtaNativeSelected && !a.Firewall.BordismClassNativeSelected && !a.Firewall.CharacteristicNumbersNativeSelected && !a.Firewall.ManifoldRepresentativeNative && !a.Firewall.NewtonPlanckCosmologyImported && !a.Firewall.NativeRegistryWritten, Detail: FormatFirewall(a.Firewall)},
			{Name: "Gate524 anomaly-inflow classifier redirect is defined", Passed: a.Next.Gate == 524, Detail: a.Next.Title + ": " + a.Next.PrimaryTask},
		}
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Checks: checks, Notes: append(statuses(), a.Truth)}
	}}
}
