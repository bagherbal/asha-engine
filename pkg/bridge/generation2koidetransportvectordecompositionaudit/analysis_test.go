package generation2koidetransportvectordecompositionaudit

import "testing"

func TestGate580KoideTransportVectorDecompositionAudit(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if abs(a.MZ.PhiDeg-257.2671800328917) > 1e-6 {
		t.Fatalf("unexpected M_Z azimuth: %s", FormatEndpoint(a.MZ))
	}
	if abs(a.Lambda12.PhiDeg-257.267382531545) > 1e-6 {
		t.Fatalf("unexpected Lambda_12 azimuth: %s", FormatEndpoint(a.Lambda12))
	}
	if abs(a.Transport.DeltaPhiDeg-0.000202498653266) > 1e-9 {
		t.Fatalf("unexpected phi drift: %s", FormatTransport(a.Transport))
	}
	if !(a.Transport.DeltaThetaDeg > 0 && a.Transport.MovesTowardCone) {
		t.Fatalf("theta should move toward cone: %s", FormatTransport(a.Transport))
	}
	if !(a.Transport.RadialDominant && a.Transport.RadialToProjectiveRatio > 100) {
		t.Fatalf("expected radial dominance: %s", FormatTransport(a.Transport))
	}
	if a.Dynamics.ConeAttractorCertified || a.Dynamics.ContinuousBetaCertified {
		t.Fatalf("attractor/beta theorem should not be certified: %s", FormatDynamics(a.Dynamics))
	}
	if a.Firewalls.DerivesKoide || a.Firewalls.IntroducesNewCarrier || a.Firewalls.PromotesObservedAsNative || !a.Firewalls.PreservesGate352 {
		t.Fatalf("firewall broken: %s", FormatFirewalls(a.Firewalls))
	}
}

func TestGate580Theorem(t *testing.T) {
	res := Generation2KoideTransportVectorDecompositionAuditTheorem().Run()
	if res.Status == "FAILED_ROUTE" {
		t.Fatalf("unexpected failed route: %+v", res)
	}
	for _, c := range res.Checks {
		if !c.Passed {
			t.Fatalf("check failed: %+v", c)
		}
	}
}
