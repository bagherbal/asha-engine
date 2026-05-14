package fieldmap

import "testing"

func TestFiniteToContinuumFieldMapDefault(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if a.ActiveRealDirections != 4 {
		t.Fatalf("expected 4 active directions, got %d", a.ActiveRealDirections)
	}
	if a.ComplexDoubletComponents != 2 {
		t.Fatalf("expected 2 complex doublet components, got %d", a.ComplexDoubletComponents)
	}
	if !a.SectorLevelDoubletDerived || !a.ScalarPotentialDerived {
		t.Fatalf("expected sector-level doublet and scalar potential evidence")
	}
	if a.LowEnergyScaleDerived || a.PhysicalUnitDerived {
		t.Fatalf("low-energy/physical scale must remain open")
	}
	if a.RegulatorClassificationDerived || a.ThresholdCorrectionAllowed {
		t.Fatalf("regulator/threshold correction classification must not be derived")
	}
	if a.HiddenObservedInput {
		t.Fatalf("field-map audit must not insert observed inputs")
	}
	if a.PrimaryClassification != ContinuumScalarCandidate {
		t.Fatalf("unexpected primary classification %q", a.PrimaryClassification)
	}
}
