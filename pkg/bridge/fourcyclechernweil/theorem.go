package fourcyclechernweil

import "github.com/bagherbal/asha-engine/pkg/theorem"

func FiniteOrientedFourCycleChernWeilCarrierConstructionSearchTheorem() theorem.Theorem {
	const id = "BRIDGE-FINITE-ORIENTED-FOUR-CYCLE-CHERN-WEIL-CARRIER-SEARCH"
	const name = "finite oriented four-cycle / Chern-Weil carrier construction search"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "build finite four-cycle/Chern-Weil search", Passed: false, Detail: err.Error()}}}
		}
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{
			{Name: "Gate 180 bridge preflight is consumed without promotion", Passed: a.PreviousGate180.Preflight.PromotableHeatKernelAnchors == 0 && a.PreviousGate180.Preflight.PromotableThresholdAnchors == 0 && !a.PreviousGate180.ChernWeilTrace.OrientedFourCycleDerived && !a.PreviousGate180.ChernWeilTrace.ChernWeilFormDerived && a.PreviousGate180.Firewall.StrictNullityAfter == 3 && a.PreviousGate180.Firewall.ConditionalNullityAfter == 2, Detail: a.PreviousGate180.Firewall.Verdict},
			{Name: "current finite candidates are substantially audited", Passed: a.Search.CandidatesAudited >= 8 && a.Search.ExactFiniteCandidates >= 7 && a.Search.FourDimensionalCandidates >= 2, Detail: FormatSearch(a.Search)},
			{Name: "no candidate provides a boundaryless oriented fundamental four-cycle", Passed: a.Search.BoundarylessCycleCandidates == 0 && a.Search.CanonicalFundamentalClasses == 0 && a.Search.IntegrationFunctionals == 0 && !a.Search.FiniteFourCycleDerived, Detail: FormatCandidates(a.Candidates)},
			{Name: "Chern-Weil carrier requirements remain blocking", Passed: allRequirementsBlocking(a.Requirements) && a.Search.GaugeBundleMaps == 0 && a.Search.CurvaturePairings == 0 && a.Search.TraceNormalizations == 0 && a.Search.IntegerChargeMaps == 0 && a.Search.CompleteChernWeilCarriers == 0, Detail: FormatRequirements(a.Requirements)},
			{Name: "grade-four and 4D predata do not equal orientability", Passed: a.Homology.GradeFourDataAvailable && a.Homology.DimensionFourVectorSpacesExist && a.Homology.BoundaryOperatorDerived && !a.Homology.NonzeroClosedFourCycleDerived && !a.Homology.CanonicalRepresentativeDerived && !a.Homology.HochschildCycleRealizesGamma, Detail: FormatHomology(a.Homology)},
			{Name: "topological seal and gauge trace are not promoted to instanton normalization", Passed: a.ChernWeil.GaugeAlgebraClosed && a.ChernWeil.RepresentationTraceRatioClosed && a.ChernWeil.TopologicalSealAvailable && !a.ChernWeil.PrincipalBundleDerived && !a.ChernWeil.ConnectionOnFourCarrierDerived && !a.ChernWeil.CurvatureTwoFormDerived && !a.ChernWeil.TracePairingDerived && !a.ChernWeil.IntegerInstantonNumberDerived && !a.ChernWeil.ContinuumNormalizationPromoted, Detail: FormatChernWeil(a.ChernWeil)},
			{Name: "nullity and physics firewall remain unchanged", Passed: !a.Firewall.UsesObservedInputForDerivation && !a.Firewall.FourCycleDerived && !a.Firewall.ChernWeilCarrierDerived && !a.Firewall.InstantonTraceBridgeDerived && !a.Firewall.AbsoluteCouplingPromoted && !a.Firewall.HeatKernelMatchingDerived && !a.Firewall.ThresholdCorrectedBetaDerived && !a.Firewall.PhysicalConstantsDerived && a.Firewall.StrictNullityBefore == 3 && a.Firewall.StrictNullityAfter == 3 && a.Firewall.ConditionalNullityBefore == 2 && a.Firewall.ConditionalNullityAfter == 2, Detail: FormatFirewall(a.Firewall)},
		}, Notes: []string{
			a.TruthStatement,
			"This is not a no-go against future Chern-Weil geometry. It proves that no existing finite carrier already supplies the oriented four-cycle/integration/bundle/curvature chain.",
			"The next target should be a local field/bundle map or a new finite carrier with orientability built in, not another spectral scalar comparison.",
		}}
	}}
}

func allRequirementsBlocking(xs []RequirementAudit) bool {
	if len(xs) == 0 {
		return false
	}
	for _, x := range xs {
		if x.DerivedByAnyCandidate || !x.Blocking {
			return false
		}
	}
	return true
}
