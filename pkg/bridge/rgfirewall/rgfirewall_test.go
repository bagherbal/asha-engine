package rgfirewall

import "testing"

func TestFiniteRGFirewallBuildsFormalFamilyWithoutPhysicalConstants(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.BoundaryDataDerived || !a.BetaDiagnosticAvailable {
		t.Fatalf("expected embedded boundary and beta diagnostic: %+v", a)
	}
	if !a.FormalRGFamilyConstructed || !a.TwoParameterUnderdetermined {
		t.Fatalf("expected symbolic underdetermined RG family")
	}
	if !a.NonUniquenessWitnessed {
		t.Fatalf("expected two non-physical samples to witness non-uniqueness")
	}
	if a.BoundaryCouplingDerived || a.BoundaryScaleDerived || a.ThresholdRuleDerived || a.FiniteRGTheoremDerived {
		t.Fatalf("firewall must not claim missing scale/coupling/threshold/RG theorems")
	}
	if a.PhysicalWeakAngleDerived || a.FineStructureDerived || a.PhysicalMassesDerived || a.HiddenObservedInputUsed {
		t.Fatalf("firewall must not derive or insert physical constants")
	}
}

func TestBoundaryNoRunningWitnessRecoversThreeEighths(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !close(a.SampleA.Sin2, 3.0/8.0, 1e-10) {
		t.Fatalf("boundary sample should recover finite diagnostic 3/8, got %.12f", a.SampleA.Sin2)
	}
	if close(a.SampleA.Sin2, a.SampleB.Sin2, 1e-10) {
		t.Fatalf("shifted log sample should differ until L is selected")
	}
	if a.SampleA.Physical || a.SampleB.Physical {
		t.Fatalf("samples are mathematical witnesses only")
	}
}
