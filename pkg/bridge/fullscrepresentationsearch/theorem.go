package fullscrepresentationsearch

import "github.com/bagherbal/asha-engine/pkg/theorem"

func FullSCFiniteAlgebraRepresentationSearchOppositeActionConstructionAuditTheorem() theorem.Theorem {
	const id = "BRIDGE-FULL-SC-FINITE-ALGEBRA-REPRESENTATION-OPPOSITE-ACTION-CONSTRUCTION-AUDIT"
	const name = "Full S_C Finite Algebra Representation Search / Opposite-Action Construction Audit"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "build Gate 271 full-S_C representation search", Passed: false, Detail: err.Error()}}}
		}
		checks := []theorem.Check{
			{Name: "Gate 270 non-vacuous target is inherited without promotion", Passed: a.Inheritance.CandidateOneFormsExposed && !a.Inheritance.CandidateOrderOnePasses && !a.Inheritance.FullSCRepresentation && !a.Inheritance.PhysicalOppositeAction && !a.Inheritance.HiggsRatioDerived && a.Inheritance.FirewallPreserved, Detail: FormatInheritance(a.Inheritance)},
			{Name: "full S_C Fock carrier and CAR operators are available", Passed: a.Carrier.ModeCount == 4 && a.Carrier.BaseComplexDimension == 16 && a.Carrier.DoubledComplexDimension == 32 && a.Carrier.BasisMasksEnumerated == 16 && a.Carrier.CARPassed && a.Carrier.ParityEvenStates == 8 && a.Carrier.ParityOddStates == 8, Detail: FormatCarrier(a.Carrier)},
			{Name: "native full-carrier lifts are audited but no associative full-S_C algebra representation is derived", Passed: len(a.Representation.Candidates) == 3 && !a.Representation.ValidFullAssociativeRepFound && a.Representation.FullSCPromotionBlocked, Detail: FormatRepresentation(a.Representation)},
			{Name: "Γ lift fails linear additivity and dΓ lift fails unital associativity", Passed: liftFailureProfileOK(a.Representation.Candidates), Detail: FormatRepresentation(a.Representation)},
			{Name: "physical opposite action remains blocked", Passed: a.Opposite.RequiresValidLeftRepresentation && a.Opposite.CandidateJAntiLinear && !a.Opposite.CandidateJPhysicalSemantics && !a.Opposite.OppositeActionConstructed && !a.Opposite.OrderOneCanBeEvaluatedPhysically, Detail: FormatOpposite(a.Opposite)},
			{Name: "full order-one condition is not promoted without representation and J", Passed: !a.OrderOne.FullSCLeftRepAvailable && !a.OrderOne.PhysicalOppositeRepAvailable && !a.OrderOne.NonVacuousOneFormsDerived && !a.OrderOne.OrderOneSatisfied && !a.OrderOne.ReevaluatedAsSpectralTriple && a.OrderOne.Gate270ToyResidualInherited > 0, Detail: FormatOrderOne(a.OrderOne)},
			{Name: "Higgs ratio and x:y selector remain blocked", Passed: !a.Ratio.XYRatioSelected && !a.Ratio.TraceRatioStable && !a.Ratio.GaugeProjectionDerived && !a.Ratio.ScalarFluctuationMapDerived && !a.Ratio.HeatKernelNormalizationDerived && !a.Ratio.HiggsRatioDerived, Detail: FormatRatio(a.Ratio)},
			{Name: "firewalls preserve seals and prevent candidate promotion", Passed: a.Firewall.EmpiricalYukawaSealPreserved && a.Firewall.SpontaneousCarrierSealPreserved && a.Firewall.NoObservedMassInserted && a.Firewall.NoVEVInserted && a.Firewall.NoCutoffScaleInserted && a.Firewall.NoConnesModelImported && a.Firewall.NoCandidatePromoted && a.Firewall.NoHiggsPredictionClaim && !a.Firewall.FiniteCorePolluted, Detail: FormatFirewall(a.Firewall)},
			{Name: "future map names the representation-classification obligations", Passed: len(a.Future.Obligations) >= 6 && a.Future.NeedAssociativeFullSCRep && a.Future.NeedPhysicalJ && a.Future.NeedOrderOnePassingNonVacuous && a.Future.NeedCanonicalXYSelector && a.Future.NeedSpectralActionProjection, Detail: FormatFuture(a.Future)},
			{Name: "summary records full-carrier audit and failed spectral-triple route", Passed: a.Summary.Gate270Inherited && a.Summary.FullCarrierEnumerated && a.Summary.CARPreflightPassed && a.Summary.NativeOperatorLiftsAudited && !a.Summary.ValidFullSCRepDerived && !a.Summary.PhysicalJDerived && !a.Summary.FullOrderOneProved && !a.Summary.NonVacuousOneFormsProved && !a.Summary.HiggsRatioDerived && a.Summary.FirewallPreserved, Detail: FormatSummary(a.Summary)},
		}
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Checks: checks, Notes: []string{
			a.TruthStatement,
			"Gate 271 does not import the standard Connes representation. It audits the native full-Fock candidates and records why Γ, dΓ, and the Λ¹ seed action are insufficient as the required associative representation on S_C.",
			"The next lawful step is representation-classification/Morita-bimodule search, not another hand-built chiral mismatch.",
		}}
	}}
}

func liftFailureProfileOK(cs []LiftCandidate) bool {
	if len(cs) != 3 {
		return false
	}
	var gamma, dgamma, one bool
	for _, c := range cs {
		switch c.Name {
		case "Γ exterior functor lift":
			gamma = c.ActsOnFullSC && c.Multiplicative && c.Unital && !c.LinearAdditive && !c.AssociativeAlgebraRep && c.DiagnosticDefect > 0
		case "dΓ creation-annihilation bilinear lift":
			dgamma = c.ActsOnFullSC && c.UsesCreationAnnihilation && c.LinearAdditive && !c.Multiplicative && !c.Unital && !c.AssociativeAlgebraRep && c.DiagnosticDefect > 0
		case "one-particle sector inclusion":
			one = !c.ActsOnFullSC && c.FaithfulOnOneParticle && c.LinearAdditive && c.Multiplicative && c.Unital && !c.AssociativeAlgebraRep
		}
	}
	return gamma && dgamma && one
}
