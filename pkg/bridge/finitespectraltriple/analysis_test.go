package finitespectraltriple

import "testing"

func TestGate217BuildDefault(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if a.Summary.Status != StatusFailedRoute {
		t.Fatalf("expected failed route: %s", FormatSummary(a.Summary))
	}
	if !a.Gate216.Gate216Inherited || !a.Gate216.Gate215SingleScaleTargetInherited || !a.Gate216.ThresholdSpectrumSealInherited || a.Gate216.FiniteMatchingRowsDerived {
		t.Fatalf("bad Gate 216 inheritance: %s", FormatGate216(a.Gate216))
	}
	if a.Gate216.ResidualSignPattern != "-+-" || a.Gate216.SpectralResidualTarget.MaxAbs() <= 0 {
		t.Fatalf("bad residual target: %s", FormatGate216(a.Gate216))
	}
}

func TestHeavySectorIsNotFiniteSpectralHilbert(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if len(a.Representations) != 2 || a.Hilbert.InternalDimensionTotal != 19 || a.Hilbert.DiracChiralCarrierDimTotal != 38 {
		t.Fatalf("unexpected heavy reps: %s", FormatRepresentations(a.Representations))
	}
	if a.Hilbert.FiniteHilbertSpaceDerived || a.Hilbert.RealStructureDerived || a.Hilbert.GradingDerived || a.Hilbert.HeavyChargeConjugation {
		t.Fatalf("heavy Hilbert semantics should be missing: %s", FormatHilbert(a.Hilbert))
	}
}

func TestNoDiracHeatKernelOrCutoffPromotion(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if a.DiracAudit.PromotableFiniteDirac != 0 || a.DiracAudit.OrderOneVerified != 0 || a.DiracAudit.CliffordG2Dictated != 0 || a.DiracAudit.MassScaleFiniteDerived != 0 {
		t.Fatalf("Dirac operator leaked: %s", FormatDiracAudit(a.DiracAudit))
	}
	if a.HeatKernel.GaugeCurvatureProjectionRowsDerived != 0 || a.HeatKernel.A4GaugeCoefficientsDerived != 0 || a.HeatKernel.ProjectedDeltaMatchRows != 0 || a.HeatKernel.TrDFMinus2Promoted {
		t.Fatalf("heat-kernel projection leaked: %s", FormatHeatKernel(a.HeatKernel))
	}
	if a.Cutoff.CutoffFunctionDerived || a.Cutoff.ThresholdSubtractionRuleDerived || a.Cutoff.MSbarImported || a.Cutoff.PhysicalDeltaMatchRows != 0 {
		t.Fatalf("cutoff/subtraction leak: %s", FormatCutoff(a.Cutoff))
	}
}

func TestFirewallsAndTheoremChecksPass(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if a.Firewall.DFFittedByHand || a.Firewall.CutoffFunctionInvented || a.Firewall.HeatKernelProjectionFitted || a.Firewall.MatchingResidualPromoted || a.Firewall.MatchingCorrectionsDerived || a.Firewall.PhysicalUnificationClaimed {
		t.Fatalf("firewall leak: %s", FormatFirewall(a.Firewall))
	}
	res := FiniteSpectralTripleHeavySectorGaugeCurvatureProjectionAuditTheorem().Verify()
	for _, c := range res.Checks {
		if !c.Passed {
			t.Fatalf("check %q failed: %s", c.Name, c.Detail)
		}
	}
}
