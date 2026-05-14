package fierzsign

import (
	"fmt"

	"github.com/bagherbal/asha-engine/pkg/theorem"
)

func CliffordLorentzFierzSignTheorem() theorem.Theorem {
	const id = "BRIDGE-CLIFFORD-LORENTZ-FIERZ-SIGN"
	const name = "Clifford/Lorentz Fierz sign construction for LR scalar channel"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "construct Clifford/Lorentz Fierz sign audit", Passed: false, Detail: err.Error()}}}
		}
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Checks: []theorem.Check{
			{Name: "kinetic-normalized current input", Passed: a.Kinetic.KineticTraceNormalizationDerived, Detail: fmt.Sprintf("sectors=%d, total kinetic trace=%.10f", len(a.Kinetic.SectorNormalizations), a.Kinetic.TotalKineticTrace)},
			{Name: "sigma/bar-sigma completeness identity", Passed: a.SigmaIdentityConstructed, Detail: fmt.Sprintf("max residual in σ^μ\u0304σ_μ identity = %.3e", a.SigmaBarSigmaIdentityResidual)},
			{Name: "spinor completeness factor", Passed: a.SpinorCompletenessFactor == 2, Detail: fmt.Sprintf("two-component spinor trace gives factor %.1f", a.SpinorCompletenessFactor)},
			{Name: "fermion reordering sign", Passed: a.FermionReorderingIncluded, Detail: fmt.Sprintf("fermion exchange sign = %.1f", a.FermionReorderingSign)},
			{Name: "universal LR scalar Fierz coefficient", Passed: a.SignedCliffordFierzCoefficientsDerived, Detail: fmt.Sprintf("c_LR^scalar = %.10f", a.UniversalLRScalarFierzCoefficient)},
			{Name: "sector signed Fierz diagnostics", Passed: a.SignedCliffordFierzCoefficientsDerived, Detail: FormatSectorCoefficients(a.SectorCoefficients)},
			{Name: "weighted signed diagnostic", Passed: a.SignedCliffordFierzCoefficientsDerived, Detail: fmt.Sprintf("Σ weight·c_signed = %.10f", a.TotalWeightedSignedCoefficient)},
			{Name: "generator propagator/action normalization", Passed: a.GeneratorPropagatorNormalizationDerived, Detail: "open; signed Fierz coefficients still need finite exchange action and propagator signs"},
			{Name: "attractive scalar-channel sign", Passed: a.AttractiveScalarChannelSignDerived, Detail: "open; negative Fierz coefficient is not yet an NJL attraction without exchange-action orientation"},
			{Name: "native four-fermion kernel", Passed: a.NativeFourFermionKernelDerived, Detail: "open; G_hat requires signed Fierz coefficients plus propagator and relative current weights"},
			{Name: "up/down splitting", Passed: a.UpDownSplittingDerived, Detail: "open; signed Fierz construction still does not distinguish top-like up from bottom-like down"},
			{Name: "hidden observed input", Passed: !a.HiddenObservedInputUsed, Detail: "no observed y_t, v, Higgs mass, coupling, or physical scale was inserted"},
		}, Notes: []string{
			a.TruthStatement,
			fmt.Sprintf("recommended next gate: %s", a.RecommendedNextGate),
			fmt.Sprintf("remaining unknowns: %s", FormatUnknowns(a.RemainingUnknowns)),
		}}
	}}
}
