package trialitygenerationpullback

import (
	"math"
	"testing"
)

func TestCarrierAndFractionalization(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault returned error: %v", err)
	}
	if !a.Carrier.Formalized || len(a.Carrier.TauEta) != 3 || !a.Carrier.GenerationBreaking || !a.Carrier.MagnitudeSpectrumDegenerate || a.Carrier.PullbackUnique {
		t.Fatalf("bad carrier: %s", FormatCarrier(a.Carrier))
	}
	if !a.Fractional.Formalized || len(a.Fractional.NormalizedWeights) != 3 || math.Abs(a.Fractional.SumWeights-1) > 1e-12 || a.Fractional.RPlusDecimal < 1.6 || a.Fractional.RPlusDecimal > 1.7 {
		t.Fatalf("bad fractionalization: %s", FormatFractionalization(a.Fractional))
	}
	if math.Abs(a.Fractional.NormalizedWeights[0]-4.0/9.0) > 1e-12 || math.Abs(a.Fractional.NormalizedWeights[1]-4.0/9.0) > 1e-12 || math.Abs(a.Fractional.NormalizedWeights[2]-1.0/9.0) > 1e-12 {
		t.Fatalf("unexpected weights: %s", FormatFractionalization(a.Fractional))
	}
}

func TestCandidatesRemainNonCanonical(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault returned error: %v", err)
	}
	if len(a.Candidates) != 4 {
		t.Fatalf("expected four candidates, got %d", len(a.Candidates))
	}
	for i := 0; i < 3; i++ {
		if !a.Candidates[i].DerivedFromTauEta || a.Candidates[i].Canonical || !a.Candidates[i].Ambiguous {
			t.Fatalf("tau_eta candidate should remain ambiguous and noncanonical: %s", FormatCandidate(a.Candidates[i]))
		}
	}
	if a.Candidates[3].DerivedFromTauEta || a.Candidates[3].Weight != 0 || a.Candidates[3].TopYtUV != 0 {
		t.Fatalf("gauge-only candidate should be diagnostic zero top: %s", FormatCandidate(a.Candidates[3]))
	}
}

func TestThresholdTransportPreflight(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault returned error: %v", err)
	}
	if len(a.Preflights) != len(a.Candidates) {
		t.Fatalf("preflight/candidate mismatch")
	}
	for _, p := range a.Preflights {
		if !p.Perturbative || !p.VacuumStable || math.IsNaN(p.RunningMassGeV) {
			t.Fatalf("bad preflight: %s", FormatPreflight(p))
		}
	}
	high := findPreflight(a.Preflights, "tau_eta_positive_high_slot")
	low := findPreflight(a.Preflights, "tau_eta_unique_low_slot")
	gauge := findPreflight(a.Preflights, "gauge_only_zero_top_envelope")
	if high == nil || low == nil || gauge == nil {
		t.Fatalf("missing preflight: %#v", a.Preflights)
	}
	if !(high.RunningMassGeV > low.RunningMassGeV && low.RunningMassGeV > gauge.RunningMassGeV) {
		t.Fatalf("expected high > low > gauge masses, got high=%s low=%s gauge=%s", FormatPreflight(*high), FormatPreflight(*low), FormatPreflight(*gauge))
	}
	if !gauge.Near125WithinOnePct || high.Near125WithinOnePct || low.Near125WithinOnePct {
		t.Fatalf("only gauge envelope should remain near 125: high=%s low=%s gauge=%s", FormatPreflight(*high), FormatPreflight(*low), FormatPreflight(*gauge))
	}
	if math.Abs(gauge.RunningMassGeV-124.9766199157) > 1e-6 {
		t.Fatalf("unexpected gauge mass: %s", FormatPreflight(*gauge))
	}
}

func TestPullbackVerdictAndFirewalls(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault returned error: %v", err)
	}
	if !a.PullbackVerdict.Formalized || a.PullbackVerdict.CanonicalTopFractionDerived || !a.PullbackVerdict.GaugeOnlyStillRequired || !a.PullbackVerdict.NonzeroTauEtaTopSpoils125 || a.PullbackVerdict.TopBoundaryStatus != StatusFailedNativeTopBoundaryNotDerived {
		t.Fatalf("bad pullback verdict: %s", FormatPullbackVerdict(a.PullbackVerdict))
	}
	if !a.Firewalls.NoObservedTopMassInserted || !a.Firewalls.NoCKMImported || !a.Firewalls.NoFlavorTextureInvented || !a.Firewalls.NoPoleMassClaimed || !a.Firewalls.NoTwoLoopClaimed || !a.Firewalls.NoFinalColliderMassClaimed || a.Firewalls.FiniteCorePolluted {
		t.Fatalf("firewall failed: %s", FormatFirewalls(a.Firewalls))
	}
	if !a.Summary.TrialityPullbackFormalized || !a.Summary.FractionalizationExtracted || !a.Summary.TopCandidatesAudited || a.Summary.NativeTopBoundaryDerived || !a.Summary.PhysicalPreflightExecuted || a.Summary.Gate322SuccessPreservedByCanonicalTop || !a.Summary.FirewallsPreserved || a.Summary.FinalMassClaimed {
		t.Fatalf("summary failed: %s", FormatSummary(a.Summary))
	}
}

func TestTheoremPasses(t *testing.T) {
	res := TrialityGenerationPullbackNativeTopYukawaBoundarySieveAuditTheorem().Run()
	if !res.Passed() {
		t.Fatalf("theorem did not pass:\n%s", res.Details())
	}
}
