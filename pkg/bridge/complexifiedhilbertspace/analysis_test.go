package complexifiedhilbertspace

import "testing"

func TestBuildDefault(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault failed: %v", err)
	}
	if !a.Summary.ComplexificationDerived || !a.Summary.AntiLinearJAvailable || !a.Summary.MajoranaCapacity {
		t.Fatalf("expected complexification/J/Majorana capacity: %s", FormatSummary(a.Summary))
	}
	if a.Summary.NativeAlgebraDerived || a.Summary.BGapMajoranaIdentified || a.Summary.FullSpectralTripleDerived {
		t.Fatalf("Gate 235 must not derive full spectral triple: %s", FormatSummary(a.Summary))
	}
	if len(a.DoubledStates) != 32 {
		t.Fatalf("expected 32 real doubled bookkeeping states, got %d", len(a.DoubledStates))
	}
}

func TestComplexificationIsNotExternalDoubling(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault failed: %v", err)
	}
	if !a.Complexification.DerivedByComplexification || a.Complexification.ExternalStatesAdded {
		t.Fatalf("complexification audit failed: %s", FormatComplexification(a.Complexification))
	}
	if a.Complexification.RealDimensionBefore != 16 || a.Complexification.ComplexDimensionAfter != 16 || a.Complexification.RealDimensionAfter != 32 {
		t.Fatalf("unexpected dimensions: %s", FormatComplexification(a.Complexification))
	}
}

func TestFiniteAlgebraNotImported(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault failed: %v", err)
	}
	if a.Algebra.ImportedConnesAlgebra || a.Algebra.ColorM3CDerived || a.Algebra.QuaternionHDerived || a.Algebra.MaximalAssociativeSubalgebraDerived || a.Algebra.FaithfulDoubledRepresentation {
		t.Fatalf("finite algebra should remain un-derived: %s", FormatAlgebra(a.Algebra))
	}
}

func TestMajoranaCapacityButNoBGapIdentification(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault failed: %v", err)
	}
	if !a.Majorana.NeutralBilinearCapacity || a.Majorana.TotallyNeutralSlotCount < 1 {
		t.Fatalf("expected neutral bilinear capacity: %s", FormatMajorana(a.Majorana))
	}
	if a.Majorana.RHNeutrinoSlotDerived || a.BGap.BGapCanonicalMajoranaEntry || a.BGap.BGapPromotedToMass || a.BGap.BGapSelectsRHNeutrino {
		t.Fatalf("B-gap must not be identified as Majorana mass: %s :: %s", FormatMajorana(a.Majorana), FormatBGap(a.BGap))
	}
}

func TestTheorem(t *testing.T) {
	res := ComplexifiedHilbertSpaceFiniteAlgebraRepresentationAuditTheorem().Verify()
	if len(res.Checks) == 0 {
		t.Fatal("expected checks")
	}
	for _, c := range res.Checks {
		if !c.Passed {
			t.Fatalf("check %q failed: %s", c.Name, c.Detail)
		}
	}
}
