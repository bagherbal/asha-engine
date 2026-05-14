package reviewerobjectionmatrix

import (
	"strings"
	"testing"
)

func TestBuildDefault(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault failed: %v", err)
	}
	if !a.Final.MatrixReady || !a.Final.BoundariesReady {
		t.Fatalf("expected ready reviewer matrix")
	}
	if a.Final.NativeFlavorDim != NativeChargedFlavorDim || a.Final.ConditionalFamilyDim != ConditionalFamilyAxiomDim {
		t.Fatalf("unexpected flavor dimensions")
	}
	if len(a.Matrix.Rows) < 10 || a.Matrix.HighRiskCount < 3 {
		t.Fatalf("matrix incomplete: %+v", a.Matrix)
	}
	if a.Next.Gate != 424 {
		t.Fatalf("expected next gate 424, got %d", a.Next.Gate)
	}
}

func TestTheoremPasses(t *testing.T) {
	res := ReviewerObjectionMatrixRebuttalReadinessExportTheorem().Run()
	if !res.Passed() {
		t.Fatalf("theorem did not pass:\n%s", res.Details())
	}
}

func TestRenderAuditContainsReviewerLanguage(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	md := RenderAudit(a)
	for _, want := range []string{"Reviewer Objection Matrix", "Rebuttal guide", StatusReviewerMatrixReady, StatusFirewallPreserved13, "derives Yukawa matrices"} {
		if !strings.Contains(md, want) {
			t.Fatalf("audit missing %q", want)
		}
	}
}

func TestRowsHaveReferencesAndBoundaries(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	for _, r := range a.Matrix.Rows {
		if len(r.GateReferences) == 0 {
			t.Fatalf("row %s lacks references", r.ID)
		}
		if strings.TrimSpace(r.Boundary) == "" {
			t.Fatalf("row %s lacks boundary", r.ID)
		}
		if strings.TrimSpace(r.SafeWording) == "" || strings.TrimSpace(r.UnsafeWording) == "" {
			t.Fatalf("row %s lacks wording", r.ID)
		}
	}
}

func TestForbiddenPhrasesProtectFirewalls(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	joined := strings.ToLower(join(a.Guide.ForbiddenPhrases))
	for _, want := range []string{"yukawa", "triality", "q4", "cosmology"} {
		if !strings.Contains(joined, want) {
			t.Fatalf("forbidden phrases missing %s", want)
		}
	}
}
