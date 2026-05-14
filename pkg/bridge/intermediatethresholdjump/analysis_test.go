package intermediatethresholdjump

import "testing"

func TestTwoStageRGAndRule(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault returned error: %v", err)
	}
	if !a.RG.Formalized || a.RG.HighScaleGeV < 1e15 || a.RG.ThresholdGeV < 1e6 || a.RG.LowScaleGeV < 240 || a.RG.TargetLambdaAtV < 0.12 || a.RG.TargetLambdaAtV > 0.14 {
		t.Fatalf("bad RG ledger: %s", FormatRG(a.RG))
	}
	if !a.Rule.Formalized || !a.Rule.NegativeJumpLowersIRMass || a.Rule.DerivedFromHeavySector || a.Rule.AppliedAtGeV != thresholdScaleGeV {
		t.Fatalf("bad threshold rule: %s", FormatRule(a.Rule))
	}
}

func TestRequiredJumpExtraction(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault returned error: %v", err)
	}
	if len(a.Lanes) != 3 || len(a.Jumps) != 3 {
		t.Fatalf("expected three lanes and three jumps, got %d/%d", len(a.Lanes), len(a.Jumps))
	}
	if a.Jumps[0].Solved || a.Jumps[1].Solved {
		t.Fatalf("high-top diagnostic lanes should not be forced into a moderate threshold solution: %s || %s", FormatJump(a.Jumps[0]), FormatJump(a.Jumps[1]))
	}
	gauge := a.Jumps[2]
	if !gauge.Solved || !gauge.Perturbative || !gauge.CorrectSign || gauge.RequiredDeltaLambda >= -0.01 || gauge.RequiredDeltaLambda <= -1.0 || gauge.TargetMassGeV < 125.09 || gauge.TargetMassGeV > 125.11 {
		t.Fatalf("bad preferred gauge-only jump: %s", FormatJump(gauge))
	}
}

func TestViabilityAndFirewalls(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault returned error: %v", err)
	}
	if !a.Viability.Formalized || a.Viability.PreferredLaneName != "gauge_only_zero_top_lower_envelope" || a.Viability.PreferredBaselineMassGeV < 150 || a.Viability.PreferredBaselineMassGeV > 165 || a.Viability.PreferredRequiredDelta >= -0.01 || a.Viability.PreferredRequiredDelta <= -1.0 || !a.Viability.JumpIsNegative || !a.Viability.JumpMagnitudeModerate || !a.Viability.MatchesScalarPortalSign || !a.Viability.CanBeGeneratedByTreePortal || a.Viability.HeavySectorDerived {
		t.Fatalf("bad viability: %s", FormatViability(a.Viability))
	}
	if !a.Firewalls.NoObservedMassClaimedAsDerivation || !a.Firewalls.NoThresholdJumpDerived || !a.Firewalls.NoPortalCouplingFitted || !a.Firewalls.NoHeavySelfQuarticFitted || !a.Firewalls.NoThresholdScaleDerived || !a.Firewalls.NoTwoLoopRGExecuted || !a.Firewalls.NoPoleMassConversionInserted || !a.Firewalls.NoFinalMassClaimed || a.Firewalls.FiniteCorePolluted {
		t.Fatalf("firewall failed: %s", FormatFirewalls(a.Firewalls))
	}
	if !a.Summary.RequiredJumpExtracted || !a.Summary.JumpHasCorrectSign || !a.Summary.JumpHasPortalMagnitude || a.Summary.HeavySectorDerived || a.Summary.FinalMassClaimed || !a.Summary.FirewallPreserved {
		t.Fatalf("summary failed: %s", FormatSummary(a.Summary))
	}
}

func TestTheoremPasses(t *testing.T) {
	res := IntermediateThresholdDecouplingQuarticJumpTransportAuditTheorem().Run()
	if !res.Passed() {
		t.Fatalf("theorem did not pass:\n%s", res.Details())
	}
}
