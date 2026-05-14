package betapermission

import "testing"

func TestBuildDefault(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault() error = %v", err)
	}
	if !a.BetaPermissionFirewallConstructed || !a.PhysicalBranchRuleConstructed || !a.ConstraintBranchRuleConstructed {
		t.Fatalf("firewall/rules not constructed")
	}
	if a.ContactRows != 7 || a.PositiveFiniteContactRows != 7 || a.DichotomyOpenRows != 7 {
		t.Fatalf("contact rows=%d positive=%d open=%d, want 7/7/7", a.ContactRows, a.PositiveFiniteContactRows, a.DichotomyOpenRows)
	}
	if a.ResolvedContactRows != 0 || a.UnresolvedContactRows != 7 || a.AllContactModesResolved || a.RepresentationOrConstraintDichotomyDerived {
		t.Fatalf("contact dichotomy must remain unresolved: resolved=%d unresolved=%d all=%t derived=%t", a.ResolvedContactRows, a.UnresolvedContactRows, a.AllContactModesResolved, a.RepresentationOrConstraintDichotomyDerived)
	}
	if a.PhysicalBranchCompleteRows != 0 || a.ConstraintBranchCompleteRows != 0 || a.BetaCorrectionRowsAllowed != 0 || a.ZeroContributionRowsProved != 0 {
		t.Fatalf("branches/beta permission must remain closed: physical=%d constraint=%d beta=%d zero=%d", a.PhysicalBranchCompleteRows, a.ConstraintBranchCompleteRows, a.BetaCorrectionRowsAllowed, a.ZeroContributionRowsProved)
	}
	if a.ThresholdCorrectedBetaDerived || a.FullFiniteBetaMatchingTensorDerived {
		t.Fatalf("threshold beta tensor must remain sealed")
	}
	if a.ResidualNullityBefore != 3 || a.ResidualNullityAfter != 3 || a.ResidualSymmetryBroken {
		t.Fatalf("residual nullity changed: before=%d after=%d broken=%t", a.ResidualNullityBefore, a.ResidualNullityAfter, a.ResidualSymmetryBroken)
	}
	if a.PhysicalWeakAngleDerived || a.FineStructureDerived || a.PhysicalMassesDerived || a.PhysicalScaleDerived || a.HiddenObservedInputUsed {
		t.Fatalf("physical predictions/observed input must stay sealed")
	}
}

func TestTheoremPasses(t *testing.T) {
	r := ContactOverlapRepresentationConstraintBetaPermissionFirewallTheorem().Run()
	if string(r.Status) != "VARIATIONAL" {
		t.Fatalf("status=%s", r.Status)
	}
	for _, c := range r.Checks {
		if !c.Passed {
			t.Fatalf("check %q failed: %s", c.Name, c.Detail)
		}
	}
}
