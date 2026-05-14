package condensate

import "testing"

func TestCompositeHiggsCondensateDirectionAudit(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault failed: %v", err)
	}
	if a.FockStates != 16 || a.FockModes != 4 {
		t.Fatalf("unexpected Fock substrate: states=%d modes=%d", a.FockStates, a.FockModes)
	}
	if a.SpatialModes != 3 || a.TemporalModes != 1 {
		t.Fatalf("expected 1+3 Fock split, got temporal=%d spatial=%d", a.TemporalModes, a.SpatialModes)
	}
	if !a.BilinearScalarCandidateAvailable {
		t.Fatalf("expected gauge-compatible bilinear scalar candidate")
	}
	if !a.ColorAmplificationAvailable {
		t.Fatalf("expected three-color amplification to be available")
	}
	if a.NativeOneLoopPotentialComputed || a.NJLGapEquationSolved || a.CondensationScaleDerived {
		t.Fatalf("condensate dynamics should remain open at this gate")
	}
}
