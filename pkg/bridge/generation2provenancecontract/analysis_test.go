package generation2provenancecontract

import (
	"strings"
	"testing"
)

func TestBuildDefault(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault failed: %v", err)
	}
	if a.Contract.RequiredFieldCount != MinimumFields {
		t.Fatalf("unexpected field count: %s", FormatContract(a.Contract))
	}
	if a.Sieve.AcceptedCount != 2 || a.Sieve.RejectedCount != 7 {
		t.Fatalf("unexpected sieve counts: %s", FormatSieve(a.Sieve))
	}
	if !a.Sieve.NoAcceptedNativeExport {
		t.Fatalf("accepted native export leaked: %s", FormatSieve(a.Sieve))
	}
}

func TestEvaluateRejectsMissingAndPromotionRoutes(t *testing.T) {
	base := ComparatorRecord{Sector: "up", Observable: "I_K", ValueKind: "symbolic", ValueExpression: "I_K", Scale: "MZ", Scheme: "MSbar", Source: "source", SourceVersion: "v", Uncertainty: "sigma", Dimensionless: true, BridgeOnly: true}
	if got := evaluate(base); !got.Passed {
		t.Fatalf("base record should pass: %s", FormatRecord(got))
	}
	missing := base
	missing.Scale = ""
	if got := evaluate(missing); got.Passed || got.Verdict != StatusFailedMissingSectorScaleScheme {
		t.Fatalf("missing scale should be rejected: %s", FormatRecord(got))
	}
	promote := base
	promote.BridgeOnly = false
	promote.NativePromotionClaim = true
	if got := evaluate(promote); got.Passed || got.Verdict != StatusFailedNativePromotionAttempt {
		t.Fatalf("native promotion should be rejected: %s", FormatRecord(got))
	}
	oriented := base
	oriented.RequiresOrientedInverse = true
	if got := evaluate(oriented); got.Passed || got.Verdict != StatusFailedBranchTagRequired {
		t.Fatalf("missing branch tag should be rejected: %s", FormatRecord(got))
	}
	oriented.BranchTag = "phi-branch:+,n=0"
	if got := evaluate(oriented); !got.Passed {
		t.Fatalf("branch-tagged oriented record should pass: %s", FormatRecord(got))
	}
}

func TestRenderAuditContainsFirewallLanguage(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault failed: %v", err)
	}
	md := RenderAudit(a)
	for _, want := range []string{
		"# Gate 457 Registry Audit",
		StatusTextureComparatorContractValidated,
		StatusFailedObservedDefaultModeRejected,
		"sector",
		"scheme",
		"source_version",
		"no observed flavor values",
	} {
		if !strings.Contains(md, want) {
			t.Fatalf("audit missing %q", want)
		}
	}
}
