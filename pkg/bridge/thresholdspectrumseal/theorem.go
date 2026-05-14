package thresholdspectrumseal

import "github.com/bagherbal/asha-engine/pkg/theorem"

func ThresholdSpectrumSealMatchingCorrectionTwoLoopPreflightTheorem() theorem.Theorem {
	const id = "BRIDGE-THRESHOLD-SPECTRUM-SEAL-MATCHING-CORRECTION-TWO-LOOP-PREFLIGHT-AUDIT"
	const name = "ThresholdSpectrumSeal / matching-correction and two-loop stability preflight audit"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.Phenomenology, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "construct Gate 213 audit", Passed: false, Detail: err.Error()}}}
		}
		checks := []theorem.Check{
			{Name: "ThresholdSpectrumSeal is introduced without unique finite-spectrum claim", Passed: a.Seal.DegeneracyQuarantined && a.Seal.SelectedPairPhenomenological && !a.Seal.FiniteUniquenessClaimed && !a.Seal.ContactOrBGapOriginClaimed, Detail: FormatSeal(a.Seal) + " :: " + FormatSubject(a.Subject)},
			{Name: "Gate-211 ranked witness is selected only as sealed test subject", Passed: a.Subject.ConditionalOnly && !a.Subject.FiniteDerived && a.Subject.Row1Rep == "(1,3,Y=1)" && a.Subject.Row2Rep == "(8,2,Y=1/2)" && a.Subject.SelectedFromOrderedRank == 1, Detail: FormatSubject(a.Subject)},
			{Name: "finite matching corrections remain obstructed", Passed: a.Matching.Status == MatchingCorrectionsFailed && !a.Matching.ThresholdMatchingCoefficientsDerived && !a.Matching.CanonicalSubtractionSchemeDerived && !a.Matching.MSbarOrDimRegImported, Detail: FormatMatching(a.Matching)},
			{Name: "exact symbolic heavy two-loop coefficients are computed as preflight only", Passed: a.TwoLoop.ExactSymbolicHeavyCoefficients && a.TwoLoop.UsesStandardQFTFormula && !a.TwoLoop.ImportedAsFiniteCore && !a.TwoLoop.SchemeIndependentForFiniteCore, Detail: FormatTwoLoop(a.TwoLoop)},
			{Name: "two-loop stability is not promoted beyond one-loop sealed phenomenology", Passed: a.Stability.OneLoopScalesValidOnlyAtOneLoop && a.Stability.RequiresFullTwoLoopIntegration && a.Stability.RequiresMatchingCorrections && a.Stability.Status == TwoLoopWarning, Detail: FormatStability(a.Stability)},
			{Name: "firewalls remain closed", Passed: a.Firewall.Gate212Inherited && a.Firewall.ThresholdSpectrumSealIntroduced && a.Firewall.LeptoquarkDynamicsSealInherited && a.Firewall.EmpiricalCarrierSealInherited && a.Firewall.EmpiricalLedgerQuarantined && !a.Firewall.UniquePhysicalSpectrumClaimed && !a.Firewall.ContactModesPromotedToCarriers && !a.Firewall.BGapPromotedToMass && !a.Firewall.MatchingCorrectionsDerived && !a.Firewall.MSbarImportedAsFiniteCore && !a.Firewall.TwoLoopCoefficientsFiniteDerived && !a.Firewall.TwoLoopScalesClaimedAsPrediction && !a.Firewall.PhysicalPredictionClaimed && !a.Firewall.ProtonLifetimeComputed, Detail: FormatFirewall(a.Firewall)},
		}
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.Phenomenology, Checks: checks, Notes: []string{a.TruthStatement, "CONDITIONAL_PHENOMENOLOGY: Gate 213 seals a chosen spectrum for preflight only; matching corrections and corrected two-loop scales remain un-derived."}}
	}}
}
