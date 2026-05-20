package historytransport

import (
	"math"
	"testing"
)

func TestBuildDefaultEndVectorAndBoundary(t *testing.T) {
	b, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if math.Abs(b.EndVector.VGeV-246.21965079413738) > 1e-9 {
		t.Fatalf("v mismatch: %.15g", b.EndVector.VGeV)
	}
	if math.Abs(b.EndVector.Sin2Theta-0.22337664470480512) > 1e-12 {
		t.Fatalf("sin2 mismatch: %.15g", b.EndVector.Sin2Theta)
	}
	if math.Abs(b.WeakAngleTransport.Sin2ThetaBoundary-0.375) > 1e-15 {
		t.Fatalf("boundary weak angle mismatch")
	}
	if !b.WeakAngleTransport.TransportRequired {
		t.Fatalf("weak angle transport should be required")
	}
	if b.GaugeBoundary.Lambda12GeV < 9e13 || b.GaugeBoundary.Lambda12GeV > 1.1e14 {
		t.Fatalf("unexpected Lambda12: %.15g", b.GaugeBoundary.Lambda12GeV)
	}
	if math.Abs(b.GaugeBoundary.G1Lambda-b.GaugeBoundary.G2Lambda) > 1e-12 {
		t.Fatalf("g1/g2 crossing failed: %.15g %.15g", b.GaugeBoundary.G1Lambda, b.GaugeBoundary.G2Lambda)
	}
	if b.GaugeBoundary.Interpretation != "threshold_needed" {
		t.Fatalf("expected threshold_needed, got %q", b.GaugeBoundary.Interpretation)
	}
	if b.GaugeBoundary.Delta3 >= 0 {
		t.Fatalf("expected negative strong inverse mismatch, got %.15g", b.GaugeBoundary.Delta3)
	}
}

func TestFlavorInvariantsAndFirewalls(t *testing.T) {
	b, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if b.FlavorTransport.JCKM < 3.0e-5 || b.FlavorTransport.JCKM > 3.3e-5 {
		t.Fatalf("unexpected J_CKM: %.15g", b.FlavorTransport.JCKM)
	}
	if math.Abs(b.FlavorTransport.KoideQe-2.0/3.0) > 1e-4 {
		t.Fatalf("Koide proxy drifted: %.15g", b.FlavorTransport.KoideQe)
	}
	if b.HistoryResidual.Interpretation == "" {
		t.Fatalf("missing residual interpretation")
	}
	for _, s := range []string{StatusNoNativeDerivationClaim, StatusNoPhysicalUnificationClaim, StatusThresholdsNotHidden} {
		if !contains(b.Statuses, s) {
			t.Fatalf("missing status %s", s)
		}
	}
}

func TestScalarTransportProducesTypedOutput(t *testing.T) {
	b, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if b.ScalarTransport.LambdaMZ <= 0 {
		t.Fatalf("lambda(MZ) must be positive")
	}
	if b.ScalarTransport.YT_MZ <= 0 {
		t.Fatalf("yt(MZ) must be positive")
	}
	if b.ScalarTransport.Approximation == "" {
		t.Fatalf("scalar approximation must be recorded")
	}
	if b.ScalarTransport.ZeroCrossingScaleGeV == nil {
		t.Fatalf("v1 approximation should expose lambda zero crossing before Lambda12")
	}
}

func contains(xs []string, target string) bool {
	for _, x := range xs {
		if x == target {
			return true
		}
	}
	return false
}
