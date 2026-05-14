package contactu4kernel

import (
	"fmt"

	"github.com/bagherbal/asha-engine/pkg/theorem"
)

func U4ProjectionKernelCanonicalQuotientRelationSearchTheorem() theorem.Theorem {
	const id = "BRIDGE-U4-PROJECTION-KERNEL-CANONICAL-QUOTIENT-RELATION-SEARCH"
	const name = "u(4) projection kernel / canonical quotient relation search"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.Variational, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "build u(4) kernel quotient search", Passed: false, Detail: err.Error()}}}
		}
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.Variational, Checks: []theorem.Check{
			{Name: "Gate 126 u(4)->contact projection no-go inherited", Passed: a.Previous.SevenRowProjectionNoGoDerived && a.U4Dimension == 16 && a.TargetContactRows == 7 && a.RequiredKernelDimension == 9, Detail: FormatSummary(a.Summary)},
			{Name: "generic nine-dimensional kernels exist but are not selected", Passed: a.GenericKernelsExist && a.GrassmannKernelDimension == 63 && !a.CanonicalKernelDerived, Detail: "kernel choice lives in Gr(9,16), dimension 9*(16-9)=63"},
			{Name: "natural sector kernels are multiple and current-side only", Passed: a.NaturalNineDimensionalKernels == 2 && a.SectorKernelAmbiguity && a.ColorBLKernelDimension == 9 && a.CentralColorKernelDimension == 9 && a.SectorKernelsWrongSemantics && a.CurrentSideQuotientOnly, Detail: FormatKernels(a.KernelCandidates)},
			{Name: "no canonical contact quotient relation is derived", Passed: !a.CanonicalQuotientRelation && !a.ContactSemanticKernelDerived && a.KernelProjectionNoGoDerived && a.Summary.QuotientRelationsDerived == 0, Detail: FormatQuotients(a.QuotientCandidates)},
			{Name: "contact beta firewall remains closed", Passed: a.RepresentationCompleteRows == 0 && a.RepresentationOpenRows == 7 && a.ContactBetaRowsAllowed == 0 && a.ContactZeroRowsProved == 0 && !a.ThresholdCorrectedBetaDerived && !a.FullBetaMatchingTensorDerived, Detail: FormatRows(a.Rows, 7)},
			{Name: "kernel search does not leak physical constants", Passed: a.ResidualNullityBefore == 3 && a.ResidualNullityAfter == 3 && !a.HiddenObservedInputUsed && !a.PhysicalWeakAngleDerived && !a.FineStructureDerived && !a.PhysicalMassesDerived && !a.PhysicalScaleDerived, Detail: fmt.Sprintf("nullity=%d->%d; no alpha/thetaW/masses/M*/g*", a.ResidualNullityBefore, a.ResidualNullityAfter)},
		}, Notes: []string{a.TruthStatement, "criteria: " + FormatCriteria(a.Criteria), "sectors: " + FormatSectors(a.Sectors), "kernels: " + FormatKernels(a.KernelCandidates), "quotients: " + FormatQuotients(a.QuotientCandidates), "rejected claims: " + Join(a.RejectedClaims), "remaining unknowns: " + Join(a.RemainingUnknowns), "Next: " + a.RecommendedNextGate}}
	}}
}
