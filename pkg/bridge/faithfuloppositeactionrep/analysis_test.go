package faithfuloppositeactionrep

import "testing"

func TestBuildDefaultExposesNonVacuousDiagnosticButNoPhysicalPromotion(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault returned error: %v", err)
	}
	if !a.Summary.CandidateOneFormsNonzero {
		t.Fatalf("expected non-vacuous diagnostic one-form: %s", FormatSummary(a.Summary))
	}
	if a.Summary.FullSCRepresentation || a.Summary.PhysicalOppositeAction || a.Summary.CandidateOrderOnePasses {
		t.Fatalf("Gate 270 must not promote candidate to physical spectral triple: %s", FormatSummary(a.Summary))
	}
}

func TestOneFormDiagnosticValues(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault returned error: %v", err)
	}
	want := [3]float64{-1, 1, 0}
	if a.OneForm.SpatialOneFormDiag != want {
		t.Fatalf("unexpected one-form diagonal: %s", FormatOneForm(a.OneForm))
	}
	if a.OneForm.FrobeniusNormSq != 2 || !a.OneForm.NonZero || !a.OneForm.CentralProbeVanishes {
		t.Fatalf("unexpected one-form norm/audit: %s", FormatOneForm(a.OneForm))
	}
}

func TestNaiveOppositeActionFailsOrderOneResidual(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault returned error: %v", err)
	}
	want := [3]float64{-1, 0, 0}
	if a.Residual.ResidualDiag != want {
		t.Fatalf("unexpected residual diagonal: %s", FormatResidual(a.Residual))
	}
	if a.Residual.FrobeniusNormSq != 1 || a.Residual.CandidatePasses || a.Residual.FullOrderOneProved {
		t.Fatalf("candidate should fail full order-one residual: %s", FormatResidual(a.Residual))
	}
}

func TestXYRatioStillUnconstrained(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault returned error: %v", err)
	}
	if len(a.Ratio.Rows) != 3 {
		t.Fatalf("expected three inherited ratio diagnostics, got %d", len(a.Ratio.Rows))
	}
	if a.Ratio.RatioStableAcrossFamily || !a.Ratio.DependsOnXY || a.Ratio.HiggsRatioDerived {
		t.Fatalf("ratio audit should remain blocked: %s", FormatRatio(a.Ratio))
	}
}

func TestFirewallAndTheorem(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault returned error: %v", err)
	}
	if !a.Firewall.CandidateNotPromoted || a.Firewall.FiniteCorePolluted || !a.Firewall.NoConnesRepresentationImported {
		t.Fatalf("firewall violation: %s", FormatFirewall(a.Firewall))
	}
	res := FaithfulOppositeActionRepresentationNonVacuousOneFormCalculusAuditTheorem().Run()
	if !res.Passed() {
		t.Fatalf("theorem did not pass:\n%s", res.Details())
	}
}
