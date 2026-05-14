package nativefinitealgebra

import "testing"

func TestGate236NativeFiniteAlgebraAudit(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault failed: %v", err)
	}
	if !a.Split.ModeLevelProjectionExists || a.Split.LeptonLikeGeneratorCount != 1 || a.Split.ColorLikeGeneratorCount != 3 {
		t.Fatalf("expected native 1⊕3 mode split, got %+v", a.Split)
	}
	if !a.Commutant.ColorMatrixAlgebraPreflight || a.Commutant.ModeProjectionCommutantDimensionC != 10 {
		t.Fatalf("expected C⊕M3(C) commutant preflight, got %+v", a.Commutant)
	}
	if a.Contact.HGenerated || a.Algebra.ExactCPlusHPlusM3Derived || a.Algebra.ConnesAlgebraImported {
		t.Fatalf("Connes algebra/H should not be derived or imported: contact=%+v algebra=%+v", a.Contact, a.Algebra)
	}
	if a.Firewall.FiniteCorePolluted || a.Firewall.InsertedSMGaugeGroup || a.Firewall.ClaimedSMAlgebraDerivation {
		t.Fatalf("firewall violated: %+v", a.Firewall)
	}
}
