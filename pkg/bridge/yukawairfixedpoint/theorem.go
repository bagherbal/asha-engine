package yukawairfixedpoint

import "github.com/bagherbal/asha-engine/pkg/theorem"

func YukawaInfraredFixedPointBasinRGAttractorReductionAuditTheorem() theorem.Theorem {
	const id = "BRIDGE-YUKAWA-INFRARED-FIXED-POINT-BASIN-RG-ATTRACTOR-REDUCTION-AUDIT"
	const name = "Yukawa Infrared Fixed-Point Basin / RG Attractor Reduction Audit"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "build Gate 353 RG attractor audit", Passed: false, Detail: err.Error()}}}
		}
		checks := []theorem.Check{
			{Name: "span inherits Gate 352 without adding a fit", Passed: a.Span.InheritedGate == 352 && !a.Span.AddsFit, Detail: FormatSpan(a.Span)},
			{Name: "one-loop third-generation RG equations are formalized", Passed: a.Equations.Formalized && a.Equations.OneLoop && a.Equations.TopBeta != "", Detail: FormatEquations(a.Equations)},
			{Name: "top quasi-fixed basin is detected but not a unique selector", Passed: a.Spiral.Audited && a.Spiral.QuasiFixedPoint && a.Spiral.ContractionRatio < 0.05 && !a.Spiral.ReductionProved, Detail: FormatSpiral(a.Spiral)},
			{Name: "r-plus boundary flows in high-top lane rather than reducing the parameter count", Passed: a.Spiral.Boundary.RPlusYtUV > 0.9 && a.Spiral.RPlusEndpointYt > 0.9 && a.Spiral.ParameterReduction == 0, Detail: FormatSpiral(a.Spiral)},
			{Name: "center criticality scan has no perturbative lambda-zero solution at M_int", Passed: a.Criticality.Formalized && !a.Criticality.PerturbativeSolution && a.Criticality.MinLambdaAtTarget > 0 && !a.Criticality.ReductionProved, Detail: FormatCriticality(a.Criticality)},
			{Name: "baryogenesis constraint is formalized but CP operator is not derived", Passed: a.Baryogenesis.Formalized && a.Baryogenesis.BGapLeptogenesisHasCapacity && !a.Baryogenesis.CPAsymmetryOperatorDerived && !a.Baryogenesis.ReductionProved, Detail: FormatBaryogenesis(a.Baryogenesis)},
			{Name: "parameter census remains at fifteen", Passed: a.Census.StartingVacuumInputs == 15 && a.Census.TotalReduction == 0 && a.Census.RemainingInputs == 15 && !a.Census.SevenSealReached, Detail: FormatCensus(a.Census)},
			{Name: "summary preserves dynamical vacuum-selection firewall", Passed: a.Summary.Executed && !a.Summary.AnyReductionProved && a.Summary.RemainingInputs == 15, Detail: FormatSummary(a.Summary)},
		}
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Checks: checks, Notes: []string{a.Truth, "Gate 353 shows that RG time is a real selection mechanism candidate, but in the installed one-loop ledger it does not yet consume any of the 15 minimal vacuum coordinates."}}
	}}
}
