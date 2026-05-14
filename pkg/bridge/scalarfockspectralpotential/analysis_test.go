package scalarfockspectralpotential

import (
	"math"
	"testing"
)

func TestBuildDefaultScalarFockComparison(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Firewall.GaugeRatioAlreadyClosed {
		t.Fatalf("Gate167 gauge ratio should be closed input: %+v", a.Firewall)
	}
	if a.UnitYukawa.ChannelCount != 8 || a.UnitYukawa.QuadraticMomentA != 8 || a.UnitYukawa.QuarticMomentB != 8 {
		t.Fatalf("bad unit Yukawa moments: %+v", a.UnitYukawa)
	}
	if a.UnitYukawa.FullDiracTraceD2 != 16 || a.UnitYukawa.FullDiracTraceD4 != 16 {
		t.Fatalf("bad doubled Dirac traces: %+v", a.UnitYukawa)
	}
}

func TestUnitIncidenceShapeDoesNotMatchContactShape(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if math.Abs(a.UnitYukawa.ChiralShape-0.125) > 1e-12 {
		t.Fatalf("unit chiral shape=%g, want 1/8", a.UnitYukawa.ChiralShape)
	}
	if a.Comparison.UnitIncidenceMatchesContact || a.Comparison.ConvergenceClosed {
		t.Fatalf("unit shape should not close scalar convergence: %+v", a.Comparison)
	}
	if a.Comparison.AbsoluteDifference <= 0.13 {
		t.Fatalf("unexpectedly small scalar shape difference: %+v", a.Comparison)
	}
}

func TestContactShapeIsAllowedYukawaConstraint(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Comparison.ContactWithinYukawaShapeRange || !a.Comparison.ConstraintOpened {
		t.Fatalf("Gate37 shape should open an allowed amplitude constraint: %+v", a.Comparison)
	}
	if !(a.Comparison.ContactEffectiveSlots > 3 && a.Comparison.ContactEffectiveSlots < 4) {
		t.Fatalf("contact effective slots=%g, expected between 3 and 4", a.Comparison.ContactEffectiveSlots)
	}
}

func TestScalarShapeIsAmplitudeSensitive(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if math.Abs(a.DeformedYukawa.ChiralShape-(53.0/289.0)) > 1e-12 {
		t.Fatalf("deformed shape=%g, want 53/289", a.DeformedYukawa.ChiralShape)
	}
	if a.DeformedYukawa.ChiralShape == a.UnitYukawa.ChiralShape {
		t.Fatalf("scalar shape should change under Yukawa-amplitude deformation")
	}
	if !a.SpectralAction.ScalarShapeAmplitudeDependent || a.SpectralAction.GaugeLikeRepresentationRigidity {
		t.Fatalf("bad scalar spectral action audit: %+v", a.SpectralAction)
	}
}
