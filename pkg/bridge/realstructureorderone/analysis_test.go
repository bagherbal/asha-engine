package realstructureorderone

import "testing"

func TestBuildDefault(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault failed: %v", err)
	}
	if !a.Summary.CandidateJAvailable || !a.Summary.CandidateKOSignsComputed || !a.Summary.JRealityReducesParameters {
		t.Fatalf("expected candidate J/KO/J-reality preflight: %s", FormatSummary(a.Summary))
	}
	if a.Summary.OrderOneDerived || a.Summary.BGapMajoranaPlacement || a.Summary.CanonicalDFDerived {
		t.Fatalf("Gate 234 must not promote spectral triple data: %s", FormatSummary(a.Summary))
	}
	if a.KO.FreeParametersBefore != 64 || a.KO.FreeParametersAfterJ != 32 {
		t.Fatalf("unexpected J-reality parameter count: %s", FormatKO(a.KO))
	}
	if !a.BGap.BGapAvailable || a.BGap.BGap < 0.1 || a.BGap.BGap > 0.11 {
		t.Fatalf("unexpected B-gap audit: %s", FormatBGap(a.BGap))
	}
	if a.Firewall.FiniteCorePolluted || a.Firewall.BGapPromotedToMass || a.Firewall.DFChosenByFit || a.Firewall.OrderOneClaimed {
		t.Fatalf("firewall violation: %s", FormatFirewall(a.Firewall))
	}
}

func TestJConstruction(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault failed: %v", err)
	}
	if a.J.J2Residual != 0 || a.J.J2Sign != 1 || !a.J.CommutesWithOccupationGamma || a.J.JGammaSign != 1 {
		t.Fatalf("bad J candidate: %s", FormatJ(a.J))
	}
	if a.J.PhysicalChargeConjugation || a.J.ParticleAntiparticleDoubling || !a.J.CandidateOnly {
		t.Fatalf("J must remain candidate-only: %s", FormatJ(a.J))
	}
}

func TestOrderOneAndBGapRemainObstructed(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault failed: %v", err)
	}
	if a.OrderOne.OrderOneVerified || a.OrderOne.SplitsColorWeakSubblocks || a.OrderOne.PromotableFiniteDirac {
		t.Fatalf("order-one should remain obstructed: %s", FormatOrderOne(a.OrderOne))
	}
	if a.BGap.BGapCanonicalMajoranaEntry || a.BGap.BGapForcedToNeutralSector || a.BGap.BGapPromotedToMajoranaMass || !a.BGap.RequiresBroaderHilbertSpace {
		t.Fatalf("B-gap Majorana placement should remain obstructed: %s", FormatBGap(a.BGap))
	}
}

func TestTheorem(t *testing.T) {
	res := RealStructureKOOrderOneCalculusAuditTheorem().Verify()
	if len(res.Checks) == 0 {
		t.Fatal("expected checks")
	}
	for _, c := range res.Checks {
		if !c.Passed {
			t.Fatalf("check %q failed: %s", c.Name, c.Detail)
		}
	}
}
