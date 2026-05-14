package higgsinverseshapeprecision

import "testing"

func TestBuildDefault(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault returned error: %v", err)
	}
	if a.Inputs.HighestInheritedGate != inheritedHighestGate || a.Inputs.NativeShape.RatString() != "1197/4624" {
		t.Fatalf("bad inputs: %s", FormatInputs(a.Inputs))
	}
}

func TestExactInverseObservedShape(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault returned error: %v", err)
	}
	if a.Inverse.ObservedShape.RatString() != "39125025/151560721" {
		t.Fatalf("bad observed shape: %s", FormatInverse(a.Inverse))
	}
	if !nearlyFloat(a.Inverse.ObservedShapeDec, 0.25814752491181403, 1e-15) {
		t.Fatalf("bad observed shape decimal: %s", FormatInverse(a.Inverse))
	}
}

func TestExactDeviation(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault returned error: %v", err)
	}
	if a.Deviation.ShapeDelta.RatString() != "504067437/700816773904" {
		t.Fatalf("bad shape delta: %s", FormatDeviation(a.Deviation))
	}
	if a.Deviation.LambdaDelta.RatString() != "504067437/1401633547808" {
		t.Fatalf("bad lambda delta: %s", FormatDeviation(a.Deviation))
	}
	if !nearlyFloat(a.Deviation.MassDeltaGeV, 0.17415714969897194, 1e-13) {
		t.Fatalf("bad mass delta: %s", FormatDeviation(a.Deviation))
	}
}

func TestSelfEnergyAndRequiredVEV(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault returned error: %v", err)
	}
	if a.SelfEnergy.RequiredRePiGeV2.RatString() != "504067437/11560000" {
		t.Fatalf("bad RePi target: %s", FormatSelfEnergy(a.SelfEnergy))
	}
	if !nearlyFloat(a.RequiredVEV.VRequiredForTargetGeV, 245.87770295825946, 1e-12) {
		t.Fatalf("bad required VEV: %s", FormatRequiredVEV(a.RequiredVEV))
	}
}

func TestFirewallsAndStatuses(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault returned error: %v", err)
	}
	if !a.Firewalls.NoPoleCorrection || !a.Firewalls.NoFittingShape || !a.Firewalls.NoColliderClaim || a.Ledger.UsesFloat64Core {
		t.Fatalf("bad firewalls: %s %s", FormatFirewalls(a.Firewalls), FormatLedger(a.Ledger))
	}
	statuses := Statuses(a)
	required := []string{StatusExactInverseShapeComputed, StatusExactDeviationComputed, StatusFailedPoleCorrectionNotComputed, StatusFailedNativeShapeNotAltered, StatusFailedColliderMassNotClaimed}
	for _, req := range required {
		found := false
		for _, got := range statuses {
			if got == req {
				found = true
				break
			}
		}
		if !found {
			t.Fatalf("missing status %s in %v", req, statuses)
		}
	}
}

func TestTheoremPasses(t *testing.T) {
	res := ExactInverseHiggsShapeDeviationPrecisionAuditTheorem().Run()
	if !res.Passed() {
		t.Fatalf("theorem did not pass:\n%s", res.Details())
	}
}
