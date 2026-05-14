package perslotmonotonicityseal

import (
	"math"
	"testing"
)

func TestGate291SealLocksRPlusOnly(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Seal.Active || !a.Seal.Phenomenological || a.Seal.NativeTheorem {
		t.Fatalf("seal provenance wrong: %+v", a.Seal)
	}
	if !a.Locked.UniqueUnderSeal || a.Locked.NativeUnique || a.Locked.SelectedBranch.Name != "r_plus" {
		t.Fatalf("expected sealed r_plus lock, got %+v", a.Locked)
	}
	if len(a.Locked.VetoedBranches) != 1 || a.Locked.VetoedBranches[0] != "r_minus" {
		t.Fatalf("expected sealed r_minus veto, got %+v", a.Locked.VetoedBranches)
	}
}

func TestGate291TraceShapeReproducesContactLambda(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if a.Trace.D2 <= 0 || a.Trace.D4 <= 0 {
		t.Fatalf("nonpositive moments: %+v", a.Trace)
	}
	want := 1197.0 / 4624.0
	if math.Abs(a.Trace.D4OverD2Squared-want) > 1e-12 {
		t.Fatalf("shape mismatch: got %.17g want %.17g", a.Trace.D4OverD2Squared, want)
	}
	if !a.Trace.ShapeMatchesContact {
		t.Fatalf("expected shape match: %+v", a.Trace)
	}
}

func TestGate291HiggsFirewallPreserved(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if a.Higgs.PhysicalHiggsPredictionClaimed || a.Higgs.HeatKernelProjectionDerived || a.Higgs.ScalarGaugeNormalization {
		t.Fatalf("higgs firewall breached: %+v", a.Higgs)
	}
	if a.Firewalls.FiniteCorePolluted || !a.Firewalls.RawProxyNotPromotedToA2A4 {
		t.Fatalf("firewall not preserved: %+v", a.Firewalls)
	}
}
