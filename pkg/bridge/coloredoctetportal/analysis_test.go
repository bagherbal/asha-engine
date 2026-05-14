package coloredoctetportal

import "testing"

func TestGate223FindsDimensionSixOctetPortal(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault: %v", err)
	}
	if !a.TensorSearch.OctetPortalFound {
		t.Fatalf("expected octet portal")
	}
	if a.TensorSearch.FalsifiesSpectrum {
		t.Fatalf("spectrum should not be falsified after portal discovery")
	}
	best := a.TensorSearch.BestPortal
	if !best.ValidPortal {
		t.Fatalf("best portal invalid: %+v", best)
	}
	if best.TotalDimHalf != 12 {
		t.Fatalf("expected total dimension 6, got half=%d", best.TotalDimHalf)
	}
	if !best.BaryonSafe {
		t.Fatalf("best portal should preserve baryon firewall: %+v", best)
	}
	if !hasChromomagneticPortal(a.TensorSearch.ValidPortals) {
		t.Fatalf("expected the scan to include the G H† e^c chromomagnetic portal among valid candidates")
	}
	if !a.Kinematics.BBNSafeForPerturbativeWilson {
		t.Fatalf("expected BBN-safe EFT range")
	}
	if !a.RelicSeal.SealGranted {
		t.Fatalf("expected conditional RelicDecaySeal")
	}
	if a.Firewall.FiniteOperatorClaimed || a.Firewall.WilsonCoefficientFixed || a.Firewall.FiniteCorePolluted {
		t.Fatalf("firewall polluted: %+v", a.Firewall)
	}
}

func hasChromomagneticPortal(xs []SearchCombination) bool {
	for _, x := range xs {
		if containsName(x.Fields, "G_{μν}") && containsName(x.Fields, "H†") && containsName(x.Fields, "e^c") {
			return true
		}
	}
	return false
}

func containsName(xs []string, s string) bool {
	for _, x := range xs {
		if x == s {
			return true
		}
	}
	return false
}
