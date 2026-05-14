package sealedthresholdstresstest

import "testing"

func TestBuildDefaultStressTestPassesWithFirewalls(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault error: %v", err)
	}
	if a.Summary.Status != StatusConditionalStressTest {
		t.Fatalf("status = %q", a.Summary.Status)
	}
	if !a.Collider.ColliderStressPassed || a.Collider.CasesAudited != 2 {
		t.Fatalf("collider stress failed: %+v", a.Collider)
	}
	if !a.ProtonDecay.NaiveSU5DimensionSixWarning {
		t.Fatalf("expected naive SU(5) proton-decay warning")
	}
	if a.ProtonDecay.XYMediatedChannelSupported {
		t.Fatalf("engine should not support X/Y mediated channel")
	}
	if !a.ProtonDecay.NaturalSuppressionByMediatorAbsence {
		t.Fatalf("expected mediator-absence suppression")
	}
	if a.UniversalCompletion.NoOneLoopPathologyBelowPlanck {
		t.Fatalf("expected universal completion pathology to be detected: %+v", a.UniversalCompletion)
	}
	if a.Firewall.AbsoluteMassPredicted || a.Firewall.PhysicalUnificationClaimed || a.Firewall.UniversalBetaSourceDerived {
		t.Fatalf("firewall leak: %+v", a.Firewall)
	}
}

func TestColliderScaleSeparationIsPeVNotTeV(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault error: %v", err)
	}
	if a.Collider.MinimumThresholdScaleTeV < 1000 {
		t.Fatalf("minimum M_B should be PeV-scale in TeV units, got %.9g", a.Collider.MinimumThresholdScaleTeV)
	}
	for _, c := range a.ColliderCases {
		if c.SeparationFromCurrentDirectLimit < 100 {
			t.Fatalf("direct-limit separation too small for %s: %.9g", c.CarrierName, c.SeparationFromCurrentDirectLimit)
		}
	}
}

func TestUniversalLandauPolePathologyIsDetected(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault error: %v", err)
	}
	for _, c := range a.UniversalCases {
		if c.U1LandauPoleAbovePlanck || c.U1LandauPoleGeV >= planckGeV {
			t.Fatalf("expected sub-Planck U1 pole for %s, got %.9g", c.CarrierName, c.U1LandauPoleGeV)
		}
		if c.SU2AsymptoticallyFreeAboveThreshold || c.SU3AsymptoticallyFreeAboveThreshold {
			t.Fatalf("expected non-Abelian asymptotic-safety obstruction for %s: %+v", c.CarrierName, c)
		}
	}
}

func TestTheoremVerifierRecordsFailedRouteWithoutFailedChecks(t *testing.T) {
	res := SealedThresholdPredictionStressTestTheorem().Run()
	if string(res.Status) != "FAILED_ROUTE" {
		t.Fatalf("status = %s", res.Status)
	}
	for _, chk := range res.Checks {
		if !chk.Passed {
			t.Fatalf("check %q failed: %s", chk.Name, chk.Detail)
		}
	}
}
