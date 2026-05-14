package complexsectorsourcephase

import "testing"

func TestBuildDefault(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault failed: %v", err)
	}
	if !a.Firewall.FirewallPreserved {
		t.Fatalf("expected native firewall preserved")
	}
	if a.Parameters.BestConditionalComplexDim != 9 {
		t.Fatalf("expected complex phase charged ledger dim 9, got %d", a.Parameters.BestConditionalComplexDim)
	}
	if !a.CPSample.NonzeroCPCapacity {
		t.Fatalf("expected nonzero CP capacity")
	}
	if a.Algebra.GeneratedComplexAlgebraDim != 9 {
		t.Fatalf("expected generated algebra dim 9, got %d", a.Algebra.GeneratedComplexAlgebraDim)
	}
	if a.Next.Gate != 418 {
		t.Fatalf("expected next gate 418, got %d", a.Next.Gate)
	}
}

func TestTheoremPasses(t *testing.T) {
	res := ComplexSectorSourceCPPhaseAxiomSieveTheorem().Run()
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
	for _, want := range []string{StatusCPCapacityActivated, StatusFailedCPValueNotPredicted, StatusFirewallPreserved13Moduli, "CP-capacity sample"} {
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
