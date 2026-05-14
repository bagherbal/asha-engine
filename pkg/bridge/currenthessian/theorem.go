package currenthessian

import (
	"fmt"

	"github.com/bagherbal/asha-engine/pkg/theorem"
)

func FiniteCurrentHessianSearchTheorem() theorem.Theorem {
	const id = "BRIDGE-FINITE-CURRENT-HESSIAN-SEARCH"
	const name = "finite current Hessian / action second-variation search"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "build current Hessian audit", Passed: false, Detail: err.Error()}}}
		}
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Checks: []theorem.Check{
			{Name: "Gate 68 exchange-selection input", Passed: !a.Previous.AnyRuleSelectedByAction, Detail: "previous gate exposed candidate exchange rules but selected none"},
			{Name: "current-sector Hessian domain", Passed: a.SectorCount == 4 && a.GeneratorCount == 16, Detail: fmt.Sprintf("sector fields=%d, generator fields=%d", a.SectorCount, a.GeneratorCount)},
			{Name: "Hessian search-space dimensions", Passed: a.SectorHessianDimension == 10 && a.GeneratorHessianDimension == 136, Detail: fmt.Sprintf("symmetric sector Hessian dim=%d, symmetric generator Hessian dim=%d, diagonal sector dim=%d", a.SectorHessianDimension, a.GeneratorHessianDimension, a.DiagonalSectorDimension)},
			{Name: "candidate Hessian diagnostics", Passed: a.CandidateCount == 4 && a.AllCandidatesPositive, Detail: FormatCandidates(a.CandidateHessians)},
			{Name: "finite-data discipline", Passed: a.AllCandidatesFiniteDataOnly && !a.HiddenObservedInputUsed, Detail: "candidate Hessians use finite sector traces/weights only; no observed couplings or masses inserted"},
			{Name: "direct/inverse ambiguity persists", Passed: a.DirectInverseAmbiguity, Detail: "positive Hessian diagnostics still disagree on sector dominance; action selection is required"},
			{Name: "minimal action template", Passed: a.MinimalActionTemplate != "", Detail: a.MinimalActionTemplate},
			{Name: "current action variables", Passed: a.ActionVariablesDefined, Detail: "open; current-sector exchange fields are not yet typed inside the finite BF/projector/contact action"},
			{Name: "second variation", Passed: a.SecondVariationComputed, Detail: "open; no δ²S/δj_Aδj_B has been computed"},
			{Name: "current Hessian derived", Passed: a.CurrentHessianDerived, Detail: "open; K remains an unknown finite kinetic operator"},
			{Name: "propagator rule", Passed: a.PropagatorRuleDerived, Detail: "open; no diagnostic Hessian is promoted to propagator"},
			{Name: "exchange kernel update", Passed: a.ExchangeKernelUpdated, Detail: "open; G_hat cannot be updated before K is selected"},
			{Name: "attractive scalar-channel theorem", Passed: a.AttractiveScalarDerived, Detail: "open; no NJL attraction or condensation follows from this gate"},
			{Name: "up/down splitting", Passed: a.UpDownSplittingDerived, Detail: "open; the current Hessian still does not select top-like up over bottom-like down"},
			{Name: "condensation claim", Passed: a.CondensationClaimAllowed, Detail: "false by design; no Higgs VEV, top condensation, or fermion mass is claimed"},
		}, Notes: []string{
			a.HessianObstruction,
			a.TruthStatement,
			fmt.Sprintf("recommended next gate: %s", a.RecommendedNextGate),
			fmt.Sprintf("remaining unknowns: %s", FormatUnknowns(a.RemainingUnknowns)),
		}}
	}}
}
