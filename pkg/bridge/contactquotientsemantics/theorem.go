package contactquotientsemantics

import (
	"fmt"

	"github.com/bagherbal/asha-engine/pkg/theorem"
)

func CurrentSideSectorQuotientSemanticsContactRowEquivalenceSearchTheorem() theorem.Theorem {
	const id = "BRIDGE-CURRENT-SIDE-SECTOR-QUOTIENT-CONTACT-ROW-EQUIVALENCE-SEARCH"
	const name = "current-side sector quotient semantics / contact-row equivalence relation search"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.Variational, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "build current-side quotient semantics search", Passed: false, Detail: err.Error()}}}
		}
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.Variational, Checks: []theorem.Check{
			{Name: "Gate 127 current-side quotient obstruction inherited", Passed: a.Previous.KernelProjectionNoGoDerived && a.Previous.CurrentSideQuotientOnly && a.U4Dimension == 16 && a.ContactRows == 7, Detail: FormatSummary(a.Summary)},
			{Name: "natural current quotients are sector patterns, not contact rows", Passed: a.NaturalCurrentQuotients == 2 && a.CurrentSideQuotientSemanticsFound && a.NaturalPatternsConflict && !a.CurrentToContactRelationDerived, Detail: FormatSectorQuotients(a.SectorQuotients)},
			{Name: "contact equivalence relation remains diagnostic only", Passed: a.ContactRowEquivalenceFound && a.RowPreservingRelationDerived && !a.CanonicalSemanticRelationDerived && a.Summary.CurrentContactRelations == 0, Detail: FormatContactRelations(a.ContactRelations)},
			{Name: "hidden assignment or arbitrary cutoff is required for non-diagnostic refinements", Passed: a.HiddenAssignmentRequired && a.ArbitraryCutoffRequired, Detail: FormatContactRelations(a.ContactRelations)},
			{Name: "contact beta firewall remains closed", Passed: a.RepresentationCompleteRows == 0 && a.RepresentationOpenRows == 7 && a.ContactBetaRowsAllowed == 0 && a.ContactZeroRowsProved == 0 && !a.ThresholdCorrectedBetaDerived && !a.FullBetaMatchingTensorDerived, Detail: FormatRows(a.Rows, 7)},
			{Name: "quotient semantics search does not leak physical constants", Passed: a.ResidualNullityBefore == 3 && a.ResidualNullityAfter == 3 && !a.HiddenObservedInputUsed && !a.PhysicalWeakAngleDerived && !a.FineStructureDerived && !a.PhysicalMassesDerived && !a.PhysicalScaleDerived, Detail: fmt.Sprintf("nullity=%d->%d; no alpha/thetaW/masses/M*/g*", a.ResidualNullityBefore, a.ResidualNullityAfter)},
		}, Notes: []string{a.TruthStatement, "criteria: " + FormatCriteria(a.Criteria), "sector quotients: " + FormatSectorQuotients(a.SectorQuotients), "contact relations: " + FormatContactRelations(a.ContactRelations), "rows: " + FormatRows(a.Rows, 7), "rejected claims: " + Join(a.RejectedClaims), "remaining unknowns: " + Join(a.RemainingUnknowns), "Next: " + a.RecommendedNextGate}}
	}}
}
