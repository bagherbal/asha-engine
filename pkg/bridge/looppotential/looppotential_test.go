package looppotential

import "testing"

func TestNativeOneLoopPotentialLedger(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault() error: %v", err)
	}
	if a.ColorAmplificationFactor != 3 {
		t.Fatalf("expected three-color amplification, got %d", a.ColorAmplificationFactor)
	}
	if a.TopLikeCoefficientSkeleton != -6 {
		t.Fatalf("expected -6 skeleton, got %d", a.TopLikeCoefficientSkeleton)
	}
	if !a.StructuralInstabilityPressureAvailable {
		t.Fatalf("expected structural instability pressure ledger to be available")
	}
	if a.MuSquaredSignDerived || a.NativeEffectivePotentialComputed {
		t.Fatalf("ledger must not claim native μ² sign or full effective potential")
	}
	if a.ImportedSMRGE || a.HiddenObservedCouplingsUsed {
		t.Fatalf("ledger must not import SM RGE table or observed couplings")
	}
}
