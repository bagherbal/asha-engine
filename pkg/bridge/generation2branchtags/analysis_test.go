package generation2branchtags

import (
	"math"
	"strings"
	"testing"
)

func TestBuildDefault(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault failed: %v", err)
	}
	if a.Sieve.AcceptedCount != 2 || a.Sieve.RejectedCount != 5 {
		t.Fatalf("unexpected sieve counts: %s", FormatSieve(a.Sieve))
	}
	if !a.Sieve.CompletePositiveAccepted || !a.Sieve.CompleteNegativeAccepted || !a.Sieve.AllAcceptedBridgeOnly {
		t.Fatalf("complete branch tags not accepted bridge-only: %s", FormatSieve(a.Sieve))
	}
}

func TestEvaluateBranchCounts(t *testing.T) {
	cosOnly := EvaluateBranch(BranchRequest{Name: "cos", Cos3Phi: 0.25, HasCosineInvariant: true, BridgeOnly: true})
	if cosOnly.Accepted || cosOnly.BranchCount != CosineOnlyBranchCount || cosOnly.Verdict != StatusFailedCosineSixBranches {
		t.Fatalf("cosine-only branch should expose six branches: %s", FormatEvaluation(cosOnly))
	}
	cpOnly := EvaluateBranch(BranchRequest{Name: "cp", Cos3Phi: 0.25, HasCosineInvariant: true, CPOddSign: +1, HasCPOddSign: true, BridgeOnly: true})
	if cpOnly.Accepted || cpOnly.BranchCount != CPSignOnlyBranchCount || cpOnly.Verdict != StatusFailedCPSignOnlyThreeSheets {
		t.Fatalf("CP sign only should leave three sheets: %s", FormatEvaluation(cpOnly))
	}
	full := EvaluateBranch(BranchRequest{Name: "full", Cos3Phi: 0.25, HasCosineInvariant: true, CPOddSign: +1, HasCPOddSign: true, C3Sheet: 1, HasC3Sheet: true, BridgeOnly: true})
	if !full.Accepted || !full.Selected || full.BranchCount != CompleteBranchCount || full.Verdict != StatusCompleteBranchTagUnique {
		t.Fatalf("complete branch tag should select one phase: %s", FormatEvaluation(full))
	}
	back := math.Cos(3 * full.Phase)
	if math.Abs(back-0.25) > 1e-12 {
		t.Fatalf("selected phase does not reproduce cos(3phi): got %.15g", back)
	}
}

func TestUnsafeSelectorsRejected(t *testing.T) {
	physical := EvaluateBranch(BranchRequest{Name: "physical", Cos3Phi: 0.25, HasCosineInvariant: true, CPOddSign: +1, HasCPOddSign: true, C3Sheet: 0, HasC3Sheet: true, UsesCKMOrPMNS: true, BridgeOnly: true})
	if physical.Accepted || physical.Verdict != StatusFailedCKMPMNSSelectorRejected {
		t.Fatalf("CKM/PMNS selector should be rejected: %s", FormatEvaluation(physical))
	}
	promote := EvaluateBranch(BranchRequest{Name: "promote", Cos3Phi: 0.25, HasCosineInvariant: true, CPOddSign: +1, HasCPOddSign: true, C3Sheet: 0, HasC3Sheet: true, BridgeOnly: false, NativePromotionClaim: true})
	if promote.Accepted || promote.Verdict != StatusFailedNativePromotionRejected {
		t.Fatalf("native promotion should be rejected: %s", FormatEvaluation(promote))
	}
	invalid := EvaluateBranch(BranchRequest{Name: "invalid", Cos3Phi: 0.25, HasCosineInvariant: true, CPOddSign: +1, HasCPOddSign: true, C3Sheet: 9, HasC3Sheet: true, BridgeOnly: true})
	if invalid.Accepted || invalid.Verdict != StatusFailedInvalidBranchTag {
		t.Fatalf("invalid C3 sheet should be rejected: %s", FormatEvaluation(invalid))
	}
}

func TestTheoremPasses(t *testing.T) {
	res := Generation2OrientedComparatorBranchTagSieveCPSignLedgerTheorem().Run()
	if !res.Passed() {
		t.Fatalf("theorem audit checks did not pass:\n%s", res.Details())
	}
}

func TestRenderAuditContainsGate459Language(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	md := RenderAudit(a)
	for _, want := range []string{
		"# Gate 459 Registry Audit",
		StatusBridgeOnlyBranchTagValidated,
		"sigma_CP",
		"n_C3",
		StatusFailedCPSignOnlyThreeSheets,
		StatusFailedNativeCPSelectorAbsent,
		"CKM/PMNS",
	} {
		if !strings.Contains(md, want) {
			t.Fatalf("audit missing %q", want)
		}
	}
}
