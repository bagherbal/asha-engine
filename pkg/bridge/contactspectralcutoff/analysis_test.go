package contactspectralcutoff

import (
	"math"
	"testing"
)

func TestContactSpectralCutoffMoments(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Cutoff.IdentifiedAsF0F2F4 {
		t.Fatal("contact moments were not identified as cutoff audit inputs")
	}
	if math.Abs(a.Cutoff.Zeta2-61.0/25.0) > 1e-15 {
		t.Fatalf("wrong zeta2: %.17g", a.Cutoff.Zeta2)
	}
	if math.Abs(a.Cutoff.Zeta4-257629.0/202500.0) > 1e-15 {
		t.Fatalf("wrong zeta4: %.17g", a.Cutoff.Zeta4)
	}
}

func TestBothBranchesSurvivePositiveX(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Sieve.BothBranchesSurvive || a.Sieve.UniqueBranchSelected {
		t.Fatalf("expected both branches to survive without unique selection: %+v", a.Sieve)
	}
	for _, sol := range a.Sieve.Solutions {
		if !sol.Survives || len(sol.PositiveRoots) != 1 {
			t.Fatalf("branch %s did not have exactly one positive root: %+v", sol.Branch.Name, sol.PositiveRoots)
		}
		if sol.SelectedX <= 0 || math.IsNaN(sol.SelectedX) || math.IsInf(sol.SelectedX, 0) {
			t.Fatalf("invalid positive X for %s: %g", sol.Branch.Name, sol.SelectedX)
		}
	}
}

func TestTraceMomentsAreBranchIndependent(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if !a.MomentLock.BranchIndependentD2D4 || !a.MomentLock.BranchDependentX {
		t.Fatalf("expected branch-independent D2/D4 but branch-dependent X: %+v", a.MomentLock)
	}
	if math.Abs(a.MomentLock.Shape-lambda) > 1e-12 {
		t.Fatalf("shape drift: got %.17g want %.17g", a.MomentLock.Shape, lambda)
	}
	if math.Abs(a.Sieve.Solutions[0].D2-a.Sieve.Solutions[1].D2) > 1e-10 {
		t.Fatalf("D2 not branch independent")
	}
	if math.Abs(a.Sieve.Solutions[0].D4-a.Sieve.Solutions[1].D4) > 1e-10 {
		t.Fatalf("D4 not branch independent")
	}
}

func TestNoHiggsPredictionClaimed(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if a.Higgs.HiggsMassRatioClaimed || a.Summary.HiggsPredictionDerived {
		t.Fatal("Gate 288 must not claim a Higgs prediction")
	}
	if !a.Firewalls.DoesNotDiscardSurvivingBranch || a.Firewalls.FiniteCorePolluted {
		t.Fatalf("firewall failure: %+v", a.Firewalls)
	}
}
