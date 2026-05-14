package exchangeselection

import (
	"fmt"

	"github.com/bagherbal/asha-engine/pkg/theorem"
)

func FiniteExchangeActionSelectionTheorem() theorem.Theorem {
	const id = "BRIDGE-FINITE-EXCHANGE-ACTION-SELECTION"
	const name = "finite exchange-action selection principle"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "build exchange-action selection audit", Passed: false, Detail: err.Error()}}}
		}
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Checks: []theorem.Check{
			{Name: "Gate 67 Casimir diagnostics inherited", Passed: a.Previous.AllCasimirDiagnosticsBuilt, Detail: "direct, inverse-nonzero, and trace-normalized Casimir diagnostics are available"},
			{Name: "candidate exchange rules exposed", Passed: a.CandidateRuleCount == 4, Detail: FormatRules(a.CandidateRules)},
			{Name: "candidate positivity", Passed: a.AllRulesPositive, Detail: "all candidate kernel rules are positive diagnostics on their supported sectors"},
			{Name: "finite-data discipline", Passed: a.AllRulesUseOnlyFiniteData && !a.HiddenObservedInputUsed, Detail: "no observed coupling, threshold, y_t, Higgs scale, or fermion mass enters the selection audit"},
			{Name: "direct/inverse ambiguity", Passed: a.DirectInverseDisagree, Detail: fmt.Sprintf("direct dominant=%s, inverse dominant=%s; disagreement prevents canonical selection", a.Previous.DominantDirectSector, a.Previous.DominantInverseSector)},
			{Name: "minimal action form exposed", Passed: a.MinimalActionFormExposed, Detail: "S[J]=1/2 <J,KJ> - <J,source>; the missing object is K from a finite action Hessian"},
			{Name: "finite action selection", Passed: a.AnyRuleSelectedByAction, Detail: "open; no finite action selects direct Casimir, inverse Casimir, trace-normalized Casimir, or unit-sector propagation"},
			{Name: "current kinetic operator", Passed: a.KineticOperatorDerived, Detail: "open; representation Casimirs are not yet the current-field Hessian"},
			{Name: "second variation", Passed: a.SecondVariationDerived, Detail: "open; BF/projector/contact action has not been varied with respect to current-sector fields"},
			{Name: "propagator rule", Passed: a.PropagatorRuleDerived, Detail: "open; rho_A denominators and exchange kernels remain unselected"},
			{Name: "exchange kernel update", Passed: a.ExchangeKernelUpdated, Detail: "open; G_hat cannot be updated from a diagnostic family without action selection"},
			{Name: "attractive scalar-channel theorem", Passed: a.AttractiveScalarChannelDerived, Detail: "open; no NJL attraction or condensation follows from this gate"},
			{Name: "up/down splitting", Passed: a.UpDownSplittingDerived, Detail: "open; all candidate rules still act on lepton/color flavor and do not select top over bottom"},
			{Name: "condensation claim", Passed: a.CondensationClaimAllowed, Detail: "false by design; no top condensation, Higgs VEV, or fermion mass is claimed"},
		}, Notes: []string{
			a.SelectionObstruction,
			a.TruthStatement,
			fmt.Sprintf("recommended next gate: %s", a.RecommendedNextGate),
			fmt.Sprintf("remaining unknowns: %s", FormatUnknowns(a.RemainingUnknowns)),
		}}
	}}
}
