package manuscriptskeletonexport

import (
	"strings"
	"testing"
)

func TestBuildDefault(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault failed: %v", err)
	}
	if !a.Final.SkeletonReady || !a.Final.ProofExportReady {
		t.Fatalf("expected ready proof-export skeleton")
	}
	if len(a.Manuscript.Sections) < 10 || len(a.Manuscript.ProofObligations) < len(a.Manuscript.Sections)*2 {
		t.Fatalf("incomplete manuscript export")
	}
	if a.Final.NativeFlavorDim != NativeChargedFlavorDim || a.Final.ConditionalFamilyDim != ConditionalFamilyAxiomDim {
		t.Fatalf("unexpected flavor dimensions")
	}
	if a.Next.Gate != 422 {
		t.Fatalf("expected next gate 422, got %d", a.Next.Gate)
	}
}

func TestTheoremPasses(t *testing.T) {
	res := ManuscriptSkeletonSectionBySectionProofExportTheorem().Run()
	if !res.Passed() {
		t.Fatalf("theorem did not pass:\n%s", res.Details())
	}
}

func TestRenderAuditContainsSkeletonAndFirewalls(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	md := RenderAudit(a)
	for _, want := range []string{"Manuscript Skeleton", "Proof obligation matrix", StatusManuscriptSkeletonReady, StatusFirewallPreserved13, "No Yukawa values"} {
		if !strings.Contains(md, want) {
			t.Fatalf("audit missing %q", want)
		}
	}
}

func TestEverySectionHasTwoProofObligations(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	counts := map[string]int{}
	for _, p := range a.Manuscript.ProofObligations {
		counts[p.SectionID]++
	}
	for _, s := range a.Manuscript.Sections {
		if counts[s.ID] < 2 {
			t.Fatalf("section %s has %d proof obligations", s.ID, counts[s.ID])
		}
	}
}
