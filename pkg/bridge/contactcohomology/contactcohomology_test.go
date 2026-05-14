package contactcohomology

import "testing"

func TestContactCohomologyObstruction(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if a.ContactRows != 7 || a.PositiveFiniteContactRows != 7 || a.OpenContactRowsAfter != 7 {
		t.Fatalf("expected seven unresolved positive contact modes, got contact=%d positive=%d open=%d", a.ContactRows, a.PositiveFiniteContactRows, a.OpenContactRowsAfter)
	}
	if !a.ZeroDifferentialSquareZero || a.ZeroDifferentialCohomologyDimension != 7 || a.ZeroDifferentialProvesCancellation {
		t.Fatalf("zero differential audit invalid: Q²=%t H=%d cancellation=%t", a.ZeroDifferentialSquareZero, a.ZeroDifferentialCohomologyDimension, a.ZeroDifferentialProvesCancellation)
	}
	if a.CanonicalDifferentialDerived || a.NontrivialNilpotentDifferentialDerived || a.CancellationLedgerDerived {
		t.Fatalf("unexpected derived constraint complex")
	}
	if a.ContactZeroRowsProved != 0 || a.ContactBetaRowsAllowed != 0 || a.ThresholdCorrectedBetaDerived {
		t.Fatalf("contact beta firewall should remain closed")
	}
	if a.ResidualNullityAfter != 3 || a.PhysicalWeakAngleDerived || a.FineStructureDerived || a.PhysicalMassesDerived || a.PhysicalScaleDerived {
		t.Fatalf("physical bridge leaked")
	}
}
