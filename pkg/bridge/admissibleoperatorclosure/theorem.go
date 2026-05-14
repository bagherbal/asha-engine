package admissibleoperatorclosure

import "github.com/bagherbal/asha-engine/pkg/theorem"

func AdmissibleOperatorClosureVacuumSelectionNoGoTheorem() theorem.Theorem {
	const id = "BRIDGE-ADMISSIBLE-OPERATOR-CLOSURE-VACUUM-SELECTION-NO-GO"
	const name = "Admissible Operator Closure / Vacuum Selection No-Go Theorem"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "build Gate 361 audit", Passed: false, Detail: err.Error()}}}
		}
		checks := []theorem.Check{
			{Name: "span inherits Gate 360 without adding fit", Passed: a.Span.InheritedGate == 360 && !a.Span.AddsFit, Detail: FormatSpan(a.Span)},
			{Name: "operator classes are enumerated", Passed: a.Sieve.Executed && len(a.Sieve.Classes) >= 7 && a.Sieve.AllNativeAudited, Detail: FormatSieve(a.Sieve)},
			{Name: "no native kinetic-safe unique vacuum selector is present", Passed: a.Sieve.NoGoApplies && !a.Sieve.AnyKineticSafeSelector, Detail: FormatSieve(a.Sieve)},
			{Name: "no-go theorem is formalized", Passed: a.NoGo.Formalized && a.NoGo.RequiresExtension && a.NoGo.VacuumInputsRemain == 15, Detail: FormatNoGo(a.NoGo)},
			{Name: "extension fork is formalized", Passed: a.Extension.Formalized && len(a.Extension.MinimalNewObjects) >= 3, Detail: FormatExtension(a.Extension)},
			{Name: "vacuum census remains preserved", Passed: a.Census.ReductionFromClosure == 0 && a.Census.RemainingInputs == 15 && !a.Census.SevenSealReached, Detail: FormatCensus(a.Census)},
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
