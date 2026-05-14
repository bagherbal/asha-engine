package kineticnorm

import (
	"fmt"

	"github.com/bagherbal/asha-engine/pkg/theorem"
)

func GeneratorKineticNormalizationTheorem() theorem.Theorem {
	const id = "BRIDGE-GENERATOR-KINETIC-NORMALIZATION-FIERZ"
	const name = "generator kinetic normalization / signed Fierz coefficient audit"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "construct generator kinetic normalization audit", Passed: false, Detail: err.Error()}}}
		}
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Checks: []theorem.Check{
			{Name: "current projection input", Passed: a.Current.UnsignedScalarProjectionCoefficientsKnown, Detail: fmt.Sprintf("sectors=%d, max current/projector residual=%.3e", len(a.Current.SectorCoefficients), a.Current.MaxIntertwinerResidual)},
			{Name: "finite kinetic trace normalization", Passed: a.KineticTraceNormalizationDerived, Detail: fmt.Sprintf("total finite current kinetic trace=%.10f", a.TotalKineticTrace)},
			{Name: "sector trace weights", Passed: a.KineticTraceNormalizationDerived, Detail: FormatSectorNormalizations(a.SectorNormalizations)},
			{Name: "unsigned unit scalar projection coefficients", Passed: a.UnsignedUnitProjectionCoefficientsDerived, Detail: fmt.Sprintf("max |c_unsigned−1| = %.3e", a.MaxUnitOverlapResidual)},
			{Name: "signed Clifford/Fierz coefficients", Passed: a.SignedCliffordFierzCoefficientsDerived, Detail: "open; trace-normalized overlaps are not Lorentz/Fierz-signed coefficients"},
			{Name: "generator propagator/action normalization", Passed: a.GeneratorPropagatorNormalizationDerived, Detail: "open; finite kinetic trace does not yet fix exchange propagator weights or signs"},
			{Name: "attractive scalar-channel sign", Passed: a.AttractiveScalarChannelSignDerived, Detail: "open; NJL attraction cannot be inferred from positive overlap normalization"},
			{Name: "native four-fermion kernel", Passed: a.NativeFourFermionKernelDerived, Detail: "open; requires signed coefficients plus propagator/action normalization"},
			{Name: "up/down splitting", Passed: a.UpDownSplittingDerived, Detail: "open; kinetic normalization still does not distinguish top-like up from bottom-like down"},
			{Name: "hidden observed input", Passed: !a.HiddenObservedInputUsed, Detail: "no observed coupling, Yukawa, v, Higgs mass, or mass scale was inserted"},
		}, Notes: []string{
			a.TruthStatement,
			fmt.Sprintf("recommended next gate: %s", a.RecommendedNextGate),
			fmt.Sprintf("remaining unknowns: %s", FormatUnknowns(a.RemainingUnknowns)),
		}}
	}}
}
