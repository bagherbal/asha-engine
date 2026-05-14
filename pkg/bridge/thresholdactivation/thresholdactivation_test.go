package thresholdactivation

import "testing"

func TestThresholdActivationDefault(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if a.CandidateCount == 0 {
		t.Fatalf("expected threshold candidates")
	}
	if !a.ScalarSectorRemainsContinuumCandidate {
		t.Fatalf("scalar/contact sector should remain a continuum-field candidate")
	}
	if !a.LeakageClassifiedAsVacuumOnly {
		t.Fatalf("leakage should be classified as vacuum-frustration-only")
	}
	if a.BGapActivationDerived {
		t.Fatalf("B-sector activation rule must remain open")
	}
	if a.ContactOverlapActivationDerived {
		t.Fatalf("contact overlap activation rule must remain open")
	}
	if a.PhysicalMassUnitDerived || a.DecouplingRuleDerived || a.ThresholdCorrectedBetaDerived {
		t.Fatalf("physical unit, decoupling rule, and threshold-corrected beta coefficients must remain open")
	}
	if a.HiddenScaleInserted {
		t.Fatalf("activation audit must not insert hidden scales")
	}
}
