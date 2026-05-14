package contactassignment

import (
	"fmt"

	"github.com/bagherbal/asha-engine/pkg/theorem"
)

func ContactSingletLeptoquarkAssignmentNaturalityPermutationObstructionTheorem() theorem.Theorem {
	const id = "BRIDGE-CONTACT-SINGLET-LEPTOQUARK-ASSIGNMENT-NATURALITY-OBSTRUCTION"
	const name = "contact singlet/leptoquark assignment naturality / permutation obstruction theorem"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.Variational, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "build contact singlet/leptoquark assignment search", Passed: false, Detail: err.Error()}}}
		}
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.Variational, Checks: []theorem.Check{
			{Name: "Gate 129 sector-pattern refinement obstruction inherited", Passed: a.Previous.SectorPatternMismatch && a.Previous.HiddenAssignmentRequired && a.ContactRows == 7 && a.CurrentPattern == "1+6", Detail: FormatSummary(a.Summary)},
			{Name: "spectral rows are distinct but only diagnostic selectors", Passed: a.SpectralRowsDistinct && a.ContactDiagnosticSelectors >= 3 && !a.CurrentNaturalSelector, Detail: FormatExtrema(a.MinSpectralRow, a.MaxSpectralRow, a.MedianSpectralRow)},
			{Name: "current-side singlet/leptoquark pattern gives no natural contact assignment", Passed: a.SingletChoiceRequired && a.PermutationRequired && !a.CanonicalAssignmentDerived && !a.CurrentDerivedAssignment && a.Summary.MinimalSingletChoices == 7 && a.Summary.MinimalPermutationChoices == 720 && a.Summary.TotalHiddenBranchChoices >= 10080, Detail: FormatCandidates(a.Candidates)},
			{Name: "Fano/observed/orientation selectors are rejected", Passed: a.FanoChoiceRequired && a.ObservedInputRejected && a.ArbitraryOrientationRejected && !a.HiddenObservedInputUsed, Detail: FormatCandidates(a.Candidates)},
			{Name: "contact beta firewall remains closed", Passed: a.RepresentationCompleteRows == 0 && a.RepresentationOpenRows == 7 && a.ContactBetaRowsAllowed == 0 && a.ContactZeroRowsProved == 0 && !a.ThresholdCorrectedBetaDerived && !a.FullBetaMatchingTensorDerived, Detail: FormatRows(a.Rows, 7)},
			{Name: "assignment search does not leak physical constants", Passed: a.ResidualNullityBefore == 3 && a.ResidualNullityAfter == 3 && !a.HiddenObservedInputUsed && !a.PhysicalWeakAngleDerived && !a.FineStructureDerived && !a.PhysicalMassesDerived && !a.PhysicalScaleDerived, Detail: fmt.Sprintf("nullity=%d->%d; no alpha/thetaW/masses/M*/g*", a.ResidualNullityBefore, a.ResidualNullityAfter)},
		}, Notes: []string{a.TruthStatement, "criteria: " + FormatCriteria(a.Criteria), "candidates: " + FormatCandidates(a.Candidates), "rows: " + FormatRows(a.Rows, 7), "rejected claims: " + Join(a.RejectedClaims), "remaining unknowns: " + Join(a.RemainingUnknowns), "Next: " + a.RecommendedNextGate}}
	}}
}
