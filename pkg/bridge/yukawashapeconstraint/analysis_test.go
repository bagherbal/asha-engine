package yukawashapeconstraint

import (
	"math"
	"testing"
)

func TestBuildDefaultYukawaShapeConstraint(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Firewall.GaugeRatioClosed || !a.Firewall.ScalarShapeTargetAvailable {
		t.Fatalf("expected closed gauge input and scalar target: %+v", a.Firewall)
	}
	if math.Abs(a.Target.FloatValue-(1197.0/4624.0)) > 1e-10 {
		t.Fatalf("target shape=%g, want 1197/4624", a.Target.FloatValue)
	}
	if math.Abs(a.Target.RequiredEffectiveSlots-(4624.0/1197.0)) > 1e-10 {
		t.Fatalf("effective slots=%g, want 4624/1197", a.Target.RequiredEffectiveSlots)
	}
}

func TestFourClassContactPatternMatchesButEightChannelDoesNot(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if a.Candidates[0].MatchesTarget {
		t.Fatalf("unit eight-channel incidence should not match")
	}
	if a.Candidates[2].MatchesTarget {
		t.Fatalf("direct duplicated contact spectrum should not match")
	}
	if !a.Best.MatchesTarget || a.Best.SlotCount != 4 {
		t.Fatalf("four-class contact pattern should conditionally match: %+v", a.Best)
	}
	if math.Abs(a.Candidates[2].Shape-a.Target.FloatValue/2) > 1e-10 {
		t.Fatalf("duplicated shape=%g, want half target %g", a.Candidates[2].Shape, a.Target.FloatValue/2)
	}
}

func TestPairCollapseAndKindAssignmentRemainOpen(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.PairCollapse.FourClassQuotientAvailable || a.PairCollapse.FourClassQuotientDerived {
		t.Fatalf("pair quotient availability/derivation mismatch: %+v", a.PairCollapse)
	}
	if a.PairCollapse.KindAssignmentAmbiguity != 6 {
		t.Fatalf("kind assignment ambiguity=%d, want 6", a.PairCollapse.KindAssignmentAmbiguity)
	}
	if a.Firewall.PairCollapseDerived || a.Firewall.KindAssignmentDerived || a.Firewall.YukawaAmplitudesDerived {
		t.Fatalf("firewall should remain closed: %+v", a.Firewall)
	}
}

func TestGenerationTextureStillUnderdetermined(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if a.Generation.GenerationCount != 3 || a.Generation.KindTextureMatrices != 4 || a.Generation.TotalGeneralTextureEntries != 36 {
		t.Fatalf("bad generation texture audit: %+v", a.Generation)
	}
	if !a.Generation.TextureUnderdetermined || a.Generation.FermionMassesDerived || a.Generation.CKMPMNSDerived {
		t.Fatalf("mass/mixing should remain underived: %+v", a.Generation)
	}
}
