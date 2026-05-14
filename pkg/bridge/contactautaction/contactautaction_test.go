package contactautaction

import "testing"

func TestContactSideAutomorphismActionEquivariantAssignmentSearch(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.SymmetrySelectorObstructionInherited || !a.FanoAutomorphismGroupDerived || a.FanoAutomorphismGroupOrder != 168 {
		t.Fatalf("expected Gate 118 symmetry obstruction and |Aut(Fano)|=168")
	}
	if !a.ContactWeightedAutomorphismGroupDerived || a.ContactWeightedAutomorphismGroupOrder != 1 || !a.ContactWeightedActionIdentityOnly || !a.ContactSpectralValuesAllDistinct {
		t.Fatalf("expected identity-only contact weighted automorphism group, got %s", FormatSummary(a.Summary))
	}
	if !a.OrderMismatchObstructionDerived || a.AutFanoActionOnContactDerived || a.AutFanoActionPreservingContactData {
		t.Fatalf("expected order mismatch obstruction without a canonical Aut(Fano) contact action")
	}
	if !a.TransportedFanoActionsPossibleAfterChoice || a.Summary.TransportedActionCount != 5040 || a.TransportedFanoActionCanonical || a.EquivariantAssignmentDerived || a.CanonicalContactFanoAssignmentDerived {
		t.Fatalf("transported actions should be noncanonical: %s", FormatSummary(a.Summary))
	}
	if a.ContactRows != 7 || a.OpenContactRowsAfter != 7 || a.RepresentationCompleteRows != 0 || a.RepresentationOpenRows != 7 || a.ContactBetaRowsAllowed != 0 || a.ContactZeroRowsProved != 0 || a.ThresholdCorrectedBetaDerived {
		t.Fatalf("contact beta firewall should remain closed")
	}
	if a.ResidualNullityAfter != 3 || a.PhysicalWeakAngleDerived || a.FineStructureDerived || a.PhysicalMassesDerived || a.PhysicalScaleDerived || a.HiddenObservedInputUsed {
		t.Fatalf("physical bridge leaked")
	}
}
