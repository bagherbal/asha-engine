package singlescalematchingaudit

import "testing"

func TestGate215BuildDefault(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if a.Summary.Status != StatusConditionalPhenomenology {
		t.Fatalf("unexpected status: %s", FormatSummary(a.Summary))
	}
	if a.GlobalScan.ClassesAudited != 22 || len(a.Fits) != 22 {
		t.Fatalf("expected 22 classes: %s", FormatGlobal(a.GlobalScan))
	}
	if !a.Gate214.MatchingEnvelopeInherited || a.Config.EpsilonU <= 0 {
		t.Fatalf("missing Gate 214 envelope: gate214=%s config=%s", FormatGate214(a.Gate214), FormatConfig(a.Config))
	}
}

func TestBestDegenerateClassIsInsideEnvelope(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	best := a.Fits[0]
	if best.ClassRank != 1 {
		t.Fatalf("expected Gate-211 ranked witness to remain best under degenerate scan: %s", FormatFit(best))
	}
	if !best.MatchingPlausible || best.MaxAbsResidual >= a.Config.EpsilonU {
		t.Fatalf("best class should fit within loop-factor envelope: fit=%s epsilon=%.12g", FormatFit(best), a.Config.EpsilonU)
	}
	if !(best.MBGeV > 2.0e6 && best.MBGeV < 3.2e6) || !(best.MStarGeV > 1.0e17 && best.MStarGeV < 2.5e17) {
		t.Fatalf("unexpected degenerate scales: %s", FormatFit(best))
	}
	if best.ResidualOverEpsilon >= 0.2 {
		t.Fatalf("expected small residual relative to envelope, got: %s", FormatFit(best))
	}
}

func TestMatchingAndFirewallRemainSealed(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if a.MatchingAudit.NativeDeltaMatchRowsDerived || a.MatchingAudit.HeatKernelMatchingMapDerived || a.MatchingAudit.CanonicalSubtractionScheme || a.MatchingAudit.ResidualInterpretedAsDerived {
		t.Fatalf("matching firewall leak: %s", FormatMatching(a.MatchingAudit))
	}
	if a.Firewall.SingleScaleForcedAsFiniteCore || a.Firewall.MatchingCorrectionsDerived || a.Firewall.MatchingResidualPromoted || a.Firewall.YukawaMatricesImported || a.Firewall.PhysicalPredictionClaimed || a.Firewall.ProtonLifetimeComputed {
		t.Fatalf("firewall leak: %s", FormatFirewall(a.Firewall))
	}
}

func TestTheoremChecksPass(t *testing.T) {
	res := SingleScaleDegenerateLimitMatchingAuditTheorem().Verify()
	for _, c := range res.Checks {
		if !c.Passed {
			t.Fatalf("check %q failed: %s", c.Name, c.Detail)
		}
	}
}
