package contactsignsplit

import (
	"fmt"

	"github.com/bagherbal/asha-engine/pkg/theorem"
)

func ContactT3RSignSplitNaturalitySpectralCutObstructionTheorem() theorem.Theorem {
	const id = "BRIDGE-CONTACT-T3R-SIGN-SPLIT-NATURALITY-SPECTRAL-CUT-OBSTRUCTION"
	const name = "Contact T3R sign-split naturality / spectral-cut obstruction theorem"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.Variational, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "build contact sign-split spectral-cut search", Passed: false, Detail: err.Error()}}}
		}
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.Variational, Checks: []theorem.Check{
			{Name: "seven-row contact spectrum remains distinct and ordered", Passed: a.ContactRows == 7 && a.DistinctSpectralRows == 7, Detail: FormatSpectrum(a.SpectrumDescending)},
			{Name: "proper spectral cuts are enumerated", Passed: a.ProperSpectralCuts == 6 && a.SpectralCutSignAssignments == 12 && a.AbstractSignAssignments == 128, Detail: FormatCuts(a.Cuts)},
			{Name: "unique largest spectral gap gives only a diagnostic 3|4 split", Passed: a.UniqueLargestGapCuts == 1 && a.LargestGapHighRows == 3 && a.LargestGapLowRows == 4 && a.CanonicalDiagnosticSplits == 1, Detail: fmt.Sprintf("largest gap %.10f pattern=%s", a.LargestGap.Gap, a.LargestGap.MultiplicityPattern)},
			{Name: "spectral cut orientation and T3R semantics are not selected", Passed: a.LargestGapOrientations == 2 && a.OrientationSelectedSplits == 0 && a.T3RSemanticSplits == 0 && a.T3RTargetOperatorsDerived == 0 && a.QuotientInducedT3RTargetOps == 0, Detail: fmt.Sprintf("orientations=%d selected=%d T3R=%d", a.LargestGapOrientations, a.OrientationSelectedSplits, a.T3RSemanticSplits)},
			{Name: "no Fock-contact operator intertwiner or contact hypercharge rows are derived", Passed: a.FullOperatorIntertwiners == 0 && a.T3RPullbackRowsDerived == 0 && a.ChiralityPullbackRowsDerived == 0 && a.BMinusLPullbackRowsDerived == 0 && a.SU2LPullbackRowsDerived == 0 && a.HyperchargeRowsDerived == 0, Detail: FormatSummary(a.Summary)},
			{Name: "contact beta firewall, residual S6 ambiguity, nullity, and no-observed-input discipline are preserved", Passed: a.BetaPermissionFirewallClosed && a.RepresentationCompleteRows == 0 && a.RepresentationOpenRows == 7 && a.ContactBetaRowsAllowed == 0 && a.ContactZeroRowsProved == 0 && !a.ThresholdCorrectedBeta && !a.FullBetaMatchingTensor && a.ResidualS6Choices == 720 && a.ResidualNullityBefore == 3 && a.ResidualNullityAfter == 3 && !a.HiddenObservedInputUsed && !a.PhysicalWeakAngleDerived && !a.FineStructureDerived && !a.PhysicalMassesDerived && !a.PhysicalScaleDerived, Detail: FormatSummary(a.Summary)},
		}, Notes: []string{a.TruthStatement, "summary: " + FormatSummary(a.Summary), "cuts: " + FormatCuts(a.Cuts), "rejected claims: " + Join(a.RejectedClaims), "remaining unknowns: " + Join(a.RemainingUnknowns), "Next: " + a.RecommendedNextGate}}
	}}
}
