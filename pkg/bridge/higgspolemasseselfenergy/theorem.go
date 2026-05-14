package higgspolemasseselfenergy

import "github.com/bagherbal/asha-engine/pkg/theorem"

func HiggsPoleSelfEnergyTargetMinimalPrecisionCorrectionAuditTheorem() theorem.Theorem {
	const id = "BRIDGE-HIGGS-POLE-SELF-ENERGY-TARGET-MINIMAL-PRECISION-CORRECTION-AUDIT"
	const name = "Higgs Pole Self-Energy Target / Minimal Precision Correction Audit"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "build Gate 332 self-energy target audit", Passed: false, Detail: err.Error()}}}
		}
		checks := []theorem.Check{
			{Name: "Gate 331 precision gap inherited", Passed: a.Inputs.HighestInheritedGate == inheritedHighestGate && a.Inputs.UsesObservedForTarget && a.Inputs.TreeMassGeV > a.Inputs.ObservedPoleGeV, Detail: FormatInputs(a.Inputs)},
			{Name: "pole equation self-energy target formalized", Passed: a.Equation.RequiredRePiGeV2 > 0 && a.Equation.DeltaPoleMinusRunGeV2 < 0, Detail: FormatEquation(a.Equation)},
			{Name: "required self-energy target computed", Passed: a.Target.RequiredRePiGeV2 > 40 && a.Target.RequiredRePiGeV2 < 50 && a.Target.RequiredMassShiftGeV < 0, Detail: FormatTarget(a.Target)},
			{Name: "required correction has natural one-loop scale", Passed: a.Capacity.RequiredIsOrderOneLoop && a.Capacity.RequiredIsSmallFraction, Detail: FormatCapacity(a.Capacity)},
			{Name: "precision ledger requires explicit SM self-energies and scheme", Passed: a.Ledger.NeedsTopLoop && a.Ledger.NeedsWZLoops && a.Ledger.NeedsCounterterms && a.Ledger.NeedsSchemeChoice && !a.Ledger.FullCalculationExecuted, Detail: FormatLedger(a.Ledger)},
			{Name: "firewall preserves no exact collider pole-mass claim", Passed: a.Audit.NoLoopIntegralsEvaluated && a.Audit.NoTwoLoopClaim && a.Audit.NoExactPoleMassClaim && a.Audit.NoFitParameterIntroduced, Detail: FormatAudit(a.Audit)},
		}
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Checks: checks, Notes: []string{a.Truth, FormatSummary(a.Summary), FormatStatuses(Statuses(a))}}
	}}
}
