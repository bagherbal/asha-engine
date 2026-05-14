package executiveabstractclaimaudit

import (
	"strings"
	"testing"
)

func TestBuildDefault(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault failed: %v", err)
	}
	if !a.Final.ExecutiveReady || !a.Final.ClaimAuditReady {
		t.Fatalf("expected ready executive claim audit")
	}
	if a.Final.NativeFlavorDim != NativeChargedFlavorDim || a.Final.ConditionalFamilyDim != ConditionalFamilyAxiomDim {
		t.Fatalf("unexpected flavor dimensions")
	}
	if a.ClaimAudit.NativeCount < 4 || a.ClaimAudit.FirewallCount < 2 || a.ClaimAudit.FailedRouteCount < 3 {
		t.Fatalf("claim audit classification incomplete: %+v", a.ClaimAudit)
	}
	if a.Next.Gate != 423 {
		t.Fatalf("expected next gate 423, got %d", a.Next.Gate)
	}
}

func TestTheoremPasses(t *testing.T) {
	res := ExecutiveAbstractClaimAuditSummaryExportTheorem().Run()
	if !res.Passed() {
		t.Fatalf("theorem did not pass:\n%s", res.Details())
	}
}

func TestRenderAuditContainsClaimLanguage(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	md := RenderAudit(a)
	for _, want := range []string{"Executive Abstract", "Claim-audit table", StatusExecutiveSummaryReady, StatusFirewallPreserved13, "no Yukawa values"} {
		if !strings.Contains(md, want) {
			t.Fatalf("audit missing %q", want)
		}
	}
}

func TestNonClaimsAndFirewallsAreExplicit(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if len(a.Abstract.NonClaims) < 5 {
		t.Fatalf("not enough non-claims")
	}
	joined := strings.ToLower(join(a.Abstract.NonClaims))
	for _, want := range []string{"yukawa", "ckm", "pmns", "cosmological", "quarantined"} {
		if !strings.Contains(joined, want) {
			t.Fatalf("non-claim ledger missing %s", want)
		}
	}
}
