package contactdualpairing

import (
	"fmt"

	"github.com/bagherbal/asha-engine/pkg/theorem"
)

func ContactSourceCurrentDualPairingNaturalityObstructionTheorem() theorem.Theorem {
	const id = "BRIDGE-CONTACT-SOURCE-CURRENT-DUAL-PAIRING-NATURALITY-OBSTRUCTION"
	const name = "contact source-current dual pairing / row-label naturality obstruction theorem"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.Variational, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "build contact source-current dual pairing search", Passed: false, Detail: err.Error()}}}
		}
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.Variational, Checks: []theorem.Check{
			{Name: "Gate 123 source-selector no-go inherited", Passed: a.Source.ContactSourceSelectorNoGoDerived && a.ContactRows == 7 && a.OpenContactRowsAfter == 7 && a.Source.ContactBetaRowsAllowed == 0 && a.Source.ContactZeroRowsProved == 0, Detail: FormatSummary(a.Summary)},
			{Name: "uniform source-current pairing is canonical but row-blind", Passed: a.UniformPairingAttempted && a.UniformPairingConstructed && a.UniformPairingCanonical && a.UniformPairingRank == 1 && a.UniformPairingRowsDistinguished == 0 && a.UniformPairingRowBlind, Detail: "canonical rank-one unit pairing cannot label seven contact rows"},
			{Name: "spectral diagonal pairing is nondegenerate but diagnostic only", Passed: a.SpectralPairingAttempted && a.SpectralPairingConstructed && a.SpectralPairingCanonical && a.SpectralPairingNonDegenerate && a.SpectralPairingRowsDistinguished == 7 && !a.SpectralPairingAddsSemantics && a.SpectralPairingDiagnosticOnly, Detail: FormatRows(a.Rows, 7)},
			{Name: "current-to-contact dual functional remains unselected", Passed: a.CurrentDualPairingAttempted && a.CurrentDualObstructionInherited && !a.CurrentToContactMapDerived && !a.CurrentFunctionalDerived && !a.SourceFunctionalDerived && a.CurrentDualRowsDerived == 0 && !a.CurrentDualPairingDerived, Detail: fmt.Sprintf("currentMap=%t currentFunctional=%t sourceFunctional=%t", a.CurrentToContactMapDerived, a.CurrentFunctionalDerived, a.SourceFunctionalDerived)},
			{Name: "Fano-labelled dual pairing requires hidden choice", Passed: a.FanoLabelledPairingAttempted && !a.FanoLabelledPairingDerived && a.RequiresHiddenFanoChoice && a.HiddenFanoChoices == 5040 && !a.NaturalRowLabelDerived && a.ContactDualPairingNoGoDerived, Detail: FormatAttempts(a.Attempts)},
			{Name: "dual pairing search does not open beta matching", Passed: a.RepresentationCompleteRows == 0 && a.RepresentationOpenRows == 7 && a.ContactBetaRowsAllowed == 0 && a.ContactZeroRowsProved == 0 && !a.ThresholdCorrectedBetaDerived && !a.FullBetaMatchingTensorDerived, Detail: "contact beta firewall remains closed"},
			{Name: "dual pairing search does not leak physical constants", Passed: a.ResidualNullityBefore == 3 && a.ResidualNullityAfter == 3 && !a.ResidualSymmetryBroken && !a.HiddenObservedInputUsed && !a.PhysicalWeakAngleDerived && !a.FineStructureDerived && !a.PhysicalMassesDerived && !a.PhysicalScaleDerived, Detail: "no alpha, physical thetaW, threshold-corrected beta tensor, W/Z/Higgs/fermion masses, M*, or g_* are used"},
		}, Notes: []string{a.TruthStatement, "criteria: " + FormatCriteria(a.Criteria), "attempts: " + FormatAttempts(a.Attempts), "rejected claims: " + Join(a.RejectedClaims), "remaining unknowns: " + Join(a.RemainingUnknowns), "Next: " + a.RecommendedNextGate}}
	}}
}
