package u1orientation

import (
	"math"
	"testing"
)

func TestChiralOrientationalAbelianSourceSearch(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault: %v", err)
	}
	if !a.Previous.LocalNonzeroCorrelation {
		t.Fatalf("expected Gate 77 local nonzero correlation")
	}
	if a.NaturalProbeCount < 8 {
		t.Fatalf("expected natural probes, got %d", a.NaturalProbeCount)
	}
	if a.NaturalNonzeroSources != 0 {
		t.Fatalf("expected natural orientational probes to cancel, got %d nonzero", a.NaturalNonzeroSources)
	}
	if a.ArbitraryNonzeroSources == 0 {
		t.Fatalf("expected non-canonical selector firewall to expose manufactured nonzero sources")
	}
	if a.CanonicalSourceDerived {
		t.Fatalf("canonical source should remain open")
	}
	if math.Abs(a.BestNaturalAbsSigned) > 1e-9 {
		t.Fatalf("best natural signed source should vanish, got %.12g", a.BestNaturalAbsSigned)
	}
}
