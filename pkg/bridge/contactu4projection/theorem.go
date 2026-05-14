package contactu4projection

import (
	"fmt"

	"github.com/bagherbal/asha-engine/pkg/theorem"
)

func ContactSevenRowTargetProjectionU4QuotientObstructionTheorem() theorem.Theorem {
	const id = "BRIDGE-CONTACT-SEVEN-ROW-TARGET-PROJECTION-U4-QUOTIENT-OBSTRUCTION"
	const name = "contact seven-row target projection / u(4)-to-contact quotient obstruction theorem"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.Variational, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "build contact u(4) projection obstruction", Passed: false, Detail: err.Error()}}}
		}
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.Variational, Checks: []theorem.Check{
			{Name: "Gate 125 target-enlargement no-go inherited", Passed: a.Previous.SevenRowTargetNoGoDerived && a.ContactRows == 7 && a.OpenContactRowsAfter == 7 && a.Previous.ContactBetaRowsAllowed == 0 && a.Previous.ContactZeroRowsProved == 0, Detail: FormatSummary(a.Summary)},
			{Name: "u(4)/Pati-Salam current inventory is typed but sixteen-dimensional", Passed: a.U4CurrentDimension == 16 && a.U4DecompositionCanonical && fmt.Sprint(a.U4SectorDimensions) == "[1 8 1 6]", Detail: FormatSectors(a.Sectors)},
			{Name: "abstract 16-to-7 projections exist but are not selected", Passed: a.RankSevenLinearMapsExist && a.GenericKernelDimension == 9 && a.ContinuousProjectionFamily && a.ContinuousProjectionFreeParameters == 63 && a.CanonicalProjectionCount == 0 && !a.U4ToContactProjectionDerived, Detail: "generic rank-seven map has 7*(16-7)=63 kernel/projection degrees and no finite selector"},
			{Name: "dimension-seven sector sums are not contact semantics", Passed: a.CentralPlusLeptoDimension == 7 && a.BLPlusLeptoDimension == 7 && a.DimensionSevenSectorSums == 2 && !a.DimensionSevenSectorSumsCanonical && a.DimensionSevenSectorSumsWrongSemantics, Detail: "1+6 sector sums are current-side matter sectors, not contact-row representation data"},
			{Name: "color/contact/spectral/Fano shortcuts do not derive quotient", Passed: !a.ColorEightToSevenQuotientDerived && !a.ContactEWFourPlusThreeDerived && !a.SpectralSevenIsU4Quotient && !a.FanoSevenIsU4Quotient && a.ObservedProjectionRejected, Detail: FormatCandidates(a.Candidates)},
			{Name: "u(4)-to-contact quotient firewall remains closed", Passed: !a.U4ToContactQuotientDerived && !a.NaturalSevenRowProjectionDerived && a.SevenRowProjectionNoGoDerived && a.RepresentationCompleteRows == 0 && a.RepresentationOpenRows == 7 && a.ContactBetaRowsAllowed == 0 && a.ContactZeroRowsProved == 0 && !a.ThresholdCorrectedBetaDerived && !a.FullBetaMatchingTensorDerived, Detail: FormatRows(a.Rows, 7)},
			{Name: "projection search does not leak physical constants", Passed: a.ResidualNullityBefore == 3 && a.ResidualNullityAfter == 3 && !a.HiddenObservedInputUsed && !a.PhysicalWeakAngleDerived && !a.FineStructureDerived && !a.PhysicalMassesDerived && !a.PhysicalScaleDerived, Detail: "no alpha, physical thetaW, W/Z/Higgs/fermion masses, M*, g_*, or fitted thresholds are used"},
		}, Notes: []string{a.TruthStatement, "criteria: " + FormatCriteria(a.Criteria), "candidates: " + FormatCandidates(a.Candidates), "rejected claims: " + Join(a.RejectedClaims), "remaining unknowns: " + Join(a.RemainingUnknowns), "Next: " + a.RecommendedNextGate}}
	}}
}
