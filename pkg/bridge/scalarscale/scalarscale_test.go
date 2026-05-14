package scalarscale

import "testing"

func TestScaleBridgeDoesNotInventPhysicalUnits(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if a.FiniteRadiusSquared <= 0 || a.DimensionlessRadialMassSq <= 0 || a.BGap <= 0 || a.ContactLeakageNormSquared <= 0 {
		t.Fatalf("expected positive finite dimensionless anchors: %+v", a)
	}
	if a.HasDimensionfulAnchor || !a.OverallScaleFree {
		t.Fatalf("scale bridge should expose a free physical unit, not derive one")
	}
	if a.ElectroweakScaleDerived || a.HiggsMassBridgeDerived || a.UniqueScaleSelected {
		t.Fatalf("bridge must not claim physical scalar scale or Higgs mass")
	}
	if a.HiddenObservedScaleInserted {
		t.Fatalf("bridge search inserted an observed physical scale")
	}
}
