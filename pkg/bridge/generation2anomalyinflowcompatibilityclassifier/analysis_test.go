package generation2anomalyinflowcompatibilityclassifier

import (
	"strings"
	"testing"
)

func TestGate524Build(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Inheritance.Gate523ReportDefined || a.Inheritance.Gate523Rows != 4 || a.Inheritance.Gate523ZeroResidualRows != 4 || !a.Inheritance.Gate523BridgeOnly || !a.Inheritance.Gate523SyntheticOnly || a.Inheritance.Gate523ObservedImported || !a.Inheritance.Gate523HeterogeneousGuard || !a.Inheritance.Gate523CrossLedgerMergeRejected || a.Inheritance.Gate523NativeManifoldSelected || !a.Inheritance.Gate517IndexSocket || !a.Inheritance.Gate517APSSocket || !a.Inheritance.Gate517InflowSocket || !a.Inheritance.Gate490GaugeAnomaliesCancel || !a.Inheritance.Gate490MixedGaugeGravityCancel || !a.Inheritance.Gate490WittenSU2Cancel {
		t.Fatalf("bad inheritance: %+v", a.Inheritance)
	}
	if !a.Inflow.NativeCapacityConfirmed || !a.Inflow.ScaleFree || !a.Inflow.MassFlavorIndependent || !a.Inflow.BulkCharacteristicClassesPresent || !a.Inflow.ChernSimonsTransgressionSocket || a.Inflow.BoundaryTheorySelected || a.Inflow.BoundaryConditionSelected || a.Inflow.EtaSpectrumDerived || a.Inflow.GlobalAnomalyCoefficientSelected {
		t.Fatalf("bad inflow: %+v", a.Inflow)
	}
	if a.Compatibility.CompatibleClassCount != 3 || !a.Compatibility.APSBoundaryFixtureCompatible || !a.Compatibility.SpinBordismFixtureCompatible || !a.Compatibility.SpinCBordismFixtureCompatible || !a.Compatibility.HeterogeneousGuardPreserved || a.Compatibility.CrossFixtureIdentityAllowed || !a.Compatibility.CrossFixtureMergeRejected || !a.Compatibility.ClassifiesButDoesNotSelect || !a.Compatibility.BoundaryCurrentConservationSocket || a.Compatibility.NativeManifoldSelected || a.Compatibility.NativeBoundarySelected || a.Compatibility.NativeBordismClassSelected {
		t.Fatalf("bad compatibility: %+v", a.Compatibility)
	}
	if a.Firewall.ObservedTopologyImported || a.Firewall.ObservedBoundaryImported || a.Firewall.ObservedBordismImported || a.Firewall.ObservedEtaImported || a.Firewall.ObservedBoundarySpectrumImported || a.Firewall.NewtonPlanckCosmologyImported || !a.Firewall.InflowCapacityNative || a.Firewall.BoundaryConditionNative || a.Firewall.EtaSpectrumNative || a.Firewall.BordismClassNative || a.Firewall.CharacteristicNumbersNative || a.Firewall.CrossFixtureMergeNative || a.Firewall.GravitationalThetaNative || a.Firewall.NativeRegistryWritten {
		t.Fatalf("firewall violation: %+v", a.Firewall)
	}
	if a.Next.Gate != 525 {
		t.Fatalf("bad next gate: %+v", a.Next)
	}
}

func TestGate524Markdown(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	md := Markdown(a)
	for _, want := range []string{"# Gate 524 Registry Audit", StatusNativeInflowCapacityConfirmed, StatusAPSBoundaryClassCompatible, StatusFirewallNativeWriteBlocked, "Gate 525"} {
		if !strings.Contains(md, want) {
			t.Fatalf("markdown missing %q", want)
		}
	}
}

func TestGate524Theorem(t *testing.T) {
	result := Generation2AnomalyInflowCompatibilityClassifierForBridgeTopologyClassesTheorem().Verify()
	for _, c := range result.Checks {
		if !c.Passed {
			t.Fatalf("check failed: %s %s", c.Name, c.Detail)
		}
	}
}
