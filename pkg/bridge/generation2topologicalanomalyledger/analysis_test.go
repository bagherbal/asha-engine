package generation2topologicalanomalyledger

import (
	"strings"
	"testing"
)

func TestBuildDefaultGate490AnomalyLedger(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Inheritance.Gate489FlavorAirlockClosed || !a.Inheritance.NoFlavorDataImported {
		t.Fatalf("bad inheritance: %+v", a.Inheritance)
	}
	if len(a.Ledger.Multiplets) != 6 || a.Ledger.LeftHandedWeylStates != 16 || a.Ledger.WeakDoubletCount != 4 || !a.Ledger.WeakDoubletCountEven {
		t.Fatalf("bad charge ledger: %+v", a.Ledger)
	}
	if !a.Anomaly.AllPerturbativeGaugeCancel || !a.Anomaly.AllMixedGaugeGravityCancel || !a.Anomaly.SU2GlobalWittenCancels || a.Anomaly.ZeroTraceCount != len(a.Anomaly.Moments) {
		t.Fatalf("anomaly traces failed: %+v", a.Anomaly)
	}
	if !a.Anomaly.ExistingGate79Consistent {
		t.Fatalf("expected Gate79 consistency: %+v", a.Anomaly)
	}
	if !a.Stability.GaugeStabilityLedgerSatisfied || a.Stability.YukawaTextureSelected || a.Stability.CKMJarlskogDerived {
		t.Fatalf("bad stability theorem: %+v", a.Stability)
	}
	if a.Firewall.ObservedMassesImported || a.Firewall.ObservedCKMImported || a.Firewall.NativeFlavorModuliChanged {
		t.Fatalf("firewall leak: %+v", a.Firewall)
	}
}

func TestGate490ExactTraceValues(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	want := map[string]string{
		"Tr(Y)":       "0",
		"Tr(Y^3)":     "0",
		"SU(2)_L^2·Y": "0",
		"SU(3)_c^2·Y": "0",
		"SU(3)_c^3":   "0",
		"SU(2)_L^3":   "0",
		"Tr(B-L)":     "0",
		"Tr((B-L)^3)": "0",
	}
	for _, m := range a.Anomaly.Moments {
		if v, ok := want[m.Symbol]; ok && m.Value != v {
			t.Fatalf("%s = %s, want %s", m.Symbol, m.Value, v)
		}
	}
}

func TestRenderAuditGate490(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	md := RenderAudit(a)
	for _, want := range []string{StatusExactGaugeAnomalyCancellation, StatusWittenSU2GlobalAnomalyCleared, "Topological Charge", "Tr(Y^3)", "SU(3)_c^3", "Gate 491"} {
		if !strings.Contains(md, want) {
			t.Fatalf("audit missing %q", want)
		}
	}
}
