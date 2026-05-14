package bfbridge

import "testing"

func TestActiveGenerationProjectionBridge(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if a.ProtectedDimension != 3 {
		t.Fatalf("protected dimension = %d", a.ProtectedDimension)
	}
	if a.ActiveDimension != 4 {
		t.Fatalf("active dimension = %d", a.ActiveDimension)
	}
	if a.ActiveCurvatureNorm <= 1e-8 {
		t.Fatalf("expected nonzero active curvature")
	}
	if a.ExistingConnectionBridgeFound {
		t.Fatalf("existing active-to-protected bridge should not be present in current finite data")
	}
	if a.CanonicalGenerationMixingFound {
		t.Fatalf("canonical generation mixing should remain open")
	}
}
