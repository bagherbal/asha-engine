package topkernel

import "testing"

func TestBuildDefault(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.UnitRightRowNormsEqual {
		t.Fatalf("expected equal unit incidence row norms")
	}
	if a.QuarkLeptonAmplification != 3 {
		t.Fatalf("expected three-color amplification 3, got %v", a.QuarkLeptonAmplification)
	}
	if !a.DiagonalGenerationKernelFound {
		t.Fatalf("expected diagonal generation kernel")
	}
	if a.TopLikeChannelSelected {
		t.Fatalf("top-like channel should not be selected yet")
	}
	if a.HiddenObservedCouplingsUsed {
		t.Fatalf("no observed couplings should be used")
	}
}
