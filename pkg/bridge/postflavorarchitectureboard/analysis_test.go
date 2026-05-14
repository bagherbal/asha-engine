package postflavorarchitectureboard

import "testing"

func TestBuildDefault(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault failed: %v", err)
	}
	if !a.Final.BoardReady {
		t.Fatalf("expected board ready")
	}
	if a.Final.NativeFlavorDim != 13 {
		t.Fatalf("expected native flavor dim 13, got %d", a.Final.NativeFlavorDim)
	}
	if a.Final.ConditionalFamilyDim != 9 {
		t.Fatalf("expected conditional family dim 9, got %d", a.Final.ConditionalFamilyDim)
	}
	if !a.Board.Ordered || len(a.Board.Nodes) < 12 {
		t.Fatalf("expected ordered board with at least 12 nodes")
	}
	if a.Next.Gate != 420 {
		t.Fatalf("expected next gate 420, got %d", a.Next.Gate)
	}
}

func TestTheoremPasses(t *testing.T) {
	res := PostFlavorArchitectureConsolidationFinalLawSpaceBoardTheorem().Run()
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
	for _, want := range []string{StatusFinalLawSpaceBoardReady, StatusNoFlavorReopening, StatusFirewallPreserved13Moduli, "Final architecture board"} {
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
