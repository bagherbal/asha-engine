package shellfunctor

import (
	"fmt"

	"github.com/bagherbal/asha-engine/pkg/theorem"
)

func FiniteShellFunctorSemigroupConstructionTheorem() theorem.Theorem {
	const id = "BRIDGE-FINITE-SHELL-FUNCTOR-SEMIGROUP-CONSTRUCTION"
	const name = "finite shell functor and semigroup construction attempt"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.Variational, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "build finite shell functor audit", Passed: false, Detail: err.Error()}}}
		}
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.Variational, Checks: []theorem.Check{
			{Name: "Gate 105 obstruction inherited", Passed: a.Coarse.ResidualNullityAfter == 3 && !a.Coarse.NativeCoarseGrainingFound, Detail: fmt.Sprintf("nullity=%d; native coarse-graining found=%t", a.Coarse.ResidualNullityAfter, a.Coarse.NativeCoarseGrainingFound)},
			{Name: "finite shell carrier constructed", Passed: a.ModeCount > 0 && a.ContinuumCount > 0 && a.OpenCount > 0, Detail: fmt.Sprintf("modes=%d; continuum=%d; open=%d; vacuum-only=%d; %s", a.ModeCount, a.ContinuumCount, a.OpenCount, a.VacuumCount, FormatModes(a.Modes))},
			{Name: "nested projection family constructed", Passed: a.NestedProjectionFamilyConstructed && a.ProjectionCount >= a.OpenCount+1, Detail: FormatProjections(a.Projections)},
			{Name: "composition closed as projection semilattice", Passed: a.CompositionClosed && a.AssociativityVerified && a.IdentityProjectionExists, Detail: FormatCompositionWitnesses(a.CompositionTable)},
			{Name: "idempotent semilattice not additive RG", Passed: a.IdempotentSemilatticeDerived && !a.AdditiveSemigroupDerived && a.NontrivialAdditiveCounterexample, Detail: "nested projections satisfy C_a∘C_b=C_min(a,b), while additive/log RG would require C_s∘C_t=C_{s+t}"},
			{Name: "functor attempts audited", Passed: len(a.FunctorAttempts) >= 4 && !a.NativeFiniteRGFunctorDerived, Detail: FormatFunctorAttempts(a.FunctorAttempts)},
			{Name: "shell ordering non-unique", Passed: !a.CanonicalShellOrderingDerived && len(a.ScheduleWitnesses) >= 2, Detail: FormatSchedules(a.ScheduleWitnesses)},
			{Name: "scale and activation still absent", Passed: !a.CanonicalScaleLogParameterDerived && !a.ThresholdActivationPredicateDerived && !a.DecouplingMatchingRuleDerived, Detail: "no finite variable selects L, no threshold predicate, no Δb_i matching map"},
			{Name: "absolute coupling flow still absent", Passed: !a.AbsoluteCouplingFlowDerived, Detail: "finite shell projections do not renormalize the overall gauge-action prefactor"},
			{Name: "residual nullity unchanged", Passed: a.ResidualNullityAfter == a.ResidualNullityBefore && !a.ResidualSymmetryBroken, Detail: fmt.Sprintf("nullity before=%d; after=%d", a.ResidualNullityBefore, a.ResidualNullityAfter)},
			{Name: "physical predictions remain sealed", Passed: !a.PhysicalWeakAngleDerived && !a.FineStructureDerived && !a.PhysicalMassesDerived && !a.HiddenObservedInputUsed, Detail: "no alpha, physical thetaW, W/Z/Higgs/fermion masses, observed scales, or fitted thresholds are used"},
		}, Notes: []string{
			a.TruthStatement,
			"rejected claims: " + Join(a.RejectedClaims),
			"remaining unknowns: " + Join(a.RemainingUnknowns),
			"Next: " + a.RecommendedNextGate,
		}}
	}}
}
