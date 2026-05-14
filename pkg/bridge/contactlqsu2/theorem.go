package contactlqsu2

import (
	"fmt"

	"github.com/bagherbal/asha-engine/pkg/theorem"
)

func LeptoquarkRealOrientationWeakDoubletSU2ActionSearchTheorem() theorem.Theorem {
	const id = "BRIDGE-LEPTOQUARK-REAL-ORIENTATION-WEAK-DOUBLET-SU2-ACTION-SEARCH"
	const name = "leptoquark real-orientation versus weak-doublet obstruction / SU(2)L action search"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.Variational, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "build leptoquark SU(2)L action search", Passed: false, Detail: err.Error()}}}
		}
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.Variational, Checks: []theorem.Check{
			{Name: "Gate 132 leptoquark real tensor inherited", Passed: a.S6ObstructionInherited && a.LeptoquarkRows == 6 && a.ColorPlanes == 3 && a.RealOrientationsPerColor == 2 && a.TotalCurrentLQSlots == 6, Detail: FormatSummary(a.Summary)},
			{Name: "orientation SO(2) action exposed", Passed: a.OrientationSO2Available && a.ColorWiseSO2Available && a.DiagonalSO2Available && a.OrientationActionAbelian, Detail: FormatPlanes(a.Planes)},
			{Name: "SO(2) orientation is not SU(2)L", Passed: !a.NonAbelianSU2TripleDerived && !a.SU2CommutationDerived && !a.SU2WeakDoubletActionDerived && !a.WeakDoubletSemanticsDerived, Detail: FormatCandidates(a.Candidates)},
			{Name: "borrowing matter SU(2)L is rejected", Passed: a.BorrowedMatterActionRejected && a.SemanticBridgeMissing && !a.CurrentNaturalSU2Action, Detail: fmt.Sprintf("borrowedRejected=%t semanticBridgeMissing=%t currentNaturalSU2=%t", a.BorrowedMatterActionRejected, a.SemanticBridgeMissing, a.CurrentNaturalSU2Action)},
			{Name: "contact leptoquark representation rows remain forbidden", Passed: a.RepresentationCompleteRows == 0 && a.RepresentationOpenRows == 7 && a.ContactBetaRowsAllowed == 0 && a.ContactZeroRowsProved == 0 && !a.ThresholdCorrectedBetaDerived && !a.FullBetaMatchingTensorDerived, Detail: fmt.Sprintf("repRows=%d open=%d beta=%d zero=%d", a.RepresentationCompleteRows, a.RepresentationOpenRows, a.ContactBetaRowsAllowed, a.ContactZeroRowsProved)},
			{Name: "SU(2)L action search does not leak physical constants", Passed: a.ResidualNullityBefore == 3 && a.ResidualNullityAfter == 3 && !a.HiddenObservedInputUsed && !a.PhysicalWeakAngleDerived && !a.FineStructureDerived && !a.PhysicalMassesDerived && !a.PhysicalScaleDerived, Detail: fmt.Sprintf("nullity=%d->%d; no alpha/thetaW/masses/M*/g*", a.ResidualNullityBefore, a.ResidualNullityAfter)},
		}, Notes: []string{a.TruthStatement, "criteria: " + FormatCriteria(a.Criteria), "planes: " + FormatPlanes(a.Planes), "candidates: " + FormatCandidates(a.Candidates), "rejected claims: " + Join(a.RejectedClaims), "remaining unknowns: " + Join(a.RemainingUnknowns), "Next: " + a.RecommendedNextGate}}
	}}
}
