package generation2koideloopdeficitpmnsorientationaudit

import (
	"math"
	"testing"
)

func TestGate587KoideLoopDeficitPMNSOrientationAudit(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if math.Abs(a.Runtime.Kappa-0.005503554191574556) > 1e-15 {
		t.Fatalf("unexpected kappa: %s", FormatRuntime(a.Runtime))
	}
	if math.Abs(a.Invariants.JPMNS-(-0.017769863116582574)) > 1e-15 {
		t.Fatalf("unexpected J_PMNS: %s", FormatInvariants(a.Invariants))
	}
	if math.Abs(a.Invariants.SqrtAbsJ-0.13330365004973635) > 1e-14 {
		t.Fatalf("unexpected sqrt J_PMNS: %s", FormatInvariants(a.Invariants))
	}
	best := a.Candidates.BestPMNSAssisted
	if best.Name != "alpha_2(M_Z)/(2π)/c13" {
		t.Fatalf("unexpected best PMNS-assisted candidate: %s", FormatCandidateSet(a.Candidates))
	}
	if math.Abs(best.RelativeResidual-(-0.008420618697162418)) > 1e-14 {
		t.Fatalf("unexpected best PMNS residual: %s", FormatCandidate(best))
	}
	if best.Certified || best.CoversKappa {
		t.Fatalf("PMNS-assisted candidate should not certify/cover kappa: %s", FormatCandidate(best))
	}
	if !a.CKM.PMNSAssistedBetterThanSqrtJCKM || !a.CKM.MidpointStillClosestNumeric {
		t.Fatalf("unexpected CKM comparison: %s", FormatCKM(a.CKM))
	}
	if !a.Uncertainty.AnyCandidateCovers || a.Decision.AnyCandidateCertified || a.Final.AnyCertified {
		t.Fatalf("candidate range should cover but no candidate should certify under uncertainty: %s", FormatDecision(a.Decision))
	}
	if a.Firewalls.DerivesKappa || a.Firewalls.DerivesPMNS || a.Firewalls.PromotesObservedAsNative || !a.Firewalls.PreservesGate352 {
		t.Fatalf("firewall broken: %s", FormatFirewalls(a.Firewalls))
	}
}

func TestGate587Theorem(t *testing.T) {
	res := Generation2KoideLoopDeficitPMNSOrientationAuditTheorem().Run()
	if res.Status == "FAILED_ROUTE" {
		t.Fatalf("unexpected failed route: %+v", res)
	}
	for _, c := range res.Checks {
		if !c.Passed {
			t.Fatalf("check failed: %+v", c)
		}
	}
}
