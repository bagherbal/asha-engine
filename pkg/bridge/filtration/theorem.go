package filtration

import (
	"fmt"

	"github.com/bagherbal/asha-engine/pkg/theorem"
)

func FiniteFiltrationOrderSelectorThresholdPredicateSearchTheorem() theorem.Theorem {
	const id = "BRIDGE-FINITE-FILTRATION-ORDER-SELECTOR-THRESHOLD-PREDICATE-SEARCH"
	const name = "finite filtration order selector and monotone threshold predicate search"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.Variational, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "build finite filtration/order audit", Passed: false, Detail: err.Error()}}}
		}
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.Variational, Checks: []theorem.Check{
			{Name: "Gate 106 shell obstruction inherited", Passed: a.Shell.ResidualNullityAfter == 3 && a.Shell.IdempotentSemilatticeDerived && !a.Shell.AdditiveSemigroupDerived, Detail: fmt.Sprintf("nullity=%d; idempotent=%t; additive=%t", a.Shell.ResidualNullityAfter, a.Shell.IdempotentSemilatticeDerived, a.Shell.AdditiveSemigroupDerived)},
			{Name: "finite filtration carrier constructed", Passed: a.ModeCount > 0 && a.ContinuumCount == 5 && a.OpenCount == 8 && a.VacuumCount == 1, Detail: fmt.Sprintf("modes=%d; continuum=%d; open=%d; vacuum=%d; %s", a.ModeCount, a.ContinuumCount, a.OpenCount, a.VacuumCount, FormatModes(a.Modes))},
			{Name: "status preorder exposes threshold antichain", Passed: a.StatusPreorderConstructed && len(a.Antichains) > 0 && !a.CanonicalTotalOrderDerived, Detail: FormatAntichains(a.Antichains)},
			{Name: "value filtrations constructed but non-unique", Passed: a.SpectralValueOrdersConstructed && a.ReverseOrderEquallyCompatible && a.NonUniqueFiltrationWitnessed, Detail: fmt.Sprintf("ascending-first=%q; descending-first=%q; %s", a.FirstAscendingOpen, a.FirstDescendingOpen, FormatOrders(a.OrderingWitnesses))},
			{Name: "no canonical orientation or cutoff selected", Passed: !a.CanonicalOrientationDerived && !a.CanonicalCutoffDerived, Detail: "ascending, descending, kind-first, and status-shell orders remain compatible with finite data"},
			{Name: "monotone predicate family audited", Passed: a.MonotonePredicateFamilyConstructed && !a.DerivedActivationPredicate, Detail: FormatPredicates(a.PredicateWitnesses)},
			{Name: "selector attempts rejected as physical thresholds", Passed: len(a.SelectorAttempts) >= 4 && !a.ThresholdCorrectedBetaDerived, Detail: FormatAttempts(a.SelectorAttempts)},
			{Name: "no decoupling or beta matching derived", Passed: !a.DerivedDecouplingMatchingRule && !a.ThresholdCorrectedBetaDerived, Detail: "no activated shell has representation-complete Δb_i matching or physical decoupling scale"},
			{Name: "native RG functor still absent", Passed: !a.NativeFiniteRGFunctorDerived, Detail: "filtration creates order predicates, not non-idempotent coupling flow"},
			{Name: "residual nullity unchanged", Passed: a.ResidualNullityAfter == a.ResidualNullityBefore && a.ResidualNullityAfter == 3 && !a.ResidualSymmetryBroken, Detail: fmt.Sprintf("nullity before=%d; after=%d", a.ResidualNullityBefore, a.ResidualNullityAfter)},
			{Name: "physical predictions remain sealed", Passed: !a.PhysicalWeakAngleDerived && !a.FineStructureDerived && !a.PhysicalMassesDerived && !a.HiddenObservedInputUsed, Detail: "no alpha, physical thetaW, W/Z/Higgs/fermion masses, observed scales, or fitted thresholds are used"},
		}, Notes: []string{
			a.TruthStatement,
			"invariant safe predicate: " + a.InvariantSafePredicate,
			"rejected claims: " + Join(a.RejectedClaims),
			"remaining unknowns: " + Join(a.RemainingUnknowns),
			"Next: " + a.RecommendedNextGate,
		}}
	}}
}
