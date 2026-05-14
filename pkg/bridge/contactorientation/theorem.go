package contactorientation

import (
	"fmt"

	"github.com/bagherbal/asha-engine/pkg/theorem"
)

func ContactSpectralGapOrientationSignChoiceObstructionTheorem() theorem.Theorem {
	const id = "BRIDGE-CONTACT-SPECTRAL-GAP-ORIENTATION-SIGN-CHOICE-OBSTRUCTION"
	const name = "Contact spectral-gap orientation / sign-choice obstruction theorem"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.Variational, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "build contact spectral-gap orientation search", Passed: false, Detail: err.Error()}}}
		}
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.Variational, Checks: []theorem.Check{
			{Name: "Gate 140 unique 3|4 largest-gap split is inherited", Passed: a.ContactRows == 7 && a.LargestGapHighRows == 3 && a.LargestGapLowRows == 4 && a.SplitPattern == "3|4", Detail: fmt.Sprintf("split=%s spectrum=%s", a.SplitPattern, contactSpectrum(a.SpectrumDescending))},
			{Name: "both sign orientations are enumerated and spectrally compatible", Passed: a.OrientationCandidates == 2 && a.SpectrallyMonotoneOrientations == 2, Detail: FormatCandidates(a.Candidates)},
			{Name: "no sign orientation is selected", Passed: a.SelectedOrientations == 0 && a.T3RSemanticOrientations == 0, Detail: FormatSourceAudits(a.SourceAudits)},
			{Name: "+/- half-sign operator on seven rows is not traceless for either 3|4 orientation", Passed: a.TracelessOrientations == 0 && a.PureHalfSignTraceMagnitudeNumerator == 1 && a.PureHalfSignTraceMagnitudeDenom == 2, Detail: FormatCandidates(a.Candidates)},
			{Name: "no Fock-contact pullback, chirality, B-L, SU2L, or hypercharge rows are derived", Passed: a.FockContactIntertwiners == 0 && a.T3RPullbackRowsDerived == 0 && a.ChiralityPullbackRowsDerived == 0 && a.BMinusLPullbackRowsDerived == 0 && a.SU2LPullbackRowsDerived == 0 && a.HyperchargeRowsDerived == 0, Detail: FormatSummary(a.Summary)},
			{Name: "contact beta firewall, S6 ambiguity, nullity, and no-observed-input discipline remain sealed", Passed: a.BetaPermissionFirewallClosed && a.RepresentationCompleteRows == 0 && a.RepresentationOpenRows == 7 && a.ContactBetaRowsAllowed == 0 && a.ContactZeroRowsProved == 0 && !a.ThresholdCorrectedBeta && !a.FullBetaMatchingTensor && a.ResidualS6Choices == 720 && a.ResidualNullityBefore == 3 && a.ResidualNullityAfter == 3 && !a.HiddenObservedInputUsed && !a.PhysicalWeakAngleDerived && !a.FineStructureDerived && !a.PhysicalMassesDerived && !a.PhysicalScaleDerived, Detail: a.TruthStatement},
		}}
	}}
}

func contactSpectrum(xs []float64) string {
	// Reuse local formatting to avoid importing implementation-only helpers from Gate 140.
	if len(xs) == 0 {
		return "[]"
	}
	s := "["
	for i, x := range xs {
		if i > 0 {
			s += ", "
		}
		s += fmt.Sprintf("%.10f", x)
	}
	return s + "]"
}
