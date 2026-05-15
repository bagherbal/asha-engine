package generation2syntheticapsindexboundaryledger

import "github.com/bagherbal/asha-engine/pkg/theorem"

func Generation2SyntheticAPSIndexBoundaryLedgerDryRunTheorem() theorem.Theorem {
	const id = AuditID
	const name = "Generation 2 synthetic APS index boundary ledger dry-run"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "construct Gate518 synthetic APS ledger", Passed: false, Detail: err.Error()}}}
		}
		checks := []theorem.Check{
			{Name: "inherits Gate517 index/eta airlock", Passed: a.Inheritance.Executed && a.Inheritance.Gate517Inherited && a.Inheritance.Gate517LocalIndexDensitySocket && a.Inheritance.Gate517APSSocket && a.Inheritance.Gate517BoundaryEtaAirlock && a.Inheritance.Gate517AnomalyInflowSocket && a.Inheritance.Gate517GlobalIndexBlocked && a.Inheritance.Gate517EtaBlocked && a.Inheritance.Gate517BoundarySpectrumBlocked && !a.Inheritance.Gate517ObservedBoundaryImported, Detail: FormatInheritance(a.Inheritance)},
			{Name: "computes synthetic APS and closed-manifold rows bridge-only", Passed: a.Ledger.Executed && a.Ledger.BridgeOnly && a.Ledger.SyntheticOnly && !a.Ledger.UsesObservedTopology && !a.Ledger.UsesBoundarySpectrum && nearly(a.Ledger.BoundaryCorrection, 2, eps) && nearly(a.Ledger.APSIndex, 9, eps) && nearly(a.Ledger.ClosedManifoldIndex, 11, eps) && nearly(a.Ledger.APSResidual, 0, eps) && nearly(a.Ledger.ClosedResidual, 0, eps) && a.Ledger.APSIndexIntegerLike && a.Ledger.ClosedIndexIntegerLike, Detail: FormatLedger(a.Ledger)},
			{Name: "enforces fail-closed topology/boundary metadata policy", Passed: a.Policy.Executed && a.Policy.RequiresBridgeOnlyTag && a.Policy.RequiresSyntheticOrExternalTag && a.Policy.RequiresSourceMetadata && a.Policy.RequiresTopologyMetadata && a.Policy.RequiresBoundaryMetadata && a.Policy.RejectsNativePromotion && a.Policy.RejectsObservedByDefault && a.Policy.RejectsMissingEtaKernelRows && a.Policy.RejectsMissingBoundaryCondition && a.Policy.RejectsMissingUncertaintyMetadata && !a.Policy.NativeIndexPredictionMade && !a.Policy.NativeEtaPredictionMade && !a.Policy.BoundaryConditionSelected && !a.Policy.BoundarySpectrumDerived && !a.Policy.ClosedManifoldNativelySelected, Detail: FormatPolicy(a.Policy)},
			{Name: "blocks synthetic APS native write and physical normalizations", Passed: a.Firewall.Executed && !a.Firewall.UsesLambdaCutoff && !a.Firewall.UsesF2Moment && !a.Firewall.UsesF4Moment && !a.Firewall.UsesNewtonConstant && !a.Firewall.UsesCosmologicalConstant && !a.Firewall.UsesPlanckScale && !a.Firewall.UsesHiggsOrElectroweakScale && !a.Firewall.UsesFlavorYukawaData && !a.Firewall.ObservedTopologyImported && !a.Firewall.ObservedBoundarySpectrumImported && !a.Firewall.SyntheticOutputNativeWrite && !a.Firewall.GlobalIndexNativePrediction && !a.Firewall.BoundaryEtaNativePrediction && !a.Firewall.PhysicalGravitationalThetaWritten && !a.Firewall.GravityCosmologyNormalizationWrite, Detail: FormatFirewall(a.Firewall)},
		}
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Checks: checks, Notes: append(statuses(), a.Truth)}
	}}
}
