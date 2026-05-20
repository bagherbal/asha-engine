package generation2koidecoordinatebetafunctionaudit

import (
	"math"
	"testing"
)

func TestGate581KoideCoordinateBetaFunctionAudit(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if math.Abs(a.MZ.DThetaDTDeg-4.25133316927e-06) > 1e-12 {
		t.Fatalf("unexpected M_Z theta beta: %s", FormatEndpointBeta(a.MZ))
	}
	if math.Abs(a.MZ.DPhiDTDeg-6.98104646218e-06) > 1e-12 {
		t.Fatalf("unexpected M_Z phi beta: %s", FormatEndpointBeta(a.MZ))
	}
	if math.Abs(a.Lambda12.DThetaDTDeg-4.22880857135e-06) > 1e-12 {
		t.Fatalf("unexpected Lambda theta beta: %s", FormatEndpointBeta(a.Lambda12))
	}
	if !(a.MZ.PointsTowardCone && a.Lambda12.PointsTowardCone) {
		t.Fatalf("runtime beta should locally point toward the cone: %s / %s", FormatEndpointBeta(a.MZ), FormatEndpointBeta(a.Lambda12))
	}
	if !(a.MZ.CommonOnlyProjectiveSpeedRad < 1e-18 && a.Lambda12.CommonOnlyProjectiveSpeedRad < 1e-18) {
		t.Fatalf("common rate should cancel projective motion: %s / %s", FormatEndpointBeta(a.MZ), FormatEndpointBeta(a.Lambda12))
	}
	if a.Cone.ConeInvariantInV1 || a.Cone.AttractorCertified {
		t.Fatalf("v1 should not certify cone invariant/attractor theorem: %s", FormatCone(a.Cone))
	}
	if !(a.Cone.MZExactConeDThetaDTDeg > 0 && a.Cone.LambdaExactConeDThetaDTDeg > 0) {
		t.Fatalf("exact cone theta beta should be nonzero in v1: %s", FormatCone(a.Cone))
	}
	if a.Firewalls.DerivesKoide || a.Firewalls.IntroducesNewCarrier || a.Firewalls.PromotesObservedAsNative || !a.Firewalls.PreservesGate352 {
		t.Fatalf("firewall broken: %s", FormatFirewalls(a.Firewalls))
	}
}

func TestGate581Theorem(t *testing.T) {
	res := Generation2KoideCoordinateBetaFunctionAuditTheorem().Run()
	if res.Status == "FAILED_ROUTE" {
		t.Fatalf("unexpected failed route: %+v", res)
	}
	for _, c := range res.Checks {
		if !c.Passed {
			t.Fatalf("check failed: %+v", c)
		}
	}
}
