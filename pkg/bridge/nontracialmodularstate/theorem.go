package nontracialmodularstate

import "github.com/bagherbal/asha-engine/pkg/theorem"

func NontracialModularStateOriginVacuumDensityMatrixDerivationAuditTheorem() theorem.Theorem {
	const id = "BRIDGE-NONTRACIAL-MODULAR-STATE-ORIGIN-VACUUM-DENSITY-MATRIX-DERIVATION"
	const name = "Nontracial Modular State Origin / Vacuum Density Matrix Derivation Audit"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "build Gate 364 audit", Passed: false, Detail: err.Error()}}}
		}
		checks := []theorem.Check{
			{Name: "topological density sourcing is audited", Passed: a.Topological.Formalized && !a.Topological.SignedTauDensityValid && !a.Topological.NativeNonTracialFound, Detail: FormatTopological(a.Topological)},
			{Name: "tau magnitude candidates are faithful and nontracial but retain 1-2 degeneracy", Passed: a.Topological.MagnitudeState.Faithful && !a.Topological.MagnitudeState.Tracial && a.Topological.MagnitudeState.ResidualDeg12 && a.Topological.SquaredMagnitudeState.ResidualDeg12, Detail: FormatTopological(a.Topological)},
			{Name: "KMS exponential state is formalized and activates nontrivial modular frequencies", Passed: a.KMS.Formalized && a.KMS.NonTrivial && !a.KMS.Mandated && allPairFrequenciesNonZero(a.KMS.State), Detail: FormatKMS(a.KMS)},
			{Name: "flow activation capacity is detected but no native mandate is present", Passed: a.Flow.Executed && a.Flow.AnyNonTrivial && !a.Flow.AnyMandatedNativeNontracial && !a.Flow.SelectsUniqueVacuum, Detail: FormatFlow(a.Flow)},
			{Name: "vacuum census is not reduced", Passed: a.Census.StartingInputs == 15 && a.Census.Reduction == 0 && a.Census.RemainingInputs == 15, Detail: FormatCensus(a.Census)},
		}
		passed := 0
		for _, c := range checks {
			if c.Passed {
				passed++
			}
		}
		status := theorem.BridgeRequired
		if passed != len(checks) {
			status = theorem.FailedRoute
		}
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: status, Checks: checks}
	}}
}
