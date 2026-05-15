package generation2topologysectorclosingledger

import (
	"strings"
	"testing"
)

func TestGate525Build(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Inheritance.Gate524InflowCapacityConfirmed || a.Inheritance.Gate524CompatibleClassCount != 3 || !a.Inheritance.Gate524HeterogeneousGuard || !a.Inheritance.Gate524CrossFixtureMergeRejected || a.Inheritance.Gate524BoundarySelected || a.Inheritance.Gate524EtaSpectrumDerived || a.Inheritance.Gate524ObservedTopologyImported || !a.Inheritance.Gate524NativeWriteBlocked {
		t.Fatalf("bad Gate524 inheritance: %+v", a.Inheritance)
	}
	if !a.Inheritance.Gate489FlavorAirlockClosed || !a.Inheritance.Gate489CKMEnvironmental || a.Inheritance.Gate489ObservedFlavorImported || !a.Inheritance.Gate508ElectroweakFirewallClosed || !a.Inheritance.Gate508Diag114NotMassRatio || !a.Inheritance.Gate508WeakAngleBlocked || !a.Inheritance.Gate508WZMassesBlocked || !a.Inheritance.Gate514GravityAirlockClosed || !a.Inheritance.Gate514CutoffBlocked || !a.Inheritance.Gate514F2F4Blocked || !a.Inheritance.Gate514NewtonBlocked || !a.Inheritance.Gate514CosmologicalBlocked {
		t.Fatalf("bad sector inheritance: %+v", a.Inheritance)
	}
	if a.Ledger.NativeLawEntries != 4 || a.Ledger.BridgeComparatorEntries != 4 || a.Ledger.EnvironmentalHistoryEntries != 4 || a.Ledger.ClosedFirewallEntries != 4 || !a.Ledger.AnomalyCancellationNative || !a.Ledger.CharacteristicClassSocketsNative || !a.Ledger.APSInflowCapacityNative || !a.Ledger.BordismClassifierBridgeReady || !a.Ledger.TopologyResidualReportBridgeReady || a.Ledger.GlobalManifoldSelected || a.Ledger.BoundaryConditionSelected || a.Ledger.EtaSpectrumDerived || a.Ledger.CharacteristicNumbersDerived || a.Ledger.HeterogeneousFixturesMerged {
		t.Fatalf("bad closing ledger: %+v", a.Ledger)
	}
	if !a.Locks.FlavorSectorClosed || !a.Locks.ElectroweakScaleSectorClosed || !a.Locks.GravityNormalizationSectorClosed || !a.Locks.TopologySectorClosed || a.Locks.ReopenFlavorFirewall || a.Locks.ReopenEWScaleFirewall || a.Locks.ReopenGravityScaleFirewall || a.Locks.ReopenTopologySelectionFirewall {
		t.Fatalf("bad locks: %+v", a.Locks)
	}
	if a.Frontier.SelectedGate != 526 || !a.Frontier.LorentzianCausalSignatureLive || a.Frontier.RequiresObservedConstants || a.Frontier.ReopensSealedFirewalls || a.Frontier.PredictsMassesOrScales || a.Frontier.SelectsManifoldTopology {
		t.Fatalf("bad frontier: %+v", a.Frontier)
	}
	if a.Firewall.ObservedFlavorImported || a.Firewall.ObservedElectroweakImported || a.Firewall.ObservedGravityCosmologyImported || a.Firewall.ObservedTopologyBoundaryImported || a.Firewall.NativeYukawaWrite || a.Firewall.NativeCKMWrite || a.Firewall.NativeWZMassWrite || a.Firewall.NativeNewtonWrite || a.Firewall.NativeCosmologicalWrite || a.Firewall.NativeManifoldWrite || a.Firewall.NativeBoundaryWrite || a.Firewall.NativeEtaWrite || a.Firewall.NativeRegistryWritten {
		t.Fatalf("firewall violation: %+v", a.Firewall)
	}
}

func TestGate525Markdown(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	md := Markdown(a)
	for _, want := range []string{"# Gate 525 Registry Audit", StatusClosingLedgerConstructed, StatusNativeFrontierSelectedLorentzian, StatusFirewallNativeWriteBlocked, "Gate 526"} {
		if !strings.Contains(md, want) {
			t.Fatalf("markdown missing %q", want)
		}
	}
}

func TestGate525Theorem(t *testing.T) {
	result := Generation2TopologySectorClosingLedgerAndNativeFrontierSelectionTheorem().Verify()
	for _, c := range result.Checks {
		if !c.Passed {
			t.Fatalf("check failed: %s %s", c.Name, c.Detail)
		}
	}
}
