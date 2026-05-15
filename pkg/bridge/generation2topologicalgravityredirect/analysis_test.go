package generation2topologicalgravityredirect

import (
	"strings"
	"testing"
)

func TestGate509Build(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Inheritance.Gate508EWFirewallClosed || !a.Inheritance.Gate490AllAnomalyTracesCancel || !a.Inheritance.Gate377ProductTripleValid {
		t.Fatalf("bad inheritance: %+v", a.Inheritance)
	}
	if !a.Anomaly.GaugeStabilityLedgerSatisfied || !a.Anomaly.PerturbativeGaugeCancel || !a.Anomaly.MixedGaugeGravityCancel || !a.Anomaly.WittenSU2GlobalCancels || !a.Anomaly.BLTracesCancel {
		t.Fatalf("anomaly cancellation not established: %+v", a.Anomaly)
	}
	if a.Anomaly.DerivesYukawaTexture || a.Anomaly.DerivesCKMOrJarlskog {
		t.Fatalf("anomaly gate must not derive flavor: %+v", a.Anomaly)
	}
	if !a.Gravity.EinsteinHilbertSocketPresent || !a.Gravity.A2ScalarCurvatureChannel || !a.Gravity.SMGravityStructuralRecovered {
		t.Fatalf("gravity socket missing: %+v", a.Gravity)
	}
	if a.Gravity.RawEHCoefficientFullyPhysical || a.Gravity.SkeletonEHCoefficientPhysical || a.Gravity.AllCoefficientsDetermined || a.Gravity.HardTOEClosure {
		t.Fatalf("gravity socket overpromoted: %+v", a.Gravity)
	}
	if a.Firewall.NewtonConstantImported || a.Firewall.NewtonConstantDerived || a.Firewall.PlanckScaleImported || a.Firewall.CosmologicalScaleImported || a.Firewall.CutoffLambdaSelected || a.Firewall.F2MomentSeparatedFromLambda || a.Firewall.EinsteinHilbertNormalizationClosed || a.Firewall.CosmologicalConstantDerived || a.Firewall.NativeGravityRegistryWritten {
		t.Fatalf("gravity firewall violated: %+v", a.Firewall)
	}
	if a.Next.Gate != 510 {
		t.Fatalf("unexpected next gate %+v", a.Next)
	}
}

func TestGate509Markdown(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	md := Markdown(a)
	for _, want := range []string{
		"# Gate 509 Registry Audit",
		StatusNativeAnomalyCancellationReaffirmed,
		StatusEinsteinHilbertSocketStructurallyPresent,
		StatusFailedNewtonConstantNotDerived,
		StatusFirewallNativeGravityWriteBlocked,
		"Gate 510",
		"No Newton constant",
	} {
		if !strings.Contains(md, want) {
			t.Fatalf("markdown missing %q", want)
		}
	}
}

func TestGate509Theorem(t *testing.T) {
	result := Generation2TopologicalAnomaliesGravitationalSpectralRedirectTheorem().Verify()
	for _, c := range result.Checks {
		if !c.Passed {
			t.Fatalf("check failed: %s %s", c.Name, c.Detail)
		}
	}
}
