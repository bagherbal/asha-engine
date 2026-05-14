package contactnaturality

import "testing"

func TestContactFanoNaturalityAutomorphismObstruction(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.IncidenceFunctorObstructionInherited || a.ContactRows != 7 || a.RepresentationOpenRows != 7 {
		t.Fatalf("expected Gate 116 incidence-open contact rows")
	}
	if !a.FanoAutomorphismGroupDerived || a.FanoAutomorphismGroupOrder != 168 || a.AutomorphismSummary.IdentityCount != 1 || a.AutomorphismSummary.NonIdentityCount != 167 {
		t.Fatalf("expected full 168-element Fano automorphism group, got order=%d identity=%d nonidentity=%d", a.FanoAutomorphismGroupOrder, a.AutomorphismSummary.IdentityCount, a.AutomorphismSummary.NonIdentityCount)
	}
	if !a.FanoPointActionTransitive || !a.FanoLineActionTransitive || a.GlobalFixedFanoPoints != 0 || a.GlobalFixedFanoLines != 0 {
		t.Fatalf("expected transitive point/line actions with no global fixed selector")
	}
	if a.ContactAutomorphismActionDerived || a.NaturalitySquareFormulable || a.InvariantContactToFanoMapDerived || a.EquivariantBijectionDerived {
		t.Fatalf("unexpected contact action/naturality/equivariant bijection derived")
	}
	if a.CanonicalAssignmentCount != 0 || a.CompatibleBijectionCount != 5040 || a.ConventionDependentBijections != 5040 {
		t.Fatalf("expected all 7! assignments to remain convention-dependent")
	}
	if a.RepresentationCompleteRows != 0 || a.RepresentationOpenRows != 7 || a.ContactBetaRowsAllowed != 0 || a.ContactZeroRowsProved != 0 || a.ThresholdCorrectedBetaDerived {
		t.Fatalf("contact beta firewall should remain closed")
	}
	if a.ResidualNullityAfter != 3 || a.PhysicalWeakAngleDerived || a.FineStructureDerived || a.PhysicalMassesDerived || a.PhysicalScaleDerived || a.HiddenObservedInputUsed {
		t.Fatalf("physical bridge leaked")
	}
}
