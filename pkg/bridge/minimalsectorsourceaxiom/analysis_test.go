package minimalsectorsourceaxiom

import "testing"

func TestBuildDefault(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault failed: %v", err)
	}
	if !a.Firewall.FirewallPreserved {
		t.Fatalf("expected native firewall preserved")
	}
	if a.Parameters.BestConditionalRealDim != 6 {
		t.Fatalf("expected real sector-source charged ledger dim 6, got %d", a.Parameters.BestConditionalRealDim)
	}
	if a.Parameters.BestConditionalCPDim != 9 {
		t.Fatalf("expected CP-capable charged ledger dim 9, got %d", a.Parameters.BestConditionalCPDim)
	}
	if a.Commutator.SampleMassCommutatorNorm <= 0 {
		t.Fatalf("expected nonzero sample commutator")
	}
	if a.Next.Gate != 417 {
		t.Fatalf("expected next gate 417, got %d", a.Next.Gate)
	}
}

func TestTheoremPasses(t *testing.T) {
	res := MinimalSectorSourceAxiomConsistencyParameterCountingSieveTheorem().Run()
	if !res.Passed() {
		t.Fatalf("theorem did not pass:\n%s", res.Details())
	}
}

func TestRenderAuditContainsKeyStatuses(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	md := RenderAudit(a)
	for _, want := range []string{StatusRealSourceSixParameterLedger, StatusFailedRealNoCPPhase, StatusFirewallPreserved13Moduli, "charged texture ledger"} {
		if !stringsContains(md, want) {
			t.Fatalf("audit missing %q", want)
		}
	}
}

func stringsContains(s, sub string) bool {
	for i := 0; i+len(sub) <= len(s); i++ {
		if s[i:i+len(sub)] == sub {
			return true
		}
	}
	return false
}
