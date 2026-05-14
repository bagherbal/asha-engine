package contactsignsource

import (
	"fmt"

	"github.com/bagherbal/asha-engine/pkg/theorem"
)

func ContactSignOrientationSourceChargeConjugationObstructionTheorem() theorem.Theorem {
	const id = "BRIDGE-CONTACT-SIGN-ORIENTATION-SOURCE-CHARGE-CONJUGATION-OBSTRUCTION"
	const name = "Contact sign-orientation source / charge-conjugation symmetry obstruction theorem"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.Variational, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "build contact sign-orientation source search", Passed: false, Detail: err.Error()}}}
		}
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.Variational, Checks: []theorem.Check{
			{Name: "Gate 141 two-orientation 3|4 diagnostic is inherited", Passed: a.ContactRows == 7 && a.LargestGapHighRows == 3 && a.LargestGapLowRows == 4 && a.OrientationCandidates == 2 && a.SplitPattern == "3|4", Detail: fmt.Sprintf("split=%s rows=%d", a.SplitPattern, a.ContactRows)},
			{Name: "orientation-source audit enumerates finite and forbidden candidates", Passed: a.OrientationSourcesAudited == 7 && a.SourcesAvailable == 3, Detail: FormatSourceAudits(a.SourceAudits)},
			{Name: "charge conjugation is an involution but exchanges orientations instead of selecting one", Passed: a.ChargeConjugationInvolutions == 1 && a.Z2OrientationDegeneracy && a.ChargeConjugationSelectedBranches == 0 && a.SourcesSelectingOrientation == 0, Detail: FormatChargeConjugation(a.ChargeConjugation)},
			{Name: "B-L, chirality, SU2L, and hypercharge still do not pull back to contact rows", Passed: a.T3RPullbackRowsDerived == 0 && a.ChiralityPullbackRowsDerived == 0 && a.BMinusLPullbackRowsDerived == 0 && a.SU2LPullbackRowsDerived == 0 && a.HyperchargeRowsDerived == 0, Detail: FormatSummary(a.Summary)},
			{Name: "no T3R-semantic orientation or traceless seven-row sign generator is derived", Passed: a.T3RSemanticOrientations == 0 && a.TracelessOrientations == 0, Detail: Join(a.RejectedClaims)},
			{Name: "contact beta firewall, S6 ambiguity, nullity, and no-observed-input discipline remain sealed", Passed: a.BetaPermissionFirewallClosed && a.RepresentationCompleteRows == 0 && a.RepresentationOpenRows == 7 && a.ContactBetaRowsAllowed == 0 && a.ContactZeroRowsProved == 0 && !a.ThresholdCorrectedBeta && !a.FullBetaMatchingTensor && a.ResidualS6Choices == 720 && a.ResidualNullityBefore == 3 && a.ResidualNullityAfter == 3 && !a.HiddenObservedInputUsed && !a.PhysicalWeakAngleDerived && !a.FineStructureDerived && !a.PhysicalMassesDerived && !a.PhysicalScaleDerived, Detail: a.TruthStatement},
		}}
	}}
}
