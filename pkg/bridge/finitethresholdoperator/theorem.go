package finitethresholdoperator

import "github.com/bagherbal/asha-engine/pkg/theorem"

func FiniteThresholdOperatorDecouplingSpectrumSearchTheorem() theorem.Theorem {
	const id = "BRIDGE-FINITE-THRESHOLD-OPERATOR-DECOUPLING-SPECTRUM-SEARCH"
	const name = "finite threshold operator / decoupling spectrum search"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "build finite threshold-operator audit", Passed: false, Detail: err.Error()}}}
		}
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{
			{Name: "Gate 177 non-universal repair remains underived", Passed: a.PreviousGate177.Firewall.NonUniversalThresholdCanFitByConstruction && !a.PreviousGate177.Firewall.NonUniversalThresholdDerived && !a.PreviousGate177.Firewall.ThresholdCorrectionsDerived && !a.PreviousGate177.Firewall.HiddenObservedInputUsedForDerivation, Detail: FormatDeltaBWitness(a.DeltaBWitness)},
			{Name: "threshold operator requirements are explicitly audited", Passed: a.Requirements.CandidatesAudited == len(a.Candidates) && a.Requirements.CandidatesAudited >= 8 && len(a.Requirements.RequiredPieces) == 5 && a.Requirements.WithFiniteSpectrum >= 4 && a.Requirements.WithGaugeRepresentation >= 3, Detail: FormatRequirements(a.Requirements)},
			{Name: "no candidate has the full finite threshold chain", Passed: a.Requirements.NoCandidateHasAllPieces && a.Requirements.CompleteThresholdOperators == 0 && a.Requirements.FiniteDerivedThresholdOps == 0 && a.Requirements.WithPhysicalMassUnit == 0 && a.Requirements.WithActivationPredicate == 0 && a.Requirements.WithDecouplingRule == 0, Detail: FormatCandidates(a.Candidates)},
			{Name: "baseline beta rows are not heavy threshold corrections", Passed: a.Requirements.BaselineRowsAlreadyCounted == 1 && a.BetaMatching.ScalarSectorRowConstructed && a.BetaMatching.ScalarSectorMatchesBaseline && !a.BetaMatching.ScalarSectorIsThresholdCorrection && a.BetaMatching.BetaCorrectionRowsAllowed == 0, Detail: "scalar/contact aggregate is a baseline complex-doublet row, not a decoupled heavy threshold"},
			{Name: "finite spectra remain open anchors rather than decoupling spectra", Passed: a.Requirements.OpenFiniteSpectrumAnchors >= 4 && !a.BetaMatching.BGapRepresentationCompleted && !a.BetaMatching.ContactOverlapRepresentationCompleted && !a.BetaMatching.ActivationRuleDerived && !a.BetaMatching.DecouplingMatchingRuleDerived, Detail: FormatRequirements(a.Requirements)},
			{Name: "all attempted combinations are rejected as strict finite theorems", Passed: len(a.Combinations) >= 5 && noStrictCombination(a.Combinations) && hasObservedCombination(a.Combinations), Detail: FormatCombinations(a.Combinations)},
			{Name: "Gate 177 Δb witness is quarantined and cannot be promoted", Passed: a.DeltaBWitness.Gate177NonUniversalFitExists && !a.DeltaBWitness.Gate177FiniteThresholdDerived && a.DeltaBWitness.MinimumNormWitnessUsesExternalFit && !a.DeltaBWitness.CanBePromotedToFiniteOperator, Detail: FormatDeltaBWitness(a.DeltaBWitness)},
			{Name: "threshold/beta firewall remains closed", Passed: !a.Firewall.ThresholdOperatorDerived && !a.Firewall.NonUniversalDeltaBDerived && !a.Firewall.ThresholdCorrectedBetaDerived && !a.Firewall.Gate177RepairPromoted && !a.Firewall.UsesObservedInputForDerivation && a.Firewall.StrictNullityBefore == 3 && a.Firewall.StrictNullityAfter == 3 && a.Firewall.ConditionalNullityBefore == 2 && a.Firewall.ConditionalNullityAfter == 2 && !a.Firewall.PhysicalConstantsDerived, Detail: FormatFirewall(a.Firewall)},
		}, Notes: []string{
			a.TruthStatement,
			"Gate 178 is a clean no-go for currently derived threshold operators: exact finite spectra exist, but no activation/decoupling/representation/beta chain exists.",
			"The non-universal Δb branch remains mathematically capable of fitting the comparison ledger, but it is not a finite theorem until a threshold operator is derived.",
		}}
	}}
}

func noStrictCombination(xs []CombinationAttempt) bool {
	for _, x := range xs {
		if x.AdmissibleAsFiniteTheorem || x.CanRepairGate177Strictly {
			return false
		}
	}
	return true
}

func hasObservedCombination(xs []CombinationAttempt) bool {
	for _, x := range xs {
		if x.UsesObservedComparison {
			return true
		}
	}
	return false
}
