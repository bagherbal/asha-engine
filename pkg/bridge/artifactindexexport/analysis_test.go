package artifactindexexport

import (
	"strings"
	"testing"
)

func TestBuildDefault(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault failed: %v", err)
	}
	if !a.Final.ArtifactIndexReady || !a.Final.ReproChecklistReady || !a.Final.RootClean {
		t.Fatalf("expected ready clean artifact index: %+v", a.Final)
	}
	if a.Coverage.LastGate != 424 || a.Coverage.GateAuditCount < 227 {
		t.Fatalf("unexpected coverage: %+v", a.Coverage)
	}
	if a.Next.Gate != 425 {
		t.Fatalf("expected next gate 425, got %d", a.Next.Gate)
	}
}

func TestTheoremPasses(t *testing.T) {
	res := ArtifactIndexReproducibilityChecklistExportTheorem().Run()
	if !res.Passed() {
		t.Fatalf("theorem did not pass:\n%s", res.Details())
	}
}

func TestRenderAuditContainsArtifactLanguage(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	md := RenderAudit(a)
	for _, want := range []string{"Artifact Index", "Reproducibility Checklist", StatusArtifactIndexReady, StatusFirewallPreserved13, "docs/audits/gates"} {
		if !strings.Contains(md, want) {
			t.Fatalf("audit missing %q", want)
		}
	}
}

func TestArtifactRowsHavePolicies(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	for _, r := range a.Tree.Rows {
		if strings.TrimSpace(r.Path) == "" || strings.TrimSpace(r.Policy) == "" || strings.TrimSpace(r.Validation) == "" {
			t.Fatalf("row incomplete: %+v", r)
		}
	}
}

func TestReproPolicyAvoidsFullSuiteByDefault(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	for _, c := range a.Repro.Commands {
		if strings.Contains(c.Command, "go test ./...") && c.RunByDefault {
			t.Fatalf("full suite must not run by default")
		}
		if strings.Contains(c.Command, "go test ./internal/app") && c.RunByDefault {
			t.Fatalf("internal/app tests must not run by default")
		}
	}
}
