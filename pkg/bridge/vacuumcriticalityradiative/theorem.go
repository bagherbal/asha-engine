package vacuumcriticalityradiative

import "github.com/bagherbal/asha-engine/pkg/theorem"

func VacuumCriticalityRadiativeHierarchySieveTheorem() theorem.Theorem {
	const id = "BRIDGE-VACUUM-CRITICALITY-RADIATIVE-HIERARCHY-SIEVE"
	const name = "Vacuum Criticality & Radiative Hierarchy Sieve"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "build Gate 350 vacuum criticality/radiative hierarchy audit", Passed: false, Detail: err.Error()}}}
		}
		checks := []theorem.Check{
			{Name: "span inherits Gate 349 without adding a fit", Passed: a.Span.InheritedGate == 349 && !a.Span.AddsNewFit, Detail: FormatSpan(a.Span)},
			{Name: "criticality equation computes a top-Yukawa target but is not native", Passed: a.Criticality.Formalized && a.Criticality.CriticalYukawa > 0 && a.Criticality.RequiresSaturationAxiom && !a.Criticality.ReductionProved, Detail: FormatCriticality(a.Criticality)},
			{Name: "native λ boundary is not a multiple-point tangency", Passed: a.Criticality.NativeLambdaBoundary > 0 && !a.Criticality.NativeBoundaryHasBetaZero, Detail: FormatCriticality(a.Criticality)},
			{Name: "tree-level zero Yukawas remain zero under standard multiplicative RG", Passed: a.Radiative.Formalized && a.Radiative.ZeroYukawaIsFixedPoint && !a.Radiative.GaugeLoopsGenerateYukawas && a.Radiative.RequiresFlavorBreakingOperator && !a.Radiative.ReductionProved, Detail: FormatRadiative(a.Radiative)},
			{Name: "matrix invariant program is identified but not promoted", Passed: a.Invariants.Identified && len(a.Invariants.CandidateInvariants) >= 4 && !a.Invariants.PromotedThisGate, Detail: FormatInvariants(a.Invariants)},
			{Name: "parameter census remains at fifteen minimal vacuum coordinates", Passed: a.Census.StartingVacuumInputs == 15 && a.Census.TotalAdditionalReduction == 0 && a.Census.RemainingVacuumInputs == 15 && !a.Census.SevenSealTargetReached, Detail: FormatCensus(a.Census)},
			{Name: "summary preserves Phase-III vacuum quarantine", Passed: a.Summary.Executed && !a.Summary.AnyReductionProved && a.Summary.RemainingVacuumInputs == 15, Detail: FormatSummary(a.Summary)},
		}
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Checks: checks, Notes: []string{a.Truth, "Gate 350 confirms that criticality and radiative hierarchy are valid dynamical research programs, but neither reduces the vacuum dimension without an additional native saturation or flavor-breaking operator."}}
	}}
}
