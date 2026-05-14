package chiraltrace

import (
	"fmt"

	"github.com/bagherbal/asha-engine/pkg/theorem"
)

func FiniteChiralBilinearMetricTheorem() theorem.Theorem {
	const id = "BRIDGE-FINITE-CHIRAL-BILINEAR-METRIC"
	const name = "finite chiral bilinear metric / Fock trace construction"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "construct finite chiral bilinear metric", Passed: false, Detail: err.Error()}}}
		}
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Checks: []theorem.Check{
			{Name: "Yukawa incidence input", Passed: a.Loop.FiniteYukawaIncidenceOperatorDerived, Detail: fmt.Sprintf("Y maps %d domain states to %d right states with %d allowed fiber entries", a.DomainDimension, a.RightDimension, a.AllowedFiberEntries)},
			{Name: "normalized chiral bilinear map", Passed: a.BilinearMetricConstructed && a.RightMetricResidual < 1e-8, Detail: fmt.Sprintf("U=Y^T/sqrt(%.10f); ||U^TU-I_R||_F=%.3e", a.CommonRightRowNormSquared, a.RightMetricResidual)},
			{Name: "scalar LR projector", Passed: a.ScalarLRProjectorConstructed, Detail: fmt.Sprintf("rank=%d, Tr(P_LR)=%.10f, complement dim=%d", a.ProjectorRank, a.ProjectorTrace, a.DomainComplementDimension)},
			{Name: "projector identities", Passed: a.ProjectorIdemResidual < 1e-8 && a.ProjectorSymResidual < 1e-8, Detail: fmt.Sprintf("||P_LR²−P_LR||_F=%.3e, ||P_LR−P_LRᵀ||_F=%.3e", a.ProjectorIdemResidual, a.ProjectorSymResidual)},
			{Name: "finite Fock/Yukawa trace rules", Passed: a.FiniteFockTraceRulesConstructed, Detail: "ordinary finite trace metric for the scalar LR incidence target is constructed"},
			{Name: "full Clifford/Lorentz trace rules", Passed: a.FullCliffordTraceRulesDerived, Detail: "open; gamma/pseudoscalar trace identities and Lorentz scalar contraction are not yet implemented"},
			{Name: "current scalar-projection coefficients", Passed: a.CurrentScalarProjectionCoefficientsKnown, Detail: "open; c_A requires current-generator action on P_LR"},
			{Name: "generator normalization", Passed: a.GeneratorNormalizationDerived, Detail: "open; x∧p/u(4) kinetic trace normalization is still missing"},
			{Name: "attractive scalar-channel sign", Passed: a.AttractiveSignDerived, Detail: "open; finite action/propagator sign is not derived"},
			{Name: "up/down splitting", Passed: a.UpDownTieResolved, Detail: "open; scalar LR projector still treats up/down incidence symmetrically"},
			{Name: "hidden observed input", Passed: !a.HiddenObservedInputUsed, Detail: "no observed Yukawa, v, Higgs mass, or fitted coupling was inserted"},
		}, Notes: []string{
			a.TruthStatement,
			fmt.Sprintf("trace identities: %s", FormatTraceIdentities(a.TraceIdentities)),
			fmt.Sprintf("requirements: %s", FormatRequirements(a.Requirements)),
			fmt.Sprintf("recommended next gate: %s", a.RecommendedNextGate),
			fmt.Sprintf("remaining unknowns: %s", FormatUnknowns(a.RemainingUnknowns)),
		}}
	}}
}
