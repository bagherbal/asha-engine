package exchangeaction

import (
	"fmt"

	"github.com/bagherbal/asha-engine/pkg/theorem"
)

func FiniteExchangeActionPropagatorSearchTheorem() theorem.Theorem {
	const id = "BRIDGE-FINITE-EXCHANGE-ACTION-PROPAGATOR-SEARCH"
	const name = "Finite exchange action / propagator normalization search"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "construct exchange-action audit", Passed: false, Detail: err.Error()}}}
		}
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Checks: []theorem.Check{
			{Name: "Gate 62 exchange-kernel input", Passed: a.Exchange.ConditionalAttractiveBranchAvailable, Detail: fmt.Sprintf("conditional unit attractive kernel=%.10f", a.Exchange.ConditionalAttractiveUnitKernel)},
			{Name: "current-sector kinetic data", Passed: a.CurrentSectorKineticDataAvailable, Detail: "finite trace weights are available from Gate 60 and inherited through Gates 61-62"},
			{Name: "unit-propagator diagnostic", Passed: a.UnitPropagatorBranchAvailable, Detail: fmt.Sprintf("G_hat_unit(diagnostic)=%.10f; dominant sector=%s", a.UnitAttractiveKernel, a.DominantUnitSector)},
			{Name: "inverse-kinetic diagnostic", Passed: a.InverseKineticBranchAvailable, Detail: fmt.Sprintf("Σ contribution/traceWeight=%.10f; diagnostic only", a.InverseKineticDiagnostic)},
			{Name: "kinetic-weighted diagnostic", Passed: a.KineticWeightedBranchAvailable, Detail: fmt.Sprintf("Σ contribution·traceWeight=%.10f; dominant sector=%s", a.KineticWeightedDiagnostic, a.DominantKineticWeightedSector)},
			{Name: "sector propagator diagnostics", Passed: len(a.SectorDiagnostics) > 0, Detail: FormatSectorDiagnostics(a.SectorDiagnostics)},
			{Name: "exchange action sign", Passed: a.ExchangeActionSignDerived, Detail: "open; finite action has not selected the attractive exchange orientation"},
			{Name: "propagator weights", Passed: a.PropagatorWeightsDerived, Detail: "open; no current-sector denominator M_A^{-2} or spectral propagator was derived"},
			{Name: "relative current couplings", Passed: a.RelativeCurrentCouplingsDerived, Detail: "open; unit and kinetic-weighted branches are diagnostics, not coupling derivations"},
			{Name: "finite exchange kernel", Passed: a.FiniteExchangeKernelDerived, Detail: "open; G_hat requires action sign plus propagator weights plus couplings"},
			{Name: "attractive scalar-channel theorem", Passed: a.AttractiveScalarChannelDerived, Detail: "open; conditional attraction is not a derived finite theorem"},
			{Name: "up/down splitting", Passed: a.UpDownSplittingDerived, Detail: "open; no finite exchange data selects top-like up over bottom-like down"},
			{Name: "regulator criticality", Passed: a.RegulatorCriticalityDerived, Detail: "open; finite NJL critical threshold C_reg is still missing"},
			{Name: "condensation claim", Passed: a.CondensationClaimAllowed, Detail: "false by design; no top condensation, Higgs VEV, or mass scale is claimed"},
			{Name: "hidden observed input", Passed: !a.HiddenObservedInputUsed, Detail: "no observed coupling, y_t, v, Higgs mass, or threshold scale was inserted"},
		}, Notes: []string{
			a.TruthStatement,
			fmt.Sprintf("recommended next gate: %s", a.RecommendedNextGate),
			fmt.Sprintf("remaining unknowns: %s", FormatUnknowns(a.RemainingUnknowns)),
		}}
	}}
}
