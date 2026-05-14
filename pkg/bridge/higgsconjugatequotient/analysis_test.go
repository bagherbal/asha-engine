package higgsconjugatequotient

import "testing"

func TestGate25BranchesAreUniqueByKind(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if a.HiggsAudit.Gate25MinimalChannels != 8 || a.HiggsAudit.FermionKindBlocks != 4 {
		t.Fatalf("bad channel/kind counts: %+v", a.HiggsAudit)
	}
	if a.HiggsAudit.KindsWithUniqueBranch != 4 || a.HiggsAudit.KindsWithBothBranches != 0 {
		t.Fatalf("expected unique branch per kind and no both-branch kind: %+v groups=%s", a.HiggsAudit, FormatGroups(a.Groups))
	}
	if !a.HiggsAudit.HyperchargeSelectsUniqueBranch {
		t.Fatalf("expected hypercharge to select unique branch: %+v", a.HiggsAudit)
	}
}

func TestHiggsConjugateQuotientRejected(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if a.HiggsAudit.HiggsConjugatePairsAvailable || a.HiggsAudit.HiggsConjugatePairCollapse {
		t.Fatalf("Higgs-conjugate quotient should be rejected: %+v", a.HiggsAudit)
	}
	if !a.Consequence.HiggsConjugatePremiseRejected {
		t.Fatalf("consequence should mark Higgs premise rejected: %+v", a.Consequence)
	}
}

func TestFourKindSupportQuotientVisibleButNotAmplitudeDerived(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	q := a.KindQuotient
	if !q.FourKindSupportQuotientVisible || !q.FourKindSupportQuotientCanonical {
		t.Fatalf("four-kind support quotient should be visible/canonical: %+v", q)
	}
	if q.UpColorChannels != 3 || q.DownColorChannels != 3 || q.NeutrinoChannels != 1 || q.ElectronChannels != 1 {
		t.Fatalf("unexpected support pattern: %+v", q)
	}
	if q.ColorAmplitudeUniversalityDerived || q.FourAmplitudeClassQuotientDerived {
		t.Fatalf("amplitude quotient should not be derived here: %+v", q)
	}
}

func TestScalarShapeAndMassFirewallRemainOpen(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Consequence.Gate169ConditionalMatchFound || a.Consequence.ContactWeightAssignments != 6 {
		t.Fatalf("Gate169 conditional target should survive with six assignments: %+v", a.Consequence)
	}
	if a.Consequence.CanonicalContactKindAssignment || a.Consequence.ScalarShapeClosed || a.Consequence.AmplitudeTextureSelected {
		t.Fatalf("scalar shape must remain open: %+v", a.Consequence)
	}
	if a.Firewall.YukawaAmplitudesDerived || a.Firewall.FermionMassesDerived || a.Firewall.PhysicalConstantsDerived {
		t.Fatalf("firewall should remain closed: %+v", a.Firewall)
	}
}
