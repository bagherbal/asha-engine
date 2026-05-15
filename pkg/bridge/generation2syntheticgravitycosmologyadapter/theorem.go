package generation2syntheticgravitycosmologyadapter

import "github.com/bagherbal/asha-engine/pkg/theorem"

func Generation2BridgeOnlyGravityCosmologyAdapterDryRunTheorem() theorem.Theorem {
	const id = AuditID
	const name = "Generation 2 bridge-only gravity/cosmology adapter dry-run"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "construct Gate515 synthetic gravity/cosmology adapter", Passed: false, Detail: err.Error()}}}
		}
		checks := []theorem.Check{
			{Name: "inherits Gate514 cutoff-renormalization airlock", Passed: a.Inheritance.Executed && a.Inheritance.Gate514Inherited && a.Inheritance.RedactedSchemaAccepted && a.Inheritance.RequiredRows == 10 && a.Inheritance.AcceptedCases == 1 && a.Inheritance.RejectedCases == 8 && a.Inheritance.Gate514NoAdapterExecuted && a.Inheritance.Gate514NativeWriteBlocked && !a.Inheritance.Gate514LambdaSelected && !a.Inheritance.Gate514NewtonDerived, Detail: FormatInheritance(a.Inheritance)},
			{Name: "uses explicit fake bridge inputs only", Passed: a.Inputs.Executed && a.Inputs.AllInputsSynthetic && a.Inputs.AllRowsBridgeOnly && a.Inputs.AllNativePromotionBlocked && !a.Inputs.ObservedDataImported && a.Inputs.LambdaCutoff == 2 && a.Inputs.F2Moment == 3 && a.Inputs.F4Moment == 5 && a.Inputs.F0Moment == 7 && a.Inputs.VacuumSubtraction == 11, Detail: FormatInputs(a.Inputs)},
			{Name: "computes synthetic a2/a0/a4 coefficients", Passed: a.Output.Executed && nearly(a.Output.F2LambdaSquared, 12, 1e-12) && nearly(a.Output.F4LambdaFourth, 80, 1e-12) && nearly(a.Output.EinsteinHilbertCoefficient, 6.0/(3.141592653589793*3.141592653589793), 1e-12) && !a.Output.NativeGravityPrediction && !a.Output.NativeCosmologyPrediction, Detail: FormatOutput(a.Output)},
			{Name: "computes synthetic residuals without observed comparator", Passed: a.Residuals.Executed && a.Residuals.ResidualsAreSynthetic && a.Residuals.ResidualsBridgeOnly && a.Residuals.ResidualsZeroByConstruction && !a.Residuals.ObservedComparatorUsed, Detail: FormatResiduals(a.Residuals)},
			{Name: "blocks native gravity/cosmology normalization write", Passed: a.Airlock.Executed && a.Airlock.NumericalAdapterExecuted && a.Airlock.SyntheticOnly && !a.Airlock.ObservedComparatorImported && !a.Airlock.LambdaNativeSelected && !a.Airlock.F2NativeSelected && !a.Airlock.F4NativeSelected && !a.Airlock.PlanckNewtonNative && !a.Airlock.CosmologicalConstantNative && !a.Airlock.VacuumSubtractionNative && !a.Airlock.NewtonConstantDerived && !a.Airlock.CosmologicalConstantDerived && !a.Airlock.NativeNormalizationWrite, Detail: FormatAirlock(a.Airlock)},
			{Name: "preserves no-observed-data firewall", Passed: a.Firewall.Executed && !a.Firewall.NewtonConstantImported && !a.Firewall.PlanckScaleImported && !a.Firewall.CutoffLambdaImported && !a.Firewall.F2MomentImported && !a.Firewall.F4MomentImported && !a.Firewall.CosmologicalConstantImported && !a.Firewall.DarkEnergyImported && !a.Firewall.VacuumSubtractionImported && !a.Firewall.ObservedComparatorImported && !a.Firewall.SyntheticOutputNativeWrite, Detail: FormatFirewall(a.Firewall)},
		}
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Checks: checks, Notes: append(statuses(), a.Truth)}
	}}
}
