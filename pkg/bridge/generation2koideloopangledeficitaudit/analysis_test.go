package generation2koideloopangledeficitaudit

import (
	"math"
	"testing"
)

func TestGate586KoideLoopAngleDeficitAudit(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if math.Abs(a.Definition.Kappa-0.005503554191574556) > 1e-15 {
		t.Fatalf("unexpected kappa: %s", FormatDefinition(a.Definition))
	}
	if a.Candidates.Best.Name != "sqrt(J_CKM)" {
		t.Fatalf("expected sqrt(J_CKM) as nearest kappa candidate: %s", FormatCandidateSet(a.Candidates))
	}
	if math.Abs(a.Orientation.SqrtJCKM.Value-0.005583004145400101) > 1e-15 {
		t.Fatalf("unexpected sqrtJ: %s", FormatOrientation(a.Orientation))
	}
	if math.Abs(a.Orientation.SqrtJCKM.RelativeResidual-0.014436117290745676) > 1e-14 {
		t.Fatalf("unexpected sqrtJ residual: %s", FormatOrientation(a.Orientation))
	}
	if !a.Orientation.SqrtJCKM.Near || a.Orientation.SqrtJCKM.Certified {
		t.Fatalf("sqrtJ should be near but not certified: %s", FormatOrientation(a.Orientation))
	}
	if !a.Corrections.Alpha2Over2Pi.Near || a.Corrections.Alpha2Over2Pi.Certified {
		t.Fatalf("alpha2/(2pi) should be near but not certified: %s", FormatCorrections(a.Corrections))
	}
	if len(a.Candidates.CertifiedCandidates) != 0 || a.Decision.CertifiedSource || a.Final.CandidateCertified {
		t.Fatalf("no kappa candidate should certify: %s", FormatDecision(a.Decision))
	}
	if a.Firewalls.DerivesKappa || a.Firewalls.PromotesObservedAsNative || !a.Firewalls.PreservesGate352 {
		t.Fatalf("firewall broken: %s", FormatFirewalls(a.Firewalls))
	}
}

func TestGate586Theorem(t *testing.T) {
	res := Generation2KoideLoopAngleDeficitAuditTheorem().Run()
	if res.Status == "FAILED_ROUTE" {
		t.Fatalf("unexpected failed route: %+v", res)
	}
	for _, c := range res.Checks {
		if !c.Passed {
			t.Fatalf("check failed: %+v", c)
		}
	}
}
