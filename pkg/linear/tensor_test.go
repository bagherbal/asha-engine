package linear

import "testing"

func TestKroneckerDimensionsAndTrace(t *testing.T) {
	a := Diagonal([]float64{1, 2})
	b := Diagonal([]float64{3, 4, 5})
	k := Kronecker(a, b)
	if k.Rows() != 6 || k.Cols() != 6 {
		t.Fatalf("unexpected Kronecker dimensions: %dx%d", k.Rows(), k.Cols())
	}
	tr, err := k.Trace()
	if err != nil {
		t.Fatal(err)
	}
	if tr != (1+2)*(3+4+5) {
		t.Fatalf("unexpected Kronecker trace: %g", tr)
	}
}
