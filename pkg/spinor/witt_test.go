package spinor

import "testing"

func TestNativeWittDecomposition(t *testing.T) {
	w, err := NativeWittDecomposition(4)
	if err != nil {
		t.Fatalf("NativeWittDecomposition: %v", err)
	}
	if w.RealDimension != 8 || w.ComplexModeCount != 4 || len(w.Pairs) != 4 {
		t.Fatalf("unexpected Witt dimensions: %+v", w)
	}
	for k, p := range w.Pairs {
		if p.ModeIndex != k || p.RealBasisA != 2*k || p.RealBasisB != 2*k+1 {
			t.Fatalf("pair %d not native consecutive pairing: %+v", k, p)
		}
		if p.BivectorLabel == "" || p.CreationFormula == "" || p.AnnihilationFormula == "" {
			t.Fatalf("pair %d missing formulas: %+v", k, p)
		}
	}
}
