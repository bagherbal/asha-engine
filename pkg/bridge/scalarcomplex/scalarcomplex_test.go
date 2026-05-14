package scalarcomplex

import "testing"

func TestScalarComplexQuaternionicStructure(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.PairCompatibleComplexAvailable {
		t.Fatalf("expected pair-compatible complex candidate")
	}
	if a.CanonicalComplexDerived {
		t.Fatalf("did not expect canonical complex orientation to be derived")
	}
	if !a.QuaternionicTripleAvailable {
		t.Fatalf("expected abstract quaternionic triple to be available")
	}
	if a.QuaternionicTripleSelected {
		t.Fatalf("did not expect anisotropic scalar response to select full quaternionic triple")
	}
	if a.FullScalarSU2Recovered {
		t.Fatalf("did not expect full scalar SU(2) theorem")
	}
}
