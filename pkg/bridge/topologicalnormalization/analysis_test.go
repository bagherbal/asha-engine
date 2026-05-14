package topologicalnormalization

import "testing"

func TestBuildDefaultTopologicalNormalization(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Input.GaugeRatioClosed || !a.Input.WeakAngleSeedClosed || !a.Input.MassGenerationSealed {
		t.Fatalf("expected inherited closed ratio, weak seed, and sealed mass route: %+v", a.Input)
	}
	if !a.Input.TopologicalSealAvailable || a.Input.ContactIndex != 1 {
		t.Fatalf("expected unit contact index/topological seal input: %+v", a.Input)
	}
}

func TestConditionalInstantonMatching(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Matching.ConditionalMatchingAvailable {
		t.Fatalf("expected conditional instanton branch: %+v", a.Matching)
	}
	if !close(a.Matching.ConditionalUInverseGStar, 1, 1e-12) || !close(a.Matching.ConditionalGStarSquared, 1, 1e-12) {
		t.Fatalf("bad conditional coupling: %+v", a.Matching)
	}
	if a.Matching.CanonicalStrictMatchingDerived || a.Matching.TopologicalSealAloneSufficient {
		t.Fatalf("strict matching should not be derived by the seal alone: %+v", a.Matching)
	}
}

func TestF0ConventionDependence(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if len(a.Conventions) != 2 {
		t.Fatalf("expected two prefactor conventions")
	}
	if !close(a.Conventions[0].ConditionalF0, 0.5, 1e-12) {
		t.Fatalf("single-trace f0=%g, want 1/2", a.Conventions[0].ConditionalF0)
	}
	if !close(a.Conventions[1].ConditionalF0, 0.25, 1e-12) {
		t.Fatalf("two-sided f0=%g, want 1/4", a.Conventions[1].ConditionalF0)
	}
	for _, c := range a.Conventions {
		if !c.SameBoundaryPhysics || !close(c.ConditionalSin2, 3.0/8.0, 1e-12) {
			t.Fatalf("convention should preserve boundary physics: %+v", c)
		}
	}
}

func TestStrictNullityRemainsThree(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if a.Firewall.StrictNullityBefore != 3 || a.Firewall.StrictNullityAfter != 3 {
		t.Fatalf("strict nullity should remain 3: %+v", a.Firewall)
	}
	if a.Firewall.ConditionalNullityAfter != 2 {
		t.Fatalf("conditional nullity should be 2: %+v", a.Firewall)
	}
	if a.Firewall.StrictAbsoluteUDerived || a.Firewall.StrictF0Derived || a.Firewall.PhysicalCouplingsDerived {
		t.Fatalf("strict physical coupling should remain sealed: %+v", a.Firewall)
	}
}
