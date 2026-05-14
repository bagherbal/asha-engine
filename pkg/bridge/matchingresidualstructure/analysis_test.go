package matchingresidualstructure

import "testing"

func TestGate216BuildDefault(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if a.Summary.Status != StatusFailedRoute {
		t.Fatalf("expected failed route: %s", FormatSummary(a.Summary))
	}
	if !a.Gate215.SingleScaleCandidateUnique || a.Gate215.PlausibleClasses != 1 || a.Gate215.ClassesAudited != 22 {
		t.Fatalf("bad Gate 215 inheritance: %s", FormatGate215(a.Gate215))
	}
	if a.Gate215.RequiredSignPattern != "-+-" || a.Gate215.RequiredMaxAbs <= 0 || a.Gate215.RequiredOverEnvelope >= 0.1 {
		t.Fatalf("unexpected target residual: %s", FormatGate215(a.Gate215))
	}
}

func TestSpectralDataAndSignResonanceRemainDiagnostic(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.SpectralData.BGapAvailable || a.SpectralData.ContactPartialModeCount != 7 || !a.SpectralData.ContactZetaLedgerAvailable || !a.SpectralData.FiniteFundamentalClassAvailable {
		t.Fatalf("bad spectral data audit: %s", FormatSpectralData(a.SpectralData))
	}
	if a.HeatKernelMap.SignOnlyResonances == 0 {
		t.Fatalf("expected at least one sign-only resonance: candidates=%s", FormatCandidates(a.TraceCandidates))
	}
	if a.HeatKernelMap.FullStructuralMatches != 0 || a.CoefficientSearch.ExactMagnitudeMatches != 0 {
		t.Fatalf("should not find canonical spectral match: heat=%s coef=%s", FormatHeatKernel(a.HeatKernelMap), FormatCoefficient(a.CoefficientSearch))
	}
}

func TestNoMatchingCorrectionOrHeatKernelFirewallLeak(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if a.SpectralData.FiniteMatchingRowsDerived || a.HeatKernelMap.DeltaMatchRowsDerived != 0 || a.HeatKernelMap.GaugeKineticTraceMapDerived || a.HeatKernelMap.ThresholdSubtractionSchemeDerived {
		t.Fatalf("heat-kernel/matching leak: spectral=%s heat=%s", FormatSpectralData(a.SpectralData), FormatHeatKernel(a.HeatKernelMap))
	}
	if a.Firewall.MatchingResidualPromoted || a.Firewall.MatchingCorrectionsDerived || a.Firewall.SpectralCoefficientTuned || a.Firewall.HeatKernelMapImported || a.Firewall.MSbarSchemeImported || a.Firewall.PhysicalMassPredictionClaimed || a.Firewall.PhysicalUnificationClaimed {
		t.Fatalf("firewall leak: %s", FormatFirewall(a.Firewall))
	}
}

func TestTheoremChecksPass(t *testing.T) {
	res := MatchingResidualStructureSpectralHeatKernelSearchTheorem().Verify()
	for _, c := range res.Checks {
		if !c.Passed {
			t.Fatalf("check %q failed: %s", c.Name, c.Detail)
		}
	}
}
