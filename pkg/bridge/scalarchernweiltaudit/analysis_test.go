package scalarchernweiltaudit

import "testing"

func TestBuildDefaultSealedScalarBundleChernWeilPreflight(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault failed: %v", err)
	}
	if !a.Curvature.PrimitiveGaugeKineticTracesStable || !a.Curvature.NeutralQTraceStable || !a.Curvature.NeutralZTraceStable {
		t.Fatalf("curvature traces not stable: %s", FormatCurvature(a.Curvature))
	}
	if !a.Grading.EtaDerivedFromSeal || a.Grading.EtaSquaredResidual != 0 || a.Grading.EtaTrace != 0 {
		t.Fatalf("eta grading invalid: %s", FormatGrading(a.Grading))
	}
	if !a.Grading.PrimitiveDiagonalGradedTracesZero {
		t.Fatalf("primitive diagonal eta traces should vanish: %s", FormatGrading(a.Grading))
	}
	if !a.Grading.NontrivialSignedNeutralCarrier || a.Grading.NeutralMixedGradedTrace != 1 || a.Grading.NeutralSplitQGradedTrace != 2 || a.Grading.NeutralSplitZGradedTrace != -2 {
		t.Fatalf("expected nontrivial neutral signed carrier: %s", FormatGrading(a.Grading))
	}
	if !a.HeatKernel.A4LocalAlgebraicIngredientPresent || a.HeatKernel.SpectralActionEvaluated || a.HeatKernel.HeatKernelCoefficientPromoted {
		t.Fatalf("heat-kernel preflight classified incorrectly: %s", FormatHeatKernel(a.HeatKernel))
	}
	if a.Firewall.ImportsTopologicalSeal8PiSquared || a.Firewall.EquatesFiniteTraceWithInstanton || a.Firewall.AbsoluteCouplingPromoted || a.Firewall.PhysicalConstantsDerived {
		t.Fatalf("firewall leaked: %s", FormatFirewall(a.Firewall))
	}
}

func TestPrimitiveTraceValues(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	for _, r := range a.Curvature.GeneratorsAudited {
		switch r.Name {
		case "T1", "T2", "T3L", "Y_phi":
			if r.TotalTrace != 1 || r.HighFiberTrace != 0.5 || r.LowFiberTrace != 0.5 || r.EtaGradedTrace != 0 {
				t.Fatalf("unexpected primitive trace record for %s: %+v", r.Name, r)
			}
		case "Q=T3+Y_phi":
			if r.TotalTrace != 2 || r.HighFiberTrace != 2 || r.LowFiberTrace != 0 || r.EtaGradedTrace != 2 {
				t.Fatalf("unexpected Q trace record: %+v", r)
			}
		case "Z=T3-Y_phi":
			if r.TotalTrace != 2 || r.HighFiberTrace != 0 || r.LowFiberTrace != 2 || r.EtaGradedTrace != -2 {
				t.Fatalf("unexpected Z trace record: %+v", r)
			}
		}
	}
}

func TestTheoremChecksPass(t *testing.T) {
	res := SealedScalarBundleChernWeilCarrierHeatKernelPreflightTheorem().Verify()
	for _, c := range res.Checks {
		if !c.Passed {
			t.Fatalf("check %q failed: %s", c.Name, c.Detail)
		}
	}
}
