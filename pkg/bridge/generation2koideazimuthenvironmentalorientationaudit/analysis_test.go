package generation2koideazimuthenvironmentalorientationaudit

import "testing"

func TestGate578KoideAzimuthAudit(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if abs(a.Frame.DotNE1) > 1e-12 || abs(a.Frame.DotNE2) > 1e-12 || abs(a.Frame.DotE1E2) > 1e-12 || !a.Frame.RightHanded {
		t.Fatalf("bad frame: %s", FormatFrame(a.Frame))
	}
	if abs(a.Transport.MZ.PhiDeg-257.2671800328917) > 1e-6 {
		t.Fatalf("unexpected MZ azimuth: %s", FormatAzimuthPoint(a.Transport.MZ))
	}
	if abs(a.Transport.Lambda12.PhiDeg-257.267382531545) > 1e-6 {
		t.Fatalf("unexpected Lambda12 azimuth: %s", FormatAzimuthPoint(a.Transport.Lambda12))
	}
	if !a.Transport.StableAt1eMinus3Deg || a.Transport.AbsDeltaPhiDeg > 3e-4 {
		t.Fatalf("azimuth should be stable in v1: %s", FormatTransport(a.Transport))
	}
	if a.Candidates.AnyCertified || a.Candidates.NearestRationalTurn != "5/7" || a.Candidates.NearestRationalDistanceDeg <= a.Candidates.CertificationThresholdDeg {
		t.Fatalf("candidate audit too permissive: %s", FormatCandidates(a.Candidates))
	}
	if a.Seal.NativeDerivation || !a.Seal.BridgeOnly || a.Seal.RemainingContinuousCoordinates != 2 {
		t.Fatalf("bad seal: %s", FormatSeal(a.Seal))
	}
	if a.Firewalls.DerivesCKM || a.Firewalls.DerivesPMNS || a.Firewalls.IdentifiesWithASHAProjectivePhase || a.Firewalls.ImportsObservedAsNative || !a.Firewalls.PreservesGate352 {
		t.Fatalf("firewall broken: %s", FormatFirewalls(a.Firewalls))
	}
}

func TestGate578CandidateThreshold(t *testing.T) {
	c := candidate("test", "270 degrees", 270, 257.2671800328917, 0.03)
	if c.Certified || c.DistanceDeg < 12 {
		t.Fatalf("270-degree candidate should fail: %s", FormatCandidate(c))
	}
	close := candidate("test", "near observed", 257.268, 257.2671800328917, 0.03)
	if !close.Certified {
		t.Fatalf("close candidate should pass threshold: %s", FormatCandidate(close))
	}
}

func TestGate578Theorem(t *testing.T) {
	res := Generation2KoideAzimuthEnvironmentalOrientationAuditTheorem().Run()
	if res.Status == "FAILED_ROUTE" {
		t.Fatalf("unexpected failed route: %+v", res)
	}
	for _, c := range res.Checks {
		if !c.Passed {
			t.Fatalf("check failed: %+v", c)
		}
	}
}
