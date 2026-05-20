package generation2koidewalloffsetsourcecandidateaudit

import (
	"math"
	"testing"
)

func TestGate585KoideWallOffsetSourceCandidateAudit(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if math.Abs(a.Target.PrimaryEpsilonRad-0.039569756309433) > 5e-15 {
		t.Fatalf("unexpected epsilon target: %s", FormatTarget(a.Target))
	}
	if a.Candidates.Best.Name != "1/(8π)" {
		t.Fatalf("expected 1/(8π) as nearest candidate: %s", FormatCandidateSet(a.Candidates))
	}
	if math.Abs(a.Loop.OneOver8Pi.Value-0.039788735772973836) > 1e-15 {
		t.Fatalf("unexpected 1/(8π): %s", FormatLoop(a.Loop))
	}
	if math.Abs(a.Loop.OneOver8Pi.RelativeResidual-0.0055340109205737065) > 1e-14 {
		t.Fatalf("unexpected 1/(8π) relative residual: %s", FormatLoop(a.Loop))
	}
	if !a.Loop.OneOver8Pi.Near || a.Loop.OneOver8Pi.Certified {
		t.Fatalf("1/(8π) should be near but not certified: %s", FormatLoop(a.Loop))
	}
	if len(a.Candidates.CertifiedCandidates) != 0 || a.Decision.CertifiedSource || a.Final.CandidateCertified {
		t.Fatalf("no source candidate should certify: %s", FormatDecision(a.Decision))
	}
	if a.Firewalls.DerivesEpsilon || a.Firewalls.PromotesObservedAsNative || !a.Firewalls.PreservesGate352 {
		t.Fatalf("firewall broken: %s", FormatFirewalls(a.Firewalls))
	}
}

func TestGate585Theorem(t *testing.T) {
	res := Generation2KoideWallOffsetSourceCandidateAuditTheorem().Run()
	if res.Status == "FAILED_ROUTE" {
		t.Fatalf("unexpected failed route: %+v", res)
	}
	for _, c := range res.Checks {
		if !c.Passed {
			t.Fatalf("check failed: %+v", c)
		}
	}
}
