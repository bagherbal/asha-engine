package publicationbundlepreflight

import (
	"strings"
	"testing"
)

func TestBuildDefault(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault failed: %v", err)
	}
	if !a.Final.Ready || !a.Final.ManifestReady || !a.Final.FirewallChecklistReady {
		t.Fatalf("expected ready publication preflight: %+v", a.Final)
	}
	if a.Manifest.MissingCount != 0 || a.Manifest.RequiredCount < 12 {
		t.Fatalf("unexpected manifest: %+v", a.Manifest)
	}
	if a.Next.Gate != 426 {
		t.Fatalf("expected next gate 426, got %d", a.Next.Gate)
	}
}

func TestTheoremPasses(t *testing.T) {
	res := FinalPaperAssemblyPublicationBundlePreflightTheorem().Run()
	if !res.Passed() {
		t.Fatalf("theorem did not pass:\n%s", res.Details())
	}
}

func TestRenderAuditContainsPublicationLanguage(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	md := RenderAudit(a)
	for _, want := range []string{"Publication Bundle Preflight", StatusBundlePreflightReady, StatusFirewallPreserved13, "CLAIM_FIREWALL_CHECKLIST", "SECTION_SOURCE_MAP"} {
		if !strings.Contains(md, want) {
			t.Fatalf("audit missing %q", want)
		}
	}
}

func TestFirewallRowsHaveForbiddenWording(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	for _, row := range a.Firewall.Rows {
		if strings.TrimSpace(row.Topic) == "" || strings.TrimSpace(row.Forbidden) == "" || strings.TrimSpace(row.GateReference) == "" {
			t.Fatalf("firewall row incomplete: %+v", row)
		}
	}
}

func TestNoClaimDriftInExports(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	joined := a.Exports.PreflightMarkdown + a.Exports.FirewallMarkdown + a.Exports.AssemblyChecklistMD
	for _, forbidden := range []string{"predicts CKM", "predicts PMNS", "Yukawa values are derived", "cosmology is predicted"} {
		if strings.Contains(joined, forbidden) {
			t.Fatalf("export contains forbidden claim %q", forbidden)
		}
	}
}
