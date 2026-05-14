package casimirkernel

import (
	"fmt"

	"github.com/bagherbal/asha-engine/pkg/theorem"
)

func CurrentSectorCasimirPropagatorDiagnosticTheorem() theorem.Theorem {
	const id = "BRIDGE-CURRENT-SECTOR-CASIMIR-PROPAGATOR-DIAGNOSTIC"
	const name = "current-sector Casimir / propagator diagnostic"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "build Casimir diagnostics", Passed: false, Detail: err.Error()}}}
		}
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Checks: []theorem.Check{
			{Name: "Gate 66 sector-operator input", Passed: a.Previous.SectorOperatorsConstructed && a.Previous.RepresentationLevelMapsDerived, Detail: "central/color/B-L/leptoquark Casimir operators C_A=sum T_a^T T_a are available"},
			{Name: "Casimir diagnostic families", Passed: a.AllCasimirDiagnosticsBuilt, Detail: FormatDiagnostics(a.Diagnostics)},
			{Name: "direct kernel family exposed", Passed: a.DirectKernelFamilyExposed, Detail: fmt.Sprintf("Σ Tr(C_A)=%.10f; dominant direct sector=%s", a.TotalDirectTrace, a.DominantDirectSector)},
			{Name: "inverse-nonzero kernel family exposed", Passed: a.InverseKernelFamilyExposed, Detail: fmt.Sprintf("Σ Tr(C_A^+)=%.10f; dominant inverse sector=%s", a.TotalInverseTrace, a.DominantInverseSector)},
			{Name: "trace-normalized family exposed", Passed: a.TraceNormalizedFamilyExposed, Detail: "C_A/Tr(C_A) gives sector probability-like diagnostics but not exchange denominators"},
			{Name: "color-sector zero-mode audit", Passed: a.ColorSectorZeroMode, Detail: "color-su3 Casimir annihilates the lepton seed and acts uniformly on the three color seeds"},
			{Name: "diagnostic ambiguity", Passed: true, Detail: a.AmbiguityStatement},
			{Name: "finite action selection", Passed: a.FiniteActionSelectionDerived, Detail: "open; no finite exchange action selects direct, inverse, or trace-normalized Casimir propagation"},
			{Name: "propagator denominators", Passed: a.PropagatorDenominatorsDerived, Detail: "open; C_A diagnostics are representation/kinetic data, not derived rho_A denominators"},
			{Name: "exchange kernel update", Passed: a.ExchangeKernelUpdated, Detail: "open; G_hat cannot be updated from diagnostics until a propagator rule is selected"},
			{Name: "attractive scalar-channel theorem", Passed: a.AttractiveScalarChannelDerived, Detail: "open; Casimir diagnostics do not derive exchange sign, NJL attraction, or condensation"},
			{Name: "up/down splitting", Passed: a.UpDownSplittingDerived, Detail: "open; current-sector Casimirs act on lepton/color flavor and still do not distinguish up-type from down-type channels"},
			{Name: "condensation claim", Passed: a.CondensationClaimAllowed, Detail: "false by design; no top condensation, Higgs VEV, or fermion mass is claimed"},
			{Name: "hidden observed input", Passed: !a.HiddenObservedInputUsed, Detail: "no observed threshold, coupling, y_t, v, Higgs mass, or fermion mass was inserted"},
		}, Notes: []string{
			a.TruthStatement,
			fmt.Sprintf("recommended next gate: %s", a.RecommendedNextGate),
			fmt.Sprintf("remaining unknowns: %s", FormatUnknowns(a.RemainingUnknowns)),
		}}
	}}
}
