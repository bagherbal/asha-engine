package exchangekernel

import (
	"fmt"

	"github.com/bagherbal/asha-engine/pkg/theorem"
)

func PropagatorExchangeKernelAuditTheorem() theorem.Theorem {
	const id = "BRIDGE-PROPAGATOR-EXCHANGE-KERNEL-AUDIT"
	const name = "Propagator/action sign and exchange-kernel audit"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "construct exchange-kernel audit", Passed: false, Detail: err.Error()}}}
		}
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Checks: []theorem.Check{
			{Name: "signed Fierz input", Passed: a.SignedFierzAvailable, Detail: fmt.Sprintf("c_LR^scalar=%.10f, weighted sum=%.10f", a.Fierz.UniversalLRScalarFierzCoefficient, a.Fierz.TotalWeightedSignedCoefficient)},
			{Name: "NJL attractiveness convention", Passed: true, Detail: a.NJLAttractiveConvention},
			{Name: "exchange-orientation branches", Passed: a.ConditionalAttractiveBranchAvailable, Detail: fmt.Sprintf("unit +J² total=%.10f; unit -J² total=%.10f", a.UnitExchangePlusTotal, a.UnitExchangeMinusTotal)},
			{Name: "conditional attractive unit kernel", Passed: a.ConditionalAttractiveBranchAvailable, Detail: fmt.Sprintf("if finite exchange has the attractive orientation, G_hat_unit=%.10f; opposite orientation gives %.10f", a.ConditionalAttractiveUnitKernel, a.ConditionalRepulsiveUnitKernel)},
			{Name: "sector exchange diagnostics", Passed: a.ConditionalAttractiveBranchAvailable, Detail: FormatSectorDiagnostics(a.SectorDiagnostics)},
			{Name: "exchange action sign", Passed: a.ExchangeSignDerived, Detail: "open; finite action has not selected +J² or -J² orientation"},
			{Name: "propagator magnitude", Passed: a.PropagatorMagnitudeDerived, Detail: "open; no finite M_A^{-2} or spectral propagator analogue has been derived"},
			{Name: "relative current couplings", Passed: a.RelativeCurrentCouplingsDerived, Detail: "open; sectors still use unit/trace-normalized weights only"},
			{Name: "attractive scalar-channel sign", Passed: a.AttractiveScalarChannelSignDerived, Detail: "open; conditional attraction is not a derived attraction theorem"},
			{Name: "native four-fermion kernel", Passed: a.NativeFourFermionKernelDerived, Detail: "open; G_hat requires exchange sign, propagator magnitude, and relative couplings"},
			{Name: "up/down splitting", Passed: a.UpDownSplittingDerived, Detail: "open; exchange-kernel audit still does not select top-like up over bottom-like down"},
			{Name: "regulator criticality", Passed: a.RegulatorCriticalityDerived, Detail: "open; C_reg for the finite NJL gap equation is not derived"},
			{Name: "condensation claim", Passed: a.CondensationClaimAllowed, Detail: "false by design; no Higgs VEV, top condensation, or mass scale is claimed"},
			{Name: "hidden observed input", Passed: !a.HiddenObservedInputUsed, Detail: "no observed y_t, v, Higgs mass, coupling, or physical scale was inserted"},
		}, Notes: []string{
			a.TruthStatement,
			fmt.Sprintf("current-current template: %s", a.CurrentCurrentTemplate),
			fmt.Sprintf("recommended next gate: %s", a.RecommendedNextGate),
			fmt.Sprintf("remaining unknowns: %s", FormatUnknowns(a.RemainingUnknowns)),
		}}
	}}
}
