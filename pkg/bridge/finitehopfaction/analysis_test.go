package finitehopfaction

import "testing"

func TestBuildDefault(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault failed: %v", err)
	}
	if !a.Summary.GeometricResonanceInherited {
		t.Fatalf("expected inherited geometric resonance: %s", FormatSummary(a.Summary))
	}
	if a.Summary.FiniteInstantonDerived || a.Summary.HopfActionMapDerived || a.Summary.HiddenOrderParameterDerived || a.Summary.IntermediateSealGranted {
		t.Fatalf("Gate 230 must not promote the route: %s", FormatSummary(a.Summary))
	}
	if a.HopfAction.Log10Gap > 0.02 {
		t.Fatalf("expected inherited tight Hopf resonance, got %.12g", a.HopfAction.Log10Gap)
	}
	if a.Firewall.FiniteCorePolluted || a.Firewall.HopfActionNormalizationFitted || a.Firewall.InstatonEquationInvented {
		t.Fatalf("firewall violation: %s", FormatFirewall(a.Firewall))
	}
}

func TestTheorem(t *testing.T) {
	res := OctonionicInstantonFiniteHopfActionMapAuditTheorem().Verify()
	if len(res.Checks) == 0 {
		t.Fatal("expected checks")
	}
	for _, c := range res.Checks {
		if !c.Passed {
			t.Fatalf("check %q failed: %s", c.Name, c.Detail)
		}
	}
}
