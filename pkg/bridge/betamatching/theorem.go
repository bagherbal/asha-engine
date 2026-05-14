package betamatching

import (
	"fmt"
	"github.com/bagherbal/asha-engine/pkg/theorem"
)

func ThresholdRepresentationCompletionBetaMatchingTensorSearchTheorem() theorem.Theorem {
	const id = "BRIDGE-THRESHOLD-REPRESENTATION-COMPLETION-BETA-MATCHING-SEARCH"
	const name = "threshold representation completion and finite beta-matching tensor search"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.Variational, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "build beta-matching tensor audit", Passed: false, Detail: err.Error()}}}
		}
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.Variational, Checks: []theorem.Check{
			{Name: "Gate 107 filtration obstruction inherited", Passed: a.Filtration.ResidualNullityAfter == 3 && a.Filtration.NonUniqueFiltrationWitnessed && !a.Filtration.ThresholdCorrectedBetaDerived, Detail: fmt.Sprintf("nullity=%d; nonunique filtration=%t", a.Filtration.ResidualNullityAfter, a.Filtration.NonUniqueFiltrationWitnessed)},
			{Name: "matching ledger constructed", Passed: len(a.Rows) >= 15 && a.SectorBaselineRows == 1 && a.VacuumRows == 1, Detail: fmt.Sprintf("rows=%d; sector=%d; vacuum=%d; %s", len(a.Rows), a.SectorBaselineRows, a.VacuumRows, FormatRows(a.Rows, 8))},
			{Name: "scalar/contact sector baseline beta row derived", Passed: a.ScalarSectorRowConstructed && a.ScalarSectorMatchesBaseline && !a.ScalarSectorIsThresholdCorrection, Detail: fmt.Sprintf("%s; baseline only, not a heavy threshold correction", FormatVector(a.ScalarSectorDeltaB))},
			{Name: "open modes remain representation-incomplete", Passed: a.IncompleteOpenRows == 8 && !a.BGapRepresentationCompleted && !a.ContactOverlapRepresentationCompleted && !a.AllOpenModesRepresentationComplete, Detail: fmt.Sprintf("incomplete open rows=%d; open=%v", a.IncompleteOpenRows, a.RepresentationOpenModes)},
			{Name: "individual heavy-threshold rows absent", Passed: a.IndividualThresholdRows == 0 && a.BetaCorrectionRowsAllowed == 0, Detail: fmt.Sprintf("individual threshold rows=%d; correction rows=%d", a.IndividualThresholdRows, a.BetaCorrectionRowsAllowed)},
			{Name: "completion attempts audited", Passed: len(a.CompletionAttempts) >= 4 && !a.FullFiniteBetaMatchingTensorDerived, Detail: FormatAttempts(a.CompletionAttempts)},
			{Name: "representation ambiguity witnessed", Passed: len(a.AmbiguityWitnesses) >= 3, Detail: FormatWitnesses(a.AmbiguityWitnesses)},
			{Name: "activation, decoupling, and scale still absent", Passed: !a.ActivationRuleDerived && !a.DecouplingMatchingRuleDerived && !a.PhysicalScaleDerived, Detail: "no finite predicate says which incomplete modes activate, how they decouple, or at what scale they match"},
			{Name: "threshold-corrected beta tensor not derived", Passed: !a.ThresholdCorrectedBetaDerived && !a.FullFiniteBetaMatchingTensorDerived, Detail: "only the baseline scalar sector row is known; no heavy Δb_i tensor is allowed"},
			{Name: "residual nullity unchanged", Passed: a.ResidualNullityAfter == a.ResidualNullityBefore && a.ResidualNullityAfter == 3 && !a.ResidualSymmetryBroken, Detail: fmt.Sprintf("nullity before=%d; after=%d", a.ResidualNullityBefore, a.ResidualNullityAfter)},
			{Name: "physical predictions remain sealed", Passed: !a.PhysicalWeakAngleDerived && !a.FineStructureDerived && !a.PhysicalMassesDerived && !a.HiddenObservedInputUsed, Detail: "no alpha, physical thetaW, W/Z/Higgs/fermion masses, observed scales, or fitted thresholds are used"},
		}, Notes: []string{a.TruthStatement, "rejected claims: " + Join(a.RejectedClaims), "remaining unknowns: " + Join(a.RemainingUnknowns), "Next: " + a.RecommendedNextGate}}
	}}
}
