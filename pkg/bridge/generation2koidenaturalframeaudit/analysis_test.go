package generation2koidenaturalframeaudit

import "testing"

func TestGate579KoideNaturalFrameAudit(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if abs(a.Frame.DotNE1) > 1e-12 || abs(a.Frame.DotNE2) > 1e-12 || abs(a.Frame.DotE1E2) > 1e-12 || !a.Frame.RightHanded {
		t.Fatalf("bad frame: %s", FormatFrame(a.Frame))
	}
	if abs(a.Compare.Pole.PhiDeg-257.2671800328917) > 1e-6 {
		t.Fatalf("unexpected pole azimuth: %s", FormatPoint(a.Compare.Pole))
	}
	if abs(a.Compare.MZ.PhiDeg-a.Compare.Pole.PhiDeg) > 1e-12 || abs(a.Compare.MZ.Q-a.Compare.Pole.Q) > 1e-14 {
		t.Fatalf("M_Z should be angle-equivalent to pole in v1: %s", FormatComparison(a.Compare))
	}
	if abs(a.Compare.Lambda12.PhiDeg-257.267382531545) > 1e-6 {
		t.Fatalf("unexpected Lambda12 azimuth: %s", FormatPoint(a.Compare.Lambda12))
	}
	if !a.Compare.LambdaCloserThanMZ || !a.Compare.AzimuthStable {
		t.Fatalf("expected Lambda12 closer and azimuth stable: %s", FormatComparison(a.Compare))
	}
	if a.Natural.BoundaryFrameCertified || a.Natural.MZYukawaFrameIndependent {
		t.Fatalf("natural frame should not be certified by v1: %s", FormatNatural(a.Natural))
	}
	if a.Firewalls.DerivesPMNS || a.Firewalls.DerivesCKM || a.Firewalls.PromotesObservedAsNative || !a.Firewalls.PreservesGate352 {
		t.Fatalf("firewall broken: %s", FormatFirewalls(a.Firewalls))
	}
}

func TestGate579Theorem(t *testing.T) {
	res := Generation2KoideNaturalFrameAuditTheorem().Run()
	if res.Status == "FAILED_ROUTE" {
		t.Fatalf("unexpected failed route: %+v", res)
	}
	for _, c := range res.Checks {
		if !c.Passed {
			t.Fatalf("check failed: %+v", c)
		}
	}
}
