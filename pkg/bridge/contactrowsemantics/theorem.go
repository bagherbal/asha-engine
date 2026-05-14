package contactrowsemantics

import (
	"fmt"

	"github.com/bagherbal/asha-engine/pkg/theorem"
)

func ContactRowSemanticsLocalVariableReconstructionTheorem() theorem.Theorem {
	const id = "BRIDGE-CONTACT-ROW-SEMANTICS-LOCAL-VARIABLE-RECONSTRUCTION"
	const name = "contact row semantics / local variable reconstruction from incidence-weighted spectrum"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.Variational, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "build contact row semantics search", Passed: false, Detail: err.Error()}}}
		}
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.Variational, Checks: []theorem.Check{
			{Name: "Gate 121 reconstruction obstruction inherited", Passed: a.ReconstructionObstructionInherited && a.QuotientForkObstructionInherited && a.NaturalityObstructionInherited && a.SymmetryObstructionInherited && a.Reconstruction.AnonymousInvariantLiftPossibleRows == 5040, Detail: "row-level semantics still cannot be lifted from anonymous invariant data without 7! choices"},
			{Name: "uniform incidence weighting is canonical but semantically inert", Passed: a.IncidenceWeightedSpectrumSearchAttempted && a.UniformFanoIncidenceDegreeAvailable && a.FanoPointDegree == 3 && a.FanoLineSize == 3 && a.IncidenceWeightCanonical && a.IncidenceWeightingPreservesRows && a.IncidenceWeightedValuesDistinct && !a.IncidenceWeightingAddsRowSemantics, Detail: FormatSummary(a.Summary)},
			{Name: "signed incidence still requires a noncanonical contact-Fano labeling", Passed: a.SignedIncidenceAttempted && !a.SignedIncidenceCanonical && a.SignedIncidenceNeedsChoice && a.SignedIncidenceChoiceCount == 5040 && !a.ContactFanoAssignmentDerived, Detail: fmt.Sprintf("signed incidence choices=%d canonical=0", a.SignedIncidenceChoiceCount)},
			{Name: "incidence moments recover spectrum but not row semantics", Passed: a.IncidenceMomentReconstructionAttempted && a.IncidenceMomentsRecoverSpectrum && !a.IncidenceMomentsRecoverRowIdentity && !a.IncidenceMomentsRecoverFanoSemantics && !a.RowSemanticsDerived && a.RowSemanticsObstructionDerived && a.IncidenceWeightedNoGoDerived, Detail: FormatAttempts(a.Attempts)},
			{Name: "local variable, constraint, and representation rows remain unconstructed", Passed: a.LocalVariableReconstructionAttempted && !a.LocalVariablesDerived && a.ConstraintSemanticMapAttempted && !a.ConstraintSemanticMapDerived && a.RepresentationRowRuleAttempted && !a.RepresentationRowRuleDerived && a.ContactRows == 7 && a.OpenContactRowsAfter == 7 && a.RepresentationCompleteRows == 0 && a.RepresentationOpenRows == 7 && a.ContactBetaRowsAllowed == 0 && a.ContactZeroRowsProved == 0 && !a.ThresholdCorrectedBetaDerived && !a.FullFiniteBetaMatchingTensorDerived, Detail: FormatRows(a.Rows, 10)},
			{Name: "contact row semantics search does not leak physical constants", Passed: a.ResidualNullityBefore == 3 && a.ResidualNullityAfter == 3 && !a.ResidualSymmetryBroken && !a.PhysicalWeakAngleDerived && !a.FineStructureDerived && !a.PhysicalMassesDerived && !a.PhysicalScaleDerived && !a.HiddenObservedInputUsed, Detail: "no alpha, physical thetaW, threshold-corrected beta tensor, W/Z/Higgs/fermion masses, M*, or g_* are used"},
		}, Notes: []string{a.TruthStatement, "criteria: " + FormatCriteria(a.Criteria), "attempts: " + FormatAttempts(a.Attempts), "rejected claims: " + Join(a.RejectedClaims), "remaining unknowns: " + Join(a.RemainingUnknowns), "Next: " + a.RecommendedNextGate}}
	}}
}
