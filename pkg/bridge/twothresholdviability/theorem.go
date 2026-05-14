package twothresholdviability

import "github.com/bagherbal/asha-engine/pkg/theorem"

func TwoThresholdRationalLatticeViabilityFilterTheorem() theorem.Theorem {
	const id = "BRIDGE-TWO-THRESHOLD-RATIONAL-LATTICE-VIABILITY-FILTER"
	const name = "two-threshold rational lattice viability filter / scale-ordered Landau safety audit"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.Phenomenology, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "construct Gate 211 audit", Passed: false, Detail: err.Error()}}}
		}
		checks := []theorem.Check{
			{Name: "Gate 210 failed-route and LeptoquarkDynamicsSeal are inherited", Passed: a.Gate210.Gate210Inherited && a.Gate210.SingleScaleFailedRoutePreserved && a.Gate210.Gate209LeptoquarkSealInherited && a.Gate210.UniversalBetaRowKilled, Detail: a.Gate210.TruthStatement},
			{Name: "quarantined u-space RG ledger is used only phenomenologically", Passed: a.Inputs.ObservedLedgerQuarantined && !a.Inputs.UsedForFiniteCoreDerivation && !a.Inputs.UniversalBetaRowAllowed && a.Inputs.TwoThresholdSystem, Detail: FormatInputs(a.Inputs)},
			{Name: "Gate-210 safe generator basis is inherited exactly", Passed: a.GeneratorAudit.InheritsGate210FilteredBasis && a.GeneratorAudit.SafeGenerators == 108 && a.GeneratorAudit.AnomalyCompatibleGenerators == a.GeneratorAudit.SafeGenerators && a.GeneratorAudit.LeptoquarkSealCompatibleRows == a.GeneratorAudit.SafeGenerators && a.GeneratorAudit.NoUniversalBetaRowInserted && a.GeneratorAudit.NoContinuousRowCoefficients, Detail: FormatGeneratorAudit(a.GeneratorAudit)},
			{Name: "both boundary targets are audited", Passed: len(a.TargetAudits) == 2 && a.TargetAudits[0].OrderedPairsAudited > 0 && a.TargetAudits[1].OrderedPairsAudited > 0, Detail: FormatTargetAudit(a.TargetAudits[0]) + " :: " + FormatTargetAudit(a.TargetAudits[1])},
			{Name: "topological branch has exact viable two-threshold witnesses", Passed: a.Summary.ViableTopological > 0 && firstTarget(a, "u_topological").ViablePairs == a.Summary.ViableTopological && allViableSolutionsPass(firstTarget(a, "u_topological")), Detail: FormatTargetAudit(firstTarget(a, "u_topological"))},
			{Name: "centroid branch is filtered separately", Passed: firstTarget(a, "u_centroid").OrderedPairsAudited > 0 && firstTarget(a, "u_centroid").ViablePairs == a.Summary.ViableCentroid, Detail: FormatTargetAudit(firstTarget(a, "u_centroid"))},
			{Name: "baryon/anomaly filters preserve the leptoquark seal", Passed: a.BaryonAnomaly.LeptoquarkDynamicsSealInherited && a.BaryonAnomaly.AllSearchRowsAnomalyCompatible && a.BaryonAnomaly.AllSearchRowsSealCompatible && a.BaryonAnomaly.ViablePairsAnomalyCompatible && a.BaryonAnomaly.ViablePairsSealCompatible && !a.BaryonAnomaly.ProtonDecayOperatorUsed && !a.BaryonAnomaly.ProtonLifetimeComputed, Detail: FormatBaryonAnomaly(a.BaryonAnomaly)},
			{Name: "B-sector/contact data are not promoted to threshold rows", Passed: a.ContactMatch.BGapAudited && a.ContactMatch.ContactPartialOverlapAudited && !a.ContactMatch.CanonicalNumericMatchFound && !a.ContactMatch.ChargeSpinMassSemanticsFound && !a.ContactMatch.ViableRowsPromotedFromContact, Detail: FormatContactMatch(a.ContactMatch)},
			{Name: "firewalls remain intact", Passed: a.Firewall.Gate210Inherited && a.Firewall.LeptoquarkDynamicsSealInherited && a.Firewall.EmpiricalLedgerQuarantined && !a.Firewall.ObservedLedgerUsedForFiniteCore && !a.Firewall.UniversalBetaRowInserted && !a.Firewall.ArbitraryRealRowCoefficientInserted && !a.Firewall.PhysicalPredictionClaimed && !a.Firewall.AbsoluteMassDerivedFromFiniteCore && !a.Firewall.ProtonDecaySealViolated && !a.Firewall.ProtonLifetimeComputed && !a.Firewall.MatchingCorrectionsDerived, Detail: FormatFirewall(a.Firewall) + " :: " + FormatSummary(a.Summary)},
		}
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.Phenomenology, Checks: checks, Notes: []string{a.Summary.Status, a.TruthStatement}}
	}}
}

func firstTarget(a Analysis, name string) TargetAudit {
	for _, t := range a.TargetAudits {
		if t.Target.Name == name {
			return t
		}
	}
	return TargetAudit{}
}

func allViableSolutionsPass(t TargetAudit) bool {
	if t.ViablePairs == 0 {
		return false
	}
	for _, s := range t.AllViableSolutions {
		if !s.Viable || !s.ExactLinearClosure || !s.ScaleOrdered || !s.DistinctThresholds || !s.SubPlanck || !s.PositiveCouplingsToMStar || !s.NoSubPlanckLandauPole || !s.AnomalyCompatible || !s.LeptoquarkSealCompatible {
			return false
		}
	}
	return true
}
