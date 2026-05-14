package coarsegrain

import (
	"fmt"

	"github.com/bagherbal/asha-engine/pkg/theorem"
)

func NativeFiniteCoarseGrainingThresholdActivationSearchTheorem() theorem.Theorem {
	const id = "BRIDGE-NATIVE-FINITE-COARSE-GRAINING-THRESHOLD-ACTIVATION-SEARCH"
	const name = "native finite coarse-graining and threshold activation operator search"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.Variational, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "build finite coarse-graining audit", Passed: false, Detail: err.Error()}}}
		}
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.Variational, Checks: []theorem.Check{
			{Name: "Gate 104 boundary seed inherited", Passed: a.FiniteBoundarySeedInherited && a.BoundaryKY > 0, Detail: fmt.Sprintf("k_Y=%.10f; sin²_*=%.10f; trace(K_*)=%.10f", a.BoundaryKY, a.BoundarySin2, a.RelativeHessianTrace)},
			{Name: "threshold inventory inherited", Passed: a.ThresholdInventoryInherited && a.ThresholdDecisionCount > 0, Detail: fmt.Sprintf("decisions=%d; continuum candidates=%d; open=%d; vacuum-only=%d", a.ThresholdDecisionCount, a.ContinuumCandidates, a.OpenThresholdModes, a.VacuumOnlyModes)},
			{Name: "candidate coarse-graining operator inventory", Passed: a.OperatorCount >= 7 && a.ProjectionOperatorsFound && a.SpectralAnchorsFound && a.StaticClassifiersFound, Detail: FormatOperators(a.CandidateOperators)},
			{Name: "native RG semigroup absent", Passed: !a.NativeCoarseGrainingFound && !a.SemigroupLawDerived, Detail: "no finite operator supplies a composable C_s∘C_t law or shell-flow law"},
			{Name: "scale/log parameter absent", Passed: !a.ScaleLogParameterDerived, Detail: "no native finite variable selects L=ln(M*/μ), ε, shell number, or M*"},
			{Name: "fixed point not selected", Passed: !a.FlowFixedPointSelected, Detail: "the boundary Hessian is a seed, not yet a fixed point of a finite flow"},
			{Name: "threshold predicate absent", Passed: !a.ThresholdActivationPredicateDerived && !a.DecouplingMatchingRuleDerived, Detail: "open threshold modes cannot be activated, integrated out, or used in Δb_i"},
			{Name: "absolute coupling running absent", Passed: !a.AbsoluteCouplingRunningDerived, Detail: "relative kinetic normalization remains invariant under the Gate 104 coupling-prefactor rescaling"},
			{Name: "RG requirements audited", Passed: len(a.Requirements) >= 7, Detail: FormatRequirements(a.Requirements)},
			{Name: "non-unique activation schedules witnessed", Passed: a.NonUniqueActivationWitnessed, Detail: FormatSchedules(a.ActivationSchedules)},
			{Name: "residual nullity unchanged", Passed: a.ResidualNullityAfter == a.ResidualNullityBefore && !a.ResidualSymmetryBroken, Detail: fmt.Sprintf("nullity before=%d; after=%d", a.ResidualNullityBefore, a.ResidualNullityAfter)},
			{Name: "physical predictions remain sealed", Passed: !a.PhysicalWeakAngleDerived && !a.FineStructureDerived && !a.PhysicalMassesDerived && !a.HiddenObservedInputUsed, Detail: "no alpha, thetaW, W/Z/Higgs/fermion masses, observed scales, or hidden matching data are used"},
		}, Notes: []string{
			a.TruthStatement,
			"rejected claims: " + Join(a.RejectedClaims),
			"remaining unknowns: " + Join(a.RemainingUnknowns),
			"Next: " + a.RecommendedNextGate,
		}}
	}}
}
