package ewprojection

import (
	"math"
	"testing"
)

func TestElectroweakProjectionInvariants(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault: %v", err)
	}
	if math.Abs(a.ChargeIdentityResidual) > 1e-10 {
		t.Fatalf("charge identity residual too large: %.3e", a.ChargeIdentityResidual)
	}
	if math.Abs(a.TraceDirectionSin2-0.25) > 1e-10 {
		t.Fatalf("left trace sin2 = %.12f", a.TraceDirectionSin2)
	}
	if math.Abs(a.HyperchargeNormalizationKY-5.0/3.0) > 1e-10 {
		t.Fatalf("kY = %.12f", a.HyperchargeNormalizationKY)
	}
	if math.Abs(a.EqualNormalizedCouplingBoundarySin2-3.0/8.0) > 1e-10 {
		t.Fatalf("boundary sin2 = %.12f", a.EqualNormalizedCouplingBoundarySin2)
	}
	if a.WeakMixingAngleDerived {
		t.Fatalf("weak mixing angle must remain bridge-open")
	}
	if a.FineStructureDerived {
		t.Fatalf("fine structure must remain bridge-open")
	}
}
