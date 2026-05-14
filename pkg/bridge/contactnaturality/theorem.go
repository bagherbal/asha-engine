package contactnaturality

import (
	"fmt"

	"github.com/bagherbal/asha-engine/pkg/theorem"
)

func ContactFanoNaturalityAutomorphismObstructionTheorem() theorem.Theorem {
	const id = "BRIDGE-CONTACT-FANO-NATURALITY-AUTOMORPHISM-OBSTRUCTION"
	const name = "contact-Fano naturality obstruction / automorphism-invariance theorem"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.Variational, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "build contact-Fano naturality obstruction", Passed: false, Detail: err.Error()}}}
		}
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.Variational, Checks: []theorem.Check{
			{Name: "Gate 116 incidence obstruction inherited", Passed: a.IncidenceFunctorObstructionInherited && a.ContactIncidence.IncidenceFunctorObstructionDerived && a.ContactRows == 7 && a.RepresentationOpenRows == 7 && a.ContactBetaRowsAllowed == 0, Detail: "seven contact rows remain incidence-open before automorphism/naturality audit"},
			{Name: "Fano automorphism group derived from incidence", Passed: a.FanoAutomorphismGroupDerived && a.FanoAutomorphismGroupOrder == 168 && a.AutomorphismSummary.IdentityCount == 1 && a.AutomorphismSummary.NonIdentityCount == 167, Detail: FormatAutomorphismSummary(a.AutomorphismSummary)},
			{Name: "Fano action is transitive, not selector-like", Passed: a.FanoPointActionTransitive && a.FanoLineActionTransitive && a.GlobalFixedFanoPoints == 0 && a.GlobalFixedFanoLines == 0 && !a.AutomorphismInvariantPointSelector && !a.AutomorphismInvariantLineSelector, Detail: fmt.Sprintf("point orbits=%v; line orbits=%v; fixed points=%d; fixed lines=%d", a.AutomorphismSummary.PointOrbitSizes, a.AutomorphismSummary.LineOrbitSizes, a.GlobalFixedFanoPoints, a.GlobalFixedFanoLines)},
			{Name: "contact-side action and naturality square not derived", Passed: !a.ContactAutomorphismActionDerived && !a.NaturalitySquareFormulable && !a.InvariantContactToFanoMapDerived && !a.EquivariantBijectionDerived, Detail: "no derived action of Aut(Fano) on contact-overlap rows; equivariance cannot be typed without adding convention"},
			{Name: "all contact-Fano bijections remain convention-dependent", Passed: a.CanonicalAssignmentCount == 0 && a.CompatibleBijectionCount == 5040 && a.ConventionDependentBijections == 5040 && a.SpectralOrderingBreaksAutomorphism, Detail: "7! compatible assignments remain; spectral labels break symmetry by convention rather than finite naturality"},
			{Name: "representation and threshold permission remain closed", Passed: a.RepresentationCompleteRows == 0 && a.RepresentationOpenRows == 7 && a.ContactBetaRowsAllowed == 0 && a.ContactZeroRowsProved == 0 && !a.ThresholdCorrectedBetaDerived && !a.FullFiniteBetaMatchingTensorDerived, Detail: FormatRows(a.Rows, 10)},
			{Name: "naturality audit does not leak physical constants", Passed: a.ResidualNullityBefore == 3 && a.ResidualNullityAfter == 3 && !a.ResidualSymmetryBroken && !a.PhysicalWeakAngleDerived && !a.FineStructureDerived && !a.PhysicalMassesDerived && !a.PhysicalScaleDerived && !a.HiddenObservedInputUsed, Detail: "no alpha, physical thetaW, W/Z/Higgs/fermion masses, M*, g_*, or observed thresholds are used"},
		}, Notes: []string{a.TruthStatement, "criteria: " + FormatCriteria(a.Criteria), "attempts: " + FormatAttempts(a.Attempts), "rejected claims: " + Join(a.RejectedClaims), "remaining unknowns: " + Join(a.RemainingUnknowns), "Next: " + a.RecommendedNextGate}}
	}}
}
