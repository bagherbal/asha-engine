package contactlqblock

import (
	"fmt"

	"github.com/bagherbal/asha-engine/pkg/theorem"
)

func ContactLeptoquarkSixBlockS6PermutationObstructionTheorem() theorem.Theorem {
	const id = "BRIDGE-CONTACT-LEPTOQUARK-SIX-BLOCK-S6-PERMUTATION-OBSTRUCTION"
	const name = "contact leptoquark six-block symmetry / S6 permutation obstruction theorem"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.Variational, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "build contact leptoquark S6 obstruction", Passed: false, Detail: err.Error()}}}
		}
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.Variational, Checks: []theorem.Check{
			{Name: "Gate 130 singlet/leptoquark obstruction inherited", Passed: a.Previous.SingletChoiceRequired && a.Previous.PermutationRequired && a.ContactRows == 7 && a.LeptoquarkRows == 6, Detail: FormatSummary(a.Summary)},
			{Name: "six-block S6 ambiguity is measured", Passed: a.S6PermutationObstruction && a.SixPermutationOrder == 720 && a.AssignmentsPerBranch == 5040 && a.TotalCurrentAssignments == 10080, Detail: FormatBlocks(a.Blocks, 2)},
			{Name: "anonymous six-block is canonical but row-blind", Passed: a.SixBlockExists && a.AnonymousBlockCanonical && !a.CurrentNaturalSixOrder && !a.CanonicalCurrentAssignmentDerived, Detail: FormatStrategies(a.Strategies)},
			{Name: "spectral six-row orderings are diagnostic not current-natural", Passed: a.SpectralOrderingAvailable && a.SpectralOrientationAmbiguous && !a.CurrentNaturalSixOrder, Detail: FormatBlocks(a.Blocks, 2)},
			{Name: "Fano/observed selectors remain forbidden", Passed: a.FanoChoiceRequired && a.ObservedInputRejected && !a.HiddenObservedInputUsed, Detail: FormatStrategies(a.Strategies)},
			{Name: "contact beta firewall remains closed", Passed: a.RepresentationCompleteRows == 0 && a.RepresentationOpenRows == 7 && a.ContactBetaRowsAllowed == 0 && a.ContactZeroRowsProved == 0 && !a.ThresholdCorrectedBetaDerived && !a.FullBetaMatchingTensorDerived, Detail: fmt.Sprintf("repRows=%d open=%d beta=%d zero=%d", a.RepresentationCompleteRows, a.RepresentationOpenRows, a.ContactBetaRowsAllowed, a.ContactZeroRowsProved)},
			{Name: "S6 search does not leak physical constants", Passed: a.ResidualNullityBefore == 3 && a.ResidualNullityAfter == 3 && !a.HiddenObservedInputUsed && !a.PhysicalWeakAngleDerived && !a.FineStructureDerived && !a.PhysicalMassesDerived && !a.PhysicalScaleDerived, Detail: fmt.Sprintf("nullity=%d->%d; no alpha/thetaW/masses/M*/g*", a.ResidualNullityBefore, a.ResidualNullityAfter)},
		}, Notes: []string{a.TruthStatement, "criteria: " + FormatCriteria(a.Criteria), "strategies: " + FormatStrategies(a.Strategies), "blocks: " + FormatBlocks(a.Blocks, 3), "rejected claims: " + Join(a.RejectedClaims), "remaining unknowns: " + Join(a.RemainingUnknowns), "Next: " + a.RecommendedNextGate}}
	}}
}
