package scalarquarticchannel

import "testing"

func TestGate305Inheritance(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault returned error: %v", err)
	}
	if !a.Input.Gate304F0Promoted || a.Input.PromotedF0Value != 7 || !a.Input.F0Positive || !a.Input.ScalarSubtractionFormalized || !a.Input.MassMapFormalized {
		t.Fatalf("bad Gate 305 inheritance: %s", FormatGate305Inheritance(a.Input))
	}
	if a.Input.F2MomentLocked || a.Input.HiggsMassPredictionClaimed || a.Input.QuarticChannelAlreadyTouched || a.Input.NumericalZHComputed || a.Input.NumericalYukawasInserted {
		t.Fatalf("Gate 306 inherited overclaimed state: %s", FormatGate305Inheritance(a.Input))
	}
}

func TestRawA4QuarticDecomposition(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault returned error: %v", err)
	}
	if !a.A4.DecompositionFormalized || !a.A4.A4SourceConfirmed || !a.A4.ScalarPower4ChannelSeen || !a.A4.DerivativeTermsRejected || !a.A4.GaugeTermsRejected || !a.A4.VacuumTermsRejected || a.A4.NumericalCoefficientUsed || len(a.A4.Components) < 4 {
		t.Fatalf("bad a4 quartic decomposition: %s", FormatRawA4(a.A4))
	}
	var quartic, rejectedKinetic, rejectedGauge bool
	for _, c := range a.A4.Components {
		if c.AcceptedForQuartic && c.ScalarPower == 4 && c.DerivativeOrder == 0 && c.GaugeCurvaturePower == 0 {
			quartic = true
		}
		if c.RejectedFromQuartic && c.DerivativeOrder == 2 {
			rejectedKinetic = true
		}
		if c.RejectedFromQuartic && c.GaugeCurvaturePower == 2 {
			rejectedGauge = true
		}
	}
	if !quartic || !rejectedKinetic || !rejectedGauge {
		t.Fatalf("a4 components did not isolate/reject correctly: %s", FormatRawA4(a.A4))
	}
}

func TestQuarticCouplingNormalization(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault returned error: %v", err)
	}
	if !a.Quartic.MapFormalized || !a.Quartic.UsesGate300ZHNormalization || !a.Quartic.UsesGate304F0Seal || !a.Quartic.RequiresPositiveZH || !a.Quartic.RequiresRawC4Carrier || !a.Quartic.RequiresYukawaAmplitudes || !a.Quartic.RequiresSignConvention {
		t.Fatalf("quartic map missing obligations: %s", FormatQuartic(a.Quartic))
	}
	if a.Quartic.NumericalLambdaComputed {
		t.Fatalf("Gate 306 must not compute numerical lambda_H: %s", FormatQuartic(a.Quartic))
	}
}

func TestF0DependencyAudit(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault returned error: %v", err)
	}
	if !a.F0.AuditFormalized || a.F0.F0Value != 7 || a.F0.F0CancelsInsideLambdaAlone || !a.F0.F0CancelsInLambdaOverGauge || !a.F0.RetainsN4F0ForAbsoluteLambda || a.F0.F2RequiredForQuartic {
		t.Fatalf("bad f0 dependency audit: %s", FormatF0(a.F0))
	}
}

func TestDimensionlessRatioSynthesis(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault returned error: %v", err)
	}
	if !a.Ratio.SynthesisFormalized || !a.Ratio.RelativeRatioCanCancelN4F0 || a.Ratio.RawRatioPromotedDirectly || !a.Ratio.NeedsC4Raw || !a.Ratio.NeedsKHRaw || !a.Ratio.NeedsTraceIndex || !a.Ratio.NeedsYukawaAmplitudeSeal || !a.Ratio.NeedsQuarticSignConvention || !a.Ratio.NeedsAbsoluteGaugeNormalization || a.Ratio.NumericalPhysicalPredictionMade {
		t.Fatalf("bad dimensionless ratio synthesis: %s", FormatRatio(a.Ratio))
	}
	if a.Ratio.RawTraceRatioNumerator != 1197 || a.Ratio.RawTraceRatioDenominator != 4624 {
		t.Fatalf("raw ratio ledger corrupted: %s", FormatRatio(a.Ratio))
	}
}

func TestChannelLedger(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault returned error: %v", err)
	}
	if !a.Channels.A4QuarticIsolated || !a.Channels.A4KineticPreserved || !a.Channels.A4GaugePreserved || !a.Channels.A2MassChannelUndisturbed || !a.Channels.F0SealUsedForA4 || !a.Channels.F2NotUsedForQuartic || !a.Channels.NoHiggsMassClaimed || !a.Channels.NoNumericalQuarticClaimed {
		t.Fatalf("bad channel ledger: %s", FormatChannels(a.Channels))
	}
}

func TestFirewalls(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault returned error: %v", err)
	}
	if !a.Firewalls.NoNumericalC4Inserted || !a.Firewalls.NoYukawaNumbersInserted || !a.Firewalls.NoNumericalZHComputed || !a.Firewalls.NoNumericalLambdaHComputed || !a.Firewalls.NoRaw1197PromotedDirectly || !a.Firewalls.NoHiggsMassPredictionClaimed || !a.Firewalls.NoAbsoluteGaugeCouplingsClaimed || !a.Firewalls.NoBGapInstantonClaimed || !a.Firewalls.F2FirewallPreserved || !a.Firewalls.F0SealPreservedForA4 || a.Firewalls.FiniteCorePolluted {
		t.Fatalf("firewall failure: %s", FormatFirewalls(a.Firewalls))
	}
}

func TestTheoremPassesWithBridgeRequiredStatus(t *testing.T) {
	res := ScalarQuarticChannelExtractionDimensionlessCouplingSieveAuditTheorem().Run()
	if !res.Passed() {
		t.Fatalf("theorem did not pass:\n%s", res.Details())
	}
}
