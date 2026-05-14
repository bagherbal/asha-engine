package perslotmonotonicityseal

import "github.com/bagherbal/asha-engine/pkg/theorem"

func PerSlotMonotonicitySealFinalSpectralSynthesisAuditTheorem() theorem.Theorem {
	const id = "BRIDGE-PER-SLOT-MONOTONICITY-SEAL-FINAL-SPECTRAL-SYNTHESIS-AUDIT"
	const name = "Per-Slot Monotonicity Seal / Final Spectral Synthesis Audit"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "build Gate 291 final spectral synthesis audit", Passed: false, Detail: err.Error()}}}
		}
		checks := []theorem.Check{
			{Name: "Gate 290 per-slot diagnostic is inherited but not native", Passed: a.Inheritance.PerSlotDiagnosticSelectsRPlus && !a.Inheritance.PerSlotRuleNativeTheorem && !a.Inheritance.BranchPreviouslySelected, Detail: a.Inheritance.Verdict},
			{Name: "PerSlotMonotonicitySeal is active and explicitly phenomenological", Passed: a.Seal.Active && a.Seal.Phenomenological && !a.Seal.NativeTheorem && a.Seal.SelectedBranch == "r_plus", Detail: FormatSeal(a.Seal)},
			{Name: "branch ledger evaluates r_plus and r_minus", Passed: len(a.Branches) == 2, Detail: FormatBranch(a.Branches[0]) + " | " + FormatBranch(a.Branches[1])},
			{Name: "sealed vacuum locks r_plus and vetoes r_minus", Passed: a.Locked.UniqueUnderSeal && !a.Locked.NativeUnique && a.Locked.SelectedBranch.Name == "r_plus" && len(a.Locked.VetoedBranches) == 1 && a.Locked.VetoedBranches[0] == "r_minus", Detail: FormatLocked(a.Locked)},
			{Name: "locked trace moments are positive and computed", Passed: a.Trace.D2 > 0 && a.Trace.D4 > 0, Detail: FormatTrace(a.Trace)},
			{Name: "raw shape proxy reproduces Gate 169 contact scalar shape", Passed: a.Trace.ShapeMatchesContact && a.Trace.LambdaResidualAbs < 1e-12, Detail: FormatTrace(a.Trace)},
			{Name: "Higgs firewall remains active", Passed: !a.Higgs.PhysicalHiggsPredictionClaimed && !a.Higgs.HeatKernelProjectionDerived && !a.Higgs.ScalarGaugeNormalization, Detail: FormatHiggs(a.Higgs)},
			{Name: "firewalls prevent sealed branch from rewriting finite theorem status", Passed: a.Firewalls.SealDoesNotRewriteGate290 && a.Firewalls.RMinusVetoMarkedSealed && a.Firewalls.RawProxyNotPromotedToA2A4 && !a.Firewalls.FiniteCorePolluted, Detail: FormatFirewalls(a.Firewalls)},
		}
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Checks: checks, Notes: []string{a.Truth, FormatSummary(a.Summary)}}
	}}
}
