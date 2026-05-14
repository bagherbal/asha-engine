package spinor

import (
	"math"
	"testing"
)

func TestCovariantPhaseFockSpace(t *testing.T) {
	f, err := NewCovariantPhaseFockSpace(4)
	if err != nil {
		t.Fatal(err)
	}
	if f.ModeCount() != 4 || f.StateCount() != 16 {
		t.Fatalf("unexpected Fock dimensions: modes=%d states=%d", f.ModeCount(), f.StateCount())
	}
	vacuum, err := f.Vacuum()
	if err != nil {
		t.Fatal(err)
	}
	if !vacuum.IsSterileVacuumCandidate(1e-12) {
		t.Fatalf("vacuum should be sterile candidate, B-L=%v", vacuum.BMinusL())
	}
	quarkSeeds, leptonSeeds := 0, 0
	for _, s := range f.OneParticleStates() {
		if math.Abs(s.BMinusL()-(1.0/3.0)) < 1e-12 {
			quarkSeeds++
		}
		if math.Abs(s.BMinusL()+1.0) < 1e-12 {
			leptonSeeds++
		}
	}
	if quarkSeeds != 3 || leptonSeeds != 1 {
		t.Fatalf("expected 3 quark seeds and 1 lepton seed, got %d/%d", quarkSeeds, leptonSeeds)
	}
}
