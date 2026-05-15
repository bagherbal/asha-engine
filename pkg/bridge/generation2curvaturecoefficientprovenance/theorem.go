package generation2curvaturecoefficientprovenance

import "github.com/bagherbal/asha-engine/pkg/theorem"

func Generation2CurvatureCoefficientProvenanceHeatKernelTraceConventionAuditTheorem() theorem.Theorem {
	const id = AuditID
	const name = "Generation 2 curvature coefficient provenance and heat-kernel trace convention audit"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "construct Gate510 curvature coefficient audit", Passed: false, Detail: err.Error()}}}
		}
		checks := []theorem.Check{
			{Name: "inherits Gate509 gravity socket without normalization", Passed: a.Inheritance.Executed && a.Inheritance.Gate509GravitySocketInherited && a.Inheritance.Gate509NormalizationBlocked && a.Inheritance.Gate509NoEmpiricalDataImported && a.Inheritance.ProductTripleValid && a.Inheritance.HeatKernelConventionDeclared && a.Inheritance.Gate377RawA2ChannelPresent && a.Inheritance.Gate377SkeletonChannelPresent && !a.Inheritance.Gate377AllCoefficientsClosed && !a.Inheritance.Gate377HardTOEClosure, Detail: FormatInheritance(a.Inheritance)},
			{Name: "audits D2 curvature endomorphism", Passed: a.Endomorphism.Executed && a.Endomorphism.CurvatureEndomorphismAudited && nearly(a.Endomorphism.CombinedA2RPart, -1.0/12.0, 1e-15) && nearly(a.Endomorphism.CombinedA2RPartMagnitude, 1.0/12.0, 1e-15) && !a.Endomorphism.FiniteDiracPartAddsCurvature && !a.Endomorphism.SignConventionClosed && !a.Endomorphism.PhysicalMetricDynamicsDerived, Detail: FormatEndomorphism(a.Endomorphism)},
			{Name: "computes native finite trace weight", Passed: a.A2.Executed && nearly(a.A2.FiniteHilbertTraceDimension, 96, 1e-15) && nearly(a.A2.A2WeightMagnitudeBefore4Pi, 8, 1e-15) && a.A2.RawDensityCoefficientPerF2Lambda2 > 0 && a.A2.Gate377RawCoefficientMatched && a.A2.DimensionlessTraceWeightNative && !a.A2.IncludesCutoffMoment && !a.A2.PhysicalCoefficientNative, Detail: FormatA2(a.A2)},
			{Name: "keeps trace convention unpromoted", Passed: a.Convention.Executed && a.Convention.RawConventionDeclared && a.Convention.SkeletonConventionDeclared && a.Convention.RawSkeletonNumericallyDifferent && !a.Convention.UniqueTraceConventionSelected && !a.Convention.CanPromoteEitherToNewtonNative, Detail: FormatConvention(a.Convention)},
			{Name: "isolates f2 Lambda squared obligation", Passed: a.Cutoff.Executed && a.Cutoff.RequiresF2LambdaSquaredProduct && a.Cutoff.F2LambdaProductNativeOnlyAsSymbol && !a.Cutoff.F2MomentSeparatedFromLambda && !a.Cutoff.CutoffLambdaSelected && !a.Cutoff.NewtonConstantDerived && !a.Cutoff.PlanckScaleImported && !a.Cutoff.CosmologicalConstantDerived && a.Cutoff.CosmologicalF4Excluded && a.Cutoff.GravityNormalizationBridgeOnly, Detail: FormatCutoff(a.Cutoff)},
			{Name: "preserves gravity and scale firewall", Passed: a.Firewall.Executed && !a.Firewall.NewtonConstantImported && !a.Firewall.NewtonConstantDerived && !a.Firewall.PlanckMassImported && !a.Firewall.CutoffLambdaSelected && !a.Firewall.F2MomentSeparatedFromLambda && !a.Firewall.EinsteinHilbertNormalizationClosed && !a.Firewall.CosmologicalConstantImported && !a.Firewall.CosmologicalConstantDerived && !a.Firewall.ElectroweakScaleImported && !a.Firewall.FlavorDataImported && !a.Firewall.NativeGravityNormalizationWritten, Detail: FormatFirewall(a.Firewall)},
		}
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Checks: checks, Notes: append(statuses(), a.Truth)}
	}}
}
