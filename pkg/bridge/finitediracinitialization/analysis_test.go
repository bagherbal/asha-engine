package finitediracinitialization

import "testing"

func TestBuildDefault(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault failed: %v", err)
	}
	if !a.Summary.DFAnsatzAvailable {
		t.Fatalf("expected D_F ansatz availability: %s", FormatSummary(a.Summary))
	}
	if a.Summary.CanonicalDFDerived || a.Summary.BGapEmbeddingDerived || a.Summary.SpectralActionReady {
		t.Fatalf("Gate 233 must not promote a physical finite Dirac operator: %s", FormatSummary(a.Summary))
	}
	if a.DiracFamily.FreeRealParameters != 64 {
		t.Fatalf("expected 64 free real entries in M, got %d", a.DiracFamily.FreeRealParameters)
	}
	if a.UnitMatrix.AnticommutatorNorm > 1e-10 || a.UnitMatrix.SelfAdjointResidual > 1e-10 {
		t.Fatalf("unit representative failed matrix identity audit: %s", FormatRepresentative(a.UnitMatrix))
	}
	if a.BGap.BGap <= 0.1 || a.BGap.BGap >= 0.11 {
		t.Fatalf("unexpected B-gap: %s", FormatBGap(a.BGap))
	}
	if a.Firewall.FiniteCorePolluted || a.Firewall.DFChosenByFit || a.Firewall.BGapPromotedToMass {
		t.Fatalf("firewall violation: %s", FormatFirewall(a.Firewall))
	}
}

func TestTheorem(t *testing.T) {
	res := FiniteDiracOperatorInitializationFockMatrixAuditTheorem().Verify()
	if len(res.Checks) == 0 {
		t.Fatal("expected checks")
	}
	for _, c := range res.Checks {
		if !c.Passed {
			t.Fatalf("check %q failed: %s", c.Name, c.Detail)
		}
	}
}
