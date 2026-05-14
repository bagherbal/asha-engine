package publicationtheorematlas

import (
	"strings"
	"testing"
)

func TestBuildDefault(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault failed: %v", err)
	}
	if !a.Final.AtlasReady || !a.Atlas.Acyclic {
		t.Fatalf("expected ready acyclic atlas")
	}
	if len(a.Atlas.Nodes) < 20 || len(a.Atlas.Edges) < 20 {
		t.Fatalf("expected publication-size graph, got nodes=%d edges=%d", len(a.Atlas.Nodes), len(a.Atlas.Edges))
	}
	if a.Final.NativeFlavorDim != NativeChargedFlavorDim || a.Final.ConditionalFamilyDim != ConditionalFamilyAxiomDim {
		t.Fatalf("unexpected flavor dimensions")
	}
	if a.Next.Gate != 421 {
		t.Fatalf("expected next gate 421, got %d", a.Next.Gate)
	}
}

func TestTheoremPasses(t *testing.T) {
	res := PublicationGradeTheoremAtlasDependencyGraphExportTheorem().Run()
	if !res.Passed() {
		t.Fatalf("theorem did not pass:\n%s", res.Details())
	}
}

func TestRenderAuditContainsGraphExportsAndFirewalls(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	md := RenderAudit(a)
	for _, want := range []string{"graph TD", "digraph ASHA_Gate420_Atlas", StatusPublicationAtlasReady, StatusFirewallPreserved13, "Failed-route index"} {
		if !strings.Contains(md, want) {
			t.Fatalf("audit missing %q", want)
		}
	}
}

func TestTopologicalOrderCoversEveryNode(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	seen := map[string]bool{}
	for _, id := range a.Atlas.TopologicalOrder {
		seen[id] = true
	}
	for _, n := range a.Atlas.Nodes {
		if !seen[n.ID] {
			t.Fatalf("topological order missing node %s", n.ID)
		}
	}
}
