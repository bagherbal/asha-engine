package contactequivrefinement

import (
	"fmt"

	"github.com/bagherbal/asha-engine/pkg/theorem"
)

func ContactRowEquivalenceRefinementSectorPatternMismatchObstructionTheorem() theorem.Theorem {
	const id = "BRIDGE-CONTACT-ROW-EQUIVALENCE-REFINEMENT-SECTOR-PATTERN-MISMATCH"
	const name = "contact-row equivalence refinement / sector-pattern mismatch obstruction theorem"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.Variational, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "build contact-row refinement mismatch search", Passed: false, Detail: err.Error()}}}
		}
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.Variational, Checks: []theorem.Check{
			{Name: "Gate 128 sector-quotient obstruction inherited", Passed: a.Previous.CurrentSideQuotientSemanticsFound && a.Previous.NaturalPatternsConflict && a.U4Dimension == 16 && a.ContactRows == 7, Detail: FormatSummary(a.Summary)},
			{Name: "current 1+6 pattern mismatches seven contact singletons", Passed: a.SectorPatternMismatch && a.CurrentPatternStable && a.ContactSingletonsStable && a.Summary.CurrentSectorPattern == "1+6" && a.Summary.ContactSingletonPattern == "1+1+1+1+1+1+1", Detail: FormatCandidates(a.Candidates)},
			{Name: "row-resolving current refinements require hidden assignments", Passed: !a.CanonicalRefinementDerived && !a.CurrentDerivedRefinement && a.HiddenAssignmentRequired && a.Summary.MinimalHiddenChoicesPerBranch == 5040 && a.Summary.TotalHiddenBranchChoices >= 10080, Detail: FormatCandidates(a.Candidates)},
			{Name: "Fano or observed refinements are rejected as non-native selectors", Passed: a.FanoChoiceRequired && a.ObservedInputRejected && a.ArbitraryCutoffRejected && !a.HiddenObservedInputUsed, Detail: FormatCandidates(a.Candidates)},
			{Name: "contact beta firewall remains closed", Passed: a.RepresentationCompleteRows == 0 && a.RepresentationOpenRows == 7 && a.ContactBetaRowsAllowed == 0 && a.ContactZeroRowsProved == 0 && !a.ThresholdCorrectedBetaDerived && !a.FullBetaMatchingTensorDerived, Detail: FormatRows(a.Rows, 7)},
			{Name: "refinement search does not leak physical constants", Passed: a.ResidualNullityBefore == 3 && a.ResidualNullityAfter == 3 && !a.HiddenObservedInputUsed && !a.PhysicalWeakAngleDerived && !a.FineStructureDerived && !a.PhysicalMassesDerived && !a.PhysicalScaleDerived, Detail: fmt.Sprintf("nullity=%d->%d; no alpha/thetaW/masses/M*/g*", a.ResidualNullityBefore, a.ResidualNullityAfter)},
		}, Notes: []string{a.TruthStatement, "criteria: " + FormatCriteria(a.Criteria), "candidates: " + FormatCandidates(a.Candidates), "rows: " + FormatRows(a.Rows, 7), "rejected claims: " + Join(a.RejectedClaims), "remaining unknowns: " + Join(a.RemainingUnknowns), "Next: " + a.RecommendedNextGate}}
	}}
}
