package currentcontact

import "testing"

func TestBuildDefault(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault() error = %v", err)
	}
	if a.SourceGeneratorCount != 16 {
		t.Fatalf("source generator count = %d, want 16", a.SourceGeneratorCount)
	}
	if a.TargetBlockSeedCount != 4 || a.TargetBlockSpanRank != 4 {
		t.Fatalf("target block = %d rank=%d, want 4 rank 4", a.TargetBlockSeedCount, a.TargetBlockSpanRank)
	}
	if !a.AbstractMapSpaceExists || a.AbstractMapDimension != 64 {
		t.Fatalf("abstract map dim = %d, exists=%v, want 64 true", a.AbstractMapDimension, a.AbstractMapSpaceExists)
	}
	if !a.AbelianAmbiguity {
		t.Fatalf("expected abelian ambiguity")
	}
	if a.ColorSectorCarrierDerived || a.LeptoquarkCarrierDerived || a.CurrentToContactMapDerived {
		t.Fatalf("unexpected derived current-to-contact carrier/map")
	}
	if a.HessianComputable || a.ExchangeKernelUpdated || a.CondensationClaimAllowed {
		t.Fatalf("gate must not allow Hessian/kernel/condensation")
	}
}
