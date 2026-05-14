package intermediatethresholdjump

import "github.com/bagherbal/asha-engine/pkg/theorem"

func IntermediateThresholdDecouplingQuarticJumpTransportAuditTheorem() theorem.Theorem {
	const id = "BRIDGE-INTERMEDIATE-THRESHOLD-DECOUPLING-QUARTIC-JUMP-TRANSPORT-AUDIT"
	const name = "Intermediate Threshold Decoupling / Quartic Jump Transport Audit"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "build Gate 314 threshold jump audit", Passed: false, Detail: err.Error()}}}
		}
		checks := []theorem.Check{
			{Name: "two-stage RG transport is formalized with GUT, PeV threshold, and electroweak scales", Passed: a.RG.Formalized && a.RG.HighScaleGeV > 1e15 && a.RG.ThresholdGeV > 1e6 && a.RG.LowScaleGeV > 240 && a.RG.BoundaryLambdaUV > 0.25 && a.RG.TargetLambdaAtV > 0.12 && a.RG.TargetLambdaAtV < 0.14, Detail: FormatRG(a.RG)},
			{Name: "finite threshold matching rule is explicit and has the correct lowering sign", Passed: a.Rule.Formalized && a.Rule.NegativeJumpLowersIRMass && !a.Rule.DerivedFromHeavySector && a.Rule.AppliedAtGeV == thresholdScaleGeV, Detail: FormatRule(a.Rule)},
			{Name: "transport lanes include legacy, tau_eta-low, and gauge-only lower-envelope diagnostics", Passed: len(a.Lanes) == 3 && a.Lanes[0].TopFraction == 1 && a.Lanes[1].TopFraction > 0.11 && a.Lanes[1].TopFraction < 0.112 && a.Lanes[2].TopFraction == 0, Detail: FormatLane(a.Lanes[0]) + " || " + FormatLane(a.Lanes[1]) + " || " + FormatLane(a.Lanes[2])},
			{Name: "required Δλ is solved for the preferred gauge-only lower-envelope lane while high-top lanes remain diagnostic", Passed: len(a.Jumps) == len(a.Lanes) && !a.Jumps[0].Solved && !a.Jumps[1].Solved && a.Jumps[2].Solved && a.Jumps[2].RequiredDeltaLambda < -0.01 && a.Jumps[2].RequiredDeltaLambda > -1.0 && a.Jumps[2].CorrectSign && a.Jumps[2].ViableOrder, Detail: FormatJump(a.Jumps[0]) + " || " + FormatJump(a.Jumps[1]) + " || " + FormatJump(a.Jumps[2])},
			{Name: "preferred gauge-only floor requires a negative moderate PeV threshold jump to reach 125.10 GeV", Passed: a.Viability.Formalized && a.Viability.PreferredLaneName == "gauge_only_zero_top_lower_envelope" && a.Viability.PreferredBaselineMassGeV > 150 && a.Viability.PreferredBaselineMassGeV < 165 && a.Viability.PreferredRequiredDelta < -0.01 && a.Viability.PreferredRequiredDelta > -1.0 && a.Viability.JumpIsNegative && a.Viability.JumpMagnitudeModerate && a.Viability.MatchesScalarPortalSign && a.Viability.CanBeGeneratedByTreePortal && !a.Viability.HeavySectorDerived, Detail: FormatViability(a.Viability)},
			{Name: "portal target is recorded without fitting or deriving the heavy sector", Passed: a.Viability.RequiredPortalRatio > 0 && a.Viability.IfLambdaHeavyEqualsOne > 0 && a.Firewalls.NoThresholdJumpDerived && a.Firewalls.NoPortalCouplingFitted && a.Firewalls.NoHeavySelfQuarticFitted && a.Firewalls.NoThresholdScaleDerived && a.Firewalls.NoTwoLoopRGExecuted && a.Firewalls.NoPoleMassConversionInserted && a.Firewalls.NoFinalMassClaimed && !a.Firewalls.FiniteCorePolluted, Detail: FormatViability(a.Viability) + " || " + FormatFirewalls(a.Firewalls)},
			{Name: "summary records a quantitative threshold obligation without claiming a final Higgs mass", Passed: a.Summary.TwoStageRGFormalized && a.Summary.JumpInsertionFormalized && a.Summary.RequiredJumpExtracted && a.Summary.JumpHasCorrectSign && a.Summary.JumpHasPortalMagnitude && !a.Summary.HeavySectorDerived && !a.Summary.FinalMassClaimed && a.Summary.FirewallPreserved, Detail: FormatSummary(a.Summary)},
		}
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Checks: checks, Notes: []string{a.Truth, "Gate 314 extracts the finite matching jump required at the inherited PeV threshold; it does not derive that jump from the finite heavy-sector graph.", "The preferred gauge-only lower-envelope lane is used to quantify the minimal discontinuous obligation because Gate 313 proved continuous RG transport alone cannot reach the 125 GeV comparison target."}}
	}}
}
