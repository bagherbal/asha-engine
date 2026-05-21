package generation2boundarysplitjetresponsefunctionalaudit

import (
	"strings"
	"testing"
)

func TestGate868JetLanes(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if a.Functional.First.Power != 1 || a.Functional.First.Rank != PiTopRank || a.Functional.First.OperatorTyped || a.Functional.First.NativeDerived {
		t.Fatalf("bad first jet: %s", FormatJet(a.Functional.First))
	}
	if a.Functional.Second.Power != 2 || a.Functional.Second.Rank != HRminRank || a.Functional.Second.OperatorTyped || a.Functional.Second.NativeDerived {
		t.Fatalf("bad second jet: %s", FormatJet(a.Functional.Second))
	}
}

func TestGate868FormalReconstructionButNoNativeFunctional(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Functional.ShapeCoherent || !near(a.Functional.ReconstructedAlpha, AlphaB) || a.Functional.Native || a.Functional.FirstJetCertified || a.Functional.SecondJetCertified {
		t.Fatalf("functional overpromoted: %s", FormatFunctional(a.Functional))
	}
}

func TestGate868SharedCoordinateAndTruncationObstructions(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.SharedS.FeedsFirstJet || !a.SharedS.FeedsSecondJet || a.SharedS.SharedJetFunctorCertified {
		t.Fatalf("bad shared S: %s", FormatSharedS(a.SharedS))
	}
	if !a.Truncation.ConstantTermAbsent || !a.Truncation.CubicAndHigherAbsent || a.Truncation.TruncationTheoremCertified || !containsAll(a.Truncation.Failures, []string{FailureNoTruncationTheorem, FailureNoHigherTermTheorem}) {
		t.Fatalf("bad truncation: %s", FormatTruncation(a.Truncation))
	}
}

func TestGate868Firewalls(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !firewallsOK(a.Firewalls) || a.Impact.CanUpdateNEff || a.Impact.CanPromoteToR3 {
		t.Fatalf("firewalls broken: %s | %s", FormatFirewalls(a.Firewalls), FormatImpact(a.Impact))
	}
}

func TestGate868Theorem(t *testing.T) {
	res := Generation2BoundarySplitJetResponseFunctionalAuditTheorem().Verify()
	if res.Status == "FAILED_ROUTE" {
		t.Fatalf("unexpected theorem failure: %+v", res)
	}
	for _, c := range res.Checks {
		if !c.Passed {
			t.Fatalf("failed check %s: %s", c.Name, c.Detail)
		}
	}
	joined := strings.Join(res.Notes, "\n")
	for _, want := range Statuses() {
		if !strings.Contains(joined, want) {
			t.Fatalf("missing status %s", want)
		}
	}
}
