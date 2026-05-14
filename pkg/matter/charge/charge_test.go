package charge

import (
	"math"
	"testing"
)

func TestBLChargePolarizesOnePlusThree(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.ChargePolarizesOnePlusThree {
		t.Fatalf("expected B-L spectrum to polarize 1+3, got clusters %v", a.OneParticleChargeClusters)
	}
	if math.Abs(a.TraceOneParticleCharge) > 1e-10 {
		t.Fatalf("expected traceless one-particle B-L charge, got %g", a.TraceOneParticleCharge)
	}
	if math.Abs(a.TraceOneParticleChargeSquared-4.0/3.0) > 1e-10 {
		t.Fatalf("unexpected B-L squared trace: %g", a.TraceOneParticleChargeSquared)
	}
	if a.DirectScalarToColorIsotropyPossible {
		t.Fatalf("active Higgs/contact weights should not be directly color-isotropic")
	}
}
