package generation2spectralmomenthierarchyairlock

import "github.com/bagherbal/asha-engine/pkg/theorem"

func Generation2SpectralMomentHierarchyCutoffSeparationAirlockAuditTheorem() theorem.Theorem {
	const id = AuditID
	const name = "Generation 2 spectral moment hierarchy and cutoff-separation airlock audit"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "construct Gate513 spectral moment hierarchy audit", Passed: false, Detail: err.Error()}}}
		}
		checks := []theorem.Check{
			{Name: "inherits a0/a2/a4 spectral boundaries", Passed: a.Inheritance.Executed && a.Inheritance.Gate512Inherited && a.Inheritance.Gate512A0PrefactorNative && a.Inheritance.Gate512CosmologicalBlocked && a.Inheritance.Gate510A2Inherited && a.Inheritance.Gate510NewtonBlocked && a.Inheritance.Gate511A4Inherited && a.Inheritance.Gate511PhysicalDynamicsBlocked && a.Inheritance.ProductTripleValid && a.Inheritance.ProductMomentLedgerAvailable && !a.Inheritance.ProductAllCoefficientsClosed, Detail: FormatInheritance(a.Inheritance)},
			{Name: "constructs three-channel heat-kernel ledger", Passed: a.Ledger.Executed && nearly(a.Ledger.FiniteTrace, 96, 1e-12) && a.Ledger.AllMatched && !a.Ledger.A0.Physical && !a.Ledger.A2.Physical && !a.Ledger.A4.Physical, Detail: FormatLedger(a.Ledger)},
			{Name: "computes stripped relative prefactor hierarchy", Passed: a.Hierarchy.Executed && a.Hierarchy.DimensionlessCombinatoric && nearly(a.Hierarchy.A2OverA0AfterFactoring, 1.0/12.0, 1e-12) && nearly(a.Hierarchy.A4OverA0AfterFactoring, 1.0/360.0, 1e-12) && nearly(a.Hierarchy.A4OverA2AfterFactoring, 1.0/30.0, 1e-12) && !a.Hierarchy.SelectsF2Moment && !a.Hierarchy.SelectsF4Moment && !a.Hierarchy.SelectsCutoffLambda && !a.Hierarchy.PhysicalNormalization, Detail: FormatHierarchy(a.Hierarchy)},
			{Name: "quarantines spectral moments and cutoff separation", Passed: a.Airlock.Executed && a.Airlock.F0MomentDimensionlessAvailable && !a.Airlock.F2MomentSelected && !a.Airlock.F4MomentSelected && !a.Airlock.F2LambdaProductSeparated && !a.Airlock.F4LambdaProductSeparated && !a.Airlock.CutoffLambdaSelected && !a.Airlock.PlanckCutoffRelationNative && !a.Airlock.NewtonConstantDerived && !a.Airlock.CosmologicalConstantDerived && !a.Airlock.VacuumSubtractionSelected && !a.Airlock.NativeNormalizationWrite, Detail: FormatAirlock(a.Airlock)},
			{Name: "preserves gravity/cosmology/electroweak/flavor firewall", Passed: a.Firewall.Executed && !a.Firewall.NewtonConstantImported && !a.Firewall.PlanckScaleImported && !a.Firewall.CutoffLambdaImported && !a.Firewall.F2MomentImported && !a.Firewall.F4MomentImported && !a.Firewall.CosmologicalConstantImported && !a.Firewall.DarkEnergyImported && !a.Firewall.ElectroweakScaleImported && !a.Firewall.FlavorDataImported && !a.Firewall.NativeSpectralNormalizationWrite, Detail: FormatFirewall(a.Firewall)},
		}
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Checks: checks, Notes: append(statuses(), a.Truth)}
	}}
}
