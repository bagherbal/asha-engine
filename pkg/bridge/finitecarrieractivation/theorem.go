package finitecarrieractivation

import "github.com/bagherbal/asha-engine/pkg/theorem"

func FiniteCarrierActivationContactToRowSemanticsObstructionAuditTheorem() theorem.Theorem {
	const id = "BRIDGE-FINITE-CARRIER-ACTIVATION-CONTACT-TO-ROW-SEMANTICS-OBSTRUCTION-AUDIT"
	const name = "finite carrier activation / contact-to-row semantics obstruction audit"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "build finite carrier activation audit", Passed: false, Detail: err.Error()}}}
		}
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{
			{Name: "Gate 204 conditional row-lattice support is inherited without physical prediction", Passed: a.PreviousGate204.Gate204Inherited && a.PreviousGate204.Gate204ConditionalSupportPreserved && a.PreviousGate204.RepresentationLatticeConstructed && a.PreviousGate204.Gate201ShapesOnLattice && a.PreviousGate204.ContactMapFailed && a.PreviousGate204.UniversalBetaSourceStillExternal && a.PreviousGate204.UniversalFitAvoided && a.PreviousGate204.NoPhysicalPredictionClaim && !a.PreviousGate204.PhysicalUnificationClaimed, Detail: FormatGate204(a.PreviousGate204)},
			{Name: "seven contact partial-overlap modes are audited only as finite spectral anchors", Passed: len(a.ContactModes) == 7 && allFinitePositive(a.ContactModes) && noContactModePromoted(a.ContactModes), Detail: FormatContactModes(a.ContactModes, 7)},
			{Name: "gauge charge semantics are absent for contact-to-row promotion", Passed: a.GaugeCharge.ContactModesAudited == 7 && a.GaugeCharge.TargetShapesAudited == len(a.PreviousGate204.TargetShapes) && a.GaugeCharge.FiniteOverlapCarrierAvailable && !a.GaugeCharge.NativeSU3DynkinIndicesDerived && !a.GaugeCharge.NativeSU2DynkinIndicesDerived && !a.GaugeCharge.NativeHyperchargeDerived && !a.GaugeCharge.CanonicalGaugeRepInheritance && !a.GaugeCharge.CanFormDiracVectorlikeDoublet && !a.GaugeCharge.CanFormWeylSU2Adjoint && a.GaugeCharge.CandidateRowsAssigned == 0 && !a.GaugeCharge.GaugeChargeSemanticsComplete, Detail: FormatGaugeCharge(a.GaugeCharge)},
			{Name: "spin-statistics semantics are absent for contact beta coefficients", Passed: a.SpinStatistics.ContactModesAudited == 7 && !a.SpinStatistics.LocalContinuumFieldClassDerived && !a.SpinStatistics.LorentzKineticOperatorDerived && !a.SpinStatistics.WeylCoefficientDerived && !a.SpinStatistics.DiracCoefficientDerived && !a.SpinStatistics.ScalarCoefficientDerived && !a.SpinStatistics.SpinStatisticsAssigned && !a.SpinStatistics.StandardBetaCoefficientSelected, Detail: FormatSpinStatistics(a.SpinStatistics)},
			{Name: "mass activation and decoupling semantics are absent independently of the VEV seal", Passed: a.MassActivation.ContactModesAudited == 7 && a.MassActivation.DimensionlessSpectralValuesAvailable && !a.MassActivation.CanonicalPhysicalMassUnitDerived && !a.MassActivation.VEVIndependentActivationDerived && !a.MassActivation.DecouplingScaleDerived && !a.MassActivation.ActivationPredicateDerived && !a.MassActivation.MatchingSchemeDerived && !a.MassActivation.ThresholdCorrectedBetaRowsAllowed, Detail: FormatMassActivation(a.MassActivation)},
			{Name: "carrier activation is classified as a three-pillar obstruction", Passed: a.Classification.RequiredPillars == 3 && a.Classification.CompletePillars == 0 && len(a.Classification.MissingPillars) == 3 && !a.Classification.CarrierActivationDerived && !a.Classification.ContactModesCanBeHeavyRows && !a.Classification.ContactModesCanBeTargetShapes && a.Classification.Verdict == "FAILED_ROUTE", Detail: FormatClassification(a.Classification)},
			{Name: "firewalls remain sealed and Gate-201 shapes remain conditional support only", Passed: a.Firewall.Gate204Inherited && a.Firewall.Gate204ConditionalSupportPreserved && a.Firewall.RepresentationLatticeConstructed && a.Firewall.Gate201ShapesRemainConditional && !a.Firewall.ContactModesPromotedToBetaRows && !a.Firewall.ContactModesAssignedToGate201Shapes && !a.Firewall.ArbitraryChargeAssignmentInserted && !a.Firewall.ArbitrarySpinStatisticInserted && !a.Firewall.ArbitraryMassScaleInserted && !a.Firewall.PhenomenologicalVEVUsedForActivation && !a.Firewall.UniversalBetaFitAttempted && !a.Firewall.ContinuousScalesSolved && !a.Firewall.PhysicalUnificationClaimed && !a.Firewall.ThresholdCorrectedPhysicalFitClaimed && !a.Firewall.AbsoluteMassPredicted && !a.Firewall.FiniteMatchingCorrectionsDerived && a.Firewall.StrictNullityBefore == a.Firewall.StrictNullityAfter && a.Firewall.PhysicalPredictionNullityBefore == a.Firewall.PhysicalPredictionNullityAfter && a.Summary.FailedRouteLogged && a.Summary.NoPhysicalPredictionClaim, Detail: FormatFirewall(a.Firewall) + " :: " + FormatSummary(a.Summary)},
		}, Notes: []string{
			a.TruthStatement,
			"Gate 205 answer: the contact carrier cannot yet become a heavy beta-row source because charge, spin-statistics, and mass-activation semantics are all absent.",
			"The Gate-201 shapes remain legal representation-row shapes from Gate 204, but not finite-derived particles or activated thresholds.",
		}}
	}}
}

func noContactModePromoted(modes []ContactMode) bool {
	for _, m := range modes {
		if m.ChargeSemantics || m.SpinStatistics || m.MassActivation || m.DecouplingLaw || m.AssignedTargetShape != "" || m.BetaRowAllowed {
			return false
		}
	}
	return true
}
