package resolventbranchsemantics

import "testing"

func TestGate281MultiplicityDoesNotOrientProjectors(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatalf("Build failed: %v", err)
	}
	if a.TraceNorm.BranchCount != 3 || a.TraceNorm.PossibleOrientations != 6 {
		t.Fatalf("unexpected branch/orientation count: %+v", a.TraceNorm)
	}
	if a.TraceNorm.AnyNativeOrientationPreferred {
		t.Fatalf("unexpected native orientation preference: %s", FormatTraceNorm(a.TraceNorm))
	}
	if !allBranchesRankTwo(a.TraceNorm) {
		t.Fatalf("expected all projectors to remain rank-2/rank-2: %s", FormatTraceNorm(a.TraceNorm))
	}
}

func TestGate281OrientationSealDoesNotLockRBranch(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatalf("Build failed: %v", err)
	}
	if !a.OrientationSeal.Active || !a.OrientationSeal.GrantsProjectorSectorMap || a.OrientationSeal.GrantsAmplitudeBranchMap {
		t.Fatalf("unexpected seal status: %+v", a.OrientationSeal)
	}
	if !a.RBranch.OrientationLocked || a.RBranch.UniqueAmplitudeBranch || a.RBranch.AlgebraicResolventToRMapDerived {
		t.Fatalf("unexpected r-branch over/under state: %+v", a.RBranch)
	}
	if err := AssertNoOverclaim(a); err != nil {
		t.Fatalf("overclaim: %v", err)
	}
}

func TestGate281TheoremChecksPass(t *testing.T) {
	th := ResolventBranchSemanticsProjectorSectorOrientationSealAuditTheorem()
	res := th.Verify()
	if res.Status == "FAILED_ROUTE" {
		t.Fatalf("theorem build failed: %+v", res)
	}
	for _, c := range res.Checks {
		if !c.Passed {
			t.Fatalf("check %q failed: %s", c.Name, c.Detail)
		}
	}
}
