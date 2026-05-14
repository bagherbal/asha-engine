package contactdualcurrenttarget

import (
	"fmt"

	"github.com/bagherbal/asha-engine/pkg/theorem"
)

func ContactDualCurrentTargetEnlargementSevenRowCarrierSearchTheorem() theorem.Theorem {
	const id = "BRIDGE-CONTACT-DUAL-CURRENT-TARGET-ENLARGEMENT-SEVEN-ROW-CARRIER-SEARCH"
	const name = "contact dual-current target enlargement / seven-row carrier search"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.Variational, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "build contact dual-current target enlargement search", Passed: false, Detail: err.Error()}}}
		}
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.Variational, Checks: []theorem.Check{
			{Name: "Gate 124 dual-pairing no-go inherited", Passed: a.Previous.ContactDualPairingNoGoDerived && a.ContactRows == 7 && a.OpenContactRowsAfter == 7 && a.Previous.ContactBetaRowsAllowed == 0 && a.Previous.ContactZeroRowsProved == 0, Detail: FormatSummary(a.Summary)},
			{Name: "existing derived current targets are not seven-row contact targets", Passed: a.UniformTargetDimension == 1 && a.UniformTargetRowBlind && a.ContactEWTargetDimension == 4 && a.ContactEWTargetDerived && !a.ContactEWTargetSevenRows && a.PatiSalamTargetDimension == 16 && a.PatiSalamTargetDerived && !a.PatiSalamTargetSevenRows && a.LeptoquarkTargetDimension == 6 && !a.LeptoquarkTargetSevenRows, Detail: fmt.Sprintf("dims: uniform=%d contactEW=%d u4=%d lepto=%d", a.UniformTargetDimension, a.ContactEWTargetDimension, a.PatiSalamTargetDimension, a.LeptoquarkTargetDimension)},
			{Name: "spectral seven-row carrier is diagnostic not current-derived", Passed: a.SpectralSevenTargetConstructed && a.SpectralSevenTargetCanonical && a.SpectralSevenTargetRowsDistinguished == 7 && !a.SpectralSevenTargetCurrentDerived && !a.SpectralSevenTargetSemantic, Detail: FormatRows(a.Rows, 7)},
			{Name: "Fano seven-row carrier still requires hidden assignment", Passed: a.FanoSevenTargetConstructed && !a.FanoSevenTargetCanonical && a.FanoSevenTargetRequiresChoice && a.FanoSevenTargetHiddenChoices == 5040 && !a.FanoSevenTargetCurrentDerived, Detail: "Fano R^7 is seven-dimensional, but contact row semantics require one of 7! labelings"},
			{Name: "anonymous seven-row carrier preserves cardinality but not semantics", Passed: a.AnonymousSevenTargetConstructed && a.AnonymousSevenTargetCanonical && !a.AnonymousSevenTargetRowSemantic, Detail: "anonymous R^7 has seven slots but no local/current/gauge semantics"},
			{Name: "no seven-row dual-current target opens beta matching", Passed: !a.DualCurrentTargetDerived && a.SevenRowTargetNoGoDerived && !a.NaturalSevenRowLabelsDerived && a.RepresentationCompleteRows == 0 && a.RepresentationOpenRows == 7 && a.ContactBetaRowsAllowed == 0 && a.ContactZeroRowsProved == 0 && !a.ThresholdCorrectedBetaDerived && !a.FullBetaMatchingTensorDerived, Detail: FormatCandidates(a.Candidates)},
			{Name: "target enlargement does not leak physical constants", Passed: a.ResidualNullityBefore == 3 && a.ResidualNullityAfter == 3 && !a.ResidualSymmetryBroken && !a.HiddenObservedInputUsed && !a.PhysicalWeakAngleDerived && !a.FineStructureDerived && !a.PhysicalMassesDerived && !a.PhysicalScaleDerived, Detail: "no alpha, physical thetaW, threshold-corrected beta tensor, W/Z/Higgs/fermion masses, M*, or g_* are used"},
		}, Notes: []string{a.TruthStatement, "criteria: " + FormatCriteria(a.Criteria), "candidates: " + FormatCandidates(a.Candidates), "rejected claims: " + Join(a.RejectedClaims), "remaining unknowns: " + Join(a.RemainingUnknowns), "Next: " + a.RecommendedNextGate}}
	}}
}
