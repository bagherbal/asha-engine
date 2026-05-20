package generation2koideloopdeficitreactorangleaudit

import (
	"math"
	"testing"
)

func TestGate588KoideLoopDeficitReactorAngleAudit(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if math.Abs(a.Runtime.Kappa-0.005503554191574556) > 1e-15 {
		t.Fatalf("unexpected kappa: %s", FormatRuntime(a.Runtime))
	}
	if math.Abs(a.Candidate.Value-0.0055375) > 1e-16 {
		t.Fatalf("unexpected reactor candidate: %s", FormatCandidate(a.Candidate))
	}
	if math.Abs(a.Candidate.RelativeResidual-0.00616797931733138) > 1e-14 {
		t.Fatalf("unexpected reactor residual: %s", FormatCandidate(a.Candidate))
	}
	if !a.Candidate.CoversKappa || a.Candidate.Certified {
		t.Fatalf("reactor candidate should cover but not certify: %s", FormatCandidate(a.Candidate))
	}
	if math.Abs(a.Inverse.Sin2Theta13Pred-0.022014216766298222) > 1e-15 {
		t.Fatalf("unexpected inverse sin2 prediction: %s", FormatInverse(a.Inverse))
	}
	if math.Abs(a.Inverse.Theta13PredDeg-8.532586786085982) > 1e-12 {
		t.Fatalf("unexpected inverse theta prediction: %s", FormatInverse(a.Inverse))
	}
	if !a.Inverse.WithinSin2OneSigma || !a.Inverse.WithinThetaOneSigma {
		t.Fatalf("inverse prediction should be inside one sigma: %s", FormatInverse(a.Inverse))
	}
	if math.Abs(a.Epsilon.SignedResidualRad-(-1.350660802035275e-06)) > 1e-16 {
		t.Fatalf("unexpected epsilon residual: %s", FormatEpsilon(a.Epsilon))
	}
	if !a.Comparison.BeatsPriorPMNSAssisted || !a.Comparison.BeatsSqrtJCKM || a.Comparison.BeatsCKMAlpha2Midpoint {
		t.Fatalf("unexpected comparison: %s", FormatComparison(a.Comparison))
	}
	if a.Operator.NativeLeptonOrientationOperatorPresent || a.Operator.DerivesKappa || a.Firewalls.PromotesObservedAsNative || !a.Firewalls.PreservesGate352 {
		t.Fatalf("firewall/operator broken: %s / %s", FormatOperator(a.Operator), FormatFirewalls(a.Firewalls))
	}
}

func TestGate588Theorem(t *testing.T) {
	res := Generation2KoideLoopDeficitReactorAngleAuditTheorem().Run()
	if res.Status == "FAILED_ROUTE" {
		t.Fatalf("unexpected failed route: %+v", res)
	}
	for _, c := range res.Checks {
		if !c.Passed {
			t.Fatalf("check failed: %+v", c)
		}
	}
}
