package seesawoverlapmatrix

import "testing"

func TestDoubledSpaceBlock(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault returned error: %v", err)
	}
	if !a.Block.Formalized || !a.Block.JSwapInstalled || a.Block.DirectSumOverlap != 0 || len(a.Block.Basis) != 3 {
		t.Fatalf("bad block: %s", FormatBlock(a.Block))
	}
}

func TestSeesawPathConstructed(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault returned error: %v", err)
	}
	if !a.Path.Constructed || a.Path.PathCount != 1 || a.Path.DirectSumPathCount != 0 || a.Path.PathMatrixRank != 1 {
		t.Fatalf("bad path: %s", FormatPath(a.Path))
	}
}

func TestOverlapIndexIsOne(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault returned error: %v", err)
	}
	if !a.Overlap.Derived || !a.Overlap.IndexVerified || a.Overlap.CanonicalOverlapIndex != 1 || a.Overlap.TraceOmegaDagOmega != 1 || a.Overlap.DirectSumIndex != 0 {
		t.Fatalf("bad overlap: %s", FormatOverlap(a.Overlap))
	}
}

func TestPortalWeightEnabledButNotPromoted(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault returned error: %v", err)
	}
	if !a.Portal.Enabled || !a.Portal.WeightsMultiplicative || !a.Portal.WithinOnePercent || a.Portal.Coefficient < 0.39 || a.Portal.Coefficient > 0.392 || a.Portal.ThresholdPromoted {
		t.Fatalf("bad portal: %s", FormatPortal(a.Portal))
	}
	if !a.Promotion.ExplicitMatrixDerived || !a.Promotion.OverlapIndexDerived || a.Promotion.HeavyPropagatorDerived || a.Promotion.HeavySelfQuarticDerived || a.Promotion.ThresholdJumpDerived {
		t.Fatalf("bad promotion: %s", FormatPromotion(a.Promotion))
	}
}

func TestFirewallsPreserved(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault returned error: %v", err)
	}
	if !a.Firewalls.NoFinalMassClaimed || !a.Firewalls.NoThresholdClaimed || !a.Firewalls.NoHeavyPropagatorClaimed || !a.Firewalls.NoHeavyQuarticClaimed || !a.Firewalls.NoLambdaMixClaimed || a.Firewalls.FiniteCorePolluted {
		t.Fatalf("bad firewalls: %s", FormatFirewalls(a.Firewalls))
	}
}

func TestTheoremPasses(t *testing.T) {
	res := SeesawOverlapMatrixConstructionMajoranaHiggsMixingSieveAuditTheorem().Run()
	if !res.Passed() {
		t.Fatalf("theorem did not pass:\n%s", res.Details())
	}
}
