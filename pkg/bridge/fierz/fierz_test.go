package fierz

import "testing"

func TestFiniteFierzAudit(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if a.U4Dimension != 16 {
		t.Fatalf("u(4) dimension = %d, want 16", a.U4Dimension)
	}
	if !a.ScalarLRTargetAvailable || a.ScalarChannelCount != 12 {
		t.Fatalf("scalar channels = %d, available=%v; want 12/true", a.ScalarChannelCount, a.ScalarLRTargetAvailable)
	}
	if len(a.ProjectionSlots) != 4 {
		t.Fatalf("projection slots = %d, want 4", len(a.ProjectionSlots))
	}
	if a.NativeFierzProjectionComplete || a.AttractiveSignDerived || a.FourFermionStrengthDerived {
		t.Fatalf("Fierz audit should not claim completed projection/sign/G_hat")
	}
	if a.HiddenObservedInputUsed {
		t.Fatalf("audit must not use observed inputs")
	}
}
