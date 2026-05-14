package contactasymmetry

import (
	"fmt"

	"github.com/bagherbal/asha-engine/pkg/theorem"
)

func ContactChargeConjugationBreakingSourceAsymmetrySelectorSearchTheorem() theorem.Theorem {
	const id = "BRIDGE-CONTACT-CHARGE-CONJUGATION-BREAKING-SOURCE-ASYMMETRY-SELECTOR-SEARCH"
	const name = "Contact charge-conjugation breaking source / asymmetry selector search"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.Variational, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "build contact C-breaking asymmetry source search", Passed: false, Detail: err.Error()}}}
		}
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.Variational, Checks: []theorem.Check{
			{Name: "Gate 142 Z2 sign degeneracy and closed beta firewall are inherited", Passed: a.ContactRows == 7 && a.LargestGapHighRows == 3 && a.LargestGapLowRows == 4 && a.OrientationCandidates == 2 && a.Z2OrientationDegeneracy && a.BetaPermissionFirewallClosed, Detail: fmt.Sprintf("split=%s rows=%d", a.SplitPattern, a.ContactRows)},
			{Name: "contact spectrum has finite asymmetry diagnostics but they are C-even", Passed: a.CardinalityImbalance == 1 && a.AsymmetryDiagnostics == 2 && a.CInvariantDiagnostics == 2 && a.CBreakingDiagnostics == 0, Detail: FormatSplitAudit(a.SplitAudit)},
			{Name: "asymmetry-source audit enumerates finite, pullback, local-field, and forbidden selectors", Passed: a.AsymmetrySourcesAudited == 7 && a.AsymmetrySourcesAvailable == 2, Detail: FormatAsymmetrySources(a.SourceAudits)},
			{Name: "no charge-conjugation-breaking finite source or C-odd contact functional is derived", Passed: !a.ChargeConjugationBroken && a.CBreakingSources == 0 && a.CoddContactFunctionals == 0 && a.SourcesSelectingOrientation == 0, Detail: FormatSummary(a.Summary)},
			{Name: "B-L, chirality, T3R, SU2L, and hypercharge still do not pull back to contact rows", Passed: a.T3RPullbackRowsDerived == 0 && a.ChiralityPullbackRowsDerived == 0 && a.BMinusLPullbackRowsDerived == 0 && a.SU2LPullbackRowsDerived == 0 && a.HyperchargeRowsDerived == 0, Detail: Join(a.RemainingUnknowns)},
			{Name: "contact beta firewall, S6 ambiguity, nullity, and no-observed-input discipline remain sealed", Passed: a.RepresentationCompleteRows == 0 && a.RepresentationOpenRows == 7 && a.ContactBetaRowsAllowed == 0 && a.ContactZeroRowsProved == 0 && !a.ThresholdCorrectedBeta && !a.FullBetaMatchingTensor && a.ResidualS6Choices == 720 && a.ResidualNullityBefore == 3 && a.ResidualNullityAfter == 3 && !a.HiddenObservedInputUsed && !a.PhysicalWeakAngleDerived && !a.FineStructureDerived && !a.PhysicalMassesDerived && !a.PhysicalScaleDerived, Detail: a.TruthStatement},
		}}
	}}
}
