package structuralfamilyboardexport

import (
	"strings"
	"testing"
)

func TestBuildDefault(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault failed: %v", err)
	}
	if !a.Final.Ready || !a.Final.NoNewPhysicsClaim {
		t.Fatalf("expected ready no-new-physics export: %+v", a.Final)
	}
	if a.Board.PromotedRows != 3 || a.Board.QuarantinedRows != 2 {
		t.Fatalf("unexpected structural board counts: %+v", a.Board)
	}
	if !a.Firewall.ForbidsMassPrediction || !a.Firewall.ForbidsMixingPrediction || !a.Firewall.ForbidsYukawaPrediction {
		t.Fatalf("firewall addendum incomplete: %+v", a.Firewall)
	}
	if a.Next.Gate != 450 {
		t.Fatalf("expected next gate 450, got %d", a.Next.Gate)
	}
}

func TestTheoremPasses(t *testing.T) {
	res := StructuralFamilyBoardManuscriptDeltaPatchTheorem().Run()
	if !res.Passed() {
		t.Fatalf("theorem did not pass:\n%s", res.Details())
	}
}

func TestRenderAuditContainsKeyPost444Language(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	md := RenderAudit(a)
	for _, want := range []string{StatusManuscriptDeltaReady, "Structural family board", "K_gen", "Generation-2 bare", "POST444_MANUSCRIPT_DELTA.md", "no observed muon/charm mass"} {
		if !strings.Contains(md, want) {
			t.Fatalf("audit missing %q", want)
		}
	}
}

func TestExportHasForbiddenClaims(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	export := strings.ToLower(a.Exports.CombinedMarkdown)
	for _, want := range []string{"predicts no observed", "yukawa", "ckm", "pmns", "cosmological"} {
		if !strings.Contains(export, want) {
			t.Fatalf("combined export missing %q", want)
		}
	}
}
