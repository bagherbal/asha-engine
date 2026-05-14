package coarsegrain

import "testing"

func TestNativeCoarseGrainingDefault(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.FiniteBoundarySeedInherited || !a.ThresholdInventoryInherited {
		t.Fatalf("expected inherited boundary seed and threshold inventory")
	}
	if a.OperatorCount < 7 {
		t.Fatalf("expected at least seven candidate operators, got %d", a.OperatorCount)
	}
	if !a.ProjectionOperatorsFound || !a.SpectralAnchorsFound || !a.StaticClassifiersFound {
		t.Fatalf("expected projection, spectral, and classifier candidates")
	}
	if a.NativeCoarseGrainingFound || a.SemigroupLawDerived || a.ScaleLogParameterDerived || a.FlowFixedPointSelected {
		t.Fatalf("native RG semigroup, scale/log parameter, and fixed point must remain open")
	}
	if a.ThresholdActivationPredicateDerived || a.DecouplingMatchingRuleDerived || a.AbsoluteCouplingRunningDerived {
		t.Fatalf("threshold predicate, matching, and absolute coupling running must remain open")
	}
	if !a.NonUniqueActivationWitnessed {
		t.Fatalf("expected non-unique activation schedule witness")
	}
	if a.ResidualNullityAfter != a.ResidualNullityBefore || a.ResidualSymmetryBroken {
		t.Fatalf("Gate 105 should not break Gate 104 residual symmetry")
	}
	if a.PhysicalWeakAngleDerived || a.FineStructureDerived || a.PhysicalMassesDerived || a.HiddenObservedInputUsed {
		t.Fatalf("physical predictions must remain sealed")
	}
}
