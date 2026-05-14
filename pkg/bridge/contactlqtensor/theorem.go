package contactlqtensor

import (
	"fmt"

	"github.com/bagherbal/asha-engine/pkg/theorem"
)

func ContactLeptoquarkSlotRepresentationTensorColorDoubletObstructionTheorem() theorem.Theorem {
	const id = "BRIDGE-CONTACT-LEPTOQUARK-SLOT-REPRESENTATION-TENSOR-COLOR-DOUBLET-OBSTRUCTION"
	const name = "contact leptoquark slot representation tensor / color-doublet semantic obstruction theorem"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.Variational, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "build contact leptoquark representation tensor obstruction", Passed: false, Detail: err.Error()}}}
		}
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.Variational, Checks: []theorem.Check{
			{Name: "Gate 131 S6 obstruction inherited", Passed: a.S6ObstructionInherited && a.LeptoquarkRows == 6 && a.Previous.SixPermutationOrder == 720 && a.Previous.AssignmentsPerBranch == 5040, Detail: FormatSummary(a.Summary)},
			{Name: "current leptoquark real tensor is derived", Passed: a.CurrentRealTensorDerived && a.CurrentLQSlots == 6 && a.ColorSlots == 3 && a.RealOrientationSlots == 2 && a.ColorTripletSemantics && a.RealOrientationSemantics, Detail: FormatSlots(a.Slots)},
			{Name: "color-doublet count trap is rejected", Passed: a.ColorWeakCountMatch && a.ColorDoubletCountTrap && a.SemanticBridgeMissing && !a.WeakDoubletSemanticsDerived && !a.HyperchargeSemanticsDerived && !a.LocalFieldSemanticsDerived, Detail: FormatCandidates(a.Candidates)},
			{Name: "no current-natural contact representation tensor", Passed: !a.CurrentNaturalRepresentation && a.RepresentationCompleteRows == 0 && a.RepresentationOpenRows == 7 && a.ContactBetaRowsAllowed == 0, Detail: fmt.Sprintf("repRows=%d open=%d beta=%d", a.RepresentationCompleteRows, a.RepresentationOpenRows, a.ContactBetaRowsAllowed)},
			{Name: "contact beta firewall remains closed", Passed: a.ContactZeroRowsProved == 0 && !a.ThresholdCorrectedBetaDerived && !a.FullBetaMatchingTensorDerived, Detail: fmt.Sprintf("zeroRows=%d thresholdBeta=%t fullTensor=%t", a.ContactZeroRowsProved, a.ThresholdCorrectedBetaDerived, a.FullBetaMatchingTensorDerived)},
			{Name: "leptoquark tensor search does not leak physical constants", Passed: a.ResidualNullityBefore == 3 && a.ResidualNullityAfter == 3 && !a.HiddenObservedInputUsed && !a.PhysicalWeakAngleDerived && !a.FineStructureDerived && !a.PhysicalMassesDerived && !a.PhysicalScaleDerived, Detail: fmt.Sprintf("nullity=%d->%d; no alpha/thetaW/masses/M*/g*", a.ResidualNullityBefore, a.ResidualNullityAfter)},
		}, Notes: []string{a.TruthStatement, "criteria: " + FormatCriteria(a.Criteria), "slots: " + FormatSlots(a.Slots), "candidates: " + FormatCandidates(a.Candidates), "rejected claims: " + Join(a.RejectedClaims), "remaining unknowns: " + Join(a.RemainingUnknowns), "Next: " + a.RecommendedNextGate}}
	}}
}
