package geometricmeanresonance

import (
	"math"
	"testing"
)

func TestBuildDefault(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault() error = %v", err)
	}
	if !a.Summary.GeometricMeanResonanceFound {
		t.Fatalf("expected geometric-mean resonance, got %+v", a.Geometric)
	}
	if a.Summary.NativeBreakingDerived || a.Summary.PatiSalamRouteOpened {
		t.Fatalf("unexpected native breaking/Pati-Salam route: %+v %+v", a.Seesaw, a.PatiSalam)
	}
	if a.Firewall.FiniteCorePolluted || a.Firewall.NewPhenomenologicalFitAdded || a.Firewall.PatiSalamImportedAsTheorem {
		t.Fatalf("firewall violation: %+v", a.Firewall)
	}
}

func TestGeometricMeanComputation(t *testing.T) {
	h := SealedHierarchy{MBGeV: heavyThresholdMBGeV, MStarGeV: topologicalBoundaryMStarGeV, VEVGeV: electroweakVEVGeV, FARequirementGeV: sealedAxionFAGeV, LambdaEFTMaxGeV: relicDecayLambdaMaxGeV, ValuesInheritedOnly: true}
	g := auditGeometricMean(h)
	expected := math.Sqrt(heavyThresholdMBGeV * topologicalBoundaryMStarGeV)
	if math.Abs(g.MIntGeV-expected)/expected > 1e-12 {
		t.Fatalf("unexpected geometric mean %.12e want %.12e", g.MIntGeV, expected)
	}
	if !g.FAMatch || !g.LambdaMatch || !g.BothTargetsBracketed {
		t.Fatalf("expected f_a and Lambda to bracket M_int within one decade: %+v", g)
	}
	if g.Targets[0].Log10Gap >= 1 || g.Targets[1].Log10Gap >= 1 {
		t.Fatalf("target gaps too large: %+v", g.Targets)
	}
}

func TestTheoremPasses(t *testing.T) {
	res := GeometricMeanIntermediateResonanceAuditTheorem().Run()
	if !res.Passed() {
		t.Fatalf("theorem failed:\n%s", res.Details())
	}
}
