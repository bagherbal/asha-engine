package twoloopintegration

import "testing"

func TestGate214BuildDefault(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if a.Summary.Status != StatusConditionalPhenomenology {
		t.Fatalf("unexpected status: %s", FormatSummary(a.Summary))
	}
	if !a.Central.Converged || a.Central.ResidualNorm > 1e-8 {
		t.Fatalf("central solve did not converge: %s", FormatSolution(a.Central))
	}
	if a.Central.LStar <= a.Gate213.OneLoopLStar {
		t.Fatalf("expected two-loop boundary to shift upward from one-loop reference: central=%s gate213=%s", FormatSolution(a.Central), FormatGate213(a.Gate213))
	}
}

func TestTwoLoopCorrectedScaleValues(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !near(a.Central.LB1, 10.309809, 2e-4) || !near(a.Central.LB2, 10.259942, 2e-4) || !near(a.Central.LStar, 35.187539, 2e-4) {
		t.Fatalf("unexpected corrected logs: %s", FormatSolution(a.Central))
	}
	if !(a.Central.MB1GeV > 2.6e6 && a.Central.MB1GeV < 2.9e6) || !(a.Central.MB2GeV > 2.4e6 && a.Central.MB2GeV < 2.8e6) || !(a.Central.MStarGeV > 1.6e17 && a.Central.MStarGeV < 1.9e17) {
		t.Fatalf("unexpected corrected scales: %s", FormatSolution(a.Central))
	}
	if a.Central.LB2 >= a.Central.LB1 {
		t.Fatalf("expected threshold order flip after two-loop correction: %s", FormatSolution(a.Central))
	}
}

func TestMatchingEnvelopeAndFirewalls(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if a.Envelope.CasesAudited != 8 || a.Envelope.ConvergedCases != 8 || a.Envelope.Status != MatchingEnvelopeStatus {
		t.Fatalf("bad envelope: %s", FormatEnvelope(a.Envelope))
	}
	if a.Envelope.MB1MinGeV >= a.Central.MB1GeV || a.Envelope.MB1MaxGeV <= a.Central.MB1GeV || a.Envelope.MStarMinGeV >= a.Central.MStarGeV || a.Envelope.MStarMaxGeV <= a.Central.MStarGeV {
		t.Fatalf("envelope does not contain central solution: central=%s envelope=%s", FormatSolution(a.Central), FormatEnvelope(a.Envelope))
	}
	if a.MatchingAudit.NativeDeltaMatchRowsDerived || a.MatchingAudit.CanonicalSubtractionSchemeDerived || a.MatchingAudit.EnvelopeImportedAsFiniteCore {
		t.Fatalf("matching firewall leak: %s", FormatMatching(a.MatchingAudit))
	}
	if a.Firewall.PhysicalPredictionClaimed || a.Firewall.FiniteMassPredictionClaimed || a.Firewall.YukawaMatricesImported || a.Firewall.MatchingCorrectionsDerived {
		t.Fatalf("firewall leak: %s", FormatFirewall(a.Firewall))
	}
}

func TestTheoremChecksPass(t *testing.T) {
	res := SealedTwoLoopRGIntegrationMatchingEnvelopeTheorem().Verify()
	for _, c := range res.Checks {
		if !c.Passed {
			t.Fatalf("check %q failed: %s", c.Name, c.Detail)
		}
	}
}

func near(got, want, tol float64) bool {
	if got < want-tol || got > want+tol {
		return false
	}
	return true
}
