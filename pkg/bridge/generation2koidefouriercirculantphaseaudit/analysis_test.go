package generation2koidefouriercirculantphaseaudit

import (
	"math"
	"testing"
)

func TestGate582KoideFourierCirculantPhaseAudit(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if math.Abs(a.MZ.DeltaDeg-132.73281996710833) > 1e-10 {
		t.Fatalf("unexpected M_Z Fourier phase: %s", FormatPoint(a.MZ))
	}
	if math.Abs(a.Lambda12.DeltaDeg-132.73261746845495) > 1e-10 {
		t.Fatalf("unexpected Lambda_12 Fourier phase: %s", FormatPoint(a.Lambda12))
	}
	if math.Abs(a.MZ.PlaneAmplitudeR-0.9999907671734557) > 1e-12 {
		t.Fatalf("unexpected M_Z Fourier amplitude: %s", FormatPoint(a.MZ))
	}
	if !(a.Transport.PhaseStable && a.Transport.AmplitudeMovesTowardOne) {
		t.Fatalf("phase should be stable and amplitude should move toward Koide: %s", FormatTransport(a.Transport))
	}
	if !(a.MZ.MaxReconstructionError < 1e-15 && a.Lambda12.MaxReconstructionError < 1e-15) {
		t.Fatalf("circulant reconstruction should be exact: %s / %s", FormatPoint(a.MZ), FormatPoint(a.Lambda12))
	}
	if a.Permutation.UniqueWithoutOrdering || a.Permutation.SimplePhaseCertified {
		t.Fatalf("permutation/simple phase firewall broken: %s", FormatPermutation(a.Permutation))
	}
	if !(a.Permutation.BestResidualDeg > a.Permutation.CertificationDeg) {
		t.Fatalf("best rational should fail certification: %s", FormatPermutation(a.Permutation))
	}
	if a.Firewalls.DerivesFourierPhase || a.Firewalls.AddsNewCarrier || a.Firewalls.PromotesObservedAsNative || !a.Firewalls.PreservesGate352 {
		t.Fatalf("firewall broken: %s", FormatFirewalls(a.Firewalls))
	}
}

func TestGate582Theorem(t *testing.T) {
	res := Generation2KoideFourierCirculantPhaseAuditTheorem().Run()
	if res.Status == "FAILED_ROUTE" {
		t.Fatalf("unexpected failed route: %+v", res)
	}
	for _, c := range res.Checks {
		if !c.Passed {
			t.Fatalf("check failed: %+v", c)
		}
	}
}
