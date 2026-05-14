package higgs

import "testing"

func TestVacuumMixingSector(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault() error: %v", err)
	}
	if len(a.HiggsFields) == 0 {
		t.Fatal("expected non-empty Higgs field list")
	}
	if a.HiggsSpanRank == 0 {
		t.Fatal("expected nonzero Higgs span rank")
	}
	if a.MaxOffDiagonalBlockResidual > 1e-8 {
		t.Fatalf("field is not purely off-diagonal: %.3e", a.MaxOffDiagonalBlockResidual)
	}
	if a.MaxSkewResidual > 1e-8 {
		t.Fatalf("field is not skew: %.3e", a.MaxSkewResidual)
	}
	if a.MaxPositiveResidual > 1e-8 {
		t.Fatalf("mixing operator is not positive: %.3e", a.MaxPositiveResidual)
	}
	if a.VacuumMixingRank == 0 {
		t.Fatal("expected nonzero contact-vacuum mixing rank")
	}
	if a.VacuumMixingRank >= a.ContactDimension {
		t.Fatalf("expected partial mixing with protected residue, got rank %d of %d", a.VacuumMixingRank, a.ContactDimension)
	}
}
