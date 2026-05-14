package branchselector

import "testing"

func TestBuildDefault(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault() error = %v", err)
	}
	if !a.LocalBundleAttemptConstructed || !a.ConstraintComplexAttemptConstructed || !a.BranchSelectorAttempted {
		t.Fatalf("expected both branch attempts and selector attempt")
	}
	if a.ContactRows != 7 || a.PositiveFiniteContactRows != 7 || a.OpenContactRows != 7 {
		t.Fatalf("contact rows=%d positive=%d open=%d, want 7/7/7", a.ContactRows, a.PositiveFiniteContactRows, a.OpenContactRows)
	}
	if a.LocalBundleBranchCompleteRows != 0 || a.ConstraintComplexCompleteRows != 0 || a.BranchSelectorDerived || a.ResolvedContactRows != 0 || a.UnresolvedContactRows != 7 {
		t.Fatalf("selector must remain open: local=%d complex=%d selector=%t resolved=%d unresolved=%d", a.LocalBundleBranchCompleteRows, a.ConstraintComplexCompleteRows, a.BranchSelectorDerived, a.ResolvedContactRows, a.UnresolvedContactRows)
	}
	if a.BetaCorrectionRowsAllowed != 0 || a.ZeroContributionRowsProved != 0 || a.ThresholdCorrectedBetaDerived || a.FullFiniteBetaMatchingTensorDerived {
		t.Fatalf("beta tensor must remain sealed")
	}
	if a.ResidualNullityBefore != 3 || a.ResidualNullityAfter != 3 || a.ResidualSymmetryBroken {
		t.Fatalf("residual nullity changed: before=%d after=%d broken=%t", a.ResidualNullityBefore, a.ResidualNullityAfter, a.ResidualSymmetryBroken)
	}
	if a.PhysicalWeakAngleDerived || a.FineStructureDerived || a.PhysicalMassesDerived || a.PhysicalScaleDerived || a.HiddenObservedInputUsed {
		t.Fatalf("physical predictions/observed input must stay sealed")
	}
}

func TestTheoremPasses(t *testing.T) {
	r := ContactModeBranchSelectorConstructionAttemptTheorem().Run()
	if string(r.Status) != "VARIATIONAL" {
		t.Fatalf("status=%s", r.Status)
	}
	for _, c := range r.Checks {
		if !c.Passed {
			t.Fatalf("check %q failed: %s", c.Name, c.Detail)
		}
	}
}
