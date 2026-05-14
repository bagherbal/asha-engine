package twothresholdminimality

import "github.com/bagherbal/asha-engine/pkg/theorem"

func TwoThresholdSolutionMinimalityFiniteOriginParentageAuditTheorem() theorem.Theorem {
	const id = "BRIDGE-TWO-THRESHOLD-SOLUTION-MINIMALITY-FINITE-ORIGIN-PARENTAGE-AUDIT"
	const name = "two-threshold solution minimality / finite-origin and multiplet-parentage audit"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "construct Gate 212 audit", Passed: false, Detail: err.Error()}}}
		}
		checks := []theorem.Check{
			{Name: "Gate 211 viable witness set is inherited", Passed: a.Gate211.Gate211Inherited && a.Gate211.ConditionalViabilityInherited && a.Summary.OrderedViablePairs == 44 && a.Summary.UnorderedPairClasses == 22, Detail: FormatSummary(a.Summary)},
			{Name: "finite-origin combinatorics do not promote contact/B-sector modes", Passed: a.FiniteOrigin.CarrierActivationSealIntact && a.FiniteOrigin.CanonicalFiniteOriginMatches == 0, Detail: FormatFiniteOrigin(a.FiniteOrigin)},
			{Name: "multiplet-parentage preflight finds no derived complete parent", Passed: !a.Parentage.UnifiedParentGaugeImported && !a.Parentage.ThresholdSplittingRuleDerived && a.Parentage.CompleteParentageDerived == 0, Detail: FormatParentage(a.Parentage)},
			{Name: "Gate-211 ranking is not a finite canonical selector", Passed: a.Degeneracy.Gate211RankedBestExists && !a.Degeneracy.Gate211RankingIsFiniteMetric && !a.Degeneracy.CanonicalUniquePairFound && a.Degeneracy.ThresholdSpectrumSealRequired, Detail: FormatDegeneracy(a.Degeneracy)},
			{Name: "firewalls remain closed", Passed: a.Firewall.Gate211Inherited && a.Firewall.LeptoquarkDynamicsSealInherited && a.Firewall.EmpiricalCarrierSealInherited && a.Firewall.EmpiricalLedgerQuarantined && !a.Firewall.ContactModesPromotedToCarriers && !a.Firewall.BGapPromotedToMass && !a.Firewall.SU5OrPatiSalamGaugeImported && !a.Firewall.MatchingCorrectionsDerived && !a.Firewall.PhysicalPredictionClaimed && !a.Firewall.ProtonLifetimeComputed && !a.Firewall.UniqueThresholdSpectrumClaimed, Detail: FormatFirewall(a.Firewall)},
		}
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: checks, Notes: []string{a.TruthStatement, "FAILED_ROUTE is the correct result: Gate 212 falsifies canonical uniqueness, not Gate 211 viability."}}
	}}
}
