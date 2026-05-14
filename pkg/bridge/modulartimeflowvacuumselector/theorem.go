package modulartimeflowvacuumselector

import "github.com/bagherbal/asha-engine/pkg/theorem"

func ModularTimeFlowVacuumSelectorExtensionAuditTheorem() theorem.Theorem {
	const id = "BRIDGE-MODULAR-TIME-FLOW-DYNAMICAL-VACUUM-SELECTOR-EXTENSION"
	const name = "Modular Time Flow / Dynamical Vacuum Selector Extension Audit"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "build Gate 362 audit", Passed: false, Detail: err.Error()}}}
		}
		checks := []theorem.Check{
			{Name: "Path B is activated without adding a fit", Passed: a.Span.InheritedGate == 361 && a.Span.Path == "B: minimal dynamical extension" && !a.Span.AddsFit, Detail: FormatSpan(a.Span)},
			{Name: "documentation shift is formalized", Passed: a.Docs.Required && a.Docs.AppliedInReadme && a.Docs.AppliedInDocs, Detail: FormatDocs(a.Docs)},
			{Name: "new flow operator class is introduced", Passed: a.Sieve.Executed && a.Sieve.Candidate.NewOperatorClass && a.Sieve.Candidate.NativeCandidate, Detail: FormatSieve(a.Sieve)},
			{Name: "flow admissibility constraints preserve the landscape", Passed: a.Sieve.PreservesLandscape && len(a.Sieve.Axioms) >= 5, Detail: FormatSieve(a.Sieve)},
			{Name: "explicit kernel remains firewalled", Passed: !a.Sieve.ExplicitKernelConstructed && !a.Sieve.VacuumSelected && a.Census.RemainingInputs == 15, Detail: FormatCensus(a.Census)},
			{Name: "Phase III program is installed", Passed: a.Program.Formalized && len(a.Program.RequiredArtifacts) >= 5 && len(a.Program.ForbiddenMoves) >= 4, Detail: FormatProgram(a.Program)},
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
