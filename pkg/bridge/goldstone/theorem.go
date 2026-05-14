package goldstone

import (
	"fmt"

	"github.com/bagherbal/asha-engine/pkg/theorem"
)

func GaugeEatingGoldstoneAuditTheorem() theorem.Theorem {
	const id = "BRIDGE-GOLDSTONE-GAUGE-EATING-AUDIT"
	const name = "finite Goldstone / gauge-eating correspondence audit"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "construct Goldstone/gauge-eating audit", Passed: false, Detail: err.Error()}}}
		}
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Checks: []theorem.Check{
			{Name: "scalar radial/angular split", Passed: a.ActiveRealDirections == 4 && a.RadialDirections == 1 && a.ScalarAngularDirections == 3, Detail: fmt.Sprintf("active scalar/contact sector: %d real directions = %d radial + %d angular", a.ActiveRealDirections, a.RadialDirections, a.ScalarAngularDirections)},
			{Name: "protected contact resonance", Passed: a.ScalarGoldstoneCountMatchesProtected, Detail: fmt.Sprintf("angular scalar directions=%d and protected contact directions=%d", a.ScalarAngularDirections, a.ProtectedContactDirections)},
			{Name: "electroweak broken-direction count", Passed: a.BrokenGaugeCountMatchesGoldstone, Detail: fmt.Sprintf("dim(SU(2)_L)+dim(U(1)_Y)-dim(U(1)_em)=%d+%d-%d=%d", a.SU2LGeneratorCount, a.HyperchargeGeneratorCount, a.UnbrokenEMDimension, a.BrokenGaugeDirections)},
			{Name: "Goldstone count-level correspondence", Passed: a.GoldstoneCountResonance, Detail: "3 scalar angular directions ↔ 3 protected contact directions ↔ 3 broken electroweak directions"},
			{Name: "canonical protected-to-broken map", Passed: a.CanonicalProtectedToBrokenMapDerived, Detail: "not derived; count equality is not yet a canonical isometry/intertwiner"},
			{Name: "SU(2)_L action on finite scalar/contact frame", Passed: a.SU2LActionOnContactScalarDerived, Detail: "not derived; current SU(2)_L action is on the audited matter doublet representation"},
			{Name: "finite covariant derivative", Passed: a.CovariantDerivativeDerived, Detail: "not derived; DΦ and kinetic normalization are required for gauge eating"},
			{Name: "gauge-boson mass matrix", Passed: a.GaugeBosonMassMatrixDerived, Detail: "not derived; no W/Z mass or mixing comparison is allowed"},
			{Name: "gauge-eating theorem", Passed: a.GaugeEatingTheoremDerived, Detail: "open; the current result is a structured count-level resonance, not a completed theorem"},
		}, Notes: []string{
			a.TruthStatement,
			fmt.Sprintf("classification: %s", a.Classification),
			fmt.Sprintf("remaining missing data: %s", FormatMissing(a.MissingData)),
		}}
	}}
}
