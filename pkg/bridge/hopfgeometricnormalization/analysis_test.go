package hopfgeometricnormalization

import (
	"math"
	"testing"
)

func TestBuildDefault(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault failed: %v", err)
	}
	if !a.Summary.GeometricHierarchySupported {
		t.Fatalf("expected conditional geometric hierarchy support: %s", FormatSummary(a.Summary))
	}
	if a.Summary.NativeHopfMapDerived || a.Summary.IntermediateSealGranted {
		t.Fatalf("native Hopf map / seal must remain ungranted: %s", FormatSummary(a.Summary))
	}
	if math.Abs(a.Geometry.Coefficient-4/math.Pi) > 1e-14 {
		t.Fatalf("coefficient is not 4/pi: %.17g", a.Geometry.Coefficient)
	}
	if a.Hierarchy.Log10Gap > 0.02 {
		t.Fatalf("expected tight near-resonance, got %.12g", a.Hierarchy.Log10Gap)
	}
	if a.Sensitivity.OnePercentBGapShiftDecades < 0.05 || a.Sensitivity.OnePercentBGapShiftDecades > 0.06 {
		t.Fatalf("unexpected 1%% B-gap sensitivity: %.12g", a.Sensitivity.OnePercentBGapShiftDecades)
	}
	if !a.Residual.HigherLoopOrMatchingCanPlausiblyCover || a.Residual.FiniteResolutionDerived {
		t.Fatalf("residual should be plausible but not derived: %s", FormatResidual(a.Residual))
	}
	if a.Firewall.FiniteCorePolluted || a.Firewall.CoefficientFitted || a.Firewall.CoefficientDerivedFromFiniteCore {
		t.Fatalf("firewall violation: %s", FormatFirewall(a.Firewall))
	}
}

func TestTheorem(t *testing.T) {
	res := HopfFibrationGeometricNormalizationBGapSensitivityAuditTheorem().Verify()
	if len(res.Checks) == 0 {
		t.Fatal("expected checks")
	}
	for _, c := range res.Checks {
		if !c.Passed {
			t.Fatalf("check %q failed: %s", c.Name, c.Detail)
		}
	}
}
