package generation2spectralmomenthierarchyairlock

import (
	"strings"
	"testing"
)

func TestGate513Build(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Inheritance.Gate512Inherited || !a.Inheritance.Gate510A2Inherited || !a.Inheritance.Gate511A4Inherited || a.Inheritance.ProductAllCoefficientsClosed {
		t.Fatalf("bad inheritance: %+v", a.Inheritance)
	}
	if !a.Ledger.AllMatched || a.Ledger.A0.Physical || a.Ledger.A2.Physical || a.Ledger.A4.Physical {
		t.Fatalf("bad ledger: %+v", a.Ledger)
	}
	if !a.Hierarchy.DimensionlessCombinatoric || !nearly(a.Hierarchy.A2OverA0AfterFactoring, 1.0/12.0, 1e-12) || !nearly(a.Hierarchy.A4OverA0AfterFactoring, 1.0/360.0, 1e-12) || !nearly(a.Hierarchy.A4OverA2AfterFactoring, 1.0/30.0, 1e-12) {
		t.Fatalf("bad hierarchy: %+v", a.Hierarchy)
	}
	if a.Hierarchy.SelectsCutoffLambda || a.Hierarchy.PhysicalNormalization || a.Airlock.F2MomentSelected || a.Airlock.F4MomentSelected || a.Airlock.CutoffLambdaSelected || a.Airlock.NewtonConstantDerived || a.Airlock.CosmologicalConstantDerived || a.Airlock.NativeNormalizationWrite {
		t.Fatalf("bad airlock: %+v %+v", a.Hierarchy, a.Airlock)
	}
	if a.Firewall.NewtonConstantImported || a.Firewall.PlanckScaleImported || a.Firewall.CutoffLambdaImported || a.Firewall.F2MomentImported || a.Firewall.F4MomentImported || a.Firewall.CosmologicalConstantImported || a.Firewall.NativeSpectralNormalizationWrite {
		t.Fatalf("firewall violation: %+v", a.Firewall)
	}
	if a.Next.Gate != 514 {
		t.Fatalf("unexpected next gate: %+v", a.Next)
	}
}

func TestGate513Markdown(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	md := Markdown(a)
	for _, want := range []string{"# Gate 513 Registry Audit", StatusRelativePrefactorHierarchyComputed, StatusFailedMomentRatiosDoNotSelectCutoff, StatusFirewallMomentNativeWriteBlocked, "1/12", "1/360", "Gate 514"} {
		if !strings.Contains(md, want) {
			t.Fatalf("markdown missing %q", want)
		}
	}
}

func TestGate513Theorem(t *testing.T) {
	result := Generation2SpectralMomentHierarchyCutoffSeparationAirlockAuditTheorem().Verify()
	for _, c := range result.Checks {
		if !c.Passed {
			t.Fatalf("check failed: %s %s", c.Name, c.Detail)
		}
	}
}
