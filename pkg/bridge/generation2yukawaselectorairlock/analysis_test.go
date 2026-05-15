package generation2yukawaselectorairlock

import (
	"strings"
	"testing"
)

func TestBuildDefaultGate489YukawaSelectorAirlock(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Inheritance.Gate488YukawaSocketInherited || a.Inheritance.Gate488NativeUpDownOperatorsFound || a.Inheritance.Gate488YukawaValuesDerived {
		t.Fatalf("bad inheritance: %+v", a.Inheritance)
	}
	if a.Ledger.NativeSelectorsPassing != 0 || a.Ledger.RephasingInvariantConstraints != 0 || a.Ledger.ObservedDataImported {
		t.Fatalf("selector ledger leaked: %+v", a.Ledger)
	}
	if !a.Variational.NativeYukawaSlotsExist || a.Variational.NativeYukawaValuesDerived || a.Variational.RankThreeUpMatrixDerived || a.Variational.RankThreeDownMatrixDerived || a.Variational.RelativeEigenbasisDerived || a.Variational.SelectorFound {
		t.Fatalf("bad variational audit: %+v", a.Variational)
	}
	if !a.Airlock.NativeYukawaSelectorBranchClosed || !a.Airlock.CKMOrientationEnvironmental || a.Airlock.NativeCKMPredictionAllowed {
		t.Fatalf("bad airlock decision: %+v", a.Airlock)
	}
	if a.Firewall.ObservedCKMImported || a.Firewall.ObservedYukawaEntriesImported || a.Firewall.NativeYukawaMatrixWritten || a.Firewall.NativeRegistryWritten {
		t.Fatalf("firewall leak: %+v", a.Firewall)
	}
}

func TestRenderAuditGate489(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	md := RenderAudit(a)
	for _, want := range []string{StatusYukawaAirlockClosedNative, StatusCKMEnvironmentalQuarantineFormal, "Yukawa selector ledger", "spectral action", "environmental", "Gate 490"} {
		if !strings.Contains(md, want) {
			t.Fatalf("audit missing %q", want)
		}
	}
}
