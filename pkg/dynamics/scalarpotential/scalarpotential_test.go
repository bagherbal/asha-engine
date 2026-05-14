package scalarpotential

import "testing"

func TestBuildDefaultScalarEffectivePotential(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault failed: %v", err)
	}
	if a.ActiveRealDimension != 4 {
		t.Fatalf("active dimension=%d, want 4", a.ActiveRealDimension)
	}
	if a.ComplexDoubletDimension != 2 {
		t.Fatalf("complex doublet dimension=%d, want 2", a.ComplexDoubletDimension)
	}
	if a.ProtectedDirectionCount != 3 {
		t.Fatalf("protected directions=%d, want 3", a.ProtectedDirectionCount)
	}
	if !a.PairDegenerate {
		t.Fatalf("expected pair-degenerate active spectrum")
	}
	if a.VacuumRadiusSquared <= 0 || a.LambdaShape <= 0 || a.DimensionlessRadialMassSq <= 0 {
		t.Fatalf("expected positive radius, lambda shape, and radial curvature")
	}
	if a.FiniteTachyonicMassDerived || a.ElectroweakScaleDerived || a.HiggsMassDerived {
		t.Fatalf("physical scalar scale/mass claims must remain bridge-only")
	}
}
