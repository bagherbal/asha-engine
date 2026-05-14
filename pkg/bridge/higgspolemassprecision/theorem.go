package higgspolemassprecision

import "github.com/bagherbal/asha-engine/pkg/theorem"

func HiggsPoleMassConversionPrecisionGapLedgerAuditTheorem() theorem.Theorem {
	const id = "BRIDGE-HIGGS-POLE-MASS-CONVERSION-PRECISION-GAP-LEDGER-AUDIT"
	const name = "Higgs Pole-Mass Conversion / Precision Gap Ledger Audit"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "build Gate 331 pole mass precision audit", Passed: false, Detail: err.Error()}}}
		}
		checks := []theorem.Check{
			{Name: "Gate 330 doubled-trace branch inherited without observed-mass fit", Passed: a.Inputs.HighestInheritedGate == inheritedHighestGate && !a.Inputs.AddsFit && nearlyEqual(a.Inputs.GStarSquared, 0.5, 1e-15), Detail: FormatInputs(a.Inputs)},
			{Name: "native tree-level Higgs proxy recomputed", Passed: a.Tree.LambdaH > 0.129 && a.Tree.LambdaH < 0.130 && a.Tree.MassGeV > 125.0 && a.Tree.MassGeV < 125.6, Detail: FormatTree(a.Tree)},
			{Name: "precision gap is quantified as sub-GeV", Passed: a.Capacity.GapIsSubGeV && a.Comparison.RelativeMassErrorPct > 0 && a.Comparison.RelativeMassErrorPct < 0.2, Detail: FormatComparison(a.Comparison)},
			{Name: "pole-mass conversion ledger is formalized but not executed", Passed: a.Pole.RequiresTopSelfEnergy && a.Pole.RequiresWeakBosonSelfEnergy && a.Pole.RequiresScalarSelfEnergy && a.Pole.RequiresRenormalizationScheme && !a.Pole.Executed, Detail: FormatPole(a.Pole)},
			{Name: "precision correction capacity can cover the residual without new structural threshold", Passed: a.Capacity.PerturbativePoleCorrectionsCanHaveThisScale && a.Capacity.TwoLoopRGCanHaveThisScale && a.Capacity.ThresholdRetuningNotRequiredForThisGap, Detail: FormatCapacity(a.Capacity)},
			{Name: "firewalls preserve no exact collider mass claim", Passed: a.Audit.NoObservedMassFitted && a.Audit.NoSelfEnergiesComputed && a.Audit.NoTwoLoopClaim && a.Audit.NoExactColliderClaim, Detail: FormatAudit(a.Audit)},
		}
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Checks: checks, Notes: []string{a.Truth, FormatSummary(a.Summary), FormatStatuses(Statuses(a))}}
	}}
}
