package universaltracedeformation

import "github.com/bagherbal/asha-engine/pkg/theorem"

func UniversalTraceDeformationTopologicalBoundaryOffsetAuditTheorem() theorem.Theorem {
	const id = "BRIDGE-UNIVERSAL-TRACE-DEFORMATION-TOPOLOGICAL-BOUNDARY-OFFSET-AUDIT"
	const name = "universal trace deformation / topological boundary offset audit"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "build universal trace deformation audit", Passed: false, Detail: err.Error()}}}
		}
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{
			{Name: "Gate 201 universal-completion shapes are inherited as conditional phenomenology", Passed: a.Firewall.Gate201Inherited && a.Firewall.Gate201UniversalShapesConditionalOnly && !a.Firewall.Gate201PhysicalPredictionClaimed && !a.Firewall.ObservedInputsUsedForFiniteDerivation, Detail: inverseGateDetail(a)},
			{Name: "universal beta row is algebraically equivalent to a boundary-u offset", Passed: a.Equivalence.UniversalBetaShiftEquivalentToBoundaryOffset && a.Equivalence.RelativeRunningUnaffectedByUniversalRow && a.Equivalence.BoundaryOffsetActsOnlyAsCommonIntercept && a.Equivalence.SignConventionChecked && a.Equivalence.RequiresKnownLeverArm && !a.Equivalence.PhysicalPredictionClaim, Detail: FormatEquivalence(a.Equivalence)},
			{Name: "Gate-201 conditional shapes imply explicit required delta_u values", Passed: len(a.RequiredOffsets) == a.PreviousGate201.UniversalShapeCount() && len(a.RequiredOffsets) > 0 && requiredOffsetsPositive(a.RequiredOffsets), Detail: FormatRequiredOffsets(a.RequiredOffsets)},
			{Name: "B-sector gap and contact zeta traces are audited without promoting them to boundary offsets", Passed: a.FiniteTrace.BGapValue > 0 && a.FiniteTrace.ContactZetaValues == 5 && a.FiniteTrace.ContactActionCandidates == 10 && a.FiniteTrace.CandidatesAudited == len(a.TraceCandidates) && a.FiniteTrace.CanonicalBoundaryOffsetCandidates == 0 && !a.FiniteTrace.BGapCanonicalOffsetDerived && !a.FiniteTrace.ContactZetaCanonicalOffsetDerived, Detail: FormatFiniteTrace(a.FiniteTrace) + " :: " + FormatTraceCandidates(a.TraceCandidates, 5)},
			{Name: "defect-adjusted boundary substitution fails to absorb the universal row canonically", Passed: len(a.AbsorptionTests) > 0 && a.FiniteTrace.CanonicalPerfectAbsorptions == 0 && !a.FiniteTrace.UniversalVolumeDefectCanonicalized, Detail: FormatAbsorptionTests(a.AbsorptionTests, 8)},
			{Name: "mass, beta-row, matching, and physical-unification firewalls remain sealed", Passed: !a.Firewall.DefectAdjustedBoundaryDerived && !a.Firewall.BGapUsedAsPhysicalMass && !a.Firewall.BGapUsedAsBetaRow && !a.Firewall.ContactZetaUsedAsBetaRow && !a.Firewall.ArbitraryCoefficientInserted && !a.Firewall.PhysicalUnificationClaimed && !a.Firewall.ThresholdCorrectedPhysicalFitClaimed && !a.Firewall.AbsoluteMassPredicted && !a.Firewall.FiniteMatchingCorrectionsDerived && a.Firewall.StrictNullityBefore == a.Firewall.StrictNullityAfter && a.Firewall.PhysicalPredictionNullityBefore == a.Firewall.PhysicalPredictionNullityAfter, Detail: FormatFirewall(a.Firewall)},
			{Name: "summary records a strict failed-route obstruction rather than a forced fit", Passed: a.Summary.TestsAudited == 6 && a.Summary.EquivalenceEstablished && a.Summary.RequiredOffsetsComputed && a.Summary.FiniteTraceCandidatesAudited && a.Summary.NoCanonicalOffsetFound && a.Summary.UniversalCompletionNotDerived && a.Summary.FailedRouteLogged && a.Summary.NoPhysicalPredictionClaim, Detail: FormatSummary(a.Summary)},
		}, Notes: []string{
			a.TruthStatement,
			"Gate 202 answer: the universal beta completion from Gate 201 can be rewritten as a topological boundary-offset variable, but the finite B-gap/contact-zeta ledgers do not yet derive its value or coefficient.",
			"Therefore the correct theorem status is FAILED_ROUTE under current axioms. The obstruction is valuable: the next gate must classify the universal beta source rather than fit it to a trace scalar.",
		}}
	}}
}

func requiredOffsetsPositive(rs []RequiredBoundaryOffset) bool {
	for _, r := range rs {
		if !(r.FromGate201 && r.ConditionalOnly && !r.FiniteDerived && r.RequiredDeltaU > 0 && r.DefectAdjustedU > 1 && r.RequiredAlphaInvShift > 0) {
			return false
		}
	}
	return true
}

func inverseGateDetail(a Analysis) string {
	return "Gate201=" + a.PreviousGate201.TruthStatement + " :: " + FormatFirewall(a.Firewall)
}
