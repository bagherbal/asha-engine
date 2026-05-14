package contactsource

import (
	"fmt"

	"github.com/bagherbal/asha-engine/pkg/theorem"
)

func ContactSemanticSourceCouplingObservableSelectorSearchTheorem() theorem.Theorem {
	const id = "BRIDGE-CONTACT-SEMANTIC-SOURCE-COUPLING-OBSERVABLE-SELECTOR-SEARCH"
	const name = "contact semantic source-coupling / observable selector search"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.Variational, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "build contact semantic source search", Passed: false, Detail: err.Error()}}}
		}
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.Variational, Checks: []theorem.Check{
			{Name: "Gate 122 row-semantics obstruction inherited", Passed: a.Semantics.RowSemanticsObstructionDerived && a.Semantics.IncidenceWeightedNoGoDerived && a.Semantics.SignedIncidenceChoiceCount == 5040 && a.ContactRows == 7 && a.OpenContactRowsAfter == 7, Detail: "contact rows remain open after incidence-weighted row-semantics search"},
			{Name: "uniform action/source coupling is canonical but row-blind", Passed: a.SemanticSourceCouplingSearchAttempted && a.UniformActionSourceAvailable && a.UniformActionSourceCanonical && a.UniformActionSourceRowBlind && !a.UniformActionSourceSelectsRows, Detail: FormatSummary(a.Summary)},
			{Name: "spectral observable distinguishes rows only diagnostically", Passed: a.SpectralObservableAttempted && a.SpectralObservableConstructed && a.SpectralObservableCanonical && a.SpectralObservableRowsDistinguished == 7 && !a.SpectralObservableAddsSemantics && a.SpectralObservableOnlyDiagnostic, Detail: FormatRows(a.Rows, 7)},
			{Name: "current/action source selectors remain unselected", Passed: a.CurrentSourceAttempted && a.CurrentSourceObstructionInherited && !a.CurrentToContactMapDerived && a.CurrentSourceRowsDerived == 0 && a.ActionCouplingSelectorAttempted && !a.ActionCouplingSelectorDerived && !a.SemanticSourceSelectorDerived && a.ContactSourceSelectorNoGoDerived, Detail: fmt.Sprintf("current map derived=%t source rows=%d", a.CurrentToContactMapDerived, a.CurrentSourceRowsDerived)},
			{Name: "source/observable selector does not open contact beta matching", Passed: a.RepresentationCompleteRows == 0 && a.RepresentationOpenRows == 7 && a.ContactBetaRowsAllowed == 0 && a.ContactZeroRowsProved == 0 && !a.ThresholdCorrectedBetaDerived && !a.FullFiniteBetaMatchingTensorDerived, Detail: FormatAttempts(a.Attempts)},
			{Name: "contact source search does not leak physical constants", Passed: a.ResidualNullityBefore == 3 && a.ResidualNullityAfter == 3 && !a.ResidualSymmetryBroken && !a.HiddenObservedInputUsed && !a.PhysicalWeakAngleDerived && !a.FineStructureDerived && !a.PhysicalMassesDerived && !a.PhysicalScaleDerived, Detail: "no alpha, physical thetaW, threshold-corrected beta tensor, W/Z/Higgs/fermion masses, M*, or g_* are used"},
		}, Notes: []string{a.TruthStatement, "criteria: " + FormatCriteria(a.Criteria), "attempts: " + FormatAttempts(a.Attempts), "rejected claims: " + Join(a.RejectedClaims), "remaining unknowns: " + Join(a.RemainingUnknowns), "Next: " + a.RecommendedNextGate}}
	}}
}
