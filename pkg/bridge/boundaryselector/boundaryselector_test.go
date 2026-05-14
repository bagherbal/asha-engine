package boundaryselector

import "testing"

func TestBoundarySelectorSearchFindsOnlyDimensionlessCandidates(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.FiniteBoundarySeedSelected || !a.RelativeKineticNormalizationComplete {
		t.Fatalf("expected inherited finite boundary seed")
	}
	if a.CandidateCount < 8 {
		t.Fatalf("expected a broad candidate operator inventory, got %d", a.CandidateCount)
	}
	if !a.AllCandidateOperatorsDimensionless || a.DimensionfulOperatorFound {
		t.Fatalf("all current candidates should be dimensionless")
	}
	if a.AbsoluteCouplingOperatorFound || a.BoundaryScaleOperatorFound || a.ThresholdSelectorFound {
		t.Fatalf("Gate 104 must not invent coupling/scale/threshold selectors")
	}
}

func TestResidualNullityKeepsPhysicalPredictionsSealed(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if a.EquationAudit.Nullity < 3 || a.EquationAudit.IndependentEquationsForPhysicalFlow != 0 {
		t.Fatalf("expected at least three residual missing variables and zero finite selector equations: %+v", a.EquationAudit)
	}
	if !a.UnitTraceConventionRejected || !a.TopologicalSealAsScaleRejected {
		t.Fatalf("expected unit-trace and topological-seal-as-scale rejections")
	}
	if a.BoundaryCouplingDerived || a.BoundaryScaleDerived || a.FiniteRGTheoremDerived || a.ThresholdRuleDerived {
		t.Fatalf("missing physical bridges must remain unproved")
	}
	if a.PhysicalWeakAngleDerived || a.FineStructureDerived || a.PhysicalMassesDerived || a.HiddenObservedInputUsed {
		t.Fatalf("physical predictions must remain sealed")
	}
}
