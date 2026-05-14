package brokengaugefields

import "testing"

func TestFiniteBrokenGaugeFieldVariablesCurvatureSearch(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.BrokenFieldsTyped {
		t.Fatalf("broken fields should be typed")
	}
	if a.BrokenOnlyClosed {
		t.Fatalf("broken directions should not be closed without Q")
	}
	if !a.RequiresPhoton {
		t.Fatalf("electromagnetic Q direction should be required")
	}
	if !a.FullEWClosed {
		t.Fatalf("full electroweak broken-basis carrier should close diagnostically")
	}
	if a.CurvatureTermDerived {
		t.Fatalf("curvature term must remain open")
	}
	if a.Diag114ActionSelected {
		t.Fatalf("diag(1,1,4) must not be action-selected")
	}
}
