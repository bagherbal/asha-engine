package generation2gaugescalarboundaryresidualpairingaudit

import "testing"

func nearly(a, b, tol float64) bool {
	if a > b {
		return a-b < tol
	}
	return b-a < tol
}

func TestBuildGate611(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if a.InheritedGauge.Verdict != StatusGate610Inherited {
		t.Fatalf("bad gauge verdict %q", a.InheritedGauge.Verdict)
	}
	if a.InheritedScalar.Verdict != StatusGate606ScalarInherited {
		t.Fatalf("bad scalar verdict %q", a.InheritedScalar.Verdict)
	}
	if !nearly(a.InheritedGauge.R3MinusOne, 0.0509933868964996, 1e-12) {
		t.Fatalf("bad R3-1 %.15g", a.InheritedGauge.R3MinusOne)
	}
	if !nearly(a.InheritedScalar.AbsLambdaLambda12, 0.049700942077683274, 1e-12) {
		t.Fatalf("bad abs lambda %.15g", a.InheritedScalar.AbsLambdaLambda12)
	}
	if !(a.ResidualComparison.RatioAOverB > 1.02 && a.ResidualComparison.RatioAOverB < 1.03) {
		t.Fatalf("bad ratio %.15g", a.ResidualComparison.RatioAOverB)
	}
	if !nearly(a.CoefficientComparisons[0].Value, 2*a.InheritedScalar.AbsLambdaLambda12, 1e-15) {
		t.Fatalf("bad 2lambda row %.15g", a.CoefficientComparisons[0].Value)
	}
	if !a.SignCompatibility[0].PositiveShift || !a.SignCompatibility[1].PositiveShift {
		t.Fatalf("expected positive shifts: %+v", a.SignCompatibility)
	}
	if a.JointVector.CertifiedRelation {
		t.Fatalf("joint vector must not certify relation")
	}
	if !a.SensitivityLedger.ScalarMoreSensitive || a.SensitivityLedger.ClosureCertified {
		t.Fatalf("bad sensitivity ledger: %+v", a.SensitivityLedger)
	}
	if a.NativeStatus.ProvesGaugeScalarThresholdTheorem || a.NativeStatus.ProvesHiggsStabilityTheorem || a.NativeStatus.ClaimsHiggsMassPrediction {
		t.Fatalf("native firewall broken: %+v", a.NativeStatus)
	}
	if a.Firewalls.ClaimsGaugeUnification || a.Firewalls.ClaimsHiggsMassPredicted || a.Firewalls.ClaimsLambdaZeroBoundaryDerived {
		t.Fatalf("firewall broken: %+v", a.Firewalls)
	}
}

func TestTheoremGate611(t *testing.T) {
	res := Generation2GaugeScalarBoundaryResidualPairingAuditTheorem().Verify()
	if res.Status == "FAILED_ROUTE" {
		t.Fatalf("unexpected failed theorem: %+v", res)
	}
	for _, c := range res.Checks {
		if !c.Passed {
			t.Fatalf("check failed: %s %s", c.Name, c.Detail)
		}
	}
}
