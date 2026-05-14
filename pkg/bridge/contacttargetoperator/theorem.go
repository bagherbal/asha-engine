package contacttargetoperator

import (
	"fmt"

	"github.com/bagherbal/asha-engine/pkg/theorem"
)

func ContactTargetOperatorReconstructionQuotientSideT3RSpectrumSearchTheorem() theorem.Theorem {
	const id = "BRIDGE-CONTACT-TARGET-OPERATOR-RECONSTRUCTION-QUOTIENT-T3R-SPECTRUM-SEARCH"
	const name = "Contact target-operator reconstruction / quotient-side T3R spectrum search"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.Variational, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "build contact target-operator search", Passed: false, Detail: err.Error()}}}
		}
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.Variational, Checks: []theorem.Check{
			{Name: "seven distinct contact spectral rows are inherited", Passed: a.ContactRows == 7 && a.ContactSpectralRows == 7 && a.ContactSpectralDistinctRows == 7, Detail: "spectrum=" + FormatSpectrum(a.ContactSpectrum)},
			{Name: "canonical contact spectral operator is diagnostic, not T3R", Passed: a.DiagnosticContactOperators >= 1 && a.CanonicalT3RTargetOperators == 0, Detail: FormatCandidates(a.Candidates)},
			{Name: "abstract quotient-side T3R splits remain noncanonical", Passed: a.AbstractT3RMultiplicitySplits == 8 && a.T3RRowSignAssignments == 128 && a.NonScalarT3RRowSignAssignments == 126, Detail: FormatSplits(a.Splits)},
			{Name: "no quotient-induced T3R target or full Fock-contact intertwiner is derived", Passed: a.QuotientInducedT3RTargetOperators == 0 && a.FullOperatorIntertwinersDerived == 0 && a.CanonicalFockKernels == 1 && a.T3RPullbackRowsDerived == 0 && a.ChiralityPullbackRowsDerived == 0, Detail: fmt.Sprintf("quotientT3R=%d intertwiners=%d kernels=%d", a.QuotientInducedT3RTargetOperators, a.FullOperatorIntertwinersDerived, a.CanonicalFockKernels)},
			{Name: "contact hypercharge and beta firewall remain closed", Passed: a.BetaPermissionFirewallClosed && a.HyperchargeRowsDerived == 0 && a.RepresentationCompleteRows == 0 && a.RepresentationOpenRows == 7 && a.ContactBetaRowsAllowed == 0 && a.ContactZeroRowsProved == 0 && !a.ThresholdCorrectedBetaDerived && !a.FullBetaMatchingTensorDerived, Detail: FormatSummary(a.Summary)},
			{Name: "residual S6 ambiguity, physical-flow nullity, and no-observed-input discipline are preserved", Passed: a.ResidualS6Choices == 720 && a.ResidualNullityBefore == 3 && a.ResidualNullityAfter == 3 && !a.HiddenObservedInputUsed && !a.PhysicalWeakAngleDerived && !a.FineStructureDerived && !a.PhysicalMassesDerived && !a.PhysicalScaleDerived, Detail: fmt.Sprintf("S6=%d nullity=%d→%d", a.ResidualS6Choices, a.ResidualNullityBefore, a.ResidualNullityAfter)},
		}, Notes: []string{a.TruthStatement, "summary: " + FormatSummary(a.Summary), "candidates: " + FormatCandidates(a.Candidates), "splits: " + FormatSplits(a.Splits), "rejected claims: " + Join(a.RejectedClaims), "remaining unknowns: " + Join(a.RemainingUnknowns), "Next: " + a.RecommendedNextGate}}
	}}
}
