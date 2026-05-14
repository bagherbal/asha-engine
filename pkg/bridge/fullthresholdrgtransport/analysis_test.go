package fullthresholdrgtransport

import "testing"

func TestProtocolFormalized(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault returned error: %v", err)
	}
	if !a.Protocol.Formalized || a.Protocol.LaneName != "gauge_only_zero_top_lower_envelope" || a.Protocol.TopYukawaUV != 0 || !a.Protocol.OneLoopOnly {
		t.Fatalf("bad protocol: %s", FormatProtocol(a.Protocol))
	}
}

func TestDerivedThresholdInserted(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault returned error: %v", err)
	}
	if !a.Insertion.Inserted || a.Insertion.DeltaLambda > -0.097 || a.Insertion.DeltaLambda < -0.098 || !a.Insertion.LowersQuartic || a.Insertion.DerivedAsFullPotential {
		t.Fatalf("bad insertion: %s", FormatInsertion(a.Insertion))
	}
}

func TestTransportNearObservedRunningProxy(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault returned error: %v", err)
	}
	if !a.Transport.Computed || !a.Transport.Perturbative || !a.Transport.VacuumStableAtEndpoint {
		t.Fatalf("transport failed: %s", FormatTransport(a.Transport))
	}
	if a.Transport.BaselineMassGeV < 158.0 || a.Transport.BaselineMassGeV > 159.0 {
		t.Fatalf("unexpected baseline mass: %s", FormatTransport(a.Transport))
	}
	if a.Transport.RunningMassGeV < 124.0 || a.Transport.RunningMassGeV > 126.0 || !a.Transport.NearObservedWithinOnePct {
		t.Fatalf("unexpected threshold mass: %s", FormatTransport(a.Transport))
	}
}

func TestPrecisionFirewalls(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault returned error: %v", err)
	}
	if !a.Precision.Formalized || !a.Precision.RunningMassNotPoleMass || !a.Precision.TwoLoopRequired || !a.Precision.PoleMatchingRequired || !a.Precision.ExactThresholdScaleRequired {
		t.Fatalf("bad precision gap: %s", FormatPrecision(a.Precision))
	}
	if !a.Firewalls.NoPoleMassClaimed || !a.Firewalls.NoTwoLoopClaimed || !a.Firewalls.NoExactThresholdScaleClaimed || !a.Firewalls.NoPhysicalTopSectorClaimed || !a.Firewalls.NoFullSigmaPotentialClaimed || !a.Firewalls.NoFinalColliderClaimed || a.Firewalls.FiniteCorePolluted {
		t.Fatalf("bad firewalls: %s", FormatFirewalls(a.Firewalls))
	}
}

func TestTheoremPasses(t *testing.T) {
	res := FullThresholdRGTransportConditionalHiggsMassPredictionAuditTheorem().Run()
	if !res.Passed() {
		t.Fatalf("theorem did not pass:\n%s", res.Details())
	}
}
