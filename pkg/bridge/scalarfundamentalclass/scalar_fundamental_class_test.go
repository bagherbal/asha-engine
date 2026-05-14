package scalarfundamentalclass

import "testing"

func TestBuildDefaultFiniteFundamentalClass(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault failed: %v", err)
	}
	if !a.Functional.FiniteMatrixFunctionalConstructed || a.Functional.ContinuumIntegralImported || a.Functional.DixmierTraceDerived {
		t.Fatalf("functional classification wrong: %s", FormatFunctional(a.Functional))
	}
	if !a.ClosedCyclic.HochschildBoundaryZeroOnAuditedDomain || a.ClosedCyclic.EtaTraceCyclicOnFullMatrixAlgebra {
		t.Fatalf("closed/cyclic audit wrong: %s", FormatClosedCyclic(a.ClosedCyclic))
	}
	if a.ClosedCyclic.EtaTraceFullMatrixCounterexampleDefect != 2 {
		t.Fatalf("expected full-matrix eta counterexample defect 2, got %g", a.ClosedCyclic.EtaTraceFullMatrixCounterexampleDefect)
	}
	if !a.Normalization.StableQuantizedInvariants || a.Normalization.NeutralQNativeDegree != 2 || a.Normalization.NeutralZNativeDegree != -2 || a.Normalization.NeutralMixedNativeDegree != 1 {
		t.Fatalf("normalization audit wrong: %s", FormatNormalization(a.Normalization))
	}
	if a.Normalization.CanonicalNormalizationFactorDerived || a.Normalization.UnitFundamentalClassDerived {
		t.Fatalf("normalization should not be forced: %s", FormatNormalization(a.Normalization))
	}
	if a.Firewall.ImportsTopologicalSeal8PiSquared || a.Firewall.AbsoluteCouplingPromoted || a.Firewall.PhysicalConstantsDerived || a.Firewall.HeatKernelA4CoefficientPromoted {
		t.Fatalf("firewall leaked: %s", FormatFirewall(a.Firewall))
	}
}

func TestMatterTensorLiftPlanRemainsSupportOnly(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if a.MatterPlan.MatterFockDimension != 16 || a.MatterPlan.ScalarBundleDimension != 4 || a.MatterPlan.TotalTensorDimension != 64 {
		t.Fatalf("unexpected tensor dimensions: %s", FormatMatterPlan(a.MatterPlan))
	}
	if !a.MatterPlan.SelectionRulesCanBeReused || !a.MatterPlan.RequiresSeparateGate {
		t.Fatalf("matter plan should be a separate support audit: %s", FormatMatterPlan(a.MatterPlan))
	}
	if a.MatterPlan.YukawaAmplitudesDerived || a.MatterPlan.MassTermsDerived {
		t.Fatalf("matter plan leaked amplitudes/masses: %s", FormatMatterPlan(a.MatterPlan))
	}
}

func TestTheoremChecksPass(t *testing.T) {
	res := FiniteFundamentalClassScalarBundleIntegrationFunctionalSearchAuditTheorem().Verify()
	for _, c := range res.Checks {
		if !c.Passed {
			t.Fatalf("check %q failed: %s", c.Name, c.Detail)
		}
	}
}
