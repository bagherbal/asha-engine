package familyaxiomclosureledger

import "testing"

func TestBuildDefault(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault failed: %v", err)
	}
	if !a.Final.FirewallPreserved {
		t.Fatalf("expected final firewall preserved")
	}
	if a.Parameters.ConditionalCompressedDim != 9 {
		t.Fatalf("expected conditional complex ledger dim 9, got %d", a.Parameters.ConditionalCompressedDim)
	}
	if a.Progression.MinimalCPGate != 417 {
		t.Fatalf("expected minimal CP gate 417, got %d", a.Progression.MinimalCPGate)
	}
	if !a.Seal.CoefficientsEnvironmental {
		t.Fatalf("expected coefficients environmental")
	}
	if a.Next.Gate != 419 {
		t.Fatalf("expected next gate 419, got %d", a.Next.Gate)
	}
}

func TestTheoremPasses(t *testing.T) {
	res := FamilyAxiomClosureLedgerFlavorFrontierSealTheorem().Run()
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
	for _, want := range []string{StatusEnvironmentalSealFormalized, StatusProjectFlavorSectorSealedComplete, StatusFirewallPreserved13Moduli, "Axiom progression ledger"} {
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
