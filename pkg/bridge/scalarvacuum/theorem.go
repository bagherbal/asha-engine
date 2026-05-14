package scalarvacuum

import (
	"fmt"

	"github.com/bagherbal/asha-engine/pkg/theorem"
)

func ScalarVacuumOrientationFiniteMinimizerSearchTheorem() theorem.Theorem {
	const id = "BRIDGE-SCALAR-VACUUM-ORIENTATION-MINIMIZER"
	const name = "scalar vacuum orientation and finite minimizer search"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "build scalar vacuum orientation audit", Passed: false, Detail: err.Error()}}}
		}
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Checks: []theorem.Check{
			{Name: "Gate 85 gauge-eating diagnostic input", Passed: a.GaugeEatingDiagnosticAvailable, Detail: "broken-generator image signature is available, but full gauge-eating theorem remains open"},
			{Name: "radial normal form selects radius", Passed: a.RadialNormalFormSelectsRadius, Detail: fmt.Sprintf("r0²=%.10f, r0=%.10f; vacuum manifold %s", a.VacuumRadiusSquared, a.VacuumRadius, a.RadialVacuumManifold)},
			{Name: "radial normal form selects vector", Passed: a.RadialNormalFormSelectsVector, Detail: "not derived; shifted radial potential is orientation-degenerate on S^3"},
			{Name: "finite scalar-response pair spectrum", Passed: a.LowPairDimension == 2 && a.HighPairDimension == 2, Detail: fmt.Sprintf("spectrum=%s; low=%.10f×%d, high=%.10f×%d, split=%.10f", FormatSpectrum(a.ActiveSpectrum), a.LowPairEigenvalue, a.LowPairDimension, a.HighPairEigenvalue, a.HighPairDimension, a.PairSplit)},
			{Name: "low active pair selected", Passed: a.LowPairSelected, Detail: fmt.Sprintf("constrained q(φ)=φᵀSφ minimizer at |φ|=r0 lies in the low pair; Emin=%.10f, Ehigh=%.10f, gap=%.10f", a.MinResponseEnergy, a.MaxResponseEnergy, a.EnergyGapAtRadius)},
			{Name: "diagnostic unitary-gauge vector is a minimizer", Passed: a.DiagnosticVacuumIsMinimizer, Detail: fmt.Sprintf("φ0=%s has q(φ0)=%.10f", FormatVector(a.DiagnosticVacuumVector), a.DiagnosticVacuumEnergy)},
			{Name: "residual phase/orientation freedom", Passed: !a.CanonicalPhaseSelected, Detail: fmt.Sprintf("low-pair minimizer leaves S^1 freedom; residual dimension=%d", a.ResidualPhaseFreedomDimension)},
			{Name: "finite vacuum orientation derived", Passed: a.FiniteVacuumOrientationDerived, Detail: "not fully derived; the low pair is selected, but the exact vector/phase inside it remains gauge or bridge data"},
			{Name: "full gauge-eating theorem", Passed: a.FullGaugeEatingTheoremDerived, Detail: "not derived; still requires protected-contact/broken-generator intertwiner plus action-selected kinetic normalizations"},
		}, Notes: []string{a.TruthStatement, fmt.Sprintf("remaining unknowns: %s", formatUnknowns(a.RemainingUnknowns)), "Next: " + a.RecommendedNextGate}}
	}}
}

func formatUnknowns(values []string) string {
	if len(values) == 0 {
		return "none"
	}
	out := ""
	for i, v := range values {
		if i > 0 {
			out += "; "
		}
		out += v
	}
	return out
}
