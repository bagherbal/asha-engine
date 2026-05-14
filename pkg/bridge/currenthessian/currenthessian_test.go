package currenthessian

import "testing"

func TestBuildDefault(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if a.SectorCount != 4 || a.GeneratorCount != 16 {
		t.Fatalf("unexpected current domain: sectors=%d generators=%d", a.SectorCount, a.GeneratorCount)
	}
	if a.SectorHessianDimension != 10 || a.GeneratorHessianDimension != 136 {
		t.Fatalf("unexpected Hessian dimensions: sector=%d generator=%d", a.SectorHessianDimension, a.GeneratorHessianDimension)
	}
	if a.CandidateCount != 4 {
		t.Fatalf("candidate count=%d", a.CandidateCount)
	}
	if !a.AllCandidatesPositive || !a.AllCandidatesFiniteDataOnly {
		t.Fatalf("candidate positivity/finite discipline failed")
	}
	if a.AnyCandidateSelected || a.CurrentHessianDerived || a.PropagatorRuleDerived || a.CondensationClaimAllowed {
		t.Fatalf("Gate 69 must not select a Hessian or claim condensation")
	}
}
