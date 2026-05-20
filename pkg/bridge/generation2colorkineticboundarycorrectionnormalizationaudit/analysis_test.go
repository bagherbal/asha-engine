package generation2colorkineticboundarycorrectionnormalizationaudit

import "testing"

func nearly(a, b, tol float64) bool {
	if a > b {
		return a-b < tol
	}
	return b-a < tol
}

func TestBuildGate610(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if a.Inherited.Verdict != StatusGate609Inherited {
		t.Fatalf("bad inherited verdict %q", a.Inherited.Verdict)
	}
	if !nearly(a.Inherited.Delta3Required, 0.32739043299998416, 1e-12) {
		t.Fatalf("bad delta %.15g", a.Inherited.Delta3Required)
	}
	if !nearly(a.FractionalCorrection.EtaAgainstUStar, 0.0946843389411641, 1e-12) {
		t.Fatalf("bad etaStar %.15g", a.FractionalCorrection.EtaAgainstUStar)
	}
	if !(a.FractionalCorrection.EtaAgainstU3 > 0.10 && a.FractionalCorrection.EtaAgainstU3 < 0.11) {
		t.Fatalf("bad etaRuntime %.15g", a.FractionalCorrection.EtaAgainstU3)
	}
	if a.GaugeCoefficientAudit.Native || a.GaugeCoefficientAudit.Certified {
		t.Fatalf("color coefficient correction must not be native/certified")
	}
	if a.FSAStatus.HasIndependentColorKineticCoefficient || a.FSAStatus.HasSectorSplitF0Moment || a.FSAStatus.HasSU3OnlyBoundaryCorrection {
		t.Fatalf("unexpected native FSA color correction: %+v", a.FSAStatus)
	}
	if a.Firewalls.ClaimsCorrectionExists || a.Firewalls.ClaimsGaugeUnification || a.Firewalls.AltersFiniteAlgebra || a.Firewalls.AddsNewColoredStates || a.Firewalls.DerivesEndpoint {
		t.Fatalf("firewall broken: %+v", a.Firewalls)
	}
}

func TestTheoremGate610(t *testing.T) {
	res := Generation2ColorKineticBoundaryCorrectionNormalizationAuditTheorem().Verify()
	if res.Status == "FAILED_ROUTE" {
		t.Fatalf("unexpected failed theorem: %+v", res)
	}
	for _, c := range res.Checks {
		if !c.Passed {
			t.Fatalf("check failed: %s %s", c.Name, c.Detail)
		}
	}
}
