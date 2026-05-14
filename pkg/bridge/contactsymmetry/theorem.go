package contactsymmetry

import (
	"fmt"

	"github.com/bagherbal/asha-engine/pkg/theorem"
)

func ContactSymmetryBreakingSelectorStabilizerSearchTheorem() theorem.Theorem {
	const id = "BRIDGE-CONTACT-SYMMETRY-BREAKING-STABILIZER-SELECTOR-SEARCH"
	const name = "contact symmetry-breaking selector / stabilizer reduction search"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.Variational, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "build contact symmetry-breaking selector search", Passed: false, Detail: err.Error()}}}
		}
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.Variational, Checks: []theorem.Check{
			{Name: "Gate 117 naturality obstruction inherited", Passed: a.NaturalityObstructionInherited && a.FanoAutomorphismGroupDerived && a.FanoAutomorphismGroupOrder == 168 && a.FanoPointActionTransitive && a.FanoLineActionTransitive, Detail: "the full transitive Fano automorphism audit is the input symmetry"},
			{Name: "stabilizer arithmetic is exact but conditional", Passed: a.StabilizerArithmeticDerived && a.PointStabilizerOrder == 24 && a.LineStabilizerOrder == 24 && a.IncidentFlagStabilizerOrder == 8 && a.StabilizerReductionPossibleAfterChoice, Detail: FormatStabilizerSummary(a.StabilizerSummary)},
			{Name: "no canonical finite symmetry-breaking object is selected", Passed: !a.CanonicalSymmetryBreakingObjectDerived && !a.CanonicalFanoPointSelected && !a.CanonicalFanoLineSelected && !a.CanonicalFanoFlagSelected && !a.CanonicalContactFanoAssignmentDerived, Detail: fmt.Sprintf("fixed points=%d; fixed lines=%d; stabilizer reduction requires a chosen point/line/flag", a.GlobalFixedFanoPoints, a.GlobalFixedFanoLines)},
			{Name: "spectral and signed-incidence diagnostics do not break Fano symmetry canonically", Passed: a.SpectralOrderingAvailable && !a.SpectralOrderingCanonicalForFano && !a.SignedFanoOrientationBreaksSymmetry && !a.ContactAutomorphismActionDerived && !a.NaturalitySquareFormulable, Detail: "spectral labels and signed incidence are finite diagnostics, not a natural contact-side Aut(Fano) action"},
			{Name: "contact representation and beta permission remain closed", Passed: a.ContactRows == 7 && a.OpenContactRowsAfter == 7 && a.RepresentationCompleteRows == 0 && a.RepresentationOpenRows == 7 && a.ContactBetaRowsAllowed == 0 && a.ContactZeroRowsProved == 0 && !a.ThresholdCorrectedBetaDerived && !a.FullFiniteBetaMatchingTensorDerived, Detail: FormatRows(a.Rows, 10)},
			{Name: "selector search does not leak physical constants", Passed: a.ResidualNullityBefore == 3 && a.ResidualNullityAfter == 3 && !a.ResidualSymmetryBroken && !a.PhysicalWeakAngleDerived && !a.FineStructureDerived && !a.PhysicalMassesDerived && !a.PhysicalScaleDerived && !a.HiddenObservedInputUsed, Detail: "no alpha, physical thetaW, W/Z/Higgs/fermion masses, M*, g_*, or observed thresholds are used"},
		}, Notes: []string{a.TruthStatement, "criteria: " + FormatCriteria(a.Criteria), "stabilizers: " + FormatStabilizers(a.Stabilizers), "rejected claims: " + Join(a.RejectedClaims), "remaining unknowns: " + Join(a.RemainingUnknowns), "Next: " + a.RecommendedNextGate}}
	}}
}
