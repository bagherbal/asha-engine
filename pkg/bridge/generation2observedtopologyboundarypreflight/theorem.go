package generation2observedtopologyboundarypreflight

import "github.com/bagherbal/asha-engine/pkg/theorem"

func Generation2ObservedTopologyBoundaryComparatorPreflightTheorem() theorem.Theorem {
	const id = AuditID
	const name = "Generation 2 observed topology and boundary comparator preflight"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "construct Gate519 topology/boundary preflight", Passed: false, Detail: err.Error()}}}
		}
		checks := []theorem.Check{
			{Name: "inherits Gate518 APS/index firewall", Passed: a.Inheritance.Executed && a.Inheritance.Gate518Inherited && a.Inheritance.Gate518SyntheticAPSDryRun && a.Inheritance.Gate518BridgeOnly && a.Inheritance.Gate518GlobalTopologyBlocked && a.Inheritance.Gate518BoundaryEtaBlocked && a.Inheritance.Gate518NativeWriteBlocked && !a.Inheritance.ObservedDataImported, Detail: FormatInheritance(a.Inheritance)},
			{Name: "defines fail-closed topology schema", Passed: a.Topology.Executed && a.Topology.RequiredRows == 7 && a.Topology.RequiresEulerCharacteristic && a.Topology.RequiresPontryaginClasses && a.Topology.RequiresSignature && a.Topology.RequiresGlobalAPSIndex && a.Topology.RequiresManifoldDimension && a.Topology.RequiresOrientationAndClosedness && a.Topology.RequiresModelID && a.Topology.RejectsNativePromotion && a.Topology.RedactedSchemaAccepted && !a.Topology.ObservedNumbersImported && a.Topology.ComparatorTargetOnly, Detail: FormatTopology(a.Topology)},
			{Name: "defines fail-closed boundary schema", Passed: a.Boundary.Executed && a.Boundary.RequiredRows == 7 && a.Boundary.RequiresBoundaryConditionType && a.Boundary.RequiresEtaInvariantValue && a.Boundary.RequiresKernelDimensionH && a.Boundary.RequiresBoundarySpectrumMetadata && a.Boundary.RequiresBoundaryOrientation && a.Boundary.RequiresBoundaryComponentCount && a.Boundary.RequiresModelID && a.Boundary.RejectsNativePromotion && a.Boundary.RedactedSchemaAccepted && !a.Boundary.ObservedNumbersImported && a.Boundary.ComparatorTargetOnly, Detail: FormatBoundary(a.Boundary)},
			{Name: "enforces mandatory provenance metadata", Passed: a.Policy.Executed && a.Policy.RequiresSource && a.Policy.RequiresSourceVersion && a.Policy.RequiresUncertainty && a.Policy.RequiresScheme && a.Policy.RequiresScaleOrTopologyContext && a.Policy.RequiresBridgeOnlyTrue && a.Policy.RequiresNativePromotionFalse && a.Policy.RequiresComparatorOnlyPurpose && a.Policy.RequiresNoTheoremInputFlag && a.Policy.RejectsMissingSource && a.Policy.RejectsMissingUncertainty && a.Policy.RejectsBridgeOnlyFalse && a.Policy.RejectsNativePromotionTrue && a.Policy.AcceptedRedactedSchemaCases == 1 && a.Policy.RejectedFailClosedCases >= 10, Detail: FormatPolicy(a.Policy)},
			{Name: "blocks native topology/boundary writes and comparator execution", Passed: a.Rejection.Executed && a.Rejection.TopologyNativePredictionBlocked && a.Rejection.BoundaryEtaNativePredictionBlock && a.Rejection.GlobalAPSIndexNativeWriteBlocked && a.Rejection.EulerCharacteristicNativeBlocked && a.Rejection.PontryaginNumberNativeBlocked && a.Rejection.SignatureNativeBlocked && a.Rejection.BoundarySpectrumNativeBlocked && a.Rejection.ClosedManifoldConditionBlocked && a.Rejection.ComparatorExecutionBlockedNow && a.Rejection.ResidualComputationBlockedNow, Detail: FormatRejection(a.Rejection)},
			{Name: "preserves topology/boundary firewall", Passed: a.Firewall.Executed && !a.Firewall.ObservedTopologyImported && !a.Firewall.ObservedBoundaryDataImported && !a.Firewall.ObservedBoundarySpectrumImported && !a.Firewall.UsesNewtonConstant && !a.Firewall.UsesPlanckScale && !a.Firewall.UsesLambdaCutoff && !a.Firewall.UsesCosmologicalConstant && !a.Firewall.UsesElectroweakScale && !a.Firewall.UsesFlavorYukawaData && !a.Firewall.NativeTopologyWrite && !a.Firewall.NativeBoundaryWrite && !a.Firewall.NativeGlobalIndexWrite && !a.Firewall.ComparatorExecuted, Detail: FormatFirewall(a.Firewall)},
		}
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Checks: checks, Notes: append(statuses(), a.Truth)}
	}}
}
