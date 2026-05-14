package intermediatebreakingaudit

import (
	"math"
	"testing"
)

func TestBuildDefaultGate228(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault() error = %v", err)
	}
	if !a.PatiSalam.CatastrophicFailure {
		t.Fatalf("expected Pati-Salam intermediate breaking to fail")
	}
	if a.PatiSalam.LifetimeYears >= superKBoundYears {
		t.Fatalf("Pati-Salam lifetime unexpectedly safe: %e", a.PatiSalam.LifetimeYears)
	}
	if !a.BGap.RequiredCOrderOne || !a.BGap.NonPerturbativeShapeWorks {
		t.Fatalf("expected B-gap exponential shape to reach M_int with an order-one required c")
	}
	if a.BGap.NativeCoefficientDerived || a.Seal.SealGranted {
		t.Fatalf("coefficient/IntermediateBreakingSeal should not be derived/granted")
	}
	if a.Firewall.FiniteCorePolluted || a.Firewall.PatiSalamUnsealedForDynamics {
		t.Fatalf("firewall was violated: %+v", a.Firewall)
	}
}

func TestRequiredCoefficientReconstructsMInt(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault() error = %v", err)
	}
	recon := a.BGap.MStarGeV * math.Exp(-a.BGap.RequiredC/a.BGap.BGap)
	if math.Abs(math.Log(recon/a.BGap.MIntTargetGeV)) > 1e-10 {
		t.Fatalf("required c does not reconstruct M_int: got %e want %e", recon, a.BGap.MIntTargetGeV)
	}
	if a.BGap.CandidateFourOverPi.Log10Gap > 0.05 {
		t.Fatalf("expected 4/pi diagnostic to be a close near resonance, gap=%f", a.BGap.CandidateFourOverPi.Log10Gap)
	}
}

func TestTheoremChecksPass(t *testing.T) {
	th := PatiSalamFalsificationBSectorHierarchyAuditTheorem()
	res := th.Verify()
	for _, c := range res.Checks {
		if !c.Passed {
			t.Fatalf("check %q failed: %s", c.Name, c.Detail)
		}
	}
}
