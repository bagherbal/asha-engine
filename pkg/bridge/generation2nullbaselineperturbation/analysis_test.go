package generation2nullbaselineperturbation

import (
	"math"
	"strings"
	"testing"
)

func TestBuildDefaultGate481Cancellation(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if a.Inheritance.AlphaVac != 1 || a.Inheritance.IKVac != 0.5 {
		t.Fatalf("bad inherited baseline: %+v", a.Inheritance)
	}
	if a.Transport.IKVacCanReplaceSectorIK || a.Transport.PerturbationsNative || a.Transport.TransportPreviouslyNative {
		t.Fatalf("transport leaked native claim: %+v", a.Transport)
	}
	if !a.Proof.BaselineAlphaCancels || !a.Proof.BaselinePhiCancels || !a.Proof.OnlyPerturbationsRemain {
		t.Fatalf("baseline did not cancel: %+v", a.Proof)
	}
	if !a.Ledger.AllRowsBridgeOnly || !a.Ledger.AllRowsSynthetic || !a.Ledger.QuarkDistanceComputed || !a.Ledger.LeptonDistanceComputed {
		t.Fatalf("bad ledger: %+v", a.Ledger)
	}
	if a.Firewall.SectorIKSolvedByBaseline || a.Firewall.DUDNativePrediction || a.Firewall.DENuNativePrediction || a.Firewall.NativeRegistryWritten {
		t.Fatalf("firewall violated: %+v", a.Firewall)
	}
}

func TestPairUsesOnlyPerturbationDeltas(t *testing.T) {
	u := row("u", 0.2, 0.4)
	d := row("d", -0.1, -0.2)
	p := pair("test", u, d)
	if !p.BaselineCancelled {
		t.Fatal("baseline must cancel")
	}
	wantDA := -0.3
	wantDP := -0.6
	if math.Abs(p.DeltaAlpha-wantDA) > 1e-12 || math.Abs(p.DeltaPhi-wantDP) > 1e-12 {
		t.Fatalf("got delta=(%.15f,%.15f)", p.DeltaAlpha, p.DeltaPhi)
	}
}

func TestRenderAuditContainsGate481Language(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	md := RenderAudit(a)
	for _, want := range []string{"# Gate 481 Registry Audit", StatusLedgerDefined, "Baseline cancellation", StatusFailedIKVacAsSectorIK, "synthetic d_ud"} {
		if !strings.Contains(md, want) {
			t.Fatalf("audit missing %q", want)
		}
	}
}
